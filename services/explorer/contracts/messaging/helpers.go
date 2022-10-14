package messaging

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MessagingRef isa bound Message Bus Upgradeable contract and the address of the contract
// nolint: golint
type MessagingRef struct {
	*MessageBusUpgradeable
	address common.Address
}

// Address is the contract address.
func (s MessagingRef) Address() common.Address {
	return s.address
}

// NewMessagingRef gets a bound Message Bus Upgradeable contract and the address of the contract
// nolint: golint
func NewMessagingRef(address common.Address, backend bind.ContractBackend) (*MessagingRef, error) {
	messageBusUpgradeable, err := NewMessageBusUpgradeable(address, backend)
	if err != nil {
		return nil, err
	}
	return &MessagingRef{
		MessageBusUpgradeable: messageBusUpgradeable,
		address:               address,
	}, nil
}

var _ vm.ContractRef = &MessagingRef{}
