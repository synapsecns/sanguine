package backfill

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"golang.org/x/sync/errgroup"
	"time"
)

// ChainBackfiller is an explorer backfiller for a chain.
type ChainBackfiller struct {
	// consumerDB is the database that the backfiller will use to store the events.
	consumerDB db.ConsumerDB
	// bridgeParser is the parser to use to parse bridge events.
	bridgeParser *consumer.BridgeParser
	// swapParsers is a map from contract address -> parser.
	swapParsers map[common.Address]*consumer.SwapParser
	// Fetcher is the Fetcher to use to fetch logs.
	Fetcher consumer.Fetcher
	// chainConfig is the chain config for the chain.
	chainConfig config.ChainConfig
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(consumerDB db.ConsumerDB, bridgeParser *consumer.BridgeParser, swapParsers map[common.Address]*consumer.SwapParser, fetcher consumer.Fetcher, chainConfig config.ChainConfig) *ChainBackfiller {
	return &ChainBackfiller{
		consumerDB:   consumerDB,
		bridgeParser: bridgeParser,
		swapParsers:  swapParsers,
		Fetcher:      fetcher,
		chainConfig:  chainConfig,
	}
}

// Backfill fetches logs from the GraphQL database, parses them, and stores them in the consumer database.
// nolint:cyclop,gocognit
func (c *ChainBackfiller) Backfill(ctx context.Context) (err error) {
	// initialize the errgroup
	g, groupCtx := errgroup.WithContext(ctx)
	startHeight := c.chainConfig.StartBlock

	// Set StartFromLastBlockStored to true to trigger backfill from last block store by explorer
	// Otherwise it will start at the block number specified in the config file.
	if c.chainConfig.StartFromLastBlockStored {
		startHeight, err = c.consumerDB.RetrieveLastBlock(ctx, c.chainConfig.ChainID)
		if err != nil {
			return fmt.Errorf("could not get last block number: %w", err)
		}
	}

	// For testing the max block height will be 12 (review this)

	endHeight, err := c.Fetcher.FetchLastBlock(ctx, c.chainConfig.ChainID)
	if err != nil {
		return fmt.Errorf("could not fetch last block: %w", err)
	}

	for currentHeight := startHeight; currentHeight <= endHeight; currentHeight += uint64(c.chainConfig.FetchBlockIncrement) {
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
					rangeEnd := funcHeight + uint64(c.chainConfig.FetchBlockIncrement) - 1
					if rangeEnd > endHeight {
						rangeEnd = endHeight
					}

					logs, err := c.Fetcher.FetchLogsInRange(groupCtx, c.chainConfig.ChainID, funcHeight, rangeEnd)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not fetch logs for chain %d: %s. Retrying in %s", c.chainConfig.ChainID, err, timeout)
						continue
					}
					// parse and store the logs
					err = c.processLogs(groupCtx, logs)
					if err != nil {
						logger.Warnf("could not process logs for chain %d: %s", c.chainConfig.ChainID, err)
					}

					// Store the last block
					err = c.consumerDB.StoreLastBlock(groupCtx, c.chainConfig.ChainID, rangeEnd)
					if err != nil {
						logger.Warnf("could not store last block for chain %d: %s", c.chainConfig.ChainID, err)
					}
					return nil
				}
			}
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("error while backfilling chain %d: %w", c.chainConfig.ChainID, err)
	}

	return nil
}

// processLogs processes the logs and stores them in the consumer database.
//
//nolint:gocognit,cyclop
func (c *ChainBackfiller) processLogs(ctx context.Context, logs []ethTypes.Log) error {
	// initialize the errgroup
	g, groupCtx := errgroup.WithContext(ctx)
	for _, log := range logs {
		log := log
		g.Go(func() error {
			var eventParser consumer.Parser
			if log.Address == common.HexToAddress(c.chainConfig.SynapseBridgeAddress) {
				eventParser = c.bridgeParser
			} else {
				if c.swapParsers[log.Address] == nil {
					logger.Warnf("no parser found for contract %s", log.Address.Hex())
					return nil
				}
				eventParser = c.swapParsers[log.Address]
			}

			err := eventParser.ParseAndStore(groupCtx, log, c.chainConfig.ChainID)
			if err != nil {
				return fmt.Errorf("could not parse and store log: %w", err)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error while processing logs: %w", err)
	}
	return nil
}
