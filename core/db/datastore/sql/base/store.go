package base

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/db"
	"gorm.io/gorm"
	"math/big"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	db *gorm.DB
}

// NewStore creates a new tore.
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s Store) DB() *gorm.DB {
	return s.db
}

// GetAllModels gets all models to migrate
//see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels,
		&RawEVMTX{})
	return allModels
}

// StoreRawTx stores a raw transaction.
func (s Store) StoreRawTx(ctx context.Context, tx *types.Transaction, chainID *big.Int, from common.Address) error {
	toAddress := ""
	if tx != nil {
		toAddress = tx.To().String()
	}

	// sanity check for making sure transaction marshaled chainid matches derived chain id (if present)
	hasID, newId := getChainID(tx)
	if hasID {
		if newId.Cmp(chainID) != 0 {
			return fmt.Errorf("chainid mismatch, expected %d, got %d", chainID, newId)
		}
	}

	marshalledTx, err := tx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("could not marshall tx to binary: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).Create(&RawEVMTX{
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

func (s Store) StoreProcessedTx(tx *types.Transaction) {
	// TODO implement me
	panic("implement me")
}

func (s Store) GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error) {
	var newNonce sql.NullInt64

	selectMaxNonce := fmt.Sprintf("max(`%s`)", NonceFieldName)

	dbTx := s.DB().WithContext(ctx).Model(&RawEVMTX{}).Select(selectMaxNonce).Where(RawEVMTX{
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
		dbTx = s.DB().WithContext(ctx).Model(&RawEVMTX{}).Where(RawEVMTX{ChainID: chainID.Uint64(), From: fromAddress.String()}).Count(&count)
		if dbTx.Error != nil {
			return 0, fmt.Errorf("error getting count on %T: %w", &RawEVMTX{}, dbTx.Error)
		}

		if count == 0 {
			return 0, db.ErrNoNonceForChain
		}
	}

	return uint64(newNonce.Int64), nil
}
