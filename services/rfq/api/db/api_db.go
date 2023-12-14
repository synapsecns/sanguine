// Package db provides the database interfaces and types for the RFQ API.
package db

import (
	"time"

	"github.com/shopspring/decimal"
)

type Quote struct {
	// ID is the unique identifier saved of each quote provided
	ID uint64 `gorm:"column:id;primaryKey;"`
	// OriginChainID is the chain which the relayer is willing to relay from
	OriginChainID uint64 `gorm:"column:origin_chain_id;index"`
	// OriginTokenAddr is the token address for which the relayer willing to relay from
	OriginTokenAddr string `gorm:"column:origin_token;index"`
	// DestChainID is the chain which the relayer is willing to relay to
	DestChainID uint64 `gorm:"column:dest_chain_id;index"`
	// DestToken is the token address for which the relayer willing to relay to
	DestTokenAddr string `gorm:"column:dest_token;index"`
	// DestAmount is the max amount of liquidity which exists for a given destination token, provided in the destination token decimals
	DestAmount decimal.Decimal `gorm:"column:dest_amount"`
	// Price is the price per origin token provided for which a relayer is indicating willingness to relay
	Price decimal.Decimal `gorm:"column:price"`
	// MaxOriginAmount is the maximum amount of origin tokens bridgeable, calculated by dividing the DestAmount by the Price
	MaxOriginAmount decimal.Decimal `gorm:"column:max_origin_amount"`
	// UpdatedAt is the time that the quote was last upserted
	UpdatedAt time.Time
}

// ApiDBReader is the interface for reading from the database.
type ApiDBReader interface {
	// GetQuote gets a quote from the database.
	GetQuotesByDestChainAndToken(destChainId uint64, destTokenAddr string) ([]*Quote, error)
	GetQuotesByOriginAndDestination(originChainId uint64, originTokenAddr string, destChainId uint64, destTokenAddr string) ([]*Quote, error)
	// GetAllQuotes retrieves all quotes from the database.
	GetAllQuotes() ([]*Quote, error)
}

// ApiDBWriter is the interface for writing to the database.
type ApiDBWriter interface {
	// UpsertQuote upserts a quote in the database.
	UpsertQuote(quote *Quote) error
}

// ApiDB is the interface for the database service.
type ApiDB interface {
	ApiDBReader
	ApiDBWriter
}
