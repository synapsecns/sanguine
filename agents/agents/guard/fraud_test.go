package guard_test

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/config"
	execConfig "github.com/synapsecns/sanguine/agents/config/executor"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/service"
)

func (g GuardSuite) getTestGuard(scribeConfig scribeConfig.Config) (testGuard *guard.Guard, sclient client.ScribeClient, err error) {
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
		MaxRetrySeconds:        60,
	}

	// Scribe setup.
	omniRPCClient := omniClient.NewOmnirpcClient(g.TestOmniRPC, g.GuardMetrics, omniClient.WithCaptureReqRes())
	originClient, err := backend.DialBackend(g.GetTestContext(), g.TestBackendOrigin.RPCAddress(), g.ScribeMetrics)
	Nil(g.T(), err)
	destinationClient, err := backend.DialBackend(g.GetTestContext(), g.TestBackendDestination.RPCAddress(), g.ScribeMetrics)
	Nil(g.T(), err)
	summitClient, err := backend.DialBackend(g.GetTestContext(), g.TestBackendSummit.RPCAddress(), g.ScribeMetrics)
	Nil(g.T(), err)

	clients := map[uint32][]backend.ScribeBackend{
		uint32(g.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(g.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
		uint32(g.TestBackendSummit.GetChainID()):      {summitClient, summitClient},
	}
	scribe, err := service.NewScribe(g.ScribeTestDB, clients, scribeConfig, g.ScribeMetrics)
	Nil(g.T(), err)
	scribeClient := client.NewEmbeddedScribe("sqlite", g.DBPath, g.ScribeMetrics)

	//nolint:errcheck
	go scribeClient.Start(g.GetTestContext())
	//nolint:errcheck
	go scribe.Start(g.GetTestContext())
	//nolint:wrapcheck
	testGuard, err = guard.NewGuard(g.GetTestContext(), testConfig, omniRPCClient, scribeClient.ScribeClient, g.GuardTestDB, g.GuardMetrics)
	sclient = scribeClient.ScribeClient
	if err != nil {
		return nil, sclient, fmt.Errorf("could not create guard: %w", err)
	}
	if testGuard == nil {
		return nil, sclient, fmt.Errorf("guard is nil")
	}

	return testGuard, sclient, nil
}

func (g GuardSuite) bumpBackends() {
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
	txContextOrigin := g.TestBackendOrigin.GetTxContext(g.GetTestContext(), g.OriginContractMetadata.OwnerPtr())
	txContextDestination := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.DestinationContractMetadata.OwnerPtr())
	g.bumpBackend(g.TestBackendSummit, g.TestContractOnSummit, txContextSummit.TransactOpts)
	g.bumpBackend(g.TestBackendOrigin, g.TestContractOnOrigin, txContextOrigin.TransactOpts)
	g.bumpBackend(g.TestBackendDestination, g.TestContractOnDestination, txContextDestination.TransactOpts)
}

// Helper to get the test backend to emit expected events.
func (g GuardSuite) bumpBackend(backend backends.SimulatedTestBackend, contract *agentstestcontract.AgentsTestContractRef, txOpts *bind.TransactOpts) {
	bumpTx, err := contract.EmitAgentsEventA(txOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	Nil(g.T(), err)
	backend.WaitForConfirmation(g.GetTestContext(), bumpTx)
}

func (g GuardSuite) updateAgentStatus(lightManager domains.LightManagerContract, bondedSigner, unbondedSigner signer.Signer, chainID uint32) {
	agentStatus, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), bondedSigner.Address())
	Nil(g.T(), err)
	agentProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), bondedSigner.Address())
	Nil(g.T(), err)
	transactor, err := unbondedSigner.GetTransactor(g.GetTestContext(), big.NewInt(int64(chainID)))
	Nil(g.T(), err)
	_, err = lightManager.UpdateAgentStatus(
		transactor,
		bondedSigner.Address(),
		agentStatus,
		agentProof,
	)
	Nil(g.T(), err)
	g.bumpBackends()
}

