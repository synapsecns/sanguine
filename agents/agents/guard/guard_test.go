package guard_test

import (
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Flaque/filet"
	awsTime "github.com/aws/smithy-go/time"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/types"
)

func RemoveGuardTempFile(t *testing.T, fileName string) {
	t.Helper()
	err := os.Remove(fileName)
	Nil(t, err)
}

func (u GuardSuite) TestGuardE2E() {
	// TODO (joeallen): FIX ME
	u.T().Skip()
	testConfig := config.AgentConfig{
		Domains: map[string]config.DomainConfig{
			"origin_client":      u.OriginDomainClient.Config(),
			"destination_client": u.DestinationDomainClient.Config(),
			"summit_client":      u.SummitDomainClient.Config(),
		},
		DomainID:       uint32(0),
		SummitDomainID: u.SummitDomainClient.Config().DomainID,
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}
	encodedTestConfig, err := testConfig.Encode()
	Nil(u.T(), err)

	tempConfigFile, err := os.CreateTemp("", "guard_temp_config.yaml")
	Nil(u.T(), err)
	defer RemoveGuardTempFile(u.T(), tempConfigFile.Name())

	numBytesWritten, err := tempConfigFile.Write(encodedTestConfig)
	Nil(u.T(), err)
	Positive(u.T(), numBytesWritten)

	decodedAgentConfig, err := config.DecodeAgentConfig(tempConfigFile.Name())
	Nil(u.T(), err)

	decodedAgentConfigBackToEncodedBytes, err := decodedAgentConfig.Encode()
	Nil(u.T(), err)

	Equal(u.T(), encodedTestConfig, decodedAgentConfigBackToEncodedBytes)

	guard, err := guard.NewGuard(u.GetTestContext(), testConfig)
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

	// TODO (joeallen): FIX ME
	// nonce := uint32(1)
	// historicalRoot, dispatchBlockNumber, err := u.OriginContract.GetHistoricalRoot(&bind.CallOpts{Context: u.GetTestContext()}, u.DestinationDomainClient.Config().DomainID, nonce)
	// Nil(u.T(), err)

	// Greater(u.T(), dispatchBlockNumber.Uint64(), uint64(0))

	// NotEqual(u.T(), historicalRoot, [32]byte{})

	// attestationKey := types.AttestationKey{
	//	Origin:      u.OriginDomainClient.Config().DomainID,
	//	Destination: u.DestinationDomainClient.Config().DomainID,
	//	Nonce:       nonce,
	//}

	// unsignedAttestation := types.NewAttestation(attestationKey.GetRawKey(), historicalRoot)
	// hashedAttestation, err := types.Hash(unsignedAttestation)
	// Nil(u.T(), err)

	// notarySignature, err := u.NotaryBondedSigner.SignMessage(u.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	// Nil(u.T(), err)

	// signedAttestation := types.NewSignedAttestation(
	//	unsignedAttestation,
	//	[]types.Signature{},
	//	[]types.Signature{notarySignature})

	// encodedSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	// Nil(u.T(), err)

	// txContextAttestationCollector := u.TestBackendAttestation.GetTxContext(u.GetTestContext(), u.AttestationContractMetadata.OwnerPtr())
	// Submit the attestation to get an AttestationSubmitted event.
	// txSubmitAttestation, err := u.AttestationContract.SubmitAttestation(txContextAttestationCollector.TransactOpts, encodedSignedAttestation)
	// Nil(u.T(), err)
	// u.TestBackendAttestation.WaitForConfirmation(u.GetTestContext(), txSubmitAttestation)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)

		return true
	})
}
