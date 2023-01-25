package guard_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
)

func (u GuardSuite) TestAttestationGuardDestinationVerifier() {
	destination := uint32(u.TestBackendDestination.GetChainID())
	origin := uint32(u.TestBackendOrigin.GetChainID())
	nonce := uint32(1)

	testDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)

	// dispatch a random update
	originAuth := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	optimisticSeconds := uint32(5)
	tx, err := u.OriginContract.Dispatch(originAuth.TransactOpts, destination, [32]byte{}, optimisticSeconds, encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

	root, dispatchBlockNumber, err := u.OriginContract.GetHistoricalRoot(&bind.CallOpts{Context: u.GetTestContext()}, destination, nonce)
	Nil(u.T(), err)
	Greater(u.T(), dispatchBlockNumber.Uint64(), uint64(0))

	attestKey := types.AttestationKey{
		Origin:      origin,
		Destination: destination,
		Nonce:       nonce,
	}
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	Nil(u.T(), err)

	notarySignature, err := u.NotaryBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(u.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(u.T(), err)

	attestationAuth := u.TestBackendAttestation.GetTxContext(u.GetTestContext(), nil)
	tx, err = u.AttestationContract.SubmitAttestation(attestationAuth.TransactOpts, rawSignedAttestation)
	Nil(u.T(), err)

	u.TestBackendAttestation.WaitForConfirmation(u.GetTestContext(), tx)

	latestNonce, err := u.AttestationDomainClient.AttestationCollector().GetLatestNonce(u.GetTestContext(), u.OriginDomainClient.Config().DomainID, u.DestinationDomainClient.Config().DomainID, u.NotaryBondedSigner)
	Nil(u.T(), err)
	Equal(u.T(), nonce, latestNonce)

	rawSignedAttestationFromCollector, err := u.AttestationContract.GetAttestation(&bind.CallOpts{Context: u.GetTestContext()}, origin, destination, nonce)
	Nil(u.T(), err)

	signedAttestationFromCollector, err := types.DecodeSignedAttestation(rawSignedAttestationFromCollector)
	Nil(u.T(), err)

	err = testDB.StoreExistingSignedInProgressAttestation(u.GetTestContext(), signedAttestationFromCollector)
	Nil(u.T(), err)

	inProgressAttestationToMarkVerified, err := testDB.RetrieveNewestInProgressAttestationIfInState(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID,
		types.AttestationStateGuardUnsignedAndUnverified)
	Nil(u.T(), err)

	nowTime := time.Now()
	submittedInProgressAttestation := types.NewInProgressAttestation(
		signedAttestationFromCollector,
		&nowTime,
		0)
	err = testDB.MarkVerifiedOnOrigin(u.GetTestContext(), submittedInProgressAttestation)
	Nil(u.T(), err)

	guardSignature, err := u.GuardBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(u.T(), err)

	guardSignedAttestation := types.NewSignedAttestation(
		unsignedAttestation,
		[]types.Signature{guardSignature},
		[]types.Signature{notarySignature})
	signedInProgressAttestation := types.NewInProgressAttestation(
		guardSignedAttestation,
		nil,
		0)
	err = testDB.UpdateGuardSignature(u.GetTestContext(), signedInProgressAttestation)
	Nil(u.T(), err)

	guardOnlySignedAttestation := types.NewSignedAttestation(
		guardSignedAttestation.Attestation(),
		guardSignedAttestation.GuardSignatures(),
		[]types.Signature{})

	rawSignedGuardOnlyAttestation, err := types.EncodeSignedAttestation(guardOnlySignedAttestation)
	Nil(u.T(), err)

	tx, err = u.AttestationContract.SubmitAttestation(attestationAuth.TransactOpts, rawSignedGuardOnlyAttestation)
	Nil(u.T(), err)

	u.TestBackendAttestation.WaitForConfirmation(u.GetTestContext(), tx)

	inProgressAttestationToSubmit := types.NewInProgressAttestation(
		inProgressAttestationToMarkVerified.SignedAttestation(),
		&nowTime,
		0)

	err = testDB.UpdateGuardSubmittedToAttestationCollectorTime(u.GetTestContext(), inProgressAttestationToSubmit)
	Nil(u.T(), err)

	err = testDB.MarkGuardConfirmedOnAttestationCollector(u.GetTestContext(), inProgressAttestationToSubmit)
	Nil(u.T(), err)

	destinationAuth := u.TestBackendDestination.GetTxContext(u.GetTestContext(), nil)

	rawSignedNotaryAndGuardAttestation, err := types.EncodeSignedAttestation(guardSignedAttestation)
	Nil(u.T(), err)
	destSubmitTx, err := u.DestinationContract.SubmitAttestation(destinationAuth.TransactOpts, rawSignedNotaryAndGuardAttestation)
	Nil(u.T(), err)
	u.TestBackendDestination.WaitForConfirmation(u.GetTestContext(), destSubmitTx)

	err = testDB.UpdateSubmittedToDestinationTime(u.GetTestContext(), submittedInProgressAttestation)
	Nil(u.T(), err)

	// Now call the guard destination verifier
	attestationGuardDestinationVerifer := guard.NewAttestationGuardDestinationVerifier(
		u.OriginDomainClient,
		u.AttestationDomainClient,
		u.DestinationDomainClient,
		testDB,
		u.GuardBondedSigner,
		u.GuardUnbondedSigner,
		1*time.Second)

	err = attestationGuardDestinationVerifer.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure the attesation has been verifier
	retrievedNewestConfirmedOnDestination, err := testDB.RetrieveNewestInProgressAttestationIfInState(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID,
		types.AttestationStateConfirmedOnDestination)

	Nil(u.T(), err)
	NotNil(u.T(), retrievedNewestConfirmedOnDestination)

	retrievedAttestation := retrievedNewestConfirmedOnDestination.SignedAttestation()
	Equal(u.T(), u.OriginDomainClient.Config().DomainID, retrievedAttestation.Attestation().Origin())
	Equal(u.T(), u.DestinationDomainClient.Config().DomainID, retrievedAttestation.Attestation().Destination())
	Equal(u.T(), root, retrievedAttestation.Attestation().Root())
	Len(u.T(), retrievedAttestation.NotarySignatures(), 1)
	Len(u.T(), retrievedAttestation.GuardSignatures(), 1)
	Greater(u.T(), retrievedNewestConfirmedOnDestination.SubmittedToAttestationCollectorTime().Unix(), int64(0))
	Equal(u.T(), types.AttestationStateConfirmedOnDestination, retrievedNewestConfirmedOnDestination.AttestationState())

	Nil(u.T(), err)
}
