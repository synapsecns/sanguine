package guard_test

import (
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/types"
)

func (u *GuardSuite) TestOriginGuardAttestationScanner() {
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
	originGuardAttestationScanner := guard.NewOriginGuardAttestationScanner(
		u.OriginDomainClient,
		u.AttestationDomainClient,
		u.DestinationDomainClient,
		testDB,
		u.GuardBondedSigner,
		u.GuardUnbondedSigner,
		1*time.Second)

	err = originGuardAttestationScanner.Update(u.GetTestContext())
	Nil(u.T(), err)

	// make sure an update has been produced
	producedAttestation, err := testDB.RetrieveNewestInProgressAttestationIfInState(
		u.GetTestContext(),
		u.OriginDomainClient.Config().DomainID,
		u.DestinationDomainClient.Config().DomainID,
		types.AttestationStateGuardInitialState)
	Nil(u.T(), err)
	Equal(u.T(), producedAttestation.SignedAttestation().Attestation().Nonce(), uint32(1))
	Equal(u.T(), types.AttestationStateGuardInitialState, producedAttestation.AttestationState())
}
