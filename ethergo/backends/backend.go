package backends

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"math/big"
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
type SimulatedTestBackend interface {
	// EnableTenderly attempts to enable tenderly for the TestBackend. Returns false if it cannot be done
	EnableTenderly() (enabled bool)
	// BackendName gets the name of the backend
	BackendName() string
	// TestingSuite allows access to T() and SetT() methods for testing
	suite.TestingSuite
	// Manager is used for concurrent signing while generating nonce
	nonce.Manager
	// ChainConfig gets the chain config
	ChainConfig() *params.ChainConfig
	// ContractVerifier are contract verification hooks
	ContractVerifier
	// WaitForConfirmation waits for a tx confirmation
	WaitForConfirmation(ctx context.Context, transaction *types.Transaction)
	// FundAccount funds an account address with an amount amount
	FundAccount(ctx context.Context, address common.Address, amount big.Int)
	// GetTxContext gets a signed transaction
	GetTxContext(ctx context.Context, address *common.Address) (auth AuthType)
	// GetFundedAccount gets a funded account with requestBalance
	GetFundedAccount(ctx context.Context, requestBalance *big.Int) *keystore.Key
	// Chain is the Chain
	//nolint:staticcheck
	chain.Chain
	// Signer is the signer for the chain
	Signer() types.Signer
}

// TestBackend provides a backend for testing.
// Deprecated: use simulated test backend.
type TestBackend interface {
	// SimulatedTestBackend is the base of a test backend
	SimulatedTestBackend
}
