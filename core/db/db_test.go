package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/types"
)

func (t *DBSuite) TestRetrieveLatestNonce() {
	const domainID = 1

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		_, err := testDB.RetrieveLatestCommittedMessageNonce(t.GetTestContext(), domainID)
		ErrorIs(t.T(), err, db.ErrNoNonceForDomain)

		nonce := 0
		leafIndex := uint32(1)

		for i := 0; i < 10; i++ {
			header := types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), uint32(i), gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
			tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))

			realMessage := types.NewMessage(header, tips, []byte(gofakeit.Sentence(10)))

			encoded, err := types.EncodeMessage(realMessage)
			Nil(t.T(), err)

			err = testDB.StoreCommittedMessage(t.GetTestContext(), domainID, types.NewCommittedMessage(leafIndex, encoded))
			Nil(t.T(), err)

			newNonce, err := testDB.RetrieveLatestCommittedMessageNonce(t.GetTestContext(), domainID)
			Nil(t.T(), err)
			Equal(t.T(), uint32(nonce), newNonce)

			nonce++
			leafIndex++
		}
	})
}

func (t *DBSuite) TestStoreMonitoring() {
	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		var destinationDomain uint32 = 2
		var nonce uint32 = 1

		header := types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), nonce, destinationDomain, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		message := types.NewMessage(header, tips, []byte(gofakeit.Sentence(10)))
		// Store a message with nonce 1 and destinationDomain of 2
		err := testDB.StoreDispatchMessage(t.GetTestContext(), message)
		Nil(t.T(), err)

		nonce++

		header = types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), nonce, destinationDomain, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		message = types.NewMessage(header, tips, []byte(gofakeit.Sentence(10)))
		// Store a message with nonce 2 and destinationDomain of 2
		err = testDB.StoreDispatchMessage(t.GetTestContext(), message)
		Nil(t.T(), err)

		nonce++

		header = types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), nonce, destinationDomain, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		message = types.NewMessage(header, tips, []byte(gofakeit.Sentence(10)))
		// Store a message with nonce 3 and destinationDomain of 2
		err = testDB.StoreDispatchMessage(t.GetTestContext(), message)
		Nil(t.T(), err)

		nonce++

		header = types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), nonce, destinationDomain+1, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		message = types.NewMessage(header, tips, []byte(gofakeit.Sentence(10)))
		// Store a message with nonce 4 and destinationDomain of 3
		// Note: this one is a different destinationDomain of what we will query for
		// so this should not be returned in the query
		err = testDB.StoreDispatchMessage(t.GetTestContext(), message)
		Nil(t.T(), err)

		attestation := types.NewAttestation(destinationDomain, 2, common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))

		// Store an accepted attestation of nonce 2 and destinationDomain of 2
		err = testDB.StoreAcceptedAttestation(t.GetTestContext(), 2, attestation)
		Nil(t.T(), err)

		// Since we only stored an accepted attestation of nonce 2 on destinationDomain 2,
		// we should get 2 messages back:
		// (1) nonce: 1, destinationDomain: 2, (2) nonce: 3, destinationDomain: 2
		delinquentMessages, err := testDB.GetDelinquentMessages(t.GetTestContext(), 2)
		Nil(t.T(), err)
		Equal(t.T(), 2, len(delinquentMessages))
		Equal(t.T(), uint32(1), delinquentMessages[0].Nonce())
		Equal(t.T(), uint32(3), delinquentMessages[1].Nonce())
	})
}
