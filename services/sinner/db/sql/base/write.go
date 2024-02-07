package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	"github.com/synapsecns/sanguine/services/sinner/types"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm/clause"
)

// StoreOriginSent stores an origin event.
func (s Store) StoreOriginSent(ctx context.Context, originSent *model.OriginSent) error {
	dbTx := s.DB().WithContext(ctx)
	dbTx = dbTx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: model.ChainIDFieldName}, {Name: model.TxHashFieldName}, {Name: model.MessageHashFieldName},
		},
		DoNothing: true,
	}).Create(originSent)
	if dbTx.Error != nil {
		return fmt.Errorf("could not store origin: %w", dbTx.Error)
	}

	return nil
}

// StoreExecuted stores an origin event.
func (s Store) StoreExecuted(ctx context.Context, executedEvent *model.Executed) error {
	dbTx := s.DB().WithContext(ctx)
	dbTx = dbTx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: model.ChainIDFieldName}, {Name: model.MessageHashFieldName},
		},
		DoNothing: true,
	}).Create(executedEvent)

	if dbTx.Error != nil {
		return fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return nil
}

// StoreLastIndexed stores the last indexed block number for a contract.
func (s Store) StoreLastIndexed(parentCtx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) (err error) {
	ctx, span := s.metrics.Tracer().Start(parentCtx, "StoreLastIndexed", trace.WithAttributes(
		attribute.String("contractAddress", contractAddress.String()),
		attribute.Int("chainID", int(chainID)),
		attribute.Int("blockNumber", int(blockNumber)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	address := contractAddress.String()

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: model.ContractAddressFieldName}, {Name: model.ChainIDFieldName}},
			DoUpdates: clause.AssignmentColumns([]string{model.BlockNumberFieldName}),
			Where: clause.Where{
				Exprs: []clause.Expression{
					clause.And(
						clause.Where{
							Exprs: []clause.Expression{
								clause.Eq{
									Column: clause.Column{Name: model.ContractAddressFieldName},
									Value:  address,
								},
								clause.Eq{
									Column: clause.Column{Name: model.ChainIDFieldName},
									Value:  chainID,
								},
							},
						},
						clause.Lt{
							Column: clause.Column{Name: model.BlockNumberFieldName},
							Value:  blockNumber,
						},
					),
				},
			},
		}).
		Create(&model.LastIndexed{
			ContractAddress: address,
			ChainID:         chainID,
			BlockNumber:     blockNumber,
		})
	if dbTx.Error != nil {
		return fmt.Errorf("could not update last indexed info: %w", dbTx.Error)
	}
	return nil
}

// StoreOrUpdateMessageStatus stores or updates the message status.
func (s Store) StoreOrUpdateMessageStatus(ctx context.Context, txHash string, messageHash string, messageType types.MessageType) error {
	dbTx := s.DB().WithContext(ctx)

	switch messageType {
	case types.Origin:
		// If the record exists, it will be updated, otherwise, a new one will be created
		dbTx = dbTx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: model.MessageHashFieldName}},
			DoUpdates: clause.AssignmentColumns([]string{model.OriginTxHashFieldName}),
		}).Create(&model.MessageStatus{
			MessageHash:       messageHash,
			OriginTxHash:      txHash,
			DestinationTxHash: "",
		})

	case types.Destination:
		// If the record exists, it will be updated, otherwise, a new one will be created with empty OriginTxHash
		dbTx = dbTx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: model.MessageHashFieldName}},
			DoUpdates: clause.AssignmentColumns([]string{model.DestinationTxHashFieldName}),
		}).Create(&model.MessageStatus{
			MessageHash:       messageHash,
			OriginTxHash:      "",
			DestinationTxHash: txHash,
		})
	default:
		return fmt.Errorf("unknown message type: %s", messageType)
	}

	if dbTx.Error != nil {
		return fmt.Errorf("could not store or update message status: %w", dbTx.Error)
	}

	return nil
}
