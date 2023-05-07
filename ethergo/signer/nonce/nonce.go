package nonce

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/mapmutex"
)

// Manager is a singleton used to generate a nonce.
// This can be explicitly overrided by setting a nonce manually (not recommended). The issue here is many of the
// deploy helper functions are called asynchronously and there's a gap between the tx.Sequence generation and the
// transactor. This solves that by wrapping the transactor in a nonce manager.
type Manager interface {
	// SignTx signs a legacy tx
	SignTx(ogTx *types.Transaction, signer types.Signer, prv *ecdsa.PrivateKey, options ...Option) (*types.Transaction, error)
	// NewKeyedTransactor wraps keyed transactor in a nonce manager.
	// right now, this only works if all txes are sent out (a safe assumption in test mode)
	// this can be obviated by signing at send time or loop + retrying on failure
	NewKeyedTransactor(realSigner *bind.TransactOpts) (*bind.TransactOpts, error)
	// NewKeyedTransactorFromKey creates a new keyed transactor from a private key.
	NewKeyedTransactorFromKey(key *ecdsa.PrivateKey) (*bind.TransactOpts, error)
	// GetNextNonce gets the next nonce for the address.
	GetNextNonce(address common.Address) (*big.Int, error)
	// ClearNonce clears the nonce for the address.
	ClearNonce(address common.Address)
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
	// get the next nonce for the account.
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

// ClearNonce clears the nonce for the account.
func (n *nonceManagerImp) ClearNonce(address common.Address) {
	// clear the nonce for the account
	n.nonceMapLock.Lock()
	n.nonceMap[address] = nil
	defer n.nonceMapLock.Unlock()
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

			copiedTx, err := util.CopyTX(transaction, util.WithNonce(nonce.Uint64()))
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

// SignTx signs a legacy tx.
func (n *nonceManagerImp) SignTx(ogTx *types.Transaction, signer types.Signer, prv *ecdsa.PrivateKey, options ...Option) (*types.Transaction, error) {
	cfg := &signTXConfig{}

	for _, opt := range options {
		opt(cfg)
	}

	address := crypto.PubkeyToAddress(prv.PublicKey)

	addressLock := n.accountMutex.Lock(address)
	defer addressLock.Unlock()

	var copyOpts []util.CopyOption
	if !cfg.skipNonceBump {
		nonce, err := n.GetNextNonce(address)
		if err != nil {
			return nil, fmt.Errorf("could not get nonce: %w", err)
		}

		copyOpts = append(copyOpts, util.WithNonce(nonce.Uint64()))
	}

	tx, err := util.CopyTX(ogTx, copyOpts...)
	if err != nil {
		return nil, fmt.Errorf("could not copy tx: %w", err)
	}

	tx, err = types.SignTx(tx, signer, prv)
	if err != nil {
		return nil, fmt.Errorf("could not sign tx: %w", err)
	}

	if !cfg.skipNonceBump {
		err = n.incrementNonce(address)
		if err != nil {
			return nil, fmt.Errorf("could not increment nonce: %w", err)
		}
	}

	return tx, nil
}

// Option is a functional option for SignTX method.
type Option func(*signTXConfig)

// signTXConfig is the config for SignTX.
type signTXConfig struct {
	skipNonceBump bool
}

// WithNoBump sets the skipNonceBump flag.
func WithNoBump(noBump bool) Option {
	return func(config *signTXConfig) {
		config.skipNonceBump = noBump
	}
}
