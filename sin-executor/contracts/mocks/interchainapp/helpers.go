// Package interchainapp provides a mock for the InterchainApp contract.
package interchainapp

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// InterchainAppMockRef is a reference to an interchain db.
type InterchainAppMockRef struct {
	*InterchainApp
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *InterchainAppMockRef) Address() common.Address {
	return i.address
}

// NewInterchainAppRef creates a new interchain client with a contract ref.
func NewInterchainAppRef(address common.Address, backend bind.ContractBackend) (*InterchainAppMockRef, error) {
	instance, err := NewInterchainApp(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &InterchainAppMockRef{
		InterchainApp: instance,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &InterchainAppMockRef{}
