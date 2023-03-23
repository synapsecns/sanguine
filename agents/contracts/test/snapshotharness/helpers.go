package snapshotharness

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SnapshotHarnessRef is a snapshot harness reference.
//
//nolint:golint
type SnapshotHarnessRef struct {
	*SnapshotHarness
	address common.Address
}

// Address gets the address of the contract.
func (s SnapshotHarnessRef) Address() common.Address {
	return s.address
}

// NewSnapshotHarnessRef creates a new snapshot harness.
func NewSnapshotHarnessRef(address common.Address, backend bind.ContractBackend) (*SnapshotHarnessRef, error) {
	contract, err := NewSnapshotHarness(address, backend)
	if err != nil {
		return nil, err
	}

	return &SnapshotHarnessRef{
		SnapshotHarness: contract,
		address:         address,
	}, nil
}

var _ vm.ContractRef = &SnapshotHarnessRef{}
