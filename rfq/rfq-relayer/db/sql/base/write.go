package base

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"
	requests "github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"
	"gorm.io/gorm/clause"
)

// StoreOriginBridgeEvent adds a new origin event to the database.
func (s Store) StoreOriginBridgeEvent(ctx context.Context, chainID uint32, log *types.Log, event *bindings.FastBridgeBridgeRequested) error {
	transactionID := common.Bytes2Hex(event.TransactionId[:]) // keccak256 hash of the request
	bridgeTransaction, err := requests.Decode(event.Request)
	if err != nil {
		return fmt.Errorf("could not decode bridge transaction: %w", err)
	}

	requestStr := common.Bytes2Hex(event.Request)
	// Check if the chain ID in the event matches the chain ID in the bridge transaction
	if chainID != bridgeTransaction.OriginChainId {
		return fmt.Errorf("chain ID '%d' not same as bridge transaction origin chain ID '%d'", chainID, bridgeTransaction.OriginChainId)
	}
	e := model.OriginBridgeEvent{
		TransactionID: transactionID,
		Request:       requestStr,
		OriginChainID: bridgeTransaction.OriginChainId,
		DestChainID:   bridgeTransaction.DestChainId,
		OriginSender:  bridgeTransaction.OriginSender.Hex(),
		DestRecipient: bridgeTransaction.DestRecipient.Hex(),
		OriginToken:   bridgeTransaction.OriginToken.Hex(),
		DestToken:     bridgeTransaction.DestToken.Hex(),
		OriginAmount:  bridgeTransaction.OriginAmount.String(),
		DestAmount:    bridgeTransaction.DestAmount.String(),
		Deadline:      bridgeTransaction.Deadline.String(),
		Nonce:         bridgeTransaction.Nonce.String(),
		BlockNumber:   log.BlockNumber,
		TxHash:        log.TxHash.Hex(),
		TxIndex:       log.TxIndex,
		BlockHash:     log.BlockHash.Hex(),
		LogIndex:      log.Index,
		Removed:       log.Removed,
	}
	// Handle any re-insertion attempts
	return s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: model.TransactionIDFieldName},
		},
		DoNothing: true,
	}).Create(&e).Error
}

// StoreDestinationBridgeEvent adds a new destination to the database.
func (s Store) StoreDestinationBridgeEvent(ctx context.Context, log *types.Log, originEvent *model.OriginBridgeEvent) error {
	e := model.DestinationBridgeEvent{
		TransactionID: originEvent.TransactionID,
		Request:       originEvent.Request,
		OriginChainID: originEvent.OriginChainID,
		DestChainID:   originEvent.DestChainID,
		BlockNumber:   log.BlockNumber,
		TxHash:        log.TxHash.Hex(),
		TxIndex:       log.TxIndex,
		BlockHash:     log.BlockHash.Hex(),
		LogIndex:      log.Index,
		Removed:       log.Removed,
	}
	// Handle any re-insertion attempts
	return s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: model.TransactionIDFieldName},
		},
		DoNothing: true,
	}).Create(&e).Error
}

// StoreLastIndexed stores the last indexed block number for a contract. It will only insert if the current stored block number is less than the block number being inserted.
func (s Store) StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error {
	var lastIndexed model.LastIndexed
	address := contractAddress.String()
	// Gorm clauses don't work on all databases
	dbTx := s.db.WithContext(ctx).
		Model(&model.LastIndexed{}).
		Where(&model.LastIndexed{
			ChainID:         chainID,
			ContractAddress: address,
		}).
		Order("block_number DESC").
		Limit(1).
		Scan(&lastIndexed)
	if dbTx.Error == nil && blockNumber > lastIndexed.BlockNumber {
		// Populate the struct
		lastIndexed.ChainID = chainID
		lastIndexed.BlockNumber = blockNumber
		lastIndexed.ContractAddress = address

		// Upsert, clauses don[t work here with gorm.model
		if s.db.Model(&model.LastIndexed{}).Where(&model.LastIndexed{
			ChainID:         chainID,
			ContractAddress: address,
		}).Updates(&lastIndexed).RowsAffected == 0 {
			s.db.Create(&lastIndexed)
		}
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last block: %w", dbTx.Error)
		}
	}
	return nil
}

// StoreToken is an UPSERT operation that updates the token metadata if it exists, or inserts it if it doesn't.
func (s Store) StoreToken(ctx context.Context, token *model.Token) error {
	return s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: model.TokenIDFieldName},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			model.TokenIDFieldName, model.ChainIDFieldName, model.AddressFieldName,
			model.SymbolFieldName, model.NameFieldName, model.DecimalsFieldName,
		}),
	}).Create(&token).Error
}

// StoreDeadlineQueueEvent inserts a deadline queue entry if it does not exist in the table.
func (s Store) StoreDeadlineQueueEvent(ctx context.Context, entry *model.DeadlineQueue) error {
	// Handle any re-insertion attempts
	dbTx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: model.TransactionIDFieldName},
		},
		DoNothing: true,
	}).Create(&entry)
	return dbTx.Error
}

// RemoveDeadlineQueueEvent removes an event from the deadline queue table.
func (s Store) RemoveDeadlineQueueEvent(ctx context.Context, transactionID string) error {
	return s.DB().WithContext(ctx).Delete(&model.DeadlineQueue{}, &model.DeadlineQueue{TransactionID: transactionID}).Error
}
