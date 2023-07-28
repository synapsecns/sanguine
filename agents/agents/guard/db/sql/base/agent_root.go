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
	chainID uint32,
	blockNumber uint64,
) error {
	dbAgentRoot := common.BytesToHash(agentRoot[:]).String()

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: AgentRootFieldName}, {Name: ChainIDFieldName},
			},
			DoNothing: true,
		}).
		Create(&AgentRoot{
			AgentRoot:   dbAgentRoot,
			ChainID:     chainID,
			BlockNumber: blockNumber,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store agent root: %w", dbTx.Error)
	}

	return nil
}
