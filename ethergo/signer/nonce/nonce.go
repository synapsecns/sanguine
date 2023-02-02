package nonce

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/mapmutex"
)

// Manager is a singleton used to generate a nonce.
// This can be explicitly overrided by setting a nonce manually (not recommended). The issue here is many of the
// deploy helper functions are called asynchronously and there's a gap between the tx.Sequence generation and the
// transactor. This solves that by wrapping the transactor in a nonce manager.
type Manager interface {
	// SignTx signs a legacy tx
	SignTx(ogTx *types.Transaction, signer types.Signer, prv *ecdsa.PrivateKey) (*types.Transaction, error)
	// NewKeyedTransactor wraps keyed transactor in a nonce manager.
	// right now, this only works if all txes are sent out (a safe assumption in test mode)
	// this can be obviated by signing at send time or loop + retrying on failure
	NewKeyedTransactor(realSigner *bind.TransactOpts) (*bind.TransactOpts, error)
	// NewKeyedTransactorFromKey creates a new keyed transactor from a private key.
	NewKeyedTransactorFromKey(key *ecdsa.PrivateKey) (*bind.TransactOpts, error)
	// GetNextNonce gets the next nonce for the address.
	GetNextNonce(address common.Address) (*big.Int, error)
}

// ChainQuery is a chain used to generate a nonce.
type ChainQuery interface {
	// PendingNonceAt returns the account nonce of the given account in the pending state.
	// This is the nonce that should be used for the next transaction.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
}

type nonceManagerImp struct {
	// chain is the chain to use for the nonce
	chain ChainQuery
	// ctx is the context used while fetching data
	//nolint: containedctx
	ctx context.Context
	// chainID is the chain id of the nonce manager.
	chainID *big.Int
	// nonceMapLock locks the nonce map for reads/writes. We use this instead of a sync
	// map because of type safety
	nonceMapLock sync.RWMutex
	// nonceMap is a map of accounts->next nonce
	nonceMap map[common.Address]*big.Int
	// accountMutex is responsible for making sure only one nonce per account is generated at a time
	// we don't use the global nonce lock here since we want to allow others to sign transactions during
	// signing. We can't paralellize this per account because the transactor may need to subtract if signing fails
	accountMutex mapmutex.StringerMapMutex
}

// NewNonceManager generates a new nonce manager. This should be called once, as nonce storage is not global.
func NewNonceManager(ctx context.Context, chain ChainQuery, chainID *big.Int) Manager {
	return &nonceManagerImp{
		//nolint: containedctx
		ctx:          ctx,
		chain:        chain,
		chainID:      core.CopyBigInt(chainID),
		nonceMap:     make(map[common.Address]*big.Int),
		nonceMapLock: sync.RWMutex{},
		accountMutex: mapmutex.NewStringerMapMutex(),
	}
}

// GetNextNonce gets the next nonce for the account.
func (n *nonceManagerImp) GetNextNonce(address common.Address) (*big.Int, error) {
	// get the next nonce for the account
	n.nonceMapLock.Lock()
	currentNonce := n.nonceMap[address]
	defer n.nonceMapLock.Unlock()

	if currentNonce == nil {
		pendingNonce, err := n.chain.PendingNonceAt(n.ctx, address)
		if err != nil {
			return nil, fmt.Errorf("could not get pending nonce for address %s: %w", address.String(), err)
		}

		logger.Debugf("got first nonce %d for account %s", pendingNonce, address)

		// set the nonce in the map
		currentNonce = big.NewInt(int64(pendingNonce))

		n.nonceMap[address] = currentNonce
	}

	return currentNonce, nil
}

// incrementNonce increments the nonce for an account. This should be called from within a accountMutex.
func (n *nonceManagerImp) incrementNonce(address common.Address) error {
	currentNonce, err := n.GetNextNonce(address)
	if err != nil {
		return fmt.Errorf("could not get current nonce: %w", err)
	}

	n.nonceMapLock.Lock()
	n.nonceMap[address] = big.NewInt(0).Add(currentNonce, big.NewInt(1))
	n.nonceMapLock.Unlock()

	return nil
}

