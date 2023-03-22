package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	agentstypes "github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

func (t *DBSuite) TestGetTimestampForMessage() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		origin := gofakeit.Uint32()
		nonceA := uint32(5)
		nonceB := uint32(10)
		nonceC := uint32(15)

		rootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberA := big.NewInt(int64(gofakeit.Uint32()))
		timestampA := big.NewInt(int64(gofakeit.Uint32()))
		stateA := agentstypes.NewState(rootA, origin, nonceA, blockNumberA, timestampA)

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofA := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}
		treeHeightA := gofakeit.Uint32()

		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberB := big.NewInt(int64(gofakeit.Uint32()))
		timestampB := big.NewInt(int64(gofakeit.Uint32()))
		stateB := agentstypes.NewState(rootB, origin, nonceB, blockNumberB, timestampB)

		snapshotRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofB := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}
		treeHeightB := gofakeit.Uint32()

		rootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberC := big.NewInt(int64(gofakeit.Uint32()))
		timestampC := big.NewInt(int64(gofakeit.Uint32()))
		stateC := agentstypes.NewState(rootC, origin, nonceC, blockNumberC, timestampC)

		snapshotRootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofC := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}
		treeHeightC := gofakeit.Uint32()

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, treeHeightA, 1)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateB, snapshotRootB, proofB, treeHeightB, 2)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateC, snapshotRootC, proofC, treeHeightC, 3)
		Nil(t.T(), err)

		attestationA := agentstypes.NewAttestation(snapshotRootA, uint8(treeHeightA), 1, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
		attestationB := agentstypes.NewAttestation(snapshotRootB, uint8(treeHeightB), 2, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
		attestationC := agentstypes.NewAttestation(snapshotRootC, uint8(treeHeightC), 3, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))

		err = testDB.StoreAttestation(t.GetTestContext(), attestationA, origin+1, 2, 2)
		Nil(t.T(), err)
		err = testDB.StoreAttestation(t.GetTestContext(), attestationB, origin+1, 1, 3)
		Nil(t.T(), err)
		err = testDB.StoreAttestation(t.GetTestContext(), attestationC, origin+1, 3, 1)
		Nil(t.T(), err)

		timestamp, err := testDB.GetTimestampForMessage(t.GetTestContext(), origin, nonceA, "")
		Nil(t.T(), err)
		Equal(t.T(), uint64(3), *timestamp)

		timestamp, err = testDB.GetTimestampForMessage(t.GetTestContext(), origin, nonceB, "")
		Nil(t.T(), err)
		Equal(t.T(), uint64(3), *timestamp)

		timestamp, err = testDB.GetTimestampForMessage(t.GetTestContext(), origin, nonceC, "")
		Nil(t.T(), err)
		Equal(t.T(), uint64(1), *timestamp)
	})
}
