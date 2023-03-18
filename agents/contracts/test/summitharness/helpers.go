package summitharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SummitHarnessRef is a summit harness reference
//
//nolint:golint
type SummitHarnessRef struct {
	*SummitHarness
	address common.Address
}

// Address gets the address of the contract.
func (r SummitHarnessRef) Address() common.Address {
	return r.address
}

// NewSummitHarnessRef creates a new summit harness.
func NewSummitHarnessRef(address common.Address, backend bind.ContractBackend) (*SummitHarnessRef, error) {
	contract, err := NewSummitHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create summit harness: %w", err)
	}

	return &SummitHarnessRef{
		SummitHarness: contract,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &SummitHarnessRef{}
