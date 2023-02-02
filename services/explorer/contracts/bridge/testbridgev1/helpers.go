package testbridgev1

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestBridgeV1Ref is a bound synapse bridge v1 contract that returns the address of that contract
//
//nolint:golint
type TestBridgeV1Ref struct {
	*TestSynapseBridgeV1
	address common.Address
}

// Address is the contract address.
func (s TestBridgeV1Ref) Address() common.Address {
	return s.address
}

// NewTestBridgeV1Ref gets a bound synapse bridge config contract that returns the address of the contract
//
//nolint:golint
func NewTestBridgeV1Ref(address common.Address, backend bind.ContractBackend) (*TestBridgeV1Ref, error) {
	bridge, err := NewTestSynapseBridgeV1(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestBridgeV1Ref{
		TestSynapseBridgeV1: bridge,
		address:             address,
	}, nil
}

var _ vm.ContractRef = &TestBridgeV1Ref{}
