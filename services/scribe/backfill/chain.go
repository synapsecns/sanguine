package backfill

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ChainBackfiller is a backfiller that fetches logs for a chain. It aggregates logs
// from a slice of ContractBackfillers.
type ChainBackfiller struct {
	// chainID is the chain ID of the chain.
	chainID uint32
	// eventDB is the database to store event data in.
	eventDB db.EventDB
	// client contains the clients used for backfilling.
	client []ScribeBackend
	// contractBackfillers is the list of contract backfillers.
	contractBackfillers []*ContractBackfiller
	// startHeights is a map from address -> start height.
	startHeights map[string]uint64
	// minBlockHeight is the minimum block height to store block time for.
	minBlockHeight uint64
	// chainConfig is the config for the backfiller.
	chainConfig config.ChainConfig
}

// Used for handling logging of various context types.
type contextKey int

const (
	chainContextKey contextKey = iota
	blocktimeContextKey
)

// NewChainBackfiller creates a new backfiller for a chain. This is done by passing through all the function parameters
// into the ChainBackfiller struct, as well as iterating through all the contracts in the chain config and creating
// ContractBackfillers for each contract.
func NewChainBackfiller(chainID uint32, eventDB db.EventDB, client []ScribeBackend, chainConfig config.ChainConfig) (*ChainBackfiller, error) {
	var contractBackfillers []*ContractBackfiller

	startHeights := make(map[string]uint64)

	if chainConfig.BlockTimeChunkCount == 0 {
		chainConfig.BlockTimeChunkCount = 40
	}

	if chainConfig.BlockTimeChunkSize == 0 {
		chainConfig.BlockTimeChunkSize = 50
	}

	if chainConfig.ContractSubChunkSize == 0 {
		chainConfig.ContractSubChunkSize = 600
	}

	if chainConfig.ContractChunkSize == 0 {
		chainConfig.ContractChunkSize = 30000
	}

	minBlockHeight := uint64(math.MaxUint64)

	for _, contract := range chainConfig.Contracts {
		contractBackfiller, err := NewContractBackfiller(chainConfig, contract.Address, eventDB, client)

		if err != nil {
			return nil, fmt.Errorf("could not create contract backfiller: %w", err)
		}
		contractBackfillers = append(contractBackfillers, contractBackfiller)
		startHeights[contract.Address] = contract.StartBlock

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
	}, nil
}

