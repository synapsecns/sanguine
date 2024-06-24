// Package db provides the database interface for the screener-api.
package db

import (
	"context"
	"errors"
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

// DB is the general database interface for the screener-api.
type DB interface {
	BlacklistedAddressDB
}

// ErrNoAddressNotFound is returned when an address is not cached.
var ErrNoAddressNotFound = errors.New("address not cached")
