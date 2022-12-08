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

func (u NotarySuite) TestOriginAttestationVerifier() {
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

	signedAttestation := types.NewSignedAttestation(unsignedInProgressAttestation.SignedAttestation().Attestation(), signature)
	signedInProgressAttestation := types.NewInProgressAttestation(signedAttestation, unsignedInProgressAttestation.OriginDispatchBlockNumber(), nil, 0, 0)
	err = testDB.UpdateSignature(u.GetTestContext(), signedInProgressAttestation)
	Nil(u.T(), err)

	err = u.domainClient.AttestationCollector().SubmitAttestation(u.GetTestContext(), u.signer, signedInProgressAttestation.SignedAttestation())
	Nil(u.T(), err)

	nowTime := time.Now()
	submittedInProgressAttestation := types.NewInProgressAttestation(signedInProgressAttestation.SignedAttestation(), signedInProgressAttestation.OriginDispatchBlockNumber(), &nowTime, 0, 0)
	err = testDB.UpdateSubmittedToAttestationCollectorTime(u.GetTestContext(), submittedInProgressAttestation)
	Nil(u.T(), err)

	// make sure an update has been produced
	inProgressAttestationToConfirm, err := testDB.RetrieveOldestUnconfirmedSubmittedInProgressAttestation(u.GetTestContext(), u.domainClient.Config().DomainID, u.destinationID)
	Nil(u.T(), err)
	Equal(u.T(), inProgressAttestationToConfirm.SignedAttestation().Attestation().Nonce(), fakeNonce)

	// call the update producing function
	originAttestationVerifier := notary.NewOriginAttestationVerifier(u.domainClient, u.destinationID, testDB, u.signer, 1*time.Second)

	err = originAttestationVerifier.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	producedAttestation, err := testDB.RetrieveInProgressAttestation(u.GetTestContext(), u.domainClient.Config().DomainID, u.destinationID, fakeNonce)
	Nil(u.T(), err)
	Equal(u.T(), fakeNonce, producedAttestation.SignedAttestation().Attestation().Nonce())
	Greater(u.T(), producedAttestation.ConfirmedOnAttestationCollectorBlockNumber(), uint64(0))
	Equal(u.T(), types.AttestationStateNotaryConfirmed, producedAttestation.AttestationState())
}
