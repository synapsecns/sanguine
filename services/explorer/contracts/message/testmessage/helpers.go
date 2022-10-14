package testmessage

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestMessageRef is a bound Message Bus Upgradeable contract and the address of the contract
// nolint: golint
type TestMessageRef struct {
	*TestMessageBusUpgradeable
	address common.Address
}

// Address is the contract address.
func (s TestMessageRef) Address() common.Address {
	return s.address
}

// NewMessageRef gets a bound Message Bus Upgradeable contract and the address of the contract
// nolint: golint
func NewMessageRef(address common.Address, backend bind.ContractBackend) (*TestMessageRef, error) {
	messageBusUpgradeable, err := NewTestMessageBusUpgradeable(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestMessageRef{
		TestMessageBusUpgradeable: messageBusUpgradeable,
		address:                   address,
	}, nil
}

var _ vm.ContractRef = &TestMessageRef{}
