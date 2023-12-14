package listener

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/db"
	"golang.org/x/sync/errgroup"
	"time"
)

type ChainListener interface {
}

type chainListener struct {
	address  common.Address
	client   client.EVM
	contract *fastbridge.FastBridgeRef
	store    db.Service
	handler  metrics.Handler
}

func NewChainListener(ctx context.Context, omnirpcClient client.EVM, store db.Service, address common.Address, handler metrics.Handler) (ChainListener, error) {
	fastBridge, err := fastbridge.NewFastBridgeRef(address, omnirpcClient)
	if err != nil {
		return nil, fmt.Errorf("could not create fast bridge contract: %w", err)
	}

	return chainListener{
		address:  address,
		handler:  handler,
		store:    store,
		client:   omnirpcClient,
		contract: fastBridge,
	}, nil
}

func (c chainListener) Listen(ctx context.Context) (err error) {
	startBlock, chainID, err := c.getMetadata(ctx)
	if err != nil {
		return fmt.Errorf("could not get metadata: %w", err)
	}

	b := newBackoffConfig()
	_ = startBlock
	_ = b
	_ = chainID

	// defaultPollInterval
	// TODO: replace w/ config param if needed
	const defaultPollInterval = 4

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled: %w", ctx.Err())
		}

	}
}

func (c chainListener) getMetadata(parentCtx context.Context) (startBlock, chainID uint64, err error) {
	var deployBlock, lastIndexed uint64
	ctx, span := c.handler.Tracer().Start(parentCtx, "getMetadata")

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: consider some kind of backoff here in case rpcs are down at boot.
	// this becomes more of an issue as we add more chains
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		deployBlock, err := c.contract.DeployBlock(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("could not get deploy block: %w", err)
		}

		startBlock = deployBlock.Uint64()
		return nil
	})

	g.Go(func() error {
		// TODO: one thing I've been going back and forth on is whether or not this method should be chain aware
		// passing in the chain ID would allow us to pull everything directly from the config, but be less testable
		// for now, this is probably the best solution for testability, but it's certainly a bit annoying we need to do
		// an rpc call in order to get the chain id
		//
		rpcChainID, err := c.client.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("could not get chain ID: %w", err)
		}
		chainID = rpcChainID.Uint64()

		lastIndexed, err = c.store.LatestBlockForChain(ctx, chainID)
		if errors.Is(err, db.ErrNoLatestBlockForChainID) {
			// TODO: consider making this negative 1, requires type change
			lastIndexed = 0
			return nil
		}
		if err != nil {
			return fmt.Errorf("could not get the latest block for chainID: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return 0, 0, fmt.Errorf("could not get metadata: %w", err)
	}

	if lastIndexed > deployBlock {
		startBlock = lastIndexed
	}

	return startBlock, chainID, nil
}

func newBackoffConfig() *backoff.Backoff {
	return &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    10 * time.Millisecond,
		Max:    1 * time.Second,
	}
}
