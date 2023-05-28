package inbox

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// InboxRef is a bound inbox contract that returns the address of the contract.
//
//nolint:golint
type InboxRef struct {
	*Inbox
	address common.Address
}

// Address is the contract address.
func (s InboxRef) Address() common.Address {
	return s.address
}

// NewInboxRef creates a new inbox contract with a contract ref.
func NewInboxRef(address common.Address, backend bind.ContractBackend) (*InboxRef, error) {
	inboxContract, err := NewInbox(address, backend)
	if err != nil {
		return nil, err
	}

	return &InboxRef{
		Inbox:   inboxContract,
		address: address,
	}, nil
}

var _ vm.ContractRef = InboxRef{}

// IInbox wraps the generated inbox interface code.
type IInbox interface {
	IInboxCaller
	IInboxFilterer
	IInboxTransactor
	vm.ContractRef
}
