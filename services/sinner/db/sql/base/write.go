package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm/clause"
)

// StoreOriginSent stores an origin event.
func (s Store) StoreOriginSent(ctx context.Context, originSent *model.OriginSent) error {

	dbTx := s.DB().WithContext(ctx)
	if s.db.Dialector.Name() == dbcommon.Sqlite.String() {
		dbTx = dbTx.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: model.ContractAddressFieldName}, {Name: model.ChainIDFieldName}, {Name: model.TxHashFieldName}, {Name: model.BlockNumberFieldName}, {Name: model.MessageHashFieldName},
			},
			DoNothing: true,
		}).CreateInBatches(originSent, 10)
	} else {
		dbTx = dbTx.Clauses(clause.Insert{
			Modifier: "IGNORE",
		}).Create(originSent)
	}

	if dbTx.Error != nil {
		return fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return nil
}

// StoreExecuted stores an origin event.
func (s Store) StoreExecuted(ctx context.Context, executedEvent *model.Executed) error {

	dbTx := s.DB().WithContext(ctx)
	if s.db.Dialector.Name() == dbcommon.Sqlite.String() {
		dbTx = dbTx.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: model.ContractAddressFieldName}, {Name: model.ChainIDFieldName}, {Name: model.TxHashFieldName}, {Name: model.BlockNumberFieldName},
			},
			DoNothing: true,
		}).CreateInBatches(executedEvent, 10)
	} else {
		dbTx = dbTx.Clauses(clause.Insert{
			Modifier: "IGNORE",
		}).Create(executedEvent)
	}

	if dbTx.Error != nil {
		return fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return nil
}

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
		Create(&model.LastIndexedInfo{
			ContractAddress: address,
			ChainID:         chainID,
			BlockNumber:     blockNumber,
		})
	if dbTx.Error != nil {
		return fmt.Errorf("could not update last indexed info: %w", dbTx.Error)
	}
	return nil
}
