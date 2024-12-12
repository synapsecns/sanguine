package multibackend

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/w3types"
)

// SimulatedBackend wraps the simulated.Backend to provide additional functionality
type SimulatedBackend struct {
	*simulated.Backend
	client simulated.Client
}

// EmptyBlock mines an empty block at time. Must be greater than previous block time.
func (b *SimulatedBackend) EmptyBlock(blockTime time.Time) {
	// Get the current block time
	header, err := b.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic("could not fetch current header")
	}

	// Calculate the time difference and adjust
	timeDiff := blockTime.Unix() - int64(header.Time)
	if timeDiff <= 0 {
		panic("block time must be greater than previous block time")
	}

	// Adjust the time and commit a new block
	b.Backend.AdjustTime(time.Duration(timeDiff) * time.Second)
	b.Backend.Commit()
}

// HeaderByNumber returns the block header with the given block number.
func (b *SimulatedBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return b.client.HeaderByNumber(ctx, number)
}

// TransactionByHash returns the transaction with the given hash.
func (b *SimulatedBackend) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return b.client.TransactionByHash(ctx, hash)
}

// SendTransaction injects a signed transaction into the pending pool for execution.
func (b *SimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return b.client.SendTransaction(ctx, tx)
}

// CallContext executes a single RPC call.
func (b *SimulatedBackend) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return b.Backend.CallContext(ctx, result, method, args...)
}

// BatchContext executes a batch of RPC calls in a single request.
func (b *SimulatedBackend) BatchContext(ctx context.Context, calls ...w3types.RPCCaller) error {
	// Simulated backend doesn't support batch calls, execute them sequentially
	for _, call := range calls {
		elem, err := call.CreateRequest()
		if err != nil {
			return err
		}
		// Use the Backend's CallContext since Client doesn't have it
		if err := b.Backend.CallContext(ctx, &elem.Result, elem.Method, elem.Args...); err != nil {
			return err
		}
		if err := call.HandleResponse(elem); err != nil {
			return err
		}
	}
	return nil
}

// NewSimulatedBackend creates a new simulated backend
func NewSimulatedBackend(alloc core.GenesisAlloc, gasLimit uint64) *SimulatedBackend {
	backend := simulated.NewBackend(alloc)
	return &SimulatedBackend{
		Backend: backend,
		client:  backend.Client(),
	}
}
