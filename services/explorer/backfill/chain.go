package backfill

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"golang.org/x/sync/errgroup"
	"time"
)

// ChainBackfiller is an explorer backfiller for a chain.
type ChainBackfiller struct {
	// consumerDB is the database that the backfiller will use to store the events.
	consumerDB db.ConsumerDB
	// bridgeParser is the parser to use to parse bridge events.
	bridgeParser *parser.BridgeParser
	// swapParsers is a map from contract address -> parser.
	swapParsers map[common.Address]*parser.SwapParser
	// Fetcher is the Fetcher to use to fetch logs.
	Fetcher fetcher.ScribeFetcher
	// chainConfig is the chain config for the chain.
	chainConfig config.ChainConfig
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(consumerDB db.ConsumerDB, bridgeParser *parser.BridgeParser, swapParsers map[common.Address]*parser.SwapParser, fetcher fetcher.ScribeFetcher, chainConfig config.ChainConfig) *ChainBackfiller {
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
	for i := range c.chainConfig.Contracts {
		contract := c.chainConfig.Contracts[i]
		g, groupCtx := errgroup.WithContext(ctx)
		g.SetLimit(c.chainConfig.MaxGoroutines)
		startHeight := uint64(contract.StartBlock)

		// Set start block to -1 to trigger backfill from last block stored by explorer,
		// otherwise backfilling will begin at the block number specified in the config file.
		if contract.StartBlock < 0 {
			startHeight, err = c.consumerDB.GetUint64(ctx, fmt.Sprintf(
				"SELECT ifNull(%s, 0) FROM last_blocks WHERE %s = %d",
				sql.BlockNumberFieldName, sql.ChainIDFieldName, c.chainConfig.ChainID,
			))
			if err != nil {
				return fmt.Errorf("could not get last block number: %w", err)
			}
		}

		endHeight, err := c.Fetcher.FetchLastIndexed(ctx, c.chainConfig.ChainID, contract.Address)
		if err != nil {
			return fmt.Errorf("could not get last indexed for contract %s: %w", contract.Address, err)
		}

		// Iterate over all blocks and fetch logs with the current contract address.
		for currentHeight := startHeight; currentHeight < endHeight; currentHeight += c.chainConfig.FetchBlockIncrement {
			funcHeight := currentHeight

			g.Go(func() error {
				b := &backoff.Backoff{
					Factor: 2,
					Jitter: true,
					Min:    1 * time.Second,
					Max:    30 * time.Second,
				}

				timeout := time.Duration(0)

				for {
					select {
					case <-groupCtx.Done():
						return fmt.Errorf("context canceled: %w", groupCtx.Err())
					case <-time.After(timeout):
						rangeEnd := funcHeight + c.chainConfig.FetchBlockIncrement - 1

						if rangeEnd > endHeight {
							rangeEnd = endHeight
						}

						// Fetch the logs from Scribe.
						logs, err := c.Fetcher.FetchLogsInRange(groupCtx, c.chainConfig.ChainID, funcHeight, rangeEnd, common.HexToAddress(contract.Address))
						if err != nil {
							timeout = b.Duration()
							logger.Warnf("could not fetch logs for chain %d: %s. Retrying in %s", c.chainConfig.ChainID, err, timeout)

							continue
						}

						var eventParser parser.Parser

						switch contract.ContractType {
						case "bridge":
							eventParser = c.bridgeParser
						case "swap":
							eventParser = c.swapParsers[common.HexToAddress(contract.Address)]
						}

						err = c.processLogs(groupCtx, logs, eventParser)
						if err != nil {
							logger.Warnf("could not process logs for chain %d: %s", c.chainConfig.ChainID, err)
						}

						return nil
					}
				}
			})
		}

		if err := g.Wait(); err != nil {
			return fmt.Errorf("error while backfilling chain %d: %w", c.chainConfig.ChainID, err)
		}
	}

	return nil
}

// processLogs processes the logs and stores them in the consumer database.
//
//nolint:gocognit,cyclop
func (c *ChainBackfiller) processLogs(ctx context.Context, logs []ethTypes.Log, eventParser parser.Parser) error {
	for i := range logs {
		err := eventParser.ParseAndStore(ctx, logs[i], c.chainConfig.ChainID)
		if err != nil {
			return fmt.Errorf("could not parse and store log: %w", err)
		}

		// TODO this can be moved out of this for loop once log order is guaranteed
		// Store the last block in clickhouse
		err = c.consumerDB.StoreLastBlock(ctx, c.chainConfig.ChainID, logs[i].BlockNumber)
		if err != nil {
			logger.Warnf("could not store last block for chain %d: %s", c.chainConfig.ChainID, err)
		}
	}

	return nil
}
