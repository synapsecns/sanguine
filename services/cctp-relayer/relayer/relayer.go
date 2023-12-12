package relayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	"github.com/synapsecns/sanguine/services/cctp-relayer/attestation"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/mockmessagetransmitter"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// chainListener is a struct that contains the necessary information for each chain level relayer.
type chainListener struct {
	// chainID is the chain ID of the chain that this relayer is responsible for.
	chainID uint32
	// closeConnection is a channel that is used to close the connection.
	closeConnection chan bool
	// stopListenChan is a channel that is used to stop listening to the log channel.
	stopListenChan chan bool
}

// CCTPRelayer listens for USDC burn events on origin chains,
// fetches attestations from Circle's API, and posts the necessary data
// on the destination chain to complete the USDC bridging process.
type CCTPRelayer struct {
	cfg           config.Config
	db            db2.CCTPRelayerDB
	scribeClient  client.ScribeClient
	grpcClient    pbscribe.ScribeServiceClient
	grpcConn      *grpc.ClientConn
	omnirpcClient omniClient.RPCClient
	// chainListeners is a map from chain ID -> chain relayer.
	chainListeners map[uint32]*chainListener
	// handler is the metrics handler.
	handler metrics.Handler
	// attestationAPI is the client for Circle's REST API.
	attestationAPI attestation.CCTPAPI
	// txSubmitter is the tx submission service
	txSubmitter submitter.TransactionSubmitter
	// boundSynapseCCTPs is a map from chain ID -> SynapseCCTP.
	boundSynapseCCTPs map[uint32]*cctp.SynapseCCTP
	// relayerAPI is the relayer api server for queueing external relay requests.
	relayerAPI *api.RelayerAPIServer
	// relayRequestChan is a channel that is used to process relay requests from the api server.
	relayRequestChan chan *api.RelayRequest
	// retryNow is used to trigger a retry immediately.
	// it circumvents the retry interval.
	// to prevent memory leaks, this has a buffer of 1.
	// callers adding to this channel should not block.
	retryNow chan bool
	// retryOnce is used to return 0 for the first retry. timer
	retryOnce sync.Once
}

// NewCCTPRelayer creates a new CCTPRelayer.
func NewCCTPRelayer(ctx context.Context, cfg config.Config, store db2.CCTPRelayerDB, scribeClient client.ScribeClient, omniRPCClient omniClient.RPCClient, handler metrics.Handler, attestationAPI attestation.CCTPAPI) (*CCTPRelayer, error) {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
	)
	if err != nil {
		return nil, fmt.Errorf("could not dial grpc: %w", err)
	}

	grpcClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := grpcClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return nil, fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return nil, fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	// Build chainListeners and bound contracts.
	chainListeners := make(map[uint32]*chainListener)
	boundSynapseCCTPs := make(map[uint32]*cctp.SynapseCCTP)
	for _, chain := range cfg.Chains {
		chainListeners[chain.ChainID] = &chainListener{
			chainID:         chain.ChainID,
			closeConnection: make(chan bool, 1),
			stopListenChan:  make(chan bool, 1),
			// processChan is buffered to prevent blocking.
		}
		cl, err := omniRPCClient.GetConfirmationsClient(ctx, int(chain.ChainID), 1)
		if err != nil {
			return nil, fmt.Errorf("could not get client: %w", err)
		}
		boundSynapseCCTPs[chain.ChainID], err = cctp.NewSynapseCCTP(chain.GetSynapseCCTPAddress(), cl)
		if err != nil {
			return nil, fmt.Errorf("could not build bound contract: %w", err)
		}
	}

	signer, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not make cctp signer: %w", err)
	}

	txSubmitter := submitter.NewTransactionSubmitter(handler, signer, omniRPCClient, store.SubmitterDB(), &cfg.SubmitterConfig)

	relayerRequestChan := make(chan *api.RelayRequest, 1000)
	relayerAPI := api.NewRelayerAPIServer(cfg.Port, cfg.Host, store, relayerRequestChan)

	return &CCTPRelayer{
		cfg:               cfg,
		db:                store,
		omnirpcClient:     omniRPCClient,
		chainListeners:    chainListeners,
		scribeClient:      scribeClient,
		grpcClient:        grpcClient,
		grpcConn:          conn,
		relayerAPI:        relayerAPI,
		relayRequestChan:  relayerRequestChan,
		retryNow:          make(chan bool, 1),
		handler:           handler,
		attestationAPI:    attestationAPI,
		txSubmitter:       txSubmitter,
		boundSynapseCCTPs: boundSynapseCCTPs,
	}, nil
}

