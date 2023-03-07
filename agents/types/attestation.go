package types

import "math/big"

// Attestation is the attestation interface.
type Attestation interface {
	// Root is the snapshot merkle tree's root.
	Root() [32]byte
	// Height is the snapshot merkle tree's height.
	Height() uint8
	// Nonce is the attestation nonce.
	Nonce() uint32
	// BlockNumber is the block number of the attestation.
	BlockNumber() *big.Int
	// Timestamp is the timestamp of the attestation.
	Timestamp() *big.Int
}

type attestation struct {
	root        [32]byte
	height      uint8
	nonce       uint32
	blockNumber *big.Int
	timestamp   *big.Int
}

// NewAttestation creates a new attestation.
func NewAttestation(root [32]byte, height uint8, nonce uint32, blockNumber *big.Int, timestamp *big.Int) Attestation {
	return &attestation{
		root:        root,
		height:      height,
		nonce:       nonce,
		blockNumber: blockNumber,
		timestamp:   timestamp,
	}
}

func (a attestation) Root() [32]byte {
	return a.root
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
