package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/scribe/types"
	"gorm.io/gorm"
)

// StoreReceipt stores a receipt.
func (s Store) StoreReceipt(ctx context.Context, receipt ethTypes.Receipt) error {
	dbTx := s.DB().WithContext(ctx).Create(&Receipt{
		Status:            receipt.Status,
		CumulativeGasUsed: receipt.CumulativeGasUsed,
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
			return nil, fmt.Errorf("could not find receipt with tx hash %s: %w", txHash.String(), dbcommon.ErrNotFound)
		}
		return nil, fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	parsedReceipt := types.NewReceipt(
		dbReceipt.Status,
		dbReceipt.CumulativeGasUsed,
		common.HexToHash(dbReceipt.TxHash),
		common.HexToAddress(dbReceipt.ContractAddress),
		dbReceipt.GasUsed,
		common.HexToHash(dbReceipt.BlockHash),
		dbReceipt.BlockNumber,
		dbReceipt.TransactionIndex,
	)

	return parsedReceipt, nil
}
