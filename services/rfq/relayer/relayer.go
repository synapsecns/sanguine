package relayer

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/sqlite"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// Relayer is the core of the relayer application.
type Relayer struct {
	cfg     relconfig.Config
	metrics metrics.Handler
	db      reldb.Service
	client  omnirpcClient.RPCClient
}

var logger = log.Logger("relayer")

// NewRelayer creates a new relayer.
func NewRelayer(ctx context.Context, metricHandler metrics.Handler, cfg relconfig.Config) (*Relayer, error) {
	omniClient := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, metricHandler, omnirpcClient.WithCaptureReqRes())

	// TODO: pull from config
	store, err := sqlite.NewSqliteStore(ctx, cfg.DBConfig, metricHandler, false)
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}

	// TODO: add bd
	rel := Relayer{
		db:      store,
		client:  omniClient,
		metrics: metricHandler,
		cfg:     cfg,
	}
	return &rel, nil
}

func (r *Relayer) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := r.startChainParser(ctx)
		if err != nil {
			return fmt.Errorf("could not start chain parser: %w", err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start: %w", err)
	}

	return nil
}

func (r *Relayer) startChainParser(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// TODO: good chance we wanna prepare these chain listeners up front and then listen later.
	for chainID, bridgeStr := range r.cfg.Bridges {
		chainID := chainID //capture func literal
		// TODO: consider getter for this convert step
		bridge := common.HexToAddress(bridgeStr)
		chainClient, err := r.client.GetChainClient(ctx, chainID)
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		parser, err := fastbridge.NewParser(bridge)
		if err != nil {
			return fmt.Errorf("could not parse: %w", err)
		}

		chainListener, err := listener.NewChainListener(chainClient, r.db, bridge, r.metrics)
		if err != nil {
			return fmt.Errorf("could not get chain listener: %w", err)
		}

		g.Go(func() error {
			err = chainListener.Listen(ctx, func(ctx context.Context, log types.Log) error {
				_, parsedEvent, ok := parser.ParseEvent(log)
				// handle unknown event
				if !ok {
					if len(log.Topics) != 0 {
						logger.Warnf("unknown event %s", log.Topics[0])
					}
					return nil
				}

				switch event := parsedEvent.(type) {
				case *fastbridge.FastBridgeBridgeRequested:
					// TODO store this if not already seen
					err = r.handleRequest(ctx, event, uint64(chainID))
					if err != nil {
						return fmt.Errorf("could not handle request: %w", err)
					}
				case *fastbridge.FastBridgeBridgeRelayed:
					panic("implement me")
				}

				return nil
			})

			if err != nil {
				return fmt.Errorf("listener failed: %w", err)
			}
			return nil
		})
	}
	return nil
}

func (r *Relayer) handleRequest(parentCtx context.Context, req *fastbridge.FastBridgeBridgeRequested, chainID uint64) (err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "getDecimals", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(req.TransactionId[:])),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: consider a mapmutex
	_, err = r.db.GetQuoteRequestByID(ctx, req.TransactionId)
	// expect no results
	if !errors.Is(err, reldb.ErrNoQuoteForID) {
		// maybe a db err? if so error out & try again later
		if err != nil {
			return fmt.Errorf("could not call db: %w", err)
		}
	}

	// TODO: these should be premade
	originClient, err := r.client.GetChainClient(ctx, int(chainID))
	if err != nil {
		return fmt.Errorf("could not get correct omnirpc client: %w", err)
	}

	fastBridge, err := fastbridge.NewFastBridgeRef(req.Raw.Address, originClient)
	if err != nil {
		return fmt.Errorf("could not get correct fast bridge: %w", err)
	}

	bridgeTx, err := fastBridge.GetBridgeTransaction(&bind.CallOpts{Context: ctx}, req.Request)
	if err != nil {
		return fmt.Errorf("could not get bridge transaction: %w", err)
	}

	decimals, err := r.getDecimals(ctx, bridgeTx)
	if err != nil {
		return fmt.Errorf("could not get decimals: %w", err)
	}

	err = r.db.StoreQuoteRequest(ctx, reldb.QuoteRequest{
		OriginTokenDecimals: decimals.originDecimals,
		DestTokenDecimals:   decimals.originDecimals,
		TransactionId:       req.TransactionId,
		Sender:              req.Sender,
		Transaction:         bridgeTx,
		Status:              reldb.Seen,
	})
	if err != nil {
		return fmt.Errorf("could not get db: %w", err)
	}

	return nil
}

func (r *Relayer) getDecimals(parentCtx context.Context, bridgeTx fastbridge.IFastBridgeBridgeTransaction) (_ *decimalsRes, err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "getDecimals", trace.WithAttributes(
		attribute.String("sender", bridgeTx.OriginSender.String()),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: add a cache for decimals.
	res := decimalsRes{}

	// TODO: cleanup duplication, but keep paralellism.
	// this is  a bit of a pain since it deals w/ 3 different fields, but shouldn't take too long
	originClient, err := r.client.GetChainClient(ctx, int(bridgeTx.OriginChainId))
	if err != nil {
		return nil, fmt.Errorf("could not get origin client: %w", err)
	}

	destClient, err := r.client.GetChainClient(ctx, int(bridgeTx.DestChainId))
	if err != nil {
		return nil, fmt.Errorf("could not get dest client: %w", err)
	}

	originERC20, err := ierc20.NewIERC20(bridgeTx.OriginToken, originClient)
	if err != nil {
		return nil, fmt.Errorf("could not get origin token")
	}

	destERC20, err := ierc20.NewIERC20(bridgeTx.DestToken, destClient)
	if err != nil {
		return nil, fmt.Errorf("could not get dest token")
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		res.originDecimals, err = originERC20.Decimals(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("could not get dest decimals: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		res.destDecimals, err = destERC20.Decimals(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("could not get origin decimals: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get decimals: %w", err)
	}

	return &res, nil
}

type decimalsRes struct {
	originDecimals, destDecimals uint8
}
