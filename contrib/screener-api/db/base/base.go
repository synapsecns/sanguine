// Package base provides a base store for the screener-api.
package base

import (
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/core/metrics"
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

// GetAllModels gets all models to migrate.
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels, &db.AddressIndicators{})

	return allModels
}

// GetAddressIndicators gets the address indicators for the given address.
func (s *Store) GetAddressIndicators(address string) (*db.AddressIndicators, error) {
	var addressIndicators db.AddressIndicators
	result := s.db.Where("address = ?", address).Find(&addressIndicators)
	if result.Error != nil {
		return nil, result.Error
	}
	return &addressIndicators, nil
}

func (s *Store) PutAddressIndicators(addressIndicators string) error {
	return s.db.Save(addressIndicators).Error
}
