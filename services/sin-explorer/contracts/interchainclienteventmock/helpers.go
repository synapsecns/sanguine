package interchainclienteventmock

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// InterchainClientEventMockRef is a bound synapse interchainclienteventmock config v2 contract that returns the address of that contract
// 233esdc
//
//nolint:golint
type InterchainClientEventMockRef struct {
	*InterchainClientEventMock
	address common.Address
}

// Address is the contract address.
func (s InterchainClientEventMockRef) Address() common.Address {
	return s.address
}

// NewInterchainClientEventMockRef gets a bound synapse interchainclienteventmock config contract that returns the address of the contract
//
//nolint:golint
func NewInterchainClientEventMockRef(address common.Address, backend bind.ContractBackend) (*InterchainClientEventMockRef, error) {
	interchainclienteventmock, err := NewInterchainClientEventMock(address, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockRef{
		InterchainClientEventMock: interchainclienteventmock,
		address:                   address,
	}, nil
}

var _ vm.ContractRef = &InterchainClientEventMockRef{}
