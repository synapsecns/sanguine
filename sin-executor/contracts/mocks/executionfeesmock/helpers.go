// Package executionfeesmock provides a mock for the ExecutionFees contract.
package executionfeesmock

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ExecutionFeesMockRef is a reference to a Executionfeesmock.
// nolint: golint
type ExecutionFeesMockRef struct {
	*ExecutionFeesMock
	address common.Address
}

// Address is the contract address.
func (s *ExecutionFeesMockRef) Address() common.Address {
	return s.address
}

// NewExecutionfeesMockRef creates a new Executionfeesmock with a contract ref.
func NewExecutionfeesMockRef(address common.Address, backend bind.ContractBackend) (*ExecutionFeesMockRef, error) {
	instance, err := NewExecutionFeesMock(address, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionFeesMockRef{
		ExecutionFeesMock: instance,
		address:           address,
	}, nil
}

var _ vm.ContractRef = &ExecutionFeesMockRef{}
