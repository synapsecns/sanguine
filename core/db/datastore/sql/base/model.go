package base

import (
	"gorm.io/gorm"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	NonceFieldName = getConsistentName("Nonce")
}

// NonceFieldName is the field name of the nonce.
var NonceFieldName string

// RawEVMTX contains a raw evm transaction that is unsigned
// note: idx_id contains a composite index of (chain_id,nonce)
type RawEVMTX struct {
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

// TableName gets the raw evm txes.
func (r RawEVMTX) TableName() string {
	return "raw_evm_txes"
}
