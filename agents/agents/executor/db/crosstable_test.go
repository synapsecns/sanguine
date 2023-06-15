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

		gasPriceA := gofakeit.Uint16()
		dataPriceA := gofakeit.Uint16()
		execBufferA := gofakeit.Uint16()
		amortAttCostA := gofakeit.Uint16()
		etherPriceA := gofakeit.Uint16()
		markupA := gofakeit.Uint16()
		gasDataA := agentstypes.NewGasData(gasPriceA, dataPriceA, execBufferA, amortAttCostA, etherPriceA, markupA)

		stateA := agentstypes.NewState(rootA, origin, nonceA, blockNumberA, timestampA, gasDataA)

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		agentRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
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
		agentRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
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
		agentRootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		proofC := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

		// Store a state with a nonce of 5, 10, and 15. (The other fields are not checked in the query we are testing).
		err := testDB.StoreState(t.GetTestContext(), stateA, snapshotRootA, proofA, 1, 1)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateB, snapshotRootB, proofB, 2, 2)
		Nil(t.T(), err)
		err = testDB.StoreState(t.GetTestContext(), stateC, snapshotRootC, proofC, 3, 3)
		Nil(t.T(), err)

		attestationA := agentstypes.NewAttestation(snapshotRootA, agentRootA, 1, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
		attestationB := agentstypes.NewAttestation(snapshotRootB, agentRootB, 2, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
		attestationC := agentstypes.NewAttestation(snapshotRootC, agentRootC, 3, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))

		// Store attestations associated with each state via snapshot root. (stateA to attestationA, etc.)
		err = testDB.StoreAttestation(t.GetTestContext(), attestationA, origin+1, 2, 2)
		Nil(t.T(), err)
		err = testDB.StoreAttestation(t.GetTestContext(), attestationB, origin+1, 1, 3)
		Nil(t.T(), err)
		err = testDB.StoreAttestation(t.GetTestContext(), attestationC, origin+1, 3, 1)
		Nil(t.T(), err)

		// We want to get the timestamp of the attestation that has the earliest `destinationBlockNumber` and has a nonce
		// greater than or equal to nonceA (5). This would be attestationB since it has a nonce of 10 and a
		// `destinationBlockNumber` of 1. Because of this, we should get attestationB's timestamp of 3.
		retrievedTimestampA, err := testDB.GetTimestampForMessage(t.GetTestContext(), origin, origin+1, nonceA)
		Nil(t.T(), err)
		Equal(t.T(), uint64(3), *retrievedTimestampA)

		// We want the timestamp of attestationB again here, since we are checking for nonce 10, and attestation has a
		// nonce of 10 and a `destinationBlockNumber` of 1. Because of this, we should get attestationB's timestamp of 3.
		retrievedTimestampB, err := testDB.GetTimestampForMessage(t.GetTestContext(), origin, origin+1, nonceB)
		Nil(t.T(), err)
		Equal(t.T(), uint64(3), *retrievedTimestampB)

		// We want the timestamp of attestationC because that is the only attestation that has a nonce greater than or
		// equal to nonceC (15). We expect to get attestationC's `destinationTimestamp` of 1.
		retrievedTimestampC, err := testDB.GetTimestampForMessage(t.GetTestContext(), origin, origin+1, nonceC)
		Nil(t.T(), err)
		Equal(t.T(), uint64(1), *retrievedTimestampC)
	})
}

func (t *DBSuite) TestGetEarliestStateInRange() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		origin := gofakeit.Uint32()
		var snapshotRoots, agentRoots []common.Hash
		for i := uint32(1); i <= 6; i++ {
			root := common.BigToHash(big.NewInt(gofakeit.Int64()))
			blockNumber := big.NewInt(int64(gofakeit.Uint32()))
			timestamp := big.NewInt(int64(gofakeit.Uint32()))

			gasPrice := gofakeit.Uint16()
			dataPrice := gofakeit.Uint16()
			execBuffer := gofakeit.Uint16()
			amortAttCost := gofakeit.Uint16()
			etherPrice := gofakeit.Uint16()
			markup := gofakeit.Uint16()
			gasData := agentstypes.NewGasData(gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)

			state := agentstypes.NewState(root, origin, i, blockNumber, timestamp, gasData)

			snapshotRoots = append(snapshotRoots, common.BigToHash(big.NewInt(gofakeit.Int64())))
			agentRoots = append(agentRoots, common.BigToHash(big.NewInt(gofakeit.Int64())))
			proof := [][]byte{[]byte(gofakeit.Word()), []byte(gofakeit.Word())}

			err := testDB.StoreState(t.GetTestContext(), state, snapshotRoots[i-1], proof, 1, 1)
			Nil(t.T(), err)
		}

		// Attestation for state 2.
		attestation0 := agentstypes.NewAttestation(snapshotRoots[1], agentRoots[1], 1, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))

		// Attestation for state 4.
		attestation1 := agentstypes.NewAttestation(snapshotRoots[3], agentRoots[2], 2, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))

		// Attestation for state 5.
		attestation2 := agentstypes.NewAttestation(snapshotRoots[4], agentRoots[3], 3, big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))

		err := testDB.StoreAttestation(t.GetTestContext(), attestation0, origin+1, 1, 1)
		Nil(t.T(), err)
		err = testDB.StoreAttestation(t.GetTestContext(), attestation1, origin+1, 2, 2)
		Nil(t.T(), err)
		err = testDB.StoreAttestation(t.GetTestContext(), attestation2, origin+1, 3, 3)
		Nil(t.T(), err)

		earliestState, err := testDB.GetEarliestStateInRange(t.GetTestContext(), origin, origin+1, 0, 5)
		Nil(t.T(), err)
		Equal(t.T(), uint32(2), (*earliestState).Nonce())

		earliestState, err = testDB.GetEarliestStateInRange(t.GetTestContext(), origin, origin+1, 0, 1)
		Nil(t.T(), err)
		Nil(t.T(), earliestState)

		earliestState, err = testDB.GetEarliestStateInRange(t.GetTestContext(), origin, origin+1, 3, 5)
		Nil(t.T(), err)
		Equal(t.T(), uint32(4), (*earliestState).Nonce())

		earliestState, err = testDB.GetEarliestStateInRange(t.GetTestContext(), origin, origin+1, 6, 6)
		Nil(t.T(), err)
		Nil(t.T(), earliestState)

		earliestState, err = testDB.GetEarliestStateInRange(t.GetTestContext(), origin, origin+1, 5, 5)
		Nil(t.T(), err)
		Equal(t.T(), uint32(5), (*earliestState).Nonce())
	})
}
