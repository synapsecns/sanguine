package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"golang.org/x/sync/errgroup"
	"math/big"
)

// TxQueueTransactor contains a submission queue for transactions on evm-based chains.
// The queue increments nonces and bumps gases for transactions on a single-chain.
// the transaction queue is thread safe
//
// Limitations:
//   - Once a transaction is submitted to the queue, it cannot be removed. Gas for the transaction is bumped until it is processed
//     as a result every call to the transactor is treated as unique. Nonce's passed into the transactor are overridden
type TxQueueTransactor struct {
	// maxGasPrice is the max gas price to bid for a transactiosn inclusion on a particular chain
	//nolint: structcheck,unused // TODO
	maxGasPrice *big.Int
	// uses1559 indicates whether eip-1559 is enabled on this chain
	//nolint: structcheck,unused // TODO
	uses1559 bool
	// chain contains the client for interacting with the chain
	chain ChainTransactor
	// intervalSeconds is how often to bump gas/resubmit on a given chain
	//nolint: structcheck,unused  // TODO
	intervalSeconds uint32
	// signer is the signer to use for signing/submission
	signer signer.Signer
	// db is the datastore used for submitting transactions
	db db.TxQueueDB
}

// ChainTransactor is the location of the chain.
type ChainTransactor interface {
	// GetBigChainID gets the chain id as a big int.
	GetBigChainID() *big.Int
	// NonceAt gets the nonce of a chain transactor
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
}

// NewTxQueue creates a new transaction submission queue.
// TODO: we still need the transactor loop.
func NewTxQueue(signer signer.Signer, db db.TxQueueDB, chain ChainTransactor) *TxQueueTransactor {
	return &TxQueueTransactor{
		db:     db,
		chain:  chain,
		signer: signer,
	}
}

// GetTransactor gets the transactor used for transacting.
func (t *TxQueueTransactor) GetTransactor(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	signerFn := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		err := t.db.StoreRawTx(ctx, tx, chainID, address)
		if err != nil {
			return nil, fmt.Errorf("could not store tx: %w", err)
		}

		parentTransactor, err := t.signer.GetTransactor(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get transactor: %w", err)
		}

		signedTx, err := parentTransactor.Signer(address, tx)
		if err != nil {
			return nil, fmt.Errorf("could not get signed tx: %w", err)
		}

		err = t.db.StoreProcessedTx(ctx, signedTx)
		if err != nil {
			return nil, fmt.Errorf("could not get signed tx: %w", err)
		}

		return signedTx, nil
	}

	return &bind.TransactOpts{
		From:   t.signer.Address(),
		Signer: signerFn,
	}, nil
}

// GetNonce uses the greatest of the database nonce or the on-chain nonce for the next transaction.
func (t TxQueueTransactor) GetNonce(parentCtx context.Context) (nonce uint64, err error) {
	g, ctx := errgroup.WithContext(parentCtx)
	// onChainNonce is the latest nonce from eth_transactionCount. DB nonce is latest nonce from db + 1
	// locks are not built into this method or the insertion level of the db
	var onChainNonce, dbNonce uint64

	g.Go(func() (err error) {
		onChainNonce, err = t.chain.NonceAt(ctx, t.signer.Address(), nil)
		if err != nil {
			return fmt.Errorf("could not get on chain nonce: %w", err)
		}

		return nil
	})

	g.Go(func() (err error) {
		dbNonce, err = t.db.GetNonceForChainID(ctx, t.signer.Address(), t.chain.GetBigChainID())
		if err != nil {
			return fmt.Errorf("could not get on chain nonce: %w", err)
		}

		dbNonce++

		return nil
	})

	err = g.Wait()
	if err != nil {
		return 0, fmt.Errorf("could not get nonce: %w", err)
	}

	if onChainNonce > dbNonce {
		return onChainNonce, nil
	}

	return dbNonce, nil
}
