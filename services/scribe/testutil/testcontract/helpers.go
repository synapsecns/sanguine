package testcontract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// Ref is a bound destination contract that returns the address of the destination contract.
// nolint: golint
type Ref struct {
	*TestContract
	address common.Address
}

// Address gets the address of the destination contract.
func (r Ref) Address() common.Address {
	return r.address
}

// NewTestContractRef creates an destination contract with a contract ref.
func NewTestContractRef(address common.Address, backend bind.ContractBackend) (*Ref, error) {
	testContract, err := NewTestContract(address, backend)
	if err != nil {
		return nil, err
	}

	return &Ref{
		TestContract: testContract,
		address:      address,
	}, nil
}

var _ vm.ContractRef = Ref{}
