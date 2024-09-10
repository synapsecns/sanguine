package listener

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	listenerDB "github.com/synapsecns/sanguine/ethergo/listener/db"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ipfs/go-log"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ContractListener listens for chain events and calls HandleLog.
//
//go:generate go run github.com/vektra/mockery/v2 --name ContractListener --output ./mocks --case=underscore
type ContractListener interface {
	// Listen starts the listener and call HandleLog for each event
	Listen(ctx context.Context, handler HandleLog) error
	// LatestBlock gets the last recorded latest block from the rpc.
	// this is NOT last indexed. It is provided as a helper for checking confirmation count
	LatestBlock() uint64
	// Address gets the address of the contract this listener is listening to
	Address() common.Address
}

// HandleLog is the handler for a log event
// in the event this errors, the range will be reparsed.
type HandleLog func(ctx context.Context, log types.Log) error

type chainListener struct {
	client       client.EVM
	address      common.Address
	initialBlock uint64
	store        listenerDB.ChainListenerDB
	handler      metrics.Handler
	backoff      *backoff.Backoff
	// IMPORTANT! These fields cannot be used until they has been set. They are NOT
	// set in the constructor
	startBlock, chainID, latestBlock  uint64
	pollInterval, pollIntervalSetting time.Duration
	// newBlockHandler is an optional handler that is called when a new block is detected.
	newBlockHandler NewBlockHandler
	finalityMode    rpc.BlockNumber
	blockWait       uint64
	// otelRecorder is the recorder for the otel metrics.
	otelRecorder iOtelRecorder
	name         string
}

var (
	logger = log.Logger("chainlistener-logger")
	// ErrNoLatestBlockForChainID is returned when no block exists for the chain.
	ErrNoLatestBlockForChainID = listenerDB.ErrNoLatestBlockForChainID
)

// NewChainListener creates a new chain listener.
func NewChainListener(omnirpcClient client.EVM, store listenerDB.ChainListenerDB, address common.Address, initialBlock uint64, handler metrics.Handler, options ...Option) (ContractListener, error) {
	c := &chainListener{
		handler:             handler,
		address:             address,
		initialBlock:        initialBlock,
		store:               store,
		client:              omnirpcClient,
		backoff:             newBackoffConfig(),
		pollIntervalSetting: time.Millisecond * 50,
		finalityMode:        rpc.LatestBlockNumber,
		blockWait:           0,
	}

	for _, option := range options {
		option(c)
	}

	return c, nil
}

// defaultPollInterval.
const (
	maxGetLogsRange = 2000
)

func (c *chainListener) Listen(ctx context.Context, handler HandleLog) (err error) {
	c.startBlock, c.chainID, err = c.getMetadata(ctx)
	if err != nil {
		return fmt.Errorf("could not get metadata: %w", err)
	}

	if c.otelRecorder == nil {
		c.otelRecorder, err = newOtelRecorder(c.handler, int(c.chainID), c.name)
		if err != nil {
			return fmt.Errorf("could not create otel recorder: %w", err)
		}
	}

	c.pollInterval = time.Duration(0)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(c.pollInterval):
			err = c.doPoll(ctx, handler)
			if err != nil {
				logger.Warn(err)
			}
		}
	}
}

func (c *chainListener) Address() common.Address {
	return c.address
}

func (c *chainListener) LatestBlock() uint64 {
	return c.latestBlock
}

