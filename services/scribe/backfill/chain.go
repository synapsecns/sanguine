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
	client ScribeBackend
	// contractBackfillers is the list of contract backfillers
	contractBackfillers []*ContractBackfiller
	// startHeights is a map from address -> start height
	startHeights map[string]uint64
	// minBlockHeight is the minimum block height to store block time for
	minBlockHeight uint64
	// chainConfig is the config for the backfiller
	chainConfig config.ChainConfig
	// blockTimeBuffer is the buffer for block times queries
	blockTimeBuffer []rpc.BatchElem
	// blockNumAttemptCount is a map from blockNum -> attempt count
	blockNumAttemptCount map[uint64]uint32
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(chainID uint32, eventDB db.EventDB, client ScribeBackend, chainConfig config.ChainConfig) (*ChainBackfiller, error) {
	// initialize the list of contract backfillers
	contractBackfillers := []*ContractBackfiller{}
	// initialize each contract backfiller and start heights
	startHeights := make(map[string]uint64)
	// initialize the block time attempt count and buffer
	blockNumAttemptCount := make(map[uint64]uint32)
	blockTimeBuffer := []rpc.BatchElem{}

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
		chainID:              chainID,
		eventDB:              eventDB,
		client:               client,
		contractBackfillers:  contractBackfillers,
		startHeights:         startHeights,
		minBlockHeight:       minBlockHeight,
		chainConfig:          chainConfig,
		blockTimeBuffer:      blockTimeBuffer,
		blockNumAttemptCount: blockNumAttemptCount,
	}, nil
}

