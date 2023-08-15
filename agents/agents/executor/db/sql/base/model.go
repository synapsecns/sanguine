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
	DestinationTimestampFieldName = namer.GetConsistentName("DestinationTimestamp")
	ExecutedFieldName = namer.GetConsistentName("Executed")
	MinimumTimeSetFieldName = namer.GetConsistentName("MinimumTimeSet")
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
	// DestinationTimestampFieldName is the destination timestamp field name.
	DestinationTimestampFieldName string
	// ExecutedFieldName is the executed field name.
	ExecutedFieldName string
	// MinimumTimeSetFieldName is the minimum time set field name.
	MinimumTimeSetFieldName string
	// MinimumTimeFieldName is the minimum time field name.
	MinimumTimeFieldName string
	// SnapshotRootFieldName is the snapshot root field name.
	SnapshotRootFieldName string
)

// PageSize is the amount of entries per page of logs.
var PageSize = 50_000

// Message is the information about a message parsed by the Executor. This is an event derived from the origin contract.
type Message struct {
	// ChainID is the chain id.
	ChainID uint32 `gorm:"column:chain_id;primaryKey;index:idx_chain_dest_nonce"`
	// Destination is the destination.
	Destination uint32 `gorm:"column:destination;primaryKey;index:idx_chain_dest_nonce"`
	// Nonce is the nonce.
	Nonce uint32 `gorm:"column:nonce;primaryKey;index:idx_chain_dest_nonce"`
	// Message is the message.
	Message []byte `gorm:"column:message"`
	// BlockNumber is the block number.
	BlockNumber uint64 `gorm:"column:block_number"`
	// Executed is if the message has been executed.
	Executed bool `gorm:"column:executed"`
	// MinimumTimeSet is if the MinimumTime field has been set from an Attestation.
	MinimumTimeSet bool `gorm:"column:minimum_time_set"`
	// MinimumTime is the minimum time that the message can be executed.
	MinimumTime uint64 `gorm:"column:minimum_time"`
}

// Attestation is the information about an attestation parsed by the Executor. This is an event derived from the destination contract.
type Attestation struct {
	// Destination is the destination of the attestation.
	Destination uint32 `gorm:"column:destination;primaryKey"`
	// SnapshotRoot is the snapshot root.
	SnapshotRoot string `gorm:"column:snapshot_root;primaryKey"`
	// DataHash is the agent root and SnapGasHash combined into a single hash.
	DataHash string `gorm:"column:data_hash"`
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
	// StateIndex is the index of the state in the Snapshot.
	StateIndex uint32 `gorm:"column:state_index"`
	// BlockNumber is the block number the state update was received on Summit.
	BlockNumber uint64 `gorm:"column:block_number"`
	// GDGasPrice is the gas price from the gas data.
	GDGasPrice uint16 `gorm:"column:gd_gas_price"`
	// GDDataPrice is the data price from the gas data.
	GDDataPrice uint16 `gorm:"column:gd_data_price"`
	// GDExecBuffer is the exec buffer from the gas data.
	GDExecBuffer uint16 `gorm:"column:gd_exec_buffer"`
	// GDAmortAttCost is the amortAttCost from the gas data.
	GDAmortAttCost uint16 `gorm:"column:gd_amort_att_cost"`
	// GDEtherPrice is the etherPrice from the gas data.
	GDEtherPrice uint16 `gorm:"column:gd_ether_price"`
	// GDMarkup is the markup from the gas data.
	GDMarkup uint16 `gorm:"column:gd_markup"`
}
