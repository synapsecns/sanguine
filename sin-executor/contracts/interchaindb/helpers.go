// Package interchaindb provides a mock for the InterchainDB contract.
package interchaindb

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// InterchainDBRef is a reference to an interchain db.
type InterchainDBRef struct {
	*InterchainDB
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *InterchainDBRef) Address() common.Address {
	return i.address
}

// NewInterchainDBRef creates a new interchain client with a contract ref.
func NewInterchainDBRef(address common.Address, backend bind.ContractBackend) (*InterchainDBRef, error) {
	instance, err := NewInterchainDB(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &InterchainDBRef{
		InterchainDB: instance,
		address:      address,
	}, nil
}

var _ vm.ContractRef = &InterchainDBRef{}
