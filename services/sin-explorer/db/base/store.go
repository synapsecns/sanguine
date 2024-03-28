package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	listenerDB "github.com/synapsecns/sanguine/ethergo/listener/db"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"gorm.io/gorm"
)

// Store implements the service.
type Store struct {
	listenerDB.ChainListenerDB
	db *gorm.DB
}

// NewStore creates a new store.
func NewStore(db *gorm.DB, metrics metrics.Handler) *Store {
	return &Store{
		ChainListenerDB: listenerDB.NewChainListenerStore(db, metrics),
		db:              db,
	}
}

// DB gets the database object for mutation outside of the lib.
func (s Store) DB() *gorm.DB {
	return s.db
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(listenerDB.GetAllModels())
	return allModels
}

// PutInterchainTransactionSent puts an interchain transaction sent into the database.
func (s Store) PutInterchainTransactionSent(ctx context.Context, sent *interchainclient.InterchainClientV1InterchainTransactionSent) error {
	// TODO: consider addnig an onconflict do nothing clause
	tx := s.DB().WithContext(ctx).Create(InterchainTransactionSent{
		TransactionId:   common.Bytes2Hex(sent.TransactionId[:]),
		DbNonce:         sent.DbNonce.Uint64(),
		EntryIndex:      sent.EntryIndex,
		DstChainId:      sent.DstChainId.Uint64(),
		SrcSender:       common.Bytes2Hex(sent.SrcSender[:]),
		DstReceiver:     common.Bytes2Hex(sent.DstReceiver[:]),
		VerificationFee: sent.VerificationFee.String(),
		ExecutionFee:    sent.ExecutionFee.String(),
		Options:         common.Bytes2Hex(sent.Options),
		Message:         common.Bytes2Hex(sent.Message),
		TransactionHash: sent.Raw.TxHash.String(),
	})

	if tx.Error != nil {
		return fmt.Errorf("could not put interchain transaction sent: %w", tx.Error)
	}

	return nil
}

// PutInterchainTransactionReceived puts an interchain transaction received into the database.
func (s Store) PutInterchainTransactionReceived(ctx context.Context, received *interchainclient.InterchainClientV1InterchainTransactionReceived) error {
	tx := s.DB().WithContext(ctx).Create(InterchainTransactionReceived{
		TransactionId:   common.Bytes2Hex(received.TransactionId[:]),
		DbNonce:         received.DbNonce.Uint64(),
		EntryIndex:      received.EntryIndex,
		SrcChainId:      received.SrcChainId.Uint64(),
		SrcSender:       common.Bytes2Hex(received.SrcSender[:]),
		DstReceiver:     common.Bytes2Hex(received.DstReceiver[:]),
		TransactionHash: received.Raw.TxHash.String(),
	})

	if tx.Error != nil {
		return fmt.Errorf("could not put interchain transaction received: %w", tx.Error)
	}

	return nil
}
