package db_test

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentstypes "github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

func (t *DBSuite) TestStoreRetrieveState() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		rootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		originA := gofakeit.Uint32()
		nonceA := gofakeit.Uint32()
		blockNumberA := big.NewInt(int64(gofakeit.Uint32()))
		timestampA := big.NewInt(int64(gofakeit.Uint32()))
		stateA := agentstypes.NewState(rootA, originA, nonceA, blockNumberA, timestampA)

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofA := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}
		treeHeightA := gofakeit.Uint32()

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, treeHeightA)
		Nil(t.T(), err)

		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		originB := gofakeit.Uint32()
		nonceB := gofakeit.Uint32()
		blockNumberB := big.NewInt(int64(gofakeit.Uint32()))
		timestampB := big.NewInt(int64(gofakeit.Uint32()))
		stateB := agentstypes.NewState(rootB, originB, nonceB, blockNumberB, timestampB)

		snapshotRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofB := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}
		treeHeightB := gofakeit.Uint32()

		err = testDB.StoreState(t.GetTestContext(), stateB, snapshotRootB, proofB, treeHeightB)
		Nil(t.T(), err)

		snapshotRootAString := snapshotRootA.String()
		rootAString := rootA.String()
		stateMaskA := types.DBState{
			SnapshotRoot: &snapshotRootAString,
			Root:         &rootAString,
			ChainID:      &originA,
			Nonce:        &nonceA,
		}

		retrievedStateA, err := testDB.GetState(t.GetTestContext(), stateMaskA)
		Nil(t.T(), err)

		Equal(t.T(), stateA, *retrievedStateA)

		blockNumberBUint64 := blockNumberB.Uint64()
		timestampBUint64 := timestampB.Uint64()
		proofBBytes, err := json.Marshal(proofB)
		Nil(t.T(), err)
		proofBJSON := json.RawMessage(proofBBytes)
		stateMaskB := types.DBState{
			OriginBlockNumber: &blockNumberBUint64,
			OriginTimestamp:   &timestampBUint64,
			Proof:             &proofBJSON,
			TreeHeight:        &treeHeightB,
		}

		retrievedStateB, err := testDB.GetState(t.GetTestContext(), stateMaskB)
		Nil(t.T(), err)

		Equal(t.T(), stateB, *retrievedStateB)

		rootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		originC := gofakeit.Uint32()
		nonceC := gofakeit.Uint32()
		blockNumberC := big.NewInt(int64(gofakeit.Uint32()))
		timestampC := big.NewInt(int64(gofakeit.Uint32()))
		stateC := agentstypes.NewState(rootC, originC, nonceC, blockNumberC, timestampC)

		proofC := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		rootD := common.BigToHash(big.NewInt(gofakeit.Int64()))
		originD := gofakeit.Uint32()
		nonceD := gofakeit.Uint32()
		blockNumberD := big.NewInt(int64(gofakeit.Uint32()))
		timestampD := big.NewInt(int64(gofakeit.Uint32()))
		stateD := agentstypes.NewState(rootD, originD, nonceD, blockNumberD, timestampD)

		proofD := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		snapshotRootCD := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofsCD := [][][]byte{proofC, proofD}
		treeHeightCD := gofakeit.Uint32()

		err = testDB.StoreStates(t.GetTestContext(), []agentstypes.State{stateC, stateD}, snapshotRootCD, proofsCD, treeHeightCD)
		Nil(t.T(), err)

		snapshotRootCDString := snapshotRootCD.String()
		rootCString := rootC.String()
		stateMaskC := types.DBState{
			SnapshotRoot: &snapshotRootCDString,
			Root:         &rootCString,
			ChainID:      &originC,
			Nonce:        &nonceC,
		}

		retrievedStateC, err := testDB.GetState(t.GetTestContext(), stateMaskC)
		Nil(t.T(), err)

		Equal(t.T(), stateC, *retrievedStateC)

		blockNumberDUint64 := blockNumberD.Uint64()
		timestampDUint64 := timestampD.Uint64()
		proofDBytes, err := json.Marshal(proofD)
		Nil(t.T(), err)
		proofDJSON := json.RawMessage(proofDBytes)
		stateMaskD := types.DBState{
			OriginBlockNumber: &blockNumberDUint64,
			OriginTimestamp:   &timestampDUint64,
			Proof:             &proofDJSON,
			TreeHeight:        &treeHeightCD,
		}

		retrievedStateD, err := testDB.GetState(t.GetTestContext(), stateMaskD)
		Nil(t.T(), err)

		Equal(t.T(), stateD, *retrievedStateD)
	})
}

