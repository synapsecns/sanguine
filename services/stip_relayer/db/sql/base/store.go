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
func (s *Store) UpdateSTIPTransactionRebated(ctx context.Context, hash string, nonce uint64) error {
	result := s.db.WithContext(ctx).Model(&db.STIPTransactions{}).Where("hash = ?", hash).Update("rebated", true).Update("nonce", nonce)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// InsertNewStipTransactions inserts new transactions into the database
func (s *Store) InsertNewStipTransactions(ctx context.Context, stipTransactions []db.STIPTransactions) error {
	for _, transaction := range stipTransactions {
		// Using FirstOrCreate
		tx := s.db.WithContext(ctx).Where(db.STIPTransactions{Hash: transaction.Hash}).FirstOrCreate(&transaction)
		if tx.Error != nil {
			return tx.Error
		}

		// Or using Clauses with ON CONFLICT for upsert behavior (PostgreSQL example)
		// tx := s.db.WithContext(ctx).Clauses(clause.OnConflict{
		//     UpdateAll: true,
		// }).Create(&transaction)
		// if tx.Error != nil {
		//     return tx.Error
		// }
	}
	return nil
}
