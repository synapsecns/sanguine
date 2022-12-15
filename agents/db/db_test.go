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

func (t *DBSuite) launchTestStoreNewInProgressAttestations(testDB db.SynapseDB) {
	t.fakeOrigin = uint32(1)
	t.fakeDestination = t.fakeOrigin + 1

	t.fakeNonces = []uint32{}
	t.fakeRoots = []*common.Hash{}
	t.fakeDispatchBlockNumbers = []uint64{}

	fakeWallet, err := wallet.FromRandom()
	Nil(t.T(), err)

	t.fakeSigner = localsigner.NewSigner(fakeWallet.PrivateKey())

	t.numMessages = 4
	for i := 0; i <= t.numMessages; i++ {
		fakeNonce := uint32(i) + 1
		fakeRoot := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		fakeDispatchBlockNumber := uint64(i) + 1

		t.fakeNonces = append(t.fakeNonces, fakeNonce)
		t.fakeRoots = append(t.fakeRoots, &fakeRoot)
		t.fakeDispatchBlockNumbers = append(t.fakeDispatchBlockNumbers, fakeDispatchBlockNumber)

		fakeAttestKey := types.AttestationKey{
			Origin:      t.fakeOrigin,
			Destination: t.fakeDestination,
			Nonce:       fakeNonce,
		}
		fakeUnsignedAttestation := types.NewAttestation(fakeAttestKey.GetRawKey(), fakeRoot)

		err := testDB.StoreNewInProgressAttestation(t.GetTestContext(), fakeUnsignedAttestation, fakeDispatchBlockNumber)
		Nil(t.T(), err)

		latestNonce, err := testDB.RetrieveLatestCachedNonce(t.GetTestContext(), t.fakeOrigin, t.fakeDestination)
		Nil(t.T(), err)
		Equal(t.T(), fakeNonce, latestNonce)

		retrievedAttestation, err := testDB.RetrieveInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination, fakeNonce)
		Nil(t.T(), err)
		NotNil(t.T(), retrievedAttestation)
		Equal(t.T(), fakeNonce, retrievedAttestation.SignedAttestation().Attestation().Nonce())
		Equal(t.T(), [32]byte(fakeRoot), retrievedAttestation.SignedAttestation().Attestation().Root())
		Equal(t.T(), t.fakeOrigin, retrievedAttestation.SignedAttestation().Attestation().Origin())
		Equal(t.T(), fakeDispatchBlockNumber, retrievedAttestation.OriginDispatchBlockNumber())
		Nil(t.T(), retrievedAttestation.SignedAttestation().NotarySignatures())
		Nil(t.T(), retrievedAttestation.SubmittedToAttestationCollectorTime())
		Equal(t.T(), types.AttestationStateNotaryUnsigned, retrievedAttestation.AttestationState())
	}
}

