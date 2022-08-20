package db_test

import (
	"math/big"
	"math/rand"
	"time"

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
		originDomain := gofakeit.Uint32()
		sender := common.BigToHash(big.NewInt(gofakeit.Int64()))
		nonce := gofakeit.Uint32()
		destinationDomain := gofakeit.Uint32()
		recipient := common.BigToHash(big.NewInt(gofakeit.Int64()))
		optimisticSeconds := gofakeit.Uint32()
		header := types.NewHeader(originDomain, sender, nonce, destinationDomain, recipient, optimisticSeconds)
		notaryTip := big.NewInt(gofakeit.Int64())
		broadcasterTip := big.NewInt(gofakeit.Int64())
		proverTip := big.NewInt(gofakeit.Int64())
		executorTip := big.NewInt(gofakeit.Int64())
		tips := types.NewTips(notaryTip, broadcasterTip, proverTip, executorTip)
		body := []byte(gofakeit.Sentence(10))

		message := types.NewMessage(header, tips, body)
		err := testDB.StoreDispatchMessage(t.GetTestContext(), message)
		Nil(t.T(), err)

		attestation := types.NewAttestation(1, gofakeit.Uint32(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))

		err = testDB.StoreAcceptedAttestation(t.GetTestContext(), gofakeit.Uint32(), attestation)
		Nil(t.T(), err)
	})
}

func (t *DBSuite) TestGetDelinquentMessage() {
	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		var nonceRange = uint32(gofakeit.Uint8())
		var destinationDomain = gofakeit.Uint32()
		var targetedDomain uint32
		var delinquentNonces []uint32
		var otherDelinquentNonces []uint32
		var header types.Header
		var message types.Message
		var attestation types.Attestation
		var storeAttestation bool

		for nonce := uint32(0); nonce <= nonceRange; nonce++ {
			// Populate the databases of DispatchMessages and AcceptedAttestations.
			// Use random cases for different scenarios of domains and if an attestation is stored.
			rand.Seed(time.Now().UnixNano())
			//nolint:gosec
			random := rand.Intn(4)
			switch random {
			// Case 0 is where we use destinationDomain and store the accepted attestation
			case 0:
				targetedDomain = destinationDomain
				storeAttestation = true
			// Case 1 is where we use destinationDomain and do not store the accepted attestation
			case 1:
				targetedDomain = destinationDomain
				storeAttestation = false
				// Keep track of what message nonces will be delinquent
				delinquentNonces = append(delinquentNonces, nonce)
			// Case 2 is where we use destinationDomain+1 and store the accepted attestation
			case 2:
				targetedDomain = destinationDomain + 1
				storeAttestation = true
			// Case 3 is where we use destinationDomain+1 and do not store the accepted attestation
			case 3:
				targetedDomain = destinationDomain + 1
				storeAttestation = false
				// Keep track of what message nonces will be delinquent
				otherDelinquentNonces = append(otherDelinquentNonces, nonce)
			}
			var err error
			header = types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), nonce, targetedDomain, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
			message = types.NewMessage(header, tips, []byte(gofakeit.Sentence(10)))
			err = testDB.StoreDispatchMessage(t.GetTestContext(), message)
			Nil(t.T(), err)
			if storeAttestation {
				attestation = types.NewAttestation(targetedDomain, nonce, common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
				err = testDB.StoreAcceptedAttestation(t.GetTestContext(), targetedDomain, attestation)
				Nil(t.T(), err)
			}
		}
		// Test to ensure the delinquent messages are successfully tracked.
		delinquentMessages, err := testDB.GetDelinquentMessages(t.GetTestContext(), destinationDomain)
		Nil(t.T(), err)
		Equal(t.T(), len(delinquentMessages), len(delinquentNonces))
		for index, delinquentMessage := range delinquentMessages {
			Equal(t.T(), delinquentMessage.Nonce(), delinquentNonces[index])
		}
		otherDelinquentMessages, err := testDB.GetDelinquentMessages(t.GetTestContext(), destinationDomain+1)
		Nil(t.T(), err)
		Equal(t.T(), len(otherDelinquentMessages), len(otherDelinquentNonces))
		for index, otherDelinquentMessage := range otherDelinquentMessages {
			Equal(t.T(), otherDelinquentMessage.Nonce(), otherDelinquentNonces[index])
		}
	})
}
