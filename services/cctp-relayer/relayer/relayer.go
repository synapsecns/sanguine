package relayer

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	"github.com/synapsecns/sanguine/services/cctp-relayer/attestation"
	db2 "github.com/synapsecns/sanguine/services/cctp-relayer/db"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

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
	// cctpHandler interacts with CCTP contracts.
	cctpHandler CCTPHandler
	// handler is the metrics handler.
	handler metrics.Handler
	// attestationAPI is the client for Circle's REST API.
	attestationAPI attestation.CCTPAPI
	// txSubmitter is the tx submission service
	txSubmitter submitter.TransactionSubmitter
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

	// Build chainListeners.
	chainListeners := make(map[uint32]*chainListener)
	for _, chain := range cfg.Chains {
		chainListeners[chain.ChainID] = &chainListener{
			chainID:         chain.ChainID,
			closeConnection: make(chan bool, 1),
			stopListenChan:  make(chan bool, 1),
			// processChan is buffered to prevent blocking.
		}
	}

	signer, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not make cctp signer: %w", err)
	}

	txSubmitter := submitter.NewTransactionSubmitter(handler, signer, omniRPCClient, store.SubmitterDB(), &cfg.SubmitterConfig)
	relayerRequestChan := make(chan *api.RelayRequest, 1000)
	relayerAPI := api.NewRelayerAPIServer(cfg.Port, cfg.Host, store, relayerRequestChan)

	cctpType, err := cfg.GetCCTPType()
	if err != nil {
		return nil, fmt.Errorf("could not get cctp type: %w", err)
	}
	var cctpHandler CCTPHandler
	switch cctpType {
	case relayTypes.SynapseMessageType:
		cctpHandler, err = NewSynapseCCTPHandler(ctx, cfg, store, omniRPCClient, txSubmitter, handler)
		if err != nil {
			return nil, fmt.Errorf("could not make synapse cctp handler: %w", err)
		}
	case relayTypes.CircleMessageType:
		cctpHandler, err = NewCircleCCTPHandler(ctx, cfg, store, omniRPCClient, txSubmitter, handler)
		if err != nil {
			return nil, fmt.Errorf("could not make circle cctp handler: %w", err)
		}
	default:
		return nil, fmt.Errorf("unknown cctp type: %s", cctpType)
	}

	return &CCTPRelayer{
		cfg:              cfg,
		db:               store,
		omnirpcClient:    omniRPCClient,
		chainListeners:   chainListeners,
		cctpHandler:      cctpHandler,
		scribeClient:     scribeClient,
		grpcClient:       grpcClient,
		grpcConn:         conn,
		relayerAPI:       relayerAPI,
		relayRequestChan: relayerRequestChan,
		retryNow:         make(chan bool, 1),
		handler:          handler,
		attestationAPI:   attestationAPI,
		txSubmitter:      txSubmitter,
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
		err := c.cctpHandler.SubmitReceiveMessage(ctx, msg)
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

			shouldProcess, err := c.cctpHandler.HandleLog(ctx, response.Log.ToLog(), chainID)
			if err != nil {
				return err
			}
			if shouldProcess {
				c.triggerProcessQueue(ctx)
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
			msg, err := c.cctpHandler.FetchAndProcessSentEvent(ctx, relayRequest.TxHash, relayRequest.Origin)
			if err != nil {
				return fmt.Errorf("could not fetch and store circle request sent from api: %w", err)
			}

			if msg != nil {
				c.triggerProcessQueue(ctx)
			}
		}
	}
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
