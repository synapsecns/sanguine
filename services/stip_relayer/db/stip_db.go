// Package db provides the database interfaces and types for the STIP Relayer.
package db

import (
	"time"
)

// ApiResponse is the model that saves Dune query API execution IDs with their results.
type ApiResponse struct {
	ExecutionID      string             `gorm:"column:execution_id;index;primaryKey"`
	STIPTransactions []STIPTransactions `gorm:"foreignKey:ExecutionID"`
}

// Row is the model that saves raw Dune response data
type STIPTransactions struct {
	ExecutionID string    `gorm:"column:execution_id;index"`
	Address     string    `gorm:"column:address"`
	Amount      float64   `gorm:"column:amount"`
	AmountUSD   float64   `gorm:"column:amount_usd"`
	ArbPrice    float64   `gorm:"column:arb_price"`
	BlockTime   time.Time `gorm:"column:block_time"`
	Direction   string    `gorm:"column:direction"`
	Hash        string    `gorm:"column:hash;index;primaryKey"`
	Token       string    `gorm:"column:token"`
	TokenPrice  float64   `gorm:"column:token_price"`
	Rebated     bool      `gorm:"column:rebated"`
	Nonce       uint64    `gorm:"column:rebated_nonce"`
}

// STIPDBReader is the interface for reading from the database.
type STIPDBReader interface {
}

// STIPDBWriter is the interface for writing to the database.
type STIPDBWriter interface {
}

// STIPDB is the interface for the database service.
type STIPDB interface {
	STIPDBReader
	STIPDBWriter
}
