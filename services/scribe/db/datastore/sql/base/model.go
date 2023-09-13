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
	BlockIndexFieldName = namer.GetConsistentName("BlockIndex")
	BlockHashFieldName = namer.GetConsistentName("BlockHash")
	ConfirmedFieldName = namer.GetConsistentName("Confirmed")
	TransactionIndexFieldName = namer.GetConsistentName("TransactionIndex")
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
	// BlockIndexFieldName is the index field name.
	BlockIndexFieldName string
	// BlockHashFieldName is the block hash field name.
	BlockHashFieldName string
	// ConfirmedFieldName is the confirmed field name.
	ConfirmedFieldName string
	// TransactionIndexFieldName is the name of the transaction block  field.
	TransactionIndexFieldName string
)

// PageSize is the amount of entries per page of logs.
var PageSize = 100

// LogColumns are all of the columns of the Log table.
const LogColumns = "contract_address,chain_id,primary_topic,topic_a,topic_b,topic_c,data,block_number,tx_hash,tx_index,block_hash,block_index,removed,confirmed"

// Log stores the log of an event.
type Log struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey;index:idx_address,priority:1,sort:desc"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;index:idx_address,priority:2,sort:desc"`
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
	BlockNumber uint64 `gorm:"column:block_number;index:idx_block_number,priority:1,sort:desc"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey;index:idx_tx_hash,priority:1,sort:desc"`
	// TxIndex is the index of the transaction in the block
	TxIndex uint64 `gorm:"tx_index"`
	// BlockHash is the hash of the block in which the transaction was included
	BlockHash string `gorm:"column:block_hash;index:idx_block_hash,priority:1,sort:desc"`
	// Index is the index of the log in the block
	BlockIndex uint64 `gorm:"column:block_index;primaryKey;index:idx_block_number,priority:2,sort:desc"`
	// Removed is true if this log was reverted due to a chain re-organization
	Removed bool `gorm:"removed"`
	// Confirmed is true if this log has been confirmed by the chain
	Confirmed bool `gorm:"confirmed"`
}

// ReceiptColumns are all of the columns of the Receipt table.
const ReceiptColumns = "chain_id,receipt_type,post_state,status,cumulative_gas_used,bloom,tx_hash,contract_address,gas_used,block_hash,block_number,transaction_index,confirmed"

// Receipt stores the receipt of a transaction.
type Receipt struct {
	// ChainID is the chain id of the receipt
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// Type is the type
	Type uint8 `gorm:"column:receipt_type"`
	// PostState is the post state
	PostState []byte `gorm:"column:post_state"`
	// Status is the status of the transaction
	Status uint64 `gorm:"column:status"`
	// CumulativeGasUsed is the total amount of gas used when this transaction was executed in the block
	CumulativeGasUsed uint64 `gorm:"column:cumulative_gas_used"`
	// Bloom is the bloom filter
	Bloom []byte `gorm:"column:bloom"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// ContractAddress is the address of the contract
	ContractAddress string `gorm:"column:contract_address"`
	// GasUsed is the amount of gas used by this transaction alone
	GasUsed uint64 `gorm:"column:gas_used"`
	// BlockHash is the hash of the block in which this transaction was included
	BlockHash string `gorm:"column:block_hash"`
	// BlockNumber is the block in which this transaction was included
	BlockNumber uint64 `gorm:"column:block_number;index:idx_block_number_receipt,priority:1,sort:desc"`
	// TransactionIndex is the index of the transaction in the block
	TransactionIndex uint64 `gorm:"column:transaction_index;index:idx_block_number_receipt,priority:2,sort:desc"`
	// Confirmed is true if this log has been confirmed by the chain
	Confirmed bool `gorm:"column:confirmed"`
}

// EthTxColumns are all of the columns of the EthTx table.
const EthTxColumns = "tx_hash,chain_id,block_hash,block_number,raw_tx,gas_fee_cap,gas_tip_cap,confirmed,transaction_index"

// EthTx contains a processed ethereum transaction.
type EthTx struct {
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// ChainID is the chain id of the transaction
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// BlockHash is the hash of the block in which the transaction was included
	BlockHash string `gorm:"column:block_hash;index:idx_tx_block_hash,priority:1,sort:desc"`
	// BlockNumber is the block in which the transaction was included
	BlockNumber uint64 `gorm:"column:block_number;index:idx_block_number_tx,priority:1,sort:desc"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
	// GasFeeCap contains the gas fee cap stored in wei
	GasFeeCap uint64
	// GasTipCap contains the gas tip cap stored in wei
	GasTipCap uint64
	// Confirmed is true if this log has been confirmed by the chain
	Confirmed bool `gorm:"column:confirmed"`
	// TransactionIndex is the index of the transaction in the block
	TransactionIndex uint64 `gorm:"column:transaction_index;index:idx_block_number_tx,priority:2,sort:desc"`
}

// LastIndexedInfo contains information on when a contract was last indexed.
type LastIndexedInfo struct {
	gorm.Model
	// ContractAddress is the contract address
	ContractAddress string `gorm:"column:contract_address;index:idx_last_indexed,priority:1;uniqueIndex:idx_contract_chain"`
	// BlockNumber is the last block number indexed
	BlockNumber uint64 `gorm:"column:block_number;index:idx_last_indexed,priority:2"`
	// ChainID is the chain id of the contract
	ChainID uint32 `gorm:"column:chain_id;index:idx_last_indexed;uniqueIndex:idx_contract_chain"`
}

