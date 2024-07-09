package relayer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/listener"
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
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// CCTPRelayer listens for USDC burn events on origin chains,
// fetches attestations from Circle's API, and posts the necessary data
// on the destination chain to complete the USDC bridging process.
type CCTPRelayer struct {
	cfg           config.Config
	db            db2.CCTPRelayerDB
	grpcConn      *grpc.ClientConn
	omnirpcClient omniClient.RPCClient
	// chainListeners contains a set of contract listeners for each chain.
	chainListeners map[uint32][]listener.ContractListener
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
// nolint: cyclop
func NewCCTPRelayer(ctx context.Context, cfg config.Config, store db2.CCTPRelayerDB, omniRPCClient omniClient.RPCClient, handler metrics.Handler, attestationAPI attestation.CCTPAPI, rawOpts ...OptionsArgsOption) (*CCTPRelayer, error) {
	opts := makeOptions(rawOpts)

	omniClient := omniClient.NewOmnirpcClient(cfg.BaseOmnirpcURL, handler, omniClient.WithCaptureReqRes())

	// setup chain listeners
	chainListeners := make(map[uint32][]listener.ContractListener)
	for _, chainCfg := range cfg.Chains {
		chainID := chainCfg.ChainID
		chainClient, err := omniClient.GetChainClient(ctx, int(chainID))
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}
		chainListener, err := listener.NewChainListener(chainClient, store, chainCfg.GetCCTPAddress(), chainCfg.CCTPStartBlock, handler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		listeners := []listener.ContractListener{chainListener}
		cctpType, err := cfg.GetCCTPType()
		if err != nil {
			return nil, fmt.Errorf("could not get cctp type: %w", err)
		}
		if cctpType == relayTypes.CircleMessageType {
			// fetch the MessageTransmitter address from TokenMessenger contract and create a new listener
			transmitterAddr, err := GetMessageTransmitterAddress(ctx, chainCfg.GetCCTPAddress(), chainClient)
			if err != nil {
				return nil, fmt.Errorf("could not get message transmitter address: %w", err)
			}
			transmitterListener, err := listener.NewChainListener(chainClient, store, transmitterAddr, chainCfg.CCTPStartBlock, handler)
			if err != nil {
				return nil, fmt.Errorf("could not get chain listener: %w", err)
			}
			listeners = append(listeners, transmitterListener)
		}
		chainListeners[chainID] = listeners
	}

	signer, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not make cctp signer: %w", err)
	}

	txSubmitter := opts.submitter
	if txSubmitter == nil {
		txSubmitter = submitter.NewTransactionSubmitter(handler, signer, omniRPCClient, store.SubmitterDB(), &cfg.SubmitterConfig)
	}
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

	// listen for contract events
	for cid, listeners := range c.chainListeners {
		for _, l := range listeners {
			listener := l
			chainID := cid
			g.Go(func() error {
				err := listener.Listen(ctx, func(ctx context.Context, log types.Log) error {
					return c.handleLog(ctx, log, chainID)
				})
				if err != nil {
					return fmt.Errorf("could not listen: %w", err)
				}
				return nil
			})
		}
	}

	g.Go(func() error {
		return c.runQueueSelector(ctx)
	})

	g.Go(func() error {
		if !c.txSubmitter.Started() {
			err := c.txSubmitter.Start(ctx)
			if err != nil && !errors.Is(err, submitter.ErrSubmitterAlreadyStarted) {
				return fmt.Errorf("could not start tx submitter: %w", err)
			}
		}
		return nil
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

func (c *CCTPRelayer) handleLog(ctx context.Context, log types.Log, chainID uint32) (err error) {
	shouldProcess, err := c.cctpHandler.HandleLog(ctx, &log, chainID)
	if err != nil {
		logger.Warn("error handling log: ", err)
	}
	if shouldProcess {
		c.triggerProcessQueue(ctx)
	}
	return nil
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
