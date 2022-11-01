package types

// FailedLog contains the logs that have failed to be indexed.
type FailedLog struct {
	// ChainID is the chain id of the contract
	ChainID uint32 `gorm:"column:chain_id;primaryKey"`
	// ContractAddress is the contract address
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// BlockIndex is the index of the log in the block
	BlockIndex uint64 `gorm:"column:block_index;primaryKey"`
	// BlockNumber is the block in which this transaction was included
	BlockNumber uint64 `gorm:"column:block_number;primaryKey"`
	// FailedAttempts is the number of times this log has failed to be indexed
	FailedAttempts uint64 `gorm:"column:failed_attempts"`
}
