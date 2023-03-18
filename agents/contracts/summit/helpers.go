package summit

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SummitRef is a bound summit contract that returns the address of the summit contract.
//
//nolint:golint
type SummitRef struct {
	*Summit
	address common.Address
}

// Address gets the address of the summit contract.
func (s SummitRef) Address() common.Address {
	return s.address
}

// NewSummitRef creates a summit contract with a contract ref.
func NewSummitRef(address common.Address, backend bind.ContractBackend) (*SummitRef, error) {
	summitContract, err := NewSummit(address, backend)
	if err != nil {
		return nil, err
	}

	return &SummitRef{
		Summit:  summitContract,
		address: address,
	}, nil
}

var _ vm.ContractRef = SummitRef{}
