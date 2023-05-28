package lightinbox

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// LightInboxRef is a bound light inbox contract that returns the address of the contract.
//
//nolint:golint
type LightInboxRef struct {
	*LightInbox
	address common.Address
}

// Address is the contract address.
func (s LightInboxRef) Address() common.Address {
	return s.address
}

// NewLightInboxRef creates a new lightinbox contract with a contract ref.
func NewLightInboxRef(address common.Address, backend bind.ContractBackend) (*LightInboxRef, error) {
	lightInboxContract, err := NewLightInbox(address, backend)
	if err != nil {
		return nil, err
	}

	return &LightInboxRef{
		LightInbox: lightInboxContract,
		address:    address,
	}, nil
}

var _ vm.ContractRef = LightInboxRef{}

// ILightInbox wraps the generated lightinbox interface code.
type ILightInbox interface {
	ILightInboxCaller
	ILightInboxFilterer
	ILightInboxTransactor
	vm.ContractRef
}
