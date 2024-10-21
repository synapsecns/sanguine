package recipientmock

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// RecipientMockRef is a bound fast bridge contract that returns the address of the contract.
//
//nolint:golint
type RecipientMockRef struct {
	*RecipientMock
	address common.Address
}

// Address gets the ocntract address.
func (f *RecipientMockRef) Address() common.Address {
	return f.address
}

// NewRecipientMockRef creates a new fast bridge mock contract with a ref.
func NewRecipientMockRef(address common.Address, backend bind.ContractBackend) (*RecipientMockRef, error) {
	recipientmock, err := NewRecipientMock(address, backend)
	if err != nil {
		return nil, err
	}

	return &RecipientMockRef{
		RecipientMock: recipientmock,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &RecipientMockRef{}
