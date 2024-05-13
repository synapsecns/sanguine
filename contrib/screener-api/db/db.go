// Package db provides the database interface for the screener-api.
package db

import (
	"context"
	"errors"
	"time"

	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
)

// BlacklistedAddressWriterDB provides methods to write blacklisted addresses to the database.
type BlacklistedAddressWriterDB interface {
	PutBlacklistedAddress(ctx context.Context, body BlacklistedAddress) error
	DeleteBlacklistedAddress(ctx context.Context, id string) error
	UpdateBlacklistedAddress(ctx context.Context, id string, body BlacklistedAddress) error
}

// BlacklistedAddressReaderDB provides methods to read blacklisted addresses from the database.
type BlacklistedAddressReaderDB interface {
	GetBlacklistedAddress(ctx context.Context, address string) (*BlacklistedAddress, error)
}

// BlacklistedAddressDB is the interface for reading and writing blacklisted addresses to the database.
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

// DB is the general database interface for the screener-api.
type DB interface {
	BlacklistedAddressDB
	RuleDB
}

// ErrNoAddressNotCached is returned when an address is not cached.
var ErrNoAddressNotCached = errors.New("address not cached")
