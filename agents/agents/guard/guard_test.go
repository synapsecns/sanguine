package guard_test

import (
	"math/big"
	"os"
	"testing"
	"time"

	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/client"

	"github.com/Flaque/filet"
	awsTime "github.com/aws/smithy-go/time"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func (g GuardSuite) TestGuardE2E() {
	testConfig := config.AgentConfig{
		Domains: map[string]config.DomainConfig{
			"origin_client":      g.OriginDomainClient.Config(),
			"destination_client": g.DestinationDomainClient.Config(),
			"summit_client":      g.SummitDomainClient.Config(),
		},
		DomainID:       uint32(0),
		SummitDomainID: g.SummitDomainClient.Config().DomainID,
		BondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(g.T(), "", g.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(g.T(), "", g.GuardUnbondedWallet.PrivateKeyHex()).Name(),
		},
		RefreshIntervalSeconds: 5,
	}
	encodedTestConfig, err := testConfig.Encode()
	Nil(g.T(), err)

	tempConfigFile, err := os.CreateTemp("", "guard_temp_config.yaml")
	Nil(g.T(), err)
	defer RemoveGuardTempFile(g.T(), tempConfigFile.Name())

	numBytesWritten, err := tempConfigFile.Write(encodedTestConfig)
	Nil(g.T(), err)
	Positive(g.T(), numBytesWritten)

	decodedAgentConfig, err := config.DecodeAgentConfig(tempConfigFile.Name())
	Nil(g.T(), err)

	decodedAgentConfigBackToEncodedBytes, err := decodedAgentConfig.Encode()
	Nil(g.T(), err)

	Equal(g.T(), encodedTestConfig, decodedAgentConfigBackToEncodedBytes)

	omniRPCClient := omniClient.NewOmnirpcClient(g.TestOmniRPC, g.GuardMetrics, omniClient.WithCaptureReqRes())
	scribeClient := client.NewEmbeddedScribe("sqlite", g.DBPath, g.ScribeMetrics)
	go func() {
		scribeErr := scribeClient.Start(g.GetTestContext())
		Nil(g.T(), scribeErr)
	}()

	guard, err := guard.NewGuard(g.GetTestContext(), testConfig, omniRPCClient, scribeClient.ScribeClient, g.GuardTestDB, g.GuardMetrics)
	Nil(g.T(), err)

	tips := types.NewTips(big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)))

	optimisticSeconds := uint32(10)

	body := []byte{byte(gofakeit.Uint32())}

	txContextOrigin := g.TestBackendOrigin.GetTxContext(g.GetTestContext(), g.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)

	txContextTestClientOrigin := g.TestBackendOrigin.GetTxContext(g.GetTestContext(), g.TestClientMetadataOnOrigin.OwnerPtr())

	gasLimit := uint64(10000000)
	version := uint32(1)
	testClientOnOriginTx, err := g.TestClientOnOrigin.SendMessage(
		txContextTestClientOrigin.TransactOpts,
		uint32(g.TestBackendDestination.GetChainID()),
		g.TestClientMetadataOnDestination.Address(),
		optimisticSeconds,
		gasLimit,
		version,
		body)

	g.Nil(err)
	g.TestBackendOrigin.WaitForConfirmation(g.GetTestContext(), testClientOnOriginTx)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(g.GetTestContext())
	}()

	g.Eventually(func() bool {
		_ = awsTime.SleepWithContext(g.GetTestContext(), time.Second*5)

		rawState, err := g.SummitContract.GetLatestAgentState(
			&bind.CallOpts{Context: g.GetTestContext()},
			g.OriginDomainClient.Config().DomainID,
			g.GuardBondedSigner.Address())

		Nil(g.T(), err)

		if len(rawState) == 0 {
			return false
		}

		state, err := types.DecodeState(rawState)
		Nil(g.T(), err)
		return state.Nonce() >= uint32(1)
	})

	// Now make sure GetLatestState works as well
	g.Eventually(func() bool {
		_ = awsTime.SleepWithContext(g.GetTestContext(), time.Second*5)

		rawState, err := g.SummitContract.GetLatestState(
			&bind.CallOpts{Context: g.GetTestContext()},
			g.OriginDomainClient.Config().DomainID)
		Nil(g.T(), err)

		if len(rawState) == 0 {
			return false
		}

		state, err := types.DecodeState(rawState)
		Nil(g.T(), err)
		return state.Nonce() >= uint32(1)
	})
}
