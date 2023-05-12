package agentsintegration_test

import (
	awsTime "github.com/aws/smithy-go/time"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	executorCfg "github.com/synapsecns/sanguine/agents/agents/executor/config"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/types"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeConfig2 "github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/config"
)

func RemoveAgentsTempFile(t *testing.T, fileName string) {
	t.Helper()
	err := os.Remove(fileName)
	Nil(t, err)
}

//nolint:cyclop,maintidx
func (u *AgentsIntegrationSuite) TestAgentsE2E() {
	// TODO (joe and lex): FIX ME
	// u.T().Skip()
	testDone := false
	defer func() {
		testDone = true
	}()

	_, testContractRef := u.TestDeployManager.GetAgentsTestContract(u.GetTestContext(), u.TestBackendDestination)
	testTransactOpts := u.TestBackendDestination.GetTxContext(u.GetTestContext(), nil)

	originClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendOrigin.RPCAddress(), u.ScribeMetrics)
	u.Nil(err)
	destinationClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendDestination.RPCAddress(), u.ScribeMetrics)
	u.Nil(err)
	summitClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendSummit.RPCAddress(), u.ScribeMetrics)
	u.Nil(err)

	originConfig := scribeConfig2.ContractConfig{
		Address:    u.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig2.ChainConfig{
		ChainID:               uint32(u.TestBackendOrigin.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig2.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig2.ContractConfig{
		Address:    u.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig2.ChainConfig{
		ChainID:               uint32(u.TestBackendDestination.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig2.ContractConfig{destinationConfig},
	}
	summitConfig := scribeConfig2.ContractConfig{
		Address:    u.BondingManagerOnSummit.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := scribeConfig2.ChainConfig{
		ChainID:               uint32(u.TestBackendSummit.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig2.ContractConfig{summitConfig},
	}
	scribeConfig := scribeConfig2.Config{
		Chains: []scribeConfig2.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		uint32(u.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(u.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
		uint32(u.TestBackendSummit.GetChainID()):      {summitClient, summitClient},
	}

	scribe, err := node.NewScribe(u.ScribeTestDB, clients, scribeConfig, u.ScribeMetrics)
	u.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", u.DBPath, u.ScribeMetrics)
	go func() {
		scribeErr := scribeClient.Start(u.GetTestContext())
		u.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		scribeError := scribe.Start(u.GetTestContext())
		if !testDone {
			u.Nil(scribeError)
		}
	}()

	chainID := uint32(u.TestBackendOrigin.GetChainID())
	destination := uint32(u.TestBackendDestination.GetChainID())
	summit := uint32(u.TestBackendSummit.GetChainID())

	excCfg := executorCfg.Config{
		SummitChainID:         summit,
		SummitAddress:         u.SummitContract.Address().String(),
		BondingManagerAddress: u.BondingManagerOnSummit.Address().String(),
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       chainID,
				OriginAddress: u.OriginContract.Address().String(),
			},
			{
				ChainID:            destination,
				DestinationAddress: u.DestinationContract.Address().String(),
			},
			{
				ChainID: summit,
			},
		},
		BaseOmnirpcURL: u.TestBackendOrigin.RPCAddress(),
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		chainID:     u.TestBackendOrigin,
		destination: u.TestBackendDestination,
		summit:      u.TestBackendSummit,
	}

	urls := map[uint32]string{
		chainID:     u.TestBackendOrigin.RPCAddress(),
		destination: u.TestBackendDestination.RPCAddress(),
		summit:      u.TestBackendSummit.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(u.GetTestContext(), excCfg, u.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls, u.ExecutorMetrics)
	Nil(u.T(), err)

	go func() {
		execErr := exec.Run(u.GetTestContext())
		if !testDone {
			Nil(u.T(), execErr)
		}
	}()

	/*attestationSavedSink := make(chan *summitharness.SummitHarnessAttestationSaved)
	savedAttestation, err := u.SummitContract.WatchAttestationSaved(&bind.WatchOpts{Context: u.GetTestContext()}, attestationSavedSink)
	Nil(u.T(), err)*/

	guardTestConfig := config.AgentConfig{
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
	notaryTestConfig := config.AgentConfig{
		Domains: map[string]config.DomainConfig{
			"origin_client":      u.OriginDomainClient.Config(),
			"destination_client": u.DestinationDomainClient.Config(),
			"summit_client":      u.SummitDomainClient.Config(),
		},
		DomainID:       u.DestinationDomainClient.Config().DomainID,
		SummitDomainID: u.SummitDomainClient.Config().DomainID,
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
		RefreshIntervalSeconds: 5,
	}
	encodedNotaryTestConfig, err := notaryTestConfig.Encode()
	Nil(u.T(), err)

	notaryTempConfigFile, err := os.CreateTemp("", "notary_temp_config.yaml")
	Nil(u.T(), err)
	defer RemoveAgentsTempFile(u.T(), notaryTempConfigFile.Name())

	numBytesWritten, err := notaryTempConfigFile.Write(encodedNotaryTestConfig)
	Nil(u.T(), err)
	Positive(u.T(), numBytesWritten)

	decodedAgentConfig, err := config.DecodeAgentConfig(notaryTempConfigFile.Name())
	Nil(u.T(), err)

	decodedAgentConfigBackToEncodedBytes, err := decodedAgentConfig.Encode()
	Nil(u.T(), err)

	Equal(u.T(), encodedNotaryTestConfig, decodedAgentConfigBackToEncodedBytes)

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig, u.GuardMetrics)
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
		err = guard.Start(u.GetTestContext())
		u.Nil(err)
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

	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig, u.NotaryMetrics)
	Nil(u.T(), err)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	waitChan := make(chan bool, 1)

	// Make sure there is one executable message in the database.
	u.Eventually(func() bool {
		mask := execTypes.DBMessage{
			ChainID:     &chainID,
			Destination: &destination,
		}
		executableMessages, err := u.ExecutorTestDB.GetExecutableMessages(u.GetTestContext(), mask, uint64(time.Now().Unix()), 1)
		u.Nil(err)
		if len(executableMessages) == 1 {
			waitChan <- true
			return true
		}
		return false
	})

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)

		rawState, err := u.SummitContract.GetLatestAgentState(
			&bind.CallOpts{Context: u.GetTestContext()},
			u.OriginDomainClient.Config().DomainID,
			u.NotaryBondedSigner.Address())
		Nil(u.T(), err)

		if len(rawState) == 0 {
			return false
		}

		state, err := types.DecodeState(rawState)
		Nil(u.T(), err)
		return state.Nonce() >= uint32(1)
	})

	/*watchCtx, cancel := context.WithCancel(u.GetTestContext())
	defer cancel()

	var retrievedAtt []byte
	select {
	// check for errors and fail
	case <-watchCtx.Done():
		retrievedAtt = []byte{}
		break
	case <-savedAttestation.Err():
		Nil(u.T(), savedAttestation.Err())
		retrievedAtt = []byte{}
		break
	// get message sent event
	case receivedAttestationSaved := <-attestationSavedSink:
		attToSubmit := receivedAttestationSaved.Attestation
		retrievedAtt = attToSubmit
		break
	}

	Greater(u.T(), len(retrievedAtt), 0)

	<-waitChan*/

	u.Eventually(func() bool {
		mask := execTypes.DBMessage{
			ChainID:     &chainID,
			Destination: &destination,
		}
		executedMessages, err := u.ExecutorTestDB.GetExecutableMessages(u.GetTestContext(), mask, uint64(time.Now().Unix()), 1)
		Nil(u.T(), err)
		if len(executedMessages) == 0 {
			return true
		}

		// Need to create a tx and wait for it to be confirmed to continue adding blocks, and therefore
		// increase the `time`.
		countBeforeIncrement, err := testContractRef.GetCount(&bind.CallOpts{Context: u.GetTestContext()})
		u.Nil(err)
		testTx, err := testContractRef.IncrementCounter(testTransactOpts.TransactOpts)
		u.Nil(err)
		u.TestBackendDestination.WaitForConfirmation(u.GetTestContext(), testTx)
		countAfterIncrement, err := testContractRef.GetCount(&bind.CallOpts{Context: u.GetTestContext()})
		u.Nil(err)
		u.Greater(countAfterIncrement.Uint64(), countBeforeIncrement.Uint64())
		return false
	})
}
