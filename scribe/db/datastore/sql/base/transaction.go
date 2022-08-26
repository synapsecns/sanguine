package base

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

// StoreEthTx stores a processed text.
func (s Store) StoreEthTx(ctx context.Context, tx *types.Transaction) error {
	marshalledTx, err := tx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("could not marshall tx to binary: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).Create(&EthTx{
		TxHash:    tx.Hash().String(),
		RawTx:     marshalledTx,
		GasFeeCap: tx.GasFeeCap().Uint64(),
		GasTipCap: tx.GasTipCap().Uint64(),
	})

	if dbTx.Error != nil {
		return fmt.Errorf("could not create raw tx: %w", dbTx.Error)
	}

	return nil
}
