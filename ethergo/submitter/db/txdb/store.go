package txdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/big"
)

// NewTXStore creates a new transaction store.
func NewTXStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	db *gorm.DB
}

// GetNonceForChainID gets the nonce for the given chain id.
func (s Store) GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error) {
	var newNonce sql.NullInt64

	selectMaxNonce := fmt.Sprintf("max(`%s`)", nonceFieldName)

	dbTx := s.DB().WithContext(ctx).Model(&ETHTX{}).Select(selectMaxNonce).Where(ETHTX{
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
		dbTx = s.DB().WithContext(ctx).Model(&ETHTX{}).Where(ETHTX{ChainID: chainID.Uint64(), From: fromAddress.String()}).Count(&count)
		if dbTx.Error != nil {
			return 0, fmt.Errorf("error getting count on %T: %w", &ETHTX{}, dbTx.Error)
		}

		if count == 0 {
			return 0, db.ErrNoNonceForChain
		}
	}

	return uint64(newNonce.Int64), nil
}

// PutTX puts a transaction in the database.
func (s Store) PutTX(ctx context.Context, tx *types.Transaction, status db.Status) error {
	marshalledTX, err := tx.MarshalJSON()
	if err != nil {
		return fmt.Errorf("could not marshall tx to json: %w", err)
	}

	msg, err := util.TxToCall(tx)
	if err != nil {
		return fmt.Errorf("could not recover signer from tx: %w", err)
	}

	if tx.To() == nil {
		return errors.New("tx has no to address")
	}

	if !tx.ChainId().IsUint64() {
		return errors.New("chainid is not uint64")
	}

	dbTX := s.DB().WithContext(ctx).Create(&ETHTX{
		From:    msg.From.String(),
		ChainID: tx.ChainId().Uint64(),
		Nonce:   tx.Nonce(),
		RawTx:   marshalledTX,
		Status:  status,
	}).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: txHashFieldName}},
		DoUpdates: clause.AssignmentColumns([]string{
			statusFieldName,
		}),
	})

	if dbTX.Error != nil {
		return fmt.Errorf("could not store tx: %w", dbTX.Error)
	}
	return nil
}

// DB gets the database.
func (s Store) DB() *gorm.DB {
	return s.db
}

var _ db.Service = &Store{}
