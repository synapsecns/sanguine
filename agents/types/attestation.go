package types

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

const (
	attestationOffsetRoot        = 0
	attestationOffsetDataHash    = 32
	attestationOffsetNonce       = 64
	attestationOffsetBlockNumber = 68
	attestationOffsetTimestamp   = 73
	attestationSize              = 78
)

// Attestation is the attestation interface.
type Attestation interface {
	// SnapshotRoot is the root of the Snapshot Merkle Tree.
	SnapshotRoot() [32]byte
	// DataHash is the agent root and SnapGasHash combined into a single hash.
	DataHash() [32]byte
	// Nonce is the attestation nonce.
	Nonce() uint32
	// BlockNumber is the block number when the attestation was created in Summit.
	BlockNumber() *big.Int
	// Timestamp is the timestamp when the attestation was created in Summit.
	Timestamp() *big.Int
	// SignAttestation signs the attestation
	SignAttestation(ctx context.Context, signer signer.Signer, valid bool) (signer.Signature, []byte, common.Hash, error)
}

type attestation struct {
	snapshotRoot [32]byte
	dataHash     [32]byte
	nonce        uint32
	blockNumber  *big.Int
	timestamp    *big.Int
}

// NewAttestation creates a new attestation.
func NewAttestation(snapshotRoot [32]byte, dataHash [32]byte, nonce uint32, blockNumber *big.Int, timestamp *big.Int) Attestation {
	return &attestation{
		snapshotRoot: snapshotRoot,
		dataHash:     dataHash,
		nonce:        nonce,
		blockNumber:  blockNumber,
		timestamp:    timestamp,
	}
}

func (a attestation) SnapshotRoot() [32]byte {
	return a.snapshotRoot
}

func (a attestation) DataHash() [32]byte {
	return a.dataHash
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

//nolint:dupl
func (a attestation) SignAttestation(ctx context.Context, signer signer.Signer, valid bool) (signer.Signature, []byte, common.Hash, error) {
	encodedAttestation, err := EncodeAttestation(a)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not encode attestation: %w", err)
	}

	var attestationSalt []byte
	if valid {
		attestationSalt = []byte("ATTESTATION_VALID_SALT")
	} else {
		attestationSalt = []byte("ATTESTATION_INVALID_SALT")
	}

	signature, hashedAttestation, err := sign(ctx, signer, encodedAttestation, attestationSalt)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not sign attestation: %w", err)
	}

	return signature, encodedAttestation, hashedAttestation, nil
}

// GetAttestationDataHash generates the data hash from the agent root and SnapGasHash.
func GetAttestationDataHash(agentRoot [32]byte, snapGasHash [32]byte) [32]byte {
	concatenatedBytes := append(agentRoot[:], snapGasHash[:]...)
	dataHash := crypto.Keccak256(concatenatedBytes)

	var dataHashB32 [32]byte
	copy(dataHashB32[:], dataHash)

	return dataHashB32
}
