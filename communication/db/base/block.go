package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/communication/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// LastIndexed is used to make sure we haven't missed any events while offline.
// since we event source - rather than use a state machine this is needed to make sure we haven't missed any events
// by allowing us to go back and source any events we may have missed.
//
// this does not inherit from gorm.model to allow us to use ChainID as a primary key.
type LastIndexed struct {
	// CreatedAt is the creation time
	CreatedAt time.Time
	// UpdatedAt is the update time
	UpdatedAt time.Time
	// DeletedAt time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// ChainID is the chain id of the chain we're watching blocks on. This is our primary index.
	ChainID uint64 `gorm:"column:chain_id;primaryKey;autoIncrement:false"`
	// BlockHeight is the highest height we've seen on the chain
	BlockNumber int `gorm:"block_number"`
}

// PutLatestBlock upserts the latest block into the database.
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

// LatestBlockForChain gets the latest block for a chain.
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
