package attestationharness

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// AttestationHarnessRef is a attestation harness reference.
//
//nolint:golint
type AttestationHarnessRef struct {
	*AttestationHarness
	address common.Address
}

// Address gets the address of the contract.
func (s AttestationHarnessRef) Address() common.Address {
	return s.address
}

// NewAttestationHarnessRef creates a new attestation harness.
func NewAttestationHarnessRef(address common.Address, backend bind.ContractBackend) (*AttestationHarnessRef, error) {
	contract, err := NewAttestationHarness(address, backend)
	if err != nil {
		return nil, err
	}

	return &AttestationHarnessRef{
		AttestationHarness: contract,
		address:            address,
	}, nil
}

var _ vm.ContractRef = &AttestationHarnessRef{}