// Backfill iterates over each contract backfiller and calls Backfill concurrently on each one.
// If `onlyOneBlock` is true, the backfiller will only backfill the block at `currentBlock`.
//
//nolint:gocognit,cyclop
func (c ChainBackfiller) Backfill(ctx context.Context, onlyOneBlock bool) error {
	var currentBlock uint64
	var err error

	// Create a new context for the chain so all chains don't halt when backfilling is completed.
	chainCtx := context.WithValue(ctx, chainContextKey, fmt.Sprintf("%d-%d", c.chainID, c.minBlockHeight))
	backfillGroup, backfillCtx := errgroup.WithContext(chainCtx)

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    10 * time.Second,
	}

	timeout := time.Duration(0)
	startTime := time.Now()

	if !onlyOneBlock {
		// Retry until block height for the current chain is retrieved.
		for {
			select {
			case <-backfillCtx.Done():
				LogEvent(ErrorLevel, "Context canceled", LogData{"cid": c.chainID, "bn": currentBlock, "bd": b.Duration(), "a": b.Attempt(), "e": backfillCtx.Err(), "bt": true})

				return fmt.Errorf("%s context canceled: %w", backfillCtx.Value(chainContextKey), backfillCtx.Err())
			case <-time.After(timeout):
				currentBlock, err = c.client[0].BlockNumber(backfillCtx)

				if err != nil {
					timeout = b.Duration()
					LogEvent(InfoLevel, "Could not get block number, bad connection to rpc likely", LogData{"cid": c.chainID, "e": err.Error()})
					continue
				}
			}
			b.Reset()

			break
		}
	} else {
		currentBlock = c.minBlockHeight + 1
	}

	for i := range c.contractBackfillers {
		contractBackfiller := c.contractBackfillers[i]
		startHeight := c.startHeights[contractBackfiller.address]

		if onlyOneBlock {
			currentBlock = startHeight
		}

		LogEvent(InfoLevel, "Starting backfilling contracts", LogData{"cid": c.chainID, "bn": currentBlock})
		backfillGroup.Go(func() error {
			timeout = time.Duration(0)
			for {
				select {
				case <-backfillCtx.Done():
					LogEvent(ErrorLevel, "Could not backfill data, context canceled", LogData{"cid": c.chainID, "ca": contractBackfiller.address, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": backfillCtx.Err()})

					return fmt.Errorf("%s chain context canceled: %w", backfillCtx.Value(chainContextKey), backfillCtx.Err())
				case <-time.After(timeout):
					err = contractBackfiller.Backfill(backfillCtx, startHeight, currentBlock)

					if err != nil {
						timeout = b.Duration()
						LogEvent(WarnLevel, "Could not backfill contract, retrying", LogData{"cid": c.chainID, "ca": contractBackfiller.address, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error()})

						continue
					}
					b.Reset()
					timeout = time.Duration(0)

					return nil
				}
			}
		})
	}

	// Set the start height to the min height from all start blocks set in config
	startHeight := c.minBlockHeight

	// Set the end height to the latest block height.
	endHeight := currentBlock

	// Check if there are any block times stored in the database for the given chain
	count, err := c.eventDB.RetrieveBlockTimesCountForChain(backfillCtx, c.chainID)
	if err != nil {
		LogEvent(ErrorLevel, "could not retrieve block times count for chain", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error(), "bt": true})

		return fmt.Errorf("could not retrieve block times count for chain: %w", err)
	}

	// Create another backfiller to start from the last stored block time if there are any block times stored.
	if count > 0 {
		LogEvent(WarnLevel, "Additional blocktime backfiller created", LogData{"cid": c.chainID, "bt": true})
		// Set the second backfiller's start height to the last stored block time.
		// This will also be used as the start height for this additional backfiller.
		endHeight, err = c.eventDB.RetrieveLastBlockStored(backfillCtx, c.chainID)
		if err != nil {
			LogEvent(ErrorLevel, "Could not retrieve last block stored", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error(), "bt": true})

			return fmt.Errorf("could not retrieve last block stored: %w", err)
		}

		// Get first stored block time to compare with the current start height.
		firstStoredBlockTime, err := c.eventDB.RetrieveFirstBlockStored(backfillCtx, c.chainID)

		if err != nil {
			LogEvent(ErrorLevel, "Could not retrieve first block stored", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error(), "bt": true})

			return fmt.Errorf("could not retrieve first block stored: %w", err)
		}

		if startHeight > firstStoredBlockTime {
			startHeight = firstStoredBlockTime
		}

		// Backfill from last stored block to current height.
		backfillGroup.Go(func() error {
			err = c.blocktimeBackfillManager(backfillCtx, decrementIfNotZero(endHeight), currentBlock)
			if err != nil {
				LogEvent(ErrorLevel, "Could not backfill block times from last stored block time", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error(), "bt": true})

				return fmt.Errorf("could not backfill block times from last stored block time: %w\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", err, c.chainID, startHeight, currentBlock, b.Attempt(), b.Duration())
			}

			LogEvent(ErrorLevel, "Completed adding later blocks", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "bt": true})

			return nil
		})
		endHeight = firstStoredBlockTime
	}

	// Backfill from the earliest block to last stored block.
	backfillGroup.Go(func() error {
		err = c.blocktimeBackfillManager(backfillCtx, decrementIfNotZero(startHeight), endHeight)
		if err != nil {
			LogEvent(ErrorLevel, "Could not backfill block times from min block height", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error(), "bt": true})

			return fmt.Errorf("could not backfill block times from min block height: %w\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", err, c.chainID, startHeight, endHeight, b.Attempt(), b.Duration())
		}

		LogEvent(WarnLevel, "Completed adding earlier blocks", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "bt": true})

		return nil
	})

	if err := backfillGroup.Wait(); err != nil {
		LogEvent(ErrorLevel, "Could not backfill with error group", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error(), "bt": true})

		return fmt.Errorf("could not backfill: %w", err)
	}
	LogEvent(WarnLevel, "Finished backfilling blocktimes and contracts", LogData{"cid": c.chainID, "sh": startHeight, "eh": currentBlock, "t": time.Since(startTime).Hours()})

	return nil
}

