package tipsharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TipsHarnessRef is a tips harness reference
//
//nolint:golint
type TipsHarnessRef struct {
	*TipsHarness
	address common.Address
}

// Address gets the address of the contract.
func (h TipsHarnessRef) Address() common.Address {
	return h.address
}

// NewTipsHarnessRef creates a new tips harness.
func NewTipsHarnessRef(address common.Address, backend bind.ContractBackend) (*TipsHarnessRef, error) {
	contract, err := NewTipsHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create tips harness: %w", err)
	}

	return &TipsHarnessRef{
		TipsHarness: contract,
		address:     address,
	}, nil
}

var _ vm.ContractRef = &TipsHarnessRef{}
