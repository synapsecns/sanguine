package mockmintburntoken

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// MockMintBurnTokenRef is a bound cctp contract that conforms to vm.ContractRef.
//
//nolint:golint
type MockMintBurnTokenRef struct {
	*MockMintBurnToken
	address common.Address
}

// Address is the contract address.
func (s MockMintBurnTokenRef) Address() common.Address {
	return s.address
}

// NewMockMintBurnTokenRef creates a new MockMintBurnTokenRef contract with a contract ref.
func NewMockMintBurnTokenRef(address common.Address, backend bind.ContractBackend) (*MockMintBurnTokenRef, error) {
	cctpContract, err := NewMockMintBurnToken(address, backend)
	if err != nil {
		return nil, err
	}

	return &MockMintBurnTokenRef{
		MockMintBurnToken: cctpContract,
		address:           address,
	}, nil
}

var _ vm.ContractRef = MockMintBurnTokenRef{}
