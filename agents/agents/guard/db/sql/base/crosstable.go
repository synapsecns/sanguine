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
// 1. Load the `AgentTree` table.
// 2. Filter the rows where the `AgentStatusRelayedState` is `Queued`, and the `Domain` is the given `chainID`.
// 3. Outer join the `AgentTree` table on the `RelayableAgentStatus` table on the `AgentAddress` fields.
// 4. Return all fields in the `AgentTree` table as well as `UpdatedFlag` from the `RelayableAgentStatus` table.
func (s Store) GetRelayableAgentStatuses(ctx context.Context, chainID uint32) ([]agentTypes.AgentTree, error) {
	agentTreesTableName, err := dbcommon.GetModelName(s.DB(), &AgentTree{})
	if err != nil {
		return nil, fmt.Errorf("failed to get agent trees table name: %w", err)
	}

	relayableAgentStatusesTableName, err := dbcommon.GetModelName(s.DB(), &RelayableAgentStatus{})
	if err != nil {
		return nil, fmt.Errorf("failed to get relayable agent statuses table name: %w", err)
	}

	query, err := interpol.WithMap(
		`
		SELECT aTable.*, rTable.{updatedFlag}
		FROM {agentTreesTable} AS aTable
		JOIN (
			SELECT * FROM {relayableAgentStatusesTable}
			WHERE {agentStatusRelayedState} = ?
			AND {domain} = ?
		) AS rTable
		ON aTable.{agentAddress} = rTable.{agentAddress}
		`,
		map[string]string{
			"domain":                      DomainFieldName,
			"updatedFlag":                 UpdatedFlagFieldName,
			"agentTreesTable":             agentTreesTableName,
			"relayableAgentStatusesTable": relayableAgentStatusesTableName,
			"agentStatusRelayedState":     AgentStatusRelayedStateFieldName,
			"agentAddress":                AgentAddressFieldName,
		})
	if err != nil {
		return nil, fmt.Errorf("failed to interpolate query: %w", err)
	}

	var dbAgentTrees []agentTreeWithStatus
	err = s.DB().WithContext(ctx).Raw(query, agentTypes.Queued, chainID).Scan(&dbAgentTrees).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get agent trees: %w", err)
	}

	// Convert db fields to agent types.
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
			AgentDomain:      chainID,
			UpdatedAgentFlag: tree.UpdatedAgentFlag,
			BlockNumber:      tree.BlockNumber,
			Proof:            proofBytes,
		})
	}
	return agentTrees, nil
}
