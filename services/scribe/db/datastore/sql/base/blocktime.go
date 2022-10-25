package base

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"
)

// StoreBlockTime stores a block time for a chain.
func (s Store) StoreBlockTime(ctx context.Context, chainID uint32, blockNumber, timestamp uint64) error {
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: ChainIDFieldName}, {Name: BlockNumberFieldName}},
			DoNothing: true,
		}).
		Model(&BlockTime{}).
		Create(&BlockTime{
			ChainID:     chainID,
			BlockNumber: blockNumber,
			Timestamp:   timestamp,
		})
	if dbTx.Error != nil {
		return fmt.Errorf("could not store block time: %w", dbTx.Error)
	}

	return nil
}

// RetrieveBlockTime retrieves a block time for a chain and block number.
func (s Store) RetrieveBlockTime(ctx context.Context, chainID uint32, blockNumber uint64) (uint64, error) {
	var blockTime BlockTime
	dbTx := s.DB().WithContext(ctx).
		Model(&BlockTime{}).
		Where(&BlockTime{
			ChainID:     chainID,
			BlockNumber: blockNumber,
		}).
		First(&blockTime)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not retrieve block time: %w", dbTx.Error)
	}

	return blockTime.Timestamp, nil
}

// RetrieveLastBlockStored retrieves the last block number that has a stored block time.
func (s Store) RetrieveLastBlockStored(ctx context.Context, chainID uint32) (uint64, error) {
	var blockTime int64
	row := s.DB().WithContext(ctx).
		Model(&BlockTime{}).
		Where(&BlockTime{
			ChainID: chainID,
		}).
		Select(fmt.Sprintf("max(%s)", BlockNumberFieldName)).Row()
	err := row.Scan(&blockTime)
	if err != nil {
		return 0, fmt.Errorf("could not retrieve last block time: %w", err)
	}

	return uint64(blockTime), nil
}

// RetrieveFirstBlockStored retrieves the first block number that has a stored block time.
func (s Store) RetrieveFirstBlockStored(ctx context.Context, chainID uint32) (uint64, error) {
	var blockTime int64
	row := s.DB().WithContext(ctx).
		Model(&BlockTime{}).
		Where(&BlockTime{
			ChainID: chainID,
		}).
		Select(fmt.Sprintf("min(%s)", BlockNumberFieldName)).Row()
	err := row.Scan(&blockTime)
	if err != nil {
		return 0, fmt.Errorf("could not retrieve first block time: %w", err)
	}

	return uint64(blockTime), nil
}
