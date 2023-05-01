package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"math/big"

	"github.com/synapsecns/sanguine/services/scribe/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreReceipt stores a receipt.
func (s Store) StoreReceipt(ctx context.Context, chainID uint32, receipt types.Receipt) error {
	dbTx := s.DB().WithContext(ctx)
	if s.DB().Dialector.Name() == dbcommon.Sqlite.String() {
		dbTx = dbTx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: TxHashFieldName}, {Name: ChainIDFieldName}},
			DoNothing: true,
		})
	} else {
		dbTx = dbTx.Clauses(clause.Insert{
			Modifier: "IGNORE",
		})
	}
	dbTx = dbTx.Create(&Receipt{
		ChainID:           chainID,
		Type:              receipt.Type,
		PostState:         receipt.PostState,
		Status:            receipt.Status,
		CumulativeGasUsed: receipt.CumulativeGasUsed,
		Bloom:             receipt.Bloom.Bytes(),
		TxHash:            receipt.TxHash.String(),
		ContractAddress:   receipt.ContractAddress.String(),
		GasUsed:           receipt.GasUsed,
		BlockHash:         receipt.BlockHash.String(),
		BlockNumber:       receipt.BlockNumber.Uint64(),
		TransactionIndex:  uint64(receipt.TransactionIndex),
		Confirmed:         false,
	})

	if dbTx.Error != nil {
		return fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	return nil
}

// ConfirmReceiptsForBlockHash confirms receipts for a given block hash.
func (s Store) ConfirmReceiptsForBlockHash(ctx context.Context, chainID uint32, blockHash common.Hash) error {
	dbTx := s.DB().WithContext(ctx).
		Model(&Receipt{}).
		Where(&Receipt{
			ChainID:   chainID,
			BlockHash: blockHash.String(),
		}).
		Update(ConfirmedFieldName, true)

	if dbTx.Error != nil {
		return fmt.Errorf("could not confirm receipt: %w", dbTx.Error)
	}

	return nil
}

// ConfirmReceiptsInRange confirms receipts in a range.
func (s Store) ConfirmReceiptsInRange(ctx context.Context, startBlock, endBlock uint64, chainID uint32) error {
	rangeQuery := fmt.Sprintf("%s BETWEEN ? AND ?", BlockNumberFieldName)
	dbTx := s.DB().WithContext(ctx).
		Model(&Receipt{ChainID: chainID}).
		Order(BlockNumberFieldName+" desc").
		Where(rangeQuery, startBlock, endBlock).
		Update(ConfirmedFieldName, true)

	if dbTx.Error != nil {
		return fmt.Errorf("could not confirm receipts: %w", dbTx.Error)
	}

	return nil
}

// DeleteReceiptsForBlockHash deletes receipts with a given block hash.
func (s Store) DeleteReceiptsForBlockHash(ctx context.Context, chainID uint32, blockHash common.Hash) error {
	dbTx := s.DB().WithContext(ctx).
		Where(&Receipt{
			ChainID:   chainID,
			BlockHash: blockHash.String(),
		}).
		Delete(&Receipt{})

	if dbTx.Error != nil {
		return fmt.Errorf("could not delete receipts: %w", dbTx.Error)
	}

	return nil
}

// receiptFilterToQuery takes in a ReceiptFilter and converts it to a database-type Receipt.
// This is used to query with `WHERE` based on the filter.
func receiptFilterToQuery(receiptFilter db.ReceiptFilter) Receipt {
	return Receipt{
		ChainID:          receiptFilter.ChainID,
		TxHash:           receiptFilter.TxHash,
		ContractAddress:  receiptFilter.ContractAddress,
		BlockHash:        receiptFilter.BlockHash,
		BlockNumber:      receiptFilter.BlockNumber,
		TransactionIndex: receiptFilter.TransactionIndex,
		Confirmed:        receiptFilter.Confirmed,
	}
}

// RetrieveReceiptsWithFilter retrieves receipts with a filter given a page.
func (s Store) RetrieveReceiptsWithFilter(ctx context.Context, receiptFilter db.ReceiptFilter, page int) (receipts []types.Receipt, err error) {
	if page < 1 {
		page = 1
	}
	dbReceipts := []Receipt{}
	query := receiptFilterToQuery(receiptFilter)
	dbTx := s.DB().WithContext(ctx).
		Model(&Receipt{}).
		Where(&query).
		Order(fmt.Sprintf("%s desc, %s desc", BlockNumberFieldName, TransactionIndexFieldName)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
		Find(&dbReceipts)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []types.Receipt{}, fmt.Errorf("could not find receipts with filter %+v: %w", receiptFilter, db.ErrNotFound)
		}
		return []types.Receipt{}, fmt.Errorf("could not retrieve receipts: %w", dbTx.Error)
	}

	parsedReceipts, err := s.buildReceiptsFromDBReceipts(ctx, dbReceipts, receiptFilter.ChainID)
	if err != nil {
		return []types.Receipt{}, fmt.Errorf("could not build receipts from db receipts: %w", err)
	}
	logger.Infof("[RECEIPT QUERY] Retrieved %d receipts with filter %+v", len(parsedReceipts), receiptFilter)
	return parsedReceipts, nil
}

