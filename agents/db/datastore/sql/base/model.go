package base

import (
	"time"

	"github.com/synapsecns/sanguine/core/dbcommon"

	"gorm.io/gorm"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	NonceFieldName = namer.GetConsistentName("Nonce")
	DomainIDFieldName = namer.GetConsistentName("DomainID")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
}

var (
	// NonceFieldName is the field name of the nonce.
	NonceFieldName string
	// DomainIDFieldName gets the chain id field name.
	DomainIDFieldName string
	// BlockNumberFieldName is the name of the block number field.
	BlockNumberFieldName string
	// LeafIndexFieldName is the field name of the leaf index.
	LeafIndexFieldName string
)

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

// BlockEndModel is used to make sure we haven't missed any events while offline.
// since we event source - rather than use a state machine this is needed to make sure we haven't missed any events
// by allowing us to go back and source any events we may have missed.
//
// this does not inherit from gorm.model to allow us to use ChainID as a primary key.
type BlockEndModel struct {
	// CreatedAt is the creation time
	CreatedAt time.Time
	// UpdatedAt is the update time
	UpdatedAt time.Time
	// DeletedAt time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// DomainID is the chain id of the chain we're watching blocks on. This is our primary index.
	DomainID uint32 `gorm:"column:domain_id;primaryKey;autoIncrement:false"`
	// BlockHeight is the highest height we've seen on the chain
	BlockNumber uint32 `gorm:"block_number"`
}
