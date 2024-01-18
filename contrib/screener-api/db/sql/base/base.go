// Package base provides a base store for the screener-api.
package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/contrib/screener-api/trmlabs"
	"github.com/synapsecns/sanguine/core/metrics"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
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
func (s *Store) GetAddressIndicators(ctx context.Context, address string, since time.Time) ([]trmlabs.AddressRiskIndicator, error) {
	var addressIndicators db.AddressIndicators
	result := s.db.WithContext(ctx).Where(&db.AddressIndicators{
		Address: strings.ToLower(address),
	}).First(&addressIndicators)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, db.ErrNoAddressNotCached
		}
		return nil, result.Error
	}

	// if the address indicators are not found, return nil
	if addressIndicators.UpdatedAt.Before(since) {
		return nil, db.ErrNoAddressNotCached
	}

	return addressIndicators.Indicators.ToTRMLabs(), nil
}

func (s *Store) PutAddressIndicators(ctx context.Context, address string, riskIndicator []trmlabs.AddressRiskIndicator) error {
	dbTx := s.db.WithContext(ctx).Model(&db.AddressIndicators{}).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: addressName},
			},
			DoUpdates: clause.AssignmentColumns([]string{addressName, indicatorName}),
		}).Create(db.MakeRecord(address, riskIndicator))
	if dbTx.Error != nil {
		return fmt.Errorf("failed to store address indicators: %w", dbTx.Error)
	}
	return nil
}
