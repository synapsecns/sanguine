package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
)

func (t *DBSuite) TestGetUpdateAgentStatusParameters() {
	t.RunOnAllDBs(func(testDB db.GuardDB) {
		addressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		addressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		addressC := common.BigToAddress(big.NewInt(gofakeit.Int64()))

		agentRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		agentRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))

		// Insert two rows into the `AgentTree` table.
		err := testDB.StoreAgentTree(
			t.GetTestContext(),
			agentRootA,
			addressA,
			gofakeit.Uint64(),
			[][32]byte{[32]byte{gofakeit.Uint8()}},
		)
		Nil()
	})
}
