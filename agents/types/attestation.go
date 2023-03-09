package types

import "math/big"

const (
	attestationOffsetRoot        = 0
	attestationOffsetDepth       = 32
	attestationOffsetNonce       = 33
	attestationOffsetBlockNumber = 37
	attestationOffsetTimestamp   = 42
	attestationSize              = 47
)

// Attestation is the attestation interface.
type Attestation interface {
	// SnapshotRoot is the root of the Snapshot Merkle Tree.
	SnapshotRoot() [32]byte
	// Height is the height of the Snapshot Merkle Tree.
	Height() uint8
	// Nonce is the attestation nonce.
	Nonce() uint32
	// BlockNumber is the block number when the attestation was created in Summit.
	BlockNumber() *big.Int
	// Timestamp is the timestamp when the attestation was created in Summit.
	Timestamp() *big.Int
}

type attestation struct {
	snapshotRoot [32]byte
	height       uint8
	nonce        uint32
	blockNumber  *big.Int
	timestamp    *big.Int
}

// NewAttestation creates a new attestation.
func NewAttestation(snapshotRoot [32]byte, height uint8, nonce uint32, blockNumber *big.Int, timestamp *big.Int) Attestation {
	return &attestation{
		snapshotRoot: snapshotRoot,
		height:       height,
		nonce:        nonce,
		blockNumber:  blockNumber,
		timestamp:    timestamp,
	}
}

func (a attestation) SnapshotRoot() [32]byte {
	return a.snapshotRoot
}

func (a attestation) Height() uint8 {
	return a.height
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
