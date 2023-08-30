package base

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/clause"
)

// StoreAgentRoot stores an agent root.
func (s Store) StoreAgentRoot(
	ctx context.Context,
	agentRoot [32]byte,
	blockNumber uint64,
) error {
	dbAgentRoot := common.BytesToHash(agentRoot[:]).String()

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: AgentRootFieldName},
			},
			DoNothing: true,
		}).
		Create(&AgentRoot{
			AgentRoot:   dbAgentRoot,
			BlockNumber: blockNumber,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store agent root: %w", dbTx.Error)
	}

	return nil
}

// GetSummitBlockNumberForRoot gets the summit block number for a given agent root.
func (s Store) GetSummitBlockNumberForRoot(ctx context.Context, agentRoot string) (uint64, error) {
	var blockNumber uint64
	dbTx := s.DB().WithContext(ctx).
		Where(fmt.Sprintf("%s = ?", AgentRootFieldName), agentRoot).
		Order(fmt.Sprintf("%s ASC", BlockNumberFieldName)).
		Limit(1).
		Model(&AgentRoot{}).
		Pluck(BlockNumberFieldName, &blockNumber)

	if dbTx.Error != nil {
		return blockNumber, fmt.Errorf("failed to get summit block number for root: %w", dbTx.Error)
	}

	return blockNumber, nil
}
