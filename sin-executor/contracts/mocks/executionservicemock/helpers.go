// Package executionservicemock provides a mock for the ExecutionService contract.
package executionservicemock

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ExecutionServicemockRef is a reference to a Executionservicemock.
// nolint: golint
type ExecutionServicemockRef struct {
	*ExecutionServiceMock
	address common.Address
}

// Address is the contract address.
func (s *ExecutionServicemockRef) Address() common.Address {
	return s.address
}

// NewExecutionserviceMockRef creates a new Executionservicemock with a contract ref.
func NewExecutionserviceMockRef(address common.Address, backend bind.ContractBackend) (*ExecutionServicemockRef, error) {
	instance, err := NewExecutionServiceMock(address, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionServicemockRef{
		ExecutionServiceMock: instance,
		address:              address,
	}, nil
}

var _ vm.ContractRef = &ExecutionServicemockRef{}
