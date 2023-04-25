package bondingmanager

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// BondingManagerRef is a bound bonding manager contract that returns the address of the contract.
//
//nolint:golint
type BondingManagerRef struct {
	*BondingManager
	address common.Address
}

// Address is the contract address.
func (s BondingManagerRef) Address() common.Address {
	return s.address
}

// NewBondingManagerRef creates a new bondingmanager contract with a contract ref.
func NewBondingManagerRef(address common.Address, backend bind.ContractBackend) (*BondingManagerRef, error) {
	bondingManagerContract, err := NewBondingManager(address, backend)
	if err != nil {
		return nil, err
	}

	return &BondingManagerRef{
		BondingManager: bondingManagerContract,
		address:        address,
	}, nil
}

var _ vm.ContractRef = BondingManagerRef{}

// IBondingManager wraps the generated bondingmanager interface code.
type IBondingManager interface {
	IBondingManagerCaller
	IBondingManagerFilterer
	IBondingManagerTransactor
	vm.ContractRef
}
