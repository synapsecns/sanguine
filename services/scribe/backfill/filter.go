package backfill

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/ethergo/util"

	"github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
)

// LogInfo is the log info.
type LogInfo struct {
	logs  []types.Log
	chunk *util.Chunk
}

// RangeFilter pre-fetches filter logs into a channel in deterministic order.
type RangeFilter struct {
	// iterator is the chunk iterator used for the range.
	iterator util.ChunkIterator
	// logs is a channel with the filtered ahead logs. This channel is not closed
	// and the user can rely on the garbage collection behavior of RangeFilter to remove it.
	logs chan *LogInfo
	// backend is the ethereum backend used to fetch logs.
	backend ScribeBackend
	// contractAddress is the contractAddress that logs are fetched for.
	contractAddress ethCommon.Address
	// done is whether the RangeFilter has completed. It cannot be restarted and the object must be recreated.
	done bool
	// doneChan is a channel that is closed when the RangeFilter has completed.
	// this is only to be used by external callers
	doneChan chan bool
	// subChunkSize is the size of each batch.
	subChunkSize int
	// chainID is the chain ID.
	chainID uint32
}

// LogFilterer is the interface for filtering logs.
type LogFilterer interface {
	// FilterLogs executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	//
	// TODO(karalabe): Deprecate when the subscription one can return past data too.
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
}

// bufferSize is how many ranges ahead should be fetched.
const bufferSize = 15

// minBackoff is the minimum backoff period between requests.
var minBackoff = 1 * time.Second

// maxBackoff is the maximum backoff period between requests.
var maxBackoff = 10 * time.Second

// NewRangeFilter creates a new filtering interface for a range of blocks. If reverse is not set, block heights are filtered from start->end.
func NewRangeFilter(address ethCommon.Address, backend ScribeBackend, startBlock, endBlock *big.Int, chunkSize int, reverse bool, subChunkSize int, chainID uint32) *RangeFilter {
	return &RangeFilter{
		iterator:        util.NewChunkIterator(startBlock, endBlock, chunkSize, reverse),
		logs:            make(chan *LogInfo, bufferSize),
		backend:         backend,
		contractAddress: address,
		doneChan:        make(chan bool),
		done:            false,
		subChunkSize:    subChunkSize,
		chainID:         chainID,
	}
}

// Start starts the filtering process. If the context is canceled, logs will stop being filtered.
// This should be run on an independent goroutine.
func (f *RangeFilter) Start(ctx context.Context) error {
	defer close(f.doneChan)
	for {
		select {
		case <-ctx.Done():
			LogEvent(InfoLevel, "Contract backfill context completed", LogData{"ca": f.contractAddress})

			if !f.done && ctx.Err() != nil {
				LogEvent(InfoLevel, "could not finish filtering range", LogData{"ca": f.contractAddress})

				return fmt.Errorf("could not finish filtering range: %w", ctx.Err())
			}

			return nil
		default:
			startTime := time.Now()
			chunk := f.iterator.NextChunk()

			if chunk == nil {
				f.done = true
				f.doneChan <- true

				return nil
			}

			logs, err := f.FilterLogs(ctx, chunk)

			if err != nil {
				return fmt.Errorf("could not filter logs: %w", err)
			}

			f.appendToChannel(ctx, logs)
			LogEvent(InfoLevel, "Contract backfill chunk completed", LogData{"ca": f.contractAddress, "sh": chunk.MinBlock(), "eh": chunk.MaxBlock(), "ts": time.Since(startTime).Seconds()})
		}
	}
}

// FilterLogs safely calls FilterLogs with the filtering implementing a backoff in the case of
// rate limiting and respecting context cancellation.
//
// nolint:cyclop
func (f *RangeFilter) FilterLogs(ctx context.Context, chunk *util.Chunk) (*LogInfo, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    minBackoff,
		Max:    maxBackoff,
	}

	attempt := 0
	timeout := time.Duration(0)

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("could not finish filtering logs: %w", ctx.Err())
		case <-time.After(timeout):
			attempt++

			if attempt > retryTolerance {
				return nil, fmt.Errorf("maximum number of filter attempts exceeded")
			}

			res, err := GetLogsInRange(ctx, f.backend, chunk.MinBlock().Uint64(), chunk.MaxBlock().Uint64(), uint64(f.subChunkSize), f.contractAddress, uint64(f.chainID))
			if err != nil {
				timeout = b.Duration()
				LogEvent(WarnLevel, "Could not filter logs for range, retrying", LogData{"sh": chunk.MinBlock(), "ca": f.contractAddress, "eh": chunk.MaxBlock(), "e": err})

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
						LogEvent(WarnLevel, "empty subchunk", LogData{"sh": chunk.MinBlock(), "ca": f.contractAddress, "eh": chunk.MaxBlock()})
						continue
					}
					logsChunk := *resLogChunk

					logs = append(logs, logsChunk...)
				}
			}

			return &LogInfo{
				logs:  logs,
				chunk: chunk,
			}, nil
		}
	}
}

// Drain fetches all logs and concatenated them into a single slice.
func (f *RangeFilter) Drain(ctx context.Context) (filteredLogs []types.Log, err error) {
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context ended: %w", ctx.Err())
		case log := <-f.GetLogChan():
			filteredLogs = append(filteredLogs, log.logs...)

			if f.done {
				return filteredLogs, nil
			}
		default:
			return filteredLogs, nil
		}
	}
}

// appendToChannel is a helper method that appends logs to a channel while respecting context cancellations.
func (f *RangeFilter) appendToChannel(ctx context.Context, logs *LogInfo) {
	select {
	case <-ctx.Done():
		return
	case f.logs <- logs:
		for _, log := range logs.logs {
			LogEvent(ErrorLevel, "appended log to channel", LogData{"ca": f.contractAddress, "tx": log.TxHash, "cid": f.chainID})
		}
	}
}

// Done returns a bool indicating whether the filtering operation is done.
func (f *RangeFilter) Done() bool {
	return f.done
}

// GetLogChan returns a log chan with the logs filtered ahead to bufferSize. Iteration oder is only guaranteed with up to one
// consumer.
func (f *RangeFilter) GetLogChan() chan *LogInfo {
	return f.logs
}
