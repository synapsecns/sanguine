package stateharness

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// StateHarnessRef is a state harness reference.
//
//nolint:golint
type StateHarnessRef struct {
	*StateHarness
	address common.Address
}

// Address gets the address of the contract.
func (s StateHarnessRef) Address() common.Address {
	return s.address
}

// NewStateHarnessRef creates a new state harness.
func NewStateHarnessRef(address common.Address, backend bind.ContractBackend) (*StateHarnessRef, error) {
	contract, err := NewStateHarness(address, backend)
	if err != nil {
		return nil, err
	}

	return &StateHarnessRef{
		StateHarness: contract,
		address:      address,
	}, nil
}

var _ vm.ContractRef = &StateHarnessRef{}