// nolint: cyclop
func (c *chainListener) doPoll(parentCtx context.Context, handler HandleLog) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "doPoll", trace.WithAttributes(attribute.Int(metrics.ChainID, int(c.chainID))))
	c.pollInterval = c.pollIntervalSetting

	// Note: in the case of an error, you don't have to handle the poll interval by calling b.duration.
	var endBlock uint64
	defer func() {
		span.SetAttributes(
			attribute.Int64("start_block", int64(c.startBlock)),
			attribute.Int64("end_block", int64(endBlock)),
			attribute.Int64("latest_block", int64(c.latestBlock)),
		)
		metrics.EndSpanWithErr(span, err)
		if err != nil {
			c.backoff.Attempt()
		} else {
			c.backoff.Reset()
		}
		c.pollInterval = c.backoff.Duration() + c.pollIntervalSetting
	}()

	oldLatestBlock := c.latestBlock
	c.latestBlock, err = c.getBlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get block number: %w", err)
	}

	// Check if latest block is the same as start block (for chains with slow block times)
	didPoll := true
	defer func() {
		span.SetAttributes(attribute.Bool("did_poll", didPoll))
	}()
	if c.latestBlock == c.startBlock {
		didPoll = false
		return nil
	}

	// check if there's a new block
	if c.newBlockHandler != nil && c.latestBlock != oldLatestBlock {
		err = c.newBlockHandler(ctx, c.latestBlock)
		if err != nil {
			return fmt.Errorf("new block handler failed: %w", err)
		}
	}

	// Handle if the listener is more than one get logs range behind the head
	// Note: this does not cover the edge case of a reorg that includes a new tx
	endBlock = c.latestBlock
	lastUnconfirmedBlock := c.latestBlock
	if c.startBlock+maxGetLogsRange < c.latestBlock {
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

	for _, newLog := range logs {
		err = handler(ctx, newLog)
		if err != nil {
			return fmt.Errorf("handle log failed, will reparse: %w", err)
		}
	}

	err = c.store.PutLatestBlock(ctx, c.chainID, endBlock)
	if err != nil {
		return fmt.Errorf("could not put latest block: %w", err)
	}
	c.otelRecorder.RecordLastBlock(endBlock)

	c.startBlock = lastUnconfirmedBlock
	return nil
}

func (c chainListener) getBlockNumber(ctx context.Context) (uint64, error) {
	block, err := c.client.HeaderByNumber(ctx, big.NewInt(c.finalityMode.Int64()))
	if err != nil {
		return 0, fmt.Errorf("could not get block by number: %w", err)
	}

	blockNumber := block.Number

	if c.blockWait > 0 {
		blockNumber.Sub(blockNumber, big.NewInt(int64(c.blockWait)))
	}

	return blockNumber.Uint64(), nil
}

func (c chainListener) getMetadata(parentCtx context.Context) (startBlock, chainID uint64, err error) {
	var lastIndexed uint64
	ctx, span := c.handler.Tracer().Start(parentCtx, "getMetadata")

	defer func() {
		span.SetAttributes(
			attribute.Int64("start_block", int64(startBlock)),
			attribute.Int64("last_indexed", int64(lastIndexed)),
			attribute.Int(metrics.ChainID, int(chainID)),
		)
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: one thing I've been going back and forth on is whether or not this method should be chain aware
	// passing in the chain ID would allow us to pull everything directly from the config, but be less testable
	// for now, this is probably the best solution for testability, but it's certainly a bit annoying we need to do
	// an rpc call in order to get the chain id
	//
	rpcChainID, err := c.client.ChainID(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("could not get chain ID: %w", err)
	}
	chainID = rpcChainID.Uint64()

	lastIndexed, err = c.getLastIndexed(ctx, chainID)
	if err != nil {
		return 0, 0, fmt.Errorf("could not get last indexed: %w", err)
	}

	if lastIndexed > c.startBlock {
		startBlock = lastIndexed
	} else {
		startBlock = c.initialBlock
	}

	return startBlock, chainID, nil
}

// TODO: consider some kind of backoff here in case rpcs are down at boot.
// this becomes more of an issue as we add more chains.
func (c chainListener) getLastIndexed(ctx context.Context, chainID uint64) (lastIndexed uint64, err error) {
	lastIndexed, err = c.store.LatestBlockForChain(ctx, chainID)
	// Workaround: TODO remove
	if errors.Is(err, ErrNoLatestBlockForChainID) || err != nil && err.Error() == ErrNoLatestBlockForChainID.Error() {
		// TODO: consider making this negative 1, requires type change
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("could not get the latest block for chainID: %w", err)
	}
	return lastIndexed, nil
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
		Addresses: []common.Address{c.address},
	}
}
