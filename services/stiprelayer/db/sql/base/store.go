package base

import (
	"context"

	"github.com/synapsecns/sanguine/services/stiprelayer/db"
	"gorm.io/gorm/clause"
)

// Write some queries here

// GetSTIPTransactionsNotRebated gets transactions that have not yet been rebated.
func (s *Store) GetSTIPTransactionsNotRebated(ctx context.Context) ([]*db.STIPTransactions, error) {
	var stipTransactions []*db.STIPTransactions

	result := s.db.WithContext(ctx).Where("rebated = ?", false).Where("do_not_process = ?", false).Find(&stipTransactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return stipTransactions, nil
}

// UpdateSTIPTransactionRebated updates the rebated status of a transaction.
func (s *Store) UpdateSTIPTransactionRebated(ctx context.Context, hash string, nonce uint64, arbTransferAmount string) error {
	result := s.db.WithContext(ctx).Model(&db.STIPTransactions{}).Where("hash = ?", hash).Update("rebated", true).Update("nonce", nonce).Update("arb_amount_rebated", arbTransferAmount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateSTIPTransactionDoNotProcess updates the rebated status of a transaction.
func (s *Store) UpdateSTIPTransactionDoNotProcess(ctx context.Context, hash string) error {
	result := s.db.WithContext(ctx).Model(&db.STIPTransactions{}).Where("hash = ?", hash).Update("do_not_process", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// InsertNewStipTransactions inserts new transactions into the database.
func (s *Store) InsertNewStipTransactions(ctx context.Context, stipTransactions []db.STIPTransactions) error {
	batchSize := 50 // Adjust batch size based on your DB's performance and limitations

	for i := 0; i < len(stipTransactions); i += batchSize {
		end := i + batchSize
		if end > len(stipTransactions) {
			end = len(stipTransactions)
		}
		batch := stipTransactions[i:end]

		// Using CreateInBatches with ON CONFLICT clause to ignore duplicates based on the hash
		tx := s.db.WithContext(ctx).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "hash"}}, // Conflict detection based on the hash column
			DoNothing: true,                            // In case of conflict, do nothing
		}).CreateInBatches(batch, len(batch))

		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}
