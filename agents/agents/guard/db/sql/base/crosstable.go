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

type agentTreeWithStatus struct {
	AgentTree
	AgentDomain      uint32
	UpdatedAgentFlag agentTypes.AgentFlagType
}

// GetRelayableAgentStatuses gets the parameters for updating the agent status with the following steps:
// 1. Outer join the `AgentTree` table on the `RelayableAgentStatus` table on the `AgentAddress` fields.
// 2. Filter the rows where the `AgentStatusRelayedState` is `Queued`.
// 3. Return each of remaining rows' `AgentRoot`, `AgentAddress`, and `Proof` fields.
func (s Store) GetRelayableAgentStatuses(ctx context.Context, chainID uint32) ([]agentTypes.AgentTree, error) {
	agentTreesTableName, err := dbcommon.GetModelName(s.DB(), &AgentTree{})
	if err != nil {
		return nil, fmt.Errorf("failed to get agent trees table name: %w", err)
	}

	relayableAgentStatusesTableName, err := dbcommon.GetModelName(s.DB(), &RelayableAgentStatus{})
	if err != nil {
		return nil, fmt.Errorf("failed to get disputes table name: %w", err)
	}

	query, err := interpol.WithMap(
		`
		SELECT aTable.*, rTable.{agentDomain}, rTable.{updatedAgentFlag}
		FROM {agentTreesTable} AS aTable
		JOIN (
			SELECT * FROM {relayableAgentStatusesTable}
			WHERE {agentStatusRelayedState} = ?
			AND {agentDomain} = ?
		) AS rTable
		ON aTable.{agentAddress} = rTable.{agentAddress}
		`,
		map[string]string{
			"agentTreesTable":             agentTreesTableName,
			"relayableAgentStatusesTable": relayableAgentStatusesTableName,
			"agentAddress":                AgentAddressFieldName,
			"agentStatusRelayedState":     AgentStatusRelayedStateFieldName,
			"agentDomain":                 AgentDomainFieldName,
		})
	if err != nil {
		return nil, fmt.Errorf("failed to interpolate query: %w", err)
	}

	var dbAgentTrees []agentTreeWithStatus
	err = s.DB().WithContext(ctx).Raw(query, agentTypes.Queued, chainID).Scan(&dbAgentTrees).Error
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
			AgentRoot:        tree.AgentRoot,
			AgentAddress:     common.HexToAddress(tree.AgentAddress),
			AgentDomain:      tree.AgentDomain,
			UpdatedAgentFlag: tree.UpdatedAgentFlag,
			BlockNumber:      tree.BlockNumber,
			Proof:            proofBytes,
		})
	}
	return agentTrees, nil
}
