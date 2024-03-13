package service

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
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
	for chainID := range r.cfg.GetChains() {
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
// nolint: cyclop
func (r *Relayer) runChainIndexer(ctx context.Context, chainID int) (err error) {
	chainListener := r.chainListeners[chainID]

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

		ctx, span := r.metrics.Tracer().Start(parentCtx, fmt.Sprintf("handleLog-%s", et), trace.WithAttributes(
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
			err = r.handleBridgeRequestedLog(ctx, event, uint64(chainID))
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		case *fastbridge.FastBridgeBridgeRelayed:
			// it wasn't me
			if event.Relayer != r.signer.Address() {
				//nolint: wrapcheck
				return r.db.UpdateQuoteRequestStatus(ctx, event.TransactionId, reldb.RelayRaceLost)
			}

			err = r.handleRelayLog(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle relay: %w", err)
			}
		case *fastbridge.FastBridgeBridgeProofProvided:
			// it wasn't me
			if event.Relayer != r.signer.Address() {
				//nolint: wrapcheck
				return r.db.UpdateQuoteRequestStatus(ctx, event.TransactionId, reldb.RelayRaceLost)
			}

			err = r.handleProofProvided(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle proof provided: %w", err)
			}
		case *fastbridge.FastBridgeBridgeDepositClaimed:
			// it wasn't me
			if event.Relayer != r.signer.Address() {
				//nolint: wrapcheck
				return r.db.UpdateQuoteRequestStatus(ctx, event.TransactionId, reldb.RelayRaceLost)
			}

			err = r.handleDepositClaimed(ctx, event, chainID)
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

const ethDecimals = 18

// getDecimals gets the decimals for the origin and dest tokens.
func (r *Relayer) getDecimals(parentCtx context.Context, bridgeTx fastbridge.IFastBridgeBridgeTransaction) (_ *decimalsRes, err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "getDecimals", trace.WithAttributes(
		attribute.String("sender", bridgeTx.OriginSender.String()),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: add a cache for decimals.
	var originERC20, destERC20 *ierc20.IERC20
	res := decimalsRes{}

	if bridgeTx.OriginToken == chain.EthAddress {
		res.originDecimals = ethDecimals
	} else {
		// TODO: cleanup duplication, but keep paralellism.
		// this is  a bit of a pain since it deals w/ 3 different fields, but shouldn't take too long
		originClient, err := r.client.GetChainClient(ctx, int(bridgeTx.OriginChainId))
		if err != nil {
			return nil, fmt.Errorf("could not get origin client: %w", err)
		}
		originERC20, err = ierc20.NewIERC20(bridgeTx.OriginToken, originClient)
		if err != nil {
			return nil, fmt.Errorf("could not get origin token")
		}
	}

	if bridgeTx.DestToken == chain.EthAddress {
		res.destDecimals = ethDecimals
	} else {
		destClient, err := r.client.GetChainClient(ctx, int(bridgeTx.DestChainId))
		if err != nil {
			return nil, fmt.Errorf("could not get dest client: %w", err)
		}
		destERC20, err = ierc20.NewIERC20(bridgeTx.DestToken, destClient)
		if err != nil {
			return nil, fmt.Errorf("could not get dest token")
		}
	}

	// return early if both dest and origin are ETH.
	if originERC20 == nil && destERC20 == nil {
		return &res, nil
	}

	// do rpc calls to fetch the erc20 decimals.
	g, ctx := errgroup.WithContext(ctx)
	if originERC20 != nil {
		g.Go(func() error {
			res.originDecimals, err = originERC20.Decimals(&bind.CallOpts{Context: ctx})
			if err != nil {
				return fmt.Errorf("could not get dest decimals: %w", err)
			}
			return nil
		})
	}

	if destERC20 != nil {
		g.Go(func() error {
			res.destDecimals, err = destERC20.Decimals(&bind.CallOpts{Context: ctx})
			if err != nil {
				return fmt.Errorf("could not get origin decimals: %w", err)
			}
			return nil
		})
	}

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get decimals: %w", err)
	}

	return &res, nil
}

type decimalsRes struct {
	originDecimals, destDecimals uint8
}

func (r *Relayer) handleDepositClaimed(ctx context.Context, event *fastbridge.FastBridgeBridgeDepositClaimed, chainID int) error {
	err := r.inventory.Rebalance(ctx, chainID, event.Token)
	if err != nil {
		return fmt.Errorf("could not rebalance: %w", err)
	}
	err = r.db.UpdateQuoteRequestStatus(ctx, event.TransactionId, reldb.ClaimCompleted)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}
