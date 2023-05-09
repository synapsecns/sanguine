package destination

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// OriginRef is a origin contract that is bound to a specific deployed contract
//
//nolint:golint
type OriginRef struct {
	*Origin
	address common.Address
}

// Address is the contract address.
func (s OriginRef) Address() common.Address {
	return s.address
}

// NewOriginRef gets a bound synapse bridge config contract that returns the address of the contract
//
//nolint:golint
func NewOriginRef(address common.Address, backend bind.ContractBackend) (*OriginRef, error) {
	swap, err := NewOrigin(address, backend)
	if err != nil {
		return nil, err
	}
	return &OriginRef{
		Origin:  swap,
		address: address,
	}, nil
}

var _ vm.ContractRef = &OriginRef{}
