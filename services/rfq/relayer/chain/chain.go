// Package chain defines the interface for interacting with a blockchain.
package chain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

// Chain is a chain helper for relayer.
// lowercase fields are private, uppercase are public.
// the plan is to move this out of relayer which is when this distinction will matter.
type Chain struct {
	ChainID       uint32
	Bridge        *fastbridge.FastBridgeRef
	Client        client.EVM
	Confirmations uint64
	listener      listener.ContractListener
	submitter     submitter.TransactionSubmitter
}

// NewChain creates a new chain.
func NewChain(ctx context.Context, chainClient client.EVM, addr common.Address, chainListener listener.ContractListener, ts submitter.TransactionSubmitter) (*Chain, error) {
	bridge, err := fastbridge.NewFastBridgeRef(addr, chainClient)
	if err != nil {
		return nil, fmt.Errorf("could not create bridge contract: %w", err)
	}
	chainID, err := chainClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get chain id: %w", err)
	}
	return &Chain{
		ChainID: uint32(chainID.Int64()),
		Bridge:  bridge,
		Client:  chainClient,
		// TODO: configure
		Confirmations: 1,
		listener:      chainListener,
		submitter:     ts,
	}, nil
}

// SubmitTransaction submits a transaction to the chain.
func (c Chain) SubmitTransaction(ctx context.Context, call submitter.ContractCallType) (nonce uint64, _ error) {
	//nolint: wrapcheck
	return c.submitter.SubmitTransaction(ctx, big.NewInt(int64(c.ChainID)), call)
}

// LatestBlock returns the latest block.
func (c Chain) LatestBlock() uint64 {
	return c.listener.LatestBlock()
}

// SubmitRelay submits a relay transaction to the destination chain after evaluating gas amount.
func (c Chain) SubmitRelay(ctx context.Context, request reldb.QuoteRequest) (uint64, *big.Int, error) {
	gasAmount := big.NewInt(0)
	var err error

	if request.Transaction.SendChainGas {
		gasAmount, err = c.Bridge.ChainGasAmount(&bind.CallOpts{Context: ctx})
		if err != nil {
			return 0, nil, fmt.Errorf("could not get chain gas amount: %w", err)
		}
	}

	nonce, err := c.SubmitTransaction(ctx, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		transactor.Value = core.CopyBigInt(gasAmount)

		tx, err = c.Bridge.Relay(transactor, request.RawRequest)
		if err != nil {
			return nil, fmt.Errorf("could not relay: %w", err)
		}

		return tx, nil
	})
	if err != nil {
		return 0, nil, fmt.Errorf("could not submit transaction: %w", err)
	}

	return nonce, gasAmount, nil
}