// TODO: Add a test for exiting the report logic early when the snapshot submitter is a guard.
func (g GuardSuite) TestFraudulentStateInSnapshot() {
	testDone := false
	defer func() {
		testDone = true
	}()

	originConfig := scribeConfig.ContractConfig{
		Address:    g.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendOrigin.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendDestination.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{destinationConfig},
	}
	summitConfig := scribeConfig.ContractConfig{
		Address:    g.InboxOnSummit.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendSummit.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{summitConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}

	// Start a new Guard.
	guard, _, err := g.getTestGuard(scribeConfig)
	Nil(g.T(), err)
	go func() {
		guardErr := guard.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), guardErr)
		}
	}()

	// Update the agent status on Origin.
	g.updateAgentStatus(g.OriginDomainClient.LightManager(), g.NotaryBondedSigner, g.NotaryUnbondedSigner, uint32(g.TestBackendOrigin.GetChainID()))

	// Verify that the agent is marked as Active.
	status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Equal(g.T(), status.Flag(), types.AgentFlagActive)
	Nil(g.T(), err)

	// Store agent trees and roots so that the agent status can be updated by the guard.
	agentRoot, err := g.SummitDomainClient.BondingManager().GetAgentRoot(g.GetTestContext())
	Nil(g.T(), err)
	blockNumber, err := g.SummitDomainClient.BlockNumber(g.GetTestContext())
	Nil(g.T(), err)
	notaryProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	err = g.GuardTestDB.StoreAgentTree(
		g.GetTestContext(),
		agentRoot,
		g.NotaryBondedSigner.Address(),
		uint64(blockNumber),
		notaryProof,
	)
	Nil(g.T(), err)
	err = g.GuardTestDB.StoreAgentRoot(g.GetTestContext(), agentRoot, uint64(blockNumber))
	Nil(g.T(), err)

	// Before submitting the attestation, ensure that there are no disputes opened.
	err = g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
	NotNil(g.T(), err)

	// Create a fraudulent snapshot.
	getState := func(nonce uint32) types.State {
		gasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())
		state := types.NewState(
			common.BigToHash(big.NewInt(gofakeit.Int64())),
			g.OriginDomainClient.Config().DomainID,
			nonce,
			big.NewInt(int64(gofakeit.Int32())),
			big.NewInt(int64(gofakeit.Int32())),
			gasData,
		)

		return state
	}
	fraudulentSnapshot := types.NewSnapshot([]types.State{getState(1), getState(2)})

	// Submit the snapshot with a guard then notary.
	guardSnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.InboxMetadataOnSummit.OwnerPtr())
	tx, err := g.SummitDomainClient.Inbox().SubmitSnapshot(txContextSummit.TransactOpts, encodedSnapshot, guardSnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	notarySnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.NotaryBondedSigner)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshot(txContextSummit.TransactOpts, encodedSnapshot, notarySnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Verify that the guard eventually marks the accused agent as Fraudulent on Origin.
	g.Eventually(func() bool {
		status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)

		if status.Flag() == types.AgentFlagFraudulent {
			return true
		}

		g.bumpBackends()
		return false
	})

	// Verify that a report has been submitted by the Guard by checking that a Dispute is now open.
	g.Eventually(func() bool {
		err = g.SummitDomainClient.BondingManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		return err == nil
	})

	// Verify that a state report was submitted on summit.
	fraudulentState := fraudulentSnapshot.States()[0]
	g.verifyStateReport(g.InboxOnSummit, 1, fraudulentState)

	// Verify that a state report was submitted on destination.
	g.verifyStateReport(g.LightInboxOnDestination, 1, fraudulentState)
}

func (g GuardSuite) TestFraudulentAttestationOnDestination() {
	testDone := false
	defer func() {
		testDone = true
	}()

	originConfig := scribeConfig.ContractConfig{
		Address:    g.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendOrigin.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendDestination.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{destinationConfig},
	}
	summitConfig := scribeConfig.ContractConfig{
		Address:    g.InboxOnSummit.Address().String(),
		StartBlock: 0,
	}
	bondingManagerConfig := scribeConfig.ContractConfig{
		Address:    g.BondingManagerOnSummit.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendSummit.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{summitConfig, bondingManagerConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}

	// Start a new Guard.
	guard, _, err := g.getTestGuard(scribeConfig)
	Nil(g.T(), err)
	go func() {
		guardErr := guard.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), guardErr)
		}
	}()

	_, gasDataContract := g.TestDeployManager.GetGasDataHarness(g.GetTestContext(), g.TestBackendDestination)
	_, attestationContract := g.TestDeployManager.GetAttestationHarness(g.GetTestContext(), g.TestBackendDestination)

	// Verify that the agent is marked as Active.
	txContextDest := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.DestinationContractMetadata.OwnerPtr())
	status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Equal(g.T(), status.Flag(), types.AgentFlagActive)
	Nil(g.T(), err)

	agentRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
	gasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())
	chainGas := types.NewChainGas(gasData, uint32(g.TestBackendOrigin.GetChainID()))
	chainGasBytes, err := types.EncodeChainGas(chainGas)
	Nil(g.T(), err)

	// Build and sign a fraudulent attestation.
	// TODO: Change from using a harness to using the Go code.
	snapGas := []*big.Int{new(big.Int).SetBytes(chainGasBytes)}
	snapGasHash, err := gasDataContract.SnapGasHash(&bind.CallOpts{Context: g.GetTestContext()}, snapGas)
	Nil(g.T(), err)
	dataHash, err := attestationContract.DataHash(&bind.CallOpts{Context: g.GetTestContext()}, agentRoot, snapGasHash)
	Nil(g.T(), err)
	attestationData := types.NewAttestation(
		common.BigToHash(big.NewInt(int64(gofakeit.Int32()))),
		dataHash,
		1,
		big.NewInt(int64(gofakeit.Int32())),
		big.NewInt(int64(gofakeit.Int32())),
	)
	attSignature, attEncoded, _, err := attestationData.SignAttestation(g.GetTestContext(), g.NotaryBondedSigner, true)
	Nil(g.T(), err)

	// Before submitting the attestation, ensure that there are no disputes opened.
	err = g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
	NotNil(g.T(), err)

	// Update the agent status of the Guard and Notary.
	g.updateAgentStatus(g.DestinationDomainClient.LightManager(), g.GuardBondedSigner, g.GuardUnbondedSigner, uint32(g.TestBackendDestination.GetChainID()))
	g.updateAgentStatus(g.DestinationDomainClient.LightManager(), g.NotaryBondedSigner, g.NotaryUnbondedSigner, uint32(g.TestBackendDestination.GetChainID()))

	// Submit the attestation.
	tx, err := g.DestinationDomainClient.LightInbox().SubmitAttestation(
		txContextDest.TransactOpts,
		attEncoded,
		attSignature,
		agentRoot,
		snapGas,
	)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendDestination.WaitForConfirmation(g.GetTestContext(), tx)

	// Verify that the guard eventually marks the accused agent as Slashed.
	g.Eventually(func() bool {
		status, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagSlashed {
			return true
		}

		// Make sure that scribe keeps producing new blocks.
		g.bumpBackends()
		return false
	})

	// Verify that a report has been submitted by the Guard by checking that a Dispute is now open.
	g.Eventually(func() bool {
		err := g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		return err == nil
	})
}

