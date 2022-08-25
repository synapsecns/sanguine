package base

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hashicorp/go-multierror"
	"github.com/ipfs/go-log"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	synapseCommon "github.com/synapsecns/synapse-node/pkg/common"
	"github.com/synapsecns/synapse-node/pkg/evm"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"github.com/synapsecns/synapse-node/testutils/debug/stacktrace"
	tenderly "github.com/synapsecns/synapse-node/testutils/debug/tenderly"
	"github.com/teivah/onecontext"
	"k8s.io/apimachinery/pkg/util/wait"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"
)

var logger = log.Logger("backend-base-logger")

// Backend contains common functions across backends and can be used to extend a backend.
type Backend struct {
	// chain is the chain to be used by the backend
	evm.Chain
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
	provider *stacktrace.Provider
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
func NewBaseBackend(ctx context.Context, t *testing.T, chn evm.Chain) (*Backend, error) {
	t.Helper()

	b := &Backend{
		Chain:    chn,
		ctx:      ctx,
		t:        t,
		Manager:  nonce.NewNonceManager(ctx, chn, chn.GetBigChainID()),
		provider: stacktrace.NewStackTraceProvider(),
	}

	return b, nil
}

// EnableTenderly turns on tenderly on the full chain.
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
			assert.Nil(b.T(), err)
			assert.NotEmpty(b.T(), code)
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

		callMessage, err := synapseCommon.TxToCall(transaction)
		if err != nil {
			logger.Warnf("could not convert tx to call: %w", err)
			return
		}

		res, err := b.CallContract(b.ctx, *callMessage, big.NewInt(0).Sub(txReceipt.BlockNumber, big.NewInt(1)))
		if err != nil {
			logger.Warnf("could not call contract: %v on tx: %s", err, transaction.Hash())
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
			logger.Debugf("tx %s reverted: %v", transaction.Hash(), vs[0].(string))
		}
	}()
}

// Context gets the context from the backend.
func (b *Backend) Context() context.Context {
	return b.ctx
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
func WaitForConfirmation(ctx context.Context, client ConfirmationClient, transaction *types.Transaction, timeout time.Duration) {
	txConfirmedCtx, cancel := context.WithCancel(ctx)
	var logOnce sync.Once
	wait.UntilWithContext(txConfirmedCtx, func(ctx context.Context) {
		tx, isPending, _ := client.TransactionByHash(txConfirmedCtx, transaction.Hash())
		logOnce.Do(func() {
			logger.Warnf("waiting for tx %s", synapseCommon.GetChainEventLogText(transaction.ChainId(), transaction))
		})
		if !isPending && tx != nil {
			cancel()
		} else if !isPending {
			_ = client.SendTransaction(ctx, transaction)
		}
	}, timeout)
}

var _ evm.Chain = &Backend{}
var _ suite.TestingSuite = &Backend{}
