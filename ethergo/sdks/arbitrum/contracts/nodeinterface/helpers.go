package nodeinterface

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NodeInterfaceRef is a reference to a NodeInterface contract
type NodeInterfaceRef struct {
	*NodeInterface
	address common.Address
}

// Address gets the address of the contract.
func (n NodeInterfaceRef) Address() common.Address {
	return n.address
}

// NewNodeInterfaceRef creates a new NodeInterfaceRef bound to a contract.
func NewNodeInterfaceRef(address common.Address, backend bind.ContractBackend) (*NodeInterfaceRef, error) {
	nodeInterface, err := NewNodeInterface(address, backend)
	if err != nil {
		return nil, err
	}

	return &NodeInterfaceRef{
		NodeInterface: nodeInterface,
		address:       address,
	}, nil
}

type INodeInterface interface {
	INodeInterfaceCaller
	INodeInterfaceTransactor
	INodeInterfaceFilterer
}