func (g GuardSuite) TestReportFraudulentStateInAttestation() {
	testDone := false
	defer func() {
		testDone = true
	}()

	// This scribe config omits the Summit and Origin chains, since we do not want to pick up the fraud coming from the
	// fraudulent snapshots, only from the Attestation submitted on the Destination that is associated with a fraudulent
	// snapshot.
	destinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendDestination.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{destinationConfig},
	}
	inboxConfig := scribeConfig.ContractConfig{
		Address:    g.InboxOnSummit.Address().String(),
		StartBlock: 0,
	}
	bondingManagerConfig := scribeConfig.ContractConfig{
		Address:    g.BondingManagerOnSummit.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendSummit.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{inboxConfig, bondingManagerConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{destinationChainConfig, summitChainConfig},
	}

	// Start a new Guard.
	guard, _, err := g.getTestGuard(scribeConfig)
	Nil(g.T(), err)
	go func() {
		guardErr := guard.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), guardErr)
		}
	}()

	// Verify that the agent is marked as Active.
	txContextDest := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.DestinationContractMetadata.OwnerPtr())
	status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Equal(g.T(), status.Flag(), types.AgentFlagActive)
	Nil(g.T(), err)

	// Create a fraudulent snapshot.
	gasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())
	fraudulentState := types.NewState(
		common.BigToHash(big.NewInt(gofakeit.Int64())),
		g.OriginDomainClient.Config().DomainID,
		1,
		big.NewInt(int64(gofakeit.Int32())),
		big.NewInt(int64(gofakeit.Int32())),
		gasData,
	)
	fraudulentSnapshot := types.NewSnapshot([]types.State{fraudulentState})

	// Before submitting the attestation, ensure that there are no disputes opened.
	err = g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
	NotNil(g.T(), err)

	// Update the agent status of the Guard and Notary.
	g.updateAgentStatus(g.DestinationDomainClient.LightManager(), g.GuardBondedSigner, g.GuardUnbondedSigner, uint32(g.TestBackendDestination.GetChainID()))
	g.updateAgentStatus(g.DestinationDomainClient.LightManager(), g.NotaryBondedSigner, g.NotaryUnbondedSigner, uint32(g.TestBackendDestination.GetChainID()))
	g.updateAgentStatus(g.OriginDomainClient.LightManager(), g.NotaryBondedSigner, g.NotaryUnbondedSigner, uint32(g.TestBackendOrigin.GetChainID()))

	// Submit the snapshot with a guard.
	guardSnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)
	transactOptsSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.InboxMetadataOnSummit.OwnerPtr())
	tx, err := g.SummitDomainClient.Inbox().SubmitSnapshot(transactOptsSummit.TransactOpts, encodedSnapshot, guardSnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Submit the snapshot with a notary.
	notarySnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.NotaryBondedSigner)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshot(transactOptsSummit.TransactOpts, encodedSnapshot, notarySnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Submit the attestation.
	notaryAttestation, err := g.SummitDomainClient.Summit().GetAttestation(g.GetTestContext(), 1)
	Nil(g.T(), err)
	attSignature, attEncoded, _, err := notaryAttestation.Attestation().SignAttestation(g.GetTestContext(), g.NotaryBondedSigner, true)
	Nil(g.T(), err)
	tx, err = g.DestinationDomainClient.LightInbox().SubmitAttestation(
		txContextDest.TransactOpts,
		attEncoded,
		attSignature,
		notaryAttestation.AgentRoot(),
		notaryAttestation.SnapGas(),
	)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendDestination.WaitForConfirmation(g.GetTestContext(), tx)

	// Verify that the guard eventually marks the accused agent as Fraudulent.
	g.Eventually(func() bool {
		status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagFraudulent {
			return true
		}

		g.bumpBackends()
		return false
	})

	// Verify that a dispute is now open on summit.
	g.Eventually(func() bool {
		err := g.SummitDomainClient.BondingManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		return err == nil
	})

	// Verify that a state report was submitted on summit.
	g.verifyStateReport(g.InboxOnSummit, 1, fraudulentState)

	// Verify that a state report was submitted on destination.
	g.verifyStateReport(g.LightInboxOnDestination, 1, fraudulentState)
}

