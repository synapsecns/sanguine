package backfill

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/synapsecns/sanguine/ethergo/util"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	indexerconfig "github.com/synapsecns/sanguine/services/explorer/config/indexer"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/db"
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
	// cctpParser is the parser to use to parse cctp events.
	cctpParser *parser.CCTPParser
	// Fetcher is the Fetcher to use to fetch logs.
	Fetcher fetcher.ScribeFetcher
	// chainConfig is the chain config for the chain.
	chainConfig indexerconfig.ChainConfig
}

type contextKey string

const (
	chainKey contextKey = "chainID"
)

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(consumerDB db.ConsumerDB, bridgeParser *parser.BridgeParser, swapParsers map[common.Address]*parser.SwapParser, messageBusParser *parser.MessageBusParser, cctpParser *parser.CCTPParser, fetcher fetcher.ScribeFetcher, chainConfig indexerconfig.ChainConfig) *ChainBackfiller {
	return &ChainBackfiller{
		consumerDB:       consumerDB,
		bridgeParser:     bridgeParser,
		swapParsers:      swapParsers,
		messageBusParser: messageBusParser,
		cctpParser:       cctpParser,
		Fetcher:          fetcher,
		chainConfig:      chainConfig,
	}
}

// Backfill fetches logs from the GraphQL database, parses them, and stores them in the consumer database.
// nolint:cyclop,gocognit
func (c *ChainBackfiller) Backfill(ctx context.Context, livefill bool, refreshRate int) (err error) {
	chainCtx := context.WithValue(ctx, chainKey, fmt.Sprintf("%d", c.chainConfig.ChainID))
	contractsGroup, contractCtx := errgroup.WithContext(chainCtx)

	if !livefill {
		for i := range c.chainConfig.Contracts {
			contract := c.chainConfig.Contracts[i]
			contractsGroup.Go(func() error {
				err := c.backfillContractLogs(contractCtx, contract)
				if err != nil {
					return fmt.Errorf("could not backfill contract logs: %w", err)
				}
				return nil
			})
		}
	} else {
		for i := range c.chainConfig.Contracts {
			contract := c.chainConfig.Contracts[i]
			contractsGroup.Go(func() error {
				b := &backoff.Backoff{
					Factor: 2,
					Jitter: true,
					Min:    1 * time.Second,
					Max:    3 * time.Second,
				}
				timeout := time.Duration(0)
				for {
					select {
					case <-chainCtx.Done():
						logger.Errorf("livefill of contract %s on chain %d failed: %v", contract.Address, c.chainConfig.ChainID, chainCtx.Err())

						return fmt.Errorf("livefill of contract %s on chain %d failed: %w", contract.Address, c.chainConfig.ChainID, chainCtx.Err())
					case <-time.After(timeout):
						err := c.backfillContractLogs(contractCtx, contract)
						if err != nil {
							timeout = b.Duration()
							logger.Warnf("could not livefill contract %s on chain %d, retrying %v", contract.Address, c.chainConfig.ChainID, err)

							continue
						}
						b.Reset()
						timeout = time.Duration(refreshRate) * time.Second
						logger.Infof("processed range for contract %s on chain %d, continuing to livefill in %d seconds - refresh rate %d ", contract.Address, c.chainConfig.ChainID, timeout, refreshRate)
					}
				}
			})
		}
	}
	if err := contractsGroup.Wait(); err != nil {
		logger.Errorf("error backfilling chain %d completed %v", c.chainConfig.ChainID, err)

		return fmt.Errorf("error while backfilling chain %d: %w", c.chainConfig.ChainID, err)
	}
	return nil
}

