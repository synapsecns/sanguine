package ierc20

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// IERC20Ref is a bound fast bridge contract that returns the address of the contract.
//
//nolint:golint
type IERC20Ref struct {
	*IERC20
	address common.Address
}

// Address gets the ocntract address.
func (f *IERC20Ref) Address() common.Address {
	return f.address
}

// NewIerc20Ref creates a new fast bridge contract witha  ref.
func NewIerc20Ref(address common.Address, backend bind.ContractBackend) (*IERC20Ref, error) {
	ierc20, err := NewIERC20(address, backend)
	if err != nil {
		return nil, err
	}

	return &IERC20Ref{
		IERC20:  ierc20,
		address: address,
	}, nil
}

var _ vm.ContractRef = &IERC20Ref{}
