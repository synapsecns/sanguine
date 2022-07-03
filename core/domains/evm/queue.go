package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/synapse-node/pkg/evm"
	"golang.org/x/sync/errgroup"
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
	// db is the datastore used for submitting transactions
	db db.TxQueueDB
}

// NewTxQueue creates a new transaction submission queue.
func NewTxQueue(ctx context.Context) *TxQueueTransactor {
	return &TxQueueTransactor{}
}

// GetTransactor gets the transactor used for.
func (t *TxQueueTransactor) GetTransactor(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	latestSigner := types.LatestSignerForChainID(chainID)

	_ = latestSigner

	// TODO
	signerFn := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		err := t.db.StoreRawTx(ctx, tx, chainID, address)
		if err != nil {
			return nil, fmt.Errorf("could not store tx: %w", err)
		}

		return nil, nil
	}

	return &bind.TransactOpts{
		From:   t.signer.Address(),
		Signer: signerFn,
	}, nil
}

// getNonce uses the greatest of the database nonce or the on-chain nonce for the next transaction.
func (t TxQueueTransactor) getNonce(parentCtx context.Context) {
	g, ctx := errgroup.WithContext(parentCtx)
	var onChainNonce, dbNonce uint64

	g.Go(func() (err error) {
		onChainNonce, err = t.chain.NonceAt(ctx, t.signer.Address(), nil)
		if err != nil {
			return fmt.Errorf("could not get on chain nonce: %w", err)
		}

		return nil
	})

	g.Go(func() (err error) {
		// dbNonce, err = t.db.StoreRawTx()
		if err != nil {
			return fmt.Errorf("could not get on chain nonce: %w", err)
		}

		return nil
	})

	_ = onChainNonce
	_ = dbNonce
}
