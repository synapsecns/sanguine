package mockerc20

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MockERC20Ref is a bound fast bridge contract that returns the address of the contract.
//
//nolint:golint
type MockERC20Ref struct {
	*MockERC20
	address common.Address
}

// Address gets the ocntract address.
func (f *MockERC20Ref) Address() common.Address {
	return f.address
}

// NewMockerc20Ref creates a new fast bridge contract witha  ref.
func NewMockerc20Ref(address common.Address, backend bind.ContractBackend) (*MockERC20Ref, error) {
	mockerc20, err := NewMockERC20(address, backend)
	if err != nil {
		return nil, err
	}

	return &MockERC20Ref{
		MockERC20: mockerc20,
		address:   address,
	}, nil
}

var _ vm.ContractRef = &MockERC20Ref{}
