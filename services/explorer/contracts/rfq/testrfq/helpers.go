package testrfq

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// RFQRef is a reference to a deployed RFQ contract.
//
//nolint:golint
type TestRFQRef struct {
	*TestFastBridge
	address common.Address
}

// MessageTransmitterRef is a reference to a deployed message transmitter contract.
//
//nolint:golint

// Address is the contract address.
func (s TestRFQRef) Address() common.Address {
	return s.address
}

// NewRFQRef creates a new RFQRef instance.
//
//nolint:golint
func NewTestRFQRef(address common.Address, backend bind.ContractBackend) (*TestRFQRef, error) {
	fastBridge, err := NewTestFastBridge(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestRFQRef{
		TestFastBridge: fastBridge,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &TestRFQRef{}
