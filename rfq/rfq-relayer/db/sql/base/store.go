package base

import (
	"github.com/synapsecns/sanguine/core/metrics"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"gorm.io/gorm"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	db             *gorm.DB
	metrics        metrics.Handler
	submitterStore submitterDB.Service
}

// UNSAFE_DB gets the underlying gorm db. This is for use only for testing.
//
//nolint:golint,revive,stylecheck
func (s Store) UNSAFE_DB() *gorm.DB {
	return s.db
}

// NewStore creates a new store.
func NewStore(db *gorm.DB, metrics metrics.Handler) *Store {
	txDB := txdb.NewTXStore(db, metrics)
	return &Store{db: db, metrics: metrics, submitterStore: txDB}
}

// DB gets the database.
func (s Store) DB() *gorm.DB {
	return s.db
}

// SubmitterDB gets the submitter database object for mutation.
func (s Store) SubmitterDB() submitterDB.Service {
	return s.submitterStore
}

var _ db.DB = Store{}
