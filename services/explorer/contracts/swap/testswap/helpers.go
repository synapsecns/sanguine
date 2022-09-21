package testswap

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SwapRef is a bound synapse swap contract that returns the address of that contract
// nolint: golint
type SwapRef struct {
	*TestSwapFlashLoan
	address common.Address
}

// Address is the contract address.
func (s SwapRef) Address() common.Address {
	return s.address
}

// NewSwapRef gets a bound swap config contract that returns the address of the contract
// nolint: golint
func NewSwapRef(address common.Address, backend bind.ContractBackend) (*SwapRef, error) {
	swap, err := NewTestSwapFlashLoan(address, backend)
	if err != nil {
		return nil, err
	}
	return &SwapRef{
		TestSwapFlashLoan: swap,
		address:           address,
	}, nil
}

var _ vm.ContractRef = &SwapRef{}
