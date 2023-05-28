package gasoracle

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// GasOracleRef is a bound gas oracle contract that returns the address of the contract.
//
//nolint:golint
type GasOracleRef struct {
	*GasOracle
	address common.Address
}

// Address is the contract address.
func (s GasOracleRef) Address() common.Address {
	return s.address
}

// NewGasOracleRef creates a new gasoracle contract with a contract ref.
func NewGasOracleRef(address common.Address, backend bind.ContractBackend) (*GasOracleRef, error) {
	gasOracleContract, err := NewGasOracle(address, backend)
	if err != nil {
		return nil, err
	}

	return &GasOracleRef{
		GasOracle: gasOracleContract,
		address:   address,
	}, nil
}

var _ vm.ContractRef = GasOracleRef{}

// IGasOracle wraps the generated gasoracle interface code.
type IGasOracle interface {
	IGasOracleCaller
	IGasOracleFilterer
	IGasOracleTransactor
	vm.ContractRef
}
