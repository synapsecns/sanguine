package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/synapse-node/pkg/evm"
	"math/big"
)

// TxQueueTransactor contains a submission queue for transactions on evm-based chains.
// The queue increments nonces and bumps gases for transactions on a single-chain.
// the transaction queue is thread safe
//
// Limitations:
// - Once a transaction is submitted to the queue, it cannot be removed. Gas for the transaction is bumped until it is processed
//   as a result every call to the transactor is treated as unique. Nonce's passed into the transactor are overridden
type TxQueueTransactor struct {
	// maxGasPrice is the max gas price to bid for a transactiosn inclusion on a particular chain
	maxGasPrice *big.Int
	// uses1559 indicates whether eip-1559 is enabled on this chain
	uses1559 bool
	// chain contains the client for interacting with the chain
	chain evm.Chain
	// intervalSeconds is how often to bump gas/resubmit on a given chain
	intervalSeconds uint32
	// signer is the signer to use for signing/submission
	signer signer.Signer
}

// NewTxQueue creates a new transaction submission queue.
func NewTxQueue(ctx context.Context) *TxQueueTransactor {
	return &TxQueueTransactor{}
}

// GetTransactor gets the transactor used for.
func (t *TxQueueTransactor) GetTransactor() (*bind.TransactOpts, error) {
	latestSigner := types.LatestSignerForChainID(big.NewInt(1))

	_ = latestSigner

	// TODO
	signerFn := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		panic("")
	}

	return &bind.TransactOpts{
		From:   t.signer.Address(),
		Signer: signerFn,
	}, nil
}
