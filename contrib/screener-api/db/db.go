// Package db provides the database interface for the screener-api.
package db

import (
	"context"
	"errors"
	"time"

	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
)

// TODO: make a general db interface.
type BlacklistedAddressWriterDB interface {
	PutBlacklistedAddress(ctx context.Context, body BlacklistedAddress) error
	DeleteBlacklistedAddress(ctx context.Context, address string) error
}

type BlacklistedAddressReaderDB interface {
	GetBlacklistedAddress(ctx context.Context, address string) (blacklisted bool, err error)
}

type BlacklistedAddressDB interface {
	BlacklistedAddressWriterDB
	BlacklistedAddressReaderDB
}

// RuleWriterDB is the interface for writing rules to the database.
type RuleWriterDB interface {
	PutAddressIndicators(ctx context.Context, address string, riskIndicator []trmlabs.AddressRiskIndicator) error
}

// RuleReaderDB is the interface for reading rules from the database.
type RuleReaderDB interface {
	GetAddressIndicators(ctx context.Context, address string, since time.Time) ([]trmlabs.AddressRiskIndicator, error)
}

// RuleDB is the interface for reading and writing rules to the database.
type RuleDB interface {
	RuleWriterDB
	RuleReaderDB
}

// DB is the interface for the database.
type DB interface {
	BlacklistedAddressDB
	RuleDB
}

// ErrNoAddressNotCached is returned when an address is not cached.
var ErrNoAddressNotCached = errors.New("address not cached")
