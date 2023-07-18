package scribe

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/logger"
	"github.com/synapsecns/sanguine/services/scribe/scribe/indexer"
	"math/big"

	"math"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/metric"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ChainIndexer is an indexer that fetches logs for a chain. It aggregates logs
// from a slice of ContractIndexers.
type ChainIndexer struct {
	// chainID is the chain ID of the chain.
	chainID uint32
	// eventDB is the database to store event data in.
	eventDB db.EventDB
	// client contains the clients used for indexing.
	client []backend.ScribeBackend
	// chainConfig is the config for the indexer.
	chainConfig config.ChainConfig
	// handler is the metrics handler for the scribe.
	handler metrics.Handler
	// blockHeightMeters is a map from address -> meter for block height.
	blockHeightMeters map[common.Address]metric.Int64Histogram
	// livefillContracts is a map from address -> livefill contract.
	livefillContracts []config.ContractConfig
}

// Used for handling logging of various context types.
type contextKey int

const maxBackoff = uint64(10)

const (
	chainContextKey contextKey = iota
)

// NewChainIndexer creates a new indexer for a chain. This is done by passing through all the function parameters
// into the ChainIndexer struct, as well as iterating through all the contracts in the chain config & creating
// ContractIndexers for each contract.
func NewChainIndexer(eventDB db.EventDB, client []backend.ScribeBackend, chainConfig config.ChainConfig, handler metrics.Handler) (*ChainIndexer, error) {
	if chainConfig.GetLogsRange == 0 {
		chainConfig.GetLogsRange = 600
	}

	if chainConfig.GetLogsBatchAmount == 0 {
		chainConfig.GetLogsBatchAmount = 2
	}

	if chainConfig.StoreConcurrency == 0 {
		chainConfig.StoreConcurrency = 20
	}

	if chainConfig.ConcurrencyThreshold == 0 {
		chainConfig.ConcurrencyThreshold = 50000
	}
	if chainConfig.LivefillRange == 0 {
		chainConfig.LivefillRange = 100
	}

	if chainConfig.LivefillFlushInterval == 0 {
		chainConfig.LivefillFlushInterval = 10800
	}

	blockHeightMeterMap := make(map[common.Address]metric.Int64Histogram)
	for _, contract := range chainConfig.Contracts {
		blockHeightMeter, err := handler.Meter().NewHistogram(fmt.Sprintf("scribe_block_meter_%d_%s", chainConfig.ChainID, contract.Address), "block_histogram", "a block height meter", "blocks")
		if err != nil {
			return nil, fmt.Errorf("error creating otel histogram %w", err)
		}
		blockHeightMeterMap[common.HexToAddress(contract.Address)] = blockHeightMeter
	}

	return &ChainIndexer{
		chainID:           chainConfig.ChainID,
		eventDB:           eventDB,
		client:            client,
		blockHeightMeters: blockHeightMeterMap,
		chainConfig:       chainConfig,
		handler:           handler,
	}, nil
}

