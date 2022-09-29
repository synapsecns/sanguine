// nolint:dupl,golint,revive,stylecheck
package base

import (
	"context"
	"fmt"
)

// StoreLastBlockTime stores the last block time stored for a chain.
// It updates the value if there is a previous last indexed value, and creates a new
// entry if there is no previous value.
// TODO: Potentially use generics since the logic is identical to that of `lastindexed.go`.
func (s Store) StoreLastBlockTime(ctx context.Context, chainID uint32, blockNumber uint64) error {
	entry := LastBlockTime{}
	dbTx := s.DB().WithContext(ctx).
		Model(&LastBlockTime{}).
		Where(&LastBlockTime{
			ChainID: chainID,
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not retrieve last block time: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		dbTx = s.DB().WithContext(ctx).
			Model(&LastBlockTime{}).
			Create(&LastBlockTime{
				ChainID:     chainID,
				BlockNumber: blockNumber,
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last block time: %w", dbTx.Error)
		}
		return nil
	}
	dbTx = s.DB().WithContext(ctx).
		Model(&LastBlockTime{}).
		Where(&LastBlockTime{
			ChainID: chainID,
		}).
		Update(BlockNumberFieldName, blockNumber)
	if dbTx.Error != nil {
		return fmt.Errorf("could not update last block time: %w", dbTx.Error)
	}
	return nil
}

// RetrieveLastBlockTime retrieves the last block time stored for a chain.
func (s Store) RetrieveLastBlockTime(ctx context.Context, chainID uint32) (uint64, error) {
	entry := LastBlockTime{}
	dbTx := s.DB().WithContext(ctx).
		Model(&LastBlockTime{}).
		Where(&LastBlockTime{
			ChainID: chainID,
		}).
		First(&entry)
	if dbTx.RowsAffected == 0 {
		return 0, nil
	}
	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not retrieve last block time: %w", dbTx.Error)
	}
	return entry.BlockNumber, nil
}
