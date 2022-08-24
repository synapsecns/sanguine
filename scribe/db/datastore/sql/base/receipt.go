package base

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/gorm"
)

// StoreReceipt stores a receipt.
func (s Store) StoreReceipt(ctx context.Context, receipt types.Receipt) error {
	dbTx := s.DB().WithContext(ctx).Create(&Receipt{
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

// RetrieveReceiptByTxHash retrieves a receipt by tx hash.
func (s Store) RetrieveReceiptByTxHash(ctx context.Context, txHash common.Hash) (receipt types.Receipt, err error) {
	dbReceipt := Receipt{}
	dbTx := s.DB().WithContext(ctx).Model(&Receipt{}).Where(&Receipt{
		TxHash: txHash.String(),
	}).First(&dbReceipt)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return types.Receipt{}, fmt.Errorf("could not find receipt with tx hash %s: %w", txHash.String(), dbcommon.ErrNotFound)
		}
		return types.Receipt{}, fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	// Retrieve Logs that match the receipt's tx hash in order to add them to the Receipt.
	logs, err := s.RetrieveLogsByTxHash(ctx, txHash)
	if err != nil {
		return types.Receipt{}, fmt.Errorf("could not retrieve logs with tx hash %s: %w", txHash.String(), err)
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
