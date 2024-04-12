package synapsemodule

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// SynapseModuleRef is a reference to a SynapseModule.
//
// nolint: golint
type SynapseModuleRef struct {
	*SynapseModule
	address common.Address
}

// Address is the contract address.
func (s *SynapseModuleRef) Address() common.Address {
	return s.address
}

// NewSynapseModuleRef creates a new SynapseModulf with a contract ref.
func NewSynapseModuleRef(address common.Address, backend bind.ContractBackend) (*SynapseModuleRef, error) {
	instance, err := NewSynapseModule(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of SynapseModule: %w", err)
	}
	return &SynapseModuleRef{
		SynapseModule: instance,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &SynapseModuleRef{}
