package db

import (
	"fmt"
	"github.com/cockroachdb/pebble"
)

// DB contains the synapse db.
type DB interface{}

// pebbleDB contains a rocksdb used to store merkle trees.
type pebbleDB struct {
	*pebble.DB
}

// NewDB creates a new db.
func NewDB(dbPath string) (DB, error) {
	db, err := pebble.Open(dbPath, &pebble.Options{})

	if err != nil {
		return nil, fmt.Errorf("could not create db: %w", err)
	}

	return pebbleDB{DB: db}, nil
}
