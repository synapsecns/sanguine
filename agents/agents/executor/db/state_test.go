package db_test

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
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

		gasPriceA := gofakeit.Uint16()
		dataPriceA := gofakeit.Uint16()
		execBufferA := gofakeit.Uint16()
		amortAttCostA := gofakeit.Uint16()
		etherPriceA := gofakeit.Uint16()
		markupA := gofakeit.Uint16()
		gasDataA := agentstypes.NewGasData(gasPriceA, dataPriceA, execBufferA, amortAttCostA, etherPriceA, markupA)

		stateA := agentstypes.NewState(rootA, originA, nonceA, blockNumberA, timestampA, gasDataA)

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofA := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, 1, blockNumberA.Uint64())
		Nil(t.T(), err)

		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		originB := gofakeit.Uint32()
		nonceB := gofakeit.Uint32()
		blockNumberB := big.NewInt(int64(gofakeit.Uint32()))
		timestampB := big.NewInt(int64(gofakeit.Uint32()))

		gasPriceB := gofakeit.Uint16()
		dataPriceB := gofakeit.Uint16()
		execBufferB := gofakeit.Uint16()
		amortAttCostB := gofakeit.Uint16()
		etherPriceB := gofakeit.Uint16()
		markupB := gofakeit.Uint16()
		gasDataB := agentstypes.NewGasData(gasPriceB, dataPriceB, execBufferB, amortAttCostB, etherPriceB, markupB)

		stateB := agentstypes.NewState(rootB, originB, nonceB, blockNumberB, timestampB, gasDataB)

		snapshotRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofB := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		err = testDB.StoreState(t.GetTestContext(), stateB, snapshotRootB, proofB, 2, blockNumberB.Uint64())
		Nil(t.T(), err)

		snapshotRootAString := snapshotRootA.String()
		rootAString := rootA.String()
		stateMaskA := db.DBState{
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
		stateMaskB := db.DBState{
			OriginBlockNumber: &blockNumberBUint64,
			OriginTimestamp:   &timestampBUint64,
			Proof:             &proofBJSON,
		}

		retrievedStateB, err := testDB.GetState(t.GetTestContext(), stateMaskB)
		Nil(t.T(), err)

		Equal(t.T(), stateB, *retrievedStateB)

		rootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		originC := gofakeit.Uint32()
		nonceC := gofakeit.Uint32()
		blockNumberC := big.NewInt(int64(gofakeit.Uint32()))
		timestampC := big.NewInt(int64(gofakeit.Uint32()))

		gasPriceC := gofakeit.Uint16()
		dataPriceC := gofakeit.Uint16()
		execBufferC := gofakeit.Uint16()
		amortAttCostC := gofakeit.Uint16()
		etherPriceC := gofakeit.Uint16()
		markupC := gofakeit.Uint16()
		gasDataC := agentstypes.NewGasData(gasPriceC, dataPriceC, execBufferC, amortAttCostC, etherPriceC, markupC)

		stateC := agentstypes.NewState(rootC, originC, nonceC, blockNumberC, timestampC, gasDataC)

		proofC := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		rootD := common.BigToHash(big.NewInt(gofakeit.Int64()))
		originD := gofakeit.Uint32()
		nonceD := gofakeit.Uint32()
		blockNumberD := big.NewInt(int64(gofakeit.Uint32()))
		timestampD := big.NewInt(int64(gofakeit.Uint32()))

		gasPriceD := gofakeit.Uint16()
		dataPriceD := gofakeit.Uint16()
		execBufferD := gofakeit.Uint16()
		amortAttCostD := gofakeit.Uint16()
		etherPriceD := gofakeit.Uint16()
		markupD := gofakeit.Uint16()
		gasDataD := agentstypes.NewGasData(gasPriceD, dataPriceD, execBufferD, amortAttCostD, etherPriceD, markupD)

		stateD := agentstypes.NewState(rootD, originD, nonceD, blockNumberD, timestampD, gasDataD)

		proofD := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		snapshotRootCD := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofsCD := [][][]byte{proofC, proofD}

		err = testDB.StoreStates(t.GetTestContext(), []agentstypes.State{stateC, stateD}, snapshotRootCD, proofsCD, blockNumberC.Uint64())
		Nil(t.T(), err)

		snapshotRootCDString := snapshotRootCD.String()
		rootCString := rootC.String()
		stateMaskC := db.DBState{
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
		stateMaskD := db.DBState{
			OriginBlockNumber: &blockNumberDUint64,
			OriginTimestamp:   &timestampDUint64,
			Proof:             &proofDJSON,
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

		gasPriceA := gofakeit.Uint16()
		dataPriceA := gofakeit.Uint16()
		execBufferA := gofakeit.Uint16()
		amortAttCostA := gofakeit.Uint16()
		etherPriceA := gofakeit.Uint16()
		markupA := gofakeit.Uint16()
		gasDataA := agentstypes.NewGasData(gasPriceA, dataPriceA, execBufferA, amortAttCostA, etherPriceA, markupA)

		stateA := agentstypes.NewState(rootA, originA, nonceA, blockNumberA, timestampA, gasDataA)

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofA := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}
		indexA := gofakeit.Uint32()

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, indexA, blockNumberA.Uint64())
		Nil(t.T(), err)

		stateMask := db.DBState{
			ChainID: &originA,
		}

		snapshotRoot, proof, stateIndex, err := testDB.GetStateMetadata(t.GetTestContext(), stateMask)
		Nil(t.T(), err)

		proofBytes, err := json.Marshal(proof)
		Nil(t.T(), err)
		var proofABytes [][]byte
		err = json.Unmarshal(proofBytes, &proofABytes)
		Nil(t.T(), err)

		Equal(t.T(), snapshotRootA, common.BytesToHash((*snapshotRoot)[:]))
		Equal(t.T(), proofA, proofABytes)
		Equal(t.T(), indexA, *stateIndex)
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

		gasPriceA := gofakeit.Uint16()
		dataPriceA := gofakeit.Uint16()
		execBufferA := gofakeit.Uint16()
		amortAttCostA := gofakeit.Uint16()
		etherPriceA := gofakeit.Uint16()
		markupA := gofakeit.Uint16()
		gasDataA := agentstypes.NewGasData(gasPriceA, dataPriceA, execBufferA, amortAttCostA, etherPriceA, markupA)

		stateA := agentstypes.NewState(rootA, origin, nonceA, blockNumberA, timestampA, gasDataA)

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofA := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberB := big.NewInt(int64(gofakeit.Uint32()))
		timestampB := big.NewInt(int64(gofakeit.Uint32()))

		gasPriceB := gofakeit.Uint16()
		dataPriceB := gofakeit.Uint16()
		execBufferB := gofakeit.Uint16()
		amortAttCostB := gofakeit.Uint16()
		etherPriceB := gofakeit.Uint16()
		markupB := gofakeit.Uint16()
		gasDataB := agentstypes.NewGasData(gasPriceB, dataPriceB, execBufferB, amortAttCostB, etherPriceB, markupB)

		stateB := agentstypes.NewState(rootB, origin, nonceB, blockNumberB, timestampB, gasDataB)

		snapshotRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofB := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		rootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberC := big.NewInt(int64(gofakeit.Uint32()))
		timestampC := big.NewInt(int64(gofakeit.Uint32()))

		gasPriceC := gofakeit.Uint16()
		dataPriceC := gofakeit.Uint16()
		execBufferC := gofakeit.Uint16()
		amortAttCostC := gofakeit.Uint16()
		etherPriceC := gofakeit.Uint16()
		markupC := gofakeit.Uint16()
		gasDataC := agentstypes.NewGasData(gasPriceC, dataPriceC, execBufferC, amortAttCostC, etherPriceC, markupC)

		stateC := agentstypes.NewState(rootC, origin, nonceC, blockNumberC, timestampC, gasDataC)

		snapshotRootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofC := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, 1, 1)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateB, snapshotRootB, proofB, 2, 1)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateC, snapshotRootC, proofC, 3, 1)
		Nil(t.T(), err)

		potentialSnapshotRoots, err := testDB.GetPotentialSnapshotRoots(t.GetTestContext(), origin, 6)
		Nil(t.T(), err)

		Equal(t.T(), 2, len(potentialSnapshotRoots))
		NotEqual(t.T(), potentialSnapshotRoots[0], potentialSnapshotRoots[1])
		True(t.T(), snapshotRootB.String() == potentialSnapshotRoots[0] || snapshotRootB.String() == potentialSnapshotRoots[1])
		True(t.T(), snapshotRootC.String() == potentialSnapshotRoots[0] || snapshotRootC.String() == potentialSnapshotRoots[1])
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

		gasPriceA := gofakeit.Uint16()
		dataPriceA := gofakeit.Uint16()
		execBufferA := gofakeit.Uint16()
		amortAttCostA := gofakeit.Uint16()
		etherPriceA := gofakeit.Uint16()
		markupA := gofakeit.Uint16()
		gasDataA := agentstypes.NewGasData(gasPriceA, dataPriceA, execBufferA, amortAttCostA, etherPriceA, markupA)

		stateA := agentstypes.NewState(rootA, origin, nonceA, blockNumberA, timestampA, gasDataA)

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofA := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberB := big.NewInt(int64(gofakeit.Uint32()))
		timestampB := big.NewInt(int64(gofakeit.Uint32()))

		gasPriceB := gofakeit.Uint16()
		dataPriceB := gofakeit.Uint16()
		execBufferB := gofakeit.Uint16()
		amortAttCostB := gofakeit.Uint16()
		etherPriceB := gofakeit.Uint16()
		markupB := gofakeit.Uint16()
		gasDataB := agentstypes.NewGasData(gasPriceB, dataPriceB, execBufferB, amortAttCostB, etherPriceB, markupB)

		stateB := agentstypes.NewState(rootB, origin, nonceB, blockNumberB, timestampB, gasDataB)

		snapshotRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofB := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		rootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberC := big.NewInt(int64(gofakeit.Uint32()))
		timestampC := big.NewInt(int64(gofakeit.Uint32()))

		gasPriceC := gofakeit.Uint16()
		dataPriceC := gofakeit.Uint16()
		execBufferC := gofakeit.Uint16()
		amortAttCostC := gofakeit.Uint16()
		etherPriceC := gofakeit.Uint16()
		markupC := gofakeit.Uint16()
		gasDataC := agentstypes.NewGasData(gasPriceC, dataPriceC, execBufferC, amortAttCostC, etherPriceC, markupC)

		stateC := agentstypes.NewState(rootC, origin, nonceC, blockNumberC, timestampC, gasDataC)

		snapshotRootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofC := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		rootD := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberD := big.NewInt(int64(gofakeit.Uint32()))
		timestampD := big.NewInt(int64(gofakeit.Uint32()))

		gasPriceD := gofakeit.Uint16()
		dataPriceD := gofakeit.Uint16()
		execBufferD := gofakeit.Uint16()
		amortAttCostD := gofakeit.Uint16()
		etherPriceD := gofakeit.Uint16()
		markupD := gofakeit.Uint16()
		gasDataD := agentstypes.NewGasData(gasPriceD, dataPriceD, execBufferD, amortAttCostD, etherPriceD, markupD)

		stateD := agentstypes.NewState(rootD, origin, nonceD, blockNumberD, timestampD, gasDataD)

		snapshotRootD := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofD := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, 1, 1)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateB, snapshotRootB, proofB, 2, 1)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateC, snapshotRootC, proofC, 3, 1)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateD, snapshotRootD, proofD, 4, 1)
		Nil(t.T(), err)

		potentialSnapshotRoots, err := testDB.GetSnapshotRootsInNonceRange(t.GetTestContext(), origin, 6, 15)
		Nil(t.T(), err)

		Equal(t.T(), 2, len(potentialSnapshotRoots))
		NotEqual(t.T(), potentialSnapshotRoots[0], potentialSnapshotRoots[1])
		True(t.T(), snapshotRootB.String() == potentialSnapshotRoots[0] || snapshotRootB.String() == potentialSnapshotRoots[1])
		True(t.T(), snapshotRootC.String() == potentialSnapshotRoots[0] || snapshotRootC.String() == potentialSnapshotRoots[1])
	})
}