// blocktimeBackfillManager is a helper function to orchestrate concurrent backfilling of block times.
func (c ChainBackfiller) blocktimeBackfillManager(ctx context.Context, startHeight uint64, endHeight uint64) error {
	currentBlock := startHeight

	// Continue to backfill block times until the current block is greater than the end height.
	for currentBlock <= endHeight {
		startTime := time.Now()
		chunkIdx := uint64(0)
		LogEvent(InfoLevel, "Starting backfilling chunks", LogData{"cid": c.chainID, "bn": currentBlock, "eh": endHeight, "bt": true})

		// Create a new context for the next batch of blocktime chunks.
		blocktimeChunkCtx := context.WithValue(ctx, blocktimeContextKey, fmt.Sprintf("%d-%d-%d", c.chainID, startHeight, endHeight))

		chunkGroup, chunkCtx := errgroup.WithContext(blocktimeChunkCtx)

		for chunkIdx < c.chainConfig.BlockTimeChunkCount {
			chunkStartHeight := currentBlock + (chunkIdx * c.chainConfig.BlockTimeChunkSize)
			chunkEndHeight := chunkStartHeight + c.chainConfig.BlockTimeChunkSize - 1

			if chunkEndHeight >= endHeight {
				chunkEndHeight = endHeight

				// This prevents any unnecessary backfillers from being created since we have reached the end height.
				chunkIdx = c.chainConfig.BlockTimeChunkCount
			}

			chunkGroup.Go(func() error {
				err := c.blocktimeBackfiller(chunkCtx, chunkStartHeight, chunkEndHeight)
				if err != nil {
					LogEvent(ErrorLevel, "Could backfill chunk", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bt": true})

					return fmt.Errorf("could not backfill chunk: %w", err)
				}

				return nil
			})

			chunkIdx++
		}

		if err := chunkGroup.Wait(); err != nil {
			LogEvent(ErrorLevel, "could not backfill chain", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "bt": true})

			return fmt.Errorf("could not backfill chain %d: %w", c.chainID, err)
		}

		// Calculate the last block stored for logging, storing, and setting the next current block.
		lastBlockStored := currentBlock + (c.chainConfig.BlockTimeChunkCount * c.chainConfig.BlockTimeChunkSize) - 1
		LogEvent(InfoLevel, "Finished backfilling chunks", LogData{"cid": c.chainID, "bn": currentBlock, "lb": lastBlockStored, "bt": true, "ts": time.Since(startTime).Seconds()})

		currentBlock = lastBlockStored + 1
	}
	LogEvent(WarnLevel, "Finished backfilling all chunks", LogData{"cid": c.chainID, "bn": currentBlock, "sh": startHeight, "eh": endHeight, "bt": true})

	return nil
}

// blocktimeBackfiller is a helper function to backfill block times for a given range of blocks.
func (c ChainBackfiller) blocktimeBackfiller(ctx context.Context, startHeight uint64, endHeight uint64) error {
	bBlockNum := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    10 * time.Second,
	}

	timeoutBlockNum := time.Duration(0)

	LogEvent(InfoLevel, "Starting backfilling blocktimes", LogData{"cid": c.chainID, "sh": startHeight, "eh": endHeight, "bt": true})

RETRY:
	select {
	case <-ctx.Done():
		LogEvent(ErrorLevel, "Context canceled", LogData{"cid": c.chainID, "sh": startHeight, "eh": endHeight, "bt": true, "a": bBlockNum.Attempt(), "bd": bBlockNum.Duration(), "e": ctx.Err()})

		return fmt.Errorf("%s context canceled: %w", ctx.Value(blocktimeContextKey), ctx.Err())
	case <-time.After(timeoutBlockNum):

		res, err := BlockTimesInRange(ctx, c.client[0], startHeight, endHeight)
		if err != nil {
			LogEvent(ErrorLevel, "Could not get block times", LogData{"cid": c.chainID, "sh": startHeight, "eh": endHeight, "bt": true, "a": bBlockNum.Attempt(), "bd": bBlockNum.Duration(), "e": err.Error()})
			timeoutBlockNum = bBlockNum.Duration()

			goto RETRY
		}

		itr := res.Iterator()
		for !itr.Done() {
			blockNumIdx, blockTime, _ := itr.Next()

			// Check if the current block's already exists in database (to prevent unnecessary requests to omnirpc).
			_, err = c.eventDB.RetrieveBlockTime(ctx, c.chainID, blockNumIdx)
			if err == nil {
				LogEvent(InfoLevel, "Skipping blocktime backfill", LogData{"cid": c.chainID, "bn": blockNumIdx, "bt": true})
				continue
			}
			// Store the block time with the block retrieved above.
			err = c.eventDB.StoreBlockTime(ctx, c.chainID, blockNumIdx, blockTime)
			if err != nil {
				LogEvent(WarnLevel, "Could not store blocktime", LogData{"cid": c.chainID, "bn": blockNumIdx, "sh": startHeight, "eh": endHeight, "bt": true, "a": bBlockNum.Attempt(), "bd": bBlockNum.Duration(), "e": err.Error()})
				timeoutBlockNum = bBlockNum.Duration()

				goto RETRY
			}
		}
		LogEvent(InfoLevel, "Exiting backfill", LogData{"cid": c.chainID})
		return nil
	}
}

func decrementIfNotZero(value uint64) uint64 {
	if value > 0 {
		return value - 1
	}

	return value
}
