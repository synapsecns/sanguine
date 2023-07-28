package base

import (
	"context"
	"fmt"
	"strconv"

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
SELECT * FROM {agentTreesTable}
OUTER JOIN {disputesTable}
ON {agentTreesTable}.{agentAddress} = {disputesTable}.{notaryAddress}
WHERE dTable.{disputeProcessedStatus} = {resolved}
`,
		map[string]string{
			"agentTreesTable":        agentTreesTableName,
			"disputesTable":          disputesTableName,
			"agentAddress":           AgentAddressFieldName,
			"notaryAddress":          NotaryAddressFieldName,
			"disputeProcessedStatus": DisputeProcessedStatusFieldName,
			"resolved":               strconv.Itoa(int(Resolved)),
		})
	if err != nil {
		return nil, fmt.Errorf("failed to interpolate query: %w", err)
	}

	var agentTrees []agentTypes.AgentTree
	err = s.DB().WithContext(ctx).Raw(query).Scan(&agentTrees).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get agent trees: %w", err)
	}
	return agentTrees, nil
}

// GetLatestConfirmedSummitBlockNumber gets the latest "confirmed" summit block number.
// That is, the summit block number corresponding to the most recent updated agent root
// on the given chain.
func (s Store) GetLatestConfirmedSummitBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	var blockNumber uint64
	agentTreesTableName, err := dbcommon.GetModelName(s.DB(), &AgentTree{})
	if err != nil {
		return blockNumber, fmt.Errorf("failed to get agnet trees table name: %w", err)
	}

	agentRootsTableName, err := dbcommon.GetModelName(s.DB(), &Dispute{})
	if err != nil {
		return blockNumber, fmt.Errorf("failed to get agent roots table name: %w", err)
	}

	query, err := interpol.WithMap(
		`
SELECT {blockNumber} FROM {agentTreesTable}
OUTER JOIN {agentRootsTable}
ON {agentTreesTable}.{agentRoot} = {agentRootsTable}.{agentRoot}
WHERE {agentRootsTable}.{chainID} = {chainID}
ORDER BY {agentTreesTable}.{blockNumber} DESC
LIMIT 1
`,
		map[string]string{
			"agentTreesTable": agentTreesTableName,
			"agentRootsTable": agentRootsTableName,
			"blockNumber":     BlockNumberFieldName,
			"agentRoot":       AgentRootFieldName,
			"chainID":         ChainIDFieldName,
		})
	if err != nil {
		return blockNumber, fmt.Errorf("failed to interpolate query: %w", err)
	}

	err = s.DB().WithContext(ctx).Raw(query, chainID).Scan(&blockNumber).Error
	if err != nil {
		return blockNumber, fmt.Errorf("failed to get latest confirmed summit block number: %w", err)
	}
	return blockNumber, nil
}
