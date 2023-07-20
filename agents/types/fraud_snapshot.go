package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// FraudSnapshot is a snapshot type with additional metadata for fraud handling.
type FraudSnapshot struct {
	// Snapshot is the underlying snapshot.
	Snapshot Snapshot

	// Domain is the domain of the snapshot.
	Domain uint32
	// Agent is the agent that signed the snapshot.
	Agent common.Address
	// Payload is the snapshot payload.
	Payload []byte
	// Signature is the signature of the snapshot payload signed by the signer.
	Signature []byte
}

// NewFraudSnapshotFromPayload returns a new FraudSnapshot from a Snapshot payload and other metadata.
func NewFraudSnapshotFromPayload(snapshotPayload []byte, domain uint32, agent common.Address, snapSignature []byte) (*FraudSnapshot, error) {
	decodedSnapshot, err := DecodeSnapshot(snapshotPayload)
	if err != nil {
		return nil, fmt.Errorf("could not decode snapshot: %w", err)
	}

	return &FraudSnapshot{
		Snapshot:  decodedSnapshot,
		Domain:    domain,
		Agent:     agent,
		Payload:   snapshotPayload,
		Signature: snapSignature,
	}, nil
}
