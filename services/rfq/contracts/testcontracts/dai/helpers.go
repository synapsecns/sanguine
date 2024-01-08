package dai

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// DaiRef is a bound synfactory bridge contract that returns the address of the contract.
// nolint: golint
type DaiRef struct {
	*Dai
	address common.Address
}

// Address is the contract address.
func (s DaiRef) Address() common.Address {
	return s.address
}

// NewDaiRef creates a new dai token.
func NewDaiRef(address common.Address, backend bind.ContractBackend) (*DaiRef, error) {
	usdcToken, err := NewDai(address, backend)
	if err != nil {
		return nil, err
	}
	return &DaiRef{
		Dai:     usdcToken,
		address: address,
	}, nil
}

var _ vm.ContractRef = &DaiRef{}
