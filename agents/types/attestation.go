package types

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"math/big"
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
	SignAttestation(ctx context.Context, signer signer.Signer) (signer.Signature, []byte, common.Hash, error)
	// GetAttestationHash gets the hash that is signed when signing an Attestation
	GetAttestationHash(ctx context.Context) ([]byte, common.Hash, error)
	// RecoverSignerAddress recovers the signer address
	RecoverSignerAddress(ctx context.Context, attestationSignature Signature) (common.Address, error)
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

func (a attestation) SignAttestation(ctx context.Context, signer signer.Signer) (signer.Signature, []byte, common.Hash, error) {
	encodedAttestation, hashedAttestation, err := a.GetAttestationHash(ctx)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not encode attestation and get its hash: %w", err)
	}

	signature, err := signer.SignMessage(ctx, core.BytesToSlice(hashedAttestation), false)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not sign attestation: %w", err)
	}
	return signature, encodedAttestation, hashedAttestation, nil
}

func (a attestation) GetAttestationHash(ctx context.Context) ([]byte, common.Hash, error) {
	encodedAttestation, err := EncodeAttestation(a)
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("could not encode attestation: %w", err)
	}

	attestationSalt := crypto.Keccak256Hash([]byte("ATTESTATION_VALID_SALT"))

	hashedEncodedAttestation := crypto.Keccak256Hash(encodedAttestation).Bytes()
	toSign := append(attestationSalt.Bytes(), hashedEncodedAttestation...)

	hashedAttestation, err := HashRawBytes(toSign)
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("could not hash attestation: %w", err)
	}

	return encodedAttestation, hashedAttestation, nil
}

func (a attestation) RecoverSignerAddress(ctx context.Context, attestationSignature Signature) (common.Address, error) {
	_, hashedAttestation, err := a.GetAttestationHash(ctx)
	if err != nil {
		return common.Address{}, fmt.Errorf("could not encode attestation and get its hash: %w", err)
	}

	encodedSignature, err := EncodeSignature(attestationSignature)
	if err != nil {
		return common.Address{}, fmt.Errorf("could not encode attestation signature: %w", err)
	}

	encodedSignature[64] -= 27
	sigPublicKey, err := crypto.Ecrecover(hashedAttestation.Bytes(), encodedSignature)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to Ecrecover attestation signature: %w", err)
	}

	unmarshalledPubKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to UnmarshalPubkey attestation pub key: %w", err)
	}
	if unmarshalledPubKey == nil {
		return common.Address{}, fmt.Errorf("failed to UnmarshalPubkey attestation pub key with no error but returned nil")
	}

	unmarshalledPubKeyAsAddress := crypto.PubkeyToAddress(*unmarshalledPubKey)

	return unmarshalledPubKeyAsAddress, nil
}