// triggerProcessQueue triggers the process queue.
// will not block if the channel is full (the tx will be processed on the next retry).
func (c *CCTPRelayer) triggerProcessQueue(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	// trigger the process queue now if we can.
	case c.retryNow <- true:
	default:
		return
	}
}

const defaultRetryInterval = 5 * time.Second

// getRetryInterval returns the retry interval.
// on the first try this is 0 and it is the configured interval (or default) after that.
func (c *CCTPRelayer) getRetryInterval() time.Duration {
	retryInterval := time.Duration(c.cfg.RetryIntervalMS)
	if retryInterval == 0 {
		retryInterval = defaultRetryInterval
	}

	// make 0 on first try
	c.retryOnce.Do(func() {
		retryInterval = 0
	})
	return retryInterval
}

func (c *CCTPRelayer) runQueueSelector(ctx context.Context) (err error) {
	for {
		select {
		case <-ctx.Done():
			//nolint: wrapcheck
			return ctx.Err()
		case <-c.retryNow:
			err = c.processQueue(ctx)
		case <-time.After(c.getRetryInterval()):
			err = c.processQueue(ctx)
		}
		if err != nil {
			logger.Warnf("could not process queue: %v", err)
		}
	}
}

// nolint: cyclop
func (c *CCTPRelayer) processQueue(parentCtx context.Context) (err error) {
	// TODO: this might be too short of a deadline depending on the number of pendingTxes in the queue
	deadlineCtx, cancel := context.WithTimeout(parentCtx, time.Second*90)
	defer cancel()

	ctx, span := c.handler.Tracer().Start(deadlineCtx, "relayer.ProcessQueue")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	attestations, err := c.db.GetMessagesByState(ctx, relayTypes.Pending, relayTypes.Attested)
	if err != nil {
		return fmt.Errorf("could not get pending messages: %w", err)
	}

	g, gctx := errgroup.WithContext(ctx)

	// add attestations to the queue
	attQueue := make(chan *relayTypes.Message, len(attestations))
	for i := range attestations {
		select {
		case <-gctx.Done():
			return fmt.Errorf("could not process: %w", gctx.Err())
		case attQueue <- &attestations[i]:
			// queue to be reprocessed
		}
	}
	// TODO: consider closing channel here?

	// process queue 7 at a time
	for i := 0; i < 7; i++ {
		g.Go(func() error {
			for {
				select {
				case <-gctx.Done():
					return fmt.Errorf("could not process: %w", gctx.Err())
				case attToProcess := <-attQueue:
					err := c.processMessage(gctx, attToProcess)
					if err != nil {
						logger.Warnf("could not process: %v", err)
					}
					continue
				default:
					return nil
				}
			}
		})
	}

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not process: %w", err)
	}

	return nil
}

// processMessage processes a message. Before each stage it checks if the current step is done.
func (c *CCTPRelayer) processMessage(parentCtx context.Context, msg *relayTypes.Message) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "processMessage", trace.WithAttributes(
		attribute.String(MessageHash, msg.MessageHash),
		attribute.Int(metrics.Origin, int(msg.OriginChainID)),
		attribute.Int(metrics.Destination, int(msg.DestChainID)),
		attribute.String(metrics.TxHash, msg.OriginTxHash),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	if msg.State == relayTypes.Pending {
		msg, err = c.fetchAttestation(ctx, msg)
		if err != nil {
			return fmt.Errorf("could not fetch attestation: %w", err)
		}
	}

	if msg.State == relayTypes.Attested {
		err := c.submitReceiveCircleToken(ctx, msg)
		if err != nil {
			return fmt.Errorf("could not submit receive circle token: %w", err)
		}
	}

	return nil
}

