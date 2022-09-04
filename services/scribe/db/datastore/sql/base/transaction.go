package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreEthTx stores a processed text.
func (s Store) StoreEthTx(ctx context.Context, tx *types.Transaction, chainID uint32, blockNumber uint64) error {
	marshalledTx, err := tx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("could not marshall tx to binary: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: TxHashFieldName}, {Name: ChainIDFieldName}},
			DoNothing: true,
		}).
		Create(&EthTx{
			TxHash:      tx.Hash().String(),
			ChainID:     chainID,
			BlockNumber: blockNumber,
			RawTx:       marshalledTx,
			GasFeeCap:   tx.GasFeeCap().Uint64(),
			GasTipCap:   tx.GasTipCap().Uint64(),
		})

	if dbTx.Error != nil {
		return fmt.Errorf("could not create raw tx: %w", dbTx.Error)
	}

	return nil
}

func ethTxFilterToQuery(ethTxFilter db.EthTxFilter) EthTx {
	return EthTx{
		ChainID:     ethTxFilter.ChainID,
		TxHash:      ethTxFilter.TxHash,
		BlockNumber: ethTxFilter.BlockNumber,
	}
}

// RetrieveEthTxsWithFilter retrieves eth transactions with a filter.
func (s Store) RetrieveEthTxsWithFilter(ctx context.Context, ethTxFilter db.EthTxFilter) ([]types.Transaction, error) {
	dbEthTxs := []EthTx{}
	query := ethTxFilterToQuery(ethTxFilter)
	dbTx := s.DB().WithContext(ctx).
		Model(&EthTx{}).
		Where(&query).
		Find(&dbEthTxs)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []types.Transaction{}, fmt.Errorf("could not find eth txs with filter %+v: %w", ethTxFilter, db.ErrNotFound)
		}
		return []types.Transaction{}, fmt.Errorf("could not retrieve eth txs: %w", dbTx.Error)
	}

	parsedEthTxs, err := buildEthTxsFromDBEthTxs(dbEthTxs)
	if err != nil {
		return []types.Transaction{}, fmt.Errorf("could not build eth txs: %w", err)
	}

	return parsedEthTxs, nil
}

// RetrieveEthTxsInRange retrieves eth transactions in a range.
func (s Store) RetrieveEthTxsInRange(ctx context.Context, ethTxFilter db.EthTxFilter, startBlock, endBlock uint64) ([]types.Transaction, error) {
	dbEthTxs := []EthTx{}
	query := ethTxFilterToQuery(ethTxFilter)
	rangeQuery := BlockNumberFieldName + " BETWEEN ? AND ?"
	dbTx := s.DB().WithContext(ctx).
		Model(&EthTx{}).
		Where(&query).
		Where(rangeQuery, startBlock, endBlock).
		Find(&dbEthTxs)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []types.Transaction{}, fmt.Errorf("could not find eth txs with filter %+v: %w", ethTxFilter, db.ErrNotFound)
		}
		return []types.Transaction{}, fmt.Errorf("could not retrieve eth txs: %w", dbTx.Error)
	}

	parsedEthTxs, err := buildEthTxsFromDBEthTxs(dbEthTxs)
	if err != nil {
		return []types.Transaction{}, fmt.Errorf("could not build eth txs: %w", err)
	}

	return parsedEthTxs, nil
}

// RetrieveEthTxByTxHash retrieves a processed transaction by tx hash and chain id.
func (s Store) RetrieveEthTxByTxHash(ctx context.Context, txHash string, chainID uint32) (types.Transaction, error) {
	dbEthTx := EthTx{}
	dbTx := s.DB().WithContext(ctx).Model(&EthTx{}).Where(&EthTx{
		ChainID: chainID,
		TxHash:  txHash,
	}).First(&dbEthTx)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return types.Transaction{}, fmt.Errorf("could not find raw tx with tx hash %s: %w", txHash, db.ErrNotFound)
		}
		return types.Transaction{}, fmt.Errorf("could not retrieve raw tx: %w", dbTx.Error)
	}

	tx := types.Transaction{}
	if err := tx.UnmarshalBinary(dbEthTx.RawTx); err != nil {
		return types.Transaction{}, fmt.Errorf("could not unmarshall raw tx: %w", err)
	}

	return tx, nil
}

func buildEthTxsFromDBEthTxs(dbEthTxs []EthTx) ([]types.Transaction, error) {
	ethTxs := []types.Transaction{}
	for _, dbEthTx := range dbEthTxs {
		ethTx := types.Transaction{}
		if err := ethTx.UnmarshalBinary(dbEthTx.RawTx); err != nil {
			return []types.Transaction{}, fmt.Errorf("could not unmarshall eth tx: %w", err)
		}
		ethTxs = append(ethTxs, ethTx)
	}

	return ethTxs, nil
}
