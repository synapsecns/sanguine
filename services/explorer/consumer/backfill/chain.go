package backfill

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"
	"time"
)

// ChainBackfiller is an explorer backfiller for a chain.
type ChainBackfiller struct {
	// consumerDB is the database that the backfiller will use to store the events.
	consumerDB db.ConsumerDB
	// chainID is the chain ID of the chain to backfill.
	chainID uint32
	// fetchBlockIncrement is the number of blocks to fetch at a time.
	fetchBlockIncrement uint64
	// bridgeParser is the parser to use to parse bridge events.
	bridgeParser *consumer.BridgeParser
	// bridgeAddress is the address of the bridge contract.
	bridgeAddress common.Address
	// swapParsers is a map from contract address -> parser.
	swapParsers map[common.Address]*consumer.SwapParser
	// bridgeConfigAddress is the address of the BridgeConfigV3 contract.
	bridgeConfigAddress common.Address
	// fetcher is the fetcher to use to fetch logs.
	fetcher consumer.Fetcher
	// FailedLogs is how many logs failed to parse.
	FailedLogs atomic.Uint32
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(chainID uint32, consumerDB db.ConsumerDB, fetchBlockIncrement uint64, bridgeParser *consumer.BridgeParser, bridgeAddress common.Address, swapParsers map[common.Address]*consumer.SwapParser, fetcher consumer.Fetcher, bridgeConfigAddress common.Address) *ChainBackfiller {
	failedLogs := atomic.Uint32{}
	failedLogs.Store(0)
	return &ChainBackfiller{
		consumerDB:          consumerDB,
		chainID:             chainID,
		fetchBlockIncrement: fetchBlockIncrement,
		bridgeParser:        bridgeParser,
		bridgeAddress:       bridgeAddress,
		swapParsers:         swapParsers,
		fetcher:             fetcher,
		bridgeConfigAddress: bridgeConfigAddress,
		FailedLogs:          failedLogs,
	}
}

// Backfill fetches logs from the GraphQL database, parses them, and stores them in the consumer database.
func (c *ChainBackfiller) Backfill(ctx context.Context, startHeight, endHeight uint64) error {
	// initialize the errgroup
	g, groupCtx := errgroup.WithContext(ctx)
	for currentHeight := startHeight; currentHeight <= endHeight; currentHeight += c.fetchBlockIncrement {
		funcHeight := currentHeight
		g.Go(func() error {
			// backoff in the case of an error
			b := &backoff.Backoff{
				Factor: 2,
				Jitter: true,
				Min:    1 * time.Second,
				Max:    30 * time.Second,
			}
			// timeout should always be 0 on the first attempt
			timeout := time.Duration(0)
			for {
				// TODO: add a notification for failure to parse and store an event
				select {
				case <-groupCtx.Done():
					return fmt.Errorf("context canceled: %w", groupCtx.Err())
				case <-time.After(timeout):
					// fetch the logs
					rangeEnd := funcHeight + c.fetchBlockIncrement - 1
					if rangeEnd > endHeight {
						rangeEnd = endHeight
					}
					logs, err := c.fetcher.FetchLogsInRange(groupCtx, c.chainID, funcHeight, rangeEnd)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not fetch logs for chain %d: %s. Retrying in %s", c.chainID, err, timeout)
						continue
					}
					// parse and store the logs
					err = c.processLogs(groupCtx, logs, false)
					if err != nil {
						logger.Warnf("could not process logs for chain %d: %s", c.chainID, err)
					}
					return nil
				}
			}
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("error while backfilling chain %d: %w", c.chainID, err)
	}
	err := c.consumerDB.StoreLastLoggedBlock(ctx, c.chainID, endHeight)
	if err != nil {
		return fmt.Errorf("could not store last confirmed block for chain %d: %w", c.chainID, err)
	}
	return nil
}

func (c *ChainBackfiller) RetryFailedLogs(ctx context.Context) error {
	// get failed logs
	failedLogs, err := c.consumerDB.RetrieveFailedLogs(ctx, c.chainID)
	if err != nil {
		return fmt.Errorf("could not retrieve failed logs for chain %d: %w", c.chainID, err)
	}
	// process the logs
	err = c.processLogs(ctx, failedLogs, true)
	if err != nil {
		return fmt.Errorf("could not process failed logs for chain %d: %w", c.chainID, err)
	}
	return nil
}

// processLogs processes the logs and stores them in the consumer database. If `retry` is true,
// then the logs that are being processed are logs that failed to parse and store the first time,
// and they should be removed from the failed logs table if they are now stored correctly.
func (c *ChainBackfiller) processLogs(ctx context.Context, logs []ethTypes.Log, retry bool) error {
	// initialize the errgroup
	g, groupCtx := errgroup.WithContext(ctx)
	for _, log := range logs {
		log := log
		g.Go(func() error {
			var eventParser consumer.Parser
			if log.Address == c.bridgeAddress {
				eventParser = c.bridgeParser
			} else {
				if c.swapParsers[log.Address] == nil {
					logger.Warnf("no parser found for contract %s", log.Address.Hex())
					return nil
				}
				eventParser = c.swapParsers[log.Address]
			}
			err := eventParser.ParseAndStore(groupCtx, log, c.chainID)
			if err != nil {
				if !retry {
					err = c.consumerDB.StoreFailedLog(ctx, log, c.chainID)
					if err != nil {
						return fmt.Errorf("could not store failed log: %w", err)
					}
					c.FailedLogs.Inc()
				} else {
					logger.Warnf("failed to retry storing the failed log: %s", err)
				}
			}
			if err == nil && retry {
				err = c.consumerDB.DeleteFailedLog(ctx, log, c.chainID)
				if err != nil {
					return fmt.Errorf("could not delete failed log: %w", err)
				}
				c.FailedLogs.Dec()
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error while processing logs: %w", err)
	}
	return nil
}
