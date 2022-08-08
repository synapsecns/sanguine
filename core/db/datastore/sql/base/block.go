package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreMessageLatestBlockEnd stores the latest message block height we've observed.
func (s Store) StoreMessageLatestBlockEnd(ctx context.Context, domainID uint32, blockNumber uint32) error {
	tx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: DomainIDFieldName}},
		DoUpdates: clause.AssignmentColumns([]string{DomainIDFieldName, BlockNumberFieldName}),
	}).Create(&BlockEndModel{
		DomainID:    domainID,
		BlockNumber: blockNumber,
	})

	if tx.Error != nil {
		return fmt.Errorf("could not block updated: %w", tx.Error)
	}
	return nil
}

// GetMessageLatestBlockEnd gets the latest block end for a given domain.
func (s Store) GetMessageLatestBlockEnd(ctx context.Context, domainID uint32) (blockNumber uint32, err error) {
	blockWatchModel := BlockEndModel{DomainID: domainID}
	err = s.db.WithContext(ctx).First(&blockWatchModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, db.ErrNoStoredBlockForChain
		}
		return 0, fmt.Errorf("could not fetch latest block: %w", err)
	}

	return blockWatchModel.BlockNumber, nil
}
