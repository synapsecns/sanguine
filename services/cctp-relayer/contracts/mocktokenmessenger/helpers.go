package mocktokenmessenger

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MockTokenMessengerRef is a bound cctp contract that conforms to vm.ContractRef.
//
//nolint:golint
type MockTokenMessengerRef struct {
	*MockTokenMessenger
	address common.Address
}

// Address is the contract address.
func (m MockTokenMessengerRef) Address() common.Address {
	return m.address
}

// NewMockTokenMessengerRef creates a new MockMessageTransmitterRef contract with a contract ref.
//
//nolint:golint
func NewMockTokenMessengerRef(address common.Address, backend bind.ContractBackend) (*MockTokenMessengerRef, error) {
	cctpContract, err := NewMockTokenMessenger(address, backend)
	if err != nil {
		return nil, err
	}

	return &MockTokenMessengerRef{
		MockTokenMessenger: cctpContract,
		address:            address,
	}, nil
}

var _ vm.ContractRef = MockTokenMessengerRef{}
