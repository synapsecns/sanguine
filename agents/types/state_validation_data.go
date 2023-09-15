package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// StateValidationData provides data necessary for processing states in a snapshot,
// i.e. verifyState() and submitStateReport() contract calls.
type StateValidationData interface {
	// Agent returns the agent that submitted the snapshot or attestation.
	Agent() common.Address
	// AgentDomain returns the domain of the agent that submitted the snapshot or attestation.
	AgentDomain() uint32
	// SnapshotPayload returns the snapshot payload.
	SnapshotPayload() []byte
	// SnapshotSignature returns the snapshot signature, if it exists.
	SnapshotSignature() []byte
	// AttestationPayload returns the attestation payload, if it exists.
	AttestationPayload() []byte
	// AttestationSignature returns the attestation signature, if it exists.
	AttestationSignature() []byte
}

// HasAttestation indicates whether the data corresponds to an attestation (or a snapshot, if false).
func HasAttestation(data StateValidationData) bool {
	return data.AttestationPayload() != nil
}

// AttestationWithMetadata is an attestation that was submitted by a Notary and was deemed fraudulent.
type AttestationWithMetadata struct {
	// Attestation is the underlying attestation.
	Attestation          Attestation
	agentDomain          uint32
	notary               common.Address
	attestationPayload   []byte
	attestationSignature []byte
	snapshotPayload      []byte
}

// NewAttestationWithMetadata creates a new AttestationWithMetadata from the attestation payload, domain, notary and attestation signature.
func NewAttestationWithMetadata(attestationPayload []byte, agentDomain uint32, notary common.Address, attSignature []byte) (*AttestationWithMetadata, error) {
	decodedAttestation, err := DecodeAttestation(attestationPayload)
	if err != nil {
		return nil, err
	}

	return &AttestationWithMetadata{
		Attestation:          decodedAttestation,
		agentDomain:          agentDomain,
		notary:               notary,
		attestationPayload:   attestationPayload,
		attestationSignature: attSignature,
	}, nil
}

// SetSnapshotPayload sets the snapshot payload.
func (f AttestationWithMetadata) SetSnapshotPayload(snapshotPayload []byte) {
	f.snapshotPayload = snapshotPayload
}

// Agent returns the agent that submitted the attestation.
func (f AttestationWithMetadata) Agent() common.Address {
	return f.notary
}

// AgentDomain returns the domain of the agent that submitted the attestation.
func (f AttestationWithMetadata) AgentDomain() uint32 {
	return f.agentDomain
}

// SnapshotPayload returns the snapshot payload.
func (f AttestationWithMetadata) SnapshotPayload() []byte {
	return f.snapshotPayload
}

// SnapshotSignature returns the nil, since the data corresponds to an attestation.
func (f AttestationWithMetadata) SnapshotSignature() []byte {
	return nil
}

// AttestationPayload returns the attestation payload.
func (f AttestationWithMetadata) AttestationPayload() []byte {
	return f.attestationPayload
}

// AttestationSignature returns the attestation signature.
func (f AttestationWithMetadata) AttestationSignature() []byte {
	return f.attestationSignature
}

// SnapshotWithMetadata is a snapshot type with additional metadata for fraud handling.
type SnapshotWithMetadata struct {
	// Snapshot is the underlying snapshot.
	Snapshot          Snapshot
	agentDomain       uint32
	agent             common.Address
	snapshotPayload   []byte
	snapshotSignature []byte
}

// NewSnapshotWithMetadata returns a new SnapshotWithMetadata from a Snapshot payload and other metadata.
func NewSnapshotWithMetadata(snapshotPayload []byte, agentDomain uint32, agent common.Address, snapshotSignature []byte) (*SnapshotWithMetadata, error) {
	decodedSnapshot, err := DecodeSnapshot(snapshotPayload)
	if err != nil {
		return nil, fmt.Errorf("could not decode snapshot: %w", err)
	}

	return &SnapshotWithMetadata{
		Snapshot:          decodedSnapshot,
		agentDomain:       agentDomain,
		agent:             agent,
		snapshotPayload:   snapshotPayload,
		snapshotSignature: snapshotSignature,
	}, nil
}

// Agent returns the agent that submitted the snapshot.
func (f SnapshotWithMetadata) Agent() common.Address {
	return f.agent
}

// AgentDomain returns the domain of the agent that submitted the snapshot.
func (f SnapshotWithMetadata) AgentDomain() uint32 {
	return f.agentDomain
}

// SnapshotPayload returns the snapshot payload.
func (f SnapshotWithMetadata) SnapshotPayload() []byte {
	return f.snapshotPayload
}

// SnapshotSignature returns the snapshot signature.
func (f SnapshotWithMetadata) SnapshotSignature() []byte {
	return f.snapshotSignature
}

// AttestationPayload returns nil, since the data corresponds to a snapshot.
func (f SnapshotWithMetadata) AttestationPayload() []byte {
	return nil
}

// AttestationSignature returns nil, since the data corresponds to a snapshot.
func (f SnapshotWithMetadata) AttestationSignature() []byte {
	return nil
}
