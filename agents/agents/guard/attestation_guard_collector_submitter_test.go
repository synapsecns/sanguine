// nolint:dupl
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

func (u GuardSuite) TestAttestationGuardCollectorSubmitter() {
	// TODO (joeallen): FIX ME
	u.T().Skip()
	testDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)

	// TODO (joeallen): FIX ME
	// origin := uint32(u.TestBackendOrigin.GetChainID())
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
	// Equal(u.T(), origin, suggestedAttestation.Origin())
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
	// unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	// hashedAttestation, err := types.Hash(unsignedAttestation)
	// Nil(u.T(), err)

	// notarySignature, err := u.NotaryBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	// Nil(u.T(), err)

	// signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{}, []types.Signature{notarySignature})

	// rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	// Nil(u.T(), err)

	// tx, err = u.AttestationContract.SubmitAttestation(auth.TransactOpts, rawSignedAttestation)
	// Nil(u.T(), err)

	// u.TestBackendAttestation.WaitForConfirmation(u.GetTestContext(), tx)

	// rawSignedAttestationFromCollector, err := u.AttestationContract.GetAttestation(&bind.CallOpts{Context: u.GetTestContext()}, origin, destination, nonce)
	// Nil(u.T(), err)

	// signedAttestationFromCollector, err := types.DecodeSignedAttestation(rawSignedAttestationFromCollector)
	// Nil(u.T(), err)

	// err = testDB.StoreExistingSignedInProgressAttestation(u.GetTestContext(), signedAttestationFromCollector)
	// Nil(u.T(), err)

	// guardSignature, err := u.GuardBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	// Nil(u.T(), err)

	// guardSignedAttestation := types.NewSignedAttestation(
	//	unsignedAttestation,
	//	[]types.Signature{guardSignature},
	//	[]types.Signature{notarySignature})
	// signedInProgressAttestation := types.NewInProgressAttestation(
	//	guardSignedAttestation,
	//	nil,
	//	0)
	// err = testDB.UpdateGuardSignature(u.GetTestContext(), signedInProgressAttestation)
	// Nil(u.T(), err)

	// Now call the guard submitter
	attestationGuardCollectorSubmitter := guard.NewAttestationGuardCollectorSubmitter(
		u.OriginDomainClient,
		u.AttestationDomainClient,
		u.DestinationDomainClient,
		testDB,
		u.GuardBondedSigner,
		u.GuardUnbondedSigner,
		1*time.Second)

	err = attestationGuardCollectorSubmitter.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure the attesation has been submitted
	retrievedNewestGuardSubmittedToCollectorUnconfirmed, err := testDB.RetrieveNewestInProgressAttestationIfInState(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID,
		types.AttestationStateGuardSubmittedToCollectorUnconfirmed)

	Nil(u.T(), err)
	NotNil(u.T(), retrievedNewestGuardSubmittedToCollectorUnconfirmed)

	retrievedAttestation := retrievedNewestGuardSubmittedToCollectorUnconfirmed.SignedAttestation()
	Equal(u.T(), u.OriginDomainClient.Config().DomainID, retrievedAttestation.Attestation().Origin())
	Equal(u.T(), u.DestinationDomainClient.Config().DomainID, retrievedAttestation.Attestation().Destination())
	// Equal(u.T(), root, retrievedAttestation.Attestation().Root())
	Len(u.T(), retrievedAttestation.NotarySignatures(), 1)
	Len(u.T(), retrievedAttestation.GuardSignatures(), 1)
	Greater(u.T(), retrievedNewestGuardSubmittedToCollectorUnconfirmed.SubmittedToAttestationCollectorTime().Unix(), int64(0))
	Equal(u.T(), types.AttestationStateGuardSubmittedToCollectorUnconfirmed, retrievedNewestGuardSubmittedToCollectorUnconfirmed.AttestationState())

	Nil(u.T(), err)
}
