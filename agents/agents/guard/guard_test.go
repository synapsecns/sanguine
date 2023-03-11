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
	"github.com/synapsecns/sanguine/agents/types"
)

func RemoveGuardTempFile(t *testing.T, fileName string) {
	t.Helper()
	err := os.Remove(fileName)
	Nil(t, err)
}

func (u GuardSuite) TestGuardE2E() {
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
			File: filet.TmpFile(u.T(), "", u.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardUnbondedWallet.PrivateKeyHex()).Name(),
		},
		RefreshIntervalSeconds: 5,
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

	tips := types.NewTips(big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)))

	optimisticSeconds := uint32(10)

	body := []byte{byte(gofakeit.Uint32())}

	txContextOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)

	txContextTestClientOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.TestClientMetadataOnOrigin.OwnerPtr())

	testClientOnOriginTx, err := u.TestClientOnOrigin.SendMessage(
		txContextTestClientOrigin.TransactOpts,
		uint32(u.TestBackendDestination.GetChainID()),
		u.TestClientMetadataOnDestination.Address(),
		optimisticSeconds,
		body)

	u.Nil(err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), testClientOnOriginTx)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)

		rawState, err := u.SummitContract.GetLatestAgentState(
			&bind.CallOpts{Context: u.GetTestContext()},
			u.OriginDomainClient.Config().DomainID,
			u.GuardBondedSigner.Address())

		Nil(u.T(), err)

		if len(rawState) == 0 {
			return false
		}

		state, err := types.DecodeState(rawState)
		Nil(u.T(), err)
		return state.Nonce() >= uint32(1)
	})

	// Now make sure GetLatestState works as well
	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)

		rawState, err := u.SummitContract.GetLatestState(
			&bind.CallOpts{Context: u.GetTestContext()},
			u.OriginDomainClient.Config().DomainID)
		Nil(u.T(), err)

		if len(rawState) == 0 {
			return false
		}

		state, err := types.DecodeState(rawState)
		Nil(u.T(), err)
		return state.Nonce() >= uint32(1)
	})
}

func (u GuardSuite) TestDeployedGuards() {
	allOriginGuards, err := u.OriginContract.AllAgents(&bind.CallOpts{Context: u.GetTestContext()}, uint32(0))
	Nil(u.T(), err)
	Equal(u.T(), 1, len(allOriginGuards))

	allDestGuards, err := u.DestinationContract.AllAgents(&bind.CallOpts{Context: u.GetTestContext()}, uint32(0))
	Nil(u.T(), err)
	Equal(u.T(), 1, len(allDestGuards))

	Equal(u.T(), allOriginGuards[0], allDestGuards[0])

	allSummitGuards, err := u.SummitContract.AllAgents(&bind.CallOpts{Context: u.GetTestContext()}, uint32(0))
	Nil(u.T(), err)

	Equal(u.T(), 1, len(allSummitGuards))

	Equal(u.T(), allSummitGuards[0], allDestGuards[0])
}
