package guard_test

import (
	"math/big"
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

func (u GuardSuite) TestGuardE2E() {
	u.T().Skip()
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
			File: filet.TmpFile(u.T(), "", u.NotaryWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.AttestationWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	guard, err := guard.NewGuard(u.GetTestContext(), testConfig)
	Nil(u.T(), err)

	dbType, err := dbcommon.DBTypeFromString(testConfig.Database.Type)
	Nil(u.T(), err)

	dbHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), dbType, testConfig.Database.ConnString)
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

	Greater(u.T(), 0, dispatchBlockNumber)

	NotEqual(u.T(), historicalRoot, [32]byte{})

	attestationKey := types.AttestationKey{
		Origin:      u.OriginDomainClient.Config().DomainID,
		Destination: u.DestinationDomainClient.Config().DomainID,
		Nonce:       nonce,
	}

	unsignedAttestation := types.NewAttestation(attestationKey.GetRawKey(), historicalRoot)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	Nil(u.T(), err)

	notarySignature, err := u.NotarySigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
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
		retrievedConfirmedInProgressAttestation, err := dbHandle.RetrieveInProgressAttestation(
			u.GetTestContext(),
			u.OriginDomainClient.Config().DomainID,
			u.DestinationDomainClient.Config().DomainID,
			nonce)
		Nil(u.T(), err)
		NotNil(u.T(), retrievedConfirmedInProgressAttestation)

		Equal(u.T(), u.OriginDomainClient.Config().DomainID, retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Origin())
		Equal(u.T(), u.OriginDomainClient.Config().DomainID, retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Destination())

		return retrievedConfirmedInProgressAttestation != nil &&
			retrievedConfirmedInProgressAttestation.SignedAttestation().Attestation().Nonce() != 0
	})
}
