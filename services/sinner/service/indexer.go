package service

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/core/retry"

	"time"

	indexerConfig "github.com/synapsecns/sanguine/services/sinner/config/indexer"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/fetcher"
	"github.com/synapsecns/sanguine/services/sinner/logger"
	"github.com/synapsecns/sanguine/services/sinner/types"
	"golang.org/x/sync/errgroup"
)

// ChainIndexer indexes message logs for a chain.
type ChainIndexer struct {
	// consumerDB is the database to store consumer data in.
	eventDB db.EventDB
	// parsers are the parsers for this chain.
	parsers Parsers
	// fetcher is the scribe fetcher.
	fetcher fetcher.ScribeFetcher
	// config is the config for the backfiller.
	config indexerConfig.ChainConfig
}

// NewChainIndexer creates a new chain indexer.
func NewChainIndexer(eventDB db.EventDB, parsers Parsers, fetcher fetcher.ScribeFetcher, config indexerConfig.ChainConfig) *ChainIndexer {
	if config.GoroutinesPerContract < 1 {
		config.GoroutinesPerContract = 1
	}

	if config.FetchBlockIncrement < 1 {
		config.FetchBlockIncrement = 10000
	}
	chainIndexer := ChainIndexer{
		eventDB,
		parsers,
		fetcher,
		config,
	}
	return &chainIndexer
}

func getBackoffConfig() *backoff.Backoff {
	return &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    3 * time.Second,
	}
}

func (c ChainIndexer) getScribeData(parentCtx context.Context, startBlock uint64, endBlock uint64, contractAddress common.Address) ([]ethTypes.Log, map[string]types.TxSupplementalInfo, error) {
	b := getBackoffConfig()
	timeout := time.Duration(0)
	for {
		select {
		case <-parentCtx.Done():
			return nil, nil, fmt.Errorf("could not get scribe data from block %d and %d. Error: %w", startBlock, endBlock, parentCtx.Err())
		case <-time.After(timeout):
			logs, err := c.fetcher.FetchLogsInRange(parentCtx, c.config.ChainID, startBlock, endBlock, contractAddress)
			if err != nil {
				logger.ReportSinnerError(fmt.Errorf("could not get logs %w", err), c.config.ChainID, logger.ScribeFetchFailure)
				timeout = b.Duration()
				continue
			}

			txs, err := c.fetcher.FetchTxsInRange(parentCtx, c.config.ChainID, startBlock, endBlock)
			if err != nil {
				logger.ReportSinnerError(fmt.Errorf("could not get txs %w", err), c.config.ChainID, logger.ScribeFetchFailure)
				timeout = b.Duration()
				continue
			}

			txMap := make(map[string]types.TxSupplementalInfo)
			for _, tx := range txs {
				txMap[tx.TxHash] = tx
			}
			return logs, txMap, nil
		}
	}
}

// Index indexes all contracts from the config for the current chain.
//
// nolint:gocognit,cyclop
func (c ChainIndexer) Index(ctx context.Context) error {
	contractGroup, contractCtx := errgroup.WithContext(ctx)

	for i := range c.config.Contracts {
		contract := c.config.Contracts[i]
		contractType, err := indexerConfig.ContractTypeFromString(contract.ContractType)
		if err != nil {
			return fmt.Errorf("error creating parser for contract type: %s", contract.ContractType)
		}
		var eventParser types.EventParser
		switch contractType {
		case indexerConfig.OriginType:
			eventParser = c.parsers.OriginParser
		case indexerConfig.ExecutionHubType:
			eventParser = c.parsers.DestinationParser
		case indexerConfig.UnknownType:
			return fmt.Errorf("could not create event parser for unknown contract type: %s", contract.ContractType)
		}
		refreshRate := time.Duration(1)

		// Create thread for current contract
		contractGroup.Go(func() error {
			for {
				select {
				case <-contractCtx.Done():
					return fmt.Errorf("could not index contract. Error: %w", contractCtx.Err())
				case <-time.After(refreshRate):
					startHeight := contract.StartBlock
					endHeight := contract.EndBlock

					// If the end block is not specified in the config (livefill) the last block stored will be used.
					if endHeight == 0 {
						// Get last stored block from sinner.
						storedStartHeight, err := c.eventDB.RetrieveLastStoredBlock(contractCtx, c.config.ChainID, common.HexToAddress(contract.Address))
						if err != nil {
							return fmt.Errorf("could not get last block number: %w, %s", err, contract.ContractType)
						}
						if storedStartHeight > startHeight {
							startHeight = storedStartHeight
						}

						// Get last indexed from Scribe.
						err = retry.WithBackoff(contractCtx, func(ctx context.Context) error {
							endHeight, err = c.fetcher.FetchLastIndexed(contractCtx, c.config.ChainID, contract.Address)
							if err != nil {
								return fmt.Errorf("could not get last indexed height, %w", err)
							}
							return nil
						})

						if err != nil {
							return fmt.Errorf("could not get last indexed height, %w", err)
						}
					}

					// Semaphore to limit the number of goroutines for each contract.
					sem := make(chan struct{}, c.config.GoroutinesPerContract)

					// Iterate through all blocks between start and finish.
					for currentHeight := startHeight; currentHeight <= endHeight; currentHeight += c.config.FetchBlockIncrement {
						if currentHeight > endHeight {
							currentHeight = endHeight
						}

						endFetchRange := currentHeight + c.config.FetchBlockIncrement
						if endFetchRange > endHeight {
							endFetchRange = endHeight
						}

						// Fetch logs and txs from scribe.
						logs, txs, contractErr := c.getScribeData(contractCtx, currentHeight, endFetchRange, common.HexToAddress(contract.Address))
						if contractErr != nil {
							return fmt.Errorf("error getting scribe data: %w", contractErr)
						}
						eventParser.UpdateTxMap(txs)
						logGroup, logCtx := errgroup.WithContext(contractCtx)

						// For each log, spin up a go routine and parse + store that data.
						for _, log := range logs {
							sem <- struct{}{}

							currentLog := log
							logGroup.Go(func() error {
								defer func() { <-sem }() // empty the chan by one semaphore

								contractErr = retry.WithBackoff(logCtx, func(parentCtx context.Context) error {
									parseErr := eventParser.ParseAndStore(parentCtx, currentLog)
									if parseErr != nil {
										return fmt.Errorf("error parsing and storing event: %w", parseErr)
									}
									return nil
								})

								if contractErr != nil {
									return fmt.Errorf("error parsing and storing event: %w", contractErr)
								}
								return nil
							})
						}
						// wait for all goroutines to finish
						if err := logGroup.Wait(); err != nil {
							return fmt.Errorf("error processing logs: %w", err)
						}

						// store last indexed
						height := currentHeight
						err = retry.WithBackoff(contractCtx, func(parentCtx context.Context) error {
							storeErr := c.eventDB.StoreLastIndexed(parentCtx, common.HexToAddress(contract.Address), c.config.ChainID, height)
							if storeErr != nil {
								return fmt.Errorf("error storing last indexed: %w", storeErr)
							}
							return nil
						})

						if err != nil {
							return fmt.Errorf("error storing last indexed: %w", err)
						}
					}

					// Backfill complete. Terminate current thread.
					if contract.EndBlock > 0 {
						return nil
					}
					// Continue livefilling
				}
			}
		})
	}
	// wait for all goroutines to finish
	if err := contractGroup.Wait(); err != nil {
		return fmt.Errorf("error processing: %w", err)
	}
	return nil
}
