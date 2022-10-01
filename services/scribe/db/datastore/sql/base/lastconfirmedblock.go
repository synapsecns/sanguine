//nolint:dupl,golint,revive,stylecheck
package base

import (
	"context"
	"fmt"
)

// StoreLastConfirmedBlock stores the last block number that has been confirmed.
// It updates the value if there is a previous last block confirmed value, and creates a new
// entry if there is no previous value.
func (s Store) StoreLastConfirmedBlock(ctx context.Context, chainID uint32, blockNumber uint64) error {
	entry := LastConfirmedBlockInfo{}
	dbTx := s.DB().WithContext(ctx).
		Model(&LastConfirmedBlockInfo{}).
		Where(&LastConfirmedBlockInfo{
			ChainID: chainID,
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not retrieve last block confirmed info: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		dbTx = s.DB().WithContext(ctx).
			Model(&LastConfirmedBlockInfo{}).
			Create(&LastConfirmedBlockInfo{
				ChainID:     chainID,
				BlockNumber: blockNumber,
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last block confirmed info: %w", dbTx.Error)
		}
		return nil
	}
	dbTx = s.DB().WithContext(ctx).
		Model(&LastConfirmedBlockInfo{}).
		Where(&LastConfirmedBlockInfo{
			ChainID: chainID,
		}).
		Update(BlockNumberFieldName, blockNumber)
	if dbTx.Error != nil {
		return fmt.Errorf("could not update last block confirmed info: %w", dbTx.Error)
	}
	return nil
}

// RetrieveLastConfirmedBlock retrieves the last block number that has been confirmed.
func (s Store) RetrieveLastConfirmedBlock(ctx context.Context, chainID uint32) (uint64, error) {
	entry := LastConfirmedBlockInfo{}
	dbTx := s.DB().WithContext(ctx).
		Model(&LastConfirmedBlockInfo{}).
		Where(&LastConfirmedBlockInfo{
			ChainID: chainID,
		}).
		First(&entry)
	if dbTx.RowsAffected == 0 {
		return 0, nil
	}
	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not retrieve last block confirmed info: %w", dbTx.Error)
	}
	return entry.BlockNumber, nil
}
