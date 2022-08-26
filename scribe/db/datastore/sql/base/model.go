package base

import (
	"database/sql"

	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/gorm"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	TxHashFieldName = namer.GetConsistentName("TxHash")
	ChainIDFieldName = namer.GetConsistentName("ChainID")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	ContractAddressFieldName = namer.GetConsistentName("ContractAddress")
}

var (
	// TxHashFieldName is the field name of the tx hash.
	TxHashFieldName string
	// ChainIDFieldName gets the chain id field name.
	ChainIDFieldName string
	// BlockNumberFieldName is the name of the block number field.
	BlockNumberFieldName string
	// ContractAddressFieldName is the address of the contract.
	ContractAddressFieldName string
)

// Log stores the log of an event.
type Log struct {
	gorm.Model
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"contract_address"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"chain_id"`
	// PrimaryTopic is the primary topic of the event. Topics[0]
	PrimaryTopic sql.NullString `gorm:"primary_topic"`
	// TopicA is the first topic. Topics[1]
	TopicA sql.NullString `gorm:"topic_a"`
	// TopicB is the second topic. Topics[2]
	TopicB sql.NullString `gorm:"topic_b"`
	// TopicC is the third topic. Topics[3]
	TopicC sql.NullString `gorm:"topic_c"`
	// Data is the data provided by the contract
	Data []byte `gorm:"data"`
	// BlockNumber is the block in which the transaction was included
	BlockNumber uint64 `gorm:"block_number"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"tx_hash"`
	// TxIndex is the index of the transaction in the block
	TxIndex uint64 `gorm:"tx_index"`
	// BlockHash is the hash of the block in which the transaction was included
	BlockHash string `gorm:"block_hash"`
	// Index is the index of the log in the block
	Index uint64 `gorm:"index"`
	// Removed is true if this log was reverted due to a chain re-organization
	Removed bool `gorm:"removed"`
}

// Receipt stores the receipt of an transaction.
type Receipt struct {
	gorm.Model
	// Type is the type
	Type uint8 `gorm:"receipt_type"`
	// PostState is the post state
	PostState []byte `gorm:"post_state"`
	// Status is the status of the transaction
	Status uint64 `gorm:"status"`
	// CumulativeGasUsed is the total amount of gas used when this transaction was executed in the block
	CumulativeGasUsed uint64 `gorm:"cumulative_gas_used"`
	// Bloom is the bloom filter
	Bloom []byte `gorm:"bloom"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"tx_hash"`
	// ContractAddress is the address of the contract
	ContractAddress string `gorm:"contract_address"`
	// GasUsed is the amount of gas used by this transaction alone
	GasUsed uint64 `gorm:"gas_used"`
	// BlockHash is the hash of the block in which this transaction was included
	BlockHash string `gorm:"block_hash"`
	// BlockNumber is the block in which this transaction was included
	BlockNumber uint64 `gorm:"block_number"`
	// TransactionIndex is the index of the transaction in the block
	TransactionIndex uint64 `gorm:"transaction_index"`
}

// EthTx contains a processed ethereum transaction.
type EthTx struct {
	TxHash string `gorm:"tx_hash"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
	// GasFeeCap contains the gas fee cap stored in wei
	GasFeeCap uint64
	// GasTipCap contains the gas tip cap stored in wei
	GasTipCap uint64
}

// LastIndexedInfo contains information on when a contract was last indexed.
type LastIndexedInfo struct {
	gorm.Model
	// ContractAddress is the contract address
	ContractAddress string `gorm:"column:contract_address"`
	// ChainID is the chain id of the contract
	ChainID uint32 `gorm:"column:chain_id"`
	// BlockNumber is the last block number indexed
	BlockNumber uint64 `gorm:"column:block_number"`
}
