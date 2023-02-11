package originharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// OriginHarnessRef is a origin harness reference
//
//nolint:golint
type OriginHarnessRef struct {
	*OriginHarness
	address common.Address
}

// Address gets the address of the contract.
func (h OriginHarnessRef) Address() common.Address {
	return h.address
}

// NewOriginHarnessRef creates a new origin harness.
func NewOriginHarnessRef(address common.Address, backend bind.ContractBackend) (*OriginHarnessRef, error) {
	contract, err := NewOriginHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create origin harness: %w", err)
	}

	return &OriginHarnessRef{
		OriginHarness: contract,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &OriginHarnessRef{}
