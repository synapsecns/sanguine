package agentsintegration_test

import (
	"math/big"
	"os"
	"testing"
	"time"

	awsTime "github.com/aws/smithy-go/time"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/config"
	execConfig "github.com/synapsecns/sanguine/agents/config/executor"
	"github.com/synapsecns/sanguine/agents/types"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/service"

	"github.com/Flaque/filet"
)

func RemoveAgentsTempFile(t *testing.T, fileName string) {
	t.Helper()
	err := os.Remove(fileName)
	Nil(t, err)
}

//nolint:cyclop,maintidx
func (u *AgentsIntegrationSuite) TestAgentsE2E() {
	testDone := false
	defer func() {
		testDone = true
	}()

	originClient, err := backend.DialBackend(u.GetTestContext(), u.TestBackendOrigin.RPCAddress(), u.ScribeMetrics)
	u.Nil(err)
	destinationClient, err := backend.DialBackend(u.GetTestContext(), u.TestBackendDestination.RPCAddress(), u.ScribeMetrics)
	u.Nil(err)
	summitClient, err := backend.DialBackend(u.GetTestContext(), u.TestBackendSummit.RPCAddress(), u.ScribeMetrics)
	u.Nil(err)

	originConfig := scribeConfig.ContractConfig{
		Address:    u.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig.ChainConfig{
		ChainID:            uint32(u.TestBackendOrigin.GetChainID()),
		GetLogsBatchAmount: 1,
		StoreConcurrency:   1,
		GetLogsRange:       1,
		Confirmations:      0,
		Contracts:          []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    u.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:            uint32(u.TestBackendDestination.GetChainID()),
		GetLogsBatchAmount: 1,
		StoreConcurrency:   1,
		GetLogsRange:       1,
		Confirmations:      0,
		Contracts:          []scribeConfig.ContractConfig{destinationConfig},
	}
	summitConfig := scribeConfig.ContractConfig{
		Address:    u.InboxOnSummit.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := scribeConfig.ChainConfig{
		ChainID:            uint32(u.TestBackendSummit.GetChainID()),
		GetLogsBatchAmount: 1,
		StoreConcurrency:   1,
		GetLogsRange:       1,
		Confirmations:      0,

		Contracts: []scribeConfig.ContractConfig{summitConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}
	clients := map[uint32][]backend.ScribeBackend{
		uint32(u.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(u.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
		uint32(u.TestBackendSummit.GetChainID()):      {summitClient, summitClient},
	}

	scribe, err := service.NewScribe(u.ScribeTestDB, clients, scribeConfig, u.ScribeMetrics)
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

	excCfg := execConfig.Config{
		SummitChainID: summit,
		SummitAddress: u.SummitContract.Address().String(),
		InboxAddress:  u.InboxOnSummit.Address().String(),
		Chains: []execConfig.ChainConfig{
			{
				ChainID:       chainID,
				OriginAddress: u.OriginContract.Address().String(),
			},
			{
				ChainID:            destination,
				LightInboxAddress:  u.LightInboxOnDestination.Address().String(),
				DestinationAddress: u.DestinationContract.Address().String(),
			},
			{
				ChainID: summit,
			},
		},
		BaseOmnirpcURL: u.TestBackendOrigin.RPCAddress(),
		UnbondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
		SubmitterConfig: submitterConfig.Config{
			ChainConfig: submitterConfig.ChainConfig{
				GasEstimate: uint64(5000000),
			},
		},
	}

	omniRPCClient := omniClient.NewOmnirpcClient(u.TestOmniRPC, u.ExecutorMetrics, omniClient.WithCaptureReqRes())

	exec, err := executor.NewExecutor(u.GetTestContext(), excCfg, u.ExecutorTestDB, scribeClient.ScribeClient, omniRPCClient, u.ExecutorMetrics)
	Nil(u.T(), err)

	go func() {
		execErr := exec.Run(u.GetTestContext())
		if !testDone {
			Nil(u.T(), execErr)
		}
	}()

	guardTestConfig := config.AgentConfig{
		Domains: map[string]config.DomainConfig{
			"origin_client":      u.OriginDomainClient.Config(),
			"destination_client": u.DestinationDomainClient.Config(),
			"summit_client":      u.SummitDomainClient.Config(),
		},
		DomainID:       uint32(0),
		SummitDomainID: u.SummitDomainClient.Config().DomainID,
		BondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
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
		BondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
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

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig, omniRPCClient, scribeClient.ScribeClient, u.GuardTestDB, u.GuardMetrics)
	Nil(u.T(), err)

	tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
	optimisticSeconds := uint32(1)
	recipientDestination := u.TestClientMetadataOnDestination.Address().Hash()
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}
	txContextOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)
	paddedRequest := big.NewInt(0)

	msgSender := common.BytesToHash(txContextOrigin.TransactOpts.From.Bytes())
	header := types.NewHeader(types.MessageFlagBase, uint32(u.TestBackendOrigin.GetChainID()), nonce, uint32(u.TestBackendDestination.GetChainID()), optimisticSeconds)
	msgRequest := types.NewRequest(uint32(0), uint64(0), big.NewInt(0))
	baseMessage := types.NewBaseMessage(msgSender, recipientDestination, tips, msgRequest, body)
	message, err := types.NewMessageFromBaseMessage(header, baseMessage)
	Nil(u.T(), err)

	// Ensure the message is not executed yet.
	executed, err := exec.CheckIfExecuted(u.GetTestContext(), message)
	Nil(u.T(), err)

	False(u.T(), executed)

	// Send a message.
	tx, err := u.OriginContract.SendBaseMessage(
		txContextOrigin.TransactOpts,
		uint32(u.TestBackendDestination.GetChainID()),
		recipientDestination,
		optimisticSeconds,
		paddedRequest,
		body,
	)
	Nil(u.T(), err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		err = guard.Start(u.GetTestContext())
		if !testDone {
			u.Nil(err)
		}
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

	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig, omniRPCClient, u.NotaryTestDB, u.NotaryMetrics)
	Nil(u.T(), err)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

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

	// Check that it was executed.
	//nolint:dupl
	u.Eventually(func() bool {
		executed, err := exec.CheckIfExecuted(u.GetTestContext(), message)
		u.Nil(err)
		if executed {
			return true
		}

		// This transaction is needed to get the simulated chain's block number to increase by 1, since StreamLogs will
		// do lastBlockNumber - 1.
		tx, err = u.TestContractOnOrigin.EmitAgentsEventA(txContextOrigin.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		u.Nil(err)
		u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

		txContextDestination := u.TestBackendDestination.GetTxContext(u.GetTestContext(), u.DestinationContractMetadata.OwnerPtr())
		txContextSummit := u.TestBackendSummit.GetTxContext(u.GetTestContext(), u.InboxMetadataOnSummit.OwnerPtr())

		tx, err = u.TestContractOnSummit.EmitAgentsEventA(txContextSummit.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		u.Nil(err)
		u.TestBackendSummit.WaitForConfirmation(u.GetTestContext(), tx)
		tx, err = u.TestContractOnDestination.EmitAgentsEventA(txContextDestination.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		u.Nil(err)
		u.TestBackendDestination.WaitForConfirmation(u.GetTestContext(), tx)

		return false
	})

	// Send a second message.
	tx, err = u.OriginContract.SendBaseMessage(
		txContextOrigin.TransactOpts,
		uint32(u.TestBackendDestination.GetChainID()),
		recipientDestination,
		optimisticSeconds,
		paddedRequest,
		body,
	)

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
		return state.Nonce() >= uint32(2)
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
		return state.Nonce() >= uint32(2)
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
		return state.Nonce() >= uint32(2)
	})

	header = types.NewHeader(types.MessageFlagBase, uint32(u.TestBackendOrigin.GetChainID()), nonce+1, uint32(u.TestBackendDestination.GetChainID()), optimisticSeconds)
	message, err = types.NewMessageFromBaseMessage(header, baseMessage)

	// Check that it was executed.
	//nolint:dupl
	u.Eventually(func() bool {
		executed, err := exec.CheckIfExecuted(u.GetTestContext(), message)
		u.Nil(err)
		if executed {
			return true
		}

		// This transaction is needed to get the simulated chain's block number to increase by 1, since StreamLogs will
		// do lastBlockNumber - 1.
		tx, err = u.TestContractOnOrigin.EmitAgentsEventA(txContextOrigin.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		u.Nil(err)
		u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

		txContextDestination := u.TestBackendDestination.GetTxContext(u.GetTestContext(), u.DestinationContractMetadata.OwnerPtr())
		txContextSummit := u.TestBackendSummit.GetTxContext(u.GetTestContext(), u.InboxMetadataOnSummit.OwnerPtr())

		tx, err = u.TestContractOnSummit.EmitAgentsEventA(txContextSummit.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		u.Nil(err)
		u.TestBackendSummit.WaitForConfirmation(u.GetTestContext(), tx)
		tx, err = u.TestContractOnDestination.EmitAgentsEventA(txContextDestination.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		u.Nil(err)
		u.TestBackendDestination.WaitForConfirmation(u.GetTestContext(), tx)

		return false
	})
}
