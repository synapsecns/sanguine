package scribe

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/scribe/indexer"

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
	// contractIndexers is the list of contract indexers.
	contractIndexers []*indexer.Indexer
	// startHeights is a map from address -> start height.
	startHeights map[string]uint64
	// minBlockHeight is the minimum block height to store block time for.
	minBlockHeight uint64
	// chainConfig is the config for the indexer.
	chainConfig config.ChainConfig
	// handler is the metrics handler for the scribe.
	handler metrics.Handler
	// blockHeightMeters is a map from address -> meter for block height.
	blockHeightMeters map[common.Address]metric.Int64Histogram
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
	var contractIndexers []*indexer.Indexer

	startHeights := make(map[string]uint64)

	if chainConfig.GetLogsRange == 0 {
		chainConfig.GetLogsRange = 600
	}

	if chainConfig.GetLogsBatchAmount == 0 {
		chainConfig.GetLogsBatchAmount = 2
	}

	if chainConfig.StoreConcurrency == 0 {
		chainConfig.StoreConcurrency = 20
	}

	if chainConfig.Confirmations == 0 {
		chainConfig.Confirmations = 2000
	}

	if chainConfig.ConcurrencyThreshold == 0 {
		chainConfig.ConcurrencyThreshold = 50000
	}
	minBlockHeight := uint64(math.MaxUint64)

	blockHeightMeterMap := make(map[common.Address]metric.Int64Histogram)
	for _, contract := range chainConfig.Contracts {
		blockHeightMeter, err := handler.Meter().NewHistogram(fmt.Sprintf("scribe_block_meter_%d_%s", chainConfig.ChainID, contract.Address), "block_histogram", "a block height meter", "blocks")
		if err != nil {
			return nil, fmt.Errorf("error creating otel histogram %w", err)
		}
		blockHeightMeterMap[common.HexToAddress(contract.Address)] = blockHeightMeter

		addresses := []common.Address{common.HexToAddress(contract.Address)}
		contractIndexer, err := indexer.NewIndexer(chainConfig, addresses, eventDB, client, handler, blockHeightMeter)
		if err != nil {
			return nil, fmt.Errorf("could not create contract indexer: %w", err)
		}
		contractIndexers = append(contractIndexers, contractIndexer)
		startHeights[contract.Address] = contract.StartBlock

		if minBlockHeight > contract.StartBlock {
			minBlockHeight = contract.StartBlock
		}
	}

	return &ChainIndexer{
		chainID:           chainConfig.ChainID,
		eventDB:           eventDB,
		client:            client,
		contractIndexers:  contractIndexers,
		blockHeightMeters: blockHeightMeterMap,
		startHeights:      startHeights,
		minBlockHeight:    minBlockHeight,
		chainConfig:       chainConfig,
		handler:           handler,
	}, nil
}

// Index iterates over each contract indexer and calls Index concurrently on each one.
// If `onlyOneBlock` is true, the indexer will only index the block at `currentBlock`.
//
//nolint:gocognit,cyclop
func (c ChainIndexer) Index(ctx context.Context, onlyOneBlock *uint64) error {
	// Create a new context for the chain so all chains don't halt when indexing is completed.
	chainCtx := context.WithValue(ctx, chainContextKey, fmt.Sprintf("%d-%d", c.chainID, c.minBlockHeight))
	indexGroup, indexCtx := errgroup.WithContext(chainCtx)

	var livefillContracts []config.ContractConfig
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
	for i := range c.chainConfig.Contracts {
		contract := c.chainConfig.Contracts[i]
		startHeight := contract.StartBlock
		contractAddress := common.HexToAddress(contract.Address)
		lastIndexed := lastIndexedMap[contractAddress]
		if lastIndexed > startHeight {
			startHeight = lastIndexed + 1
		}

		// Does not consider if the config's start block is within the livefill threshold for simplicity. In this case, a indexer will bring the contract to head and it will be passed to livefill.
		if *latestBlock-c.chainConfig.Confirmations > lastIndexed {
			livefillContracts = append(livefillContracts, contract)
			continue
		}

		// If current contract is not within the livefill threshold, start a indexer for it.
		contractIndexer, err := indexer.NewIndexer(c.chainConfig, []common.Address{contractAddress}, c.eventDB, c.client, c.handler, c.blockHeightMeters[contractAddress])
		if err != nil {
			return fmt.Errorf("could not create contract indexer: %w", err)
		}

		indexGroup.Go(func() error {
			err := c.IndexToBlock(indexCtx, onlyOneBlock, startHeight, contractIndexer)
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
		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    1 * time.Second,
			Max:    time.Duration(maxBackoff) * time.Second,
		}
		livefillBlockMeter, err := c.handler.Meter().NewHistogram(fmt.Sprintf("scribe_block_meter_%d_livefill", c.chainConfig.ChainID), "block_histogram", "a block height meter", "blocks")
		if err != nil {
			return fmt.Errorf("error creating otel histogram %w", err)
		}

		livefillIndexer, err := indexer.NewIndexer(c.chainConfig, getAddressesFromConfig(livefillContracts), c.eventDB, c.client, c.handler, livefillBlockMeter)
		if err != nil {
			return fmt.Errorf("could not create contract indexer: %w", err)
		}
		for {
			select {
			case <-indexCtx.Done():
				return fmt.Errorf("%s chain context canceled: %w", indexCtx.Value(chainContextKey), indexCtx.Err())
			case newLivefillContract := <-readyToLivefill:
				livefillContracts = append(livefillContracts, newLivefillContract)
				// Update indxer's config to include new contract.
				livefillIndexer.UpdateAddress(getAddressesFromConfig(livefillContracts))
			case <-time.After(timeout):
				var endHeight *uint64
				var err error
				livefillLastIndexed, err := c.eventDB.RetrieveLastIndexedMultiple(chainCtx, contractAddresses, c.chainConfig.ChainID)
				if err != nil {
					//LogEvent(ErrorLevel, "Could not get last indexed map", LogData{"cid": c.chainID, "e": err.Error()})
					timeout = b.Duration()
					continue
				}

				startHeight := getMinFromMap(livefillLastIndexed)
				endHeight, err = c.getLatestBlock(indexCtx, true)

				err = livefillIndexer.Index(indexCtx, startHeight, *endHeight)
				if err != nil {
					timeout = b.Duration()
					//LogEvent(ErrorLevel, "Could not livefill contracts, retrying", LogData{"cid": c.chainID, "ca": livefillContracts, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error()})
					continue
				}

				timeout = 1 * time.Second
				//LogEvent(InfoLevel, "Continuing to livefill contracts", LogData{"t": timeout, "cid": c.chainID, "ca": livefillIndexer, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt()})
			}
		}
	})

	if err := indexGroup.Wait(); err != nil {
		return fmt.Errorf("could not index: %w", err)
	}
	//LogEvent(WarnLevel, "Finished indexing blocktimes and contracts", LogData{"cid": c.chainID, "t": time.Since(startTime).Hours()})

	return nil
}

