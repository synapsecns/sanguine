package base

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm/clause"

	"github.com/synapsecns/sanguine/services/cctp-relayer/types"
)

// GetLastBlockNumber gets the last block number that had a message in the database.
func (s Store) GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	var message types.Message
	var lastBlockNumber sql.NullInt64

	dbTx := s.DB().WithContext(ctx).
		Model(&message).
		Where(fmt.Sprintf("%s = ?", OriginChainIDFieldName), chainID).
		Select(fmt.Sprintf("MAX(%s)", BlockNumberFieldName)).
		Find(&lastBlockNumber)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get last block number: %w", dbTx.Error)
	}

	// explicitly return 0 on nil
	if !lastBlockNumber.Valid {
		return 0, nil
	}

	return uint64(lastBlockNumber.Int64), nil
}

// StoreMessage stores a message in the database.
func (s Store) StoreMessage(ctx context.Context, msg types.Message) error {
	dbTx := s.DB().WithContext(ctx).Clauses(
		clause.OnConflict{DoNothing: true}).Create(&msg)
	if dbTx.Error != nil {
		return fmt.Errorf("failed to store message: %w", dbTx.Error)
	}
	return nil
}
