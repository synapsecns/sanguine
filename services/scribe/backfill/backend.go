package backfill

import (
	"context"
	"fmt"
	"github.com/benbjohnson/immutable"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/ethergo/util"
	"golang.org/x/exp/constraints"
	"math"
	"math/big"
	"time"
)

// ScribeBackend is the set of functions that the scribe needs from a client.
type ScribeBackend interface {
	// ChainID gets the chain id from the rpc server.
	ChainID(ctx context.Context) (*big.Int, error)
	// BlockByNumber retrieves a block from the database by number, caching it
	// (associated with its hash) if found.
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	// TransactionByHash checks the pool of pending transactions in addition to the
	// blockchain. The isPending return value indicates whether the transaction has been
	// mined yet. Note that the transaction may not be part of the canonical chain even if
	// it's not pending.
	TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error)
	// TransactionReceipt returns the receipt of a mined transaction. Note that the
	// transaction may not be included in the current canonical chain even if a receipt
	// exists.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	// BlockNumber gets the latest block number.
	BlockNumber(ctx context.Context) (uint64, error)
	// FilterLogs executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	//
	// TODO(karalabe): Deprecate when the subscription one can return past data too.
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	// HeaderByNumber returns the block header with the given block number.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// Batch batches multiple
	Batch(ctx context.Context, calls ...w3types.Caller) error
}

// DialBackend returns a scribe backend.
func DialBackend(ctx context.Context, url string) (ScribeBackend, error) {
	c, err := rpc.DialContext(ctx, url)
	if err != nil {
		// nolint: wrapcheck
		return nil, err
	}

	ethClient := ethclient.NewClient(c)
	w3Client := w3.NewClient(c)

	return &scribeBackendImpl{
		Client: ethClient,
		w3:     w3Client,
	}, nil
}

type scribeBackendImpl struct {
	*ethclient.Client
	w3 *w3.Client
}

// Batch batches multiple w3 calls.
func (c *scribeBackendImpl) Batch(ctx context.Context, calls ...w3types.Caller) error {
	//nolint: wrapcheck
	return c.w3.CallCtx(ctx, calls...)
}

var _ ScribeBackend = &scribeBackendImpl{}

// BlockTimesInRange gets all block times in a range with a single batch request
// in successful cases an immutable map is returned of [height->time], otherwise an error is returned.
func BlockTimesInRange(ctx context.Context, backend ScribeBackend, startHeight uint64, endHeight uint64) (*immutable.Map[uint64, uint64], error) {
	// performance impact will be negligible here because of external constraints on blocksize
	blocks := makeRange(startHeight, endHeight)
	bulkSize := len(blocks)
	calls := make([]w3types.Caller, bulkSize)
	results := make([]types.Header, bulkSize)

	for i, blockNumber := range blocks {
		calls[i] = eth.HeaderByNumber(new(big.Int).SetUint64(blockNumber)).Returns(&results[i])
	}

	if err := backend.Batch(ctx, calls...); err != nil {
		return nil, fmt.Errorf("could not fetch blocks in range %d to %d: %w", startHeight, endHeight, err)
	}

	// use an immutable map for additional safety to the caller, don't allocate until batch returns successfully
	res := immutable.NewMapBuilder[uint64, uint64](nil)
	for _, result := range results {
		res.Set(result.Number.Uint64(), result.Time)
	}

	return res.Map(), nil
}

// GetLogsInRange gets all logs in a range with a single batch request
// in successful cases an immutable list is returned, otherwise an error is returned.
func GetLogsInRange(ctx context.Context, backend ScribeBackend, startHeight uint64, endHeight uint64, subChunkSize uint64, contractAddress common.Address) (*immutable.List[*[]types.Log], error) {
	blockRange := (endHeight - startHeight) + 1
	subChunkCount := int(math.Ceil(float64(blockRange) / float64(subChunkSize)))
	iterator := util.NewChunkIterator(big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), int(subChunkSize)-1, true)
	calls := make([]w3types.Caller, subChunkCount)
	results := make([][]types.Log, subChunkCount)
	subChunkIdx := uint64(0)
	chunk := iterator.NextChunk()
	for chunk != nil {
		filter := ethereum.FilterQuery{
			FromBlock: chunk.StartBlock,
			ToBlock:   chunk.EndBlock,
			Addresses: []common.Address{contractAddress},
		}
		calls[subChunkIdx] = eth.Logs(filter).Returns(&results[subChunkIdx])
		subChunkIdx++
		chunk = iterator.NextChunk()
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Minute*5)
	defer cancel()

	if err := backend.Batch(timeoutCtx, calls...); err != nil {
		return nil, fmt.Errorf("could not fetch logs in range %d to %d: %w", startHeight, endHeight, err)
	}
	// use an immutable list for additional safety to the caller, don't allocate until batch returns successfully
	res := immutable.NewListBuilder[*[]types.Log]()
	for _, result := range results {
		logChunk := result
		if len(result) > 0 {
			res.Append(&logChunk)
		}
	}
	return res.List(), nil
}

// make range.
func makeRange[T constraints.Integer](min, max T) []T {
	a := make([]T, max-min+1)
	for i := range a {
		a[i] = min + T(i)
	}
	return a
}
