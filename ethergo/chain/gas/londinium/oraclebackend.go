package londinium

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/backend"
)

// wrappedOracleBackend is an oracle backend that wraps gasprice.OracleBackend.
type wrappedOracleBackend struct {
	chain backend.OracleBackendChain
}

// NewOracleBackendFromChain there are conflicting method names between evm.Chain and gasprice.OracleBackend, this
// wraps chain to create a gasprice.OracleBackend.
func NewOracleBackendFromChain(unwrapped backend.OracleBackendChain) OracleBackend {
	return wrappedOracleBackend{chain: unwrapped}
}

// HeaderByNumber wraps oracle backend.
//
//nolint:wrapcheck
func (w wrappedOracleBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	return w.chain.HeaderByNumber(ctx, big.NewInt(int64(number)))
}

// BlockByNumber wraps the oracle backend.
//
//nolint:wrapcheck
func (w wrappedOracleBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	return w.chain.BlockByNumber(ctx, big.NewInt(int64(number)))
}

// ChainConfig gets the chain config for the current chain.
func (w wrappedOracleBackend) ChainConfig() *params.ChainConfig {
	return w.chain.ChainConfig()
}

var _ OracleBackend = wrappedOracleBackend{}
