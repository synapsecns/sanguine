package base

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/committee/db"
	"gorm.io/gorm/clause"
)

// VerificationRequest is a request to sign a message.
type VerificationRequest struct {
	// TXHash is the hash of the transaction
	TXHash string `gorm:"column:tx_hash;index;size:256"`
	// TransactionID is the id of the transaction
	TransactionID string `gorm:"column:transaction_id;index;size:256;primaryKey"`
	// DataHash is the hash of the data
	Entry string `gorm:"column:entry"`
	// CreatedAt is the time the transaction was created
	CreatedAt time.Time
	// Status is the status of the transaction
	Status db.SynapseRequestStatus `gorm:"column:status;index"`
	// OriginChainID is the chain id the transaction hash was sent on
	OriginChainID int `gorm:"column:origin_chain_id;index"`
	// DestinationChainID is the chain id the transaction hash will be sent on
	DestinationChainID int `gorm:"column:destination_chain_id;index"`
}

// ToServiceSignRequest converts a VerificationRequest to a db.SignRequest.
func (s *VerificationRequest) ToServiceSignRequest() db.SignRequest {
	return db.SignRequest{
		OriginChainID:   big.NewInt(int64(s.OriginChainID)),
		DestChainID:     big.NewInt(int64(s.DestinationChainID)),
		Status:          s.Status,
		TXHash:          common.HexToHash(s.TXHash),
		Entry:           common.Hex2Bytes(s.Entry),
		SignedEntryHash: common.HexToHash(s.TransactionID),
	}
}

func toSignRequest(originChainID int, sr synapsemodule.SynapseModuleVerificationRequested) VerificationRequest {
	return VerificationRequest{
		OriginChainID:      originChainID,
		TXHash:             sr.Raw.TxHash.String(),
		TransactionID:      common.Bytes2Hex(sr.EthSignedEntryHash[:]),
		Entry:              common.Bytes2Hex(sr.Entry),
		DestinationChainID: int(sr.DestChainId.Int64()),
		Status:             db.Seen,
	}
}

// UpdateSignRequestStatus updates a sign request status.
func (s Store) UpdateSignRequestStatus(ctx context.Context, txid common.Hash, status db.SynapseRequestStatus) error {
	tx := s.DB().WithContext(ctx).Model(&VerificationRequest{}).Where(fmt.Sprintf("%s = ?", transactionIDFieldName), common.Bytes2Hex(txid[:])).Update(statusFieldName, status)
	if tx.Error != nil {
		return fmt.Errorf("could not update sign request status: %w", tx.Error)
	}
	return nil
}

// StoreInterchainTransactionReceived stores an interchain transaction received.
func (s Store) StoreInterchainTransactionReceived(ctx context.Context, originChainID int, sr synapsemodule.SynapseModuleVerificationRequested) error {
	signRequest := toSignRequest(originChainID, sr)

	tx := s.DB().WithContext(ctx).
		Model(&VerificationRequest{}).
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
	var signRequests []VerificationRequest

	inArgs := make([]int, len(matchStatuses))
	for i := range matchStatuses {
		inArgs[i] = int(matchStatuses[i].Int())
	}

	// TODO: consider pagination
	tx := s.DB().WithContext(ctx).Model(&VerificationRequest{}).Where(fmt.Sprintf("%s IN ?", statusFieldName), inArgs).Find(&signRequests)
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
