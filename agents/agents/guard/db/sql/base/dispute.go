package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
	"gorm.io/gorm/clause"
	"math/big"
)

// StoreDispute stores an dispute.
func (s Store) StoreDispute(
	ctx context.Context,
	disputeIndex *big.Int,
	disputeProcessedStatus guard.DisputeProcessedStatus,
	guardAddress common.Address,
	notaryIndex uint32,
	notaryAddress common.Address,
) error {
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: db.AgentRootFieldName}, {Name: db.DisputeIndexFieldName},
			},
			DoNothing: true,
		}).
		Create(&db.Dispute{
			DisputeIndex:           disputeIndex.Uint64(),
			DisputeProcessedStatus: disputeProcessedStatus,
			GuardAddress:           guardAddress.String(),
			NotaryIndex:            uint64(notaryIndex),
			NotaryAddress:          notaryAddress.String(),
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store dispute: %w", dbTx.Error)
	}

	return nil
}

// GetDisputesToPropagate returns disputes with `DisputeProcessStatus` equal to `Resolved`.
func (s Store) GetDisputesToPropagate(
	ctx context.Context,
) ([]db.Dispute, error) {
	var disputes []db.Dispute
	dbTx := s.DB().WithContext(ctx).
		Where(fmt.Sprintf("%s = ?", db.DisputeProcessedStatusFieldName), guard.Resolved).
		Find(&disputes)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get disputes to propagate: %w", dbTx.Error)
	}

	return disputes, nil
}
