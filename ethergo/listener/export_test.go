package listener

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/listener/db"
)

// TestChainListener wraps chain listener for testing.
type TestChainListener interface {
	ContractListener
	GetMetadata(parentCtx context.Context) (startBlock, chainID uint64, err error)
}

// GetMetadata wraps chain listener for testing.
func (c chainListener) GetMetadata(ctx context.Context) (startBlock, chainID uint64, err error) {
	return c.getMetadata(ctx)
}

type TestChainListenerArgs struct {
	Address      common.Address
	InitialBlock uint64
	Client       client.EVM
	Store        db.ChainListenerDB
	Handler      metrics.Handler
}

func NewTestChainListener(args TestChainListenerArgs) TestChainListener {
	return &chainListener{
		client:       args.Client,
		address:      args.Address,
		initialBlock: args.InitialBlock,
		store:        args.Store,
		handler:      args.Handler,
	}
}
