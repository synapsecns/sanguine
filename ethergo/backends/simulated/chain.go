package simulated

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated/multibackend"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"golang.org/x/sync/errgroup"
	"math/big"
)

// Client is a simulated client for a simulated backend.
type Client struct {
	*multibackend.SimulatedBackend
}

// FeeHistory is not implemented on this backend.
func (s Client) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	// TODO implement me
	panic("cannot implement on this backend")
}

// PendingBalanceAt calls balance at since simulated backends are monotonic.
func (s Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	//nolint: wrapcheck
	return s.SimulatedBackend.BalanceAt(ctx, account, nil)
}

// PendingStorageAt gets the storage at since simulated backends cannot have non-final storage.
func (s Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	//nolint: wrapcheck
	return s.SimulatedBackend.StorageAt(ctx, account, key, nil)
}

// PendingTransactionCount always returns 0 since simulated backends cannot have pending transactions.
func (s Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	return 0, nil
}

// SyncProgress panics since this state is not accessible on the simulated backend.
func (s Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	panic("not implemented")
}

// NetworkID wraps network id on underlying backend.
func (s Client) NetworkID(ctx context.Context) (*big.Int, error) {
	//nolint: errwrap
	return s.ChainID(ctx)
}

// ChainConfig gets the chain config for the backend.
func (s Client) ChainConfig() *params.ChainConfig {
	return s.Blockchain().Config()
}

// ChainID returns the chain id.
func (s Client) ChainID(_ context.Context) (*big.Int, error) {
	return s.ChainConfig().ChainID, nil
}

// Close closes the connection with the chain.
func (s Client) Close() {
	// do nothing
}

// CallContext panics here to bypass interface requirements for testing.
func (s Client) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	panic("CallContext is not supported on the simulated backend")
}

// BatchCallContext panics here to bypass interface requirements for testing.
func (s Client) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	panic("BatchCallContext is not supported on the simulated backend")
}

// BatchContext panics here to bypass interface requirements for testing.
func (s Client) BatchContext(ctx context.Context, calls ...w3types.Caller) error {
	panic("BatchCallContext is not supported on the simulated backend")
}

// BlockNumber gets the latest block number.
func (s Client) BlockNumber(ctx context.Context) (uint64, error) {
	latestBlock, err := s.BlockByNumber(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get block: %w", err)
	}
	return latestBlock.NumberU64(), nil
}

// SuggestGasPrice follows the rpc behavior for SuggestGasPrice. Because we rely
// on the legacy behavior of eth_sendTransaction and don't utilize the base fee (we need to figure
// out a way to do this cross-chain), we emulate the rpc eth_sugestGasPrice behavior here.
// TODO: find out if not emulating the rpc here is intended behavior on geth's end, patch via pr if not.
func (s Client) SuggestGasPrice(ctx context.Context) (gasPrice *big.Int, err error) {
	g, ctx := errgroup.WithContext(ctx)
	var (
		// baseFee is the base fee to add to the gas price estimation
		baseFee *big.Int
		// estimatedPrice is the gas price estimate
		estimatedPrice *big.Int
	)

	g.Go(func() error {
		latestBlock, err := s.BlockByNumber(ctx, nil)
		if err != nil {
			return fmt.Errorf("could not get latest block to add to base fee: %w", err)
		}
		baseFee = latestBlock.BaseFee()
		return nil
	})

	g.Go(func() error {
		estimatedPrice, err = s.SimulatedBackend.SuggestGasPrice(ctx)
		if err != nil {
			return fmt.Errorf("could not get suggested gas price")
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get gas price: %w", err)
	}

	if baseFee == nil {
		return estimatedPrice, nil
	}

	return big.NewInt(0).Add(estimatedPrice, baseFee), nil
}

var _ client.EVMClient = &Client{}
