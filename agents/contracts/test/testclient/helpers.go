package testclient

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestClientRef is a test client reference
//
//nolint:golint
type TestClientRef struct {
	*TestClient
	address common.Address
}

// Address gets the address of the contract.
func (a TestClientRef) Address() common.Address {
	return a.address
}

// NewTestClientRef creates a new test client.
func NewTestClientRef(address common.Address, backend bind.ContractBackend) (*TestClientRef, error) {
	contract, err := NewTestClient(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create test client: %w", err)
	}

	return &TestClientRef{
		TestClient: contract,
		address:    address,
	}, nil
}

var _ vm.ContractRef = &TestClientRef{}
