package summit

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SummitRef is a summit contract that is bound to a specific deployed contract
//
//nolint:golint
type SummitRef struct {
	*Summit
	address common.Address
}

// Address is the contract address.
func (s SummitRef) Address() common.Address {
	return s.address
}

// NewSummitRef creates the SummitRef type.
//
//nolint:golint
func NewSummitRef(address common.Address, backend bind.ContractBackend) (*SummitRef, error) {
	swap, err := NewSummit(address, backend)
	if err != nil {
		return nil, err
	}
	return &SummitRef{
		Summit:  swap,
		address: address,
	}, nil
}

var _ vm.ContractRef = &SummitRef{}
