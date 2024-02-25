// Package optionslibexport provides a mock for the OptionsLibMocks contract.
package optionslibexport

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// OptionsLibExportRef is a reference to an interchain db.
type OptionsLibExportRef struct {
	*OptionsLibMocks
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *OptionsLibExportRef) Address() common.Address {
	return i.address
}

// NewOptionsLibExportRef creates a new interchain client with a contract ref.
func NewOptionsLibExportRef(address common.Address, backend bind.ContractBackend) (*OptionsLibExportRef, error) {
	instance, err := NewOptionsLibMocks(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &OptionsLibExportRef{
		OptionsLibMocks: instance,
		address:         address,
	}, nil
}

var _ vm.ContractRef = &OptionsLibExportRef{}