func (t *DBSuite) launchTestUpdateSignatures(testDB db.SynapseDB) {
	t.fakeSignatures = []types.Signature{}
	for i := 0; i <= t.numMessages; i++ {
		fakeNonce := t.fakeNonces[i]
		fakeRoot := t.fakeRoots[i]
		fakeDispatchBlockNumber := t.fakeDispatchBlockNumbers[i]

		inProgressAttestation, err := testDB.RetrieveOldestUnsignedInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination)
		Nil(t.T(), err)
		NotNil(t.T(), inProgressAttestation)
		Equal(t.T(), fakeDispatchBlockNumber, inProgressAttestation.OriginDispatchBlockNumber())
		Nil(t.T(), inProgressAttestation.SubmittedToAttestationCollectorTime())
		Nil(t.T(), inProgressAttestation.SignedAttestation().NotarySignatures())
		Equal(t.T(), fakeNonce, inProgressAttestation.SignedAttestation().Attestation().Nonce())
		Equal(t.T(), [32]byte(*fakeRoot), inProgressAttestation.SignedAttestation().Attestation().Root())
		Equal(t.T(), t.fakeOrigin, inProgressAttestation.SignedAttestation().Attestation().Origin())
		Equal(t.T(), t.fakeDestination, inProgressAttestation.SignedAttestation().Attestation().Destination())
		Equal(t.T(), types.AttestationStateNotaryUnsigned, inProgressAttestation.AttestationState())

		hashedUpdate, err := types.Hash(inProgressAttestation.SignedAttestation().Attestation())
		Nil(t.T(), err)

		signature, err := t.fakeSigner.SignMessage(t.GetTestContext(), core.BytesToSlice(hashedUpdate), false)
		Nil(t.T(), err)
		t.fakeSignatures = append(t.fakeSignatures, signature)

		signedAttestation := types.NewSignedAttestation(inProgressAttestation.SignedAttestation().Attestation(), []types.Signature{}, []types.Signature{signature})
		signedInProgressAttestation := types.NewInProgressAttestation(signedAttestation, inProgressAttestation.OriginDispatchBlockNumber(), nil, 0)
		err = testDB.UpdateSignature(t.GetTestContext(), signedInProgressAttestation)
		Nil(t.T(), err)

		retrievedAttestation, err := testDB.RetrieveInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination, fakeNonce)
		Nil(t.T(), err)
		NotNil(t.T(), retrievedAttestation)
		Equal(t.T(), fakeNonce, retrievedAttestation.SignedAttestation().Attestation().Nonce())
		Equal(t.T(), [32]byte(*fakeRoot), retrievedAttestation.SignedAttestation().Attestation().Root())
		Equal(t.T(), t.fakeOrigin, retrievedAttestation.SignedAttestation().Attestation().Origin())
		Equal(t.T(), fakeDispatchBlockNumber, retrievedAttestation.OriginDispatchBlockNumber())
		inProgressSigBytes, err := types.EncodeSignature(signedInProgressAttestation.SignedAttestation().NotarySignatures()[0])
		Nil(t.T(), err)
		retrievedSigBytes, err := types.EncodeSignature(retrievedAttestation.SignedAttestation().NotarySignatures()[0])
		Nil(t.T(), err)
		Equal(t.T(), inProgressSigBytes, retrievedSigBytes)
		Nil(t.T(), retrievedAttestation.SubmittedToAttestationCollectorTime())
		Equal(t.T(), types.AttestationStateNotarySignedUnsubmitted, retrievedAttestation.AttestationState())
	}
	inProgressAttestation, err := testDB.RetrieveOldestUnsignedInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination)
	NotNil(t.T(), err)
	Nil(t.T(), inProgressAttestation)
}

