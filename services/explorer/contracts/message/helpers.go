package message

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MessageRef is a bound Message Bus Upgradeable contract and the address of the contract
// nolint: golint
type MessageRef struct {
	*MessageBusUpgradeable
	address common.Address
}

// Address is the contract address.
func (s MessageRef) Address() common.Address {
	return s.address
}

// NewMessageRef gets a bound Message Bus Upgradeable contract and the address of the contract
// nolint: golint
func NewMessageRef(address common.Address, backend bind.ContractBackend) (*MessageRef, error) {
	messageBusUpgradeable, err := NewMessageBusUpgradeable(address, backend)
	if err != nil {
		return nil, err
	}
	return &MessageRef{
		MessageBusUpgradeable: messageBusUpgradeable,
		address:               address,
	}, nil
}

var _ vm.ContractRef = &MessageRef{}