// RetrieveReceiptsInRange retrieves receipts that match an inputted filter and are within a range given a page.
func (s Store) RetrieveReceiptsInRange(ctx context.Context, receiptFilter db.ReceiptFilter, startBlock, endBlock uint64, page int) (receipts []types.Receipt, err error) {
	if page < 1 {
		page = 1
	}
	dbReceipts := []Receipt{}
	query := receiptFilterToQuery(receiptFilter)
	rangeQuery := fmt.Sprintf("%s BETWEEN ? AND ?", BlockNumberFieldName)
	dbTx := s.DB().WithContext(ctx).
		Model(&Receipt{}).
		Where(&query).
		Where(rangeQuery, startBlock, endBlock).
		Order(fmt.Sprintf("%s desc, %s desc", BlockNumberFieldName, TransactionIndexFieldName)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
		Find(&dbReceipts)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []types.Receipt{}, fmt.Errorf("could not find receipts with filter %+v: %w", receiptFilter, db.ErrNotFound)
		}
		return []types.Receipt{}, fmt.Errorf("could not retrieve receipts: %w", dbTx.Error)
	}

	parsedReceipts, err := s.buildReceiptsFromDBReceipts(ctx, dbReceipts, receiptFilter.ChainID)
	if err != nil {
		return []types.Receipt{}, fmt.Errorf("could not build receipts from db receipts: %w", err)
	}

	return parsedReceipts, nil
}

func (s Store) buildReceiptsFromDBReceipts(ctx context.Context, dbReceipts []Receipt, chainID uint32) ([]types.Receipt, error) {
	receipts := []types.Receipt{}
	for i := range dbReceipts {
		dbReceipt := dbReceipts[i]
		// Retrieve Logs that match the receipt's tx hash in order to add them to the Receipt.
		logFilter := db.BuildLogFilter(nil, nil, &dbReceipt.TxHash, nil, nil, nil, nil)
		logFilter.ChainID = chainID

		logs := []*types.Log{}
		page := 1
		for {
			// TODO DELETE
			logger.Infof("[RECEIPT QUERY] logFilter: %v, page: %d", logFilter, page)
			logGroup, err := s.RetrieveLogsWithFilter(ctx, logFilter, page)
			if err != nil {
				return []types.Receipt{}, fmt.Errorf("could not retrieve logs with tx hash %s and chain id %d: %w", dbReceipt.TxHash, chainID, err)
			}
			if len(logGroup) == 0 {
				break
			}
			page++
			logs = append(logs, logGroup...)
		}
		logger.Infof("[RECEIPT QUERY] logs collected: %d, %v, page: %d", len(logs), logFilter, page)

		parsedReceipt := types.Receipt{
			Type:              dbReceipt.Type,
			PostState:         dbReceipt.PostState,
			Status:            dbReceipt.Status,
			CumulativeGasUsed: dbReceipt.CumulativeGasUsed,
			Bloom:             types.BytesToBloom(dbReceipt.Bloom),
			Logs:              logs,
			TxHash:            common.HexToHash(dbReceipt.TxHash),
			ContractAddress:   common.HexToAddress(dbReceipt.ContractAddress),
			GasUsed:           dbReceipt.GasUsed,
			BlockHash:         common.HexToHash(dbReceipt.BlockHash),
			BlockNumber:       big.NewInt(int64(dbReceipt.BlockNumber)),
			TransactionIndex:  uint(dbReceipt.TransactionIndex),
		}
		logger.Infof("[RECEIPT QUERY] parsedReceipt:, %v", parsedReceipt)

		receipts = append(receipts, parsedReceipt)
	}
	logger.Infof("[RECEIPT QUERY] parsedReceipt: %d", len(receipts))

	return receipts, nil
}

// RetrieveReceiptCountForChain retrieves the count of receipts per chain.
func (s Store) RetrieveReceiptCountForChain(ctx context.Context, chainID uint32) (int64, error) {
	var count int64
	dbTx := s.DB().WithContext(ctx).
		Model(&Receipt{}).
		Where(&Receipt{ChainID: chainID}).
		Count(&count)

	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not count receipts: %w", dbTx.Error)
	}

	return count, nil
}
