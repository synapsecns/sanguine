package attestationharness

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// AttestationHarnessRef is an attestation harness reference
//
//nolint:golint
type AttestationHarnessRef struct {
	*AttestationHarness
	address common.Address
}

// Address gets the address of the contract.
func (a AttestationHarnessRef) Address() common.Address {
	return a.address
}

// NewAttestationHarnessRef creates a new attestation harness.
func NewAttestationHarnessRef(address common.Address, backend bind.ContractBackend) (*AttestationHarnessRef, error) {
	contract, err := NewAttestationHarness(address, backend)
	if err != nil {
		return nil, fmt.Errorf("could not create attestation harness: %w", err)
	}

	return &AttestationHarnessRef{
		AttestationHarness: contract,
		address:            address,
	}, nil
}

var _ vm.ContractRef = &AttestationHarnessRef{}
