package metaswap

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MetaSwapRef is a bound synapse bridge config v2 contract that returns the address of that contract
//
//nolint:golint
type MetaSwapRef struct {
	*MetaSwap
	address common.Address
}

// Address is the contract address.
func (s MetaSwapRef) Address() common.Address {
	return s.address
}

// NewMetaSwapRef gets a bound synapse bridge config contract that returns the address of the contract
//
//nolint:golint
func NewMetaSwapRef(address common.Address, backend bind.ContractBackend) (*MetaSwapRef, error) {
	swap, err := NewMetaSwap(address, backend)
	if err != nil {
		return nil, err
	}
	return &MetaSwapRef{
		MetaSwap: swap,
		address:  address,
	}, nil
}

var _ vm.ContractRef = &MetaSwapRef{}
