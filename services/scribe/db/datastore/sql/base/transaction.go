package base

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm/clause"
)

// StoreEthTx stores a processed text.
func (s Store) StoreEthTx(ctx context.Context, tx *types.Transaction, chainID uint32) error {
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
			TxHash:    tx.Hash().String(),
			ChainID:   chainID,
			RawTx:     marshalledTx,
			GasFeeCap: tx.GasFeeCap().Uint64(),
			GasTipCap: tx.GasTipCap().Uint64(),
		})

	if dbTx.Error != nil {
		return fmt.Errorf("could not create raw tx: %w", dbTx.Error)
	}

	return nil
}
