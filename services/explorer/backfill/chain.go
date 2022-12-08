package backfill

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"golang.org/x/sync/errgroup"
)

// ChainBackfiller is an explorer backfiller for a chain.
type ChainBackfiller struct {
	// consumerDB is the database that the backfiller will use to store the events.
	consumerDB db.ConsumerDB
	// bridgeParser is the parser to use to parse bridge events.
	bridgeParser *parser.BridgeParser
	// swapParsers is a map from contract address -> parser.
	swapParsers map[common.Address]*parser.SwapParser
	// messageBusParser is the parser to use to parse message bus events.
	messageBusParser *parser.MessageBusParser
	// Fetcher is the Fetcher to use to fetch logs.
	Fetcher fetcher.ScribeFetcher
	// chainConfig is the chain config for the chain.
	chainConfig config.ChainConfig
}

type contextKey string

const (
	chainKey contextKey = "chainID"
)

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(consumerDB db.ConsumerDB, bridgeParser *parser.BridgeParser, swapParsers map[common.Address]*parser.SwapParser, messageBusParser *parser.MessageBusParser, fetcher fetcher.ScribeFetcher, chainConfig config.ChainConfig) *ChainBackfiller {
	return &ChainBackfiller{
		consumerDB:       consumerDB,
		bridgeParser:     bridgeParser,
		swapParsers:      swapParsers,
		messageBusParser: messageBusParser,
		Fetcher:          fetcher,
		chainConfig:      chainConfig,
	}
}

// Backfill fetches logs from the GraphQL database, parses them, and stores them in the consumer database.
// nolint:cyclop,gocognit
func (c *ChainBackfiller) Backfill(ctx context.Context) (err error) {
	chainCtx := context.WithValue(ctx, chainKey, fmt.Sprintf("%d", c.chainConfig.ChainID))
	contractsGroup, _ := errgroup.WithContext(chainCtx)

	for i := range c.chainConfig.Contracts {
		contract := c.chainConfig.Contracts[i]
		var eventParser parser.Parser

		switch contract.ContractType {
		case "bridge":
			eventParser = c.bridgeParser
		case "swap":
			eventParser = c.swapParsers[common.HexToAddress(contract.Address)]
		case "messagebus":
			eventParser = c.messageBusParser
		}
		contractsGroup.Go(func() error {
			g, groupCtx := errgroup.WithContext(chainCtx)

			// Create a new context for the chain so all chains don't halt when backfilling is completed.
			g.SetLimit(c.chainConfig.MaxGoroutines)
			startHeight := uint64(contract.StartBlock)

			// Set start block to -1 to trigger backfill from last block stored by explorer,
			// otherwise backfilling will begin at the block number specified in the config file.
			if contract.StartBlock < 0 {
				startHeight, err = c.consumerDB.GetUint64(ctx, fmt.Sprintf(
					"SELECT ifNull(%s, 0) FROM last_blocks WHERE %s = %d AND %s = '%s'",
					sql.BlockNumberFieldName, sql.ChainIDFieldName, c.chainConfig.ChainID, sql.ContractAddressFieldName, contract.Address,
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
						Max:    10 * time.Second,
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

							err = c.processLogs(groupCtx, logs, eventParser)
							if err != nil {
								logger.Warnf("could not process logs for chain %d: %s", c.chainConfig.ChainID, err)
								continue
							}
							fmt.Println("processed Chunk logs for chain", contract.ContractType, c.chainConfig.ChainID, "from block", funcHeight, "to block", rangeEnd)
							return nil
						}
					}
				})
			}

			if err := g.Wait(); err != nil {
				return fmt.Errorf("error while backfilling chain %d: %w", c.chainConfig.ChainID, err)
			}
			logger.Infof("backfilling contract %s completed", contract.Address)

			return nil
		})
	}
	if err := contractsGroup.Wait(); err != nil {
		return fmt.Errorf("error while backfilling chain %d: %w", c.chainConfig.ChainID, err)
	}
	logger.Infof("backfilling chain %d completed", c.chainConfig.ChainID)
	return nil
}

// processLogs processes the logs and stores them in the consumer database.
//
//nolint:gocognit,cyclop
func (c *ChainBackfiller) processLogs(ctx context.Context, logs []ethTypes.Log, eventParser parser.Parser) error {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    10 * time.Second,
	}

	timeout := time.Duration(0)
	logIdx := 0
	for {
		select {
		case <-ctx.Done():

			return fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):
			if logIdx >= len(logs) {
				return nil
			}
			err := eventParser.ParseAndStore(ctx, logs[logIdx], c.chainConfig.ChainID)
			if err != nil {
				logger.Errorf("could not parse and store log %d, %s: %s", c.chainConfig.ChainID, logs[logIdx].Address, err)
				timeout = b.Duration()
				continue
			}

			// Store the last block in clickhouse
			err = c.consumerDB.StoreLastBlock(ctx, c.chainConfig.ChainID, logs[logIdx].BlockNumber, logs[logIdx].Address.String())
			if err != nil {
				logger.Errorf("could not store last block for chain %d: %s", c.chainConfig.ChainID, err)
				timeout = b.Duration()
				continue
			}

			logIdx++

			// Reset the backoff after successful log parse run to prevent bloated back offs.
			b.Reset()
			timeout = time.Duration(0)
		}
	}
}
