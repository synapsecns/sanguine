package relayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	messagetransmitter "github.com/synapsecns/sanguine/services/cctp-relayer/contracts/messagetransmitter"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type circleCCTPHandler struct {
	cfg              config.Config
	db               db2.CCTPRelayerDB
	omniRPCClient    omniClient.RPCClient
	boundCircleCCTPs map[uint32]*messagetransmitter.MessageTransmitter
	txSubmitter      submitter.TransactionSubmitter
	relayerAddress   common.Address
	handler          metrics.Handler
}

// NewCircleCCTPHandler creates a new CircleCCTPHandler.
func NewCircleCCTPHandler(ctx context.Context, cfg config.Config, db db2.CCTPRelayerDB, omniRPCClient omniClient.RPCClient, txSubmitter submitter.TransactionSubmitter, handler metrics.Handler) (CCTPHandler, error) {
	boundCircleCCTPs := make(map[uint32]*messagetransmitter.MessageTransmitter)
	for _, chain := range cfg.Chains {
		cl, err := omniRPCClient.GetConfirmationsClient(ctx, int(chain.ChainID), 1)
		if err != nil {
			return nil, fmt.Errorf("could not get client: %w", err)
		}
		boundCircleCCTPs[chain.ChainID], err = messagetransmitter.NewMessageTransmitter(chain.GetCircleCCTPAddress(), cl)
		if err != nil {
			return nil, fmt.Errorf("could not build bound contract: %w", err)
		}
	}
	signer, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not make cctp signer: %w", err)
	}
	return &circleCCTPHandler{
		cfg:              cfg,
		db:               db,
		omniRPCClient:    omniRPCClient,
		boundCircleCCTPs: boundCircleCCTPs,
		txSubmitter:      txSubmitter,
		relayerAddress:   signer.Address(),
		handler:          handler,
	}, nil
}

func (c *circleCCTPHandler) HandleLog(ctx context.Context, log *types.Log, chainID uint32) (processQueue bool, err error) {
	if log == nil {
		return false, fmt.Errorf("log is nil")
	}

	// shouldn't be possible: maybe remove?
	if len(log.Topics) == 0 {
		return false, fmt.Errorf("not enough topics")
	}

	switch log.Topics[0] {
	case messagetransmitter.MessageSentTopic:
		msg, err := c.handleMessageSent(ctx, log, chainID)
		if err != nil {
			return false, fmt.Errorf("could not store message sent: %w", err)
		}
		if msg != nil {
			processQueue = true
		}
		return processQueue, nil
	case messagetransmitter.MessageReceivedTopic:
		err = c.handleMessageReceived(ctx, log, chainID)
		if err != nil {
			return false, fmt.Errorf("could not handle message received: %w", err)
		}
		return false, nil
	default:
		logger.Warnf("unknown topic %s", log.Topics[0])
		return false, nil
	}
}

