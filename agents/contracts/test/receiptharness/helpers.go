package receiptharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// ReceiptHarnessRef is a receipt harness reference
//
//nolint:golint
type ReceiptHarnessRef struct {
	*ReceiptHarness
	address common.Address
}

// Address gets the address of the contract.
func (h ReceiptHarnessRef) Address() common.Address {
	return h.address
}

// NewReceiptHarnessRef creates a new receipt harness.
func NewReceiptHarnessRef(address common.Address, backend bind.ContractBackend) (*ReceiptHarnessRef, error) {
	contract, err := NewReceiptHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create receipt harness: %w", err)
	}

	return &ReceiptHarnessRef{
		ReceiptHarness: contract,
		address:        address,
	}, nil
}

var _ vm.ContractRef = &ReceiptHarnessRef{}
