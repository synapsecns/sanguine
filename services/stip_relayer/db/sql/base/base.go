package base

import (
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/stip_relayer/db"
	"gorm.io/gorm"
)

// Store is a store that implements an underlying gorm db.
type Store struct {
	db      *gorm.DB
	metrics metrics.Handler
}

// NewStore creates a new store.
func NewStore(db *gorm.DB, metrics metrics.Handler) *Store {
	return &Store{db: db, metrics: metrics}
}

// DB gets the database object for mutation outside of the lib.
func (s Store) DB() *gorm.DB {
	return s.db
}

// GetAllModels gets all models to migrate.
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels, &db.ApiResponse{}, &db.STIPTransactions{})
	return allModels
}

var _ db.STIPDB = &Store{}
