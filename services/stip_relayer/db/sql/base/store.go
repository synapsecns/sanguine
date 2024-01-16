package base

import (
	"context"

	"github.com/synapsecns/sanguine/services/stip_relayer/db"
)

// Write some queries here

// GetSTIPTransactionsNotRebated gets transactions that have not yet been rebated
func (s *Store) GetSTIPTransactionsNotRebated(ctx context.Context) ([]*db.STIPTransactions, error) {
	var stipTransactions []*db.STIPTransactions

	result := s.db.WithContext(ctx).Where("rebated = ?", false).Find(&stipTransactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return stipTransactions, nil
}

// UpdateSTIPTransactionRebated updates the rebated status of a transaction
func (s *Store) UpdateSTIPTransactionRebated(ctx context.Context, hash string) error {
	result := s.db.WithContext(ctx).Model(&db.STIPTransactions{}).Where("hash = ?", hash).Update("rebated", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// InsertNewStipTransactions inserts new transactions into the database
func (s *Store) InsertNewStipTransactions(ctx context.Context, stipTransactions []db.STIPTransactions) error {
	result := s.db.WithContext(ctx).Create(&stipTransactions)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