func (t *DBSuite) TestGetStateMetadata() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		rootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		originA := gofakeit.Uint32()
		nonceA := gofakeit.Uint32()
		blockNumberA := big.NewInt(int64(gofakeit.Uint32()))
		timestampA := big.NewInt(int64(gofakeit.Uint32()))
		stateA := agentstypes.NewState(rootA, originA, nonceA, blockNumberA, timestampA)

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofA := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}
		treeHeightA := gofakeit.Uint32()

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, treeHeightA)
		Nil(t.T(), err)

		stateMask := types.DBState{
			ChainID: &originA,
		}

		snapshotRoot, proof, treeHeight, err := testDB.GetStateMetadata(t.GetTestContext(), stateMask)
		Nil(t.T(), err)

		proofBytes, err := json.Marshal(proof)
		Nil(t.T(), err)
		var proofABytes [][]byte
		err = json.Unmarshal(proofBytes, &proofABytes)
		Nil(t.T(), err)

		Equal(t.T(), snapshotRootA, common.BytesToHash((*snapshotRoot)[:]))
		Equal(t.T(), proofA, proofABytes)
		Equal(t.T(), treeHeightA, *treeHeight)
	})
}

func (t *DBSuite) TestGetPotentialSnapshotRoots() {
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

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, treeHeightA)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateB, snapshotRootB, proofB, treeHeightB)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateC, snapshotRootC, proofC, treeHeightC)
		Nil(t.T(), err)

		potentialSnapshotRoots, err := testDB.GetPotentialSnapshotRoots(t.GetTestContext(), origin, 6)
		Nil(t.T(), err)

		Equal(t.T(), 2, len(potentialSnapshotRoots))
		NotEqual(t.T(), potentialSnapshotRoots[0], potentialSnapshotRoots[1])
		True(t.T(), snapshotRootA.String() == potentialSnapshotRoots[0] || snapshotRootA.String() == potentialSnapshotRoots[1])
		True(t.T(), snapshotRootB.String() == potentialSnapshotRoots[0] || snapshotRootB.String() == potentialSnapshotRoots[1])
	})
}

func (t *DBSuite) TestGetSnapshotRootsInNonceRange() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		origin := gofakeit.Uint32()
		nonceA := uint32(5)
		nonceB := uint32(10)
		nonceC := uint32(15)
		nonceD := uint32(20)

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

		rootD := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberD := big.NewInt(int64(gofakeit.Uint32()))
		timestampD := big.NewInt(int64(gofakeit.Uint32()))
		stateD := agentstypes.NewState(rootD, origin, nonceD, blockNumberD, timestampD)

		snapshotRootD := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofD := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}
		treeHeightD := gofakeit.Uint32()

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, treeHeightA)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateB, snapshotRootB, proofB, treeHeightB)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateC, snapshotRootC, proofC, treeHeightC)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateD, snapshotRootD, proofD, treeHeightD)
		Nil(t.T(), err)

		potentialSnapshotRoots, err := testDB.GetSnapshotRootsInNonceRange(t.GetTestContext(), origin, 6, 15)
		Nil(t.T(), err)

		Equal(t.T(), 2, len(potentialSnapshotRoots))
		NotEqual(t.T(), potentialSnapshotRoots[0], potentialSnapshotRoots[1])
		True(t.T(), snapshotRootA.String() == potentialSnapshotRoots[0] || snapshotRootA.String() == potentialSnapshotRoots[1])
		True(t.T(), snapshotRootB.String() == potentialSnapshotRoots[0] || snapshotRootB.String() == potentialSnapshotRoots[1])
	})
}
