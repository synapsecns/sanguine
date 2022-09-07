package bridgeconfig

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// BridgeConfigSwapRef is a bound synapse bridge config v2 contract that returns the address of that contract
//nolint: golint
type BridgeConfigSwapRef struct {
	*BridgeConfigSwap
	address common.Address
}

// Address is the contract address.
func (s BridgeConfigSwapRef) Address() common.Address {
	return s.address
}

// NewBridgeConfigSwapRef gets a bound synapse bridge config contract that returns the address of the contract
//nolint: golint
func NewBridgeConfigSwapRef(address common.Address, backend bind.ContractBackend) (*BridgeConfigSwapRef, error) {
	bridgeConfigSwap, err := NewBridgeConfigSwap(address, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigSwapRef{
		BridgeConfigSwap: bridgeConfigSwap,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &NewBridgeConfigSwapRef{}
