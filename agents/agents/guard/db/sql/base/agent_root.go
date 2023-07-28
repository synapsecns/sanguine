package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
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
				{Name: db.AgentRootFieldName}, {Name: db.ChainIDFieldName},
			},
			DoNothing: true,
		}).
		Create(&db.AgentRoot{
			AgentRoot:   dbAgentRoot,
			ChainID:     chainID,
			BlockNumber: blockNumber,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store agent root: %w", dbTx.Error)
	}

	return nil
}

// GetSummitBlockNumber returns the summit block number for the given agent root.
func (s Store) GetSummitBlockNumber(
	ctx context.Context,
	agentRoot [32]byte,
) (uint64, error) {
	dbAgentRoot := common.BytesToHash(agentRoot[:]).String()

	var agentRootRow db.AgentRoot
	dbTx := s.DB().WithContext(ctx).
		Where(fmt.Sprintf("%s = ?", db.AgentAddressFieldName), dbAgentRoot).
		First(&agentRootRow)

	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get summit block number: %w", dbTx.Error)
	}

	return agentRootRow.BlockNumber, nil
}
