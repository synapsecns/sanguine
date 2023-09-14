package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
	"github.com/synapsecns/sanguine/agents/types"
)

func (t *DBSuite) TestGetRelayableAgentStatuses() {
	// Test by setting up multiple addresses and agent roots in the database. Then, once each agent tree is stored with
	// the agent root and address, get the agent trees that match up with a specific chain ID.
	t.RunOnAllDBs(func(testDB db.GuardDB) {
		addressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		addressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		addressC := common.BigToAddress(big.NewInt(gofakeit.Int64()))

		agentRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		agentRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		agentRootC := common.BigToHash(big.NewInt(gofakeit.Int64()))

		// Insert three rows into the `AgentTree` table.
		err := testDB.StoreAgentTree(
			t.GetTestContext(),
			agentRootA,
			addressA,
			gofakeit.Uint64(),
			[][32]byte{{gofakeit.Uint8()}},
		)
		Nil(t.T(), err)
		err = testDB.StoreAgentTree(
			t.GetTestContext(),
			agentRootB,
			addressB,
			gofakeit.Uint64(),
			[][32]byte{{gofakeit.Uint8()}},
		)
		Nil(t.T(), err)
		err = testDB.StoreAgentTree(
			t.GetTestContext(),
			agentRootC,
			addressC,
			gofakeit.Uint64(),
			[][32]byte{{gofakeit.Uint8()}},
		)
		Nil(t.T(), err)

		// Insert three rows into `RelayableAgentStatus`, two will have matching agent address to `AgentTree` rows and with status `Queued`.
		chainA := gofakeit.Uint32()
		chainB := chainA + 1
		err = testDB.StoreRelayableAgentStatus(
			t.GetTestContext(),
			addressA,
			types.AgentFlagUnknown,
			types.AgentFlagActive,
			chainA,
		)
		Nil(t.T(), err)
		err = testDB.StoreRelayableAgentStatus(
			t.GetTestContext(),
			addressB,
			types.AgentFlagUnknown,
			types.AgentFlagActive,
			chainA,
		)
		Nil(t.T(), err)
		err = testDB.StoreRelayableAgentStatus(
			t.GetTestContext(),
			addressC,
			types.AgentFlagUnknown,
			types.AgentFlagActive,
			chainB,
		)
		Nil(t.T(), err)

		// Get the matching agent tree from the database.
		agentTrees, err := testDB.GetRelayableAgentStatuses(t.GetTestContext(), chainA)
		Nil(t.T(), err)

		Equal(t.T(), 2, len(agentTrees))
	})
}
