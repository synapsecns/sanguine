package gasdataharness

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// GasDataHarnessRef is a gasData harness reference.
//
//nolint:golint
type GasDataHarnessRef struct {
	*GasDataHarness
	address common.Address
}

// Address gets the address of the contract.
func (s GasDataHarnessRef) Address() common.Address {
	return s.address
}

// NewGasDataHarnessRef creates a new gasData harness.
func NewGasDataHarnessRef(address common.Address, backend bind.ContractBackend) (*GasDataHarnessRef, error) {
	contract, err := NewGasDataHarness(address, backend)
	if err != nil {
		return nil, err
	}

	return &GasDataHarnessRef{
		GasDataHarness: contract,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &GasDataHarnessRef{}
