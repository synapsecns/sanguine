// Package interchainclient provides the interchainclient interface.
package interchainclient

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// InterchainClientRef is a reference to an interchain client.
type InterchainClientRef struct {
	*InterchainClientV1
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *InterchainClientRef) Address() common.Address {
	return i.address
}

// NewInterchainClientRef creates a new interchain client with a contract ref.
func NewInterchainClientRef(address common.Address, backend bind.ContractBackend) (*InterchainClientRef, error) {
	instance, err := NewInterchainClientV1(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &InterchainClientRef{
		InterchainClientV1: instance,
		address:            address,
	}, nil
}

var _ vm.ContractRef = &InterchainClientRef{}
