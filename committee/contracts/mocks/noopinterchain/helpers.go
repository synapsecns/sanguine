package noopinterchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// NoOpInterchainRef is a reference to a NoOpInterchain.
// nolint: golint
type NoOpInterchainRef struct {
	*NoOpInterchain
	address common.Address
}

// Address is the contract address.
func (s *NoOpInterchainRef) Address() common.Address {
	return s.address
}

// NewNoOpInterchainRef creates a new NoOpInterchain with a contract ref.
func NewNoOpInterchainRef(address common.Address, backend bind.ContractBackend) (*NoOpInterchainRef, error) {
	instance, err := NewNoOpInterchain(address, backend)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchainRef{
		NoOpInterchain: instance,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &NoOpInterchainRef{}
