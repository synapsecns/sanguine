package base

import "github.com/jackc/pgtype"

type TokenSwap struct {
	// ContractAddress is the address of the contract that generated the event
	ContractAddress string `gorm:"column:contract_address;primaryKey"`
	// ChainID is the chain id of the contract that generated the event
	ChainID uint32 `gorm:"column:chain_id;primaryKey;auto_increment:false"`
	// TxHash is the hash of the transaction
	TxHash string `gorm:"column:tx_hash;primaryKey"`
	// TxIndex is the index of the transaction in the block
	TxIndex uint64 `gorm:"tx_index"`
	// Buyer is the buyer of the tkoen
	Buyer string
	// TokensSold is the amount of tokens sold
	TokensSold pgtype.Numeric
	// TokensBought is the amount of tokens bought
	TokensBought pgtype.Numeric
	// TokenIndexFrom is the token index we sold from
	TokenIndexFrom uint32
	// TokenIndexTo is the token index to
	TokenIndexTo uint32
}
