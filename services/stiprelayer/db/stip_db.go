// Package db provides the database interfaces and types for the STIP Relayer.
package db

import (
	"context"
	"math/big"
	"time"

	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

// STIPTransactions is the model that saves raw Dune response data.
type STIPTransactions struct {
	ExecutionID      string    `gorm:"column:execution_id;index"`
	Address          string    `gorm:"column:address"`
	Amount           float64   `gorm:"column:amount"`
	AmountUSD        float64   `gorm:"column:amount_usd"`
	ArbPrice         float64   `gorm:"column:arb_price"`
	BlockTime        time.Time `gorm:"column:block_time"`
	Direction        string    `gorm:"column:direction"`
	Hash             string    `gorm:"column:hash;index;primaryKey"`
	Module           string    `gorm:"column:module"`
	Token            string    `gorm:"column:token"`
	TokenPrice       float64   `gorm:"column:token_price"`
	Rebated          bool      `gorm:"column:rebated"`
	Nonce            uint64    `gorm:"column:nonce"`
	DoNotProcess     bool      `gorm:"column:do_not_process"`
	ArbAmountRebated string    `gorm:"column:arb_amount_rebated"`
}

// STIPDBReader is the interface for reading from the database.
type STIPDBReader interface {
	GetSTIPTransactionsNotRebated(ctx context.Context) ([]*STIPTransactions, error)
	GetTotalArbRebated(ctx context.Context, address string) (*big.Int, error)
}

// STIPDBWriter is the interface for writing to the database.
type STIPDBWriter interface {
	UpdateSTIPTransactionRebated(ctx context.Context, hash string, nonce uint64, arbAmountRebated string) error
	InsertNewStipTransactions(ctx context.Context, stipTransactions []STIPTransactions) error
	UpdateSTIPTransactionDoNotProcess(ctx context.Context, hash string) error
}

// STIPDB is the interface for the database service.
type STIPDB interface {
	STIPDBReader
	STIPDBWriter
	// SubmitterDB returns the submitter database service.
	SubmitterDB() submitterDB.Service
}
