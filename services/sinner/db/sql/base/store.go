package base

import (
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"gorm.io/gorm"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	db      *gorm.DB
	metrics metrics.Handler
}

// UNSAFE_DB gets the underlying gorm db. This is for use only for testing.
//
//nolint:golint,revive,stylecheck
func (s Store) UNSAFE_DB() *gorm.DB {
	return s.db
}

// NewStore creates a new store.
func NewStore(db *gorm.DB, metrics metrics.Handler) *Store {
	return &Store{db: db, metrics: metrics}
}

// DB gets the database.
func (s Store) DB() *gorm.DB {
	return s.db
}

var _ db.EventDB = Store{}
