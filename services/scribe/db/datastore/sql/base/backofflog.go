package base

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm/clause"
)

// StoreBackoffLog stores a backoff log.
func (s Store) StoreBackoffLog(ctx context.Context, chainID uint32, contractAddress sql.NullString, blockNumber uint64, backoffType string, backoffAttempt uint32) error {
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: TxHashFieldName}, {Name: ChainIDFieldName}},
			DoNothing: true,
		}).Create(&BackoffLog{
		ChainID:         chainID,
		ContractAddress: contractAddress,
		BlockNumber:     blockNumber,
		BackoffType:     backoffType,
		BackoffAttempt:  backoffAttempt,
	})

	if dbTx.Error != nil {
		return fmt.Errorf("could not store backoff log: %w", dbTx.Error)
	}

	return nil
}

// GetBackoffLogs gets backoff logs matching chain id and backoff type.
func (s Store) GetBackoffLogs(ctx context.Context, chainID uint32, backoffType string) ([]BackoffLog, error) {
	var entries []BackoffLog
	dbTx := s.DB().WithContext(ctx).
		Model(&BackoffLog{}).
		Where(&BackoffLog{
			ChainID:     chainID,
			BackoffType: backoffType,
		}).
		Scan(&entries)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("could not get backoff logs: %w", dbTx.Error)
	}
	return entries, nil
}

// StoreBackoffCount stores a backoff count.
func (s Store) StoreBackoffCount(ctx context.Context, chainID uint32, backoffType string, count uint64) error {
	entry := BackoffCount{}
	dbTx := s.DB().WithContext(ctx).
		Model(&BackoffCount{}).
		Where(&BackoffCount{
			ChainID:     chainID,
			BackoffType: backoffType,
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not get backoff count: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		dbTx = s.DB().WithContext(ctx).
			Model(&BackoffCount{}).
			Create(&BackoffCount{
				ChainID:     chainID,
				BackoffType: backoffType,
				Count:       count,
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not store backoff count: %w", dbTx.Error)
		}
		return nil
	}
	dbTx = s.DB().WithContext(ctx).
		Model(&BackoffCount{}).
		Where(&BackoffCount{
			ChainID:     chainID,
			BackoffType: backoffType,
		}).
		Update(CountFieldName, count)
	if dbTx.Error != nil {
		return fmt.Errorf("could not update backoff count: %w", dbTx.Error)
	}
	return nil
}

// GetBackoffCount gets a backoff count.
func (s Store) GetBackoffCount(ctx context.Context, chainID uint32, backoffType string) (uint64, error) {
	entry := BackoffCount{}
	dbTx := s.DB().WithContext(ctx).
		Model(&BackoffCount{}).
		Where(&BackoffCount{
			ChainID:     chainID,
			BackoffType: backoffType,
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not get backoff count: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		return 0, nil
	}
	return entry.Count, nil
}
