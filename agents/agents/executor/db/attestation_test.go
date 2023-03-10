package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

func (t *DBSuite) TestStoreRetrieveAttestation() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		heightA := gofakeit.Uint8()
		nonceA := gofakeit.Uint32()
		blockNumberA := big.NewInt(int64(gofakeit.Uint32()))
		timestampA := big.NewInt(int64(gofakeit.Uint32()))
		attestationA := agentsTypes.NewAttestation(snapshotRootA, heightA, nonceA, blockNumberA, timestampA)
		destinationA := gofakeit.Uint32()
		destinationBlockNumberA := gofakeit.Uint64()
		destinationTimestampA := gofakeit.Uint64()

		err := testDB.StoreAttestation(t.GetTestContext(), attestationA, destinationA, destinationBlockNumberA, destinationTimestampA)
		Nil(t.T(), err)

		snapshotRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		heightB := gofakeit.Uint8()
		nonceB := gofakeit.Uint32()
		blockNumberB := big.NewInt(int64(gofakeit.Uint32()))
		timestampB := big.NewInt(int64(gofakeit.Uint32()))
		attestationB := agentsTypes.NewAttestation(snapshotRootB, heightB, nonceB, blockNumberB, timestampB)
		destinationB := gofakeit.Uint32()
		destinationBlockNumberB := gofakeit.Uint64()
		destinationTimestampB := gofakeit.Uint64()

		err = testDB.StoreAttestation(t.GetTestContext(), attestationB, destinationB, destinationBlockNumberB, destinationTimestampB)
		Nil(t.T(), err)

		snapshotRootAString := snapshotRootA.String()
		attestationMaskA := types.DBAttestation{
			Destination:      &destinationA,
			SnapshotRoot:     &snapshotRootAString,
			Height:           &heightA,
			AttestationNonce: &nonceA,
		}

		retrievedAttestationA, err := testDB.GetAttestation(t.GetTestContext(), attestationMaskA)
		Nil(t.T(), err)
		Equal(t.T(), attestationA, *retrievedAttestationA)

		blockNumberBUint64 := blockNumberB.Uint64()
		timestampBUint64 := timestampB.Uint64()
		attestationMaskB := types.DBAttestation{
			SummitBlockNumber:      &blockNumberBUint64,
			SummitTimestamp:        &timestampBUint64,
			DestinationBlockNumber: &destinationBlockNumberB,
			DestinationTimestamp:   &destinationTimestampB,
		}

		retrievedAttestationB, err := testDB.GetAttestation(t.GetTestContext(), attestationMaskB)
		Nil(t.T(), err)

		Equal(t.T(), attestationB, *retrievedAttestationB)
	})
}

