package backend

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

// OracleBackendChain is the backend for an oracle. It wraps this backend so it can be used by OracleBackend.
//
//go:generate go run github.com/vektra/mockery/v2 --name OracleBackendChain --output ./mocks --case=underscore
type OracleBackendChain interface {
	// ChainConfig gets the chain config
	ChainConfig() *params.ChainConfig
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction. This is only used for non-determinstic pricing
	SuggestGasPrice(ctx context.Context) (*big.Int, error)

	// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
	// a timely execution of a transaction. This is only used for non-deterministic pricing.
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
}
