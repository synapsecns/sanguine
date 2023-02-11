package guard_test

import (
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Flaque/filet"
	awsTime "github.com/aws/smithy-go/time"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

func RemoveGuardTempFile(t *testing.T, fileName string) {
	t.Helper()
	err := os.Remove(fileName)
	Nil(t, err)
}

func (u GuardSuite) TestGuardE2E() {
	testConfig := config.GuardConfig{
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		DestinationDomains: map[string]config.DomainConfig{
			"destination_client": u.DestinationDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardUnbondedWallet.PrivateKeyHex()).Name(),
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

	tempConfigFile, err := os.CreateTemp("", "guard_temp_config.yaml")
	Nil(u.T(), err)
	defer RemoveGuardTempFile(u.T(), tempConfigFile.Name())

	numBytesWritten, err := tempConfigFile.Write(encodedTestConfig)
	Nil(u.T(), err)
	Positive(u.T(), numBytesWritten)

	decodedGuardConfig, err := config.DecodeGuardConfig(tempConfigFile.Name())
	Nil(u.T(), err)

	decodedGuardConfigBackToEncodedBytes, err := decodedGuardConfig.Encode()
	Nil(u.T(), err)

	Equal(u.T(), encodedTestConfig, decodedGuardConfigBackToEncodedBytes)

	guard, err := guard.NewGuard(u.GetTestContext(), testConfig)
	Nil(u.T(), err)

	dbType, err := dbcommon.DBTypeFromString(testConfig.Database.Type)
	Nil(u.T(), err)

	dbHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), dbType, testConfig.Database.ConnString, "guard")
	Nil(u.T(), err)

	originAuth := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	tx, err := u.OriginContract.Dispatch(
		originAuth.TransactOpts,
		u.DestinationDomainClient.Config().DomainID,
		[32]byte{},
		gofakeit.Uint32(),
		encodedTips,
		[]byte(gofakeit.Paragraph(3, 2, 1, " ")))
	Nil(u.T(), err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

	nonce := uint32(1)

	historicalRoot, dispatchBlockNumber, err := u.OriginContract.GetHistoricalRoot(&bind.CallOpts{Context: u.GetTestContext()}, u.DestinationDomainClient.Config().DomainID, nonce)
	Nil(u.T(), err)

	Greater(u.T(), dispatchBlockNumber.Uint64(), uint64(0))

	NotEqual(u.T(), historicalRoot, [32]byte{})

	attestationKey := types.AttestationKey{
		Origin:      u.OriginDomainClient.Config().DomainID,
		Destination: u.DestinationDomainClient.Config().DomainID,
		Nonce:       nonce,
	}

	unsignedAttestation := types.NewAttestation(attestationKey.GetRawKey(), historicalRoot)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	Nil(u.T(), err)

	notarySignature, err := u.NotaryBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	Nil(u.T(), err)

	signedAttestation := types.NewSignedAttestation(
		unsignedAttestation,
		[]types.Signature{},
		[]types.Signature{notarySignature})

	encodedSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	Nil(u.T(), err)

	txContextAttestationCollector := u.TestBackendAttestation.GetTxContext(u.GetTestContext(), u.AttestationContractMetadata.OwnerPtr())
	// Submit the attestation to get an AttestationSubmitted event.
	txSubmitAttestation, err := u.AttestationContract.SubmitAttestation(txContextAttestationCollector.TransactOpts, encodedSignedAttestation)
	Nil(u.T(), err)
	u.TestBackendAttestation.WaitForConfirmation(u.GetTestContext(), txSubmitAttestation)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		retrievedInProgressAttestation, err := dbHandle.RetrieveNewestInProgressAttestationIfInState(
			u.GetTestContext(),
			u.OriginDomainClient.Config().DomainID,
			u.DestinationDomainClient.Config().DomainID,
			types.AttestationStateConfirmedOnDestination)

		isTrue := err == nil &&
			retrievedInProgressAttestation != nil &&
			retrievedInProgressAttestation.SignedAttestation().Attestation().Nonce() == nonce &&
			u.OriginDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Origin() &&
			u.DestinationDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Destination() &&
			historicalRoot == retrievedInProgressAttestation.SignedAttestation().Attestation().Root() &&
			len(retrievedInProgressAttestation.SignedAttestation().NotarySignatures()) == 1 &&
			len(retrievedInProgressAttestation.SignedAttestation().GuardSignatures()) == 1 &&
			retrievedInProgressAttestation.AttestationState() == types.AttestationStateConfirmedOnDestination

		return isTrue
	})
}
