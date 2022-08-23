package preset

import (
	"context"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"math/big"
	"testing"
)

// Backend is a backend with preset parameters. It can be created either with the geth
// testnet or ganache backends.
type Backend struct {
	config     *params.ChainConfig
	rpcURL     string
	name       string
	privateKey string
}

// GetBigChainID gets the chain id for a preset backend.
func (b Backend) GetBigChainID() *big.Int {
	return b.config.ChainID
}

// GetChainID gets the preset chain id as a uint.
func (b Backend) GetChainID() uint {
	return uint(b.config.ChainID.Int64())
}

// Geth creates a new geth version of the preset backend.
func (b Backend) Geth(ctx context.Context, t *testing.T) *geth.Backend {
	t.Helper()
	return geth.NewEmbeddedBackendWithConfig(ctx, t, b.config)
}
