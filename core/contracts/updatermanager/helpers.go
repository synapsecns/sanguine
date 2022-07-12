package updatermanager

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// UpdaterManagerRef is a bound update manager that returns the address of the contract
//nolint: golint
type UpdaterManagerRef struct {
	*UpdaterManager
	address common.Address
}

// Address is the contract address.
func (u UpdaterManagerRef) Address() common.Address {
	return u.address
}

// NewUpdaterManagerRef gets a new update manager reference contracts.
func NewUpdaterManagerRef(address common.Address, backend bind.ContractBackend) (*UpdaterManagerRef, error) {
	contract, err := NewUpdaterManager(address, backend)
	if err != nil {
		return nil, err
	}

	return &UpdaterManagerRef{
		UpdaterManager: contract,
		address:        address,
	}, nil
}
