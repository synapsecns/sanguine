package basemessageharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// BaseMessageHarnessRef is a base message harness reference
//
//nolint:golint
type BaseMessageHarnessRef struct {
	*BaseMessageHarness
	address common.Address
}

// Address gets the address of the contract.
func (m BaseMessageHarnessRef) Address() common.Address {
	return m.address
}

// NewBaseMessageHarnessRef creates a new message harness bound to a contract.
func NewBaseMessageHarnessRef(address common.Address, backend bind.ContractBackend) (*BaseMessageHarnessRef, error) {
	contract, err := NewBaseMessageHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create base message harness: %w", err)
	}

	return &BaseMessageHarnessRef{
		BaseMessageHarness: contract,
		address:            address,
	}, nil
}

var _ vm.ContractRef = BaseMessageHarnessRef{}
