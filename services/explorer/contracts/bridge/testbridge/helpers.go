package testbridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestBridgeRef is a bound synapse bridge config v2 contract that returns the address of that contract
//
//nolint:golint
type TestBridgeRef struct {
	*TestSynapseBridge
	address common.Address
}

// Address is the contract address.
func (s TestBridgeRef) Address() common.Address {
	return s.address
}

// NewTestBridgeRef gets a bound synapse bridge config contract that returns the address of the contract
//
//nolint:golint
func NewTestBridgeRef(address common.Address, backend bind.ContractBackend) (*TestBridgeRef, error) {
	bridge, err := NewTestSynapseBridge(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestBridgeRef{
		TestSynapseBridge: bridge,
		address:           address,
	}, nil
}

var _ vm.ContractRef = &TestBridgeRef{}