// LastConfirmedBlockInfo contains information on when a chain last had a block pass the required confirmation
// threshold and was validated.
type LastConfirmedBlockInfo struct {
	gorm.Model
	// ChainID is the chain id of the contract
	ChainID uint32 `gorm:"column:chain_id"`
	// BlockNumber is the last block number indexed
	BlockNumber uint64 `gorm:"column:block_number"`
}

// BlockTime contains the timestamp of a block.
type BlockTime struct {
	// ChainID is the chain id of the contract
	ChainID uint32 `gorm:"column:chain_id;primaryKey;index:idx_block_time_chain,priority:1"`
	// BlockNumber is the block number
	BlockNumber uint64 `gorm:"column:block_number;primaryKey"`
	// Timestamp is the timestamp of the block
	Timestamp uint64 `gorm:"column:timestamp"`
}

// LastBlockTime contains the last block that had its timestamp stored.
type LastBlockTime struct {
	gorm.Model
	// ChainID is the chain id of the contract
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// BlockNumber is the block number
	BlockNumber uint64 `gorm:"column:block_number"`
}

// LogAtHead stores the log of an event that occurred near the tip of the chain.
type LogAtHead struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey;index:idx_head_address,priority:1,sort:desc"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;index:idx_head_address,priority:2,sort:desc"`
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
	BlockNumber uint64 `gorm:"column:block_number;index:idx_head_block_number,priority:1,sort:desc"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey;index:idx_head_tx_hash,priority:1,sort:desc"`
	// TxIndex is the index of the transaction in the block
	TxIndex uint64 `gorm:"tx_index"`
	// BlockHash is the hash of the block in which the transaction was included
	BlockHash string `gorm:"column:block_hash;index:idx_head_block_hash,priority:1,sort:desc"`
	// Index is the index of the log in the block
	BlockIndex uint64 `gorm:"column:block_index;primaryKey;index:idx_head_block_number,priority:2,sort:desc"`
	// Removed is true if this log was reverted due to a chain re-organization
	Removed bool `gorm:"removed"`
	// Confirmed is true if this log has been confirmed by the chain
	Confirmed bool `gorm:"confirmed"`
	// InsertTime is the time at which this log was inserted
	InsertTime uint64 `gorm:"column:insert_time"`
}

// ReceiptAtHead stores the receipt of a transaction at the tip.
type ReceiptAtHead struct {
	// ChainID is the chain id of the receipt
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// Type is the type
	Type uint8 `gorm:"column:receipt_type"`
	// PostState is the post state
	PostState []byte `gorm:"column:post_state"`
	// Status is the status of the transaction
	Status uint64 `gorm:"column:status"`
	// CumulativeGasUsed is the total amount of gas used when this transaction was executed in the block
	CumulativeGasUsed uint64 `gorm:"column:cumulative_gas_used"`
	// Bloom is the bloom filter
	Bloom []byte `gorm:"column:bloom"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// ContractAddress is the address of the contract
	ContractAddress string `gorm:"column:contract_address"`
	// GasUsed is the amount of gas used by this transaction alone
	GasUsed uint64 `gorm:"column:gas_used"`
	// BlockHash is the hash of the block in which this transaction was included
	BlockHash string `gorm:"column:block_hash"`
	// BlockNumber is the block in which this transaction was included
	BlockNumber uint64 `gorm:"column:block_number;index:idx_head_block_number_receipt,priority:1,sort:desc"`
	// TransactionIndex is the index of the transaction in the block
	TransactionIndex uint64 `gorm:"column:transaction_index;index:idx_head_block_number_receipt,priority:2,sort:desc"`
	// Confirmed is true if this log has been confirmed by the chain
	Confirmed bool `gorm:"column:confirmed"`
	// InsertTime is the time at which this receipt was inserted
	InsertTime uint64 `gorm:"column:insert_time"`
}

// EthTxAtHead contains a processed ethereum transaction at the tip of the chain.
type EthTxAtHead struct {
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// ChainID is the chain id of the transaction
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// BlockHash is the hash of the block in which the transaction was included
	BlockHash string `gorm:"column:block_hash;index:idx_head_tx_block_hash,priority:1,sort:desc"`
	// BlockNumber is the block in which the transaction was included
	BlockNumber uint64 `gorm:"column:block_number;index:idx_head_block_number_tx,priority:1,sort:desc"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
	// GasFeeCap contains the gas fee cap stored in wei
	GasFeeCap uint64
	// GasTipCap contains the gas tip cap stored in wei
	GasTipCap uint64
	// Confirmed is true if this log has been confirmed by the chain
	Confirmed bool `gorm:"column:confirmed"`
	// TransactionIndex is the index of the transaction in the block
	TransactionIndex uint64 `gorm:"column:transaction_index;index:idx_head_block_number_tx,priority:2,sort:desc"`
	// InsertTime is the time at which this tx was inserted
	InsertTime uint64 `gorm:"column:insert_time"`
}
