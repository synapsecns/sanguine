package base

import (
	"github.com/synapsecns/sanguine/committee/db"
	"time"
)

// SignRequest is a request to sign a message.
type SignRequest struct {
	// TXHash is the hash of the transaction
	TXHash string `gorm:"column:tx_hash;index;size:256"`
	// TXHash is the hash of the transaction
	Transaction string `gorm:"column:transaction;primaryKey;size:256"`
	// CreatedAt is the time the transaction was created
	CreatedAt time.Time
	// ChainID is the chain id the transaction hash will be sent on
	ChainID uint64 `gorm:"column:chain_id;index"`
	// Nonce is the nonce of the raw evm tx
	Nonce uint64 `gorm:"column:nonce;index"`
	// Status is the status of the transaction
	Status db.SynapseRequestStatus `gorm:"column:status;index"`
	// SignRequest is a reference to the SignRequest
	SignRequest *SignRequest `gorm:"foreignKey:Transaction"`
}

type Signature struct {
	// Signature is the actual signature
	Signature []byte `gorm:"column:signature"`
	// SenderAddress is the address of the sender
	SenderAddress string `gorm:"column:sender_address;index;size:256"`
	// Transaction is the actual transaction
	Transaction string `gorm:"column:transaction"`
	// SignRequest is a reference to the SignRequest
	SignRequest *SignRequest `gorm:"foreignKey:Transaction"`
}
