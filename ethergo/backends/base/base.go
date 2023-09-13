package base

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hashicorp/go-multierror"
	"github.com/ipfs/go-log"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/debug"
	"github.com/synapsecns/sanguine/ethergo/debug/tenderly"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/teivah/onecontext"
	"k8s.io/apimachinery/pkg/util/wait"
)

var logger = log.Logger("backend-base-logger")

// Backend contains common functions across backends and can be used to extend a backend.
type Backend struct {
	// chain is the chain to be used by the backend
	chain.Chain
	// Manager is the nonce manager
	nonce.Manager
	// ctx is the context of the backend
	//nolint: containedctx
	ctx context.Context
	// tb contains the testing object
	t *testing.T
	// tenderly is the tenderly backend
	tenderly *tenderly.Tenderly
	// provider is the stack trace provider
	provider *debug.Provider
}

// T returns the testing object.
func (b *Backend) T() *testing.T {
	return b.t
}

// SetT sets the testing object.
func (b *Backend) SetT(t *testing.T) {
	t.Helper()
	b.t = t
}

// NewBaseBackend creates a new base backend.
//
//nolint:staticcheck
func NewBaseBackend(ctx context.Context, t *testing.T, chn chain.Chain) (*Backend, error) {
	t.Helper()

	b := &Backend{
		Chain:    chn,
		ctx:      ctx,
		t:        t,
		Manager:  nonce.NewNonceManager(ctx, chn, chn.GetBigChainID()),
		provider: debug.NewStackTraceProvider(),
	}

	return b, nil
}

// EnableTenderly turns on tenderly on the full chain.
// Note: tenderly must be installed and you must be logged in.
func (b *Backend) EnableTenderly() bool {
	if b.tenderly != nil {
		return true
	}
	var err error
	b.tenderly, err = tenderly.NewTenderly(b.ctx)
	if err != nil {
		logger.Warnf("could not enable tenderly %v, skipping", err)
		return false
	}

	err = b.tenderly.StartListener(b)
	if err != nil {
		logger.Warnf("listener returned error: %v", err)
		return false
	}
	return true
}

// Client fetches an eth client fro the backend.
func (b *Backend) Client() client.EVMClient {
	return b.Chain
}

// see: https://git.io/JGsC1
// taken from geth, used to speed up tests.
const (
	VeryLightScryptN = 2
	VeryLightScryptP = 1
)

// MockAccount creates a new mock account.
// TODO: dry this up w/ mocks.
func MockAccount(t *testing.T) *keystore.Key {
	t.Helper()

	kstr := keystore.NewKeyStore(filet.TmpDir(t, ""), VeryLightScryptN, VeryLightScryptP)
	password := gofakeit.Password(true, true, true, false, false, 10)
	acct, err := kstr.NewAccount(password)
	assert.Nil(t, err)

	data, err := os.ReadFile(acct.URL.Path)
	assert.Nil(t, err)

	key, err := keystore.DecryptKey(data, password)
	assert.Nil(t, err)
	return key
}

// MockAccount creates a new mock account.
func (b *Backend) MockAccount() *keystore.Key {
	return MockAccount(b.t)
}

var logOnce sync.Once

// EnableLocalDebug enables local tx debugging. It is exported so it can be disabled
// and disabled by default to speed up the ci.
// Note: there's currently a bug causing this to fail if tenderly is disabled.
var EnableLocalDebug = os.Getenv("CI") == ""

// VerifyContract calls the contract verification hook (e.g. tenderly).
func (b *Backend) VerifyContract(contractType contracts.ContractType, contract contracts.DeployedContract) (resError error) {
	// TODO actually verify the contract against abi locally: https://pkg.go.dev/github.com/iden3/tx-forwarder/eth/contracts/verifier
	// until then we go ahead and run a code at to ensure the correct address was used, this helps avoid extremely hard to debug prob
	go func() {
		code, err := b.Client().CodeAt(b.ctx, contract.Address(), nil)
		if !errors.Is(err, context.Canceled) {
			require.Nil(b.T(), err)
			require.NotEmpty(b.T(), code, "contract of type %s (metadata %s) not found", contractType.ContractName(), contract.String())
		}
	}()
	var errMux sync.Mutex
	var wg sync.WaitGroup

	//nolint: nestif
	// disable this on CI, as it dramatically slows down builds
	if EnableLocalDebug {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := b.provider.AddContract(b.ctx, b.Chain, contractType, contract)
			if err != nil {
				errMux.Lock()
				resError = multierror.Append(resError, err)
				errMux.Unlock()
			}
		}()
	}

	if b.tenderly != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := b.tenderly.VerifyContract(b.ctx, b, contractType, contract)
			if err != nil {
				errMux.Lock()
				resError = multierror.Append(resError, err)
				errMux.Unlock()
			}
		}()
	}

	wg.Wait()
	return errors.Wrap(resError, "error verifying contract")
}

