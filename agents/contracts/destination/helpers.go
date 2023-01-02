package destination

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// DestinationRef is a bound destination contract that returns the address of the destination contract.
//
//nolint:golint
type DestinationRef struct {
	*Destination
	address common.Address
}

// Address gets the address of the destination contract.
func (r DestinationRef) Address() common.Address {
	return r.address
}

// NewDestinationRef creates an destination contract with a contract ref.
func NewDestinationRef(address common.Address, backend bind.ContractBackend) (*DestinationRef, error) {
	destinationContract, err := NewDestination(address, backend)
	if err != nil {
		return nil, err
	}

	return &DestinationRef{
		Destination: destinationContract,
		address:     address,
	}, nil
}

var _ vm.ContractRef = DestinationRef{}
