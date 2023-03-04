package notary_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/types"
)

func (u *NotarySuite) TestOriginAttestationScanner() {
	// TODO (joeallen): FIX ME
	u.T().Skip()
	destinationDomain := uint32(u.TestBackendDestination.GetChainID())

	testDB, err := sqlite.NewSqliteStore(u.GetTestContext(), filet.TmpDir(u.T(), ""))
	Nil(u.T(), err)

	// dispatch a random update
	auth := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	optimisticSeconds := uint32(5)
	tx, err := u.OriginContract.Dispatch(auth.TransactOpts, destinationDomain, [32]byte{}, optimisticSeconds, encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

	// call the update producing function
	originAttestationScanner := notary.NewOriginAttestationScanner(
		u.OriginDomainClient,
		u.AttestationDomainClient,
		u.DestinationDomainClient,
		testDB,
		u.NotaryBondedSigner,
		u.NotaryUnbondedSigner,
		1*time.Second)

	err = originAttestationScanner.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	producedAttestation, err := testDB.RetrieveNewestInProgressAttestationIfInState(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID,
		types.AttestationStateNotaryUnsigned)
	Nil(u.T(), err)
	Equal(u.T(), producedAttestation.SignedAttestation().Attestation().Nonce(), uint32(1))
	Equal(u.T(), types.AttestationStateNotaryUnsigned, producedAttestation.AttestationState())
}
