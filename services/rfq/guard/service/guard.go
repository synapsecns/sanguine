package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/listener"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/guard/guardconfig"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb/connect"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

var logger = log.Logger("guard")

// Guard monitors calls to prove() and verifies them.
type Guard struct {
	cfg            guardconfig.Config
	metrics        metrics.Handler
	db             guarddb.Service
	client         omniClient.RPCClient
	chainListeners map[int]listener.ContractListener
	contracts      map[int]*fastbridge.FastBridgeRef
	txSubmitter    submitter.TransactionSubmitter
	otelRecorder   iOtelRecorder
}

// NewGuard creates a new Guard.
//
//nolint:cyclop
func NewGuard(ctx context.Context, metricHandler metrics.Handler, cfg guardconfig.Config, txSubmitter submitter.TransactionSubmitter) (*Guard, error) {
	omniClient := omniClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omniClient.WithCaptureReqRes())
	chainListeners := make(map[int]listener.ContractListener)

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}
	store, err := connect.Connect(ctx, dbType, cfg.Database.DSN, metricHandler)
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}

	// setup chain listeners
	contracts := make(map[int]*fastbridge.FastBridgeRef)
	for chainID := range cfg.GetChains() {
		rfqAddr, err := cfg.GetRFQAddress(chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get rfq address: %w", err)
		}
		chainClient, err := omniClient.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}

		contract, err := fastbridge.NewFastBridgeRef(common.HexToAddress(rfqAddr), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create fast bridge contract: %w", err)
		}
		startBlock, err := contract.DeployBlock(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("could not get deploy block: %w", err)
		}
		chainListener, err := listener.NewChainListener(chainClient, store, common.HexToAddress(rfqAddr), uint64(startBlock.Int64()), metricHandler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		chainListeners[chainID] = chainListener

		// setup FastBridge contract on this chain
		contracts[chainID], err = fastbridge.NewFastBridgeRef(common.HexToAddress(rfqAddr), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create bridge contract: %w", err)
		}
	}

	// build submitter from config if one is not supplied
	if txSubmitter == nil {
		sg, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
		if err != nil {
			return nil, fmt.Errorf("could not get signer: %w", err)
		}
		txSubmitter = submitter.NewTransactionSubmitter(metricHandler, sg, omniClient, store.SubmitterDB(), &cfg.SubmitterConfig)
	}

	otelRecorder, err := newOtelRecorder(metricHandler, txSubmitter.Address(), store)
	if err != nil {
		return nil, fmt.Errorf("could not get otel recorder: %w", err)
	}

	return &Guard{
		cfg:            cfg,
		metrics:        metricHandler,
		db:             store,
		client:         omniClient,
		chainListeners: chainListeners,
		contracts:      contracts,
		txSubmitter:    txSubmitter,
		otelRecorder:   otelRecorder,
	}, nil
}

// Start starts the guard.
func (g *Guard) Start(ctx context.Context) (err error) {
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		err := g.startChainIndexers(ctx)
		if err != nil {
			return fmt.Errorf("could not start chain indexers: %w", err)
		}
		return nil
	})
	group.Go(func() error {
		err = g.runDBSelector(ctx)
		if err != nil {
			return fmt.Errorf("could not start db selector: %w", err)
		}
		return nil
	})

	group.Go(func() error {
		if !g.txSubmitter.Started() {
			err = g.txSubmitter.Start(ctx)
			// defensive coding against potential race.
			if err != nil && !errors.Is(err, submitter.ErrSubmitterAlreadyStarted) {
				return fmt.Errorf("could not start tx submitter: %w", err)
			}
		}
		return nil
	})

	err = group.Wait()
	if err != nil {
		return fmt.Errorf("could not wait for group: %w", err)
	}

	return nil
}

func (g *Guard) runDBSelector(ctx context.Context) (err error) {
	interval := g.cfg.GetDBSelectorInterval()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("could not run db selector: %w", ctx.Err())
		case <-time.After(interval):
			err := g.processDB(ctx)
			if err != nil {
				return err
			}
		}
	}
}

func (g *Guard) startChainIndexers(ctx context.Context) (err error) {
	group, ctx := errgroup.WithContext(ctx)

	for chainID := range g.cfg.GetChains() {
		chainID := chainID // capture func literal

		group.Go(func() error {
			err := g.runChainIndexer(ctx, chainID)
			if err != nil {
				return fmt.Errorf("could not runChainIndexer chain indexer for chain %d: %w", chainID, err)
			}
			return nil
		})
	}

	err = group.Wait()
	if err != nil {
		return fmt.Errorf("could not run chain indexers")
	}

	return nil
}

//nolint:cyclop
func (g Guard) runChainIndexer(ctx context.Context, chainID int) (err error) {
	chainListener := g.chainListeners[chainID]

	parser, err := fastbridge.NewParser(chainListener.Address())
	if err != nil {
		return fmt.Errorf("could not parse: %w", err)
	}

	err = chainListener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		et, parsedEvent, ok := parser.ParseEvent(log)
		// handle unknown event
		if !ok {
			if len(log.Topics) != 0 {
				logger.Warnf("unknown event %s", log.Topics[0])
			}
			return nil
		}

		ctx, span := g.metrics.Tracer().Start(parentCtx, fmt.Sprintf("handleLog-%s", et), trace.WithAttributes(
			attribute.String(metrics.TxHash, log.TxHash.String()),
			attribute.Int(metrics.Origin, chainID),
			attribute.String(metrics.Contract, log.Address.String()),
			attribute.String("block_hash", log.BlockHash.String()),
			attribute.Int64("block_number", int64(log.BlockNumber)),
		))

		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		switch event := parsedEvent.(type) {
		case *fastbridge.FastBridgeBridgeRequested:
			err = g.handleBridgeRequestedLog(ctx, event, chainID)
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		case *fastbridge.FastBridgeBridgeProofProvided:
			err = g.handleProofProvidedLog(ctx, event, chainID)
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		case *fastbridge.FastBridgeBridgeProofDisputed:
			err = g.handleProofDisputedLog(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("listener failed: %w", err)
	}
	return nil
}

func (g *Guard) processDB(ctx context.Context) (err error) {
	provens, err := g.db.GetPendingProvensByStatus(ctx, guarddb.ProveCalled)
	if err != nil {
		return fmt.Errorf("could not get pending provens: %w", err)
	}

	for _, proven := range provens {
		err := g.handleProveCalled(ctx, proven)
		if err != nil {
			return fmt.Errorf("could not handle prove called: %w", err)
		}
	}

	return nil
}
