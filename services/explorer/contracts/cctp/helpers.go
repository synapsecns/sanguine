package cctp

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// CCTPRef is a reference to a deployed CCTP contract.
//
//nolint:golint
type CCTPRef struct {
	*SynapseCCTP
	address common.Address
}

// MessageTransmitterRef is a reference to a deployed message transmitter contract.
//
//nolint:golint
type MessageTransmitterRef struct {
	*MessageTransmitter
	address common.Address
}

// Address is the contract address.
func (s CCTPRef) Address() common.Address {
	return s.address
}

// Address is the contract address.
func (s MessageTransmitterRef) Address() common.Address {
	return s.address
}

// NewCCTPRef creates a new CCTPRef instance.
//
//nolint:golint
func NewCCTPRef(address common.Address, backend bind.ContractBackend) (*CCTPRef, error) {
	synapseCCTP, err := NewSynapseCCTP(address, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPRef{
		SynapseCCTP: synapseCCTP,
		address:     address,
	}, nil
}

// NewMessageTransmitterRef creates a new MessageTransmitterRef instance.
func NewMessageTransmitterRef(address common.Address, backend bind.ContractBackend) (*MessageTransmitterRef, error) {
	messageTransmitter, err := NewMessageTransmitter(address, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterRef{
		MessageTransmitter: messageTransmitter,
		address:            address,
	}, nil
}

var _ vm.ContractRef = &CCTPRef{}
var _ vm.ContractRef = &MessageTransmitterRef{}
