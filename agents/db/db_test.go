package db_test

import (
	"math/big"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

func (t *DBSuite) TestRetrieveLatestNonce() {
	const domainID = 1

	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		_, err := testDB.RetrieveLatestCommittedMessageNonce(t.GetTestContext(), domainID)
		ErrorIs(t.T(), err, db.ErrNoNonceForDomain)

		nonce := 0

		for i := 0; i < 10; i++ {
			header := types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), uint32(i), gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
			tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))

			realMessage := types.NewMessage(header, tips, []byte(gofakeit.Sentence(10)))

			encoded, err := types.EncodeMessage(realMessage)
			Nil(t.T(), err)

			err = testDB.StoreCommittedMessage(t.GetTestContext(), domainID, types.NewCommittedMessage(encoded))
			Nil(t.T(), err)

			newNonce, err := testDB.RetrieveLatestCommittedMessageNonce(t.GetTestContext(), domainID)
			Nil(t.T(), err)
			Equal(t.T(), uint32(nonce), newNonce)

			nonce++
		}
	})
}

func (t *DBSuite) TestStoreMonitoring() {
	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		header := types.NewHeader(gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32(), gofakeit.Uint32(), common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		message := types.NewMessage(header, tips, []byte(gofakeit.Sentence(10)))

		err := testDB.StoreDispatchMessage(t.GetTestContext(), message)
		Nil(t.T(), err)
		origin := uint32(1)
		destination := origin + 1
		nonce := gofakeit.Uint32()
		root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		attestKey := types.AttestationKey{
			Origin:      origin,
			Destination: destination,
			Nonce:       nonce,
		}
		attestation := types.NewAttestation(attestKey.GetRawKey(), root)

		err = testDB.StoreAcceptedAttestation(t.GetTestContext(), attestation)
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
			// TODO (joe): Look into this test more after we are further along with the refactor of
			// having the GlobalNotaryRegistry and having attestations from many origins
			if storeAttestation {
				attestKey := types.AttestationKey{
					Origin:      targetedDomain,
					Destination: targetedDomain,
					Nonce:       nonce,
				}

				attestation = types.NewAttestation(attestKey.GetRawKey(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
				err = testDB.StoreAcceptedAttestation(t.GetTestContext(), attestation)
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

func (t *DBSuite) TestNotaryHappyPath() {
	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		fakeOrigin := uint32(1)
		fakeDestination := fakeOrigin + 1

		fakeNonces := []uint32{}
		fakeRoots := []common.Hash{}
		fakeDispatchBlockNumbers := []uint64{}
		fakeSignatures := []types.Signature{}
		fakeSumbittedTimes := []time.Time{}
		fakeConfirmedBlockNumbers := []uint64{}

		fakeWallet, err := wallet.FromRandom()
		Nil(t.T(), err)

		fakeSigner := localsigner.NewSigner(fakeWallet.PrivateKey())

		numMessages := 4
		for i := 0; i <= numMessages; i++ {
			fakeNonce := gofakeit.Uint32()
			fakeRoot := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
			fakeDispatchBlockNumber := uint64(i)

			fakeNonces = append(fakeNonces, fakeNonce)
			fakeRoots = append(fakeRoots, fakeRoot)
			fakeDispatchBlockNumbers = append(fakeDispatchBlockNumbers, fakeDispatchBlockNumber)

			fakeAttestKey := types.AttestationKey{
				Origin:      fakeOrigin,
				Destination: fakeDestination,
				Nonce:       fakeNonce,
			}
			fakeUnsignedAttestation := types.NewAttestation(fakeAttestKey.GetRawKey(), fakeRoot)

			err := testDB.StoreNewInProgressAttestation(t.GetTestContext(), fakeUnsignedAttestation, fakeDispatchBlockNumber)
			Nil(t.T(), err)

			latestNonce, err := testDB.RetrieveLatestCachedNonce(t.GetTestContext(), fakeOrigin, fakeDestination)
			Nil(t.T(), err)
			Equal(t.T(), fakeNonce, latestNonce)
		}
		for i := 0; i <= numMessages; i++ {
			fakeNonce := fakeNonces[i]
			fakeRoot := fakeRoots[i]
			fakeDispatchBlockNumber := fakeDispatchBlockNumbers[i]

			inProgressAttestation, err := testDB.RetrieveOldestUnsignedInProgressAttestation(t.GetTestContext(), fakeOrigin, fakeDestination)
			Nil(t.T(), err)
			Equal(t.T(), fakeDispatchBlockNumber, inProgressAttestation.OriginDispatchBlockNumber())
			Equal(t.T(), uint64(0), inProgressAttestation.ConfirmedOnAttestationCollectorBlockNumber())
			Equal(t.T(), nil, inProgressAttestation.SubmittedToAttestationCollectorTime())
			Equal(t.T(), nil, inProgressAttestation.SignedAttestation().Signature())
			Equal(t.T(), fakeNonce, inProgressAttestation.SignedAttestation().Attestation().Nonce())
			Equal(t.T(), fakeRoot, inProgressAttestation.SignedAttestation().Attestation().Root())
			Equal(t.T(), fakeOrigin, inProgressAttestation.SignedAttestation().Attestation().Origin())
			Equal(t.T(), fakeDestination, inProgressAttestation.SignedAttestation().Attestation().Destination())

			hashedUpdate, err := inProgressAttestation.SignedAttestation().Attestation().Hash()
			Nil(t.T(), err)

			signature, err := fakeSigner.SignMessage(t.GetTestContext(), core.BytesToSlice(hashedUpdate), false)
			Nil(t.T(), err)
			fakeSignatures = append(fakeSignatures, signature)

			signedAttestation := types.NewSignedAttestation(inProgressAttestation.SignedAttestation().Attestation(), signature)
			signedInProgressAttestation := types.NewInProgressAttestation(signedAttestation, inProgressAttestation.OriginDispatchBlockNumber(), nil, 0)
			err = testDB.UpdateSignature(t.GetTestContext(), signedInProgressAttestation)
			Nil(t.T(), err)
		}
		inProgressAttestation, err := testDB.RetrieveOldestUnsignedInProgressAttestation(t.GetTestContext(), fakeOrigin, fakeDestination)
		Nil(t.T(), err)
		Nil(t.T(), inProgressAttestation)

		for i := 0; i <= numMessages; i++ {
			fakeNonce := fakeNonces[i]
			fakeRoot := fakeRoots[i]
			fakeDispatchBlockNumber := fakeDispatchBlockNumbers[i]
			fakeSignature := fakeSignatures[i]

			inProgressAttestation, err := testDB.RetrieveOldestUnsubmittedSignedInProgressAttestation(t.GetTestContext(), fakeOrigin, fakeDestination)
			Nil(t.T(), err)
			Equal(t.T(), fakeDispatchBlockNumber, inProgressAttestation.OriginDispatchBlockNumber())
			Equal(t.T(), uint64(0), inProgressAttestation.ConfirmedOnAttestationCollectorBlockNumber())
			Equal(t.T(), nil, inProgressAttestation.SubmittedToAttestationCollectorTime())
			Equal(t.T(), fakeSignature, inProgressAttestation.SignedAttestation().Signature())
			Equal(t.T(), fakeNonce, inProgressAttestation.SignedAttestation().Attestation().Nonce())
			Equal(t.T(), fakeRoot, inProgressAttestation.SignedAttestation().Attestation().Root())
			Equal(t.T(), fakeOrigin, inProgressAttestation.SignedAttestation().Attestation().Origin())
			Equal(t.T(), fakeDestination, inProgressAttestation.SignedAttestation().Attestation().Destination())

			nowTime := time.Now()
			fakeSumbittedTimes = append(fakeSumbittedTimes, nowTime)
			submittedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestation.SignedAttestation(), inProgressAttestation.OriginDispatchBlockNumber(), &nowTime, 0)
			err = testDB.UpdateSubmittedToAttestationCollectorTime(t.GetTestContext(), submittedInProgressAttestation)
			Nil(t.T(), err)
		}
		inProgressAttestation, err = testDB.RetrieveOldestUnsubmittedSignedInProgressAttestation(t.GetTestContext(), fakeOrigin, fakeDestination)
		Nil(t.T(), err)
		Nil(t.T(), inProgressAttestation)

		for i := 0; i <= numMessages; i++ {
			fakeNonce := fakeNonces[i]
			fakeRoot := fakeRoots[i]
			fakeDispatchBlockNumber := fakeDispatchBlockNumbers[i]
			fakeSignature := fakeSignatures[i]
			fakeSubmittedTime := fakeSumbittedTimes[i]

			inProgressAttestation, err := testDB.RetrieveOldestUnconfirmedSubmittedInProgressAttestation(t.GetTestContext(), fakeOrigin, fakeDestination)
			Nil(t.T(), err)
			Equal(t.T(), fakeDispatchBlockNumber, inProgressAttestation.OriginDispatchBlockNumber())
			Equal(t.T(), uint64(0), inProgressAttestation.ConfirmedOnAttestationCollectorBlockNumber())
			Equal(t.T(), fakeSubmittedTime, inProgressAttestation.SubmittedToAttestationCollectorTime())
			Equal(t.T(), fakeSignature, inProgressAttestation.SignedAttestation().Signature())
			Equal(t.T(), fakeNonce, inProgressAttestation.SignedAttestation().Attestation().Nonce())
			Equal(t.T(), fakeRoot, inProgressAttestation.SignedAttestation().Attestation().Root())
			Equal(t.T(), fakeOrigin, inProgressAttestation.SignedAttestation().Attestation().Origin())
			Equal(t.T(), fakeDestination, inProgressAttestation.SignedAttestation().Attestation().Destination())

			fakeConfirmedBlockNumer := uint64(numMessages) + uint64(i)
			fakeConfirmedBlockNumbers = append(fakeConfirmedBlockNumbers, fakeConfirmedBlockNumer)
			confirmedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestation.SignedAttestation(), inProgressAttestation.OriginDispatchBlockNumber(), inProgressAttestation.SubmittedToAttestationCollectorTime(), fakeConfirmedBlockNumer)
			err = testDB.UpdateConfirmedOnAttestationCollectorBlockNumber(t.GetTestContext(), confirmedInProgressAttestation)
			Nil(t.T(), err)
		}
		inProgressAttestation, err = testDB.RetrieveOldestUnconfirmedSubmittedInProgressAttestation(t.GetTestContext(), fakeOrigin, fakeDestination)
		Nil(t.T(), err)
		Nil(t.T(), inProgressAttestation)

		for i := 0; i <= numMessages; i++ {
			fakeNonce := fakeNonces[i]
			fakeRoot := fakeRoots[i]
			fakeDispatchBlockNumber := fakeDispatchBlockNumbers[i]
			fakeSignature := fakeSignatures[i]
			fakeSubmittedTime := fakeSumbittedTimes[i]
			fakeConfirmedBlockNumber := fakeConfirmedBlockNumbers[i]

			inProgressAttestation, err := testDB.RetrieveInProgressAttestation(t.GetTestContext(), fakeOrigin, fakeDestination, fakeNonce)
			Nil(t.T(), err)
			Equal(t.T(), fakeDispatchBlockNumber, inProgressAttestation.OriginDispatchBlockNumber())
			Equal(t.T(), fakeConfirmedBlockNumber, inProgressAttestation.ConfirmedOnAttestationCollectorBlockNumber())
			Equal(t.T(), fakeSubmittedTime, inProgressAttestation.SubmittedToAttestationCollectorTime())
			Equal(t.T(), fakeSignature, inProgressAttestation.SignedAttestation().Signature())
			Equal(t.T(), fakeNonce, inProgressAttestation.SignedAttestation().Attestation().Nonce())
			Equal(t.T(), fakeRoot, inProgressAttestation.SignedAttestation().Attestation().Root())
			Equal(t.T(), fakeOrigin, inProgressAttestation.SignedAttestation().Attestation().Origin())
			Equal(t.T(), fakeDestination, inProgressAttestation.SignedAttestation().Attestation().Destination())
		}
	})
}
