package types

import "github.com/ethereum/go-ethereum/common"

// FraudAttestation is an attestation that was submitted by a Notary and was deemed fraudulent.
type FraudAttestation struct {
	// Attestation is the underlying attestation.
	Attestation Attestation

	// Domain is the domain of the Notary who signed the attestation.
	Domain uint32
	// Notary is the Notary who signed and submitted the attestation.
	Notary common.Address
	// Payload is the attestation payload.
	Payload []byte
	// Signature is the signature of the attestation payload signed by the Notary.
	Signature []byte
}

// NewFraudAttestationFromPayload creates a new FraudAttestation from the attestation payload, domain, notary and attestation signature.
func NewFraudAttestationFromPayload(attestationPayload []byte, domain uint32, notary common.Address, attSignature []byte) (*FraudAttestation, error) {
	decodedAttestation, err := DecodeAttestation(attestationPayload)
	if err != nil {
		return nil, err
	}

	return &FraudAttestation{
		Attestation: decodedAttestation,
		Domain:      domain,
		Notary:      notary,
		Payload:     attestationPayload,
		Signature:   attSignature,
	}, nil
}
