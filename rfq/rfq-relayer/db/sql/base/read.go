package base

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"
)

// GetOriginBridgeEvent gets a origin bridge event.
func (s Store) GetOriginBridgeEvent(ctx context.Context, transactionID string) (*model.OriginBridgeEvent, error) {
	var event model.OriginBridgeEvent
	err := s.DB().WithContext(ctx).
		Model(&model.OriginBridgeEvent{}).Where(&model.OriginBridgeEvent{TransactionID: transactionID}).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// GetDestinationBridgeEvent gets a destination bridge event.
func (s Store) GetDestinationBridgeEvent(ctx context.Context, transactionID string) (*model.DestinationBridgeEvent, error) {
	var event model.DestinationBridgeEvent
	err := s.DB().WithContext(ctx).
		Model(&model.DestinationBridgeEvent{}).Where(&model.DestinationBridgeEvent{TransactionID: transactionID}).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// GetLastIndexed gets the last indexed block by contract.
func (s Store) GetLastIndexed(ctx context.Context, chainID uint32, address common.Address) (uint64, error) {
	var lastIndexed model.LastIndexed
	err := s.DB().WithContext(ctx).
		Model(&model.LastIndexed{}).
		Where(&model.LastIndexed{
			ContractAddress: address.String(),
			ChainID:         chainID,
		}).First(&lastIndexed).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}

	if err != nil {
		return 0, fmt.Errorf("could not retrieve last stored block: %w", err)
	}
	return lastIndexed.BlockNumber, nil
}

// GetToken gets a token's metadata with a token's ID.
func (s Store) GetToken(ctx context.Context, tokenID string) (*model.Token, error) {
	var token model.Token
	err := s.DB().WithContext(ctx).
		Model(&model.Token{}).Where(&model.Token{TokenID: tokenID}).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// GetDeadlineQueueEvents gets all events in the deadline queue table.
func (s Store) GetDeadlineQueueEvents(ctx context.Context) ([]*model.DeadlineQueue, error) {
	var queueEvents []*model.DeadlineQueue
	err := s.DB().WithContext(ctx).
		Model(&model.DeadlineQueue{}).
		Scan(&queueEvents).Error
	if err != nil && err.Error() == "record not found" {
		logger.Infof("No events in deadline queue db store. Proceeding...")
		return []*model.DeadlineQueue{}, nil
	}
	if err != nil {
		return nil, err
	}
	return queueEvents, nil
}
