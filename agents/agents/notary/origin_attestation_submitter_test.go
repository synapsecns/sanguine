package notary_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
)

func (u *NotarySuite) TestOriginAttestationSubmitter() {
	// TODO (joeallen): FIX ME
	u.T().Skip()
	testDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)

	fakeNonce := uint32(1)
	fakeRoot := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))

	fakeAttestKey := types.AttestationKey{
		Origin:      u.OriginDomainClient.Config().DomainID,
		Destination: u.DestinationDomainClient.Config().DomainID,
		Nonce:       fakeNonce,
	}
	fakeUnsignedAttestation := types.NewAttestation(fakeAttestKey.GetRawKey(), fakeRoot)

	err = testDB.StoreNewInProgressAttestation(u.GetTestContext(), fakeUnsignedAttestation)
	Nil(u.T(), err)

	unsignedInProgressAttestation, err := testDB.RetrieveInProgressAttestation(u.GetTestContext(), u.OriginDomainClient.Config().DomainID, u.DestinationDomainClient.Config().DomainID, fakeNonce)
	Nil(u.T(), err)

	hashedAttestation, err := types.Hash(unsignedInProgressAttestation.SignedAttestation().Attestation())
	Nil(u.T(), err)

	signature, err := u.NotaryBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(u.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedInProgressAttestation.SignedAttestation().Attestation(), []types.Signature{}, []types.Signature{signature})
	signedInProgressAttestation := types.NewInProgressAttestation(signedAttestation, nil, 0)
	err = testDB.UpdateNotarySignature(u.GetTestContext(), signedInProgressAttestation)
	Nil(u.T(), err)

	// call the update producing function
	originAttestationSubmitter := notary.NewOriginAttestationSubmitter(
		u.OriginDomainClient,
		u.AttestationDomainClient,
		u.DestinationDomainClient,
		testDB,
		u.NotaryBondedSigner,
		u.NotaryUnbondedSigner,
		1*time.Second)

	err = originAttestationSubmitter.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	producedAttestation, err := testDB.RetrieveNewestInProgressAttestationIfInState(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID,
		types.AttestationStateNotarySubmittedUnconfirmed)
	Nil(u.T(), err)
	Equal(u.T(), producedAttestation.SignedAttestation().Attestation().Nonce(), fakeNonce)
	NotNil(u.T(), producedAttestation.SubmittedToAttestationCollectorTime())
	Equal(u.T(), types.AttestationStateNotarySubmittedUnconfirmed, producedAttestation.AttestationState())
}
