package base

import "github.com/synapsecns/sanguine/core/dbcommon"

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
}

var (
	// ChainIDFieldName gets the chain id field name.
	ChainIDFieldName string
	// DestinationFieldName is the field name of the destination.
	DestinationFieldName string
	// NonceFieldName is the field name of the tx hash.
	NonceFieldName string
	// RootFieldName is the name of the block number field.
	RootFieldName string
	// BlockNumberFieldName is the index field name.
	BlockNumberFieldName string
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
}

// Attestation is the information about an attestation parsed by the Executor.
type Attestation struct {
	// ChainID is the chain id.
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// Destination is the destination.
	Destination uint32 `gorm:"column:destination;primaryKey"`
	// Nonce is the nonce.
	Nonce uint32 `gorm:"column:nonce;primaryKey"`
	// Root is the root.
	Root string `gorm:"column:root;primaryKey"`
	// DestinationBlockNumber is the block number that the attestation was submitted on the destination.
	DestinationBlockNumber uint64 `gorm:"column:block_number"`
	// DestinationBlockTime is the timestamp of the block that the attestation was submitted on the destination.
	DestinationBlockTime uint64 `gorm:"column:destination_block_time"`
}
