package executor

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3"
	"math/big"
)

// ExecutorBackend is the backend for the executor.
//
//nolint:golint,revive
type ExecutorBackend interface {
	// BlockNumber gets the latest block number.
	BlockNumber(ctx context.Context) (uint64, error)
	// BlockByNumber retrieves a block from the database by number, caching it
	// (associated with its hash) if found.
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}

// DialBackend returns an executor backend.
func DialBackend(ctx context.Context, url string) (ExecutorBackend, error) {
	c, err := rpc.DialContext(ctx, url)
	if err != nil {
		// nolint: wrapcheck
		return nil, err
	}

	ethClient := ethclient.NewClient(c)
	w3Client := w3.NewClient(c)

	return &executorBackendImpl{
		Client: ethClient,
		w3:     w3Client,
	}, nil
}

type executorBackendImpl struct {
	*ethclient.Client
	w3 *w3.Client
}
