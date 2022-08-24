package notarymanager

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NotaryManagerRef is a bound update manager that returns the address of the contract
// nolint: golint
type NotaryManagerRef struct {
	*NotaryManager
	address common.Address
}

// Address is the contract address.
func (u NotaryManagerRef) Address() common.Address {
	return u.address
}

// NewNotaryManagerRef gets a new update manager reference contracts.
func NewNotaryManagerRef(address common.Address, backend bind.ContractBackend) (*NotaryManagerRef, error) {
	contract, err := NewNotaryManager(address, backend)
	if err != nil {
		return nil, err
	}

	return &NotaryManagerRef{
		NotaryManager: contract,
		address:       address,
	}, nil
}
