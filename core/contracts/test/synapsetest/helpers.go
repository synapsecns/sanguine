package synapsetest

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SynapseTestRef is a synapse test reference
//nolint: golint
type SynapseTestRef struct {
	*SynapseTest
	address common.Address
}

// Address gets the address of the contract.
func (h SynapseTestRef) Address() common.Address {
	return h.address
}

// NewSynapseTestRef creates a new synapse test contract.
func NewSynapseTestRef(address common.Address, backend bind.ContractBackend) (*SynapseTestRef, error) {
	contract, err := NewSynapseTest(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create synapse test: %w", err)
	}

	return &SynapseTestRef{
		SynapseTest: contract,
		address:     address,
	}, nil
}

var _ vm.ContractRef = &SynapseTestRef{}
