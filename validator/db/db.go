package db

import (
	"github.com/cockroachdb/pebble"
	_ "github.com/cockroachdb/pebble"
	_ "github.com/synapsecns/synapse-node/config"
)

type DB interface {
}

// pebbleDB contains a rocksdb used to store merkle trees
type pebbleDB struct {
	*pebble.DB
}

// NewDB creates a new db
func NewDB(dbPath string) (DB, error) {
	db, err := pebble.Open(dbPath, &pebble.Options{})
	_ = db
	_ = err

	return nil, nil
}
