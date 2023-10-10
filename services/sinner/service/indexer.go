package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/core/retry"

	indexerConfig "github.com/synapsecns/sanguine/services/sinner/config/indexer"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/fetcher"
	"github.com/synapsecns/sanguine/services/sinner/logger"
	"github.com/synapsecns/sanguine/services/sinner/types"
	"golang.org/x/sync/errgroup"
	"time"
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

func NewChainIndexer(eventDB db.EventDB, parsers Parsers, fetcher fetcher.ScribeFetcher, config indexerConfig.ChainConfig) *ChainIndexer {
	if config.MaxGoroutines < 1 {
		config.MaxGoroutines = 1
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
			return nil, nil, fmt.Errorf("could not get scribe data from block %d and %d. Error: %v", startBlock, endBlock, parentCtx.Err())
		case <-time.After(timeout):
			logs, err := c.fetcher.FetchLogsInRange(parentCtx, c.config.ChainID, startBlock, endBlock, contractAddress)
			if err != nil {
				logger.ReportSinnerError(fmt.Errorf("could not get logs %v", err), c.config.ChainID, logger.ScribeFetchFailure)
				timeout = b.Duration()
				continue
			}

			txs, err := c.fetcher.FetchTxsInRange(parentCtx, c.config.ChainID, startBlock, endBlock)
			if err != nil {
				logger.ReportSinnerError(fmt.Errorf("could not get txs %v", err), c.config.ChainID, logger.ScribeFetchFailure)
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

func (c ChainIndexer) Index(ctx context.Context) error {
	contractGroup, contractCtx := errgroup.WithContext(ctx)

	for i := range c.config.Contracts {
		contract := c.config.Contracts[i]
		contractType, err := indexerConfig.ContractTypeFromString(contract.ContractType)
		if err != nil {
			return fmt.Errorf("could not create event parser for unknown contract type: %s", contract.ContractType)
		}
		var eventParser types.EventParser
		switch contractType {
		case indexerConfig.OriginType:
			eventParser = c.parsers.OriginParser
		case indexerConfig.ExecutionHubType:
			eventParser = c.parsers.DestinationParser
		}
		refreshRate := time.Duration(1)

		// Create thread for current contract
		contractGroup.Go(func() error {
			for {
				select {
				case <-contractCtx.Done():
					return fmt.Errorf("could not index contract. Error: %v", contractCtx.Err())
				case <-time.After(refreshRate):
					startHeight := contract.StartBlock
					endHeight := contract.EndBlock

					// If the end block is not specified in the config (livefill) the last block stored will be used.
					if endHeight == 0 {
						storedStartHeight, err := c.eventDB.RetrieveLastStoredBlock(contractCtx, c.config.ChainID, common.HexToAddress(contract.Address))
						if err != nil {
							return fmt.Errorf("could not get last block number: %w, %s", err, contract.ContractType)
						}
						if storedStartHeight > startHeight {
							startHeight = storedStartHeight
						}

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

					sem := make(chan struct{}, c.config.MaxGoroutines) // semaphore to limit the number of goroutines

					for currentHeight := startHeight; currentHeight <= endHeight; currentHeight += c.config.FetchBlockIncrement {
						if currentHeight > endHeight {
							currentHeight = endHeight
						}

						logs, txs, contractErr := c.getScribeData(contractCtx, currentHeight, currentHeight+c.config.FetchBlockIncrement, common.HexToAddress(contract.Address))
						if contractErr != nil {
							return fmt.Errorf("error getting scribe data: %v", contractErr)
						}
						eventParser.UpdateTxMap(txs)
						logGroup, logCtx := errgroup.WithContext(contractCtx)

						for _, log := range logs {
							sem <- struct{}{}

							currentLog := log
							logGroup.Go(func() error {
								defer func() { <-sem }() // empty the chan by one semaphore

								contractErr = retry.WithBackoff(logCtx, func(parentCtx context.Context) error {

									return eventParser.ParseAndStore(parentCtx, currentLog)
								})

								if contractErr != nil {
									return fmt.Errorf("error parsing and storing event: %v", contractErr)
								}
								return nil
							})
						}
						// wait for all goroutines to finish
						if err := logGroup.Wait(); err != nil {
							return fmt.Errorf("error processing logs: %v", err)
						}

						// store last indexed
						err = retry.WithBackoff(contractCtx, func(parentCtx context.Context) error {
							return c.eventDB.StoreLastIndexed(parentCtx, common.HexToAddress(contract.Address), c.config.ChainID, currentHeight)
						})
						if err != nil {
							return fmt.Errorf("error storing last indexed: %v", err)
						}
					}

					// Backfill complete. Terminate current thread.
					if contract.EndBlock == 0 {
						return nil
					}
					// Continue livefilling
				}
			}
		})

		return nil
	}
	// wait for all goroutines to finish
	if err := contractGroup.Wait(); err != nil {
		return fmt.Errorf("error processing: %v", err)
	}
	return nil
}
