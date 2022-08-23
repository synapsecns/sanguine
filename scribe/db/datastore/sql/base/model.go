package base

import (
	"database/sql"

	"gorm.io/gorm"
)

// Log stores the log of an event.
type Log struct {
	gorm.Model
	// Address is the address of the contract that generated the event
	Address string `gorm:"address"`
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
	// Status is the status of the transaction
	Status uint64 `gorm:"status"`
	// CumulativeGasUsed is the total amount of gas used when this transaction was executed in the block
	CumulativeGasUsed uint64 `gorm:"cumulative_gas_used"`
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
