package guard_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/types"
)

func (u GuardSuite) TestAttestationCollectorAttestationScanner() {
	// TODO (joeallen): FIX ME
	u.T().Skip()
	// TODO (joeallen): FIX ME
	testDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)

	origin := uint32(u.TestBackendOrigin.GetChainID())
	destination := uint32(u.TestBackendDestination.GetChainID())
	// nonce := uint32(1)

	// dispatch a random update
	originAuth := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	optimisticSeconds := uint32(5)
	tx, err := u.OriginContract.Dispatch(originAuth.TransactOpts, destination, [32]byte{}, optimisticSeconds, encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

	// TODO (joeallen): FIX ME
	// suggestedAttestationRaw, err := u.OriginContract.SuggestAttestation(&bind.CallOpts{Context: u.GetTestContext()}, destination)
	// Nil(u.T(), err)
	// suggestedAttestation, err := types.DecodeAttestation(suggestedAttestationRaw)
	// Nil(u.T(), err)
	//Equal(u.T(), origin, suggestedAttestation.Origin())
	//Equal(u.T(), destination, suggestedAttestation.Destination())
	//Equal(u.T(), nonce, suggestedAttestation.Nonce())

	// err = testDB.StoreNewGuardInProgressAttestation(u.GetTestContext(), suggestedAttestation)
	// Nil(u.T(), err)

	// auth := u.TestBackendAttestation.GetTxContext(u.GetTestContext(), nil)

	// root := suggestedAttestation.Root()

	// attestKey := types.AttestationKey{
	//	Origin:      origin,
	//	Destination: destination,
	//	Nonce:       nonce,
	//}
	// TODO (joeallen): FIX ME
	// unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), [32]byte{})
	// hashedAttestation, err := types.Hash(unsignedAttestation)
	// Nil(u.T(), err)

	// notarySignature, err := u.NotaryBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	// Nil(u.T(), err)

	// signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{}, []types.Signature{notarySignature})

	// rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	// Nil(u.T(), err)

	// tx, err = u.AttestationContract.SubmitAttestation(auth.TransactOpts, rawSignedAttestation)
	// Nil(u.T(), err)

	u.TestBackendAttestation.WaitForConfirmation(u.GetTestContext(), tx)

	// call the update producing function
	attestationCollectorAttestationScanner := guard.NewAttestationCollectorAttestationScanner(
		u.AttestationDomainClient,
		origin,
		destination,
		testDB,
		u.GuardUnbondedSigner,
		1*time.Second)

	err = attestationCollectorAttestationScanner.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	retrievedConfirmedInProgressAttestation, err := testDB.RetrieveNewestInProgressAttestationIfInState(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID,
		types.AttestationStateGuardUnsignedAndVerified)

	Nil(u.T(), err)
	NotNil(u.T(), retrievedConfirmedInProgressAttestation)

	retrievedSignedAttestation := retrievedConfirmedInProgressAttestation.SignedAttestation()
	Equal(u.T(), u.OriginDomainClient.Config().DomainID, retrievedSignedAttestation.Attestation().Origin())
	Equal(u.T(), u.DestinationDomainClient.Config().DomainID, retrievedSignedAttestation.Attestation().Destination())
	// Equal(u.T(), root, retrievedSignedAttestation.Attestation().Root())
	Len(u.T(), retrievedSignedAttestation.NotarySignatures(), 1)
	Len(u.T(), retrievedSignedAttestation.GuardSignatures(), 0)
	Equal(u.T(), types.AttestationStateGuardUnsignedAndVerified, retrievedConfirmedInProgressAttestation.AttestationState())

	Nil(u.T(), err)
}
