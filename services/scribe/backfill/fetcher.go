package backfill

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/ethergo/util"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/config"

	"github.com/jpillora/backoff"
)

// LogFetcher pre-fetches filter logs into a channel in deterministic order.
type LogFetcher struct {
	// iterator is the chunk iterator used for the range.
	iterator util.ChunkIterator
	// for logging
	startBlock *big.Int
	// for logging
	endBlock *big.Int
	// fetchedLogsChan is a channel with the fetched chunks of logs.
	fetchedLogsChan chan []types.Log
	// backend is the ethereum backend used to fetch logs.
	backend ScribeBackend
	// contractAddress is the contractAddress that logs are fetched for.
	contractAddress ethCommon.Address
	// chainConfig holds the chain config (config data for the chain)
	chainConfig *config.ChainConfig
}

// bufferSize is how many getLogs*batch amount chunks ahead should be fetched.
const bufferSize = 3

// NewLogFetcher creates a new filtering interface for a range of blocks. If reverse is not set, block heights are filtered from start->end.
func NewLogFetcher(address ethCommon.Address, backend ScribeBackend, startBlock, endBlock *big.Int, chainConfig *config.ChainConfig) *LogFetcher {
	// The ChunkIterator is inclusive of the start and ending block resulting in potentially confusing behavior when
	// setting the range size in the config. For example, setting a range of 1 would result in two blocks being queried
	// instead of 1. This is accounted for by subtracting 1.
	chunkSize := int(chainConfig.GetLogsRange) - 1
	return &LogFetcher{
		iterator:        util.NewChunkIterator(startBlock, endBlock, chunkSize, true),
		startBlock:      startBlock,
		endBlock:        endBlock,
		fetchedLogsChan: make(chan []types.Log, bufferSize),
		backend:         backend,
		contractAddress: address,
		chainConfig:     chainConfig,
	}
}

// GetChunkArr gets the appropriate amount of block chunks (getLogs ranges).
func (f *LogFetcher) GetChunkArr() (chunkArr []*util.Chunk) {
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
func (f *LogFetcher) Start(ctx context.Context) error {
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
				close(f.fetchedLogsChan)
				return nil
			}
			logs, err := f.FetchLogs(ctx, chunks)
			if err != nil {
				return fmt.Errorf("could not filter logs: %w", err)
			}

			select {
			case <-ctx.Done():
				return fmt.Errorf("context canceled while adding log to chan %w", ctx.Err())
			case f.fetchedLogsChan <- logs:
			}
		}
	}
}

// FetchLogs safely calls FilterLogs with the filtering implementing a backoff in the case of
// rate limiting and respects context cancellation.
//
// nolint:cyclop
func (f *LogFetcher) FetchLogs(ctx context.Context, chunks []*util.Chunk) ([]types.Log, error) {
	backoffConfig := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    10 * time.Second,
	}

	attempt := 0
	timeout := time.Duration(0)

	startHeight := chunks[0].StartBlock.Uint64()
	endHeight := chunks[len(chunks)-1].EndBlock.Uint64()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context was canceled before logs could be filtered")
		case <-time.After(timeout):
			attempt++
			if attempt > retryTolerance {
				return nil, fmt.Errorf("maximum number of filter attempts exceeded")
			}

			logs, err := f.getAndUnpackLogs(ctx, chunks, backoffConfig, startHeight, endHeight)
			if err != nil {
				LogEvent(WarnLevel, "Could not get and unpack logs for range, retrying", LogData{"sh": startHeight, "ca": f.contractAddress, "eh": endHeight, "cid": f.chainConfig.ChainID, "e": err})
				continue
			}

			return logs, nil
		}
	}
}

func (f *LogFetcher) getAndUnpackLogs(ctx context.Context, chunks []*util.Chunk, backoffConfig *backoff.Backoff, startHeight, endHeight uint64) ([]types.Log, error) {
	result, err := GetLogsInRange(ctx, f.backend, f.contractAddress, uint64(f.chainConfig.ChainID), chunks)
	if err != nil {
		backoffConfig.Duration()
		LogEvent(WarnLevel, "Could not filter logs for range, retrying", LogData{"sh": startHeight, "ca": f.contractAddress, "eh": endHeight, "cid": f.chainConfig.ChainID, "e": err})
		return nil, err
	}
	var logs []types.Log
	resultIterator := result.Iterator()
	for !resultIterator.Done() {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context canceled while unpacking logs from request: %w", ctx.Err())
		default:
			_, logChunk := resultIterator.Next()
			if logChunk == nil || len(*logChunk) == 0 {
				LogEvent(WarnLevel, "empty subchunk", LogData{"sh": startHeight, "ca": f.contractAddress, "cid": f.chainConfig.ChainID, "eh": endHeight})
				continue
			}

			logs = append(logs, *logChunk...)
		}
	}

	return logs, nil
}
