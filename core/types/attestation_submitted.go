package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// AttestationSubmitted is the type emitted by
// the AttestationCollector when an attestation is submitted.
type AttestationSubmitted interface {
	// Notary gets the notary of the attestation.
	Notary() common.Hash
	// Attestation gets the raw bytes of the attestation.
	Attestation() []byte
}

type attestationSubmitted struct {
	notary      common.Hash
	attestation []byte
}

// NewAttestationSubmitted creates a new attestation submitted type.
func NewAttestationSubmitted(notary common.Hash, attestation []byte) AttestationSubmitted {
	return attestationSubmitted{
		notary:      notary,
		attestation: attestation,
	}
}

func (a attestationSubmitted) Notary() common.Hash {
	return a.notary
}

func (a attestationSubmitted) Attestation() []byte {
	return a.attestation
}
