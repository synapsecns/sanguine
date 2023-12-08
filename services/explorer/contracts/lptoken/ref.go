package lptoken

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// LPTokenRef is a reference to a deployed LPToken contract, used primarily for testing.
type LPTokenRef struct {
	*LPToken
	address common.Address
}

// NewLPTokenRef creates a LPTokenContract ref.
func NewLPTokenRef(address common.Address, backend bind.ContractBackend) (*LPTokenRef, error) {
	lpToken, err := NewLPToken(address, backend)
	if err != nil {
		return nil, err
	}
	return &LPTokenRef{
		LPToken: lpToken,
		address: address,
	}, nil
}

// Address gets the contract address.
func (s LPTokenRef) Address() common.Address {
	return s.address
}

var _ vm.ContractRef = &LPTokenRef{}
