package base

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/services/scribe/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreReceipt stores a receipt.
func (s Store) StoreReceipt(ctx context.Context, receipt types.Receipt, chainID uint32) error {
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: TxHashFieldName}, {Name: ChainIDFieldName}},
			DoNothing: true,
		}).
		Create(&Receipt{
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
		})

	if dbTx.Error != nil {
		return fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	return nil
}

func receiptFilterToQuery(receiptFilter db.ReceiptFilter) Receipt {
	return Receipt{
		ChainID:          receiptFilter.ChainID,
		TxHash:           receiptFilter.TxHash,
		ContractAddress:  receiptFilter.ContractAddress,
		BlockHash:        receiptFilter.BlockHash,
		BlockNumber:      receiptFilter.BlockNumber,
		TransactionIndex: receiptFilter.TransactionIndex,
	}
}

// RetrieveReceiptsWithFilter retrieves receipts with a filter.
func (s Store) RetrieveReceiptsWithFilter(ctx context.Context, receiptFilter db.ReceiptFilter) (receipts []types.Receipt, err error) {
	dbReceipts := []Receipt{}
	query := receiptFilterToQuery(receiptFilter)
	dbTx := s.DB().WithContext(ctx).Model(&Receipt{}).Where(&query).Find(&dbReceipts)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []types.Receipt{}, fmt.Errorf("could not find receipts with filter %+v: %w", receiptFilter, db.ErrNotFound)
		}
		return []types.Receipt{}, fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	parsedReceipts, err := s.buildReceiptsFromDBReceipts(ctx, dbReceipts, receiptFilter.ChainID)
	if err != nil {
		return []types.Receipt{}, fmt.Errorf("could not build receipts from db receipts: %w", err)
	}

	return parsedReceipts, nil
}

// RetrieveReceiptsInRange retrieves receipts in a range.
func (s Store) RetrieveReceiptsInRange(ctx context.Context, receiptFilter db.ReceiptFilter, startBlock, endBlock uint64) (receipts []types.Receipt, err error) {
	dbReceipts := []Receipt{}
	query := receiptFilterToQuery(receiptFilter)
	rangeQuery := BlockNumberFieldName + " BETWEEN ? AND ?"
	dbTx := s.DB().WithContext(ctx).Model(&Receipt{}).Where(&query).Where(rangeQuery, startBlock, endBlock).Find(&dbReceipts)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []types.Receipt{}, fmt.Errorf("could not find receipts with filter %+v: %w", receiptFilter, db.ErrNotFound)
		}
		return []types.Receipt{}, fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	parsedReceipts, err := s.buildReceiptsFromDBReceipts(ctx, dbReceipts, receiptFilter.ChainID)
	if err != nil {
		return []types.Receipt{}, fmt.Errorf("could not build receipts from db receipts: %w", err)
	}

	return parsedReceipts, nil
}

// RetrieveReceiptByTxHash retrieves a receipt by tx hash and chain id.
func (s Store) RetrieveReceiptByTxHash(ctx context.Context, txHash common.Hash, chainID uint32) (receipt types.Receipt, err error) {
	dbReceipt := Receipt{}
	dbTx := s.DB().WithContext(ctx).Model(&Receipt{}).Where(&Receipt{
		ChainID: chainID,
		TxHash:  txHash.String(),
	}).First(&dbReceipt)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return types.Receipt{}, fmt.Errorf("could not find receipt with tx hash %s: %w", txHash.String(), db.ErrNotFound)
		}
		return types.Receipt{}, fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	// Retrieve Logs that match the receipt's tx hash in order to add them to the Receipt.
	logs, err := s.RetrieveLogsByTxHash(ctx, txHash, chainID)
	if err != nil {
		return types.Receipt{}, fmt.Errorf("could not retrieve logs with tx hash %s and chain id %d: %w", txHash.String(), chainID, err)
	}

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

	return parsedReceipt, nil
}

// RetrieveReceiptsByContractAddress retrieves all receipts with a given contract address and chain id.
func (s Store) RetrieveReceiptsByContractAddress(ctx context.Context, contractAddress common.Address, chainID uint32) (receipts []types.Receipt, err error) {
	dbReceipts := []Receipt{}
	dbTx := s.DB().WithContext(ctx).Model(&Receipt{}).Where(&Receipt{
		ChainID:         chainID,
		ContractAddress: contractAddress.String(),
	}).Find(&dbReceipts)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []types.Receipt{}, fmt.Errorf("could not find receipts with contract address %s: %w", contractAddress.String(), db.ErrNotFound)
		}
		return []types.Receipt{}, fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	parsedReceipts, err := s.buildReceiptsFromDBReceipts(ctx, dbReceipts, chainID)
	if err != nil {
		return []types.Receipt{}, fmt.Errorf("could not build receipts from db receipts: %w", err)
	}

	return parsedReceipts, nil
}

// UnsafeRetrieveAllReceipts retrieves all receipts in the database. When `specific` is true, you can specify
// a chainID to specifically search for. This is only used for testing.
func (s Store) UnsafeRetrieveAllReceipts(ctx context.Context, specific bool, chainID uint32) (receipts []*types.Receipt, err error) {
	dbReceipts := []Receipt{}
	var dbTx *gorm.DB
	if specific {
		dbTx = s.DB().WithContext(ctx).Model(&Receipt{}).Where(&Receipt{
			ChainID: chainID,
		}).Find(&dbReceipts)
	} else {
		dbTx = s.DB().WithContext(ctx).Model(&Receipt{}).Find(&dbReceipts)
	}

	if dbTx.Error != nil {
		return nil, fmt.Errorf("could not retrieve receipts: %w", dbTx.Error)
	}

	for _, dbReceipt := range dbReceipts {
		// Retrieve Logs that match the receipt's tx hash in order to add them to the Receipt.
		logs, err := s.RetrieveLogsByTxHash(ctx, common.HexToHash(dbReceipt.TxHash), dbReceipt.ChainID)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve logs with tx hash %s and chain id %d: %w", dbReceipt.TxHash, dbReceipt.ChainID, err)
		}

		parsedReceipt := &types.Receipt{
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

		receipts = append(receipts, parsedReceipt)
	}

	return receipts, nil
}

func (s Store) buildReceiptsFromDBReceipts(ctx context.Context, dbReceipts []Receipt, chainID uint32) ([]types.Receipt, error) {
	receipts := []types.Receipt{}
	for _, dbReceipt := range dbReceipts {
		// Retrieve Logs that match the receipt's tx hash in order to add them to the Receipt.
		logs, err := s.RetrieveLogsByTxHash(ctx, common.HexToHash(dbReceipt.TxHash), chainID)
		if err != nil {
			return []types.Receipt{}, fmt.Errorf("could not retrieve logs with tx hash %s and chain id %d: %w", dbReceipt.TxHash, chainID, err)
		}

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

		receipts = append(receipts, parsedReceipt)
	}

	return receipts, nil
}
