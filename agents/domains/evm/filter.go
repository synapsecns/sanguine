package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"
	"time"
)

// LogInfo is the log info.
type LogInfo struct {
	// logs are logs
	logs []types.Log
	// chunk are chunks
	chunk *util.Chunk
}

// RangeFilter pre-fetches filter logs into a channel in deterministic order.
type RangeFilter struct {
	// iterator is the chunk iterator used for the range
	iterator util.ChunkIterator
	// logs is a channel with the filtered ahead logs. This channel is not closed
	// and the user can rely on the garbage collection behavior of RangeFilter to remove it.
	logs chan *LogInfo
	// filterer contains the interface used to fetch logs for the given contract. Logs are fteched
	// for contractAddress
	filterer bind.ContractFilterer
	// contractAddress is the contractAddress that logs are fetched for
	contractAddress ethCommon.Address
	// done is whether or not the RangeFilter has completed. It cannot be restarted and the object must be recreated
	done bool
}

// bufferSize is how many ranges ahead should be fetched.
const bufferSize = 10

// maxAttempts is that maximum number of times a filter attempt should be made before giving up.
const maxAttempts = 5

// minBackoff is the minimum backoff period between requests.
var minBackoff = 1 * time.Second

// maxBackoff is the maximum backoff period between requests.
var maxBackoff = 30 * time.Second

// NewRangeFilter creates a new filtering interface for a range of blocks. If reverse is not set, block heights are filtered from start->end.
func NewRangeFilter(address ethCommon.Address, filterer bind.ContractFilterer, startBlock, endBlock *big.Int, chunkSize int, reverse bool) *RangeFilter {
	return &RangeFilter{
		iterator:        util.NewChunkIterator(startBlock, endBlock, chunkSize, reverse),
		logs:            make(chan *LogInfo, bufferSize),
		filterer:        filterer,
		contractAddress: address,
		done:            false,
	}
}

// Start starts the filtering process. If the context is canceled, logs will stop being filtered. Returns an error.
// this should be run on an independent goroutine.
func (f *RangeFilter) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			if !f.done && ctx.Err() != nil {
				return fmt.Errorf("could not finish filtering range: %w", ctx.Err())
			}
			return nil
		default:
			chunk := f.iterator.NextChunk()
			if chunk == nil {
				f.done = true
				return nil
			}

			logs, err := f.FilterLogs(ctx, chunk)

			if err != nil {
				return fmt.Errorf("could not filter logs: %w", err)
			}

			f.appendToChannel(ctx, logs)
		}
	}
}

// FilterLogs safely calls FilterLogs with the filtering implementing a backoff in the case of
// rate limiting and respecting context cancellation.
func (f *RangeFilter) FilterLogs(ctx context.Context, chunk *util.Chunk) (*LogInfo, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    minBackoff,
		Max:    maxBackoff,
	}

	attempt := 0
	// timeout should always be 0 on the first attmept
	timeout := time.Duration(0)
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("could not finish filtering logs: %w", ctx.Err())
		case <-time.After(timeout):
			attempt++
			if attempt > maxAttempts {
				return nil, errors.New("maximum number of filter attempts exceeded")
			}

			logs, err := f.filterer.FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: chunk.MinBlock(),
				ToBlock:   chunk.MaxBlock(),
				Addresses: []ethCommon.Address{f.contractAddress},
			})

			if err != nil {
				timeout = b.Duration()
				logger.Warnf("could not filter logs for range %d to %d: %v", chunk.MinBlock(), chunk.MaxBlock(), err)
				continue
			}

			return &LogInfo{
				logs:  logs,
				chunk: chunk,
			}, nil
		}
	}
}

// Drain fetches all logs and concatenated them into a single slice.
// Deprecated: use the channel.
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
		}
	}
}

// appendToChannel is a helper method that appends logs to a channel while respecting context cancellations.
func (f *RangeFilter) appendToChannel(ctx context.Context, logs *LogInfo) {
	select {
	case <-ctx.Done():
		return
	case f.logs <- logs:
	}
}

// Done returns a bool indicating whether or not the filtering operation is done.
func (f *RangeFilter) Done() bool {
	return f.done
}

// GetLogChan retursn a log chan with the logs filtered ahead to bufferSize. Iteration oder is only guaranteed with up to one
// consumer.
func (f *RangeFilter) GetLogChan() chan *LogInfo {
	return f.logs
}
