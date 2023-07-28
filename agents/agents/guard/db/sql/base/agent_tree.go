package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
	"gorm.io/gorm/clause"
)

// StoreAgentTree stores an agent tree.
func (s Store) StoreAgentTree(
	ctx context.Context,
	agentRoot [32]byte,
	agentAddress common.Address,
	proof [][32]byte,
) error {
	dbAgentRoot := common.BytesToHash(agentRoot[:]).String()

	proofJSON, err := json.Marshal(proof)
	if err != nil {
		return fmt.Errorf("failed to marshal proof: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: db.AgentRootFieldName}, {Name: db.AgentAddressFieldName},
			},
			DoNothing: true,
		}).
		Create(&db.AgentTree{
			AgentRoot:    dbAgentRoot,
			AgentAddress: agentAddress.String(),
			Proof:        proofJSON,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store agent tree: %w", dbTx.Error)
	}

	return nil
}

// SeenLaterAgentRoot
