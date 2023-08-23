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
				{Name: DisputeIndexFieldName},
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

// UpdateDisputeProcessedStatus updates the disputed processed status for a dispute.
func (s Store) UpdateDisputeProcessedStatus(
	ctx context.Context,
	guardAddress *common.Address,
	notaryAddress *common.Address,
	flag agentTypes.DisputeProcessedStatus,
) error {
	disputeMask := Dispute{
		DisputeProcessedStatus: agentTypes.Opened,
	}
	if guardAddress != nil {
		disputeMask.GuardAddress = guardAddress.String()
	} else if notaryAddress != nil {
		disputeMask.NotaryAddress = notaryAddress.String()
	} else {
		return fmt.Errorf("guardAddress or notaryAddress must be set")
	}

	dbTx := s.DB().WithContext(ctx).Debug().
		Model(&Dispute{}).
		Where(disputeMask).
		Update(DisputeProcessedStatusFieldName, flag)
	if dbTx.Error != nil {
		return fmt.Errorf("failed to update dispute processed status %w", dbTx.Error)
	}
	return nil
}
