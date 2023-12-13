package db

import (
	"time"

	"github.com/shopspring/decimal"
)

type QuoteModel struct {
	// ID is the unique identifier saved of each quote provided
	ID uint64 `gorm:"column:id;primaryKey;"`
	// DestChainID is the chain which the relayer is willing to provide liquidity for
	DestChainID uint64 `gorm:"column:dest_chain_id;index"`
	// DestToken is the token address for which the relayer is providing liquidity
	DestTokenAddr string `gorm:"column:token;index"`
	// DestAmount is the max amount of liquidity which exists for a given destination token, provided in the destination token decimals
	DestAmount decimal.Decimal `gorm:"column:dest_amount"`
	// Price is the price per origin token provided for which a relayer is indicating willingness to relay
	Price decimal.Decimal `gorm:"column:price"`
	// UpdatedAt is the time that the quote was last upserted
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// ApiDBReader is the interface for reading from the database.
type ApiDBReader interface {
}

// ApiDBWriter is the interface for writing to the database.
type ApiDBWriter interface {
}

// ApiDB is the interface for the database service.
type ApiDB interface {
	ApiDBReader
	ApiDBWriter
}
