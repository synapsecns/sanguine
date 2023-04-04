package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm/clause"
)

// StoreLastIndexed stores the last indexed block number for a contract.
// It updates the value if there is a previous last indexed value, and creates a new
// entry if there is no previous value.
func (s Store) StoreLastIndexed(parentCtx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) (err error) {
	ctx, span := s.metrics.Tracer().Start(parentCtx, "StoreLastIndexed", trace.WithAttributes(
		attribute.String("contractAddress", contractAddress.String()),
		attribute.Int("chainID", int(chainID)),
		attribute.Int("blockNumber", int(blockNumber)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: ContractAddressFieldName}, {Name: ChainIDFieldName}},
			DoUpdates: clause.AssignmentColumns([]string{BlockNumberFieldName}),
			Where: clause.Where{
				Exprs: []clause.Expression{
					clause.And(
						clause.Where{
							Exprs: []clause.Expression{
								clause.Eq{
									Column: clause.Column{Name: ContractAddressFieldName},
									Value:  contractAddress.String(),
								},
								clause.Eq{
									Column: clause.Column{Name: ChainIDFieldName},
									Value:  chainID,
								},
							},
						},
						clause.Lt{
							Column: clause.Column{Name: BlockNumberFieldName},
							Value:  blockNumber,
						},
					),
				},
			},
		}).
		Create(&LastIndexedInfo{
			ContractAddress: contractAddress.String(),
			ChainID:         chainID,
			BlockNumber:     blockNumber,
		})
	if dbTx.Error != nil {
		return fmt.Errorf("could not update last indexed info: %w", dbTx.Error)
	}
	return nil
}

// RetrieveLastIndexed retrieves the last indexed block number for a contract.
func (s Store) RetrieveLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32) (uint64, error) {
	entry := LastIndexedInfo{}
	dbTx := s.DB().WithContext(ctx).
		Model(&LastIndexedInfo{}).
		Where(&LastIndexedInfo{
			ContractAddress: contractAddress.String(),
			ChainID:         chainID,
		}).
		First(&entry)
	if dbTx.RowsAffected == 0 {
		logger.Warnf("no last indexed info found for contract %s on chain %d. Providing 0.", contractAddress.String(), chainID)
		return 0, nil
	}
	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not retrieve last indexed info: %w", dbTx.Error)
	}
	return entry.BlockNumber, nil
}
