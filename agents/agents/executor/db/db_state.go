package db

import "encoding/json"

// DBState is the executor type for interacting with the database representation of a state.
//
//nolint:golint,revive
type DBState struct {
	// SnapshotRoot is the snapshot root.
	SnapshotRoot *string
	// Root is the origin Merkle tree's root.
	Root *string
	// ChainID is the origin chain id.
	ChainID *uint32
	// Nonce is the origin Merkle tree's nonce.
	Nonce *uint32
	// OriginBlockNumber is the block number that the state was taken from on the origin.
	OriginBlockNumber *uint64
	// OriginTimestamp is the timestamp of the block that the state was taken from on the origin.
	OriginTimestamp *uint64
	// Proof is the Snapshot Merkle Tree proof for the state.
	Proof *json.RawMessage
	// StateIndex is the index of the state in the Snapshot.
	StateIndex *uint32
	// BlockNumber is the block number the state was received at on Summit.
	BlockNumber *uint64
	// GDGasPrice is the gas price from the gas data.
	GDGasPrice *uint16
	// GDDataPrice is the data price from the gas data.
	GDDataPrice *uint16
	// GDExecBuffer is the exec buffer from the gas data.
	GDExecBuffer *uint16
	// GDAmortAttCost is the amortAttCost from the gas data.
	GDAmortAttCost *uint16
	// GDEtherPrice is the etherPrice from the gas data.
	GDEtherPrice *uint16
	// GDMarkup is the markup from the gas data.
	GDMarkup *uint16
}
