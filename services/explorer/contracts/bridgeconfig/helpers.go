package bridgeconfig

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// BridgeConfigRef is a bound synapse bridge config v2 contract that returns the address of that contract
//
//nolint:golint
type BridgeConfigRef struct {
	*BridgeConfigV3
	address common.Address
}

// Address is the contract address.
func (s BridgeConfigRef) Address() common.Address {
	return s.address
}

// NewBridgeConfigRef gets a bound synapse bridge config contract that returns the address of the contract
//
//nolint:golint
func NewBridgeConfigRef(address common.Address, backend bind.ContractBackend) (*BridgeConfigRef, error) {
	bridgeConfigSwap, err := NewBridgeConfigV3(address, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeConfigRef{
		BridgeConfigV3: bridgeConfigSwap,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &BridgeConfigRef{}
