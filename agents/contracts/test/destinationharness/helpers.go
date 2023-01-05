package destinationharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// DestinationHarnessRef is a destination harness reference
//
//nolint:golint
type DestinationHarnessRef struct {
	*DestinationHarness
	address common.Address
}

// Address gets the address of the contract.
func (r DestinationHarnessRef) Address() common.Address {
	return r.address
}

// NewDestinationHarnessRef creates a new destination harness.
func NewDestinationHarnessRef(address common.Address, backend bind.ContractBackend) (*DestinationHarnessRef, error) {
	contract, err := NewDestinationHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create destination harness: %w", err)
	}

	return &DestinationHarnessRef{
		DestinationHarness: contract,
		address:            address,
	}, nil
}

var _ vm.ContractRef = &DestinationHarnessRef{}