type statementInboxContract interface {
	GetReportsAmount(opts *bind.CallOpts) (*big.Int, error)
	GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
		StatementPayload []byte
		ReportSignature  []byte
	}, error)
}

// Verify that a state report was submitted on the given contract.
//
//nolint:unparam
func (g GuardSuite) verifyStateReport(contract statementInboxContract, expectedNumReports int64, expectedState types.State) {
	g.Eventually(func() bool {
		numReports, err := contract.GetReportsAmount(&bind.CallOpts{Context: g.GetTestContext()})
		Nil(g.T(), err)

		if numReports.Int64() < expectedNumReports {
			return false
		}
		if numReports.Int64() != expectedNumReports {
			g.T().Fatalf("too many reports; expected %d, got %v", expectedNumReports, numReports.Int64())
		}

		stateReportIdx := big.NewInt(numReports.Int64() - 1)
		stateReport, err := contract.GetGuardReport(&bind.CallOpts{Context: g.GetTestContext()}, stateReportIdx)
		Nil(g.T(), err)

		expected, err := expectedState.Encode()
		Nil(g.T(), err)
		return Equal(g.T(), stateReport.StatementPayload, expected)
	})
}

func (g GuardSuite) TestInvalidReceipt() {
	testDone := false
	defer func() {
		testDone = true
	}()

	originConfig := scribeConfig.ContractConfig{
		Address:    g.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendOrigin.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendDestination.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{destinationConfig},
	}
	summitConfig := scribeConfig.ContractConfig{
		Address:    g.InboxOnSummit.Address().String(),
		StartBlock: 0,
	}
	bondingManagerConfig := scribeConfig.ContractConfig{
		Address:    g.BondingManagerOnSummit.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendSummit.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{summitConfig, bondingManagerConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}

	// Start a new Guard.
	guard, _, err := g.getTestGuard(scribeConfig)
	Nil(g.T(), err)
	go func() {
		guardErr := guard.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), guardErr)
		}
	}()

	// Send a message on Origin for it to be included in a valid state.
	summitTip := big.NewInt(int64(gofakeit.Uint32()))
	attestationTip := big.NewInt(int64(gofakeit.Uint32()))
	executorTip := big.NewInt(int64(gofakeit.Uint32()))
	deliveryTip := big.NewInt(int64(gofakeit.Uint32()))
	tips := types.NewTips(summitTip, attestationTip, executorTip, deliveryTip)
	optimisticSeconds := uint32(1)
	recipientDestination := g.TestClientMetadataOnDestination.Address().Hash()
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}
	txContextOrigin := g.TestBackendOrigin.GetTxContext(g.GetTestContext(), g.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)
	paddedRequest := big.NewInt(0)

	msgSender := common.BytesToHash(txContextOrigin.TransactOpts.From.Bytes())
	header := types.NewHeader(types.MessageFlagBase, uint32(g.TestBackendOrigin.GetChainID()), nonce, uint32(g.TestBackendDestination.GetChainID()), optimisticSeconds)
	msgRequest := types.NewRequest(uint32(0), uint64(0), big.NewInt(0))
	baseMessage := types.NewBaseMessage(msgSender, recipientDestination, tips, msgRequest, body)
	message, err := types.NewMessageFromBaseMessage(header, baseMessage)
	Nil(g.T(), err)

	tx, err := g.OriginContract.SendBaseMessage(
		txContextOrigin.TransactOpts,
		uint32(g.TestBackendDestination.GetChainID()),
		recipientDestination,
		optimisticSeconds,
		paddedRequest,
		body,
	)
	Nil(g.T(), err)
	g.TestBackendOrigin.WaitForConfirmation(g.GetTestContext(), tx)

	// Submit the snapshot with a guard.
	latestOriginState, err := g.OriginDomainClient.Origin().SuggestLatestState(g.GetTestContext())
	Nil(g.T(), err)
	snapshot := types.NewSnapshot([]types.State{latestOriginState})
	guardSnapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.InboxMetadataOnSummit.OwnerPtr())
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshot(txContextSummit.TransactOpts, encodedSnapshot, guardSnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Submit the snapshot with a notary.
	notarySnapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(g.GetTestContext(), g.NotaryBondedSigner)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshot(txContextSummit.TransactOpts, encodedSnapshot, notarySnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Build and sign a receipt.
	snapshotRoot, _, err := snapshot.SnapshotRootAndProofs()
	Nil(g.T(), err)
	g.updateAgentStatus(g.DestinationDomainClient.LightManager(), g.NotaryBondedSigner, g.NotaryUnbondedSigner, uint32(g.TestBackendDestination.GetChainID()))
	messageHash, err := message.ToLeaf()
	Nil(g.T(), err)
	receipt := types.NewReceipt(
		g.OriginDomainClient.Config().DomainID,
		g.DestinationDomainClient.Config().DomainID,
		messageHash,
		snapshotRoot,
		0,
		g.NotaryBondedWallet.Address(),
		common.BigToAddress(big.NewInt(gofakeit.Int64())),
		common.BigToAddress(big.NewInt(gofakeit.Int64())),
	)
	rcptSignature, rcptPayload, _, err := receipt.SignReceipt(g.GetTestContext(), g.NotaryBondedSigner, true)
	Nil(g.T(), err)

	// Submit the receipt.
	bodyHash, err := baseMessage.BodyLeaf()
	Nil(g.T(), err)
	var bodyHashB32 [32]byte
	copy(bodyHashB32[:], bodyHash)
	headerHash, err := header.Leaf()
	Nil(g.T(), err)
	paddedTips, err := types.EncodeTipsBigInt(tips)
	Nil(g.T(), err)

	txContextSummit = g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
	tx, err = g.SummitDomainClient.Inbox().SubmitReceipt(
		txContextSummit.TransactOpts,
		rcptPayload,
		rcptSignature,
		paddedTips,
		headerHash,
		bodyHashB32,
	)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Verify that the guard eventually marks the accused agent as Fraudulent.
	g.Eventually(func() bool {
		status, err := g.DestinationDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagFraudulent {
			return true
		}

		g.bumpBackends()
		return false
	})

	// TODO: Uncomment once updating agent status is implemented.
	// g.Eventually(func() bool {
	// 	status, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	// 	Nil(g.T(), err)
	// 	if status.Flag() == types.AgentFlagSlashed {
	// 		return true
	// 	}

	//  g.bumpBackends()
	// 	return false
	// })

	// Verify that a report has been submitted by the Guard by checking that a Dispute is now open.
	g.Eventually(func() bool {
		err := g.SummitDomainClient.BondingManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		return err == nil
	})
}

//nolint:maintidx,cyclop
func (g GuardSuite) TestUpdateAgentStatusOnRemote() {
	// This test requires a call to anvil's evm.IncreaseTime() cheat code, so we should
	// set up the backends with anvil.

	testDone := false
	defer func() {
		testDone = true
	}()

	// This scribe config omits the Summit and Origin chains, since we do not want to pick up the fraud coming from the
	// fraudulent snapshots, only from the Attestation submitted on the Destination that is associated with a fraudulent
	// snapshot.
	originConfig := scribeConfig.ContractConfig{
		Address:    g.OriginContract.Address().String(),
		StartBlock: 0,
	}
	lightManagerConfig := scribeConfig.ContractConfig{
		Address:    g.LightManagerOnOrigin.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig.ChainConfig{
		ChainID:            uint32(g.TestBackendOrigin.GetChainID()),
		GetLogsBatchAmount: 1,
		StoreConcurrency:   1,
		GetLogsRange:       1,
		Confirmations:      1,
		Contracts:          []scribeConfig.ContractConfig{originConfig, lightManagerConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    g.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	lightInboxDestinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	lightManagerDestinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightManagerOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendDestination.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{destinationConfig, lightInboxDestinationConfig, lightManagerDestinationConfig},
	}
	summitConfig := scribeConfig.ContractConfig{
		Address:    g.SummitContract.Address().String(),
		StartBlock: 0,
	}
	inboxConfig := scribeConfig.ContractConfig{
		Address:    g.InboxOnSummit.Address().String(),
		StartBlock: 0,
	}
	bondingManagerConfig := scribeConfig.ContractConfig{
		Address:    g.BondingManagerOnSummit.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := scribeConfig.ChainConfig{
		ChainID:       uint32(g.TestBackendSummit.GetChainID()),
		Confirmations: 1,
		Contracts:     []scribeConfig.ContractConfig{summitConfig, inboxConfig, bondingManagerConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}

	// Start a new Guard.
	guard, scribeClient, err := g.getTestGuard(scribeConfig)
	Nil(g.T(), err)
	go func() {
		guardErr := guard.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), guardErr)
		}
	}()

	// Scribe setup.
	omniRPCClient := omniClient.NewOmnirpcClient(g.TestOmniRPC, g.GuardMetrics, omniClient.WithCaptureReqRes())
	chainID := uint32(g.TestBackendOrigin.GetChainID())
	destination := uint32(g.TestBackendDestination.GetChainID())
	summit := uint32(g.TestBackendSummit.GetChainID())
	fmt.Printf("omnirpc: %v", omniRPCClient.GetDefaultEndpoint(int(summit)))

	excCfg := execConfig.Config{
		SummitChainID: summit,
		SummitAddress: g.SummitContract.Address().String(),
		InboxAddress:  g.InboxOnSummit.Address().String(),
		Chains: []execConfig.ChainConfig{
			{
				ChainID:       chainID,
				OriginAddress: g.OriginContract.Address().String(),
			},
			{
				ChainID:            destination,
				LightInboxAddress:  g.LightInboxOnDestination.Address().String(),
				DestinationAddress: g.DestinationContract.Address().String(),
			},
			{
				ChainID:            summit,
				DestinationAddress: g.DestinationContractOnSummit.Address().String(),
			},
		},
		BaseOmnirpcURL: g.TestBackendOrigin.RPCAddress(),
		UnbondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(g.T(), "", g.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
		SubmitterConfig: submitterConfig.Config{
			ChainConfig: submitterConfig.ChainConfig{
				GasEstimate: uint64(5000000),
			},
		},
	}

	// This function will allow us to override the current time perceived by Executor.
	var currentTime *time.Time
	nowFunc := func() time.Time {
		if currentTime == nil {
			return time.Now()
		}
		return *currentTime
	}
	getChainTimeFunc := func(ctx context.Context, backend executor.Backend) (uint64, error) {
		return uint64(nowFunc().Unix()), nil
	}

	// Start a new Executor.
	exec, err := executor.NewExecutor(g.GetTestContext(), excCfg, g.ExecutorTestDB, scribeClient, omniRPCClient, g.ExecutorMetrics)
	Nil(g.T(), err)
	exec.NowFunc = nowFunc
	exec.GetChainTimeFunc = getChainTimeFunc

	go func() {
		execErr := exec.Run(g.GetTestContext())
		if !testDone {
			Nil(g.T(), execErr)
		}
	}()

	// Verify that the agent is marked as Active.
	txContextDest := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.DestinationContractMetadata.OwnerPtr())
	status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Equal(g.T(), status.Flag(), types.AgentFlagActive)
	Nil(g.T(), err)

	// Create a fraudulent snapshot.
	gasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())
	fraudulentState := types.NewState(
		common.BigToHash(big.NewInt(gofakeit.Int64())),
		g.OriginDomainClient.Config().DomainID,
		1,
		big.NewInt(int64(gofakeit.Int32())),
		big.NewInt(int64(gofakeit.Int32())),
		gasData,
	)
	fraudulentSnapshot := types.NewSnapshot([]types.State{fraudulentState})

	// Before submitting the attestation, ensure that there are no disputes opened.
	err = g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
	NotNil(g.T(), err)

	// Update the agent status of the Guard and Notaries.
	g.updateAgentStatus(g.DestinationDomainClient.LightManager(), g.GuardBondedSigner, g.GuardUnbondedSigner, destination)
	g.updateAgentStatus(g.DestinationDomainClient.LightManager(), g.NotaryBondedSigner, g.NotaryUnbondedSigner, destination)
	g.updateAgentStatus(g.OriginDomainClient.LightManager(), g.NotaryBondedSigner, g.NotaryUnbondedSigner, uint32(g.TestBackendOrigin.GetChainID()))
	g.updateAgentStatus(g.DestinationDomainClient.LightManager(), g.NotaryOnDestinationBondedSigner, g.NotaryOnDestinationUnbondedSigner, destination)
	g.updateAgentStatus(g.OriginDomainClient.LightManager(), g.NotaryOnDestinationBondedSigner, g.NotaryOnDestinationUnbondedSigner, uint32(g.TestBackendOrigin.GetChainID()))

	fmt.Printf("BondedNotaryAddress: %v\n", g.NotaryBondedSigner.Address())

	// Submit the snapshot with a guard.
	guardSnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.InboxMetadataOnSummit.OwnerPtr())
	tx, err := g.SummitDomainClient.Inbox().SubmitSnapshot(txContextSummit.TransactOpts, encodedSnapshot, guardSnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)
	g.bumpBackends()

	// Submit the snapshot with a notary.
	notarySnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.NotaryBondedSigner)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshot(txContextSummit.TransactOpts, encodedSnapshot, notarySnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)
	g.bumpBackends()

	// Submit the attestation.
	notaryAttestation, err := g.SummitDomainClient.Summit().GetAttestation(g.GetTestContext(), 1)
	Nil(g.T(), err)
	attSignature, attEncoded, _, err := notaryAttestation.Attestation().SignAttestation(g.GetTestContext(), g.NotaryBondedSigner, true)
	Nil(g.T(), err)
	tx, err = g.DestinationDomainClient.LightInbox().SubmitAttestation(
		txContextDest.TransactOpts,
		attEncoded,
		attSignature,
		notaryAttestation.AgentRoot(),
		notaryAttestation.SnapGas(),
	)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendDestination.WaitForConfirmation(g.GetTestContext(), tx)
	g.bumpBackends()

	// Verify that the guard eventually marks the accused agent as Fraudulent.
	g.Eventually(func() bool {
		status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagFraudulent {
			return true
		}

		g.bumpBackends()
		return false
	})

	// Verify that a report has been submitted by the Guard by checking that a Dispute is now open.
	g.Eventually(func() bool {
		err := g.SummitDomainClient.BondingManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		return err == nil
	})

	// Get the origin state so we can submit it on the Summit.
	originStateRaw, err := g.OriginContract.SuggestLatestState(&bind.CallOpts{Context: g.GetTestContext()})
	g.Nil(err)
	originState, err := types.DecodeState(originStateRaw)
	g.Nil(err)
	snapshot := types.NewSnapshot([]types.State{originState})

	//nolint:wrapcheck
	submitAndVerifySnapshot := func(originState types.State, agentSigner signer.Signer) error {
		agentSnapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(g.GetTestContext(), agentSigner)
		if err != nil {
			return err
		}
		txContextSummit = g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
		tx, err = g.SummitDomainClient.Inbox().SubmitSnapshot(
			txContextSummit.TransactOpts,
			encodedSnapshot,
			agentSnapshotSignature,
		)
		if err != nil {
			return err
		}
		g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)
		g.bumpBackends()

		latestStateRaw, err := g.SummitContract.GetLatestAgentState(&bind.CallOpts{Context: g.GetTestContext()}, uint32(g.TestBackendOrigin.GetChainID()), agentSigner.Address())
		if err != nil {
			return err
		}
		latestState, err := types.DecodeState(latestStateRaw)
		if err != nil {
			return err
		}
		latestStateHash, err := latestState.Hash()
		if err != nil {
			return err
		}
		originStateHash, err := originState.Hash()
		if err != nil {
			return err
		}
		latestStateHashHex := common.BytesToHash(latestStateHash[:])
		originStateHashHex := common.BytesToHash(originStateHash[:])
		if latestStateHash != originStateHash {
			return fmt.Errorf("latest state hash mismatch; expected %v, got %v", originStateHashHex, latestStateHashHex)
		}
		return nil
	}

	// Submit snapshot with Guard.
	g.Eventually(func() bool {
		err := submitAndVerifySnapshot(originState, g.GuardBondedSigner)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	})
	time.Sleep(5 * time.Second)

	// Submit snapshot with Notary.
	g.Eventually(func() bool {
		err := submitAndVerifySnapshot(originState, g.NotaryOnOriginBondedSigner)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	})

	// Wait for the executor to have attestations before increasing time.
	summitChainID := uint32(g.TestBackendSummit.GetChainID())
	attestationNonce := uint32(2)
	g.Eventually(func() bool {
		attest, err := g.ExecutorTestDB.GetAttestation(g.GetTestContext(), db.DBAttestation{
			Destination:      &summitChainID,
			AttestationNonce: &attestationNonce,
		})
		Nil(g.T(), err)
		return attest != nil
	})

	// Increase EVM time to allow agent status to be updated to Slashed on summit.
	optimisticPeriodSeconds := int64(86400)
	offset := optimisticPeriodSeconds / 2
	increaseEvmTime := func(backend backends.SimulatedTestBackend, seconds int64) {
		anvilClient, err := anvil.Dial(g.GetTestContext(), backend.RPCAddress())
		Nil(g.T(), err)
		client, err := omniRPCClient.GetClient(g.GetTestContext(), g.TestBackendSummit.GetBigChainID())
		Nil(g.T(), err)
		headerPreIncrease, err := client.HeaderByNumber(g.GetTestContext(), nil)
		Nil(g.T(), err)
		// targetTimestamp := int64(headerPreIncrease.Time) + seconds
		err = anvilClient.IncreaseTime(g.GetTestContext(), seconds)
		targetTimestamp := 0
		// err = anvilClient.SetBlockTimestampInterval(g.GetTestContext(), seconds)
		Nil(g.T(), err)
		err = anvilClient.Mine(g.GetTestContext(), 1)
		Nil(g.T(), err)
		headerPostIncrease, err := client.HeaderByNumber(g.GetTestContext(), nil)
		Nil(g.T(), err)
		fmt.Printf("header time moved from (%d, block=%s) to (%d, block=%s) [target=%d]\n", headerPreIncrease.Time, headerPreIncrease.Number.String(), headerPostIncrease.Time, headerPostIncrease.Number.String(), targetTimestamp)
	}
	increaseEvmTime(g.TestBackendSummit, optimisticPeriodSeconds+offset)
	g.bumpBackends()

	// Increase executor time so that the manager message may be executed.
	updatedTime := time.Now().Add(time.Duration(optimisticPeriodSeconds+offset) * time.Second)
	currentTime = &updatedTime
	fmt.Printf("updated executor time from %d to %d\n", time.Now().Unix(), updatedTime.Unix())

	// Verify that the accused agent is eventually Slashed on Summit.
	g.Eventually(func() bool {
		status, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagSlashed {
			return true
		}
		time.Sleep(5 * time.Second)
		g.bumpBackends()
		return false
	})

	// Get the origin state so we can submit it on the Summit.
	originStateRaw, err = g.OriginContract.SuggestLatestState(&bind.CallOpts{Context: g.GetTestContext()})
	g.Nil(err)
	originState, err = types.DecodeState(originStateRaw)
	g.Nil(err)
	snapshot = types.NewSnapshot([]types.State{originState})

	// Submit snapshot with Guard.
	guardSnapshotSignature, encodedSnapshot, _, err = snapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	g.Nil(err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshot(
		txContextSummit.TransactOpts,
		encodedSnapshot,
		guardSnapshotSignature,
	)
	g.Nil(err)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)
	g.bumpBackends()

	// Submit snapshot with Notary.
	notarySnapshotSignature, encodedSnapshot, _, err = snapshot.SignSnapshot(g.GetTestContext(), g.NotaryOnOriginBondedSigner)
	g.Nil(err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshot(
		txContextSummit.TransactOpts,
		encodedSnapshot,
		notarySnapshotSignature,
	)
	g.Nil(err)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)
	g.bumpBackends()

	// Create a new attestation with the agent root corresponding to newly Slashed status.
	latestAgentRoot, err := g.SummitDomainClient.BondingManager().GetAgentRoot(g.GetTestContext())
	Nil(g.T(), err)
	_, gasDataContract := g.TestDeployManager.GetGasDataHarness(g.GetTestContext(), g.TestBackendDestination)
	_, attestationContract := g.TestDeployManager.GetAttestationHarness(g.GetTestContext(), g.TestBackendDestination)
	chainGas := types.NewChainGas(originState.GasData(), uint32(g.TestBackendOrigin.GetChainID()))
	chainGasBytes, err := types.EncodeChainGas(chainGas)
	Nil(g.T(), err)
	// TODO: Change from using a harness to using the Go code.
	snapGas := []*big.Int{new(big.Int).SetBytes(chainGasBytes)}
	snapGasHash, err := gasDataContract.SnapGasHash(&bind.CallOpts{Context: g.GetTestContext()}, snapGas)
	Nil(g.T(), err)
	dataHash, err := attestationContract.DataHash(&bind.CallOpts{Context: g.GetTestContext()}, latestAgentRoot, snapGasHash)
	Nil(g.T(), err)
	notaryAttestation, err = g.SummitDomainClient.Summit().GetAttestation(g.GetTestContext(), 2)
	Nil(g.T(), err)
	attestation := types.NewAttestation(
		notaryAttestation.Attestation().SnapshotRoot(),
		dataHash,
		2,
		notaryAttestation.Attestation().BlockNumber(),
		notaryAttestation.Attestation().Timestamp(),
	)
	attEncoded, err = attestation.Encode()
	Nil(g.T(), err)
	notaryAttestation, err = types.NewNotaryAttestation(attEncoded, latestAgentRoot, snapGas)
	Nil(g.T(), err)

	// Submit the attestation.
	attSignature, attEncoded, _, err = attestation.SignAttestation(g.GetTestContext(), g.NotaryOnDestinationBondedSigner, true)
	Nil(g.T(), err)
	tx, err = g.DestinationDomainClient.LightInbox().SubmitAttestation(
		txContextDest.TransactOpts,
		attEncoded,
		attSignature,
		latestAgentRoot,
		notaryAttestation.SnapGas(),
	)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendDestination.WaitForConfirmation(g.GetTestContext(), tx)

	// Advance time on destination and call passAgentRoot() so that the latest agent root is accepted.
	increaseEvmTime(g.TestBackendDestination, optimisticPeriodSeconds+offset)
	g.bumpBackends()
	txContextDestination := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.DestinationContractMetadata.OwnerPtr())
	tx, err = g.DestinationDomainClient.Destination().PassAgentRoot(txContextDestination.TransactOpts)
	g.Nil(err)
	g.TestBackendDestination.WaitForConfirmation(g.GetTestContext(), tx)
	g.bumpBackends()

	// Verify that the guard eventually marks the accused agent as Slashed.
	g.Eventually(func() bool {
		status, err := g.DestinationDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagSlashed {
			return true
		}

		g.bumpBackends()
		return false
	})
}
