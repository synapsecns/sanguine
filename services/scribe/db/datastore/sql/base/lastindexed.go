package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// StoreLastIndexed stores the last indexed block number for a contract.
// It updates the value if there is a previous last indexed value, and creates a new
// entry if there is no previous value.
func (s Store) StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error {
	entry := LastIndexedInfo{}
	dbTx := s.DB().WithContext(ctx).
		Model(&LastIndexedInfo{}).
		Where(&LastIndexedInfo{
			ContractAddress: contractAddress.String(),
			ChainID:         chainID,
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not retrieve last indexed info: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		dbTx = s.DB().WithContext(ctx).
			Model(&LastIndexedInfo{}).
			Create(&LastIndexedInfo{
				ContractAddress: contractAddress.String(),
				ChainID:         chainID,
				BlockNumber:     blockNumber,
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last indexed info: %w", dbTx.Error)
		}
		return nil
	}
	dbTx = s.DB().WithContext(ctx).
		Model(&LastIndexedInfo{}).
		Where(&LastIndexedInfo{
			ContractAddress: contractAddress.String(),
			ChainID:         chainID,
		}).
		Update(BlockNumberFieldName, blockNumber)
	if dbTx.Error != nil {
		return fmt.Errorf("could not update last indexed info: %w", dbTx.Error)
	}
	return nil
}

// RetrieveLastIndexed retrieves the last indexed block number for a contract.
func (s Store) RetrieveLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32) (uint64, error) {
	entry := LastIndexedInfo{}
	dbTx := s.DB().WithContext(ctx).
		Model(&LastIndexedInfo{}).
		Where(&LastIndexedInfo{
			ContractAddress: contractAddress.String(),
			ChainID:         chainID,
		}).
		First(&entry)
	if dbTx.RowsAffected == 0 {
		logger.Warnf("no last indexed info found for contract %s on chain %d. Providing 0.", contractAddress.String(), chainID)
		return 0, nil
	}
	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not retrieve last indexed info: %w", dbTx.Error)
	}
	return entry.BlockNumber, nil
}
