package usdt

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// USDTRef is a bound synfactory bridge contract that returns the address of the contract.
// nolint: golint
type USDTRef struct {
	*TetherToken
	address common.Address
}

// Address is the contract address.
func (s USDTRef) Address() common.Address {
	return s.address
}

// NewUSDTRef creates a new tether token.
func NewUSDTRef(address common.Address, backend bind.ContractBackend) (*USDTRef, error) {
	tetherToken, err := NewTetherToken(address, backend)
	if err != nil {
		return nil, err
	}
	return &USDTRef{
		TetherToken: tetherToken,
		address:     address,
	}, nil
}

var _ vm.ContractRef = &USDTRef{}
