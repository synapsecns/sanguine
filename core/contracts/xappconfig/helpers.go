package xappconfig

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// XAppConfigRef is a bound xappconfig contract that returns the address of the contract.
//nolint: golint
type XAppConfigRef struct {
	*XAppConfig
	address common.Address
}

// Address is the contract address.
func (s XAppConfigRef) Address() common.Address {
	return s.address
}

// NewXAppConfigRef creates a new xappconfig contract with a contract ref.
func NewXAppConfigRef(address common.Address, backend bind.ContractBackend) (*XAppConfigRef, error) {
	contract, err := NewXAppConfig(address, backend)
	if err != nil {
		return nil, err
	}
	return &XAppConfigRef{
		XAppConfig: contract,
		address:    address,
	}, nil
}

var _ vm.ContractRef = XAppConfigRef{}
