package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/chain/gas"
	"math/big"
	"time"
)

// Chain is a chain.
//
// Deprecated: will be removed in a future version
//
//go:generate go run github.com/vektra/mockery/v2 --name Chain --output ./mocks --case=underscore
type Chain interface {
	client.MeteredEVMClient
	// GetBigChainID gets the chain id as a big int.
	GetBigChainID() *big.Int
	// GetChainID gets chain id
	GetChainID() uint
	// RPCAddress gets the rpc address of the chain (if available)
	RPCAddress() string
	// GetHeightWatcher gets the block height watcher for the chain
	GetHeightWatcher() chainwatcher.BlockHeightWatcher
	// ChainID retrieves the current chain ID for transaction replay protection.
	ChainID(ctx context.Context) (*big.Int, error)
	// PendingNonceAt returns the account nonce of the given account in the pending state.
	// This is the nonce that should be used for the next transaction.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// ChainName gets the chain name
	ChainName() string
	// Estimator returns the gas price estimator
	// Deprecated: use gas setter
	Estimator() gas.PriceEstimator
	// GasSetter gets the gas setter
	GasSetter() gas.Setter
	// ChainConfig gets the chain config.
	ChainConfig() *params.ChainConfig
	// SetChainConfig sets the config for a chain
	SetChainConfig(config *params.ChainConfig)
	// HeaderByTime gets the closest block to the given time.
	HeaderByTime(ctx context.Context, startBlock *big.Int, searchTime time.Time) (*ethTypes.Header, error)
}
