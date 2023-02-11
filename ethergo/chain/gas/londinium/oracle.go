package londinium

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// HeightOracleBackend returns deterministic prices for a given gas amount.
type HeightOracleBackend struct {
	OracleBackend
	height uint
}

// HeaderByNumber overrides the default behavior on header by number when getting the latest block.
//
//nolint:wrapcheck
func (d HeightOracleBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	if number == rpc.LatestBlockNumber {
		return d.OracleBackend.HeaderByNumber(ctx, rpc.BlockNumber(d.height))
	}
	return d.OracleBackend.HeaderByNumber(ctx, number)
}

// NewOracleBackendFromHeight creates a gas price oracleBackend for deterministically generating gas prices from a given height
// this hijacks the HeaderByNumber method (when it's passed -1) and returns the block height passed in the constructor.
func NewOracleBackendFromHeight(height uint64, oracle OracleBackend) HeightOracleBackend {
	return HeightOracleBackend{
		OracleBackend: oracle,
		height:        uint(height),
	}
}
