package gasoraclemock

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// GasOracleMockRef is a reference to a Gasoraclemock.
// nolint: golint
type GasOracleMockRef struct {
	*SynapseGasOracleMock
	address common.Address
}

// Address is the contract address.
func (s *GasOracleMockRef) Address() common.Address {
	return s.address
}

// NewGasOracleMockRef creates a new Gasoraclemock with a contract ref.
func NewGasOracleMockRef(address common.Address, backend bind.ContractBackend) (*GasOracleMockRef, error) {
	instance, err := NewSynapseGasOracleMock(address, backend)
	if err != nil {
		return nil, err
	}
	return &GasOracleMockRef{
		SynapseGasOracleMock: instance,
		address:              address,
	}, nil
}

var _ vm.ContractRef = &GasOracleMockRef{}