func (t *DBSuite) launchTestSubmittedToAttestationCollectorTimes(testDB db.SynapseDB) {
	t.fakeSumbittedTimes = []time.Time{}
	for i := 0; i <= t.numMessages; i++ {
		fakeNonce := t.fakeNonces[i]
		fakeRoot := t.fakeRoots[i]
		fakeDispatchBlockNumber := t.fakeDispatchBlockNumbers[i]
		fakeSignature := t.fakeSignatures[i]

		inProgressAttestation, err := testDB.RetrieveOldestUnsubmittedSignedInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination)
		Nil(t.T(), err)
		NotNil(t.T(), inProgressAttestation)
		Equal(t.T(), fakeDispatchBlockNumber, inProgressAttestation.OriginDispatchBlockNumber())
		Nil(t.T(), inProgressAttestation.SubmittedToAttestationCollectorTime())
		inProgressSigBytes, err := types.EncodeSignature(inProgressAttestation.SignedAttestation().NotarySignatures()[0])
		Nil(t.T(), err)
		fakeSigBytes, err := types.EncodeSignature(fakeSignature)
		Nil(t.T(), err)
		Equal(t.T(), fakeSigBytes, inProgressSigBytes)
		Equal(t.T(), fakeNonce, inProgressAttestation.SignedAttestation().Attestation().Nonce())
		Equal(t.T(), [32]byte(*fakeRoot), inProgressAttestation.SignedAttestation().Attestation().Root())
		Equal(t.T(), t.fakeOrigin, inProgressAttestation.SignedAttestation().Attestation().Origin())
		Equal(t.T(), t.fakeDestination, inProgressAttestation.SignedAttestation().Attestation().Destination())
		Equal(t.T(), types.AttestationStateNotarySignedUnsubmitted, inProgressAttestation.AttestationState())

		nowTime := time.Now()
		t.fakeSumbittedTimes = append(t.fakeSumbittedTimes, nowTime)
		submittedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestation.SignedAttestation(), inProgressAttestation.OriginDispatchBlockNumber(), &nowTime, 0)
		err = testDB.UpdateSubmittedToAttestationCollectorTime(t.GetTestContext(), submittedInProgressAttestation)
		Nil(t.T(), err)

		retrievedAttestation, err := testDB.RetrieveInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination, fakeNonce)
		Nil(t.T(), err)
		NotNil(t.T(), retrievedAttestation)
		Equal(t.T(), fakeNonce, retrievedAttestation.SignedAttestation().Attestation().Nonce())
		Equal(t.T(), [32]byte(*fakeRoot), retrievedAttestation.SignedAttestation().Attestation().Root())
		Equal(t.T(), t.fakeOrigin, retrievedAttestation.SignedAttestation().Attestation().Origin())
		Equal(t.T(), fakeDispatchBlockNumber, retrievedAttestation.OriginDispatchBlockNumber())
		retrievedSigBytes, err := types.EncodeSignature(retrievedAttestation.SignedAttestation().NotarySignatures()[0])
		Nil(t.T(), err)
		Equal(t.T(), fakeSigBytes, retrievedSigBytes)
		NotNil(t.T(), retrievedAttestation.SubmittedToAttestationCollectorTime())
		Equal(t.T(), nowTime.Unix(), retrievedAttestation.SubmittedToAttestationCollectorTime().Unix())
		Equal(t.T(), types.AttestationStateNotarySubmittedUnconfirmed, retrievedAttestation.AttestationState())
	}
	inProgressAttestation, err := testDB.RetrieveOldestUnsubmittedSignedInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination)
	NotNil(t.T(), err)
	Nil(t.T(), inProgressAttestation)
}

func (t *DBSuite) launchTestMarkConfirmedOnAttestationCollector(testDB db.SynapseDB) {
	for i := 0; i <= t.numMessages; i++ {
		fakeNonce := t.fakeNonces[i]
		fakeRoot := t.fakeRoots[i]
		fakeDispatchBlockNumber := t.fakeDispatchBlockNumbers[i]
		fakeSignature := t.fakeSignatures[i]
		fakeSubmittedTime := t.fakeSumbittedTimes[i]

		inProgressAttestation, err := testDB.RetrieveOldestUnconfirmedSubmittedInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination)
		Nil(t.T(), err)
		NotNil(t.T(), inProgressAttestation)
		Equal(t.T(), fakeDispatchBlockNumber, inProgressAttestation.OriginDispatchBlockNumber())
		NotNil(t.T(), inProgressAttestation.SubmittedToAttestationCollectorTime())
		Equal(t.T(), fakeSubmittedTime.Unix(), inProgressAttestation.SubmittedToAttestationCollectorTime().Unix())
		inProgressSigBytes, err := types.EncodeSignature(inProgressAttestation.SignedAttestation().NotarySignatures()[0])
		Nil(t.T(), err)
		fakeSigBytes, err := types.EncodeSignature(fakeSignature)
		Nil(t.T(), err)
		Equal(t.T(), fakeSigBytes, inProgressSigBytes)
		Equal(t.T(), fakeNonce, inProgressAttestation.SignedAttestation().Attestation().Nonce())
		Equal(t.T(), [32]byte(*fakeRoot), inProgressAttestation.SignedAttestation().Attestation().Root())
		Equal(t.T(), t.fakeOrigin, inProgressAttestation.SignedAttestation().Attestation().Origin())
		Equal(t.T(), t.fakeDestination, inProgressAttestation.SignedAttestation().Attestation().Destination())
		Equal(t.T(), types.AttestationStateNotarySubmittedUnconfirmed, inProgressAttestation.AttestationState())

		confirmedInProgressAttestation := types.NewInProgressAttestation(inProgressAttestation.SignedAttestation(), inProgressAttestation.OriginDispatchBlockNumber(), inProgressAttestation.SubmittedToAttestationCollectorTime(), 0)
		err = testDB.MarkConfirmedOnAttestationCollector(t.GetTestContext(), confirmedInProgressAttestation)
		Nil(t.T(), err)

		retrievedConfirmedInProgressAttestation, err := testDB.RetrieveNewestConfirmedInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination)
		Nil(t.T(), err)
		NotNil(t.T(), retrievedConfirmedInProgressAttestation)
		Equal(t.T(), fakeDispatchBlockNumber, retrievedConfirmedInProgressAttestation.OriginDispatchBlockNumber())
		Equal(t.T(), fakeSubmittedTime.Unix(), retrievedConfirmedInProgressAttestation.SubmittedToAttestationCollectorTime().Unix())
		confirmedInProgressSigBytes, err := types.EncodeSignature(retrievedConfirmedInProgressAttestation.SignedAttestation().NotarySignatures()[0])
		Nil(t.T(), err)
		confirmedFakeSigBytes, err := types.EncodeSignature(fakeSignature)
		Nil(t.T(), err)
		Equal(t.T(), confirmedFakeSigBytes, confirmedInProgressSigBytes)
		Equal(t.T(), fakeNonce, retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Nonce())
		Equal(t.T(), [32]byte(*fakeRoot), retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Root())
		Equal(t.T(), t.fakeOrigin, retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Origin())
		Equal(t.T(), t.fakeDestination, retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Destination())
		Equal(t.T(), types.AttestationStateNotaryConfirmed, retrievedConfirmedInProgressAttestation.AttestationState())
	}
	inProgressAttestation, err := testDB.RetrieveOldestUnconfirmedSubmittedInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination)
	NotNil(t.T(), err)
	Nil(t.T(), inProgressAttestation)
}

