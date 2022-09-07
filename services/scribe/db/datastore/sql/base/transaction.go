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

// ethTxFilterToQuery converts an ethTxFilter to a database-type EthTx.
// This is used to query with `WHERE` based on the filter.
func ethTxFilterToQuery(ethTxFilter db.EthTxFilter) EthTx {
	return EthTx{
		ChainID:     ethTxFilter.ChainID,
		TxHash:      ethTxFilter.TxHash,
		BlockNumber: ethTxFilter.BlockNumber,
	}
}

// RetrieveEthTxsWithFilter retrieves eth transactions with a filter given a page.
func (s Store) RetrieveEthTxsWithFilter(ctx context.Context, ethTxFilter db.EthTxFilter, page int) ([]types.Transaction, error) {
	if page < 1 {
		page = 1
	}
	dbEthTxs := []EthTx{}
	query := ethTxFilterToQuery(ethTxFilter)
	dbTx := s.DB().WithContext(ctx).
		Model(&EthTx{}).
		Where(&query).
		Order(fmt.Sprintf("%s, %s", BlockIndexFieldName, BlockIndexFieldName)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
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

// RetrieveEthTxsInRange retrieves eth transactions that match an inputted filter and are within a range given a page.
func (s Store) RetrieveEthTxsInRange(ctx context.Context, ethTxFilter db.EthTxFilter, startBlock, endBlock uint64, page int) ([]types.Transaction, error) {
	if page < 1 {
		page = 1
	}
	dbEthTxs := []EthTx{}
	query := ethTxFilterToQuery(ethTxFilter)
	rangeQuery := fmt.Sprintf("%s BETWEEN ? AND ?", BlockNumberFieldName)
	dbTx := s.DB().WithContext(ctx).
		Model(&EthTx{}).
		Where(&query).
		Where(rangeQuery, startBlock, endBlock).
		Order(fmt.Sprintf("%s, %s", BlockIndexFieldName, BlockIndexFieldName)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
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
