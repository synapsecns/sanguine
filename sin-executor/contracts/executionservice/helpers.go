// Package executionservice provides a mock for the ExecutionService contract.
package executionservice

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ExecutionServiceRef is a reference to an interchain db.
type ExecutionServiceRef struct {
	*ExecutionService
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *ExecutionServiceRef) Address() common.Address {
	return i.address
}

// NewExecutionServiceRef creates a new interchain client with a contract ref.
func NewExecutionServiceRef(address common.Address, backend bind.ContractBackend) (*ExecutionServiceRef, error) {
	instance, err := NewExecutionService(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &ExecutionServiceRef{
		ExecutionService: instance,
		address:          address,
	}, nil
}

var _ vm.ContractRef = &ExecutionServiceRef{}
