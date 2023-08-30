package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
)

func (t *DBSuite) TestGetSummitBlockNumberForRoot() {
	t.RunOnAllDBs(func(testDB db.GuardDB) {
		// Store two agent roots.
		agentRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		agentRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))

		blockNumberA := gofakeit.Uint64()
		blockNumberB := gofakeit.Uint64()

		err := testDB.StoreAgentRoot(
			t.GetTestContext(),
			agentRootA,
			blockNumberA,
		)
		Nil(t.T(), err)

		err = testDB.StoreAgentRoot(
			t.GetTestContext(),
			agentRootB,
			blockNumberB,
		)
		Nil(t.T(), err)

		// Call GetSummitBlockNumberForRoot for each agent root.
		blockNumber, err := testDB.GetSummitBlockNumberForRoot(t.GetTestContext(), agentRootA.String())
		Nil(t.T(), err)
		Equal(t.T(), blockNumberA, blockNumber)

		blockNumber, err = testDB.GetSummitBlockNumberForRoot(t.GetTestContext(), agentRootB.String())
		Nil(t.T(), err)
		Equal(t.T(), blockNumberB, blockNumber)
	})
}
