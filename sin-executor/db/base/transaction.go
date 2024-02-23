package base

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/db"
	"gorm.io/gorm/clause"
)

type InterchainTransaction struct {
	// TransactionID is the transaction id.
	TransactionID string `gorm:"column:transaction_id;primaryKey"`
	// SrcChainID is the source chain id.
	SrcChainID uint64 `gorm:"column:src_chain_id;index"`
	// DstChainID is the destination chain id.
	DstChainID uint64 `gorm:"column:dst_chain_id;index"`
	// Status is the status of the transaction.
	Status db.ExecutableStatus `gorm:"column:status;index"`
	// Nonce is the nonce of the transaction.
	GasAirdrop string `gorm:"column:gas_airdrop;index"`
	// Options is the options of the transaction.
	GasLimit string `gorm:"column:gas_limit;index"`
	// EncodedTx is the encoded transaction.
	EncodedTx string `gorm:"column:encoded_tx"`
}

func (s InterchainTransaction) ToTransactionSent() (db.TransactionSent, error) {
	airdrop, ok := new(big.Int).SetString(s.GasAirdrop, 10)
	if !ok {
		return db.TransactionSent{}, fmt.Errorf("could not convert gas airdrop to big.Int")
	}

	gasLimit, ok := new(big.Int).SetString(s.GasLimit, 10)
	if !ok {
		return db.TransactionSent{}, fmt.Errorf("could not convert gas limit to big.Int")
	}

	return db.TransactionSent{
		Status:        s.Status,
		EncodedTX:     common.Hex2Bytes(s.EncodedTx),
		SrcChainID:    big.NewInt(int64(s.SrcChainID)),
		DstChainID:    big.NewInt(int64(s.DstChainID)),
		TransactionID: common.HexToHash(s.TransactionID),
		Options: interchainclient.OptionsV1{
			GasAirdrop: airdrop,
			GasLimit:   gasLimit,
		},
	}, nil
}

func fromInterchainTX(chainID *big.Int, interchainTx *interchainclient.InterchainClientV1InterchainTransactionSent, encodedTX []byte) InterchainTransaction {
	return InterchainTransaction{
		EncodedTx:     common.Bytes2Hex(encodedTX),
		TransactionID: common.Bytes2Hex(interchainTx.TransactionId[:]),
		GasAirdrop:    options.GasAirdrop.String(),
		GasLimit:      options.GasLimit.String(),
		SrcChainID:    chainID.Uint64(),
		DstChainID:    interchainTx.DstChainId.Uint64(),
		Status:        db.Seen,
	}
}

func (s Store) StoreInterchainTransaction(ctx context.Context, originChainID *big.Int, interchainTx *interchainclient.InterchainClientV1InterchainTransactionSent, encodedTX []byte) error {
	dbTx := s.db.WithContext(ctx).Model(&InterchainTransaction{}).Clauses(clause.OnConflict{DoNothing: true}).Create(fromInterchainTX(originChainID, interchainTx, encodedTX))
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
		marshaled, err := result.ToTransactionSent()
		if err != nil {
			return []db.TransactionSent{}, fmt.Errorf("could not marshal db result: %w", err)
		}
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