// Run starts the CCTPRelayer.
func (c *CCTPRelayer) Run(parentCtx context.Context) error {
	g, ctx := errgroup.WithContext(parentCtx)

	// Listen for USDC burn events on origin chains.
	for _, chain := range c.cfg.Chains {
		chain := chain
		g.Go(func() error {
			return c.streamLogs(ctx, c.grpcClient, c.grpcConn, chain.ChainID, chain.SynapseCCTPAddress, nil)
		})
	}

	g.Go(func() error {
		return c.runQueueSelector(ctx)
	})

	g.Go(func() error {
		err := c.txSubmitter.Start(ctx)
		if err != nil {
			err = fmt.Errorf("could not start tx submitter: %w", err)
		}
		return err
	})

	g.Go(func() error {
		err := c.processAPIRequests(ctx)
		if err != nil {
			err = fmt.Errorf("could not process api requests: %w", err)
		}
		return err
	})

	g.Go(func() error {
		err := c.relayerAPI.Start(ctx)
		if err != nil {
			err = fmt.Errorf("could not start relayer api: %w", err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error in cctp relayer: %w", err)
	}

	return nil
}

// Stop stops the CCTPRelayer.
func (c *CCTPRelayer) Stop(chainID uint32) {
	c.chainListeners[chainID].closeConnection <- true
	c.chainListeners[chainID].stopListenChan <- true
}

// Listens for USDC send events on origin chain, and registers relayTypes.Messages to be signed.
//
//nolint:cyclop
func (c *CCTPRelayer) streamLogs(ctx context.Context, grpcClient pbscribe.ScribeServiceClient, conn *grpc.ClientConn, chainID uint32, address string, toBlockNumber *uint64) error {
	lastStoredBlock, err := c.db.GetLastBlockNumber(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not get last stored block: %w", err)
	}

	fromBlock := strconv.FormatUint(lastStoredBlock, 16)

	toBlock := "latest"
	if toBlockNumber != nil {
		toBlock = strconv.FormatUint(*toBlockNumber, 16)
	}

	stream, err := grpcClient.StreamLogs(ctx, &pbscribe.StreamLogsRequest{
		Filter: &pbscribe.LogFilter{
			ContractAddress: &pbscribe.NullableString{Kind: &pbscribe.NullableString_Data{Data: address}},
			ChainId:         chainID,
		},
		FromBlock: fromBlock,
		ToBlock:   toBlock,
	})
	if err != nil {
		return fmt.Errorf("could not stream logs: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done: %w", ctx.Err())
		case <-c.chainListeners[chainID].closeConnection:
			err := stream.CloseSend()
			if err != nil {
				return fmt.Errorf("could not close stream: %w", err)
			}

			err = conn.Close()
			if err != nil {
				return fmt.Errorf("could not close connection: %w", err)
			}

			return nil
		default:
			response, err := stream.Recv()
			if err != nil {
				return fmt.Errorf("could not receive: %w", err)
			}

			err = c.handleLog(ctx, response.Log.ToLog(), chainID)
			if err != nil {
				return err
			}
		}
	}
}

// processAPIRequests processes requests from the API.
func (c *CCTPRelayer) processAPIRequests(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done: %w", ctx.Err())
		default:
			var relayRequest *api.RelayRequest
			select {
			case relayRequest = <-c.relayRequestChan:
				// continue below
			case <-ctx.Done():
				return fmt.Errorf("context done while attempting to add to channel: %w", ctx.Err())
			}
			msg, err := c.fetchAndStoreCircleRequestSent(ctx, relayRequest.TxHash, relayRequest.Origin)
			if err != nil {
				return fmt.Errorf("could not fetch and store circle request sent from api: %w", err)
			}

			if msg != nil {
				c.triggerProcessQueue(ctx)
			}
		}
	}
}

// This takes in a log from the SynapseCCTP contract, determines the topic and then performs an action based on that topic.
// Note that the log could correspond to a send or receive event.
func (c *CCTPRelayer) handleLog(ctx context.Context, log *types.Log, chainID uint32) (err error) {
	if log == nil {
		return fmt.Errorf("log is nil")
	}

	// shouldn't be possible: maybe remove?
	if len(log.Topics) == 0 {
		return fmt.Errorf("not enough topics")
	}

	switch log.Topics[0] {
	// since this is the last stopic that comes out of the message, we use it to kick off the send loop
	case cctp.CircleRequestSentTopic:
		msg, err := c.fetchAndStoreCircleRequestSent(ctx, log.TxHash, chainID)
		if err != nil {
			return fmt.Errorf("could not fetch and store circle request sent: %w", err)
		}

		if msg != nil {
			c.triggerProcessQueue(ctx)
		}

		return nil
	case cctp.CircleRequestFulfilledTopic:
		err = c.handleCircleRequestFulfilled(ctx, log, chainID)
		if err != nil {
			return fmt.Errorf("could not store circle request fulfilled: %w", err)
		}
		return nil
	default:
		return nil
	}
}

