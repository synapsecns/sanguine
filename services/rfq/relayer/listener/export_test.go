package listener

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/db"
)

// TestChainListener wraps chain listener for testing
type TestChainListener interface {
	ChainListener
	GetMetadata(parentCtx context.Context) (startBlock, chainID uint64, err error)
}

// GetMetadata wraps chain listener for testing
func (c chainListener) GetMetadata(ctx context.Context) (startBlock, chainID uint64, err error) {
	return c.getMetadata(ctx)
}

type TestChainListenerArgs struct {
	Address  common.Address
	Client   client.EVM
	Contract *fastbridge.FastBridgeRef
	Store    db.Service
	Handler  metrics.Handler
}

func NewTestChainListener(args TestChainListenerArgs) TestChainListener {
	return chainListener{
		address:  args.Address,
		client:   args.Client,
		contract: args.Contract,
		store:    args.Store,
		handler:  args.Handler,
	}
}
