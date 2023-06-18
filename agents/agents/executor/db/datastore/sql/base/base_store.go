package base

import (
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"gorm.io/gorm"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	db *gorm.DB
}

// NewStore creates a new store.
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// DB gets the database.
func (s Store) DB() *gorm.DB {
	return s.db
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels,
		&Message{}, &Attestation{}, &State{},
	)
	return allModels
}

var _ db.ExecutorDB = Store{}