// Backfill iterates over each contract backfiller and calls Backfill concurrently on each one.
// If `onlyOneBlock` is true, the backfiller will only backfill the block at `endHeight`.
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

	// Init endHeight
	var endHeight uint64
	var err error

	// Retry until block height for the current chain is retrieved.
	if !onlyOneBlock {
		for {
			select {
			case <-groupCtxBackfill.Done():
				return fmt.Errorf("context canceled: %w", groupCtxBackfill.Err())
			case <-time.After(timeout):
				// get the end height for the backfill
				endHeight, err = c.client.BlockNumber(groupCtxBackfill)
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
		endHeight = c.minBlockHeight + 1
	}

	// iterate over each contract backfiller
	for i := range c.contractBackfillers {
		// capture func literal
		contractBackfiller := c.contractBackfillers[i]
		// get the start height for the backfill
		startHeight := c.startHeights[contractBackfiller.address]
		if onlyOneBlock {
			endHeight = startHeight
		}
		// call Backfill concurrently
		gBackfill.Go(func() error {
			// timeout should always be 0 on the first attempt
			timeout = time.Duration(0)
			for {
				select {
				case <-groupCtxBackfill.Done():
					logger.Warnf("could not backfill data: %v\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", groupCtxBackfill.Err(), c.chainID, startHeight, endHeight, b.Attempt(), b.Duration())
					return fmt.Errorf("context canceled: %w", groupCtxBackfill.Err())
				case <-time.After(timeout):
					err = contractBackfiller.Backfill(groupCtxBackfill, startHeight, endHeight)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not backfill data: %v\nChain: %d\nStart Block: %d\nEnd Block: %d\nBackoff Atempts: %f\nBackoff Duration: %d", err, c.chainID, startHeight, endHeight, b.Attempt(), b.Duration())
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

	// Backfill the block times

	// Start the buffer handler for the batch block time queries
	go func() {
		err := c.bufferHandler(ctx, &c.blockTimeBuffer)
		if err != nil {
			logger.Errorf("could not handle buffer: %v", err)
		}
	}()

	// Set the start height to the minimum block height of all contracts
	startHeight := c.minBlockHeight

	// Start at the block before the minimum block height
	if startHeight != 0 {
		startHeight--
	}

	// Current block
	blockNum := startHeight
blockTimeLoop:
	for {
		select {
		case <-ctx.Done():
			logger.Warnf("context canceled %s: %v\nChain: %d\nBlock: %d", big.NewInt(int64(blockNum)).String(), ctx.Err(), c.chainID, blockNum)
			return fmt.Errorf("context canceled: %w", ctx.Err())
		default:
			// Check if the current block's already exists in database.
			_, err = c.eventDB.RetrieveBlockTime(ctx, c.chainID, blockNum)
			if err == nil {
				logger.Infof("skipping storing blocktime for block %s: %v\nChain: %d\nBlock: %d", big.NewInt(int64(blockNum)).String(), err, c.chainID, blockNum)
				blockNum++
				continue
			}

			// Store the block time info in the blockTimeBuffer for batch querying
			var batchElemErr error
			var head *types.Header
			batchElem := rpc.BatchElem{
				Method: "eth_getBlockByNumber",
				Args:   []interface{}{hexutil.EncodeBig(big.NewInt(int64(blockNum))), false},
				Result: &head,
				Error:  batchElemErr,
			}
			c.blockTimeBuffer = append(c.blockTimeBuffer, batchElem)
			fmt.Println("appending to buffer", blockNum)
			fmt.Println("CMON", len(c.blockTimeBuffer))

			// Move on to the next block.
			blockNum++

			// If done with the range, exit the loop.
			if blockNum > endHeight {
				break blockTimeLoop
			}
		}
	}

	// wait for all the backfillers to finish
	if err = gBackfill.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}
	return nil
}

// bufferHandler will be run concurrently as it takes in BatchElems until it hits a
// certain threshold, then it will process the BatchElems and stores the results in
// the eventDB.
func (c ChainBackfiller) bufferHandler(ctx context.Context, blockTimeBuffer *[]rpc.BatchElem) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled: %w", ctx.Err())
		default:
			fmt.Println("REALLY HOPE THIS NUM ISN'T ZERO:", uint32(len(*blockTimeBuffer)))
			if uint32(len(*blockTimeBuffer)) >= c.chainConfig.BlockTimeBatchSize {
				fmt.Println("HERE")

				// batch call to get the block times
				err := c.client.BatchCallContext(ctx, *blockTimeBuffer)
				if err != nil {
					return fmt.Errorf("could not batch call: %w", err)
				}
				// iterate over the batch elems that have results and no errors,
				// then store the results in the eventDB.
				for i := 0; i < len(*blockTimeBuffer); i++ {
					elem := &(*blockTimeBuffer)[i]
					// store the block time
					blockNum, err := hexutil.DecodeBig(elem.Args[0].(string))
					if err != nil {
						return fmt.Errorf("could not decode block number: %w", err)
					}
					header := elem.Result.(**types.Header)
					_ = header
					c.blockNumAttemptCount[blockNum.Uint64()]++
					fmt.Println("SOMESOMESOME", c.blockNumAttemptCount[blockNum.Uint64()])
					if elem.Error != nil {
						logger.Warnf("could not get block time: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %d", elem.Error, c.chainID, blockNum.Uint64(), c.blockNumAttemptCount[blockNum.Uint64()])
						continue
						// todo
					}
					if elem.Result == nil {
						continue
					}
					err = c.eventDB.StoreBlockTime(ctx, c.chainID, blockNum.Uint64(), (**header).Time)
					//if err != nil {
					//	logger.Warnf("could not store block time - block %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %d", blockNum.String(), err, c.chainID, blockNum, c.blockNumAttemptCount[blockNum.Uint64()])
					//	continue
					//}

					// store the last block time
					err = c.eventDB.StoreLastBlockTime(ctx, c.chainID, blockNum.Uint64())
					if err != nil {
						logger.Warnf("could not store last block time %s: %v\nChain: %d\nBlock: %d\nBackoff Atempts: %d", blockNum.String(), err, c.chainID, blockNum, c.blockNumAttemptCount[blockNum.Uint64()])
						continue
					}
					// remove the batch elem from the buffer since it has been processed
					tempBuffer := remove(*blockTimeBuffer, i)
					blockTimeBuffer = &tempBuffer
					i--
				}
			}
		}
	}
}

func remove(slice []rpc.BatchElem, s int) []rpc.BatchElem {
	return append(slice[:s], slice[s+1:]...)
}
