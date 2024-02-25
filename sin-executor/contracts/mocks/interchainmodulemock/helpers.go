// Package interchainmodulemock provides a mock for the InterchainModule contract.
package interchainmodulemock

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// InterchainModuleMockRef is a reference to an interchain db.
type InterchainModuleMockRef struct {
	*InterchainModuleMock
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *InterchainModuleMockRef) Address() common.Address {
	return i.address
}

// NewInterchainModuleMockRef creates a new interchain client with a contract ref.
func NewInterchainModuleMockRef(address common.Address, backend bind.ContractBackend) (*InterchainModuleMockRef, error) {
	instance, err := NewInterchainModuleMock(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &InterchainModuleMockRef{
		InterchainModuleMock: instance,
		address:              address,
	}, nil
}

var _ vm.ContractRef = &InterchainModuleMockRef{}
