package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/clause"
	"math/big"
)

// StoreDispute stores an dispute.
func (s Store) StoreDispute(
	ctx context.Context,
	agentRoot [32]byte,
	disputeIndex *big.Int,
	resolved bool,
	guardIndex *big.Int,
	guardAddress common.Address,
	notaryIndex *big.Int,
	notaryAddress common.Address,
) error {
	dbAgentRoot := common.BytesToHash(agentRoot[:]).String()

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: AgentRootFieldName}, {Name: DisputeIndexFieldName},
			},
			DoNothing: true,
		}).
		Create(&Dispute{
			AgentRoot:     dbAgentRoot,
			DisputeIndex:  disputeIndex.Uint64(),
			Resolved:      resolved,
			GuardIndex:    guardIndex.Uint64(),
			GuardAddress:  guardAddress.String(),
			NotaryIndex:   notaryIndex.Uint64(),
			NotaryAddress: notaryAddress.String(),
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store dispute: %w", dbTx.Error)
	}

	return nil
}