func (t *DBSuite) TestGetAttestationMinimumTimestamp() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		destination := gofakeit.Uint32()

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		heightA := gofakeit.Uint8()
		nonceA := gofakeit.Uint32()
		blockNumberA := big.NewInt(int64(gofakeit.Uint32()))
		timestampA := big.NewInt(int64(gofakeit.Uint32()))
		attestationA := agentsTypes.NewAttestation(snapshotRootA, heightA, nonceA, blockNumberA, timestampA)

		destinationBlockNumberA := gofakeit.Uint64()
		destinationTimestampA := gofakeit.Uint64()

		err := testDB.StoreAttestation(t.GetTestContext(), attestationA, destination, destinationBlockNumberA, destinationTimestampA)
		Nil(t.T(), err)

		snapshotRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		heightB := gofakeit.Uint8()
		nonceB := gofakeit.Uint32()
		blockNumberB := big.NewInt(int64(gofakeit.Uint32()))
		timestampB := big.NewInt(int64(gofakeit.Uint32()))
		attestationB := agentsTypes.NewAttestation(snapshotRootB, heightB, nonceB, blockNumberB, timestampB)

		destinationBlockNumberB := destinationBlockNumberA + 1
		destinationTimestampB := destinationTimestampA + 1

		err = testDB.StoreAttestation(t.GetTestContext(), attestationB, destination, destinationBlockNumberB, destinationTimestampB)
		Nil(t.T(), err)

		snapshotRootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		heightC := gofakeit.Uint8()
		nonceC := gofakeit.Uint32()
		blockNumberC := big.NewInt(int64(gofakeit.Uint32()))
		timestampC := big.NewInt(int64(gofakeit.Uint32()))
		attestationC := agentsTypes.NewAttestation(snapshotRootC, heightC, nonceC, blockNumberC, timestampC)

		destinationBlockNumberC := uint64(0)
		destinationTimestampC := uint64(0)

		err = testDB.StoreAttestation(t.GetTestContext(), attestationC, destination, destinationBlockNumberC, destinationTimestampC)
		Nil(t.T(), err)

		mask := types.DBAttestation{
			Destination: &destination,
		}

		roots := []string{snapshotRootA.String(), snapshotRootB.String()}

		minimumTimestamp, err := testDB.GetAttestationMinimumTimestamp(t.GetTestContext(), mask, roots)
		Nil(t.T(), err)

		Equal(t.T(), destinationTimestampA, *minimumTimestamp)
	})
}

func (t *DBSuite) TestGetEarliestSnapshotFromAttestation() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		destination := gofakeit.Uint32()

		snapshotRootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		heightA := gofakeit.Uint8()
		nonceA := gofakeit.Uint32()
		blockNumberA := big.NewInt(int64(gofakeit.Uint32()))
		timestampA := big.NewInt(int64(gofakeit.Uint32()))
		attestationA := agentsTypes.NewAttestation(snapshotRootA, heightA, nonceA, blockNumberA, timestampA)

		destinationBlockNumberA := gofakeit.Uint64()
		destinationTimestampA := gofakeit.Uint64()

		err := testDB.StoreAttestation(t.GetTestContext(), attestationA, destination, destinationBlockNumberA, destinationTimestampA)
		Nil(t.T(), err)

		snapshotRootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		heightB := gofakeit.Uint8()
		nonceB := gofakeit.Uint32()
		blockNumberB := big.NewInt(int64(gofakeit.Uint32()))
		timestampB := big.NewInt(int64(gofakeit.Uint32()))
		attestationB := agentsTypes.NewAttestation(snapshotRootB, heightB, nonceB, blockNumberB, timestampB)

		destinationBlockNumberB := destinationBlockNumberA + 1
		destinationTimestampB := destinationTimestampA + 1

		err = testDB.StoreAttestation(t.GetTestContext(), attestationB, destination, destinationBlockNumberB, destinationTimestampB)
		Nil(t.T(), err)

		snapshotRootC := common.BigToHash(big.NewInt(gofakeit.Int64()))
		heightC := gofakeit.Uint8()
		nonceC := gofakeit.Uint32()
		blockNumberC := big.NewInt(int64(gofakeit.Uint32()))
		timestampC := big.NewInt(int64(gofakeit.Uint32()))
		attestationC := agentsTypes.NewAttestation(snapshotRootC, heightC, nonceC, blockNumberC, timestampC)

		destinationBlockNumberC := uint64(0)
		destinationTimestampC := uint64(0)

		err = testDB.StoreAttestation(t.GetTestContext(), attestationC, destination, destinationBlockNumberC, destinationTimestampC)
		Nil(t.T(), err)

		mask := types.DBAttestation{
			Destination: &destination,
		}

		roots := []string{snapshotRootA.String(), snapshotRootB.String()}

		earliestSnapshot, err := testDB.GetEarliestSnapshotFromAttestation(t.GetTestContext(), mask, roots)
		Nil(t.T(), err)

		Equal(t.T(), snapshotRootA, common.BytesToHash((*earliestSnapshot)[:]))
	})
}
