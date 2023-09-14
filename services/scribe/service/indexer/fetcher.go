package indexer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/logger"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/ethergo/util"

	"github.com/ethereum/go-ethereum/core/types"

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
	fetchedLogsChan chan types.Log
	// backend is the ethereum backend used to fetch logs.
	backend backend.ScribeBackend
	// indexerConfig holds the chain config (config data for the chain)
	indexerConfig *scribeTypes.IndexerConfig
	// topics is the list of topics to filter logs by.
	topics [][]common.Hash
	// bufferSize prevents from overloading the scribe indexer with too many logs as well as upstream RPCs with too many requests.
	bufferSize int
}

// NewLogFetcher creates a new filtering interface for a range of blocks. If reverse is not set, block heights are filtered from start->end.
func NewLogFetcher(backend backend.ScribeBackend, startBlock, endBlock *big.Int, indexerConfig *scribeTypes.IndexerConfig, ascending bool) *LogFetcher {
	// The ChunkIterator is inclusive of the start and ending block resulting in potentially confusing behavior when
	// setting the range size in the config. For example, setting a range of 1 would result in two blocks being queried
	// instead of 1. This is accounted for by subtracting 1.
	chunkSize := int(indexerConfig.GetLogsRange) - 1

	// Using the specified StoreConcurrency value from the config, as the buffer size for the fetchedLogsChan
	bufferSize := indexerConfig.StoreConcurrency
	if bufferSize > 100 {
		bufferSize = 100
	}
	if bufferSize == 0 {
		bufferSize = 3 // default buffer size
	}
	return &LogFetcher{
		iterator:        util.NewChunkIterator(startBlock, endBlock, chunkSize, ascending),
		startBlock:      startBlock,
		endBlock:        endBlock,
		fetchedLogsChan: make(chan types.Log, bufferSize),
		backend:         backend,
		indexerConfig:   indexerConfig,
		bufferSize:      bufferSize,
		topics:          indexerConfig.Topics,
	}
}

// GetChunkArr gets the appropriate amount of block chunks (getLogs ranges).
func (f *LogFetcher) GetChunkArr() (chunkArr []*util.Chunk) {
	for i := uint64(0); i < f.indexerConfig.GetLogsBatchAmount; i++ {
		chunk := f.iterator.NextChunk()
		if chunk == nil {
			return chunkArr
		}
		chunkArr = append(chunkArr, chunk)

		// Stop appending chunks if the max height of the current chunk exceeds the concurrency threshold
		if chunk.EndBlock.Uint64() > f.endBlock.Uint64()-f.indexerConfig.ConcurrencyThreshold {
			logger.ReportScribeState(f.indexerConfig.ChainID, chunk.EndBlock.Uint64(), f.indexerConfig.Addresses, logger.ConcurrencyThresholdReached)
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

			default:
				// insert logs into channel
				for i := range logs {
					f.fetchedLogsChan <- logs[i]
				}
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
		Max:    8 * time.Second,
	}

	attempt := 0
	timeout := time.Duration(0)

	for {
		select {
		case <-ctx.Done():
			logger.ReportIndexerError(ctx.Err(), *f.indexerConfig, logger.GetLogsError)
			return nil, fmt.Errorf("context was canceled before logs could be fetched")
		case <-time.After(timeout):
			attempt++
			if attempt > retryTolerance {
				logger.ReportIndexerError(fmt.Errorf("retry max reached"), *f.indexerConfig, logger.GetLogsError)
				return nil, fmt.Errorf("maximum number of fetch logs attempts exceeded")
			}

			logs, err := f.getAndUnpackLogs(ctx, chunks, backoffConfig)
			if err != nil {
				logger.ReportIndexerError(err, *f.indexerConfig, logger.GetLogsError)
				timeout = backoffConfig.Duration()
				continue
			}

			return logs, nil
		}
	}
}

func (f *LogFetcher) getAndUnpackLogs(ctx context.Context, chunks []*util.Chunk, backoffConfig *backoff.Backoff) ([]types.Log, error) {
	result, err := backend.GetLogsInRange(ctx, f.backend, f.indexerConfig.Addresses, uint64(f.indexerConfig.ChainID), chunks, f.indexerConfig.Topics)
	if err != nil {
		backoffConfig.Duration()
		return nil, fmt.Errorf("could not get logs: %w", err)
	}

	var logs []types.Log
	resultIterator := result.Iterator()
	for !resultIterator.Done() {
		select {
		case <-ctx.Done():
			logger.ReportIndexerError(ctx.Err(), *f.indexerConfig, logger.GetLogsError)
			return nil, fmt.Errorf("context canceled while unpacking logs from request: %w", ctx.Err())
		default:
			_, logChunk := resultIterator.Next()
			if logChunk == nil || len(*logChunk) == 0 {
				logger.ReportIndexerError(fmt.Errorf("empty log chunk"), *f.indexerConfig, logger.EmptyGetLogsChunk)
				continue
			}

			logs = append(logs, *logChunk...)
		}
	}

	return logs, nil
}

// GetFetchedLogsChan returns the fetchedLogsChan channel as a pointer for access by the indexer and tests.
func (f *LogFetcher) GetFetchedLogsChan() *chan types.Log {
	return &f.fetchedLogsChan
}
