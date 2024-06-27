package relayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/contrib/screener-api/client"

	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/messagetransmitter"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type synapseCCTPHandler struct {
	cfg               config.Config
	db                db2.CCTPRelayerDB
	omniRPCClient     omniClient.RPCClient
	boundSynapseCCTPs map[uint32]*cctp.SynapseCCTP
	txSubmitter       submitter.TransactionSubmitter
	handler           metrics.Handler
	// screener is used to screen addresses.
	screener client.ScreenerClient
}

// NewSynapseCCTPHandler creates a new SynapseCCTPHandler.
func NewSynapseCCTPHandler(ctx context.Context, cfg config.Config, db db2.CCTPRelayerDB, omniRPCClient omniClient.RPCClient, txSubmitter submitter.TransactionSubmitter, handler metrics.Handler) (_ CCTPHandler, err error) {
	boundSynapseCCTPs := make(map[uint32]*cctp.SynapseCCTP)
	for _, chain := range cfg.Chains {
		cl, err := omniRPCClient.GetConfirmationsClient(ctx, int(chain.ChainID), 1)
		if err != nil {
			return nil, fmt.Errorf("could not get client: %w", err)
		}
		boundSynapseCCTPs[chain.ChainID], err = cctp.NewSynapseCCTP(chain.GetCCTPAddress(), cl)
		if err != nil {
			return nil, fmt.Errorf("could not build bound contract: %w", err)
		}
	}

	var ss client.ScreenerClient
	if cfg.ScreenerAPIUrl != "" {
		ss, err = client.NewClient(handler, cfg.ScreenerAPIUrl)
		if err != nil {
			return nil, fmt.Errorf("error creating screener client: %w", err)
		}
	}

	return &synapseCCTPHandler{
		cfg:               cfg,
		db:                db,
		screener:          ss,
		omniRPCClient:     omniRPCClient,
		boundSynapseCCTPs: boundSynapseCCTPs,
		txSubmitter:       txSubmitter,
		handler:           handler,
	}, nil
}

const screenerRuleset = "cctp"

func (s *synapseCCTPHandler) HandleLog(ctx context.Context, log *types.Log, chainID uint32) (processQueue bool, err error) {
	if log == nil {
		return false, fmt.Errorf("log is nil")
	}

	// shouldn't be possible: maybe remove?
	if len(log.Topics) == 0 {
		return false, fmt.Errorf("not enough topics")
	}

	switch log.Topics[0] {
	// since this is the last stopic that comes out of the message, we use it to kick off the send loop
	case cctp.CircleRequestSentTopic:
		msg, err := s.FetchAndProcessSentEvent(ctx, log.TxHash, chainID)
		if err != nil {
			return false, fmt.Errorf("could not fetch and store circle request sent: %w", err)
		}

		if msg != nil {
			processQueue = true
		}

		return processQueue, nil
	case cctp.CircleRequestFulfilledTopic:
		err = s.handleCircleRequestFulfilled(ctx, log, chainID)
		if err != nil {
			return false, fmt.Errorf("could not store circle request fulfilled: %w", err)
		}
		return false, nil
	default:
		logger.Warnf("unknown topic %s", log.Topics[0])
		return false, nil
	}
}

