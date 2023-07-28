package base

import (
	"context"
	"fmt"
	"github.com/imkira/go-interpol"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// GetUpdateAgentStatusParameters gets the parameters for updating the agent status with the following steps:
// 1. Outer join the `AgentTree` table on the `Dispute` table on the `NotaryAddress` <-> `AgentAddress` fields.
// 2. Filter the rows where the `DisputeProcessedStatus` is `Resolved`.
// 3. Return each of remaining rows' `AgentRoot`, `AgentAddress`, and `Proof` fields.
func (s Store) GetUpdateAgentStatusParameters(ctx context.Context) ([]db.AgentRoot, error) {
	agentTreesTableName, err := dbcommon.GetModelName(s.DB(), &db.AgentTree{})
	if err != nil {
		return nil, fmt.Errorf("failed to get agnet trees table name: %w", err)
	}

	disputesTableName, err := dbcommon.GetModelName(s.DB(), &db.Dispute{})
	if err != nil {
		return nil, fmt.Errorf("failed to get disputes table name: %w", err)
	}

	query, err := interpol.WithMap(
		`
(SELECT * FROM {agentTreesTable}) AS atTable
FULL OUTER JOIN
(SELECT * FROM {disputesTable}) AS dTable
ON atTable.{agentAddress} = dTable.{notaryAddress}
`,
		map[string]string{})
}
