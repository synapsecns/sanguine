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
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgev2"
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
	cfg          guardconfig.Config
	metrics      metrics.Handler
	db           guarddb.Service
	client       omniClient.RPCClient
	contractsV1  map[int]*fastbridge.FastBridgeRef
	contractsV2  map[int]*fastbridgev2.FastBridgeV2Ref
	listenersV1  map[int]listener.ContractListener
	listenersV2  map[int]listener.ContractListener
	txSubmitter  submitter.TransactionSubmitter
	otelRecorder iOtelRecorder
}

// NewGuard creates a new Guard.
//
//nolint:cyclop
func NewGuard(ctx context.Context, metricHandler metrics.Handler, cfg guardconfig.Config, txSubmitter submitter.TransactionSubmitter) (*Guard, error) {
	omniClient := omniClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omniClient.WithCaptureReqRes())

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}
	store, err := connect.Connect(ctx, dbType, cfg.Database.DSN, metricHandler)
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}

	contractsV1 := make(map[int]*fastbridge.FastBridgeRef)
	contractsV2 := make(map[int]*fastbridgev2.FastBridgeV2Ref)
	listenersV1 := make(map[int]listener.ContractListener)
	listenersV2 := make(map[int]listener.ContractListener)

	// setup chain listeners
	for chainID := range cfg.GetChains() {
		// setup v1
		rfqAddrV1, err := cfg.GetRFQAddressV1(chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get rfq address: %w", err)
		}
		if rfqAddrV1 != nil {
			chainClient, err := omniClient.GetChainClient(ctx, chainID)
			if err != nil {
				return nil, fmt.Errorf("could not get chain client: %w", err)
			}
			contract, err := fastbridge.NewFastBridgeRef(common.HexToAddress(*rfqAddrV1), chainClient)
			if err != nil {
				return nil, fmt.Errorf("could not create fast bridge contract: %w", err)
			}
			startBlock, err := contract.DeployBlock(&bind.CallOpts{Context: ctx})
			if err != nil {
				return nil, fmt.Errorf("could not get deploy block: %w", err)
			}
			chainListener, err := listener.NewChainListener(chainClient, store, common.HexToAddress(*rfqAddrV1), uint64(startBlock.Int64()), metricHandler, listener.WithName("guard"))
			if err != nil {
				return nil, fmt.Errorf("could not get chain listener: %w", err)
			}
			listenersV1[chainID] = chainListener
			contractsV1[chainID] = contract
		}

		// setup v2
		rfqAddrV2, err := cfg.GetRFQAddressV2(chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get rfq address: %w", err)
		}
		chainClient, err := omniClient.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}
		contract, err := fastbridgev2.NewFastBridgeV2Ref(common.HexToAddress(rfqAddrV2), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create fast bridge contract: %w", err)
		}
		startBlock, err := contract.DeployBlock(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("could not get deploy block: %w", err)
		}
		chainListener, err := listener.NewChainListener(chainClient, store, common.HexToAddress(rfqAddrV2), uint64(startBlock.Int64()), metricHandler, listener.WithName("guard"))
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		listenersV2[chainID] = chainListener
		contractsV2[chainID] = contract
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
		cfg:          cfg,
		metrics:      metricHandler,
		db:           store,
		client:       omniClient,
		contractsV1:  contractsV1,
		contractsV2:  contractsV2,
		listenersV1:  listenersV1,
		listenersV2:  listenersV2,
		txSubmitter:  txSubmitter,
		otelRecorder: otelRecorder,
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
		//nolint: copyloopvar
		chainID := chainID // capture loop variable

		// only run v1 if it is set
		v1Addr, err := g.cfg.GetRFQAddressV1(chainID)
		if err != nil {
			return fmt.Errorf("could not get rfq address v1: %w", err)
		}
		if v1Addr != nil {
			group.Go(func() error {
				err := g.runChainIndexerV1(ctx, chainID)
				if err != nil {
					return fmt.Errorf("could not runChainIndexer chain indexer for chain %d [v1]: %w", chainID, err)
				}
				return nil
			})
		}

		group.Go(func() error {
			err := g.runChainIndexerV2(ctx, chainID)
			if err != nil {
				return fmt.Errorf("could not runChainIndexer chain indexer for chain %d [v2]: %w", chainID, err)
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
func (g Guard) runChainIndexerV1(ctx context.Context, chainID int) (err error) {
	chainListener := g.listenersV1[chainID]

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
			err = g.handleProofProvidedLog(ctx, event, chainID, log.Address)
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

//nolint:cyclop
func (g Guard) runChainIndexerV2(ctx context.Context, chainID int) (err error) {
	chainListener := g.listenersV2[chainID]

	parser, err := fastbridgev2.NewParser(chainListener.Address())
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
		case *fastbridgev2.FastBridgeV2BridgeRequested:
			v1Event := &fastbridge.FastBridgeBridgeRequested{
				TransactionId: event.TransactionId,
				Sender:        event.Sender,
				Request:       event.Request,
				DestChainId:   event.DestChainId,
				OriginToken:   event.OriginToken,
				DestToken:     event.DestToken,
				DestAmount:    event.DestAmount,
				SendChainGas:  event.SendChainGas,
				Raw:           event.Raw,
			}
			err = g.handleBridgeRequestedLog(ctx, v1Event, chainID)
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		// following events match ABIs exactly, so no need to differentiate
		case *fastbridgev2.FastBridgeV2BridgeProofProvided:
			v1Event := &fastbridge.FastBridgeBridgeProofProvided{
				TransactionId:   event.TransactionId,
				Relayer:         event.Relayer,
				TransactionHash: event.TransactionHash,
				Raw:             event.Raw,
			}
			err = g.handleProofProvidedLog(ctx, v1Event, chainID, log.Address)
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		case *fastbridgev2.FastBridgeV2BridgeProofDisputed:
			v1Event := &fastbridge.FastBridgeBridgeProofDisputed{
				TransactionId: event.TransactionId,
				Relayer:       event.Relayer,
				Raw:           event.Raw,
			}
			err = g.handleProofDisputedLog(ctx, v1Event)
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

func (g *Guard) isV2Address(chainID int, addr common.Address) bool {
	rfqAddr, err := g.cfg.GetRFQAddressV2(chainID)
	if err != nil {
		return false
	}
	return addr == common.HexToAddress(rfqAddr)
}