// fetchAndStoreCircleRequestSent handles the CircleRequestSent event.
//
//nolint:cyclop
func (s *synapseCCTPHandler) FetchAndProcessSentEvent(parentCtx context.Context, txhash common.Hash, originChain uint32) (msg *relayTypes.Message, err error) {
	ctx, span := s.handler.Tracer().Start(parentCtx, "FetchAndProcessSentEvent", trace.WithAttributes(
		attribute.String(metrics.TxHash, txhash.String()),
		attribute.Int(metrics.ChainID, int(originChain)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// check if message already exists before we do anything
	msg, err = s.db.GetMessageByOriginHash(ctx, txhash)
	// if we already have the message, we can just return it
	if err == nil {
		return msg, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("could not get message by origin hash: %w", err)
	}

	ethClient, err := s.omniRPCClient.GetConfirmationsClient(ctx, int(originChain), 1)
	if err != nil {
		return nil, fmt.Errorf("could not get chain client: %w", err)
	}

	// TODO: consider pulling from scribe?
	// TODO: this function is particularly prone to silent errors
	receipt, err := ethClient.TransactionReceipt(ctx, txhash)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction receipt: %w", err)
	}

	// From this receipt, we expect two different logs:
	// - `messageSentEvent` gives us the raw bytes of the CCTP message
	// - `circleRequestSentEvent` gives us auxiliary data for SynapseCCTP
	var messageSentEvent *messagetransmitter.MessageTransmitterMessageSent
	var circleRequestSentEvent *cctp.SynapseCCTPEventsCircleRequestSent

	for _, log := range receipt.Logs {
		// this should never happen
		if len(log.Topics) == 0 {
			continue
		}

		switch log.Topics[0] {
		case cctp.CircleRequestSentTopic:
			// TODO: do we need to make sure log.Address matches our log.Address?
			eventParser, err := cctp.NewSynapseCCTPEvents(log.Address, ethClient)
			if err != nil {
				return nil, fmt.Errorf("could not create event parser: %w", err)
			}

			circleRequestSentEvent, err = eventParser.ParseCircleRequestSent(*log)
			if err != nil {
				return nil, fmt.Errorf("could not parse circle request sent: %w", err)
			}
		case messagetransmitter.MessageSentTopic:
			eventParser, err := messagetransmitter.NewMessageTransmitterFilterer(log.Address, ethClient)
			if err != nil {
				return nil, fmt.Errorf("could not create event parser: %w", err)
			}

			messageSentEvent, err = eventParser.ParseMessageSent(*log)
			if err != nil {
				return nil, fmt.Errorf("could not parse message sent: %w", err)
			}
		}
	}

	if messageSentEvent == nil {
		return nil, fmt.Errorf("no message sent event found")
	}

	if circleRequestSentEvent == nil {
		return nil, fmt.Errorf("no circle request sent event found")
	}

	rawMsg := relayTypes.Message{
		OriginTxHash:  txhash.String(),
		OriginChainID: originChain,
		DestChainID:   uint32(circleRequestSentEvent.ChainId.Int64()),
		Message:       messageSentEvent.Message,
		MessageHash:   crypto.Keccak256Hash(messageSentEvent.Message).String(),
		RequestID:     common.Bytes2Hex(circleRequestSentEvent.RequestID[:]),

		//Attestation: //comes from the api
		RequestVersion:   circleRequestSentEvent.RequestVersion,
		FormattedRequest: circleRequestSentEvent.FormattedRequest,
		BlockNumber:      uint64(receipt.BlockNumber.Int64()),
	}

	// Store the requested message.
	rawMsg.State = relayTypes.Pending

	if s.screener != nil {
		blocked, err := s.screener.ScreenAddress(ctx, circleRequestSentEvent.Sender.String())
		if err != nil {
			return nil, fmt.Errorf("error screening address: %w", err)
		}

		if blocked {
			rawMsg.State = relayTypes.WillNotComplete
		}
	}

	err = s.db.StoreMessage(ctx, rawMsg)
	if err != nil {
		return nil, fmt.Errorf("could not store pending message: %w", err)
	}

	return &rawMsg, nil
}

func (s *synapseCCTPHandler) SubmitReceiveMessage(parentCtx context.Context, msg *relayTypes.Message) (err error) {
	ctx, span := s.handler.Tracer().Start(parentCtx, "SubmitReceiveMessage", trace.WithAttributes(
		attribute.String(MessageHash, msg.MessageHash),
		attribute.Int(metrics.Origin, int(msg.OriginChainID)),
		attribute.Int(metrics.Destination, int(msg.DestChainID)),
		attribute.String(metrics.TxHash, msg.OriginTxHash),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	contract, ok := s.boundSynapseCCTPs[msg.DestChainID]
	if !ok {
		return fmt.Errorf("could not find destination chain %d", msg.DestChainID)
	}

	// TODO: functionalize this
	ridBytes := common.Hex2Bytes(msg.RequestID)
	var rid [32]byte
	copy(rid[:], ridBytes)

	isFulfilled, err := contract.IsRequestFulfilled(&bind.CallOpts{Context: ctx}, rid)
	if err != nil {
		return fmt.Errorf("could not check if request is fulfilled: %w", err)
	}
	if isFulfilled {
		msg.State = relayTypes.Complete
		err = s.db.StoreMessage(ctx, *msg)
		if err != nil {
			return fmt.Errorf("could not store completed message: %w", err)
		}

		return nil
	}
	// end: functionalization

	var nonce uint64
	var destTxHash common.Hash
	nonce, err = s.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(msg.DestChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		gasAmount, err := contract.ChainGasAmount(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("could not get chain gas amount: %w", err)
		}
		transactor.Value = gasAmount

		tx, err = contract.ReceiveCircleToken(transactor, msg.Message, msg.Attestation, msg.RequestVersion, msg.FormattedRequest)
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

	// Store the completed message.
	// Note: this can cause double submission sometimes
	msg.State = relayTypes.Submitted
	msg.DestNonce = int(nonce)
	msg.DestTxHash = destTxHash.String()
	err = s.db.StoreMessage(ctx, *msg)
	if err != nil {
		return fmt.Errorf("could not store completed message: %w", err)
	}
	return nil
}

// handleCircleRequestFulfilled handles the CircleRequestFulfilled event.
//
//nolint:cyclop
func (s *synapseCCTPHandler) handleCircleRequestFulfilled(parentCtx context.Context, log *types.Log, destChain uint32) (err error) {
	ctx, span := s.handler.Tracer().Start(parentCtx, "handleCircleRequestFulfilled", trace.WithAttributes(
		attribute.String(metrics.TxHash, log.TxHash.String()),
		attribute.Int(metrics.Destination, int(destChain)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	if len(log.Topics) == 0 {
		return fmt.Errorf("no topics found")
	}

	// Parse the request id from the log.
	ethClient, err := s.omniRPCClient.GetConfirmationsClient(ctx, int(destChain), 1)
	if err != nil {
		return fmt.Errorf("could not get chain client: %w", err)
	}
	if log.Topics[0] != cctp.CircleRequestFulfilledTopic {
		return fmt.Errorf("log topic does not match CircleRequestFulfilledTopic")
	}
	eventParser, err := cctp.NewSynapseCCTPEvents(log.Address, ethClient)
	if err != nil {
		return fmt.Errorf("could not create event parser: %w", err)
	}
	circleRequestFulfilledEvent, err := eventParser.ParseCircleRequestFulfilled(*log)
	if err != nil {
		return fmt.Errorf("could not parse circle request fulfilled: %w", err)
	}

	err = s.storeCircleRequestFulfilled(ctx, log, circleRequestFulfilledEvent, destChain)
	if err != nil {
		return fmt.Errorf("could not store circle request fulfilled: %w", err)
	}

	return nil
}

// storeCircleRequestFullfilled fetches pending message from db, and marks as complete if found.
// If the message is not found, it will be created from the given log.
func (s *synapseCCTPHandler) storeCircleRequestFulfilled(ctx context.Context, log *types.Log, event *cctp.SynapseCCTPEventsCircleRequestFulfilled, destChain uint32) error {
	var msg *relayTypes.Message
	requestID := common.Bytes2Hex(event.RequestID[:])
	msg, err := s.db.GetMessageByRequestID(ctx, requestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Reconstruct what we can from the given log.
			msg = &relayTypes.Message{
				OriginChainID: event.OriginDomain,
				DestChainID:   destChain,
				RequestID:     requestID,
				BlockNumber:   log.BlockNumber,
			}
		} else {
			return fmt.Errorf("could not get message by request id: %w", err)
		}
	}

	// Mark as Complete and store the message.
	msg.State = relayTypes.Complete
	msg.DestTxHash = log.TxHash.String()
	err = s.db.StoreMessage(ctx, *msg)
	if err != nil {
		return fmt.Errorf("could not store complete message: %w", err)
	}
	return nil
}
