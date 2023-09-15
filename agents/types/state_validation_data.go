package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// StateValidationData provides data necessary for processing states in a snapshot,
// i.e. verifyState() and submitStateReport() contract calls.
type StateValidationData interface {
	SnapshotPayload() []byte
	SnapshotSignature() []byte
	AttestationPayload() []byte
	AttestationSignature() []byte
	Agent() common.Address
	AgentDomain() uint32
}

func HasAttestation(data StateValidationData) bool {
	return data.AttestationPayload() != nil
}

// AttestationWithMetadata is an attestation that was submitted by a Notary and was deemed fraudulent.
type AttestationWithMetadata struct {
	// Attestation is the underlying attestation.
	Attestation Attestation
	// AgentDomain is the domain of the Notary who signed the attestation.
	agentDomain uint32
	// Notary is the Notary who signed and submitted the attestation.
	notary common.Address
	// Payload is the attestation payload.
	payload []byte
	// Signature is the signature of the attestation payload signed by the Notary.
	signature []byte
	// SnapshotPayload is the snapshot payload for the snapshot that the attestation corresponds to.
	snapshotPayload []byte
}

// NewAttestationWithMetadata creates a new AttestationWithMetadata from the attestation payload, domain, notary and attestation signature.
func NewAttestationWithMetadata(attestationPayload []byte, agentDomain uint32, notary common.Address, attSignature []byte) (*AttestationWithMetadata, error) {
	decodedAttestation, err := DecodeAttestation(attestationPayload)
	if err != nil {
		return nil, err
	}

	return &AttestationWithMetadata{
		Attestation: decodedAttestation,
		agentDomain: agentDomain,
		notary:      notary,
		payload:     attestationPayload,
		signature:   attSignature,
	}, nil
}

func (f AttestationWithMetadata) Agent() common.Address {
	return f.notary
}

func (f AttestationWithMetadata) AgentDomain() uint32 {
	return f.agentDomain
}

func (f AttestationWithMetadata) SetSnapshotPayload(snapshotPayload []byte) {
	f.snapshotPayload = snapshotPayload
}

func (f AttestationWithMetadata) SnapshotPayload() []byte {
	return f.snapshotPayload
}

func (f AttestationWithMetadata) SnapshotSignature() []byte {
	return nil
}

func (f AttestationWithMetadata) AttestationPayload() []byte {
	return f.payload
}

func (f AttestationWithMetadata) AttestationSignature() []byte {
	return f.signature
}

// SnapshotWithMetadata is a snapshot type with additional metadata for fraud handling.
type SnapshotWithMetadata struct {
	// Snapshot is the underlying snapshot.
	Snapshot Snapshot
	// AgentDomain is the domain of the agent that submitted the snapshot.
	agentDomain uint32
	// Agent is the agent that signed the snapshot.
	agent common.Address
	// Payload is the snapshot payload.
	payload []byte
	// Signature is the signature of the snapshot payload signed by the signer.
	signature []byte
}

// NewSnapshotWithMetadata returns a new SnapshotWithMetadata from a Snapshot payload and other metadata.
func NewSnapshotWithMetadata(snapshotPayload []byte, agentDomain uint32, agent common.Address, snapSignature []byte) (*SnapshotWithMetadata, error) {
	decodedSnapshot, err := DecodeSnapshot(snapshotPayload)
	if err != nil {
		return nil, fmt.Errorf("could not decode snapshot: %w", err)
	}

	return &SnapshotWithMetadata{
		Snapshot:    decodedSnapshot,
		agentDomain: agentDomain,
		agent:       agent,
		payload:     snapshotPayload,
		signature:   snapSignature,
	}, nil
}

func (f SnapshotWithMetadata) Agent() common.Address {
	return f.agent
}

func (f SnapshotWithMetadata) AgentDomain() uint32 {
	return f.agentDomain
}

func (f SnapshotWithMetadata) SnapshotPayload() []byte {
	return f.payload
}

func (f SnapshotWithMetadata) SnapshotSignature() []byte {
	return f.signature
}

func (f SnapshotWithMetadata) AttestationPayload() []byte {
	return nil
}

func (f SnapshotWithMetadata) AttestationSignature() []byte {
	return nil
}
