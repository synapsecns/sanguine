package messageharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MessageHarnessRef is a message harness reference
//
//nolint:golint
type MessageHarnessRef struct {
	*MessageHarness
	address common.Address
}

// Address gets the address of the contract.
func (m MessageHarnessRef) Address() common.Address {
	return m.address
}

// NewMessageHarnessRef creates a new message harness bound to a contract.
func NewMessageHarnessRef(address common.Address, backend bind.ContractBackend) (*MessageHarnessRef, error) {
	contract, err := NewMessageHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create message harness: %w", err)
	}

	return &MessageHarnessRef{
		MessageHarness: contract,
		address:        address,
	}, nil
}

var _ vm.ContractRef = MessageHarnessRef{}
