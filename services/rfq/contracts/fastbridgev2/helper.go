package fastbridgev2

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// FastBridgeV2Ref is a bound fast bridge contract that returns the address of the contract.
//
//nolint:golint
type FastBridgeV2Ref struct {
	*FastBridgeV2
	address common.Address
}

// Address gets the ocntract address.
func (f *FastBridgeV2Ref) Address() common.Address {
	return f.address
}

// NewFastBridgeV2Ref creates a new fast bridge contract witha  ref.
func NewFastBridgeV2Ref(address common.Address, backend bind.ContractBackend) (*FastBridgeV2Ref, error) {
	fastBridge, err := NewFastBridgeV2(address, backend)
	if err != nil {
		return nil, err
	}

	return &FastBridgeV2Ref{
		FastBridgeV2: fastBridge,
		address:      address,
	}, nil
}

var _ vm.ContractRef = &FastBridgeV2Ref{}
