package bridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// BridgeRef is a bound synapse bridge config v2 contract that returns the address of that contract
//
//nolint:golint
type BridgeRef struct {
	*SynapseBridge
	address common.Address
}

// Address is the contract address.
func (s BridgeRef) Address() common.Address {
	return s.address
}

// NewBridgeRef gets a bound synapse bridge config contract that returns the address of the contract
//
//nolint:golint
func NewBridgeRef(address common.Address, backend bind.ContractBackend) (*BridgeRef, error) {
	bridge, err := NewSynapseBridge(address, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeRef{
		SynapseBridge: bridge,
		address:       address,
	}, nil
}

var _ vm.ContractRef = &BridgeRef{}