// Index iterates over each contract indexer and calls Index concurrently on each one.
// If `onlyOneBlock` is true, the indexer will only index the block at `currentBlock`.
//
//nolint:gocognit,cyclop,unparam
func (c *ChainIndexer) Index(ctx context.Context, onlyOneBlock *uint64) error {
	// Create a new context for the chain so all chains don't halt when indexing is completed.
	chainCtx := context.WithValue(ctx, chainContextKey, fmt.Sprintf("%d", c.chainID))
	indexGroup, indexCtx := errgroup.WithContext(chainCtx)

	// var livefillContracts []config.ContractConfig
	readyToLivefill := make(chan config.ContractConfig)

	latestBlock, err := c.getLatestBlock(indexCtx, true)
	if err != nil {
		return fmt.Errorf("could not get current block number while indexing: %w", err)
	}

	var contractAddresses []common.Address
	for i := range c.chainConfig.Contracts {
		contractAddresses = append(contractAddresses, common.HexToAddress(c.chainConfig.Contracts[i].Address))
	}

	// Gets all last indexed infos for the contracts on the current chain to determine which contracts need to be initially livefilled.
	lastIndexedMap, err := c.eventDB.RetrieveLastIndexedMultiple(chainCtx, contractAddresses, c.chainConfig.ChainID)
	if err != nil {
		return fmt.Errorf("could not get last indexed map: %w", err)
	}

	for j := range c.chainConfig.Contracts {
		contract := c.chainConfig.Contracts[j]
		contractAddress := common.HexToAddress(contract.Address)
		lastIndexed := lastIndexedMap[contractAddress]

		// Does not consider if the config's start block is within the livefill threshold for simplicity.
		// In this case, an indexer will bring the contract to head, and it will be passed to livefill.
		// If there is no last indexed info for the contract, it will not be passed to livefill.
		if *latestBlock-c.chainConfig.LivefillThreshold > lastIndexed && lastIndexed > 0 {
			c.livefillContracts = append(c.livefillContracts, contract)
			continue
		}

		// If current contract is not within the livefill threshold, start an indexer for it.
		contractIndexer, err := indexer.NewIndexer(c.chainConfig, []common.Address{contractAddress}, c.eventDB, c.client, c.handler, c.blockHeightMeters[contractAddress], false)
		if err != nil {
			return fmt.Errorf("could not create contract indexer: %w", err)
		}

		indexGroup.Go(func() error {
			err := c.IndexToBlock(indexCtx, onlyOneBlock, contract.StartBlock, contractIndexer)
			if err != nil {
				return fmt.Errorf("could not index to livefill: %w", err)
			}
			readyToLivefill <- contract

			// TODO make sure metrics are killed when indexing is done
			return nil
		})
	}

	// Livefill contracts that are within the livefill threshold.
	indexGroup.Go(func() error {
		timeout := time.Duration(0)
		b := createBackoff()
		livefillBlockMeter, err := c.handler.Meter().NewHistogram(fmt.Sprintf("scribe_block_meter_%d_livefill", c.chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
		if err != nil {
			return fmt.Errorf("error creating otel histogram %w", err)
		}

		livefillIndexer, err := indexer.NewIndexer(c.chainConfig, getAddressesFromConfig(c.livefillContracts), c.eventDB, c.client, c.handler, livefillBlockMeter, false)
		if err != nil {
			return fmt.Errorf("could not create contract indexer: %w", err)
		}
		for {
			select {
			case <-indexCtx.Done():
				return fmt.Errorf("%s chain context canceled: %w", indexCtx.Value(chainContextKey), indexCtx.Err())
			case newLivefillContract := <-readyToLivefill:
				c.livefillContracts = append(c.livefillContracts, newLivefillContract)
				// Update indexer's config to include new contract.
				livefillIndexer.UpdateAddress(getAddressesFromConfig(c.livefillContracts))
			case <-time.After(timeout):
				if len(c.livefillContracts) == 0 {
					timeout = b.Duration()
					continue
				}
				var endHeight *uint64
				var err error
				livefillLastIndexed, err := c.eventDB.RetrieveLastIndexedMultiple(chainCtx, contractAddresses, c.chainConfig.ChainID)
				if err != nil {
					logger.ReportIndexerError(err, livefillIndexer.GetIndexerConfig(), logger.LivefillIndexerError)
					timeout = b.Duration()
					continue
				}
				startHeight := getMinFromMap(livefillLastIndexed)

				endHeight, err = c.getLatestBlock(indexCtx, true)
				if err != nil {
					logger.ReportIndexerError(err, livefillIndexer.GetIndexerConfig(), logger.GetBlockError)
					timeout = b.Duration()
					continue
				}

				// Don't reindex the head block.
				if startHeight == *endHeight {
					timeout = 1 * time.Second
					continue
				}

				err = livefillIndexer.Index(indexCtx, startHeight, *endHeight)
				if err != nil {
					timeout = b.Duration()
					logger.ReportIndexerError(err, livefillIndexer.GetIndexerConfig(), logger.LivefillIndexerError)
					continue
				}

				// Default refresh rate for livefill is 1 second.
				timeout = 1 * time.Second
			}
		}
	})

	if err := indexGroup.Wait(); err != nil {
		return fmt.Errorf("could not index: %w", err)
	}
	return nil
}

// nolint:unparam
func (c *ChainIndexer) getLatestBlock(ctx context.Context, confirmations bool) (*uint64, error) {
	var currentBlock uint64
	var err error
	b := createBackoff()
	timeout := time.Duration(0)
	for {
		select {
		case <-ctx.Done():

			return nil, fmt.Errorf("%s context canceled: %w", ctx.Value(chainContextKey), ctx.Err())
		case <-time.After(timeout):
			currentBlock, err = c.client[0].BlockNumber(ctx)

			if err != nil {
				timeout = b.Duration()
				logger.ReportScribeError(err, c.chainID, logger.GetBlockError)
				continue
			}
			if confirmations {
				currentBlock -= c.chainConfig.Confirmations
			}
		}

		return &currentBlock, nil
	}
}

// IndexToBlock takes a contract indexer and indexs a contract up until it reaches the livefill threshold. This function should be generally used for calling a indexer with a single contract.
func (c *ChainIndexer) IndexToBlock(parentContext context.Context, onlyOneBlock *uint64, contractStartBlock uint64, indexer *indexer.Indexer) error {
	timeout := time.Duration(0)
	b := createBackoff()
	for {
		select {
		case <-parentContext.Done():
			return fmt.Errorf("%s chain context canceled: %w", parentContext.Value(chainContextKey), parentContext.Err())
		case <-time.After(timeout):
			var endHeight *uint64
			var err error
			startHeight, endHeight, err := c.getStartHeight(parentContext, onlyOneBlock, contractStartBlock, indexer)
			if err != nil {
				return err
			}
			err = indexer.Index(parentContext, startHeight, *endHeight)
			if err != nil {
				timeout = b.Duration()
				// if the config has set the contract to refresh at a slower rate than the timeout, use the refresh rate instead.
				if indexer.RefreshRate() > maxBackoff {
					timeout = time.Duration(indexer.RefreshRate()) * time.Second
				}
				logger.ReportIndexerError(err, indexer.GetIndexerConfig(), logger.BackfillIndexerError)
				continue
			}
			if onlyOneBlock != nil {
				return nil
			}

			livefillReady, err := c.isReadyForLivefill(parentContext, indexer)
			if err != nil {
				return fmt.Errorf("could not get last indexed: %w", err)
			}
			if livefillReady {
				return nil
			}

			timeout = time.Duration(indexer.RefreshRate()) * time.Second
		}
	}
}

func getMinFromMap(inputMap map[common.Address]uint64) uint64 {
	minValue := uint64(math.MaxUint64)

	for i := range inputMap {
		if inputMap[i] < minValue {
			minValue = inputMap[i]
		}
	}

	return minValue
}

func getAddressesFromConfig(contractConfigs []config.ContractConfig) []common.Address {
	var addresses []common.Address
	for i := range contractConfigs {
		contract := common.HexToAddress(contractConfigs[i].Address)
		addresses = append(addresses, contract)
	}

	return addresses
}

func createBackoff() *backoff.Backoff {
	return &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    time.Duration(maxBackoff) * time.Second,
	}
}

func (c *ChainIndexer) isReadyForLivefill(parentContext context.Context, indexer *indexer.Indexer) (bool, error) {
	// get last indexed to check livefill threshold
	lastBlockIndexed, err := c.eventDB.RetrieveLastIndexed(parentContext, indexer.GetIndexerConfig().Addresses[0], c.chainConfig.ChainID, false)
	if err != nil {
		return false, fmt.Errorf("could not get last indexed: %w", err)
	}
	endHeight, err := c.getLatestBlock(parentContext, true)
	if err != nil {
		return false, fmt.Errorf("could not get current block number while indexing: %w", err)
	}
	return int64(lastBlockIndexed) >= int64(*endHeight)-int64(c.chainConfig.LivefillThreshold), nil
}

func (c *ChainIndexer) getStartHeight(parentContext context.Context, onlyOneBlock *uint64, givenStart uint64, indexer *indexer.Indexer) (uint64, *uint64, error) {
	lastIndexed, err := c.eventDB.RetrieveLastIndexed(parentContext, indexer.GetIndexerConfig().Addresses[0], c.chainConfig.ChainID, false)
	if err != nil {
		return 0, nil, fmt.Errorf("could not get last block indexed: %w", err)
	}

	// If the last indexed block is greater than the contract start block, start indexing from the last indexed block.
	startHeight := givenStart
	if lastIndexed > startHeight {
		startHeight = lastIndexed + 1
	}

	var endHeight *uint64
	// onlyOneBlock is used for amending single blocks with a blockhash discrepancies or for testing.
	if onlyOneBlock != nil {
		startHeight = *onlyOneBlock
		endHeight = onlyOneBlock
	} else {
		endHeight, err = c.getLatestBlock(parentContext, true)
		if err != nil {
			return 0, nil, fmt.Errorf("could not get current block number while indexing: %w", err)
		}
	}

	return startHeight, endHeight, nil
}

// LivefillToTip stores data for all contracts all the way to the tip in a separate table.
//
// nolint:cyclop
func (c *ChainIndexer) LivefillToTip(parentContext context.Context) error {
	timeout := time.Duration(0)
	b := createBackoff()
	addresses := getAddressesFromConfig(c.chainConfig.Contracts)
	tipLivefillBlockMeter, err := c.handler.Meter().NewHistogram(fmt.Sprintf("scribe_block_meter_%d_tip_livefill", c.chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
	if err != nil {
		return fmt.Errorf("error creating otel histogram %w", err)
	}

	tipLivefillIndexer, err := indexer.NewIndexer(c.chainConfig, addresses, c.eventDB, c.client, c.handler, tipLivefillBlockMeter, true)
	if err != nil {
		return fmt.Errorf("could not create contract indexer: %w", err)
	}
	flushDuration := time.Duration(c.chainConfig.LivefillFlushInterval) * time.Second
	for {
		select {
		case <-parentContext.Done():
			return fmt.Errorf("context canceled: %w", parentContext.Err())
		case <-time.After(flushDuration):
			deleteBefore := time.Now().Add(-flushDuration).UnixNano()
			err := c.eventDB.FlushLogsFromHead(parentContext, deleteBefore)
			if err != nil {
				return fmt.Errorf("could not flush logs from head: %w", err)
			}
		case <-time.After(timeout):

			endHeight, err := c.getLatestBlock(parentContext, false)
			if err != nil {
				logger.ReportIndexerError(err, tipLivefillIndexer.GetIndexerConfig(), logger.GetBlockError)
				timeout = b.Duration()
				continue
			}

			tipLivefillLastIndexed, err := c.eventDB.RetrieveLastIndexed(parentContext, common.BigToAddress(big.NewInt(0)), c.chainConfig.ChainID, false)
			if err != nil {
				logger.ReportIndexerError(err, tipLivefillIndexer.GetIndexerConfig(), logger.LivefillIndexerError)
				timeout = b.Duration()
				continue
			}
			startHeight := tipLivefillLastIndexed
			if startHeight == 0 {
				startHeight = *endHeight - c.chainConfig.Confirmations
			}

			err = tipLivefillIndexer.Index(parentContext, startHeight, *endHeight)
			if err != nil {
				timeout = b.Duration()
				logger.ReportIndexerError(err, tipLivefillIndexer.GetIndexerConfig(), logger.LivefillIndexerError)
				continue
			}

			// Default refresh rate for tip livefill is 1 second.
			timeout = 1 * time.Second
		}
	}
}
