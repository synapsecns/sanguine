package swap

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SwapRef is a bound synapse bridge config v2 contract that returns the address of that contract
//
//nolint:golint
type SwapRef struct {
	*SwapFlashLoan
	address common.Address
}

// Address is the contract address.
func (s SwapRef) Address() common.Address {
	return s.address
}

// NewSwapRef gets a bound synapse bridge config contract that returns the address of the contract
//
//nolint:golint
func NewSwapRef(address common.Address, backend bind.ContractBackend) (*SwapRef, error) {
	swap, err := NewSwapFlashLoan(address, backend)
	if err != nil {
		return nil, err
	}
	return &SwapRef{
		SwapFlashLoan: swap,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &SwapRef{}