func (c *circleCCTPHandler) FetchAndProcessSentEvent(parentCtx context.Context, txHash common.Hash, chainID uint32) (msg *relayTypes.Message, err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "FetchAndProcessSentEvent", trace.WithAttributes(
		attribute.String(metrics.TxHash, txHash.String()),
		attribute.Int(metrics.ChainID, int(chainID)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// check if message already exists before we do anything
	msg, err = c.db.GetMessageByOriginHash(ctx, txHash)
	// if we already have the message, we can just return it
	if err == nil {
		return msg, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("could not get message by origin hash: %w", err)
	}

	ethClient, err := c.omniRPCClient.GetConfirmationsClient(ctx, int(chainID), 1)
	if err != nil {
		return nil, fmt.Errorf("could not get chain client: %w", err)
	}

	// TODO: consider pulling from scribe?
	// TODO: this function is particularly prone to silent errors
	receipt, err := ethClient.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction receipt: %w", err)
	}

	for _, log := range receipt.Logs {
		// this should never happen
		if len(log.Topics) == 0 {
			continue
		}

		if log.Topics[0] == messagetransmitter.MessageSentTopic {
			msg, err = c.handleMessageSent(ctx, log, chainID)
			if err != nil {
				return nil, fmt.Errorf("could not handle message sent: %w", err)
			}
			return msg, nil
		}
	}
	return nil, fmt.Errorf("no message sent event found")
}

func (c *circleCCTPHandler) SubmitReceiveMessage(parentCtx context.Context, msg *relayTypes.Message) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "SubmitReceiveMessage", trace.WithAttributes(
		attribute.String("message_hash", msg.MessageHash),
		attribute.Int(metrics.ChainID, int(msg.DestChainID)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	contract, ok := c.boundCircleCCTPs[msg.DestChainID]
	if !ok {
		return fmt.Errorf("no contract found for chain %d", msg.DestChainID)
	}

	//TODO: check if already executed on destination?

	var nonce uint64
	var destTxHash common.Hash
	nonce, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(msg.DestChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = contract.ReceiveMessage(transactor, msg.Message, msg.Attestation)
		if err != nil {
			return nil, fmt.Errorf("could not submit transaction: %w", err)
		}

		destTxHash = tx.Hash()
		return tx, nil
	})
	if err != nil {
		err = fmt.Errorf("could not submit transaction: %w", err)
		return err
	}
	span.SetAttributes(
		attribute.String("dest_tx_hash", destTxHash.String()),
		attribute.Int("nonce", int(nonce)),
	)

	// Store the completed message.
	// Note: this can cause double submission sometimes
	msg.State = relayTypes.Submitted
	msg.DestNonce = int(nonce)
	msg.DestTxHash = destTxHash.String()
	err = c.db.StoreMessage(ctx, *msg)
	if err != nil {
		return fmt.Errorf("could not store completed message: %w", err)
	}

	return nil
}

func (c *circleCCTPHandler) handleMessageSent(parentCtx context.Context, log *types.Log, chainID uint32) (msg *relayTypes.Message, err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "handleMessageSent", trace.WithAttributes(
		attribute.String(metrics.TxHash, log.TxHash.Hex()),
		attribute.Int(metrics.ChainID, int(chainID)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	ethClient, err := c.omniRPCClient.GetConfirmationsClient(ctx, int(chainID), 1)
	if err != nil {
		return nil, fmt.Errorf("could not get chain client: %w", err)
	}

	eventParser, err := messagetransmitter.NewMessageTransmitterFilterer(log.Address, ethClient)
	if err != nil {
		return nil, fmt.Errorf("could not create event parser: %w", err)
	}

	messageSentEvent, err := eventParser.ParseMessageSent(*log)
	if err != nil {
		return nil, fmt.Errorf("could not parse message sent: %w", err)
	}

	// check that we sent the tx
	tx, _, err := ethClient.TransactionByHash(ctx, log.TxHash)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction by hash: %w", err)
	}
	call, err := util.TxToCall(tx)
	if err != nil {
		return nil, fmt.Errorf("could not convert transaction to call: %w", err)
	}
	if call.From != c.relayerAddress {
		span.AddEvent(fmt.Sprintf("sender %s does not match relayer address %s", call.From.String(), c.relayerAddress.String()))
		//nolint:nilnil
		return nil, nil
	}

	destDomain, err := parseDestDomain(messageSentEvent.Message)
	if err != nil {
		return nil, fmt.Errorf("could not parse destination chain from raw message")
	}
	span.SetAttributes(attribute.Int("dest_domain", int(destDomain)))
	destChainID, err := CircleDomainToChainID(destDomain, IsTestnetChainID(chainID))
	if err != nil {
		return nil, fmt.Errorf("could not convert circle domain to chain ID: %w", err)
	}

	rawMsg := relayTypes.Message{
		OriginTxHash:  log.TxHash.Hex(),
		OriginChainID: chainID,
		DestChainID:   destChainID,
		Message:       messageSentEvent.Message,
		MessageHash:   crypto.Keccak256Hash(messageSentEvent.Message).String(),

		//Attestation: //comes from the api
		BlockNumber: log.BlockNumber,
	}

	span.SetAttributes(
		attribute.String("message_hash", rawMsg.MessageHash),
		attribute.Int(metrics.Origin, int(rawMsg.OriginChainID)),
		attribute.Int(metrics.Destination, int(rawMsg.DestChainID)),
	)

	// Store the requested message.
	rawMsg.State = relayTypes.Pending
	err = c.db.StoreMessage(ctx, rawMsg)
	if err != nil {
		return nil, fmt.Errorf("could not store pending message: %w", err)
	}

	return &rawMsg, nil
}

func (c *circleCCTPHandler) handleMessageReceived(parentCtx context.Context, log *types.Log, chainID uint32) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "handleMessageReceived", trace.WithAttributes(
		attribute.String(metrics.TxHash, log.TxHash.Hex()),
		attribute.Int(metrics.ChainID, int(chainID)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	if len(log.Topics) == 0 {
		return fmt.Errorf("no topics found")
	}

	// Parse the request id from the log.
	ethClient, err := c.omniRPCClient.GetConfirmationsClient(ctx, int(chainID), 1)
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}

	eventParser, err := messagetransmitter.NewMessageTransmitterFilterer(log.Address, ethClient)
	if err != nil {
		return fmt.Errorf("could not create event parser: %w", err)
	}

	event, err := eventParser.ParseMessageReceived(*log)
	if err != nil {
		return fmt.Errorf("could not parse circle request fulfilled: %w", err)
	}

	// check that we sent the tx
	tx, _, err := ethClient.TransactionByHash(ctx, log.TxHash)
	if err != nil {
		return fmt.Errorf("could not get transaction by hash: %w", err)
	}
	call, err := util.TxToCall(tx)
	if err != nil {
		return fmt.Errorf("could not convert transaction to call: %w", err)
	}
	if call.From != c.relayerAddress {
		span.AddEvent(fmt.Sprintf("sender %s does not match relayer address %s", call.From.String(), c.relayerAddress.String()))
		return nil
	}

	messageHash := crypto.Keccak256Hash(event.MessageBody)
	span.SetAttributes(
		attribute.String("message_hash", messageHash.String()),
		attribute.Int(metrics.Origin, int(event.SourceDomain)),
	)

	// convert the source domain to a chain ID
	originChainID, err := CircleDomainToChainID(event.SourceDomain, IsTestnetChainID(chainID))
	if err != nil {
		return fmt.Errorf("could not convert circle domain to chain ID: %w", err)
	}

	msg, err := c.db.GetMessageByHash(ctx, messageHash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Reconstruct what we can from the given log.
			msg = &relayTypes.Message{
				OriginChainID: originChainID,
				DestChainID:   chainID,
				BlockNumber:   log.BlockNumber,
				MessageHash:   messageHash.String(),
			}
			span.AddEvent("message not found; reconstructing")
		} else {
			return fmt.Errorf("could not get message by hash: %w", err)
		}
	}

	if msg == nil {
		return fmt.Errorf("no message found")
	}

	// Mark as Complete and store the message.
	msg.State = relayTypes.Complete
	msg.DestTxHash = log.TxHash.String()
	err = c.db.StoreMessage(ctx, *msg)
	if err != nil {
		return fmt.Errorf("could not store complete message: %w", err)
	}
	return nil
}
