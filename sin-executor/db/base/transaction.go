package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/db"
	"math/big"
)

type InterchainTransaction struct {
	// TransactionID is the transaction id.
	TransactionID string `gorm:"primaryKey"`
	// SrcSender is the sender of the transaction.
	SrcSender string `gorm:"column:src_sender;index"`
	// DstReceiver is the receiver of the transaction.
	DstReceiver string `gorm:"column:dst_receiver;index"`
	// SrcChainID is the source chain id.
	SrcChainID uint64 `gorm:"column:src_chain_id;index"`
	// DstChainID is the destination chain id.
	DstChainID uint64 `gorm:"column:dst_chain_id;index"`
	// Message is the message of the transaction.
	Message string `gorm:"column:message"`
	// Status is the status of the transaction.
	Status db.ExecutableStatus `gorm:"column:status;index"`
	// Nonce is the nonce of the transaction.
	Nonce uint64 `gorm:"column:nonce;index"`
	// DBWriterNonce is the nonce of the transaction in the database.
	DBWriterNonce uint64 `gorm:"column:db_writer_nonce;index"`
}

func (s InterchainTransaction) ToTransactionSent() db.TransactionSent {
	return db.TransactionSent{
		Status: s.Status,
		InterchainClientV1InterchainTransactionSent: interchainclient.InterchainClientV1InterchainTransactionSent{
			TransactionId: common.HexToHash(s.TransactionID),
			SrcSender:     common.HexToHash(s.SrcSender),
			DstReceiver:   common.HexToHash(s.DstReceiver[:]),
			SrcChainId:    big.NewInt(int64(s.SrcChainID)),
			Message:       common.Hex2Bytes(s.Message),
			Nonce:         s.Nonce,
			DstChainId:    big.NewInt(int64(s.DstChainID)),
			DbWriterNonce: big.NewInt(int64(s.DBWriterNonce)),
		},
	}
}

func fromInterchainTX(interchainTx interchainclient.InterchainClientV1InterchainTransactionSent) InterchainTransaction {
	return InterchainTransaction{
		TransactionID: common.Bytes2Hex(interchainTx.TransactionId[:]),
		SrcSender:     common.Bytes2Hex(interchainTx.SrcSender[:]),
		DstReceiver:   common.Bytes2Hex(interchainTx.DstReceiver[:]),
		SrcChainID:    interchainTx.SrcChainId.Uint64(),
		DstChainID:    interchainTx.DstChainId.Uint64(),
		Message:       common.Bytes2Hex(interchainTx.Message[:]),
		Status:        db.Seen,
		Nonce:         interchainTx.Nonce,
		DBWriterNonce: interchainTx.DbWriterNonce.Uint64(),
	}
}

func (s Store) StoreInterchainTransaction(ctx context.Context, interchainTx interchainclient.InterchainClientV1InterchainTransactionSent) error {
	dbTx := s.db.WithContext(ctx).Model(&InterchainTransaction{}).Create(fromInterchainTX(interchainTx))
	if dbTx.Error != nil {
		return fmt.Errorf("could not store interchain transaction: %w", dbTx.Error)
	}

	return nil
}

func (s Store) GetInterchainTXsByStatus(ctx context.Context, matchStatuses ...db.ExecutableStatus) (res []db.TransactionSent, err error) {
	var interchainTransactions []InterchainTransaction

	inArgs := make([]int, len(matchStatuses))
	for i := range matchStatuses {
		inArgs[i] = int(matchStatuses[i].Int())
	}

	// TODO: consider pagination
	tx := s.DB().WithContext(ctx).Model(&InterchainTransaction{}).Where(fmt.Sprintf("%s IN ?", statusFieldName), inArgs).Find(&interchainTransactions)
	if tx.Error != nil {
		return []db.TransactionSent{}, fmt.Errorf("could not get db results: %w", tx.Error)
	}

	for _, result := range interchainTransactions {
		marshaled := result.ToTransactionSent()
		res = append(res, marshaled)
	}
	return res, nil
}

func (s Store) UpdateInterchainTransactionStatus(ctx context.Context, transactionid [32]byte, status db.ExecutableStatus) error {
	tx := s.DB().WithContext(ctx).Model(&InterchainTransaction{}).Where(fmt.Sprintf("%s = ?", transactionIDFieldName), common.Bytes2Hex(transactionid[:])).Update(statusFieldName, status)
	if tx.Error != nil {
		return fmt.Errorf("could not update sign request status: %w", tx.Error)
	}
	return nil
}
