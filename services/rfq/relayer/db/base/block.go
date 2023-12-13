package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/services/rfq/relayer/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PutLatestBlock upserts the latest block into the database
func (s Store) PutLatestBlock(ctx context.Context, chainID, height uint64) error {
	tx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: chainIDFieldName}},
		DoUpdates: clause.AssignmentColumns([]string{chainIDFieldName, blockNumberFieldName}),
	}).Create(&LastIndexed{
		ChainID:     chainID,
		BlockNumber: int(height),
	})

	if tx.Error != nil {
		return fmt.Errorf("could not block updated: %w", tx.Error)
	}
	return nil
}

func (s Store) LatestBlockForChain(ctx context.Context, chainID uint64) (uint64, error) {
	blockWatchModel := LastIndexed{ChainID: chainID}
	err := s.db.WithContext(ctx).First(&blockWatchModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, db.ErrNoLatestBlockForChainID
		}
		return 0, fmt.Errorf("could not fetch latest block: %w", err)
	}

	return uint64(blockWatchModel.BlockNumber), nil
}
