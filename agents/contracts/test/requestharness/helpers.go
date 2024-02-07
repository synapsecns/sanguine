package requestharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// RequestHarnessRef is a request harness reference
//
//nolint:golint
type RequestHarnessRef struct {
	*RequestHarness
	address common.Address
}

// Address gets the address of the contract.
func (m RequestHarnessRef) Address() common.Address {
	return m.address
}

// NewRequestHarnessRef creates a new request harness bound to a contract.
func NewRequestHarnessRef(address common.Address, backend bind.ContractBackend) (*RequestHarnessRef, error) {
	contract, err := NewRequestHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create request harness: %w", err)
	}

	return &RequestHarnessRef{
		RequestHarness: contract,
		address:        address,
	}, nil
}

var _ vm.ContractRef = RequestHarnessRef{}
