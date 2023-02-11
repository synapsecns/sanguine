package headerharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// HeaderHarnessRef is a header harness reference
//
//nolint:golint
type HeaderHarnessRef struct {
	*HeaderHarness
	address common.Address
}

// Address gets the address of the contract.
func (h HeaderHarnessRef) Address() common.Address {
	return h.address
}

// NewHeaderHarnessRef creates a new header harness.
func NewHeaderHarnessRef(address common.Address, backend bind.ContractBackend) (*HeaderHarnessRef, error) {
	contract, err := NewHeaderHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create header harness: %w", err)
	}

	return &HeaderHarnessRef{
		HeaderHarness: contract,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &HeaderHarnessRef{}
