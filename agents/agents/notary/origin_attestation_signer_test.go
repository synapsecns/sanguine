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
)

func (u *NotarySuite) TestOriginAttestationSigner() {
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

	// call the update producing function
	originAttestationSigner := notary.NewOriginAttestationSigner(
		u.OriginDomainClient,
		u.AttestationDomainClient,
		u.DestinationDomainClient,
		testDB,
		u.NotaryBondedSigner,
		u.NotaryUnbondedSigner,
		time.Second)

	err = originAttestationSigner.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	producedAttestation, err := testDB.RetrieveNewestInProgressAttestationIfInState(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID,
		types.AttestationStateNotarySignedUnsubmitted)
	Nil(u.T(), err)
	Equal(u.T(), producedAttestation.SignedAttestation().Attestation().Nonce(), fakeNonce)
	NotNil(u.T(), producedAttestation.SignedAttestation().NotarySignatures()[0])
	Equal(u.T(), types.AttestationStateNotarySignedUnsubmitted, producedAttestation.AttestationState())
}
