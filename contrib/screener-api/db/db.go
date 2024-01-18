package db

import (
	"context"
	"errors"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"time"
)

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

// ErrNoAddressNotCached is returned when an address is not cached.
var ErrNoAddressNotCached = errors.New("address not cached")
