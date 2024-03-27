package interchaindb

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// InterchainDBRef is a reference to a InterchainDB.
//
// nolint: golint
type InterchainDBRef struct {
	*InterchainDB
	address common.Address
}

// Address is the contract address.
func (s *InterchainDBRef) Address() common.Address {
	return s.address
}

// NewInterchainDBRef creates a new SynapseModulf with a contract ref.
func NewInterchainDBRef(address common.Address, backend bind.ContractBackend) (*InterchainDBRef, error) {
	instance, err := NewInterchainDB(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainDB: %w", err)
	}
	return &InterchainDBRef{
		InterchainDB: instance,
		address:      address,
	}, nil
}

var _ vm.ContractRef = &InterchainDBRef{}
