package mockmessagetransmitter

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MockMessageTransmitterRef is a bound cctp contract that conforms to vm.ContractRef.
//
//nolint:golint
type MockMessageTransmitterRef struct {
	*MockMessageTransmitter
	address common.Address
}

// Address is the contract address.
func (s MockMessageTransmitterRef) Address() common.Address {
	return s.address
}

// NewMockMessageTransmitterRef creates a new MockMessageTransmitterRef contract with a contract ref.
func NewMockMessageTransmitterRef(address common.Address, backend bind.ContractBackend) (*MockMessageTransmitterRef, error) {
	cctpContract, err := NewMockMessageTransmitter(address, backend)
	if err != nil {
		return nil, err
	}

	return &MockMessageTransmitterRef{
		MockMessageTransmitter: cctpContract,
		address:                address,
	}, nil
}

var _ vm.ContractRef = MockMessageTransmitterRef{}
