package base

import (
	"encoding/json"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	ChainIDFieldName = namer.GetConsistentName("ChainID")
	DestinationFieldName = namer.GetConsistentName("Destination")
	NonceFieldName = namer.GetConsistentName("Nonce")
	RootFieldName = namer.GetConsistentName("Root")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	DestinationBlockNumberFieldName = namer.GetConsistentName("DestinationBlockNumber")
	ExecutedFieldName = namer.GetConsistentName("Executed")
	MinimumTimeFieldName = namer.GetConsistentName("MinimumTime")
	SnapshotRootFieldName = namer.GetConsistentName("SnapshotRoot")
	AttestationNonceFieldName = namer.GetConsistentName("AttestationNonce")
}

var (
	// ChainIDFieldName gets the chain id field name.
	ChainIDFieldName string
	// DestinationFieldName is the field name of the destination.
	DestinationFieldName string
	// NonceFieldName is the field name of the tx hash.
	NonceFieldName string
	// AttestationNonceFieldName is the field name of the attestation nonce.
	AttestationNonceFieldName string
	// RootFieldName is the name of the block number field.
	RootFieldName string
	// BlockNumberFieldName is the name of the block number field.
	BlockNumberFieldName string
	// DestinationBlockNumberFieldName is the index field name.
	DestinationBlockNumberFieldName string
	// ExecutedFieldName is the executed field name.
	ExecutedFieldName string
	// MinimumTimeFieldName is the minimum time field name.
	MinimumTimeFieldName string
	// SnapshotRootFieldName is the snapshot root field name.
	SnapshotRootFieldName string
)

// PageSize is the amount of entries per page of logs.
var PageSize = 50_000

// Message is the information about a message parsed by the Executor.
type Message struct {
	// ChainID is the chain id.
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// Destination is the destination.
	Destination uint32 `gorm:"column:destination;primaryKey"`
	// Nonce is the nonce.
	Nonce uint32 `gorm:"column:nonce;primaryKey"`
	// Message is the message.
	Message []byte `gorm:"column:message"`
	// BlockNumber is the block number.
	BlockNumber uint64 `gorm:"column:block_number"`
	// Executed is if the message has been executed.
	Executed bool `gorm:"column:executed"`
	// AttestationNonce is the nonce of the attestation that was used to set the minimum time.
	// If this is 0, then the minimum time was not set by an attestation.
	AttestationNonce uint32 `gorm:"column:attestation_nonce"`
	// MinimumTime is the minimum time that the message can be executed.
	MinimumTime uint64 `gorm:"column:minimum_time"`
}

// Attestation is the information about an attestation parsed by the Executor.
type Attestation struct {
	// Destination is the destination of the attestation.
	Destination uint32 `gorm:"column:destination;primaryKey"`
	// SnapshotRoot is the snapshot root.
	SnapshotRoot string `gorm:"column:snapshot_root;primaryKey"`
	// Height is the height of the snapshot Merkle tree.
	Height uint8 `gorm:"column:height"`
	// AttestationNonce is the nonce of the attestation.
	AttestationNonce uint32 `gorm:"column:attestation_nonce;primaryKey"`
	// SummitBlockNumber is the block number when the attestation was created in Summit.
	SummitBlockNumber uint64 `gorm:"column:summit_block_number"`
	// SummitTimestamp is the timestamp of the block when the attestation was created in Summit.
	SummitTimestamp uint64 `gorm:"column:summit_timestamp"`
	// DestinationBlockNumber is the block number that the attestation was submitted on the destination.
	DestinationBlockNumber uint64 `gorm:"column:destination_block_number"`
	// DestinationTimestamp is the timestamp of the block that the attestation was submitted on the destination.
	DestinationTimestamp uint64 `gorm:"column:destination_timestamp"`
}

// State is the information about a state, received from the `Summit` and parsed by the Executor.
type State struct {
	// SnapshotRoot is the snapshot root.
	SnapshotRoot string `gorm:"column:snapshot_root;primaryKey"`
	// Root is the origin Merkle tree's root.
	Root string `gorm:"column:root;primaryKey"`
	// ChainID is the origin chain id.
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// Nonce is the origin Merkle tree's nonce.
	Nonce uint32 `gorm:"column:nonce;primaryKey"`
	// OriginBlockNumber is the block number that the state was taken from on the origin.
	OriginBlockNumber uint64 `gorm:"column:origin_block_number"`
	// OriginTimestamp is the timestamp of the block that the state was taken from on the origin.
	OriginTimestamp uint64 `gorm:"column:origin_timestamp"`
	// Proof is the Snapshot Merkle Tree proof for the state.
	Proof json.RawMessage `gorm:"column:proof"`
	// TreeHeight is the height of the Snapshot Merkle Tree that the state belongs to.
	TreeHeight uint32 `gorm:"column:tree_height"`
	// StateIndex is the index of the state in the Snapshot.
	StateIndex uint32 `gorm:"column:state_index"`
}
