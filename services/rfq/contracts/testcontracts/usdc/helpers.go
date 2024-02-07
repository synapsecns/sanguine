package usdc

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// USDCRef is a bound synfactory bridge contract that returns the address of the contract.
// nolint: golint
type USDCRef struct {
	*FiatTokenV2
	address common.Address
}

// Address is the contract address.
func (s USDCRef) Address() common.Address {
	return s.address
}

// NewUSDCRef creates a new tether token.
func NewUSDCRef(address common.Address, backend bind.ContractBackend) (*USDCRef, error) {
	usdcToken, err := NewFiatTokenV2(address, backend)
	if err != nil {
		return nil, err
	}
	return &USDCRef{
		FiatTokenV2: usdcToken,
		address:     address,
	}, nil
}

var _ vm.ContractRef = &USDCRef{}
