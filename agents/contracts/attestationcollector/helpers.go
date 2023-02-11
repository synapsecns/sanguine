package attestationcollector

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// AttestationCollectorRef is a bound attestatoin collector contract that returns the address of the attestation collector contract.
//
//nolint:golint
type AttestationCollectorRef struct {
	*AttestationCollector
	address common.Address
}

// Address gets the address of the attestation contract.
func (a AttestationCollectorRef) Address() common.Address {
	return a.address
}

// NewAttestationCollectorRef creates an attestation contract with a contract ref.
func NewAttestationCollectorRef(address common.Address, backend bind.ContractBackend) (*AttestationCollectorRef, error) {
	attestationContract, err := NewAttestationCollector(address, backend)
	if err != nil {
		return nil, err
	}

	return &AttestationCollectorRef{
		AttestationCollector: attestationContract,
		address:              address,
	}, nil
}

var _ vm.ContractRef = AttestationCollectorRef{}
