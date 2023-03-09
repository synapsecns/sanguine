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
		chainIDA := gofakeit.Uint32()
		destinationA := gofakeit.Uint32()
		nonceA := gofakeit.Uint32()
		blockNumberA := gofakeit.Uint64()
		blockTimeA := gofakeit.Uint64()
		attestKeyA := agentsTypes.AttestationKey{
			Origin:      chainIDA,
			Destination: destinationA,
			Nonce:       nonceA,
		}
		rootA := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		attestationA := agentsTypes.NewAttestation(attestKeyA.GetRawKey(), rootA)

		err := testDB.StoreAttestation(t.GetTestContext(), attestationA, blockNumberA, blockTimeA)
		Nil(t.T(), err)

		chainIDB := gofakeit.Uint32()
		destinationB := gofakeit.Uint32()
		nonceB := gofakeit.Uint32()
		blockNumberB := gofakeit.Uint64()
		blockTimeB := gofakeit.Uint64()
		attestKeyB := agentsTypes.AttestationKey{
			Origin:      chainIDB,
			Destination: destinationB,
			Nonce:       nonceB,
		}
		rootB := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		attestationB := agentsTypes.NewAttestation(attestKeyB.GetRawKey(), rootB)

		err = testDB.StoreAttestation(t.GetTestContext(), attestationB, blockNumberB, blockTimeB)
		Nil(t.T(), err)

		attestationMaskA := types.DBAttestation{
			ChainID:     &chainIDA,
			Destination: &destinationA,
		}

		retrievedAttestationA, err := testDB.GetAttestation(t.GetTestContext(), attestationMaskA)
		Nil(t.T(), err)

		encodeAttestationA, err := agentsTypes.EncodeAttestation(attestationA)
		Nil(t.T(), err)
		encodeRetrievedAttestationA, err := agentsTypes.EncodeAttestation(*retrievedAttestationA)
		Nil(t.T(), err)

		Equal(t.T(), encodeAttestationA, encodeRetrievedAttestationA)

		attestationMaskB := types.DBAttestation{
			DestinationBlockNumber: &blockNumberB,
			DestinationTimestamp:   &blockTimeB,
		}

		retrievedAttestationB, err := testDB.GetAttestation(t.GetTestContext(), attestationMaskB)
		Nil(t.T(), err)

		encodeAttestationB, err := agentsTypes.EncodeAttestation(attestationB)
		Nil(t.T(), err)
		encodeRetrievedAttestationB, err := agentsTypes.EncodeAttestation(*retrievedAttestationB)
		Nil(t.T(), err)

		Equal(t.T(), encodeAttestationB, encodeRetrievedAttestationB)
	})
}

func (t *DBSuite) TestAttestationBlockNumberBlockTime() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		destination := gofakeit.Uint32()
		nonce := gofakeit.Uint32()
		number := gofakeit.Uint64()
		time := gofakeit.Uint64()
		attestKey := agentsTypes.AttestationKey{
			Origin:      chainID,
			Destination: destination,
			Nonce:       nonce,
		}
		root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		attestation := agentsTypes.NewAttestation(attestKey.GetRawKey(), root)

		err := testDB.StoreAttestation(t.GetTestContext(), attestation, number, time)
		Nil(t.T(), err)

		wrongMask := types.DBAttestation{
			ChainID:     &destination,
			Destination: &chainID,
		}

		blockNumber, err := testDB.GetAttestationBlockNumber(t.GetTestContext(), wrongMask)
		Nil(t.T(), err)
		Nil(t.T(), blockNumber)

		blockTime, err := testDB.GetAttestationBlockTime(t.GetTestContext(), wrongMask)
		Nil(t.T(), err)
		Nil(t.T(), blockTime)

		mask := types.DBAttestation{
			ChainID:     &chainID,
			Destination: &destination,
			Nonce:       &nonce,
		}

		blockNumber, err = testDB.GetAttestationBlockNumber(t.GetTestContext(), mask)
		Nil(t.T(), err)
		Equal(t.T(), number, *blockNumber)

		blockTime, err = testDB.GetAttestationBlockTime(t.GetTestContext(), mask)
		Nil(t.T(), err)
		Equal(t.T(), time, *blockTime)
	})
}

func (t *DBSuite) TestGetAttestationsInNonceRange() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		for i := uint32(5); i <= 50; i += 5 {
			destination := gofakeit.Uint32()
			nonce := i
			number := gofakeit.Uint64()
			time := gofakeit.Uint64()
			attestKey := agentsTypes.AttestationKey{
				Origin:      chainID,
				Destination: destination,
				Nonce:       nonce,
			}
			root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
			attestation := agentsTypes.NewAttestation(attestKey.GetRawKey(), root)

			err := testDB.StoreAttestation(t.GetTestContext(), attestation, number, time)
			Nil(t.T(), err)
		}

		// Get attestations in nonce range
		mask := types.DBAttestation{
			ChainID: &chainID,
		}
		attestations, err := testDB.GetAttestationsAboveOrEqualNonce(t.GetTestContext(), mask, 6, 1)
		Nil(t.T(), err)
		Equal(t.T(), 9, len(attestations))

		attestations, err = testDB.GetAttestationsAboveOrEqualNonce(t.GetTestContext(), mask, 5, 1)
		Nil(t.T(), err)
		Equal(t.T(), 10, len(attestations))

		attestations, err = testDB.GetAttestationsAboveOrEqualNonce(t.GetTestContext(), mask, 100, 1)
		Nil(t.T(), err)
		Equal(t.T(), 0, len(attestations))
	})
}

func (t *DBSuite) TestGetEarliestAttestationsNonceInNonceRange() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		destination := gofakeit.Uint32()
		nonce := uint32(1)
		attestKey := agentsTypes.AttestationKey{
			Origin:      chainID,
			Destination: destination,
			Nonce:       nonce,
		}
		root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		attestation := agentsTypes.NewAttestation(attestKey.GetRawKey(), root)

		err := testDB.StoreAttestation(t.GetTestContext(), attestation, 2, 2)
		Nil(t.T(), err)

		nonce = uint32(5)
		attestKey = agentsTypes.AttestationKey{
			Origin:      chainID,
			Destination: destination,
			Nonce:       nonce,
		}
		root = common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		attestation = agentsTypes.NewAttestation(attestKey.GetRawKey(), root)

		err = testDB.StoreAttestation(t.GetTestContext(), attestation, 1, 1)
		Nil(t.T(), err)

		nonce = uint32(10)
		attestKey = agentsTypes.AttestationKey{
			Origin:      chainID,
			Destination: destination,
			Nonce:       nonce,
		}
		root = common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		attestation = agentsTypes.NewAttestation(attestKey.GetRawKey(), root)

		err = testDB.StoreAttestation(t.GetTestContext(), attestation, 3, 3)
		Nil(t.T(), err)
		// Get attestations in nonce range
		mask := types.DBAttestation{
			ChainID: &chainID,
		}

		nonceNum, err := testDB.GetEarliestAttestationsNonceInNonceRange(t.GetTestContext(), mask, 1, 10)
		Nil(t.T(), err)

		Equal(t.T(), uint32(5), *nonceNum)
	})
}