var (
	errorSig     = []byte{0x08, 0xc3, 0x79, 0xa0} // Keccak256("Error(string)")[:4]
	abiString, _ = abi.NewType("string", "", nil)
)

// WaitForConfirmation waits for transaction confirmation.
// nolint: cyclop
func (b *Backend) WaitForConfirmation(parentCtx context.Context, transaction *types.Transaction) {
	ctx, cancel := onecontext.Merge(b.ctx, parentCtx)
	defer cancel()

	//nolint: contextcheck
	WaitForConfirmation(ctx, b.Client(), transaction, time.Millisecond*500)
	// check or an error, if there is one log it
	go func() {
		txReceipt, err := b.TransactionReceipt(b.ctx, transaction.Hash())
		if err != nil {
			logger.Warnf("could not get tx receipt: %v on tx %s", err, transaction.Hash())
			return
		}

		callMessage, err := util.TxToCall(transaction)
		if err != nil {
			logger.Warnf("could not convert tx to call: %w", err)
			return
		}

		res, err := b.CallContract(b.ctx, *callMessage, big.NewInt(0).Sub(txReceipt.BlockNumber, big.NewInt(1)))
		if err != nil {
			errMessage := fmt.Sprintf("could not call contract: %v on tx: %s", err, transaction.Hash())
			if b.RPCAddress() != "" {
				errMessage += fmt.Sprintf("\nFor more info run (before the process stops): cast run --rpc-url %s %s --trace-printer", b.RPCAddress(), transaction.Hash())
			}
			logger.Error(errMessage)
			return
		}

		// if we are using something with an address, try to generate a stacktrace
		if b.RPCAddress() != "" && EnableLocalDebug {
			stackTrace, err := b.provider.GenerateStackTrace(b, transaction)
			if err != nil {
				logOnce.Do(func() {
					logger.Warnf("could not generate stack trace for tx: %s", transaction.Hash())
				})
			} else {
				fmt.Println(stackTrace)
				return
			}
		}
		if bytes.Equal(res, errorSig) {
			vs, err := abi.Arguments{{Type: abiString}}.UnpackValues(res[4:])
			if err != nil {
				logger.Errorf("could not unpack revert: %w", err)
				return
			}

			//nolint: forcetypeassert
			errMessage := fmt.Sprintf("tx %s reverted: %v", transaction.Hash(), vs[0].(string))
			if b.RPCAddress() != "" {
				errMessage += fmt.Sprintf("\nFor more info run (before the process stops): cast run --rpc-url %s %s --trace-printer", b.RPCAddress(), transaction.Hash())
			}
			logger.Error(errMessage)
		}
	}()
}

// Context gets the context from the backend.
func (b *Backend) Context() context.Context {
	return b.ctx
}

// ImpersonateAccount impersonates an account.
func (b *Backend) ImpersonateAccount(_ context.Context, _ common.Address, _ func(opts *bind.TransactOpts) *types.Transaction) error {
	return errors.New("account impersonation is not implemented on this backend")
}

// ConfirmationClient waits for confirmation.
//
//go:generate go run github.com/vektra/mockery/v2 --name ConfirmationClient --output ./mocks --case=underscore
type ConfirmationClient interface {
	ethereum.TransactionReader
	ethereum.TransactionSender
}

// WaitForConfirmation is a helper that can be called by various inheriting funcs.
// it blocks until the transaction is confirmed.
// nolint: cyclop
func WaitForConfirmation(ctx context.Context, client ConfirmationClient, transaction *types.Transaction, timeout time.Duration) {
	// if tx is nil , we should panic here so we can see the call context
	_ = transaction.Hash()

	txConfirmedCtx, cancel := context.WithCancel(ctx)
	var logOnce sync.Once
	wait.UntilWithContext(txConfirmedCtx, func(ctx context.Context) {
		tx, isPending, _ := client.TransactionByHash(txConfirmedCtx, transaction.Hash())
		logOnce.Do(func() {
			logger.Debugf("waiting for tx %s", transaction.Hash())
		})
		if !isPending && tx != nil {
			receipt, err := client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				if receipt.Status == types.ReceiptStatusFailed {
					rawJSON, _ := transaction.MarshalJSON()
					logger.Errorf("transaction %s with body %s reverted", transaction, string(rawJSON))
				}
			}

			cancel()
		} else if !isPending {
			_ = client.SendTransaction(ctx, transaction)
		}
	}, timeout)
}

//nolint:staticcheck
var _ chain.Chain = &Backend{}
