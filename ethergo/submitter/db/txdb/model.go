package txdb

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"time"
)

// define common field names. See package docs  for an explanation of why we have to do this.
// note: some models share names. In cases where they do, we run the check against all names.
// This is cheap because it's only done at startup.
func init() {
	namer := dbcommon.NewNamer(GetAllModels())
	txHashFieldName = namer.GetConsistentName("TXHash")
	chainIDFieldName = namer.GetConsistentName("ChainID")
	nonceFieldName = namer.GetConsistentName("Nonce")
	statusFieldName = namer.GetConsistentName("Status")
	createdAtFieldName = namer.GetConsistentName("CreatedAt")
	fromFieldName = namer.GetConsistentName("From")
	idFieldName = namer.GetConsistentName("ID")
}

var (
	// txHashFieldName is the field name of the tx hash.
	txHashFieldName string
	// chainIDFieldName is the field name of the to address.
	chainIDFieldName string
	// nonceFieldName is the field name of the nonce.
	nonceFieldName string
	// statusFieldName is the field name of the status.
	statusFieldName string
	// createdAtFieldName is the field name of the created at time.
	createdAtFieldName string
	// fromFieldName is the field name of the from address.
	fromFieldName string
	// idFieldName is the field name of the id.
	idFieldName string
)

// ETHTX contains a raw evm transaction that is unsigned.
type ETHTX struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement:true"`
	// CreatedAt is the time the transaction was created
	CreatedAt time.Time
	// TXHash is the hash of the transaction
	TXHash string `gorm:"column:tx_hash;uniqueIndex;size:256"`
	// From is the sender of the transaction
	From string `gorm:"column:from;index"`
	// ChainID is the chain id the transaction hash will be sent on
	ChainID uint64 `gorm:"column:chain_id;index"`
	// Nonce is the nonce of the raw evm tx
	Nonce uint64 `gorm:"column:nonce;index"`
	// RawTx is the raw serialized transaction
	RawTx []byte `gorm:"column:raw_tx"`
	// Status is the status of the transaction
	Status db.Status `gorm:"column:status;index"`
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = []interface{}{&ETHTX{}}
	return allModels
}
