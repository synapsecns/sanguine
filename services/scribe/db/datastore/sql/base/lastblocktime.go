// nolint:dupl,golint,revive,stylecheck
package base

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/model"
	"time"
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

func (s Store) RetrieveLastBlockStoredVerbose(ctx context.Context, chainID uint32) (*model.Block, error) {
	entry := LastBlockTime{}
	dbTx := s.DB().WithContext(ctx).
		Model(&LastBlockTime{}).
		Where(&LastBlockTime{
			ChainID: chainID,
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("could not retrieve last block time: %w", dbTx.Error)
	}
	output := &model.Block{
		ChainID:     int(entry.ChainID),
		BlockNumber: int(entry.BlockNumber),
		CreatedAt:   entry.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   entry.UpdatedAt.Format(time.RFC3339),
	}
	return output, nil
}
