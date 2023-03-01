package pingpongclient

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// PingPongClientRef is a ping pong test client reference
//
//nolint:golint
type PingPongClientRef struct {
	*PingPongClient
	address common.Address
}

// Address gets the address of the contract.
func (a PingPongClientRef) Address() common.Address {
	return a.address
}

// NewPingPongClientRef creates a new test client.
func NewPingPongClientRef(address common.Address, backend bind.ContractBackend) (*PingPongClientRef, error) {
	contract, err := NewPingPongClient(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create ping pong test client: %w", err)
	}

	return &PingPongClientRef{
		PingPongClient: contract,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &PingPongClientRef{}
