package guard_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	awsTime "github.com/aws/smithy-go/time"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

func (u NotarySuite) TestNotaryE2E() {
	u.T().Skip()
	testConfig := config.NotaryConfig{
		DestinationID: u.destinationID,
		Domains: map[string]config.DomainConfig{
			"test": u.domainClient.Config(),
		},
		Signer: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.wallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	ud, err := notary.NewNotary(u.GetTestContext(), testConfig)
	Nil(u.T(), err)

	dbType, err := dbcommon.DBTypeFromString(testConfig.Database.Type)
	Nil(u.T(), err)

	dbHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), dbType, testConfig.Database.ConnString)
	Nil(u.T(), err)

	auth := u.testBackend.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	tx, err := u.originContract.Dispatch(auth.TransactOpts, testConfig.DestinationID, [32]byte{}, gofakeit.Uint32(), encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = ud.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		retrievedConfirmedInProgressAttestation, err := dbHandle.RetrieveNewestConfirmedInProgressAttestation(u.GetTestContext(), u.domainClient.Config().DomainID, testConfig.DestinationID)
		Nil(u.T(), err)
		NotNil(u.T(), retrievedConfirmedInProgressAttestation)

		Equal(u.T(), u.domainClient.Config().DomainID, retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Origin())
		Equal(u.T(), testConfig.DestinationID, retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Destination())
		Equal(u.T(), types.AttestationStateNotaryConfirmed, retrievedConfirmedInProgressAttestation.AttestationState())

		return retrievedConfirmedInProgressAttestation != nil &&
			retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Nonce() != 0
	})
}
