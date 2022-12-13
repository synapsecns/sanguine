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

func (u NotarySuite) TestOriginAttestationSubmitter() {
	testDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)

	fakeNonce := uint32(1)
	fakeRoot := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	fakeDispatchBlockNumber := uint64(1)

	fakeAttestKey := types.AttestationKey{
		Origin:      u.domainClient.Config().DomainID,
		Destination: u.destinationID,
		Nonce:       fakeNonce,
	}
	fakeUnsignedAttestation := types.NewAttestation(fakeAttestKey.GetRawKey(), fakeRoot)

	err = testDB.StoreNewInProgressAttestation(u.GetTestContext(), fakeUnsignedAttestation, fakeDispatchBlockNumber)
	Nil(u.T(), err)

	unsignedInProgressAttestation, err := testDB.RetrieveInProgressAttestation(u.GetTestContext(), u.domainClient.Config().DomainID, u.destinationID, fakeNonce)
	Nil(u.T(), err)

	hashedUpdate, err := types.Hash(unsignedInProgressAttestation.SignedAttestation().Attestation())
	Nil(u.T(), err)

	signature, err := u.signer.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedUpdate), false)
	Nil(u.T(), err)

	signedAttestation := types.NewSignedAttestation(unsignedInProgressAttestation.SignedAttestation().Attestation(), []types.Signature{}, []types.Signature{signature})
	signedInProgressAttestation := types.NewInProgressAttestation(signedAttestation, unsignedInProgressAttestation.OriginDispatchBlockNumber(), nil, 0, 0)
	err = testDB.UpdateSignature(u.GetTestContext(), signedInProgressAttestation)
	Nil(u.T(), err)

	// call the update producing function
	originAttestationSubmitter := notary.NewOriginAttestationSubmitter(u.domainClient, u.destinationID, testDB, u.signer, 1*time.Second)

	err = originAttestationSubmitter.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	producedAttestation, err := testDB.RetrieveOldestUnconfirmedSubmittedInProgressAttestation(u.GetTestContext(), u.domainClient.Config().DomainID, u.destinationID)
	Nil(u.T(), err)
	Equal(u.T(), producedAttestation.SignedAttestation().Attestation().Nonce(), fakeNonce)
	NotNil(u.T(), producedAttestation.SubmittedToAttestationCollectorTime())
	Equal(u.T(), types.AttestationStateNotarySubmittedUnconfirmed, producedAttestation.AttestationState())
}
