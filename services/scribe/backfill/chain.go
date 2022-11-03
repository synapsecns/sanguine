package backfill

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
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
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(chainID uint32, eventDB db.EventDB, client []ScribeBackend, chainConfig config.ChainConfig) (*ChainBackfiller, error) {
	// initialize the list of contract backfillers
	contractBackfillers := []*ContractBackfiller{}
	// initialize each contract backfiller and start heights
	startHeights := make(map[string]uint64)

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
			err = c.backfillBlockTimes(groupCtxBlockTime, zeroCheck(endHeight), currentBlock)
			if err != nil {
				return fmt.Errorf("could not backfill block times from last stored block time: %w\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", err, c.chainID, startHeight, currentBlock, b.Attempt(), b.Duration())
			}
			return nil
		})
	}

	// Backfill from earliest block to last stored block
	gBlockTime.Go(func() error {
		err = c.backfillBlockTimes(groupCtxBlockTime, zeroCheck(startHeight), endHeight)
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

func (c ChainBackfiller) backfillBlockTimes(ctx context.Context, startHeight uint64, endHeight uint64) error {
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
			tempBatchLimit := 10
			var batchArr []rpc.BatchElem
			for batchIdx := 0; batchIdx <= batchIdx+tempBatchLimit; batchIdx++ {
				// Check if the current block's already exists in database.
				_, err := c.eventDB.RetrieveBlockTime(ctx, c.chainID, blockNum)
				if err == nil {
					loggerBlocktime.Infof("skipping storing blocktime for block %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), err, c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
					blockNum++
					// Make sure the count doesn't increase unnecessarily.
					bBlockNum.Reset()
					continue
				}

				// Store the block time info in the batchArr for batch querying
				var batchElemErr error
				var head *types.Header
				batchElem := rpc.BatchElem{
					Method: "eth_getBlockByNumber",
					Args:   []interface{}{hexutil.EncodeBig(big.NewInt(int64(blockNum))), false},
					Result: &head,
					Error:  batchElemErr,
				}
				batchArr = append(batchArr, batchElem)

			}
			err := c.client[0].BatchCallContext(ctx, batchArr)
			if err != nil {
				timeoutBlockNum = bBlockNum.Duration()
				loggerBlocktime.Warnf("could not get block time at block %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), err, c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
				continue
			}

			// Iterate through the batchArr and store the block time info in the database
			for _, batchElem := range batchArr {
				if batchElem.Error != nil {
					timeoutBlockNum = bBlockNum.Duration()
					loggerBlocktime.Warnf("could not get block time at block %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), batchElem.Error, c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
					continue
				}
				head := batchElem.Result.(*types.Header)
				err = c.eventDB.StoreBlockTime(ctx, c.chainID, blockNum, head.Time)
				if err != nil {
					timeoutBlockNum = bBlockNum.Duration()
					loggerBlocktime.Warnf("could not store block time at block %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), err, c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
					continue
				}
				loggerBlocktime.Infof("stored block time at block %s\nChain: %d\nBlock: %d\nBackoff Atempts: %f\nBackoff Duration: %d", big.NewInt(int64(blockNum)).String(), c.chainID, blockNum, bBlockNum.Attempt(), bBlockNum.Duration())
				blockNum++
				// Make sure the count doesn't increase unnecessarily.
				bBlockNum.Reset()
			}

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
