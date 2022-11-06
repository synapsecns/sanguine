package backfill

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ChainBackfiller is a backfiller that fetches logs for a chain. It aggregates logs
// from a slice of ContractBackfillers.
type ChainBackfiller struct {
	// chainID is the chain ID of the chain
	chainID uint32
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client []ScribeBackend
	// contractBackfillers is the list of contract backfillers
	contractBackfillers []*ContractBackfiller
	// startHeights is a map from address -> start height
	startHeights map[string]uint64
	// minBlockHeight is the minimum block height to store block time for
	minBlockHeight uint64
	// chainConfig is the config for the backfiller
	chainConfig config.ChainConfig
	// BlockTimeChunkCount is the number of chunks to process at a time.
	BlockTimeChunkCount uint64
	// BlockTimeChunkSize is the number of blocks to process per chunk.
	BlockTimeChunkSize uint64
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(chainID uint32, eventDB db.EventDB, client []ScribeBackend, chainConfig config.ChainConfig) (*ChainBackfiller, error) {
	// initialize the list of contract backfillers
	contractBackfillers := []*ContractBackfiller{}
	// initialize each contract backfiller and start heights
	startHeights := make(map[string]uint64)

	if chainConfig.BlockTimeChunkCount == 0 {
		chainConfig.BlockTimeChunkCount = 10
	}

	// start with max uint64
	minBlockHeight := uint64(math.MaxUint64)
	for _, contract := range chainConfig.Contracts {
		contractBackfiller, err := NewContractBackfiller(chainConfig.ChainID, contract.Address, eventDB, client)
		if err != nil {
			return nil, fmt.Errorf("could not create contract backfiller: %w", err)
		}
		contractBackfillers = append(contractBackfillers, contractBackfiller)
		startHeights[contract.Address] = contract.StartBlock

		// Compare if current minBlockHeight is greater than current StartBlock set in the yaml
		if minBlockHeight > contract.StartBlock {
			minBlockHeight = contract.StartBlock
		}
	}

	return &ChainBackfiller{
		chainID:             chainID,
		eventDB:             eventDB,
		client:              client,
		contractBackfillers: contractBackfillers,
		startHeights:        startHeights,
		minBlockHeight:      minBlockHeight,
		chainConfig:         chainConfig,
		BlockTimeChunkCount: chainConfig.BlockTimeChunkCount,
		BlockTimeChunkSize:  20,
	}, nil
}

// Backfill iterates over each contract backfiller and calls Backfill concurrently on each one.
// If `onlyOneBlock` is true, the backfiller will only backfill the block at `currentBlock`.
//
//nolint:gocognit,cyclop
func (c ChainBackfiller) Backfill(ctx context.Context, onlyOneBlock bool) error {
	// initialize the errgroups for backfilling contracts and getting latest blocknumber.
	gBackfill, groupCtxBackfill := errgroup.WithContext(ctx)
	// backoff in the case of an error
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    30 * time.Second,
	}

	// Starting with 0 time out.
	timeout := time.Duration(0)

	// Init currentBlock
	var currentBlock uint64
	var err error

	// Retry until block height for the current chain is retrieved.
	if !onlyOneBlock {
		for {
			select {
			case <-groupCtxBackfill.Done():
				return fmt.Errorf("context canceled: %w", groupCtxBackfill.Err())
			case <-time.After(timeout):
				// get the end height for the backfill
				currentBlock, err = c.client[0].BlockNumber(groupCtxBackfill)
				if err != nil {
					timeout = b.Duration()
					logger.Warnf("could not get block number, bad connection to rpc likely: %v", err)
					continue
				}
			}
			b.Reset()
			break
		}
	} else {
		currentBlock = c.minBlockHeight + 1
	}

	// iterate over each contract backfiller
	for i := range c.contractBackfillers {
		// capture func literal
		contractBackfiller := c.contractBackfillers[i]
		// get the start height for the backfill
		startHeight := c.startHeights[contractBackfiller.address]
		if onlyOneBlock {
			currentBlock = startHeight
		}
		logger.Infof("Starting backfilling contracts on %d up to block %d ", c.chainID, currentBlock)

		// call Backfill concurrently
		gBackfill.Go(func() error {
			// timeout should always be 0 on the first attempt
			timeout = time.Duration(0)
			for {
				select {
				case <-groupCtxBackfill.Done():
					logger.Warnf("could not backfill data: %v\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", groupCtxBackfill.Err(), c.chainID, startHeight, currentBlock, b.Attempt(), b.Duration())
					return fmt.Errorf("context canceled: %w", groupCtxBackfill.Err())
				case <-time.After(timeout):
					err = contractBackfiller.Backfill(groupCtxBackfill, startHeight, currentBlock)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not backfill data: %v\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", err, c.chainID, startHeight, currentBlock, b.Attempt(), b.Duration())
						continue
					}
					// Reset backoff and timeout
					b.Reset()
					timeout = time.Duration(0)
					return nil
				}
			}
		})
	}

	// Initialize the errgroup for backfilling block times
	gBlockTime, groupCtxBlockTime := errgroup.WithContext(ctx)

	// Set the start height to the min height from all start blocks set in config
	startHeight := c.minBlockHeight

	// Set the end height to the latest block height
	endHeight := currentBlock

	// Check if there are any block times stored in the database for the given chain
	count, err := c.eventDB.RetrieveBlockTimesCountForChain(groupCtxBlockTime, c.chainID)
	if err != nil {
		return fmt.Errorf("could not retrieve block times count for chain: %w", err)
	}

	// Create another backfiller to start from the last stored block time if there are any block times stored.
	if count > 0 {
		loggerBlocktime.Warnf("creating additional backfiller to start at last stored blocktime on chain %d", c.chainID)

		// Set the second backfiller's start height to the last stored block time
		// This will also be used as the start height for this additional backfiller
		endHeight, err = c.eventDB.RetrieveLastBlockStored(groupCtxBlockTime, c.chainID)
		if err != nil {
			return fmt.Errorf("could not retrieve last block stored: %w", err)
		}

		// Get first stored block time to compare with the current start height
		firstStoredBlockTime, err := c.eventDB.RetrieveFirstBlockStored(groupCtxBlockTime, c.chainID)
		if err != nil {
			return fmt.Errorf("could not retrieve first block stored: %w", err)
		}

		// Get the min of the last stored blocktime and the min start block from the contracts
		if startHeight > firstStoredBlockTime {
			startHeight = firstStoredBlockTime
		}

		// Backfill from last stored block to current height
		gBlockTime.Go(func() error {
			err = c.blocktimeBackfillManager(groupCtxBlockTime, zeroCheck(endHeight), currentBlock)
			if err != nil {
				return fmt.Errorf("could not backfill block times from last stored block time: %w\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", err, c.chainID, startHeight, currentBlock, b.Attempt(), b.Duration())
			}
			return nil
		})
	}

	// Backfill from earliest block to last stored block
	gBlockTime.Go(func() error {
		err = c.blocktimeBackfillManager(groupCtxBlockTime, zeroCheck(startHeight), endHeight)
		if err != nil {
			return fmt.Errorf("could not backfill block times from min block height: %w\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", err, c.chainID, startHeight, endHeight, b.Attempt(), b.Duration())
		}
		return nil
	})

	// wait for all the blocktimes to finish
	if err := gBlockTime.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}
	loggerBlocktime.Infof("Finished backfilling blocktimes on %d up to block %d ", c.chainID, currentBlock)

	// wait for all the backfillers to finish
	if err := gBackfill.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}
	logger.Infof("Finished backfilling contracts on %d up to block %d ", c.chainID, currentBlock)
	return nil
}

