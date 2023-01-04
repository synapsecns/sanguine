package testmessagebus

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestMessageBusRef is a bound Messaging Bus Upgradeable contract and the address of the contract
//
//nolint:golint
type TestMessageBusRef struct {
	*TestMessageBusUpgradeable
	address common.Address
}

// Address is the contract address.
func (s TestMessageBusRef) Address() common.Address {
	return s.address
}

// NewTestMessageBusRef gets a bound Messaging Bus Upgradeable contract and the address of the contract
//
//nolint:golint
func NewTestMessageBusRef(address common.Address, backend bind.ContractBackend) (*TestMessageBusRef, error) {
	messageBusUpgradeable, err := NewTestMessageBusUpgradeable(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestMessageBusRef{
		TestMessageBusUpgradeable: messageBusUpgradeable,
		address:                   address,
	}, nil
}

var _ vm.ContractRef = &TestMessageBusRef{}
