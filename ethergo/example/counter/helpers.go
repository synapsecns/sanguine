package counter

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// CounterRef is a bound counter.
//
//nolint:golint, revive
type CounterRef struct {
	*Counter
	address common.Address
}

// Address gets the address of the contract.
func (a CounterRef) Address() common.Address {
	return a.address
}

// NewCounterRef creates a new test client.
func NewCounterRef(address common.Address, backend bind.ContractBackend) (*CounterRef, error) {
	contract, err := NewCounter(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create ping pong test client: %w", err)
	}

	return &CounterRef{
		Counter: contract,
		address: address,
	}, nil
}

var _ vm.ContractRef = &CounterRef{}
