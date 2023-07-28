package base

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	agentTypes "github.com/synapsecns/sanguine/agents/types"
	"gorm.io/gorm/clause"
)

// StoreDispute stores an dispute.
func (s Store) StoreDispute(
	ctx context.Context,
	disputeIndex *big.Int,
	disputeProcessedStatus agentTypes.DisputeProcessedStatus,
	guardAddress common.Address,
	notaryIndex uint32,
	notaryAddress common.Address,
) error {
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: AgentRootFieldName}, {Name: DisputeIndexFieldName},
			},
			DoNothing: true,
		}).
		Create(&Dispute{
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
