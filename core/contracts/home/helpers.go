package home

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// HomeRef is a bound home contract that returns the address of the contract.
//nolint: golint
type HomeRef struct {
	*Home
	address common.Address
}

// Address is the contract address.
func (s HomeRef) Address() common.Address {
	return s.address
}

// NewHomeRef creates a new home contract with a contract ref.
func NewHomeRef(address common.Address, backend bind.ContractBackend) (*HomeRef, error) {
	homeContract, err := NewHome(address, backend)
	if err != nil {
		return nil, err
	}
	return &HomeRef{
		Home:    homeContract,
		address: address,
	}, nil
}

var _ vm.ContractRef = HomeRef{}