// makeEventParser returns a parser for a contract using it's config.
// in the event one is not present, this function will return an error.
func (c *ChainBackfiller) makeEventParser(contract indexerconfig.ContractConfig) (eventParser parser.Parser, err error) {
	contractType, err := indexerconfig.ContractTypeFromString(contract.ContractType)
	if err != nil {
		return nil, fmt.Errorf("could not create event parser for unknown contract type: %s", contract.ContractType)
	}
	switch contractType {
	case indexerconfig.BridgeContractType:
		eventParser = c.bridgeParser
	case indexerconfig.SwapContractType:
		eventParser = c.swapParsers[common.HexToAddress(contract.Address)]
	case indexerconfig.MessageBusContractType:
		eventParser = c.messageBusParser
	case indexerconfig.MetaSwapContractType:
		eventParser = c.swapParsers[common.HexToAddress(contract.Address)]
	case indexerconfig.CCTPContractType:
		eventParser = c.cctpParser
	}
	return eventParser, nil
}

// backfillContractLogs creates a backfiller for a given contract with an independent context
// nolint:cyclop,gocognit
func (c *ChainBackfiller) backfillContractLogs(parentCtx context.Context, contract indexerconfig.ContractConfig) (err error) {
	// make the event parser
	eventParser, err := c.makeEventParser(contract)
	if err != nil {
		return err
	}

	startHeight := uint64(contract.StartBlock)

	// Set start block to -1 to trigger backfill from last block stored by explorer,
	// otherwise backfilling will begin at the block number specified in the config file.
	if contract.StartBlock < 0 {
		startHeight, err = c.consumerDB.GetLastStoredBlock(parentCtx, c.chainConfig.ChainID, contract.Address)
		if err != nil {
			return fmt.Errorf("could not get last block number: %w, %s", err, contract.ContractType)
		}
	}
	var endHeight uint64
	err = c.retryWithBackoff(parentCtx, func(ctx context.Context) error {
		// TODO change to get last unconfirmed block
		endHeight, err = c.Fetcher.FetchLastIndexed(parentCtx, c.chainConfig.ChainID, contract.Address)
		if err != nil {
			return fmt.Errorf("could not get last indexed height, %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not get last indexed for contract %s: %w, %v", contract.Address, err, c.chainConfig)
	}

	currentHeight := startHeight
	// Iterate over all blocks and fetch logs with the current contract address.
	for currentHeight <= endHeight {
		// Create context for backfilling chunks
		g, groupCtx := errgroup.WithContext(parentCtx)

		chunkStart := currentHeight
		chunkEnd := currentHeight + (c.chainConfig.FetchBlockIncrement-1)*uint64(c.chainConfig.MaxGoroutines)
		if chunkEnd > endHeight {
			chunkEnd = endHeight
		}

		iterator := util.NewChunkIterator(big.NewInt(int64(chunkStart)), big.NewInt(int64(chunkEnd)), int(c.chainConfig.FetchBlockIncrement)-1, true)
		for subChunk := iterator.NextChunk(); subChunk != nil; subChunk = iterator.NextChunk() {
			chunkVar := subChunk
			g.Go(func() error {
				b := &backoff.Backoff{
					Factor: 2,
					Jitter: true,
					Min:    1 * time.Second,
					Max:    3 * time.Second,
				}

				timeout := time.Duration(0)

				for {
					select {
					case <-groupCtx.Done():
						return fmt.Errorf("context canceled: %w", groupCtx.Err())
					case <-time.After(timeout):
						rangeEnd := chunkVar.EndBlock.Uint64()

						// Fetch the logs from Scribe.
						logs, err := c.Fetcher.FetchLogsInRange(groupCtx, c.chainConfig.ChainID, chunkVar.StartBlock.Uint64(), rangeEnd, common.HexToAddress(contract.Address))
						if err != nil {
							timeout = b.Duration()
							logger.Warnf("could not fetch logs for chain %d: %v. Retrying in %s", c.chainConfig.ChainID, err, timeout)

							continue
						}

						parsedLogs, err := ProcessLogs(groupCtx, logs, c.chainConfig.ChainID, eventParser)
						if err != nil {
							timeout = b.Duration()
							logger.Warnf("could not process logs for chain %d: %s", c.chainConfig.ChainID, err)
							continue
						}

						if len(parsedLogs) > 0 {
							g.Go(func() error {
								return c.storeParsedLogs(groupCtx, parsedLogs)
							})
						}
						return nil
					}
				}
			})
		}

		if err := g.Wait(); err != nil {
			return fmt.Errorf("error while backfilling chain %d: %w", c.chainConfig.ChainID, err)
		}
		logger.Infof("backfilling contract %s chunk completed, %d to %d", contract.Address, chunkStart, chunkEnd)
		// Store the last block in clickhouse
		err = c.retryWithBackoff(parentCtx, func(ctx context.Context) error {
			err = c.consumerDB.StoreLastBlock(parentCtx, c.chainConfig.ChainID, chunkEnd, contract.Address)
			if err != nil {
				return fmt.Errorf("error storing last block, %w", err)
			}
			return nil
		})
		if err != nil {
			logger.Errorf("could not store last block for chain %d: %s %d, %s, %s", c.chainConfig.ChainID, err, chunkEnd, contract.Address, contract.ContractType)
			return fmt.Errorf("could not store last block for chain %d: %w", c.chainConfig.ChainID, err)
		}
		currentHeight = chunkEnd + 1
	}

	return nil
}

// ProcessLogs processes the logs and stores them in the consumer database.
//
//nolint:gocognit,cyclop
func ProcessLogs(ctx context.Context, logs []ethTypes.Log, chainID uint32, eventParser parser.Parser) (parsedLogs []interface{}, _ error) {
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
			return parsedLogs, fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):
			if logIdx >= len(logs) {
				return parsedLogs, nil
			}
			parsedLog, err := eventParser.Parse(ctx, logs[logIdx], chainID)
			if err != nil || parsedLog == nil {
				// TODO: this should really, REALLY use errors.IS and wrap the underlying error
				if strings.Contains(err.Error(), parser.ErrUnknownTopic) {
					logger.Warnf("could not parse log (ErrUnknownTopic) %d, %s %s blocknumber: %d, %s", chainID, logs[logIdx].TxHash, logs[logIdx].Address, logs[logIdx].BlockNumber, err)
				} else { // retry
					logger.Errorf("could not parse log %d, %s blocknumber: %d, %s", chainID, logs[logIdx].Address, logs[logIdx].BlockNumber, err)
					timeout = b.Duration()
					continue
				}
			}

			parsedLogs = append(parsedLogs, parsedLog)

			logIdx++

			// Reset the backoff after successful log parse run to prevent bloated back offs.
			b.Reset()
			timeout = time.Duration(0)
		}
	}
}

func (c *ChainBackfiller) storeParsedLogs(ctx context.Context, parsedEvents []interface{}) error {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    3 * time.Millisecond,
		Max:    2 * time.Second,
	}
	timeout := time.Duration(0)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled while storing events: %w", ctx.Err())
		case <-time.After(timeout):
			err := c.consumerDB.StoreEvents(ctx, parsedEvents)
			if err != nil {
				logger.Errorf("Error storing events: %v", err)
				timeout = b.Duration()
				continue
			}
			return nil
		}
	}
}

const maxAttempt = 20

type retryableFunc func(ctx context.Context) error

// retryWithBackoff will retry to get data with a backoff.
func (c *ChainBackfiller) retryWithBackoff(ctx context.Context, doFunc retryableFunc) error {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    3 * time.Second,
	}

	timeout := time.Duration(0)
	attempts := 0
	for attempts < maxAttempt {
		select {
		case <-ctx.Done():
			return fmt.Errorf("%w while retrying", ctx.Err())
		case <-time.After(timeout):
			err := doFunc(ctx)
			if err != nil {
				timeout = b.Duration()
				attempts++
			} else {
				return nil
			}
		}
	}
	return fmt.Errorf("max attempts reached while retrying")
}