// fetchAndStoreCircleRequestSent handles the CircleRequestSent event.
//
//nolint:cyclop
func (c *CCTPRelayer) fetchAndStoreCircleRequestSent(parentCtx context.Context, txhash common.Hash, originChain uint32) (msg *relayTypes.Message, err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "fetchAndStoreCircleRequestSent", trace.WithAttributes(
		attribute.String(metrics.TxHash, txhash.String()),
		attribute.Int(metrics.ChainID, int(originChain)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// check if message already exist before we do anything
	msg, err = c.db.GetMessageByOriginHash(ctx, txhash)
	// if we already have the message, we can just return it
	if err == nil {
		return msg, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("could not get message by origin hash: %w", err)
	}

	ethClient, err := c.omnirpcClient.GetConfirmationsClient(ctx, int(originChain), 1)
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
	var messageSentEvent *mockmessagetransmitter.MessageTransmitterEventsMessageSent
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
			// TODO: this shouldn't be coming from a mock contract, generate from the abstract contract itself
		case mockmessagetransmitter.MessageSentTopic:
			eventParser, err := mockmessagetransmitter.NewMessageTransmitterEvents(log.Address, ethClient)
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
	err = c.db.StoreMessage(ctx, rawMsg)
	if err != nil {
		return nil, fmt.Errorf("could not store pending message: %w", err)
	}

	return &rawMsg, nil
}

// handleCircleRequestFulfilled handles the CircleRequestFulfilled event.
//
//nolint:cyclop
func (c *CCTPRelayer) handleCircleRequestFulfilled(parentCtx context.Context, log *types.Log, destChain uint32) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "handleCircleRequestFulfilled", trace.WithAttributes(
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
	ethClient, err := c.omnirpcClient.GetConfirmationsClient(ctx, int(destChain), 1)
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

	err = c.storeCircleRequestFulfilled(ctx, log, circleRequestFulfilledEvent, destChain)
	if err != nil {
		return fmt.Errorf("could not store circle request fulfilled: %w", err)
	}

	return nil
}

// storeCircleRequestFullfilled fetches pending message from db, and marks as complete if found.
// If the message is not found, it will be created from the given log.
func (c *CCTPRelayer) storeCircleRequestFulfilled(ctx context.Context, log *types.Log, event *cctp.SynapseCCTPEventsCircleRequestFulfilled, destChain uint32) error {
	var msg *relayTypes.Message
	requestID := common.Bytes2Hex(event.RequestID[:])
	msg, err := c.db.GetMessageByRequestID(ctx, requestID)
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
	err = c.db.StoreMessage(ctx, *msg)
	if err != nil {
		return fmt.Errorf("could not store complete message: %w", err)
	}
	return nil
}

func (c *CCTPRelayer) fetchAttestation(parentCtx context.Context, msg *relayTypes.Message) (_ *relayTypes.Message, err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "fetchAttestation", trace.WithAttributes(
		attribute.String(MessageHash, msg.MessageHash),
		attribute.Int(metrics.Origin, int(msg.OriginChainID)),
		attribute.Int(metrics.Destination, int(msg.DestChainID)),
		attribute.String(metrics.TxHash, msg.OriginTxHash),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	err = retry.WithBackoff(ctx, func(ctx context.Context) (err error) {
		msg.Attestation, err = c.attestationAPI.GetAttestation(ctx, msg.MessageHash)
		return
	}, retry.WithMax(time.Duration(c.cfg.HTTPBackoffMaxElapsedTimeMs)*time.Millisecond))
	if err != nil {
		return
	}

	// Store the attested message.
	msg.State = relayTypes.Attested
	err = c.db.StoreMessage(ctx, *msg)
	if err != nil {
		return nil, fmt.Errorf("could not store attested message: %w", err)
	}

	return msg, nil
}

func (c *CCTPRelayer) submitReceiveCircleToken(parentCtx context.Context, msg *relayTypes.Message) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "submitReceiveCircleToken", trace.WithAttributes(
		attribute.String(MessageHash, msg.MessageHash),
		attribute.Int(metrics.Origin, int(msg.OriginChainID)),
		attribute.Int(metrics.Destination, int(msg.DestChainID)),
		attribute.String(metrics.TxHash, msg.OriginTxHash),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	contract, ok := c.boundSynapseCCTPs[msg.DestChainID]
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
		err = c.db.StoreMessage(ctx, *msg)
		if err != nil {
			return fmt.Errorf("could not store completed message: %w", err)
		}

		return nil
	}
	// end: functionalization

	var nonce uint64
	var destTxHash common.Hash
	nonce, err = c.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(msg.DestChainID)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
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
	err = c.db.StoreMessage(ctx, *msg)
	if err != nil {
		return fmt.Errorf("could not store completed message: %w", err)
	}
	return nil
}
