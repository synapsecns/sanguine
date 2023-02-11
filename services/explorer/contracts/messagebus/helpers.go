package messagebus

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MessageBusRef is a bound Message Bus Upgradeable contract and the address of the contract
//
//nolint:golint
type MessageBusRef struct {
	*MessageBusUpgradeable
	address common.Address
}

// Address is the contract address.
func (s MessageBusRef) Address() common.Address {
	return s.address
}

// NewMessageBusRef gets a bound Message Bus Upgradeable contract and the address of the contract
//
//nolint:golint
func NewMessageBusRef(address common.Address, backend bind.ContractBackend) (*MessageBusRef, error) {
	messageBusUpgradeable, err := NewMessageBusUpgradeable(address, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusRef{
		MessageBusUpgradeable: messageBusUpgradeable,
		address:               address,
	}, nil
}

var _ vm.ContractRef = &MessageBusRef{}
