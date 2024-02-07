package testcctp

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// TestCCTPRef is a reference to a deployed test CCTP contract.
//
//nolint:golint
type TestCCTPRef struct {
	*TestSynapseCCTP
	address common.Address
}

// Address is the contract address.
func (s TestCCTPRef) Address() common.Address {
	return s.address
}

// NewTestCCTPRef creates a new TestCCTPRef instance.
//
//nolint:golint
func NewTestCCTPRef(address common.Address, backend bind.ContractBackend) (*TestCCTPRef, error) {
	synapseCCTP, err := NewTestSynapseCCTP(address, backend)
	if err != nil {
		return nil, err
	}
	return &TestCCTPRef{
		TestSynapseCCTP: synapseCCTP,
		address:         address,
	}, nil
}

var _ vm.ContractRef = &TestCCTPRef{}
