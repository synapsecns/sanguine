package base

import (
	"github.com/synapsecns/sanguine/services/explorer/db"
	"gorm.io/gorm"
)

// Store implements service.
type Store struct {
	db *gorm.DB
}

// NewStore creates a new store from a gorm db.
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// DB gets the underlying gorm db.
func (s Store) DB() *gorm.DB {
	return s.db
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels,
		&BridgeModel{},
		&SwapModel{},
	)
	return allModels
}

var _ db.EventDB = Store{}
