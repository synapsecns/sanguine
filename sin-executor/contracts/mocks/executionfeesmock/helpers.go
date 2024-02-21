package executionfeesmock

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ExecutionfeesmockRef is a reference to a Executionfeesmock.
// nolint: golint
type ExecutionFeesmockRef struct {
	*ExecutionFeesMock
	address common.Address
}

// Address is the contract address.
func (s *ExecutionFeesmockRef) Address() common.Address {
	return s.address
}

// NewExecutionfeesMockRef creates a new Executionfeesmock with a contract ref.
func NewExecutionfeesMockRef(address common.Address, backend bind.ContractBackend) (*ExecutionFeesmockRef, error) {
	instance, err := NewExecutionFeesMock(address, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionFeesmockRef{
		ExecutionFeesMock: instance,
		address:           address,
	}, nil
}

var _ vm.ContractRef = &ExecutionFeesmockRef{}
