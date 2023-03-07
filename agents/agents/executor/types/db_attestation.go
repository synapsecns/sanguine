package types

import "github.com/ethereum/go-ethereum/common"

// DBAttestation is the executor type for interacting with the database representation of an attestation.
type DBAttestation struct {
	// SnapshotRoot is the snapshot root of the attestation.
	SnapshotRoot *common.Hash
	// ChainID is the chain ID of the chain that the attestation is for.
	ChainID *uint32
	// Destination is the destination chain id of the attestation.
	Destination *uint32
	// Height is the height of the Snapshot Merkle Tree
	Height *uint8
	// OriginBlockNumber is the block number that the state was taken from on the origin.
	OriginBlockNumber *uint64
	// OriginTimestamp is the timestamp of the block that the state was taken from on the origin.
	OriginTimestamp *uint64
	// NotValid is if the attestation is not valid.
	NotValid *bool
}
