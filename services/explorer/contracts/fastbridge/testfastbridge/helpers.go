package testfastbridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestFastBridgeRef is a reference to a deployed TestFastBridge contract.
//
//nolint:golint
type TestFastBridgeRef struct {
	*TestFastBridge
	address common.Address
}

// MessageTransmitterRef is a reference to a deployed message transmitter contract.
//
//nolint:golint

// Address is the contract address.
func (s TestFastBridgeRef) Address() common.Address {
	return s.address
}

// NewTestFastBridgeRef creates a new TestFastBridgeRef instance.
//
//nolint:golint
func NewTestFastBridgeRef(address common.Address, backend bind.ContractBackend) (*TestFastBridgeRef, error) {
	fastBridge, err := NewTestFastBridge(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestFastBridgeRef{
		TestFastBridge: fastBridge,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &TestFastBridgeRef{}
