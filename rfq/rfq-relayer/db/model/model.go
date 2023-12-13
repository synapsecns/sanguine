// Package model holds all the models for the database.
package model

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"gorm.io/gorm"
)

// GetAllModels gets all models to migrate.
func GetAllModels() (allModels []interface{}) {
	return append([]interface{}{&OriginBridgeEvent{}, &Token{}, &LastIndexed{}, &DeadlineQueue{}, &DestinationBridgeEvent{}}, txdb.GetAllModels()...)
}

func init() {
	allModels := GetAllModels()
	namer := dbcommon.NewNamer(allModels)
	TransactionIDFieldName = namer.GetConsistentName("TransactionID")
	ContractAddressFieldName = namer.GetConsistentName("ContractAddress")
	ChainIDFieldName = namer.GetConsistentName("ChainID")
	AddressFieldName = namer.GetConsistentName("Address")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	TokenIDFieldName = namer.GetConsistentName("TokenID")
	SymbolFieldName = namer.GetConsistentName("Symbol")
	NameFieldName = namer.GetConsistentName("Name")
	DecimalsFieldName = namer.GetConsistentName("Decimals")
}

var (
	// TransactionIDFieldName is the field name of the transactionID.
	TransactionIDFieldName string
	// ContractAddressFieldName is the field name of the contract address.
	ContractAddressFieldName string
	// ChainIDFieldName is the field name of the chain id.
	ChainIDFieldName string
	// AddressFieldName is the field name of the address.
	AddressFieldName string
	// BlockNumberFieldName is the field name of the block number.
	BlockNumberFieldName string
	// TokenIDFieldName is the field name of the token id, used in the token and balance tables.
	TokenIDFieldName string
	// SymbolFieldName is the field name of the symbol for the token table.
	SymbolFieldName string
	// NameFieldName is the field name of the name for the token table.
	NameFieldName string
	// DecimalsFieldName is the field name of the decimals for the token table.
	DecimalsFieldName string
)

// OriginBridgeEvent is the table that holds every origin fast bridge transaction (bridge).
type OriginBridgeEvent struct {
	TransactionID string `gorm:"column:transaction_id;primaryKey"`
	Request       string
	OriginChainID uint32
	DestChainID   uint32
	OriginSender  string
	DestRecipient string
	OriginToken   string
	DestToken     string
	OriginAmount  string
	DestAmount    string
	Deadline      string
	Nonce         string
	BlockNumber   uint64
	TxHash        string
	TxIndex       uint
	BlockHash     string
	LogIndex      uint
	Removed       bool
}

// DestinationBridgeEvent is the table that holds every destination fast bridge transaction (relay).
type DestinationBridgeEvent struct {
	TransactionID string `gorm:"column:transaction_id;primaryKey"`
	Request       string
	OriginChainID uint32
	DestChainID   uint32
	BlockNumber   uint64
	TxHash        string
	TxIndex       uint
	BlockHash     string
	LogIndex      uint
	Removed       bool
}

// Token holds token metadata.
type Token struct {
	TokenID  string `gorm:"column:token_id;primaryKey"` // Hash of ChainID and TokenAddress
	ChainID  uint32
	Address  string
	Symbol   string
	Name     string
	Decimals int
}

// DeadlineQueue is the underlying table in the deadline queue.
type DeadlineQueue struct {
	Timestamp     int64
	TransactionID string `gorm:"column:transaction_id;primaryKey"`
}

// LastIndexed contains information on when a contract was last indexed.
type LastIndexed struct {
	gorm.Model
	// ContractAddress is the contract address
	ContractAddress string `gorm:"column:contract_address;index:idx_last_indexed,priority:1;uniqueIndex:idx_contract_chain"`
	// BlockNumber is the last block number indexed.
	// NOTE: SQLite's max int is below the max uint64 in golang. However, we can assume that no chain will have over 9 quintillion blocks.
	BlockNumber uint64 `gorm:"column:block_number;index:idx_last_indexed,priority:2"`
	// ChainID is the chain id of the contract
	ChainID uint32 `gorm:"column:chain_id;index:idx_last_indexed;uniqueIndex:idx_contract_chain"`
}
