package backfill

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/synapsecns/sanguine/ethergo/util"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/config"

	"github.com/jpillora/backoff"
)

// RangeFilter pre-fetches filter logs into a channel in deterministic order.
type RangeFilter struct {
	// iterator is the chunk iterator used for the range.
	iterator util.ChunkIterator
	// for logging
	startBlock *big.Int
	// for logging
	endBlock *big.Int
	// logs is a channel with the filtered ahead logs. This channel is not closed
	// and the user can rely on the garbage collection behavior of RangeFilter to remove it.
	logs chan []types.Log
	// backend is the ethereum backend used to fetch logs.
	backend ScribeBackend
	// contractAddress is the contractAddress that logs are fetched for.
	contractAddress ethCommon.Address
	// doneChan is a channel that is closed when the RangeFilter has completed.
	// this is only to be used by external callers
	doneChan chan bool
	// chainConfig holds the chain config (config data for the chain)
	chainConfig *config.ChainConfig
}

// bufferSize is how many ranges ahead should be fetched.
const bufferSize = 15

// minBackoff is the minimum backoff period between requests.
var minBackoff = 1 * time.Second

// maxBackoff is the maximum backoff period between requests.
var maxBackoff = 10 * time.Second

// NewRangeFilter creates a new filtering interface for a range of blocks. If reverse is not set, block heights are filtered from start->end.
func NewRangeFilter(address ethCommon.Address, backend ScribeBackend, startBlock, endBlock *big.Int, chainConfig *config.ChainConfig) *RangeFilter {
	// The ChunkIterator is inclusive of the start and ending block resulting in potentially confusing behavior when
	// setting the range size in the config. For example, setting a range of 1 would result in two blocks being queried
	// instead of 1. This is accounted for by subtracting 1.
	chunkSize := int(chainConfig.GetLogsRange) - 1
	return &RangeFilter{
		iterator:        util.NewChunkIterator(startBlock, endBlock, chunkSize, true),
		startBlock:      startBlock,
		endBlock:        endBlock,
		logs:            make(chan []types.Log, bufferSize),
		backend:         backend,
		contractAddress: address,
		doneChan:        make(chan bool),
		chainConfig:     chainConfig,
	}
}

// closeOnDone closes the done channel when the process is finished.
func (f *RangeFilter) closeOnDone() {
	f.doneChan <- true
}

// GetChunkArr gets the appropriate amount of block chunks (getLogs ranges).
func (f *RangeFilter) GetChunkArr() (chunkArr []*util.Chunk) {
	for i := uint64(0); i < f.chainConfig.GetLogsBatchAmount; i++ {
		chunk := f.iterator.NextChunk()
		if chunk == nil {
			return chunkArr
		}
		chunkArr = append(chunkArr, chunk)

		// Stop appending chunks if the max height of the current chunk exceeds the concurrency threshold
		if chunk.EndBlock.Uint64() > f.endBlock.Uint64()-f.chainConfig.ConcurrencyThreshold {
			return chunkArr
		}
	}
	return chunkArr
}

// Start starts the log fetching process. If the context is canceled, logs will stop being filtered.
// 1. Within an infinite for loop, chunks of getLogs blocks are constructed and used to get logs. This flow is paused
// when the logs channel's buffer of 15 is reached.
// 2. Each time the logs are received, a wait group is used to ensure that there is no race condition
// where channels could be closed before a log could be saved.
// 3. When the range to get logs is completed (GetChunkArr returns a zero array), the wait group is used to ensure
// that all logs are added to the logs channel before returning and terminating the function.
// 4. Completing the Start function triggers the closeOnDone function, which sends a boolean in the done channel
// that signals that the fetcher has completed. The consumer of these logs then performs a drain to fully empty the logs
// channel. See contract.go to learn more how the logs from this file are consumed.
func (f *RangeFilter) Start(ctx context.Context) error {
	var wg sync.WaitGroup

	defer f.closeOnDone()

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				LogEvent(WarnLevel, "could not finish filtering range", LogData{"ca": f.contractAddress, "sh": f.startBlock.String(), "eh": f.endBlock.String(), "cid": &f.chainConfig.ChainID})
				return fmt.Errorf("could not finish filtering range: %w", ctx.Err())
			}

			return nil
		default:
			chunks := f.GetChunkArr()

			if len(chunks) == 0 {
				wg.Wait()
				return nil
			}
			logs, err := f.FilterLogs(ctx, chunks)
			if err != nil {
				return fmt.Errorf("could not filter logs: %w", err)
			}

			wg.Add(1)
			go func(logs []types.Log) {
				defer wg.Done()
				f.logs <- logs
			}(logs)
			LogEvent(InfoLevel, "Contract backfill chunk completed", LogData{"ca": f.contractAddress, "sh": chunks[0].MinBlock(), "eh": chunks[0].MaxBlock(), "cid": &f.chainConfig.ChainID})
		}
	}
}

// FilterLogs safely calls FilterLogs with the filtering implementing a backoff in the case of
// rate limiting and respects context cancellation.
//
// nolint:cyclop
func (f *RangeFilter) FilterLogs(ctx context.Context, chunks []*util.Chunk) ([]types.Log, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    minBackoff,
		Max:    maxBackoff,
	}

	attempt := 0
	timeout := time.Duration(0)

	// for logging purposes
	startHeight := chunks[0].StartBlock.Uint64()
	endHeight := chunks[len(chunks)-1].EndBlock.Uint64()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("could not finish filtering logs: %w", ctx.Err())
		case <-time.After(timeout):
			attempt++

			if attempt > retryTolerance {
				return nil, fmt.Errorf("maximum number of filter attempts exceeded")
			}

			res, err := GetLogsInRange(ctx, f.backend, f.contractAddress, uint64(f.chainConfig.ChainID), chunks)
			if err != nil {
				timeout = b.Duration()
				LogEvent(WarnLevel, "Could not filter logs for range, retrying", LogData{"sh": startHeight, "ca": f.contractAddress, "eh": endHeight, "cid": &f.chainConfig.ChainID, "e": err})

				continue
			}

			var logs []types.Log
			itr := res.Iterator()
			for !itr.Done() {
				select {
				case <-ctx.Done():
					return nil, fmt.Errorf("could not finish filtering logs: %w", ctx.Err())
				default:
					_, resLogChunk := itr.Next()

					if resLogChunk == nil || len(*resLogChunk) == 0 {
						LogEvent(WarnLevel, "empty subchunk", LogData{"sh": startHeight, "ca": f.contractAddress, "cid": &f.chainConfig.ChainID, "eh": endHeight})
						continue
					}
					logsChunk := *resLogChunk

					logs = append(logs, logsChunk...)
				}
			}

			return logs, nil
		}
	}
}

// Drain fetches empties the log chan. For use once the doneChan is emitted.
func (f *RangeFilter) Drain(ctx context.Context) (filteredLogs []types.Log, err error) {
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context ended: %w", ctx.Err())
		case log := <-f.logs:
			filteredLogs = append(filteredLogs, log...)
		default:
			return filteredLogs, nil
		}
	}
}
