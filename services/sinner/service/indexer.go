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
	// default refresh rate
	refreshRate time.Duration
}

// NewChainIndexer creates a new chain indexer.
func NewChainIndexer(eventDB db.EventDB, parsers Parsers, fetcher fetcher.ScribeFetcher, config indexerConfig.ChainConfig, refreshRate time.Duration) *ChainIndexer {
	chainIndexer := ChainIndexer{
		eventDB,
		parsers,
		fetcher,
		config,
		refreshRate,
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

func (c ChainIndexer) getScribeData(parentCtx context.Context, startBlock uint64, endBlock uint64, contractAddress common.Address, page int) ([]ethTypes.Log, map[string]types.TxSupplementalInfo, error) {
	b := getBackoffConfig()
	timeout := time.Duration(0)
	for {
		select {
		case <-parentCtx.Done():
			return nil, nil, fmt.Errorf("could not get scribe data from block %d and %d. Error: %w", startBlock, endBlock, parentCtx.Err())
		case <-time.After(timeout):
			logs, txs, err := c.fetcher.FetchLogsAndTransactionsRange(parentCtx, c.config.ChainID, startBlock, endBlock, contractAddress, page)
			if err != nil {
				logger.ReportSinnerError(fmt.Errorf("could not get logs and txs %w", err), c.config.ChainID, logger.ScribeFetchFailure)
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
func (c ChainIndexer) Index(ctx context.Context) error {
	contractGroup, contractCtx := errgroup.WithContext(ctx)

	for _, contract := range c.config.Contracts {
		if err := c.processContract(contractCtx, contractGroup, contract); err != nil {
			return err
		}
	}
	return contractGroup.Wait()
}

// processContract indexes a contract by creating an appropriate parser and then indexing the contract in a goroutine.
func (c ChainIndexer) processContract(contractCtx context.Context, contractGroup *errgroup.Group, contract indexerConfig.ContractConfig) error {
	eventParser, err := c.createEventParser(contract)
	if err != nil {
		return err
	}
	contractGroup.Go(func() error {
		return c.indexContractEvents(contractCtx, contract, eventParser)
	})
	return nil
}

// createEventParser creates an event parser for the given contract.
func (c ChainIndexer) createEventParser(contract indexerConfig.ContractConfig) (types.EventParser, error) {
	contractType, err := indexerConfig.ContractTypeFromString(contract.ContractType)
	if err != nil {
		return nil, fmt.Errorf("error creating parser for contract type: %s", contract.ContractType)
	}
	switch contractType {
	case indexerConfig.OriginType:
		return c.parsers.OriginParser, nil
	case indexerConfig.ExecutionHubType:
		return c.parsers.DestinationParser, nil
	case indexerConfig.UnknownType:
		return nil, fmt.Errorf("could not create event parser for unknown contract type: %s", contract.ContractType)
	default:
		return nil, fmt.Errorf("unsupported contract type: %s", contract.ContractType)
	}
}

// indexContractEvents indexes all events for a contract.
func (c ChainIndexer) indexContractEvents(contractCtx context.Context, contract indexerConfig.ContractConfig, eventParser types.EventParser) error {
	for {
		select {
		case <-contractCtx.Done():
			return fmt.Errorf("could not index contract. Error: %w", contractCtx.Err())
		case <-time.After(c.refreshRate):
			startHeight, endHeight, err := c.fetchBlockRange(contractCtx, contract)
			if err != nil {
				return err
			}
			if err := c.processBlocksInRange(contractCtx, startHeight, endHeight, contract, eventParser); err != nil {
				return err
			}

			// Terminate if backfill is complete.
			if contract.EndBlock > 0 {
				return nil
			}
		}
	}
}

// fetchBlockRange gets the block range to fetch for the contract.
func (c ChainIndexer) fetchBlockRange(contractCtx context.Context, contract indexerConfig.ContractConfig) (uint64, uint64, error) {
	startHeight := contract.StartBlock
	endHeight := contract.EndBlock

	// If the end block is not specified in the config (livefill) the last block stored will be used.
	if endHeight == 0 {
		storedStartHeight, err := c.eventDB.RetrieveLastStoredBlock(contractCtx, c.config.ChainID, common.HexToAddress(contract.Address))
		if err != nil {
			return 0, 0, fmt.Errorf("could not get last block number: %w, %s", err, contract.ContractType)
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
			return 0, 0, fmt.Errorf("could not get last indexed height, %w", err)
		}
	}

	return startHeight, endHeight, nil
}

// processBlocksInRange handles all the core indexing logic and processes all blocks in the range for the contract.
func (c ChainIndexer) processBlocksInRange(contractCtx context.Context, startHeight uint64, endHeight uint64, contract indexerConfig.ContractConfig, eventParser types.EventParser) error {
	currPage := 1

	// Loop until all logs are processed in the provided range.
	for {
		logs, txs, contractErr := c.getScribeData(contractCtx, startHeight, endHeight, common.HexToAddress(contract.Address), currPage)
		if contractErr != nil {
			return fmt.Errorf("error getting scribe data: %w", contractErr)
		}

		// No more logs to process, finished draining logs/txs with pagination.
		if len(logs) == 0 {
			return nil
		}

		for _, log := range logs {
			tx, ok := txs[log.TxHash.String()]
			if !ok {
				tx = types.TxSupplementalInfo{}
			}
			contractErr = retry.WithBackoff(contractCtx, func(parentCtx context.Context) error {
				parseErr := eventParser.ParseAndStore(parentCtx, log, tx)
				if parseErr != nil {
					return fmt.Errorf("error parsing and storing event: %w", parseErr)
				}
				return nil
			})

			if contractErr != nil {
				return fmt.Errorf("error parsing and storing event: %w", contractErr)
			}
		}

		// Store last indexed block number
		currentHeight := logs[len(logs)-1].BlockNumber
		err := retry.WithBackoff(contractCtx, func(parentCtx context.Context) error {
			storeErr := c.eventDB.StoreLastIndexed(parentCtx, common.HexToAddress(contract.Address), c.config.ChainID, currentHeight)
			if storeErr != nil {
				return fmt.Errorf("error storing last indexed: %w", storeErr)
			}
			return nil
		})

		if err != nil {
			return fmt.Errorf("error storing last indexed: %w", err)
		}

		// Move to next page
		currPage++
	}
}
