package homeharness

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// HomeHarnessRef is a home harness reference
//nolint: golint
type HomeHarnessRef struct {
	*HomeHarness
	address common.Address
}

// Address gets the address of the contract.
func (h HomeHarnessRef) Address() common.Address {
	return h.address
}

// NewHomeHarnessRef creates a new home harness.
func NewHomeHarnessRef(address common.Address, backend bind.ContractBackend) (*HomeHarnessRef, error) {
	contract, err := NewHomeHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create home harness: %w", err)
	}

	return &HomeHarnessRef{
		HomeHarness: contract,
		address:     address,
	}, nil
}

var _ vm.ContractRef = &HomeHarnessRef{}
