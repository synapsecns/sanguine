// Package base provides a base store for the screener-api.
package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/core/metrics"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	allModels = append(allModels, &db.BlacklistedAddress{})

	return allModels
}

// GetBlacklistedAddress queries the db for the blacklisted address.
// Returns true if the address is blacklisted, false otherwise.
// Not used currently.
func (s *Store) GetBlacklistedAddress(ctx context.Context, address string) (*db.BlacklistedAddress, error) {
	var blacklistedAddress db.BlacklistedAddress

	if err := s.db.WithContext(ctx).Where("address = ?", address).
		First(&blacklistedAddress).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, db.ErrNoAddressNotFound
		}
		return nil, fmt.Errorf("failed to get blacklisted address: %w", err)
	}

	return &blacklistedAddress, nil
}

// PutBlacklistedAddress puts the blacklisted address in the underlying db.
func (s *Store) PutBlacklistedAddress(ctx context.Context, body db.BlacklistedAddress) error {
	dbTx := s.db.WithContext(ctx).Model(&db.BlacklistedAddress{}).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: idName},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				idName, typeName, dataName, addressName, networkName, tagName, remarkName},
			),
		}).Create(&body)
	if dbTx.Error != nil {
		return fmt.Errorf("failed to store blacklisted address: %w", dbTx.Error)
	}

	return nil
}

// UpdateBlacklistedAddress updates the blacklisted address in the underlying db.
func (s *Store) UpdateBlacklistedAddress(ctx context.Context, id string, body db.BlacklistedAddress) error {
	dbTx := s.db.WithContext(ctx).Model(&db.BlacklistedAddress{}).
		Where("id = ?", id).Updates(body)
	if dbTx.Error != nil {
		return fmt.Errorf("failed to update blacklisted address: %w", dbTx.Error)
	}

	return nil
}

// DeleteBlacklistedAddress deletes the blacklisted address from the underlying db.
func (s *Store) DeleteBlacklistedAddress(ctx context.Context, id string) error {
	if dbTx := s.db.WithContext(ctx).Where(
		"id = ?", id).Delete(&db.BlacklistedAddress{}); dbTx.Error != nil {
		return fmt.Errorf("failed to delete blacklisted address: %w", dbTx.Error)
	}
	return nil
}
