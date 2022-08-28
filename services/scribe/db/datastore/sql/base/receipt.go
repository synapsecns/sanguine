package base

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/agents/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

// StoreReceipt stores a receipt.
func (s Store) StoreReceipt(ctx context.Context, receipt types.Receipt, chainID uint32) error {
	dbTx := s.DB().WithContext(ctx).Create(&Receipt{
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

// RetrieveReceipt retrieves a receipt by tx hash and chain id.
func (s Store) RetrieveReceipt(ctx context.Context, txHash common.Hash, chainID uint32) (receipt types.Receipt, err error) {
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
	logs, err := s.RetrieveLogs(ctx, txHash, chainID)
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

// RetrieveAllReceipts_Test retrieves all receipts. Should only be used for testing.
//
//nolint:golint, revive, stylecheck
func (s Store) RetrieveAllReceipts_Test(ctx context.Context, specific bool, chainID uint32, address string) (receipts []types.Receipt, err error) {
	dbReceipts := []Receipt{}
	var dbTx *gorm.DB
	if specific {
		dbTx = s.DB().WithContext(ctx).Model(&Receipt{}).Where(&Receipt{
			ChainID:         chainID,
			ContractAddress: address,
		}).Find(&dbReceipts)
	} else {
		dbTx = s.DB().WithContext(ctx).Model(&Receipt{}).Find(&dbReceipts)
	}

	if dbTx.Error != nil {
		return nil, fmt.Errorf("could not retrieve receipts: %w", dbTx.Error)
	}

	for _, dbReceipt := range dbReceipts {
		// Retrieve Logs that match the receipt's tx hash in order to add them to the Receipt.
		logs, err := s.RetrieveLogs(ctx, common.HexToHash(dbReceipt.TxHash), dbReceipt.ChainID)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve logs with tx hash %s and chain id %d: %w", dbReceipt.TxHash, dbReceipt.ChainID, err)
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
