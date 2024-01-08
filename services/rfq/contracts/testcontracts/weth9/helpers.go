package weth9

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// Weth9Ref is a bound WETH9 bridge that returns the address of the contract.
// nolint: golint
type Weth9Ref struct {
	*WETH9
	address common.Address
}

// Address is the contract address.
func (s Weth9Ref) Address() common.Address {
	return s.address
}

// NewWeth9Ref creates a new weth9 token.
func NewWeth9Ref(address common.Address, backend bind.ContractBackend) (*Weth9Ref, error) {
	wethToken, err := NewWETH9(address, backend)
	if err != nil {
		return nil, err
	}
	return &Weth9Ref{
		WETH9:   wethToken,
		address: address,
	}, nil
}

var _ vm.ContractRef = &Weth9Ref{}