func (t *DBSuite) launchTestVerifyAllAreConfirmed(testDB db.SynapseDB) {
	for i := 0; i <= t.numMessages; i++ {
		fakeNonce := t.fakeNonces[i]
		fakeRoot := t.fakeRoots[i]
		fakeDispatchBlockNumber := t.fakeDispatchBlockNumbers[i]
		fakeSignature := t.fakeSignatures[i]
		fakeSubmittedTime := t.fakeSumbittedTimes[i]

		inProgressAttestation, err := testDB.RetrieveInProgressAttestation(t.GetTestContext(), t.fakeOrigin, t.fakeDestination, fakeNonce)
		Nil(t.T(), err)
		NotNil(t.T(), inProgressAttestation)
		Equal(t.T(), fakeDispatchBlockNumber, inProgressAttestation.OriginDispatchBlockNumber())
		Equal(t.T(), fakeSubmittedTime.Unix(), inProgressAttestation.SubmittedToAttestationCollectorTime().Unix())
		inProgressSigBytes, err := types.EncodeSignature(inProgressAttestation.SignedAttestation().NotarySignatures()[0])
		Nil(t.T(), err)
		fakeSigBytes, err := types.EncodeSignature(fakeSignature)
		Nil(t.T(), err)
		Equal(t.T(), fakeSigBytes, inProgressSigBytes)
		Equal(t.T(), fakeNonce, inProgressAttestation.SignedAttestation().Attestation().Nonce())
		Equal(t.T(), [32]byte(*fakeRoot), inProgressAttestation.SignedAttestation().Attestation().Root())
		Equal(t.T(), t.fakeOrigin, inProgressAttestation.SignedAttestation().Attestation().Origin())
		Equal(t.T(), t.fakeDestination, inProgressAttestation.SignedAttestation().Attestation().Destination())
		Equal(t.T(), types.AttestationStateNotaryConfirmed, inProgressAttestation.AttestationState())
	}
}

func (t *DBSuite) TestNotaryHappyPath() {
	t.RunOnAllDBs(func(testDB db.SynapseDB) {
		t.launchTestStoreNewInProgressAttestations(testDB)
		t.launchTestUpdateSignatures(testDB)
		t.launchTestSubmittedToAttestationCollectorTimes(testDB)
		t.launchTestMarkConfirmedOnAttestationCollector(testDB)
		t.launchTestVerifyAllAreConfirmed(testDB)
	})
}
