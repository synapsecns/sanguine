package mocktokenminter

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MockTokenMinterRef is a bound cctp contract that conforms to vm.ContractRef.
//
//nolint:golint
type MockTokenMinterRef struct {
	*MockTokenMinter
	address common.Address
}

// Address is the contract address.
func (m MockTokenMinterRef) Address() common.Address {
	return m.address
}

// NewMockTokenMinterRef creates a new MockMessageTransmitterRef contract with a contract ref.
//
//nolint:golint
func NewMockTokenMinterRef(address common.Address, backend bind.ContractBackend) (*MockTokenMinterRef, error) {
	cctpContract, err := NewMockTokenMinter(address, backend)
	if err != nil {
		return nil, err
	}

	return &MockTokenMinterRef{
		MockTokenMinter: cctpContract,
		address:         address,
	}, nil
}

var _ vm.ContractRef = MockTokenMinterRef{}
