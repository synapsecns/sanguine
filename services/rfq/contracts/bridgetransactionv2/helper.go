package bridgetransactionv2

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// BridgeTransactionV2Ref is a bound fast bridge contract that returns the address of the contract.
//
//nolint:golint
type BridgeTransactionV2Ref struct {
	*BridgeTransactionV2Harness
	address common.Address
}

// Address gets the ocntract address.
func (f *BridgeTransactionV2Ref) Address() common.Address {
	return f.address
}

// NewBridgeTransactionV2Ref creates a new fast bridge mock contract with a ref.
func NewBridgeTransactionV2Ref(address common.Address, backend bind.ContractBackend) (*BridgeTransactionV2Ref, error) {
	bridgetransactionv2, err := NewBridgeTransactionV2Harness(address, backend)
	if err != nil {
		return nil, err
	}

	return &BridgeTransactionV2Ref{
		BridgeTransactionV2Harness: bridgetransactionv2,
		address:                    address,
	}, nil
}

var _ vm.ContractRef = &BridgeTransactionV2Ref{}
