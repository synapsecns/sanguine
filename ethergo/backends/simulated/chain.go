package simulated

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated/multibackend"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
	"math/big"
)

// Client is a simulated client for a simulated backend.
type Client struct {
	*multibackend.SimulatedBackend
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

// CallContext calls the call context method on the underlying client.
func (s Client) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	panic("CallContext is not supported on the simulated backend")
}

// BatchCallContext calls the batch call method on the underlying client.
func (s Client) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
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
