package debug

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	client2 "github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/tenderly/tenderly-cli/stacktrace"
	"math/big"
	"os"
)

// Provider provides stack trace errors.
type Provider struct {
	// contractSource is the source of the contract
	*ContractSource
}

// NewStackTraceProvider creates a new stack trace provider. Note: this currently doesn't work
// see: https://ethereum.stackexchange.com/questions/25479/how-to-map-evm-trace-to-contract-source for details.
// this will be fixed in a future version.
func NewStackTraceProvider() *Provider {
	contractSource := NewContractSource()
	return &Provider{ContractSource: contractSource}
}

// Backend is the backend used for generating stack traces.
type Backend interface {
	// RPCAddress is the rpc address
	RPCAddress() string
	// GetBigChainID gets the chainid
	GetBigChainID() *big.Int
}

// GenerateStackTrace generates a stack trace for a failed tx.
func (p Provider) GenerateStackTrace(backend Backend, tx *types.Transaction) (stackTrace string, err error) {
	if tx.To() == nil || tx.To().String() == (common.Address{}).String() {
		return stackTrace, errors.New("cannot generate stack trace for a tx with no contract")
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error while parsing stack trace %w", err)
		}
	}()

	// temporary measure. Tenderly uses logger.fatalf which calls exit
	// to fix this we're going to have to fork tenderly and use ast to override log.fatalf
	// see: https://stackoverflow.com/a/39509732 for details
	if os.Getenv("TENDERLY") == "" {
		return "", errors.New("tenderly must be enabled in order to retrieve stack traces")
	}

	client, err := MakeClient(backend.RPCAddress(), backend.GetBigChainID().String(), "", client2.ConfigFromID(backend.GetBigChainID()))
	if err != nil {
		return stackTrace, fmt.Errorf("could not connect to rpc server: %w", err)
	}

	trace, err := client.GetTransactionVMTrace(tx.Hash().String())
	if err != nil {
		return stackTrace, fmt.Errorf("could not get trace: %w", err)
	}

	core := stacktrace.NewCore(p)

	frames, err := core.GenerateStackTrace(tx.To().String(), trace)
	if err != nil {
		return stackTrace, fmt.Errorf("could not generate stack trace: %w", err)
	}

	for _, frame := range frames {
		stackTrace = fmt.Sprintf("%s \n %s", stackTrace, frame.String())
	}

	return stackTrace, nil
}
