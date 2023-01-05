package agentstestcontract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// AgentsTestContractRef is a bound agents test contract that returns the address of the agents test contract.
//
//nolint:golint,revive
type AgentsTestContractRef struct {
	*AgentsTestContract
	address common.Address
}

// Address gets the address of the agents test contract.
func (r AgentsTestContractRef) Address() common.Address {
	return r.address
}

// NewAgentsTestContractRef creates an agents test contract with a contract ref.
func NewAgentsTestContractRef(address common.Address, backend bind.ContractBackend) (*AgentsTestContractRef, error) {
	testContract, err := NewAgentsTestContract(address, backend)
	if err != nil {
		return nil, err
	}

	return &AgentsTestContractRef{
		AgentsTestContract: testContract,
		address:            address,
	}, nil
}

var _ vm.ContractRef = AgentsTestContractRef{}