func (c ChainBackfiller) blocktimeBackfillManager(ctx context.Context, startHeight uint64, endHeight uint64) error {
	// Initialize the errgroup for backfilling block times
	gBlocktimeChunkBackfiller, groupBlocktimeBackfiller := errgroup.WithContext(ctx)
	currentBlock := startHeight

	// Continue to backfill block times until the current block is greater than the end height
	for currentBlock <= endHeight {
		exitFlag := false

		// Creates a backfiller for the number of chunks specified in the config
		for i := uint64(0); i < c.BlockTimeChunkCount; i++ {
			// Set the start height for the current chunk
			chunkStartHeight := currentBlock + (i * c.BlockTimeChunkSize)

			// Set the end height for the current chunk
			chunkEndHeight := chunkStartHeight + c.BlockTimeChunkSize - 1

			// Handle if the current chunk end height is greater than the total end height
			if chunkEndHeight > endHeight {
				chunkEndHeight = endHeight
				exitFlag = true
			}

			// Create a new backfiller for the current chunk
			gBlocktimeChunkBackfiller.Go(func() error {
				err := c.blocktimeBackfiller(groupBlocktimeBackfiller, chunkStartHeight, chunkEndHeight)
				if err != nil {
					return fmt.Errorf("could not backfill chunk : %w", err)
				}
				return nil
			})

			// Prevents any unnecessary backfillers from being created since we have reached the end height.
			if exitFlag {
				break
			}
		}
		// Wait for all the backfillers to finish
		if err := gBlocktimeChunkBackfiller.Wait(); err != nil {
			return fmt.Errorf("could not backfill: %w", err)
		}
		// Calculate the last block stored for logging, storing, and setting the next current block.
		lastBlockStored := currentBlock + (c.BlockTimeChunkCount * c.BlockTimeChunkSize) - 1
		logger.Infof("Finished backfilling chunks on %d from block %d up to block %d ", c.chainID, currentBlock, lastBlockStored)

		// store the last block time
		err := c.eventDB.StoreLastBlockTime(ctx, c.chainID, lastBlockStored)
		if err != nil {
			loggerBlocktime.Warnf("could not store last block time %s: %v\nChain: %d\nCurrent Block: %d, Last Block: %d", err, c.chainID, currentBlock, lastBlockStored)
		}

		// Increment the current block by the number of chunks just backfilled.
		currentBlock = lastBlockStored + 1
	}
	return nil
}

