package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/committee/db"
	"gorm.io/gorm/clause"
	"time"
)

// SignRequest is a request to sign a message.
type SignRequest struct {
	// TXHash is the hash of the transaction
	TXHash string `gorm:"column:tx_hash;index;size:256"`
	// TransactionID is the id of the transaction
	TransactionID string `gorm:"column:transaction_id;index;size:256"`
	// TXHash is the hash of the transaction
	Transaction string `gorm:"column:transaction"`
	// CreatedAt is the time the transaction was created
	CreatedAt time.Time
	// Nonce is the nonce of the raw evm tx
	Nonce uint64 `gorm:"column:nonce;index"`
	// Status is the status of the transaction
	Status db.SynapseRequestStatus `gorm:"column:status;index"`
	// Sender is the sender of the transaction
	Sender string `gorm:"column:sender;index;size:256"`
	// Receiver is the receiver of the transaction
	Receiver string `gorm:"column:receiver;index;size:256"`
	// OriginChainID is the chain id the transaction hash was sent on
	OriginChainID int `gorm:"column:origin_chain_id;index"`
	// DestinationChainID is the chain id the transaction hash will be sent on
	DestinationChainID int `gorm:"column:destination_chain_id;index"`
	// ModuleRequiredResponses is the number of responses required for the module
	ModuleRequiredResponses int `gorm:"column:module_required_responses"`
	// Signature is the signature of the transaction
	Signature []*Signature `gorm:"foreignKey:transaction_id;references:transaction_id"`
}

func (s *SignRequest) ToServiceSignRequest() db.SignRequest {
	signatureMap := make(map[common.Address][]byte)
	for _, signature := range s.Signature {
		signatureMap[common.HexToAddress(signature.SenderAddress)] = signature.Signature
	}
	return db.SignRequest{
		TXHash:                  common.HexToHash(s.TXHash),
		TransactionID:           common.HexToHash(s.TransactionID),
		Transaction:             common.FromHex(s.Transaction),
		Nonce:                   s.Nonce,
		Status:                  s.Status,
		Sender:                  common.HexToAddress(s.Sender),
		Receiver:                common.HexToAddress(s.Receiver),
		OriginChainID:           s.OriginChainID,
		DestinationChainID:      s.DestinationChainID,
		ModuleRequiredResponses: s.ModuleRequiredResponses,
		Signature:               signatureMap,
	}
}

func (s Store) toSignRequest(ctx context.Context, sr synapsemodule.SynapseModuleModuleMessageSent) (SignRequest, error) {
	decodedTx, err := s.rawTxDecoder(ctx, sr.Transaction)
	if err != nil {
		return SignRequest{}, fmt.Errorf("could not decode transaction: %w", err)
	}

	return SignRequest{
		TXHash:                  sr.Raw.TxHash.String(),
		TransactionID:           common.Bytes2Hex(decodedTx.TransactionId[:]),
		Transaction:             common.Bytes2Hex(sr.Transaction),
		OriginChainID:           int(decodedTx.SrcChainId.Int64()),
		DestinationChainID:      int(decodedTx.DstChainId.Int64()),
		Sender:                  decodedTx.SrcSender.String(),
		Receiver:                common.Bytes2Hex(decodedTx.DstReceiver[:]),
		Nonce:                   decodedTx.Nonce,
		Status:                  db.Seen,
		ModuleRequiredResponses: int(decodedTx.RequiredModuleResponses.Int64()),
	}, nil
}

// Signature is a signature of a transaction.
type Signature struct {
	// Signature is the actual signature
	Signature []byte `gorm:"column:signature"`
	// SenderAddress is the address of the sender
	SenderAddress string `gorm:"column:sender_address;index;size:256;primaryKey"`
	// Transaction is the actual transaction
	TransactionID string `gorm:"column:transaction_id;primaryKey;size:256"`
	// SignRequest is a reference to the SignRequest
	SignRequest *SignRequest `gorm:"foreignKey:transaction_id;references:transaction_id"`
}

func (s Store) StoreInterchainTransactionReceived(ctx context.Context, sr synapsemodule.SynapseModuleModuleMessageSent) error {
	signRequest, err := s.toSignRequest(ctx, sr)
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
