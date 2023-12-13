package fastbridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// FastBridgeRef is a bound fast bridge contract that returns the address of the contract.
//
//nolint:golint
type FastBridgeRef struct {
	*FastBridge
	address common.Address
}

// Address gets the ocntract address
func (f *FastBridgeRef) Address() common.Address {
	return f.address
}

// NewFastBridgeRef creates a new fast bridge contract witha  ref.
func NewFastBridgeRef(address common.Address, backend bind.ContractBackend) (*FastBridgeRef, error) {
	fastBridge, err := NewFastBridge(address, backend)
	if err != nil {
		return nil, err
	}

	return &FastBridgeRef{
		FastBridge: fastBridge,
		address:    address,
	}, nil
}

var _ vm.ContractRef = &FastBridgeRef{}
