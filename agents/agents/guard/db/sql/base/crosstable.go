package base

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/imkira/go-interpol"
	agentTypes "github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// GetUpdateAgentStatusParameters gets the parameters for updating the agent status with the following steps:
// 1. Outer join the `AgentTree` table on the `Dispute` table on the `NotaryAddress` <-> `AgentAddress` fields.
// 2. Filter the rows where the `DisputeProcessedStatus` is `Resolved`.
// 3. Return each of remaining rows' `AgentRoot`, `AgentAddress`, and `Proof` fields.
func (s Store) GetUpdateAgentStatusParameters(ctx context.Context) ([]agentTypes.AgentTree, error) {
	agentTreesTableName, err := dbcommon.GetModelName(s.DB(), &AgentTree{})
	if err != nil {
		return nil, fmt.Errorf("failed to get agent trees table name: %w", err)
	}

	disputesTableName, err := dbcommon.GetModelName(s.DB(), &Dispute{})
	if err != nil {
		return nil, fmt.Errorf("failed to get disputes table name: %w", err)
	}

	query, err := interpol.WithMap(
		`
		SELECT * FROM {agentTreesTable} AS aTable
		JOIN (
			SELECT * FROM {disputesTable} WHERE {disputeProcessedStatus} = ?
		) AS dTable
		ON aTable.{agentAddress} = dTable.{notaryAddress}
		`,
		map[string]string{
			"agentTreesTable":        agentTreesTableName,
			"disputesTable":          disputesTableName,
			"agentAddress":           AgentAddressFieldName,
			"notaryAddress":          NotaryAddressFieldName,
			"disputeProcessedStatus": DisputeProcessedStatusFieldName,
		})
	if err != nil {
		return nil, fmt.Errorf("failed to interpolate query: %w", err)
	}

	var dbAgentTrees []AgentTree
	err = s.DB().WithContext(ctx).Raw(query, agentTypes.Resolved).Scan(&dbAgentTrees).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get agent trees: %w", err)
	}

	// Convert DB fields to agent types.
	agentTrees := []agentTypes.AgentTree{}
	for _, tree := range dbAgentTrees {
		var proofBytes [][32]byte
		err = json.Unmarshal(tree.Proof, &proofBytes)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal proof: %w", err)
		}
		agentTrees = append(agentTrees, agentTypes.AgentTree{
			AgentRoot:    tree.AgentRoot,
			AgentAddress: common.HexToAddress(tree.AgentAddress),
			BlockNumber:  tree.BlockNumber,
			Proof:        proofBytes,
		})
	}
	return agentTrees, nil
}
