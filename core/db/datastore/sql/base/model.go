package base

import (
	"gorm.io/gorm"
)

// RawEVMTX contains a raw evm transaction that is unsigned
// note: idx_id contains a composite index of (chain_id,nonce)
type RawEVMTX struct {
	gorm.Model
	// To is the contract address the transaction was sent to.
	To string `gorm:"index"`
	// ChainID is the chain id the transaction hash will be sent on
	ChainID uint64 `gorm:"column:chain_id,index:idx_id,unique"`
	// Nonce is the nonce of the raw evm tx
	Nonce uint64 `gorm:"index:idx_id,unique"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
}
