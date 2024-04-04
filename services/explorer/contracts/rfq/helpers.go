package rfq

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// RFQRef is a reference to a deployed RFQ contract.
//
//nolint:golint
type RFQRef struct {
	*FastBridge
	address common.Address
}

// MessageTransmitterRef is a reference to a deployed message transmitter contract.
//
//nolint:golint

// Address is the contract address.
func (s RFQRef) Address() common.Address {
	return s.address
}

// NewRFQRef creates a new RFQRef instance.
//
//nolint:golint
func NewRFQRef(address common.Address, backend bind.ContractBackend) (*RFQRef, error) {
	fastBridge, err := NewFastBridge(address, backend)
	if err != nil {
		return nil, err
	}
	return &RFQRef{
		FastBridge: fastBridge,
		address:    address,
	}, nil
}

var _ vm.ContractRef = &RFQRef{}
