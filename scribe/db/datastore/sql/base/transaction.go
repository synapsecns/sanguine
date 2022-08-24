package base

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// StoreProcessedTx stores a processed text.
func (s Store) StoreProcessedTx(ctx context.Context, tx *types.Transaction) error {
	marshalledTx, err := tx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("could not marshall tx to binary: %w", err)
	}

	signer := types.LatestSignerForChainID(tx.ChainId())
	sender, err := types.Sender(signer, tx)
	if err != nil {
		return fmt.Errorf("could not get sender for tx: %s: %w", tx.Hash(), err)
	}

	parentID, err := s.getRawTXIDByParams(ctx, tx.Nonce(), tx.ChainId(), sender)
	if err != nil {
		return fmt.Errorf("could not get parent tx: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).Create(&ProcessedEthTx{
		TxHash:    tx.Hash().String(),
		RawTx:     marshalledTx,
		RawEthTx:  parentID,
		GasFeeCap: tx.GasFeeCap().Uint64(),
		GasTipCap: tx.GasTipCap().Uint64(),
	})

	if dbTx.Error != nil {
		return fmt.Errorf("could not create raw tx: %w", dbTx.Error)
	}

	return nil
}

// StoreRawTx stores a raw transaction.
func (s Store) StoreRawTx(ctx context.Context, tx *types.Transaction, chainID *big.Int, from common.Address) error {
	toAddress := ""
	if tx != nil {
		toAddress = tx.To().String()
	}

	// sanity check for making sure transaction marshaled chainid matches derived chain id (if present)
	hasID, newID := getChainID(tx)
	if hasID {
		if newID.Cmp(chainID) != 0 {
			return fmt.Errorf("chainid mismatch, expected %d, got %d", chainID, newID)
		}
	}

	marshalledTx, err := tx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("could not marshall tx to binary: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).Create(&RawEthTX{
		From:    from.String(),
		To:      toAddress,
		ChainID: chainID.Uint64(),
		Nonce:   tx.Nonce(),
		RawTx:   marshalledTx,
	})

	if dbTx.Error != nil {
		return fmt.Errorf("could not create raw tx: %w", dbTx.Error)
	}

	return nil
}

// getRawTXIDByParams by nonce/chain id gets the raw transaction id by a combination of the nonce and chain id
// this is used for storing processed txes.
func (s Store) getRawTXIDByParams(ctx context.Context, nonce uint64, chainID *big.Int, sender common.Address) (id uint, err error) {
	var res RawEthTX
	dbTx := s.DB().Select("ID").WithContext(ctx).Model(&RawEthTX{}).Where(RawEthTX{
		ChainID: chainID.Uint64(),
		Nonce:   nonce,
		From:    sender.String(),
	}).Find(&res)

	if dbTx.RowsAffected == 0 {
		return 0, dbcommon.ErrNotFound
	}

	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not get %T by chainID: %d and nonce: %d. error: %w", &RawEthTX{}, chainID.Uint64(), nonce, dbTx.Error)
	}

	return res.ID, nil
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

// GetNonceForChainID gets a nonce for a chainid.
func (s Store) GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error) {
	var newNonce sql.NullInt64

	selectMaxNonce := "max(`nonce`)"

	dbTx := s.DB().WithContext(ctx).Model(&RawEthTX{}).Select(selectMaxNonce).Where(RawEthTX{
		From:    fromAddress.String(),
		ChainID: chainID.Uint64(),
	}).Scan(&newNonce)

	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not get nonce for chain id: %w", dbTx.Error)
	}

	// if no nonces, return the corresponding error.
	if newNonce.Int64 == 0 {
		// we need to check if any nonces exist first
		var count int64
		dbTx = s.DB().WithContext(ctx).Model(&RawEthTX{}).Where(RawEthTX{ChainID: chainID.Uint64(), From: fromAddress.String()}).Count(&count)
		if dbTx.Error != nil {
			return 0, fmt.Errorf("error getting count on %T: %w", &RawEthTX{}, dbTx.Error)
		}

		if count == 0 {
			return 0, dbcommon.ErrNoNonceForChain
		}
	}

	return uint64(newNonce.Int64), nil
}
