package base

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/clause"
)

// StoreAgentTree stores an agent tree.
func (s Store) StoreAgentTree(
	ctx context.Context,
	agentRoot [32]byte,
	agentAddress common.Address,
	blockNumber uint64,
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
				{Name: AgentRootFieldName}, {Name: AgentAddressFieldName},
			},
			DoNothing: true,
		}).
		Create(&AgentTree{
			AgentRoot:    dbAgentRoot,
			AgentAddress: agentAddress.String(),
			BlockNumber:  blockNumber,
			Proof:        proofJSON,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store agent tree: %w", dbTx.Error)
	}

	return nil
}
