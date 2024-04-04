// Package executionservice provides a mock for the ExecutionService contract.
package executionservice

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SynapseExecutionServiceV1HarnessRef is a reference to an interchain db.
// nolint: golint
type SynapseExecutionServiceV1HarnessRef struct {
	*SynapseExecutionServiceV1Harness
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *SynapseExecutionServiceV1HarnessRef) Address() common.Address {
	return i.address
}

// NewSynapseExecutionServiceV1HarnessRef creates a new interchain client with a contract ref.
func NewSynapseExecutionServiceV1HarnessRef(address common.Address, backend bind.ContractBackend) (*SynapseExecutionServiceV1HarnessRef, error) {
	instance, err := NewSynapseExecutionServiceV1Harness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &SynapseExecutionServiceV1HarnessRef{
		SynapseExecutionServiceV1Harness: instance,
		address:                          address,
	}, nil
}

var _ vm.ContractRef = &SynapseExecutionServiceV1HarnessRef{}
