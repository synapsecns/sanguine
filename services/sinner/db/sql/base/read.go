package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	graphqlModel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"
	"gorm.io/gorm"
)

// RetrieveMessageStatus retrieve message status.
func (s Store) RetrieveMessageStatus(ctx context.Context, messageHash string) (graphqlModel.MessageStatus, error) {
	var record model.MessageStatus

	err := s.DB().WithContext(ctx).
		Model(&model.MessageStatus{}).
		Where(&model.MessageStatus{MessageHash: messageHash}).
		First(&record).Error

	if err != nil {
		return graphqlModel.MessageStatus{}, fmt.Errorf("could not retrieve message status: %w", err)
	}

	payload := graphqlModel.MessageStatus{
		MessageHash:       &record.MessageHash,
		OriginTxHash:      &record.OriginTxHash,
		DestinationTxHash: &record.DestinationTxHash,
	}

	// Determine the last seen state
	if record.DestinationTxHash == "" {
		ls := graphqlModel.MessageStateLastSeenOrigin
		payload.LastSeen = &ls
	} else {
		ls := graphqlModel.MessageStateLastSeenDestination
		payload.LastSeen = &ls
	}

	return payload, nil
}

// RetrieveMessagesByStatus retrieves pending messages.
func (s Store) RetrieveMessagesByStatus(ctx context.Context, messageStatus graphqlModel.MessageState, page int) ([]*graphqlModel.MessageStatus, error) {
	var records []model.MessageStatus

	var queryStr string
	switch messageStatus {
	case graphqlModel.MessageStatePending:
		queryStr = fmt.Sprintf("%s = '' OR %s IS NULL", model.DestinationTxHashFieldName, model.DestinationTxHashFieldName)
	case graphqlModel.MessageStateCompleted:
		queryStr = fmt.Sprintf("%s != '' AND %s IS NOT NULL", model.DestinationTxHashFieldName, model.DestinationTxHashFieldName)
	}

	err := s.DB().WithContext(ctx).
		Model(&model.MessageStatus{}).
		Where(queryStr).
		Offset((page - 1) * model.PageSize).
		Limit(model.PageSize).
		Find(&records).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []*graphqlModel.MessageStatus{}, nil
	}

	if err != nil {
		return []*graphqlModel.MessageStatus{}, fmt.Errorf("could not retrieve message status: %w", err)
	}

	payload := make([]*graphqlModel.MessageStatus, 0, len(records))
	for _, record := range records {
		ms := &graphqlModel.MessageStatus{
			MessageHash:       &record.MessageHash,
			OriginTxHash:      &record.OriginTxHash,
			DestinationTxHash: &record.DestinationTxHash,
		}

		if record.OriginTxHash == "" {
			ls := graphqlModel.MessageStateLastSeenUnknown
			ms.LastSeen = &ls
		} else {
			ls := graphqlModel.MessageStateLastSeenOrigin
			ms.LastSeen = &ls
		}
		payload = append(payload, ms)
	}

	return payload, nil
}

// RetrieveLastStoredBlock gets the last stored block.
func (s Store) RetrieveLastStoredBlock(ctx context.Context, chainID uint32, address common.Address) (uint64, error) {
	var lastIndexed model.LastIndexed
	err := s.DB().WithContext(ctx).
		Model(&model.LastIndexed{}).
		Where(&model.LastIndexed{
			ContractAddress: address.String(),
			ChainID:         chainID,
		}).
		Order("block_number DESC").First(&lastIndexed).Error

	if err != nil && err.Error() == "record not found" {
		return 0, nil
	}

	if err != nil {
		return 0, fmt.Errorf("could not retrieve last stored block: %w", err)
	}

	return lastIndexed.BlockNumber, nil
}

// RetrieveOriginSent gets the Origin Sent event.
func (s Store) RetrieveOriginSent(ctx context.Context, filter model.OriginSent) ([]model.OriginSent, error) {
	var originSentRecord []model.OriginSent
	err := s.DB().WithContext(ctx).
		Model(&model.OriginSent{}).
		Where(filter).
		Scan(&originSentRecord).Error

	if err != nil {
		return []model.OriginSent{}, fmt.Errorf("could not retrieve Origin Sent event: %w", err)
	}
	return originSentRecord, nil
}

// RetrieveExecuted gets a Executed event.
func (s Store) RetrieveExecuted(ctx context.Context, filter model.Executed) ([]model.Executed, error) {
	var executedRecord []model.Executed
	err := s.DB().WithContext(ctx).
		Model(&model.Executed{}).
		Where(filter).
		Scan(&executedRecord).Error

	if err != nil {
		return []model.Executed{}, fmt.Errorf("could not retrieve Executed event: %w", err)
	}
	return executedRecord, nil
}
