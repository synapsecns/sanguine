package interchainappmock

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// InterchainAppMockRef is a reference to an interchain db.
type InterchainAppMockRef struct {
	*InterchainAppMock
	// address of the interchain client
	address common.Address
}

// Address is the contract address.
func (i *InterchainAppMockRef) Address() common.Address {
	return i.address
}

// NewInterchainAppMockRef creates a new interchain client with a contract ref.
func NewInterchainAppMockRef(address common.Address, backend bind.ContractBackend) (*InterchainAppMockRef, error) {
	instance, err := NewInterchainAppMock(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create instance of InterchainClient: %w", err)
	}
	return &InterchainAppMockRef{
		InterchainAppMock: instance,
		address:           address,
	}, nil
}

var _ vm.ContractRef = &InterchainAppMockRef{}
