package lightmanagerharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// LightManagerHarnessRef is a light manager harness reference
//
//nolint:golint
type LightManagerHarnessRef struct {
	*LightManagerHarness
	address common.Address
}

// Address gets the address of the contract.
func (h LightManagerHarnessRef) Address() common.Address {
	return h.address
}

// NewLightManagerHarnessRef creates a new light manager harness.
func NewLightManagerHarnessRef(address common.Address, backend bind.ContractBackend) (*LightManagerHarnessRef, error) {
	contract, err := NewLightManagerHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create light manager harness: %w", err)
	}

	return &LightManagerHarnessRef{
		LightManagerHarness: contract,
		address:             address,
	}, nil
}

var _ vm.ContractRef = &LightManagerHarnessRef{}
