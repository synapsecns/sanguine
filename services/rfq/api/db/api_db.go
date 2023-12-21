// Package db provides the database interfaces and types for the RFQ API.
package db

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

// Quote is the database model for a quote.
type Quote struct {
	// OriginChainID is the chain which the relayer is willing to relay from
	OriginChainID uint64 `gorm:"column:origin_chain_id;index;primaryKey"`
	// OriginTokenAddr is the token address for which the relayer willing to relay from
	OriginTokenAddr string `gorm:"column:origin_token;index;primaryKey"`
	// DestChainID is the chain which the relayer is willing to relay to
	DestChainID uint64 `gorm:"column:dest_chain_id;index;primaryKey"`
	// DestToken is the token address for which the relayer willing to relay to
	DestTokenAddr string `gorm:"column:dest_token;index;primaryKey"`
	// DestAmount is the max amount of liquidity which exists for a given destination token, provided in the destination token decimals
	DestAmount decimal.Decimal `gorm:"column:dest_amount"`
	// MaxOriginAmount is the maximum amount of origin tokens bridgeable
	MaxOriginAmount decimal.Decimal `gorm:"column:max_origin_amount"`
	// FixedFee is the fixed fee for the quote, provided in the destination token terms
	FixedFee decimal.Decimal `gorm:"column:fixed_fee"`
	// Address of the relayer providing the quote
	RelayerAddr string `gorm:"column:relayer_address;primaryKey"`
	// UpdatedAt is the time that the quote was last upserted
	UpdatedAt time.Time
}

// APIDBReader is the interface for reading from the database.
type APIDBReader interface {
	// GetQuotesByDestChainAndToken gets quotes from the database by destination chain and token.
	GetQuotesByDestChainAndToken(ctx context.Context, destChainID uint64, destTokenAddr string) ([]*Quote, error)
	// GetQuotesByOriginAndDestination gets quotes from the database by origin and destination.
	GetQuotesByOriginAndDestination(ctx context.Context, originChainID uint64, originTokenAddr string, destChainID uint64, destTokenAddr string) ([]*Quote, error)
	// GetQuotesByRelayerAddress gets quotes from the database by relayer address.
	GetQuotesByRelayerAddress(ctx context.Context, relayerAddress string) ([]*Quote, error)
	// GetAllQuotes retrieves all quotes from the database.
	GetAllQuotes(ctx context.Context) ([]*Quote, error)
}

// APIDBWriter is the interface for writing to the database.
type APIDBWriter interface {
	// UpsertQuote upserts a quote in the database.
	UpsertQuote(ctx context.Context, quote *Quote) error
}

// APIDB is the interface for the database service.
type APIDB interface {
	APIDBReader
	APIDBWriter
}
