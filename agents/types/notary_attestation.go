package types

import (
	"fmt"
	"math/big"
)

// NotaryAttestation is the notary attestation interface.
type NotaryAttestation interface {
	// AttPayload is raw bytes of the attestation payload.
	AttPayload() []byte
	// AgentRoot is the agent root.
	AgentRoot() [32]byte
	// SnapGas is the snap gas.
	SnapGas() []*big.Int
	// Attestation is the decoded attestation from the AttPayload
	Attestation() Attestation
}

type notaryAttestation struct {
	attPayload  []byte
	agentRoot   [32]byte
	snapGas     []*big.Int
	attestation Attestation
}

// NewNotaryAttestation creates a new notary attestation.
func NewNotaryAttestation(attPayload []byte, agentRoot [32]byte, snapGas []*big.Int) (NotaryAttestation, error) {
	attestation, err := DecodeAttestation(attPayload)
	if err != nil {
		return nil, fmt.Errorf("could not decode attestation: %w", err)
	}

	return &notaryAttestation{
		attPayload:  attPayload,
		agentRoot:   agentRoot,
		snapGas:     snapGas,
		attestation: attestation,
	}, nil
}

func (a notaryAttestation) AttPayload() []byte {
	return a.attPayload
}

func (a notaryAttestation) AgentRoot() [32]byte {
	return a.agentRoot
}

func (a notaryAttestation) SnapGas() []*big.Int {
	return a.snapGas
}

func (a notaryAttestation) Attestation() Attestation {
	return a.attestation
}
