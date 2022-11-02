package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/types"
	"gorm.io/gorm"
)

// StoreFailedLog checks if the failed log exists in the database. If so, increment its FailedAttempt counter,
// otherwise, add it to the table with a FailedAttempt value of one.
func (s Store) StoreFailedLog(ctx context.Context, chainID uint32, contractAddress common.Address, txHash common.Hash, blockIndex uint64, blockNumber uint64) error {
	entry := FailedLog{}
	dbTx := s.DB().WithContext(ctx).
		Model(&FailedLog{}).
		Where(&FailedLog{
			types.FailedLog{
				ChainID:         chainID,
				ContractAddress: contractAddress.String(),
				TxHash:          txHash.String(),
				BlockIndex:      blockIndex,
				BlockNumber:     blockNumber,
			},
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not scan failed log: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		dbTx = s.DB().WithContext(ctx).
			Create(&FailedLog{
				types.FailedLog{
					ChainID:         chainID,
					ContractAddress: contractAddress.String(),
					TxHash:          txHash.String(),
					BlockIndex:      blockIndex,
					BlockNumber:     blockNumber,
					FailedAttempts:  1,
				},
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not create failed log: %w", dbTx.Error)
		}
		return nil
	}
	dbTx = s.DB().WithContext(ctx).
		Model(&FailedLog{}).
		Where(&FailedLog{
			types.FailedLog{
				ChainID:         chainID,
				ContractAddress: contractAddress.String(),
				TxHash:          txHash.String(),
				BlockIndex:      blockIndex,
				BlockNumber:     blockNumber,
			},
		}).
		Update(FailedAttemptsFieldName, gorm.Expr(fmt.Sprintf("%s + 1", FailedAttemptsFieldName)))
	if dbTx.Error != nil {
		return fmt.Errorf("could not update failed log: %w", dbTx.Error)
	}
	return nil
}

// DeleteFailedLog deletes a failed log from the FailedLog table.
func (s Store) DeleteFailedLog(ctx context.Context, chainID uint32, contractAddress common.Address, txHash common.Hash, blockIndex uint64, blockNumber uint64) error {
	entry := FailedLog{}
	dbTx := s.DB().WithContext(ctx).
		Model(&FailedLog{}).
		Where(&FailedLog{
			types.FailedLog{
				ChainID:         chainID,
				ContractAddress: contractAddress.String(),
				TxHash:          txHash.String(),
				BlockIndex:      blockIndex,
				BlockNumber:     blockNumber,
			},
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not scan failed log: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		// If the failed log does not exist, return nil.
		return nil
	}
	dbTx = s.DB().WithContext(ctx).
		Where(&FailedLog{
			types.FailedLog{
				ChainID:         chainID,
				ContractAddress: contractAddress.String(),
				TxHash:          txHash.String(),
				BlockIndex:      blockIndex,
				BlockNumber:     blockNumber,
			},
		}).
		Delete(&FailedLog{})
	if dbTx.Error != nil {
		return fmt.Errorf("could not delete failed log: %w", dbTx.Error)
	}
	return nil
}

// GetFailedAttempts returns the number of failed attempts for a failed log.
func (s Store) GetFailedAttempts(ctx context.Context, chainID uint32, contractAddress common.Address, txHash common.Hash, blockIndex uint64, blockNumber uint64) (uint64, error) {
	entry := FailedLog{}
	dbTx := s.DB().WithContext(ctx).
		Model(&FailedLog{}).
		Where(&FailedLog{
			types.FailedLog{
				ChainID:         chainID,
				ContractAddress: contractAddress.String(),
				TxHash:          txHash.String(),
				BlockIndex:      blockIndex,
				BlockNumber:     blockNumber,
			},
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not scan failed log: %w", dbTx.Error)
	}
	return entry.FailedAttempts, nil
}

// failedLogFilterToQuery converts an failedLogFilter to a database-type FailedLog.
// This is used to query with `WHERE` based on the filter.
func failedLogFilterToQuery(failedLogFilter db.FailedLogFilter) FailedLog {
	return FailedLog{
		types.FailedLog{
			ChainID:         failedLogFilter.ChainID,
			ContractAddress: failedLogFilter.ContractAddress,
			TxHash:          failedLogFilter.TxHash,
			BlockIndex:      failedLogFilter.BlockIndex,
			BlockNumber:     failedLogFilter.BlockNumber,
		},
	}
}

// GetFailedLogsFromFilter returns a list of failed logs that match the given filter.
func (s Store) GetFailedLogsFromFilter(ctx context.Context, failedLogFilter db.FailedLogFilter) ([]*types.FailedLog, error) {
	var failedLogs []*FailedLog
	query := failedLogFilterToQuery(failedLogFilter)
	dbTx := s.DB().WithContext(ctx).
		Model(&FailedLog{}).
		Where(&query).
		Scan(&failedLogs)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("could not scan failed logs: %w", dbTx.Error)
	}

	var failedLogsTypes []*types.FailedLog
	for _, failedLog := range failedLogs {
		storeLog := failedLog.FailedLog
		failedLogsTypes = append(failedLogsTypes, &storeLog)
	}
	return failedLogsTypes, nil
}
