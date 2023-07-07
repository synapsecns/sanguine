package arbgasinfo

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ArbGasInfoRef is a reference to a ArbGasInfo contract.
//
// nolint: golint
type ArbGasInfoRef struct {
	*ArbGasInfo
	address common.Address
}

// Address gets the address of the contract.
func (a ArbGasInfoRef) Address() common.Address {
	return a.address
}

// NewArbGasInfoRef creates a new ArbGasInfoRef bound to a contract.
func NewArbGasInfoRef(address common.Address, backend bind.ContractBackend) (IArbGasInfo, error) {
	gasInfo, err := NewArbGasInfo(address, backend)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ArbGasInfo contract: %w", err)
	}

	return &ArbGasInfoRef{
		ArbGasInfo: gasInfo,
		address:    address,
	}, nil
}

// IArbGasInfo is a thin wrapper around ArbGasInfoCaller that allows interfacing with the contract.
type IArbGasInfo interface {
	IArbGasInfoCaller
}