// NewKeyedTransactorFromKey wraps keyed transactor in a nonce manager.
// right now, this only works if all txes are sent out (a safe assumption in test mode)
// this can be obviated by signing at send time or loop + retrying on failure.
func (n *nonceManagerImp) NewKeyedTransactorFromKey(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	realSigner, err := bind.NewKeyedTransactorWithChainID(key, n.chainID)
	if err != nil {
		return nil, fmt.Errorf("could not create signer: %w", err)
	}
	return n.NewKeyedTransactor(realSigner)
}

// NewKeyedTransactor wraps keyed transactor in a nonce manager.
// right now, this only works if all txes are sent out (a safe assumption in test mode)
// this can be obviated by signing at send time or loop + retrying on failure.
func (n *nonceManagerImp) NewKeyedTransactor(realSigner *bind.TransactOpts) (*bind.TransactOpts, error) {
	return &bind.TransactOpts{
		From: realSigner.From,
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			// lock the account
			acctMutex := n.accountMutex.Lock(address)
			defer acctMutex.Unlock()

			// get the next nonce for the account
			nonce, err := n.GetNextNonce(realSigner.From)
			if err != nil {
				return nil, fmt.Errorf("could not get next nonce: %w", err)
			}

			copiedTx, err := n.copyTxWithNonce(transaction, nonce.Uint64())
			if err != nil {
				return nil, fmt.Errorf("could not copy tx: %w", err)
			}
			signedTx, err := realSigner.Signer(address, copiedTx)
			// sign was unsuccesfull, don't increment nonce
			if err != nil {
				return nil, fmt.Errorf("could not sign tx: %w", err)
			}

			err = n.incrementNonce(address)
			if err != nil {
				return nil, err
			}
			return signedTx, nil
		},
	}, nil
}

// copyTxWithNonce copies a transaction but changes the nonce.
func (n *nonceManagerImp) copyTxWithNonce(unsignedTx *types.Transaction, nonce uint64) (*types.Transaction, error) {
	// tx is immutable except within the confines of type. Here we manually copy over the inner values

	// these are overwritten, but copied over anyway for parity
	v, r, s := unsignedTx.RawSignatureValues()

	switch unsignedTx.Type() {
	case types.LegacyTxType:
		return types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: unsignedTx.GasPrice(),
			Gas:      unsignedTx.Gas(),
			To:       unsignedTx.To(),
			Value:    unsignedTx.Value(),
			Data:     unsignedTx.Data(),
			V:        v,
			R:        r,
			S:        s,
		}), nil
	case types.AccessListTxType:
		return nil, fmt.Errorf("unsupported tx type %d", types.AccessListTxType)
	case types.DynamicFeeTxType:
		return types.NewTx(&types.DynamicFeeTx{
			ChainID:    unsignedTx.ChainId(),
			Nonce:      nonce,
			GasTipCap:  unsignedTx.GasTipCap(),
			GasFeeCap:  unsignedTx.GasFeeCap(),
			Gas:        unsignedTx.Gas(),
			To:         unsignedTx.To(),
			Value:      unsignedTx.Value(),
			Data:       unsignedTx.Data(),
			AccessList: unsignedTx.AccessList(),
			V:          v,
			R:          r,
			S:          s,
		}), nil
	}
	return nil, errors.New("an unexpected error occurred")
}

// SignTx signs a legacy tx.
func (n *nonceManagerImp) SignTx(ogTx *types.Transaction, signer types.Signer, prv *ecdsa.PrivateKey) (*types.Transaction, error) {
	address := crypto.PubkeyToAddress(prv.PublicKey)

	addressLock := n.accountMutex.Lock(address)
	defer addressLock.Unlock()

	nonce, err := n.GetNextNonce(address)
	if err != nil {
		return nil, fmt.Errorf("could not get nonce: %w", err)
	}

	tx, err := n.copyTxWithNonce(ogTx, nonce.Uint64())
	if err != nil {
		return nil, fmt.Errorf("could not copy tx: %w", err)
	}

	tx, err = types.SignTx(tx, signer, prv)
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	err = n.incrementNonce(address)
	if err != nil {
		return nil, fmt.Errorf("could not increment nonce: %w", err)
	}

	return tx, nil
}
