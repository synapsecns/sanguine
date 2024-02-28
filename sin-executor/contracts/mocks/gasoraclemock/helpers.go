// Package gasoraclemock provides a mock for the Gasoraclemock contract.
package gasoraclemock

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// GasOracleMockMockRef is a reference to an interchain db.
type GasOracleMockMockRef struct {
	*GasOracleMock
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *GasOracleMockMockRef) Address() common.Address {
	return i.address
}

// NewGasOracleMockRef creates a new interchain client with a contract ref.
func NewGasOracleMockRef(address common.Address, backend bind.ContractBackend) (*GasOracleMockMockRef, error) {
	instance, err := NewGasOracleMock(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &GasOracleMockMockRef{
		GasOracleMock: instance,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &GasOracleMockMockRef{}
