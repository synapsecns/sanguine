package base

import (
	"context"
	"fmt"
	"math/big"

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

// getChainID gets the chain id from non-legacy transaction types
// it is used to check chainids against the chainid passed in the raw id.
func getChainID(tx *types.Transaction) (hasType bool, chainID *big.Int) {
	switch tx.Type() {
	case types.LegacyTxType:
		return false, nil
	default:
		chainID = tx.ChainId()
		if chainID == nil || chainID.Cmp(big.NewInt(0)) == 0 {
			return false, nil
		}

		return true, tx.ChainId()
	}
}
