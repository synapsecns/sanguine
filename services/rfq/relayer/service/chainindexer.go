package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/contracts/ierc20"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/util"
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
			// blocking lock on the txid mutex to ensure state transitions are not overrwitten
			unlocker := r.handlerMtx.Lock(hexutil.Encode(event.TransactionId[:]))
			defer unlocker.Unlock()

			// it wasn't me
			if event.Relayer != r.signer.Address() {
				//nolint: wrapcheck
				return r.setRelayRaceLost(ctx, event.TransactionId)
			}

			err = r.handleRelayLog(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle relay: %w", err)
			}
		case *fastbridge.FastBridgeBridgeProofProvided:
			unlocker := r.handlerMtx.Lock(hexutil.Encode(event.TransactionId[:]))
			defer unlocker.Unlock()

			// it wasn't me
			if event.Relayer != r.signer.Address() {
				//nolint: wrapcheck
				return r.setRelayRaceLost(ctx, event.TransactionId)
			}

			err = r.handleProofProvided(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle proof provided: %w", err)
			}
		case *fastbridge.FastBridgeBridgeDepositClaimed:
			unlocker := r.handlerMtx.Lock(hexutil.Encode(event.TransactionId[:]))
			defer unlocker.Unlock()

			// it wasn't me
			if event.Relayer != r.signer.Address() {
				//nolint: wrapcheck
				return r.setRelayRaceLost(ctx, event.TransactionId)
			}

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

var ethDecimals uint8 = 18

// getDecimals gets the decimals for the origin and dest tokens.
func (r *Relayer) getDecimalsFromBridgeTx(parentCtx context.Context, bridgeTx fastbridge.IFastBridgeBridgeTransaction) (originDecimals *uint8, destDecimals *uint8, err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "getDecimals", trace.WithAttributes(
		attribute.String("sender", bridgeTx.OriginSender.String()),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// fetch the token decimals in parallel
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		originDecimals, err = r.getDecimals(ctx, bridgeTx.OriginToken, bridgeTx.OriginChainId)
		if err != nil {
			return fmt.Errorf("could not get origin decimals: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		destDecimals, err = r.getDecimals(ctx, bridgeTx.DestToken, bridgeTx.DestChainId)
		if err != nil {
			return fmt.Errorf("could not get dest decimals: %w", err)
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return nil, nil, fmt.Errorf("could not get decimals: %w", err)
	}

	return originDecimals, destDecimals, nil
}

// getDecimals gets the decimals for a token on a chain.
// It will attempt to load a result from the cache first.
// The cache will be updated if the result is fetched from RPC.
func (r *Relayer) getDecimals(ctx context.Context, addr common.Address, chainID uint32) (decimals *uint8, err error) {
	// attempt to load decimal from cache
	key := getDecimalsKey(addr, chainID)
	decimals, ok := r.decimalsCache.Load(key)
	if ok {
		return decimals, nil
	}

	if addr == util.EthAddress {
		return &ethDecimals, nil
	}

	// fetch decimals from RPC
	client, err := r.client.GetChainClient(ctx, int(chainID))
	if err != nil {
		return nil, fmt.Errorf("could not get client for chain %d: %w", chainID, err)
	}
	erc20, err := ierc20.NewIERC20(addr, client)
	if err != nil {
		return nil, fmt.Errorf("could not get token at %s: %w", addr.String(), err)
	}
	dec, err := erc20.Decimals(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("could not get decimals: %w", err)
	}

	// update the cache
	r.decimalsCache.Store(key, &dec)
	return &dec, nil
}

func getDecimalsKey(addr common.Address, chainID uint32) string {
	return fmt.Sprintf("%s-%d", addr.Hex(), chainID)
}

func (r *Relayer) handleDepositClaimed(ctx context.Context, event *fastbridge.FastBridgeBridgeDepositClaimed) error {
	err := r.db.UpdateQuoteRequestStatus(ctx, event.TransactionId, reldb.ClaimCompleted, nil)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

func (r *Relayer) setRelayRaceLost(ctx context.Context, transactionID [32]byte) error {
	err := r.db.UpdateQuoteRequestStatus(ctx, transactionID, reldb.RelayRaceLost, nil)
	// quote does not exist, no need to update status
	if err != nil && (errors.Is(err, reldb.ErrNoQuoteForID) || strings.Contains(err.Error(), reldb.ErrNoQuoteForID.Error())) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("could not set relay race lost: %w", err)
	}
	return nil
}
