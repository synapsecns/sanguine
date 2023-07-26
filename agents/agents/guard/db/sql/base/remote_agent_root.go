package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/clause"
)

// StoreRemoteAgentRoot stores a remote agent root.
func (s Store) StoreRemoteAgentRoot(ctx context.Context, agentRoot [32]byte, chainID uint32) error {
	dbAgentRoot := common.BytesToHash(agentRoot[:]).String()

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: AgentRootFieldName}, {Name: ChainIDFieldName},
			},
			DoNothing: true,
		}).
		Create(&RemoteAgentRoot{
			AgentRoot: dbAgentRoot,
			ChainID:   chainID,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store remote agent root: %w", dbTx.Error)
	}

	return nil
}