func (c *ChainIndexer) getLatestBlock(ctx context.Context, confirmations bool) (*uint64, error) {
	var currentBlock uint64
	var err error
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    time.Duration(maxBackoff) * time.Second,
	}
	timeout := time.Duration(0)
	for {
		select {
		case <-ctx.Done():

			return nil, fmt.Errorf("%s context canceled: %w", ctx.Value(chainContextKey), ctx.Err())
		case <-time.After(timeout):
			currentBlock, err = c.client[0].BlockNumber(ctx)

			if err != nil {
				timeout = b.Duration()
				//LogEvent(InfoLevel, "Could not get block number, bad connection to rpc likely", LogData{"cid": c.chainID, "e": err.Error()})
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
func (c *ChainIndexer) IndexToBlock(parentContext context.Context, onlyOneBlock *uint64, startHeight uint64, indexer *indexer.Indexer) error {
	timeout := time.Duration(0)
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    time.Duration(maxBackoff) * time.Second,
	}
	for {
		select {
		case <-parentContext.Done():
			return fmt.Errorf("%s chain context canceled: %w", parentContext.Value(chainContextKey), parentContext.Err())
		case <-time.After(timeout):
			var endHeight *uint64
			var err error

			// onlyOneBlock is used for amending single blocks with a blockhash discrepancies or for testing.
			if onlyOneBlock != nil {
				startHeight = *onlyOneBlock
				endHeight = onlyOneBlock
			} else {
				endHeight, err = c.getLatestBlock(parentContext, true)
				if err != nil {
					return fmt.Errorf("could not get current block number while indexing: %w", err)
				}
			}

			err = indexer.Index(parentContext, startHeight, *endHeight)
			if err != nil {
				timeout = b.Duration()
				// if the config has set the contract to refresh at a slower rate than the timeout, use the refresh rate instead.
				if indexer.RefreshRate() > maxBackoff {
					timeout = time.Duration(indexer.RefreshRate()) * time.Second
				}
				//LogEvent(ErrorLevel, "Could not index contract, retrying", LogData{"cid": c.chainID, "ca": indexer.indexerConfig.Contracts, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error()})
				continue
			}

			// get last indexed to check livefill threshold
			lastBlockIndexed, err := c.eventDB.RetrieveLastIndexed(parentContext, indexer.IndexerConfig().Contracts[0], c.chainConfig.ChainID)
			if err != nil {
				return fmt.Errorf("could not get last indexed: %w", err)
			}
			endHeight, err = c.getLatestBlock(parentContext, true)
			if err != nil {
				return fmt.Errorf("could not get current block number while indexing: %w", err)
			}
			if lastBlockIndexed > *endHeight-c.chainConfig.Confirmations {
				return nil
			}

			timeout = time.Duration(indexer.RefreshRate()) * time.Second
			//LogEvent(InfoLevel, "Continuing to livefill contract", LogData{"t": timeout, "cid": c.chainID, "ca": indexer.indexerConfig.Contracts, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt()})
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
