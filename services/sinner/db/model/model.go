// Package model holds all the models for the database.
package model

import (
	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/gorm"
)

// PageSize is the size of the page when returning graphql responses.
const PageSize = 100

// GetAllModels gets all models to migrate.
func GetAllModels() (allModels []interface{}) {
	return []interface{}{&OriginSent{}, &Executed{}, &MessageStatus{}, &LastIndexed{}}
}

func init() {
	allModels := GetAllModels()
	namer := dbcommon.NewNamer(allModels)
	TxHashFieldName = namer.GetConsistentName("TxHash")
	ChainIDFieldName = namer.GetConsistentName("ChainID")
	BlockNumberFieldName = namer.GetConsistentName("BlockNumber")
	ContractAddressFieldName = namer.GetConsistentName("ContractAddress")
	MessageHashFieldName = namer.GetConsistentName("MessageHash")

	OriginTxHashFieldName = namer.GetConsistentName("OriginTxHash")
	DestinationTxHashFieldName = namer.GetConsistentName("DestinationTxHash")
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
	// MessageHashFieldName is the name of the message hash field.
	MessageHashFieldName string
	// OriginTxHashFieldName is the name of the origin tx hash field.
	OriginTxHashFieldName string
	// DestinationTxHashFieldName is the name of the destination tx hash field.
	DestinationTxHashFieldName string
)

// MessageStatus is the table holding the status of each message.
type MessageStatus struct {
	// MessageHash is the message hash.
	MessageHash string `gorm:"type:varchar(66);column:message_hash;primaryKey"`
	// OriginTxHash is the txhash when the origin event was emitted.
	OriginTxHash string `gorm:"column:origin_txhash"`
	// DestinationTxHash is the txhash when the destination event was emitted.
	DestinationTxHash string `gorm:"column:successful_destination_txhash"`
}

// OriginSent is the information about a message parsed by the Executor. This is an event derived from the origin contract.
type OriginSent struct {
	// ContractAddress is the address of the contract that generated the event.
	ContractAddress string `gorm:"column:contract_address"`
	// BlockNumber is the block number in which the tx occurred.
	BlockNumber uint64 `gorm:"column:block_number"`
	// TxHash is the hash of the tx.
	TxHash string `gorm:"column:tx_hash;primaryKey;index:idx_tx_hash_origin,priority:1,sort:desc"`
	// TxIndex is the index of the tx in a block.
	TxIndex uint `gorm:"column:tx_index"`
	// Sender is the address of the sender of the tx.
	Sender string `gorm:"column:sender"`
	// Timestamp is the timestamp of the tx.
	Timestamp uint64 `gorm:"column:timestamp"`
	// Recipient is the address of the recipient of the tx.
	Recipient string `gorm:"column:recipient"`
	// MessageLeaf is the keccaked message leaf.
	MessageLeaf string `gorm:"column:message_leaf"`
	// MessageID is the keccaked message.
	MessageID string `gorm:"column:message_id"`
	// MessageHash is the message hash.
	MessageHash string `gorm:"type:varchar(66);column:message_hash;primaryKey"`
	// ChainID is the chain id.
	ChainID uint32 `gorm:"column:chain_id;primaryKey;index:idx_tx_hash_origin,priority:2,sort:desc"`
	// Destination is the destination chain id.
	DestinationChainID uint32 `gorm:"column:destination_chain_id"`

	// Nonce is the nonce.
	Nonce uint32 `gorm:"column:nonce"`
	// Message is the message.
	Message string `gorm:"column:message"`
	// OptimisticSeconds is if the optimistic seconds.
	OptimisticSeconds uint32 `gorm:"column:optimistic_seconds"`

	// MessageFlag is the message flag (system or otherwise).
	MessageFlag uint8 `gorm:"column:message_flag"`
	// SummitTip gets the tips for the agent work on summit
	SummitTip string `gorm:"column:summit_tip"`
	// AttestationTip gets the tips for the doing the attestation
	AttestationTip string `gorm:"column:attestation_tip"`
	// ExecutionTip gets the tips for executing the message
	ExecutionTip string `gorm:"column:execution_tip"`
	// DeliveryTip gets the tips for delivering the message receipt to summit
	DeliveryTip string `gorm:"column:delivery_tip"`
	// Version is the base message version to pass to the recipient.
	Version uint32 `gorm:"column:version"`
	// GasLimit is the minimum amount of gas units to supply for execution.
	GasLimit uint64 `gorm:"column:gas_limit"`
	// GasDrop is the minimum amount of gas token to drop to the recipient.
	GasDrop string `gorm:"column:gas_drop"`
}

// Executed is the information about a message executed on execution hub.
type Executed struct {
	// ContractAddress is the address of the contract that generated the event.
	ContractAddress string `gorm:"column:contract_address"`
	// BlockNumber is the block number in which the tx occurred.
	BlockNumber uint64 `gorm:"column:block_number"`
	// TxHash is the hash of the tx.
	TxHash string `gorm:"column:tx_hash;index:idx_tx_hash_executed,priority:1,sort:desc"`
	// TxIndex is the index of the tx in a block.
	TxIndex uint `gorm:"column:tx_index"`
	// MessageHash is the message hash.
	MessageHash string `gorm:"type:varchar(66);column:message_hash;primaryKey"`
	// ChainID is the chain id.
	ChainID uint32 `gorm:"column:chain_id;primaryKey;index:idx_tx_hash_executed,priority:2,sort:desc"`
	// RemoteDomain is the destination.
	RemoteDomain uint32 `gorm:"column:destination_chain_id"`
	// Success is the status of success of the message.
	Success bool `gorm:"column:success"`
	// Sender is the address of the sender of the tx.
	Sender string `gorm:"column:sender"`
	// Timestamp is the timestamp of the tx.
	Timestamp uint64 `gorm:"column:timestamp"`
}

// LastIndexed contains information on when a contract was last indexed.
type LastIndexed struct {
	gorm.Model
	// ContractAddress is the contract address
	ContractAddress string `gorm:"column:contract_address;index:idx_last_indexed,priority:1;uniqueIndex:idx_contract_chain"`
	// BlockNumber is the last block number indexed
	BlockNumber uint64 `gorm:"column:block_number;index:idx_last_indexed,priority:2"`
	// ChainID is the chain id of the contract
	ChainID uint32 `gorm:"column:chain_id;index:idx_last_indexed;uniqueIndex:idx_contract_chain"`
}
