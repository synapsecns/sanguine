package notary_test

import (
	"math/big"
	"os"
	"testing"
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

func RemoveNotaryTempFile(t *testing.T, fileName string) {
	t.Helper()
	err := os.Remove(fileName)
	Nil(t, err)
}

func (u *NotarySuite) TestNotaryE2E() {
	testConfig := config.NotaryConfig{
		DestinationDomain: u.DestinationDomainClient.Config(),
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	encodedTestConfig, err := testConfig.Encode()
	Nil(u.T(), err)

	tempConfigFile, err := os.CreateTemp("", "notary_temp_config.yaml")
	Nil(u.T(), err)
	defer RemoveNotaryTempFile(u.T(), tempConfigFile.Name())

	numBytesWritten, err := tempConfigFile.Write(encodedTestConfig)
	Nil(u.T(), err)
	Positive(u.T(), numBytesWritten)

	decodedNotaryConfig, err := config.DecodeNotaryConfig(tempConfigFile.Name())
	Nil(u.T(), err)

	decodedNotaryConfigBackToEncodedBytes, err := decodedNotaryConfig.Encode()
	Nil(u.T(), err)

	Equal(u.T(), encodedTestConfig, decodedNotaryConfigBackToEncodedBytes)

	notary, err := notary.NewNotary(u.GetTestContext(), testConfig)
	Nil(u.T(), err)

	dbType, err := dbcommon.DBTypeFromString(testConfig.Database.Type)
	Nil(u.T(), err)

	dbHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), dbType, testConfig.Database.ConnString, "notary")
	Nil(u.T(), err)

	auth := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	tx, err := u.OriginContract.Dispatch(auth.TransactOpts, testConfig.DestinationDomain.DomainID, [32]byte{}, gofakeit.Uint32(), encodedTips, []byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		retrievedConfirmedInProgressAttestation, err := dbHandle.RetrieveNewestInProgressAttestationIfInState(
			u.GetTestContext(),
			u.OriginDomainClient.Config().DomainID,
			testConfig.DestinationDomain.DomainID,
			types.AttestationStateNotaryConfirmed)

		return err == nil &&
			retrievedConfirmedInProgressAttestation != nil &&
			u.OriginDomainClient.Config().DomainID == retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Origin() &&
			testConfig.DestinationDomain.DomainID == retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Destination() &&
			types.AttestationStateNotaryConfirmed == retrievedConfirmedInProgressAttestation.AttestationState() &&
			retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Nonce() != 0
	})
}
