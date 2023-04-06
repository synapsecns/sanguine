package backends

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"math/big"
	"testing"
)

// AuthType is the type used for authentication.
type AuthType struct {
	// transaction options
	*bind.TransactOpts
	// PrivateKey used for the tx
	PrivateKey *ecdsa.PrivateKey
}

// SimulatedTestBackend is a strict subset of TestBackend that all backends must comply with.
// TODO: we need one of these for testnets so we can run e2e tests. This should source addresses from a single address.
//
//go:generate go run github.com/vektra/mockery/v2 --name SimulatedTestBackend --output ./mocks --case=underscore
type SimulatedTestBackend interface {
	// EnableTenderly attempts to enable tenderly for the TestBackend. Returns false if it cannot be done
	EnableTenderly() (enabled bool)
	// BackendName gets the name of the backend
	BackendName() string
	// T is the testing.T
	T() *testing.T
	// SetT sets the testing.T
	SetT(t *testing.T)
	// Manager is used for concurrent signing while generating nonce
	nonce.Manager
	// ContractVerifier are contract verification hooks
	ContractVerifier
	// WaitForConfirmation waits for a tx confirmation
	WaitForConfirmation(ctx context.Context, transaction *types.Transaction)
	// FundAccount funds an account address with an amount amount
	FundAccount(ctx context.Context, address common.Address, amount big.Int)
	// GetTxContext gets a signed transaction. If the address is `nil`, will fund a new account.
	GetTxContext(ctx context.Context, address *common.Address) (auth AuthType)
	// GetFundedAccount gets a funded account with requestBalance
	GetFundedAccount(ctx context.Context, requestBalance *big.Int) *keystore.Key
	// Chain is the Chain
	//nolint:staticcheck
	chain.Chain
	// Signer is the signer for the chain
	Signer() types.Signer
	// ImpersonateAccount impersonates an account. This is only supported on the anvil backend backends.
	ImpersonateAccount(ctx context.Context, address common.Address, transact func(opts *bind.TransactOpts) *types.Transaction) error
}
