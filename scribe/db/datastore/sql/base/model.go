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

// RawEthTX contains a raw evm transaction that is unsigned
// note: idx_id contains a composite index of (chain_id,nonce)
type RawEthTX struct {
	gorm.Model
	// From is the sender of the transaction
	From string `gorm:"from"`
	// To is the contract address the transaction was sent to.
	To string `gorm:"index"`
	// ChainID is the chain id the transaction hash will be sent on
	ChainID uint64 `gorm:"column:chain_id;uniqueIndex:idx_id"`
	// Nonce is the nonce of the raw evm tx
	Nonce uint64 `gorm:"column:nonce;uniqueIndex:idx_id"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
}

// ProcessedEthTx contains a processed ethereum transaction.
type ProcessedEthTx struct {
	TxHash string `gorm:"txhash;uniqueIndex:idx_txhash;size:66"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
	// RawEthTx is the txid that caused the event
	RawEthTx uint
	// OriginatingEvent is the event that originated the tx
	EthTx RawEthTX `gorm:"foreignkey:RawEthTx"`
	// GasFeeCap contains the gas fee cap stored in wei
	GasFeeCap uint64
	// GasTipCap contains the gas tip cap stored in wei
	GasTipCap uint64
}
