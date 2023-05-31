package bondingmanagerharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// BondingManagerHarnessRef is a bonding manager harness reference
//
//nolint:golint
type BondingManagerHarnessRef struct {
	*BondingManagerHarness
	address common.Address
}

// Address gets the address of the contract.
func (h BondingManagerHarnessRef) Address() common.Address {
	return h.address
}

// NewBondingManagerHarnessRef creates a new bonding manager harness.
func NewBondingManagerHarnessRef(address common.Address, backend bind.ContractBackend) (*BondingManagerHarnessRef, error) {
	contract, err := NewBondingManagerHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create bonding manager harness: %w", err)
	}

	return &BondingManagerHarnessRef{
		BondingManagerHarness: contract,
		address:               address,
	}, nil
}

var _ vm.ContractRef = &BondingManagerHarnessRef{}
