package listener

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/db"
	"golang.org/x/sync/errgroup"
	"math/big"
	"time"
)

type ChainListener interface {
}

type chainListener struct {
	client   client.EVM
	contract *fastbridge.FastBridgeRef
	store    db.Service
	handler  metrics.Handler
	backoff  *backoff.Backoff
	// IMPORTANT! These fields cannot be used until they has been set. They are NOT
	// set in the constructor
	startBlock, chainID uint64
	pollInterval        time.Duration
	//latestBlock         uint64
}

var logger = log.Logger("chainlistener-logger")

func NewChainListener(ctx context.Context, omnirpcClient client.EVM, store db.Service, address common.Address, handler metrics.Handler) (ChainListener, error) {
	fastBridge, err := fastbridge.NewFastBridgeRef(address, omnirpcClient)
	if err != nil {
		return nil, fmt.Errorf("could not create fast bridge contract: %w", err)
	}

	return chainListener{
		handler:  handler,
		store:    store,
		client:   omnirpcClient,
		contract: fastBridge,
		backoff:  newBackoffConfig(),
	}, nil
}

// defaultPollInterval
const (
	// TODO: replace w/ config param if needed
	defaultPollInterval = 4
	maxGetLogsRange     = 2000
)

func (c *chainListener) Listen(ctx context.Context) (err error) {
	c.startBlock, c.chainID, err = c.getMetadata(ctx)
	if err != nil {
		return fmt.Errorf("could not get metadata: %w", err)
	}

	c.pollInterval = time.Duration(0)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled: %w", ctx.Err())
		case <-time.After(time.Second * c.pollInterval):
			err = c.doPoll(ctx)
			if err != nil {
				logger.Warn(err)
			}
		}

	}
}

func (c *chainListener) doPoll(parentCtx context.Context) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "getMetadata")
	c.pollInterval = defaultPollInterval

	// Note: in the case of an error, you don't have to handle the poll interval by calling b.duration.
	defer func() {
		metrics.EndSpanWithErr(span, err)
		if err != nil {
			c.pollInterval = c.backoff.Duration()
		}
	}()

	latestBlock, err := c.client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get block number: %w", err)
	}

	// Check if latest block is the same as start block (for chains with slow block times)

	if latestBlock == c.startBlock {
		return
	}

	// Handle if the listener is more than one get logs range behind the head
	// Note: this does not cover the edge case of a reorg that includes a new tx
	endBlock := latestBlock
	lastUnconfirmedBlock := latestBlock
	if c.startBlock+maxGetLogsRange < latestBlock {
		endBlock = c.startBlock + maxGetLogsRange
		// This will be used as the bottom of the range in the next iteration
		lastUnconfirmedBlock = endBlock
		c.pollInterval = 0
	}

	filterQuery := c.buildFilterQuery(c.startBlock, endBlock)
	logs, err := c.client.FilterLogs(ctx, filterQuery)
	if err != nil {
		return fmt.Errorf("could not filter logs: %w", err)
	}

	_ = logs // TODO: do something with logs

	c.startBlock = lastUnconfirmedBlock
	return nil
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

func (c chainListener) buildFilterQuery(fromBlock, toBlock uint64) ethereum.FilterQuery {
	return ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{c.contract.Address()},
		Topics:    [][]common.Hash{{fastbridge.BridgeRequestedTopic, fastbridge.BridgeRelayedTopic}},
	}
}
