package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/committee/db"
	"gorm.io/gorm/clause"
	"math/big"
	"time"
)

// SignRequest is a request to sign a message.
type SignRequest struct {
	// TXHash is the hash of the transaction
	TXHash string `gorm:"column:tx_hash;index;size:256"`
	// TransactionID is the id of the transaction
	TransactionID string `gorm:"column:transaction_id;index;size:256;primaryKey"`
	// EntryHash is the hash of the entry
	EntryHash string `gorm:"column:entry_hash;index;size:256"`
	// DataHash is the hash of the data
	DataHash string `gorm:"column:data_hash;index;size:256"`
	// CreatedAt is the time the transaction was created
	CreatedAt time.Time
	// Nonce is the nonce of the raw evm tx
	Nonce uint64 `gorm:"column:nonce;index"`
	// Status is the status of the transaction
	Status db.SynapseRequestStatus `gorm:"column:status;index"`
	// Sender is the sender of the transaction
	Sender string `gorm:"column:sender;index;size:256"`
	// OriginChainID is the chain id the transaction hash was sent on
	OriginChainID int `gorm:"column:origin_chain_id;index"`
	// DestinationChainID is the chain id the transaction hash will be sent on
	DestinationChainID int `gorm:"column:destination_chain_id;index"`
}

// ToServiceSignRequest converts a SignRequest to a db.SignRequest.
func (s *SignRequest) ToServiceSignRequest() db.SignRequest {
	return db.SignRequest{
		InterchainEntry: synapsemodule.InterchainEntry{
			SrcChainId:  big.NewInt(int64(s.OriginChainID)),
			SrcWriter:   common.HexToHash(s.Sender),
			WriterNonce: big.NewInt(int64(s.Nonce)),
			DataHash:    common.HexToHash(s.DataHash),
		},
		DestChainID:     big.NewInt(int64(s.DestinationChainID)),
		Status:          s.Status,
		SignedEntryHash: common.HexToHash(s.EntryHash),
	}
}

func toSignRequest(sr synapsemodule.SynapseModuleVerificationRequested) (SignRequest, error) {
	return SignRequest{
		TXHash:             sr.Raw.TxHash.String(),
		TransactionID:      common.Bytes2Hex(sr.SignableEntryHash[:]),
		OriginChainID:      int(sr.Entry.SrcChainId.Int64()),
		DestinationChainID: int(sr.DestChainId.Int64()),
		Sender:             common.Bytes2Hex(sr.Entry.SrcWriter[:]),
		Nonce:              sr.Entry.WriterNonce.Uint64(),
		DataHash:           common.Bytes2Hex(sr.Entry.DataHash[:]),
		EntryHash:          common.Bytes2Hex(sr.SignableEntryHash[:]),
		Status:             db.Seen,
	}, nil
}

// UpdateSignRequestStatus updates a sign request status.
func (s Store) UpdateSignRequestStatus(ctx context.Context, txid common.Hash, status db.SynapseRequestStatus) error {
	tx := s.DB().WithContext(ctx).Model(&SignRequest{}).Where(fmt.Sprintf("%s = ?", transactionIDFieldName), common.Bytes2Hex(txid[:])).Update(statusFieldName, status)
	if tx.Error != nil {
		return fmt.Errorf("could not update sign request status: %w", tx.Error)
	}
	return nil
}

// StoreInterchainTransactionReceived stores an interchain transaction received.
func (s Store) StoreInterchainTransactionReceived(ctx context.Context, sr synapsemodule.SynapseModuleVerificationRequested) error {
	signRequest, err := toSignRequest(sr)
	if err != nil {
		return fmt.Errorf("could not convert to sign request: %w", err)
	}

	tx := s.DB().WithContext(ctx).
		Model(&SignRequest{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: transactionIDFieldName}},
			DoNothing: true,
		}).
		Create(&signRequest)

	if tx.Error != nil {
		return fmt.Errorf("could not create sign request: %w", tx.Error)
	}

	return nil
}

// GetQuoteResultsByStatus gets quote results by status.
func (s Store) GetQuoteResultsByStatus(ctx context.Context, matchStatuses ...db.SynapseRequestStatus) (res []db.SignRequest, err error) {
	var signRequests []SignRequest

	inArgs := make([]int, len(matchStatuses))
	for i := range matchStatuses {
		inArgs[i] = int(matchStatuses[i].Int())
	}

	// TODO: consider pagination
	tx := s.DB().WithContext(ctx).Model(&SignRequest{}).Where(fmt.Sprintf("%s IN ?", statusFieldName), inArgs).Find(&signRequests)
	if tx.Error != nil {
		return []db.SignRequest{}, fmt.Errorf("could not get db results: %w", tx.Error)
	}

	for _, result := range signRequests {
		marshaled := result.ToServiceSignRequest()
		if err != nil {
			return []db.SignRequest{}, fmt.Errorf("could not get quotes")
		}
		res = append(res, marshaled)
	}
	return res, nil
}
