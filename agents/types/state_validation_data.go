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

// FraudAttestation is an attestation that was submitted by a Notary and was deemed fraudulent.
type FraudAttestation struct {
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

// NewFraudAttestationFromPayload creates a new FraudAttestation from the attestation payload, domain, notary and attestation signature.
func NewFraudAttestationFromPayload(attestationPayload []byte, agentDomain uint32, notary common.Address, attSignature []byte) (*FraudAttestation, error) {
	decodedAttestation, err := DecodeAttestation(attestationPayload)
	if err != nil {
		return nil, err
	}

	return &FraudAttestation{
		Attestation: decodedAttestation,
		agentDomain: agentDomain,
		notary:      notary,
		payload:     attestationPayload,
		signature:   attSignature,
	}, nil
}

func (f FraudAttestation) Agent() common.Address {
	return f.notary
}

func (f FraudAttestation) AgentDomain() uint32 {
	return f.agentDomain
}

func (f FraudAttestation) SetSnapshotPayload(snapshotPayload []byte) {
	f.snapshotPayload = snapshotPayload
}

func (f FraudAttestation) SnapshotPayload() []byte {
	return f.snapshotPayload
}

func (f FraudAttestation) SnapshotSignature() []byte {
	return nil
}

func (f FraudAttestation) AttestationPayload() []byte {
	return f.payload
}

func (f FraudAttestation) AttestationSignature() []byte {
	return f.signature
}

// FraudSnapshot is a snapshot type with additional metadata for fraud handling.
type FraudSnapshot struct {
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

// NewFraudSnapshotFromPayload returns a new FraudSnapshot from a Snapshot payload and other metadata.
func NewFraudSnapshotFromPayload(snapshotPayload []byte, agentDomain uint32, agent common.Address, snapSignature []byte) (*FraudSnapshot, error) {
	decodedSnapshot, err := DecodeSnapshot(snapshotPayload)
	if err != nil {
		return nil, fmt.Errorf("could not decode snapshot: %w", err)
	}

	return &FraudSnapshot{
		Snapshot:    decodedSnapshot,
		agentDomain: agentDomain,
		agent:       agent,
		payload:     snapshotPayload,
		signature:   snapSignature,
	}, nil
}

func (f FraudSnapshot) Agent() common.Address {
	return f.agent
}

func (f FraudSnapshot) AgentDomain() uint32 {
	return f.agentDomain
}

func (f FraudSnapshot) SnapshotPayload() []byte {
	return f.payload
}

func (f FraudSnapshot) SnapshotSignature() []byte {
	return f.signature
}

func (f FraudSnapshot) AttestationPayload() []byte {
	return nil
}

func (f FraudSnapshot) AttestationSignature() []byte {
	return nil
}
