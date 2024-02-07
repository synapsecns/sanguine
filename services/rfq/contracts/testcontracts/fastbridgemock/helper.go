package fastbridgemock

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// FastBridgeMockRef is a bound fast bridge contract that returns the address of the contract.
//
//nolint:golint
type FastBridgeMockRef struct {
	*FastBridgeMock
	address common.Address
}

// Address gets the ocntract address.
func (f *FastBridgeMockRef) Address() common.Address {
	return f.address
}

// NewFastBridgeMockRef creates a new fast bridge mock contract with a ref.
func NewFastBridgeMockRef(address common.Address, backend bind.ContractBackend) (*FastBridgeMockRef, error) {
	fastbridgemock, err := NewFastBridgeMock(address, backend)
	if err != nil {
		return nil, err
	}

	return &FastBridgeMockRef{
		FastBridgeMock: fastbridgemock,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &FastBridgeMockRef{}
