package client

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core/metrics"
	"math/big"
)

// Client is the set of functions that the scribe needs from a client.
type Client interface {
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

type clientImpl struct {
	*ethclient.Client
	w3 *w3.Client
}

// DialBackend returns a scribe backend.
func DialBackend(ctx context.Context, url string, handler metrics.Handler) (Client, error) {
	c, err := metrics.RPCClient(ctx, handler, url)
	if err != nil {
		return nil, fmt.Errorf("failed to create rpc client: %w", err)
	}

	ethClient := ethclient.NewClient(c)
	w3Client := w3.NewClient(c)

	return &clientImpl{
		Client: ethClient,
		w3:     w3Client,
	}, nil
}

// Batch batches multiple w3 calls.
func (c *clientImpl) Batch(ctx context.Context, calls ...w3types.Caller) error {
	//nolint: wrapcheck
	return c.w3.CallCtx(ctx, calls...)
}
