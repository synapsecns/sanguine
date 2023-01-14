package testmetaswap

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestMetaSwapRef is a bound synapse bridge config v2 contract that returns the address of that contract
//
//nolint:golint
type TestMetaSwapRef struct {
	*TestMetaSwap
	address common.Address
}

// Address is the contract address.
func (s TestMetaSwapRef) Address() common.Address {
	return s.address
}

// NewTestMetaSwapRef gets a bound synapse bridge config contract that returns the address of the contract
//
//nolint:golint
func NewTestMetaSwapRef(address common.Address, backend bind.ContractBackend) (*TestMetaSwapRef, error) {
	swap, err := NewTestMetaSwap(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestMetaSwapRef{
		TestMetaSwap: swap,
		address:      address,
	}, nil
}

var _ vm.ContractRef = &TestMetaSwapRef{}
