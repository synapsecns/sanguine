package testswap

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestSwapRef is a bound synapse swap contract that returns the address of that contract
//
//nolint:golint
type TestSwapRef struct {
	*TestSwapFlashLoan
	address common.Address
}

// Address is the contract address.
func (s TestSwapRef) Address() common.Address {
	return s.address
}

// NewTestSwapRef gets a bound swap config contract that returns the address of the contract
//
//nolint:golint
func NewTestSwapRef(address common.Address, backend bind.ContractBackend) (*TestSwapRef, error) {
	swap, err := NewTestSwapFlashLoan(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestSwapRef{
		TestSwapFlashLoan: swap,
		address:           address,
	}, nil
}

var _ vm.ContractRef = &TestSwapRef{}
