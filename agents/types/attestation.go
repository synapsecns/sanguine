package types

import "math/big"

const (
	attestationOffsetRoot        = 0
	attestationOffsetAgentRoot   = 32
	attestationOffsetNonce       = 64
	attestationOffsetBlockNumber = 68
	attestationOffsetTimestamp   = 73
	attestationSize              = 78
)

// Attestation is the attestation interface.
type Attestation interface {
	// SnapshotRoot is the root of the Snapshot Merkle Tree.
	SnapshotRoot() [32]byte
	// AgentRoot is the root of the Agent Merkle Tree.
	AgentRoot() [32]byte
	// Nonce is the attestation nonce.
	Nonce() uint32
	// BlockNumber is the block number when the attestation was created in Summit.
	BlockNumber() *big.Int
	// Timestamp is the timestamp when the attestation was created in Summit.
	Timestamp() *big.Int
}

type attestation struct {
	snapshotRoot [32]byte
	agentRoot    [32]byte
	nonce        uint32
	blockNumber  *big.Int
	timestamp    *big.Int
}

// NewAttestation creates a new attestation.
func NewAttestation(snapshotRoot [32]byte, agentRoot [32]byte, nonce uint32, blockNumber *big.Int, timestamp *big.Int) Attestation {
	return &attestation{
		snapshotRoot: snapshotRoot,
		agentRoot:    agentRoot,
		nonce:        nonce,
		blockNumber:  blockNumber,
		timestamp:    timestamp,
	}
}

func (a attestation) SnapshotRoot() [32]byte {
	return a.snapshotRoot
}

func (a attestation) AgentRoot() [32]byte {
	return a.agentRoot
}

func (a attestation) Nonce() uint32 {
	return a.nonce
}

func (a attestation) BlockNumber() *big.Int {
	return a.blockNumber
}

func (a attestation) Timestamp() *big.Int {
	return a.timestamp
}
