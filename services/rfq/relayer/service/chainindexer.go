package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// startChainIndexers starts the chain indexers for each bridge.
// these listen on the chain using the chain listeners and then handle the events.
func (r *Relayer) startChainIndexers(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// TODO: good chance we wanna prepare these chain listeners up front and then listen later.
	for chainID := range r.cfg.Bridges {
		chainID := chainID // capture func literal

		g.Go(func() error {
			err := r.runChainIndexer(ctx, chainID)
			if err != nil {
				return fmt.Errorf("could not runChainIndexer chain indexer for chain %d: %w", chainID, err)
			}
			return nil
		})
	}
	return nil
}

// runChainIndexer runs the chain indexer for a given chain.
// any events that an action exists for are indexed.
func (r *Relayer) runChainIndexer(ctx context.Context, chainID int) (err error) {
	chainListener := r.chainListeners[chainID]

	parser, err := fastbridge.NewParser(chainListener.Address())
	if err != nil {
		return fmt.Errorf("could not parse: %w", err)
	}

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
			err = r.handleBridgeRequestedLog(ctx, event, uint64(chainID))
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		case *fastbridge.FastBridgeBridgeRelayed:
			err = r.handleRelayLog(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle relay: %w", err)
			}
		case *fastbridge.FastBridgeBridgeProofProvided:
			err = r.handleProofProvided(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle proof provided: %w", err)
			}
		case *fastbridge.FastBridgeBridgeDepositClaimed:
			err = r.handleDepositClaimed(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle deposit claimed: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("listener failed: %w", err)
	}
	return nil
}

// handleBridgeRequestedLog handles the BridgeRequestedLog event.
// Step: 1
//
// This is the first event emitted in the bridge process. It is emitted when a user calls bridge on chain.
// To process it, we decode the bridge transaction and store all the data, marking it as seen.
func (r *Relayer) handleBridgeRequestedLog(parentCtx context.Context, req *fastbridge.FastBridgeBridgeRequested, chainID uint64) (err error) {
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

	// TODO: you can just pull these out of inventory. If they don't exist mark as invalid.
	decimals, err := r.getDecimals(ctx, bridgeTx)
	if err != nil {
		return fmt.Errorf("could not get decimals: %w", err)
	}

	err = r.db.StoreQuoteRequest(ctx, reldb.QuoteRequest{
		BlockNumber:         req.Raw.BlockNumber,
		RawRequest:          req.Request,
		OriginTokenDecimals: decimals.originDecimals,
		DestTokenDecimals:   decimals.destDecimals,
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

func (r *Relayer) handleRelayLog(ctx context.Context, req *fastbridge.FastBridgeBridgeRelayed) (err error) {
	reqID, err := r.db.GetQuoteRequestByID(ctx, req.TransactionId)
	if err != nil {
		return fmt.Errorf("could not get quote request: %v", err)
	}
	// we might've accidentally gotten this later, if so we'll just ignore it
	if reqID.Status != reldb.RelayStarted {
		logger.Warnf("got relay log for request that was not relay started (transaction id: %s, txhash: %s)", hexutil.Encode(reqID.TransactionId[:]), req.Raw.TxHash)
		return nil
	}

	// TODO: this can still get re-orged
	err = r.db.UpdateQuoteRequestStatus(ctx, req.TransactionId, reldb.RelayCompleted)
	if err != nil {
		return fmt.Errorf("could not update request status: %v", err)
	}
	return nil
}

// getDecimals gets the decimals for the origin and dest tokens.
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

func (r *Relayer) handleDepositClaimed(ctx context.Context, event *fastbridge.FastBridgeBridgeDepositClaimed) error {
	err := r.db.UpdateQuoteRequestStatus(ctx, event.TransactionId, reldb.ClaimCompleted)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

func (r *Relayer) handleProofProvided(ctx context.Context, req *fastbridge.FastBridgeBridgeProofProvided) (err error) {
	// TODO: this can still get re-orged
	// ALso: we should make sure the previous status  is ProvePosting
	err = r.db.UpdateQuoteRequestStatus(ctx, req.TransactionId, reldb.ProvePosted)
	if err != nil {
		return fmt.Errorf("could not update request status: %v", err)
	}
	return nil
}
