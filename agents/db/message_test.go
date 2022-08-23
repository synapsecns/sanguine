package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

func (t *DBSuite) TestStoreRetreiveMessageLatestBlockEnd() {
	const testDomain = 10

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		height, err := testDB.GetMessageLatestBlockEnd(t.GetTestContext(), testDomain)
		ErrorIs(t.T(), err, db.ErrNoStoredBlockForChain, "expected an error when no height is stored")
		Zerof(t.T(), height, "expected non-existent height")

		testHeight := uint32(gofakeit.Uint16())

		// store again
		err = testDB.StoreMessageLatestBlockEnd(t.GetTestContext(), testDomain, testHeight)
		Nil(t.T(), err)

		// store a different height on another chain to see if we break anything
		err = testDB.StoreMessageLatestBlockEnd(t.GetTestContext(), uint32(testDomain+1+gofakeit.Uint16()), testHeight)
		Nil(t.T(), err)

		height, err = testDB.GetMessageLatestBlockEnd(t.GetTestContext(), testDomain)
		Nil(t.T(), err)
		Equal(t.T(), height, testHeight)
	})
}

func (t *DBSuite) TestStoreRetrieveSignedAttestaion() {
	domain := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	fakeV := new(big.Int).SetUint64(uint64(gofakeit.Uint8()))
	fakeR := big.NewInt(gofakeit.Int64())
	fakeS := big.NewInt(gofakeit.Int64())

	signedAttestation := types.NewSignedAttestation(
		types.NewAttestation(domain, nonce, root),
		types.NewSignature(fakeV, fakeR, fakeS),
	)

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		err := testDB.StoreSignedAttestations(t.GetTestContext(), signedAttestation)
		Nil(t.T(), err)

		retrievedAttestation, err := testDB.RetrieveSignedAttestationByNonce(t.GetTestContext(), domain, nonce)
		Nil(t.T(), err)

		Equal(t.T(), signedAttestation.Attestation().Root(), retrievedAttestation.Attestation().Root())
		Equal(t.T(), signedAttestation.Attestation().Domain(), retrievedAttestation.Attestation().Domain())
		Equal(t.T(), signedAttestation.Attestation().Nonce(), retrievedAttestation.Attestation().Nonce())

		Equal(t.T(), signedAttestation.Signature().V(), retrievedAttestation.Signature().V())
		Equal(t.T(), signedAttestation.Signature().R(), retrievedAttestation.Signature().R())
		Equal(t.T(), signedAttestation.Signature().S(), retrievedAttestation.Signature().S())
	})
}
