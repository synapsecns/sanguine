package replicamanager

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ReplicaManagerRef is a bound replica manager contract that returns the address of the replica manager contract.
//nolint: golint
type ReplicaManagerRef struct {
	*ReplicaManager
	address common.Address
}

// Address gets the address of the replica manager contract.
func (r ReplicaManagerRef) Address() common.Address {
	return r.address
}

// NewReplicaManagerRef creates an replica manager contract with a contract ref.
func NewReplicaManagerRef(address common.Address, backend bind.ContractBackend) (*ReplicaManagerRef, error) {
	replicaManagerContract, err := NewReplicaManager(address, backend)
	if err != nil {
		return nil, err
	}

	return &ReplicaManagerRef{
		ReplicaManager: replicaManagerContract,
		address:        address,
	}, nil
}

var _ vm.ContractRef = ReplicaManagerRef{}
