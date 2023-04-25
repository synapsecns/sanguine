package lightmanager

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// LightManagerRef is a bound light manager contract that returns the address of the contract.
//
//nolint:golint
type LightManagerRef struct {
	*LightManager
	address common.Address
}

// Address is the contract address.
func (s LightManagerRef) Address() common.Address {
	return s.address
}

// NewLightManagerRef creates a new lightmanager contract with a contract ref.
func NewLightManagerRef(address common.Address, backend bind.ContractBackend) (*LightManagerRef, error) {
	lightManagerContract, err := NewLightManager(address, backend)
	if err != nil {
		return nil, err
	}

	return &LightManagerRef{
		LightManager: lightManagerContract,
		address:      address,
	}, nil
}

var _ vm.ContractRef = LightManagerRef{}

// ILightManager wraps the generated lightmanager interface code.
type ILightManager interface {
	ILightManagerCaller
	ILightManagerFilterer
	ILightManagerTransactor
	vm.ContractRef
}
