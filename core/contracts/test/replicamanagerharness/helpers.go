package replicamanagerharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ReplicaManagerHarnessRef is a replica manager harness reference
//nolint: golint
type ReplicaManagerHarnessRef struct {
	*ReplicaManagerHarness
	address common.Address
}

// Address gets the address of the contract.
func (r ReplicaManagerHarnessRef) Address() common.Address {
	return r.address
}

// NewReplicaManagerHarnessRef creates a new replica manager harness.
func NewReplicaManagerHarnessRef(address common.Address, backend bind.ContractBackend) (*ReplicaManagerHarnessRef, error) {
	contract, err := NewReplicaManagerHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create replica manager harness: %w", err)
	}

	return &ReplicaManagerHarnessRef{
		ReplicaManagerHarness: contract,
		address:               address,
	}, nil
}

var _ vm.ContractRef = &ReplicaManagerHarnessRef{}
