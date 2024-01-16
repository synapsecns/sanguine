// Package db provides the database interfaces and types for the STIP Relayer.
package db

import (
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
	// OriginFastBridgeAddress is the address of the fast bridge contract on the origin chain
	OriginFastBridgeAddress string `gorm:"column:origin_fast_bridge_address"`
	// DestFastBridgeAddress is the address of the fast bridge contract on the destination chain
	DestFastBridgeAddress string `gorm:"column:dest_fast_bridge_address"`
	// UpdatedAt is the time that the quote was last upserted
	UpdatedAt time.Time
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
