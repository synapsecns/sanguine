package nodeinterface

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/parser/abiutil"
	"math/big"
)

func init() {
	gasEstimateComponentsMethod = abiutil.MustGetMethodByName("gasEstimateComponents", NodeInterfaceMetaData)
}

var gasEstimateComponentsMethod abi.Method

// NodeInterfaceRef is a reference to a NodeInterface contract.
//
// nolint: golint
type NodeInterfaceRef struct {
	*NodeInterface
	address common.Address
	// required for GasEstimateComponents
	backend bind.ContractBackend
}

// Address gets the address of the contract.
func (n NodeInterfaceRef) Address() common.Address {
	return n.address
}

// NewNodeInterfaceRef creates a new NodeInterfaceRef bound to a contract.
// this returns an interface to prevent calling uncallable transactor methods.
func NewNodeInterfaceRef(address common.Address, backend bind.ContractBackend) (INodeInterface, error) {
	nodeInterface, err := NewNodeInterface(address, backend)
	if err != nil {
		return nil, err
	}

	return &NodeInterfaceRef{
		NodeInterface: nodeInterface,
		backend:       backend,
		address:       address,
	}, nil
}

// GetGasEstimateComponents returns the gas estimate components for a transaction.
// this is necessary since, in order to properly support transactions that pass in msg.value, w/ custom msg.froms, etc
// the contract interface specifies a mutable payable interface for the GasEstimateComponents method.
//
// The problem is that even though this returns some values, the current version of abigen will not generate
// retrieval callers so we need ot manually construct these here.
func (n NodeInterfaceRef) GetGasEstimateComponents(opts *bind.TransactOpts, toAddress common.Address, contractCreation bool, data []byte) (gasEstimate uint64, gasEstimateForL1 uint64, baseFee *big.Int, l1BaseFeeEstimate *big.Int, err error) {
	// see: https://github.com/Tenderly/nitro/blob/7b1d0d334e358e8d837c883f12bbf26d1900003e/system_tests/estimation_test.go#L159 for details
	estimateCallData := append([]byte{}, gasEstimateComponentsMethod.ID...)
	packed, err := gasEstimateComponentsMethod.Inputs.Pack(toAddress, contractCreation, data)
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("failed to pack inputs: %w", err)
	}

	estimateCallData = append(estimateCallData, packed...)

	msg := ethereum.CallMsg{
		From: opts.From,
		To:   &n.address,
		Data: estimateCallData,
		// see: https://github.com/Tenderly/nitro/blob/7b1d0d334e358e8d837c883f12bbf26d1900003e/system_tests/estimation_test.go#L155C16-L155C25
		Gas:       uint64(100000000),
		GasPrice:  core.CopyBigInt(opts.GasPrice),
		GasFeeCap: core.CopyBigInt(opts.GasFeeCap),
		GasTipCap: core.CopyBigInt(opts.GasTipCap),
		Value:     core.CopyBigInt(opts.Value),
	}

	returnData, err := n.backend.CallContract(opts.Context, msg, nil)
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("failed to call contract: %w", err)
	}

	outputs, err := gasEstimateComponentsMethod.Outputs.Unpack(returnData)
	if err != nil {
		return 0, 0, nil, nil, fmt.Errorf("failed to unpack outputs: %w", err)
	}

	var ok bool
	gasEstimate, ok = outputs[0].(uint64)
	if !ok {
		return 0, 0, nil, nil, fmt.Errorf("failed to convert output at index 0 to uint64")
	}
	gasEstimateForL1, ok = outputs[1].(uint64)
	if !ok {
		return 0, 0, nil, nil, fmt.Errorf("failed to convert output at index 1 to uint64")
	}
	baseFee, ok = outputs[2].(*big.Int)
	if !ok {
		return 0, 0, nil, nil, fmt.Errorf("failed to convert output at index 2 to *big.Int")
	}
	l1BaseFeeEstimate, ok = outputs[3].(*big.Int)
	if !ok {
		return 0, 0, nil, nil, fmt.Errorf("failed to convert output at index 3 to *big.Int")
	}

	return gasEstimate, gasEstimateForL1, baseFee, l1BaseFeeEstimate, nil
}

// INodeInterface INodeInterfaceCaller is a thin wrapper around NodeInterface that allows interfacing with the contract.
type INodeInterface interface {
	INodeInterfaceCaller
	vm.ContractRef
	// GetGasEstimateComponents returns the gas estimate components for a transaction.
	GetGasEstimateComponents(opts *bind.TransactOpts, toAddress common.Address, contractCreation bool, data []byte) (gasEstimate uint64, gasEstimateForL1 uint64, baseFee *big.Int, l1BaseFeeEsimate *big.Int, err error)
}
