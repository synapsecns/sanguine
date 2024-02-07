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
func (s *SnapshotWithMetadata) Agent() common.Address {
	return s.agent
}

// AgentDomain returns the domain of the agent that submitted the snapshot.
func (s *SnapshotWithMetadata) AgentDomain() uint32 {
	return s.agentDomain
}

// SnapshotPayload returns the snapshot payload.
func (s *SnapshotWithMetadata) SnapshotPayload() []byte {
	return s.snapshotPayload
}

// SnapshotSignature returns the snapshot signature.
func (s *SnapshotWithMetadata) SnapshotSignature() []byte {
	return s.snapshotSignature
}

// AttestationPayload returns nil, since the data corresponds to a snapshot.
func (s *SnapshotWithMetadata) AttestationPayload() []byte {
	return nil
}

// AttestationSignature returns nil, since the data corresponds to a snapshot.
func (s *SnapshotWithMetadata) AttestationSignature() []byte {
	return nil
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
func (a *AttestationWithMetadata) SetSnapshotPayload(snapshotPayload []byte) {
	a.snapshotPayload = snapshotPayload
}

// Agent returns the agent that submitted the attestation.
func (a *AttestationWithMetadata) Agent() common.Address {
	return a.notary
}

// AgentDomain returns the domain of the agent that submitted the attestation.
func (a *AttestationWithMetadata) AgentDomain() uint32 {
	return a.agentDomain
}

// SnapshotPayload returns the snapshot payload.
func (a *AttestationWithMetadata) SnapshotPayload() []byte {
	return a.snapshotPayload
}

// SnapshotSignature returns the nil, since the data corresponds to an attestation.
func (a *AttestationWithMetadata) SnapshotSignature() []byte {
	return nil
}

// AttestationPayload returns the attestation payload.
func (a *AttestationWithMetadata) AttestationPayload() []byte {
	return a.attestationPayload
}

// AttestationSignature returns the attestation signature.
func (a *AttestationWithMetadata) AttestationSignature() []byte {
	return a.attestationSignature
}