func (c ChainBackfiller) blocktimeBackfiller(ctx context.Context, startHeight uint64, endHeight uint64) error {
	// Init backoff for backfilling block times
	bBlockNum := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    30 * time.Second,
	}

	// timeout should always be 0 on the first attempt
	timeoutBlockNum := time.Duration(0)

	// Current block
	blockNum := startHeight
	loggerBlocktime.Infof("Starting backfilling blocktimes on %d from block %d  to block %d ", c.chainID, startHeight, endHeight)

	for {
		select {
		case <-ctx.Done():
			loggerBlocktime.Warnf("gBlockTime context canceled %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), ctx.Err(), c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
			return fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeoutBlockNum):
			// Check if the current block's already exists in database.
			_, err := c.eventDB.RetrieveBlockTime(ctx, c.chainID, blockNum)
			if err == nil {
				loggerBlocktime.Infof("skipping storing blocktime for block %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), err, c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
				blockNum++
				// Make sure the count doesn't increase unnecessarily.
				bBlockNum.Reset()
				continue
			}

			// Get information on the current block for further processing.
			rawBlock, err := c.client[0].HeaderByNumber(ctx, big.NewInt(int64(blockNum)))
			if err != nil {
				timeoutBlockNum = bBlockNum.Duration()
				loggerBlocktime.Warnf("could not get block time at block %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), err, c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
				continue
			}

			// Store the block time with the block retrieved above.
			err = c.eventDB.StoreBlockTime(ctx, c.chainID, blockNum, rawBlock.Time)
			if err != nil {
				timeoutBlockNum = bBlockNum.Duration()
				loggerBlocktime.Warnf("could not store block time - block %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), err, c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
				continue
			}

			// Move on to the next block.
			blockNum++

			// Reset the backoff after successful block parse run to prevent bloated back offs.
			bBlockNum.Reset()
			timeoutBlockNum = time.Duration(0)

			// If done with the range, exit go routine.
			if blockNum > endHeight {
				loggerBlocktime.Infof("Exiting backfill on chain %d on block %d ", c.chainID, blockNum)
				return nil
			}
		}
	}
}

// Used for setting start heights for backfilling.
func zeroCheck(value uint64) uint64 {
	if value > 0 {
		return value - 1
	}
	return value
}
