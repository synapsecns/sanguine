package guard_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
)

func (g GuardSuite) getTestGuard(scribeConfig scribeConfig.Config) (*guard.Guard, error) {
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

	// Scribe setup.
	omniRPCClient := omniClient.NewOmnirpcClient(g.TestOmniRPC, g.GuardMetrics, omniClient.WithCaptureReqRes())
	originClient, err := backfill.DialBackend(g.GetTestContext(), g.TestBackendOrigin.RPCAddress(), g.ScribeMetrics)
	Nil(g.T(), err)
	destinationClient, err := backfill.DialBackend(g.GetTestContext(), g.TestBackendDestination.RPCAddress(), g.ScribeMetrics)
	Nil(g.T(), err)
	summitClient, err := backfill.DialBackend(g.GetTestContext(), g.TestBackendSummit.RPCAddress(), g.ScribeMetrics)
	Nil(g.T(), err)

	clients := map[uint32][]backfill.ScribeBackend{
		uint32(g.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(g.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
		uint32(g.TestBackendSummit.GetChainID()):      {summitClient, summitClient},
	}

	scribe, err := node.NewScribe(g.ScribeTestDB, clients, scribeConfig, g.ScribeMetrics)
	Nil(g.T(), err)
	scribeClient := client.NewEmbeddedScribe("sqlite", g.DBPath, g.ScribeMetrics)
	go scribeClient.Start(g.GetTestContext())
	go scribe.Start(g.GetTestContext())

	return guard.NewGuard(g.GetTestContext(), testConfig, omniRPCClient, scribeClient.ScribeClient, g.GuardTestDB, g.GuardMetrics)
}

// Helper to get the test backend to emit expected events.
func (g GuardSuite) bumpBackend(backend backends.SimulatedTestBackend, contract *agentstestcontract.AgentsTestContractRef, txOpts *bind.TransactOpts) {
	bumpTx, err := contract.EmitAgentsEventA(txOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	Nil(g.T(), err)
	backend.WaitForConfirmation(g.GetTestContext(), bumpTx)
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
		ChainID:               uint32(g.TestBackendOrigin.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(g.TestBackendDestination.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{destinationConfig},
	}
	summitConfig := scribeConfig.ContractConfig{
		Address:    g.InboxOnSummit.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(g.TestBackendSummit.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{summitConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}

	// Start a new Guard.
	guard, err := g.getTestGuard(scribeConfig)
	Nil(g.T(), err)
	go func() {
		guardErr := guard.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), guardErr)
		}
	}()

	// Update the agent status on Origin.
	notaryStatus, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	notaryProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	err = g.OriginDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.NotaryUnbondedSigner,
		g.NotaryBondedSigner,
		notaryStatus,
		notaryProof,
	)
	Nil(g.T(), err)

	// Verify that the agent is marked as Active
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
	status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Equal(g.T(), status.Flag(), types.AgentFlagActive)
	Nil(g.T(), err)

	// Before submitting the attestation, ensure that there are no disputes opened.
	err = g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
	NotNil(g.T(), err)

	// Create a fraudulent snapshot
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

	// Submit the snapshot with a guard then notary
	guardSnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)
	tx, err := g.SummitDomainClient.Inbox().SubmitSnapshotCtx(g.GetTestContext(), g.GuardUnbondedSigner, encodedSnapshot, guardSnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	notarySnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.NotaryBondedSigner)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshotCtx(g.GetTestContext(), g.NotaryUnbondedSigner, encodedSnapshot, notarySnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Verify that the guard eventually marks the accused agent as Fraudulent on Origin
	g.Eventually(func() bool {
		status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)

		if status.Flag() == types.AgentFlagFraudulent {
			return true
		}

		g.bumpBackend(g.TestBackendSummit, g.TestContractOnSummit, txContextSummit.TransactOpts)
		return false
	})

	// TODO: Once we add updating agent statuses fully, uncomment this.
	// Verify that the guard eventually marks the accused agent as Fraudulent on Summit
	//g.Eventually(func() bool {
	//	status, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	//	Nil(g.T(), err)
	//
	//	if status.Flag() == types.AgentFlagSlashed {
	//		return true
	//	}
	//
	//	g.bumpBackend(g.TestBackendSummit, g.TestContractOnSummit, txContextSummit.TransactOpts)
	//	return false
	//})

	// Verify that a report has been submitted by the Guard by checking that a Dispute is now open.
	g.Eventually(func() bool {
		err = g.SummitDomainClient.BondingManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		if err != nil {
			return false
		}

		return true
	})
	// TODO: Add a unit test for testing the case where multiple states are in the same snapshot to ensure they are
	// handled correctly.
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
		ChainID:               uint32(g.TestBackendOrigin.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(g.TestBackendDestination.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{destinationConfig},
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
		ChainID:               uint32(g.TestBackendSummit.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{summitConfig, bondingManagerConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}

	// Start a new Guard.
	guard, err := g.getTestGuard(scribeConfig)
	Nil(g.T(), err)
	go func() {
		guardErr := guard.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), guardErr)
		}
	}()

	_, gasDataContract := g.TestDeployManager.GetGasDataHarness(g.GetTestContext(), g.TestBackendDestination)
	_, attestationContract := g.TestDeployManager.GetAttestationHarness(g.GetTestContext(), g.TestBackendDestination)

	// Verify that the agent is marked as Active
	txContextDest := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.DestinationContractMetadata.OwnerPtr())
	status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Equal(g.T(), status.Flag(), types.AgentFlagActive)
	Nil(g.T(), err)

	agentRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
	gasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())
	chainGas := types.NewChainGas(gasData, uint32(g.TestBackendOrigin.GetChainID()))
	chainGasBytes, err := types.EncodeChainGas(chainGas)
	Nil(g.T(), err)

	// Build and sign a fraudulent attestation
	// TODO: Change from using a harness to using the Go code.
	snapGas := []*big.Int{new(big.Int).SetBytes(chainGasBytes)}
	snapGasHash, err := gasDataContract.SnapGasHash(&bind.CallOpts{Context: g.GetTestContext()}, snapGas)
	Nil(g.T(), err)
	dataHash, err := attestationContract.DataHash(&bind.CallOpts{Context: g.GetTestContext()}, agentRoot, snapGasHash)
	Nil(g.T(), err)
	fraudAttestation := types.NewAttestation(
		common.BigToHash(big.NewInt(int64(gofakeit.Int32()))),
		dataHash,
		1,
		big.NewInt(int64(gofakeit.Int32())),
		big.NewInt(int64(gofakeit.Int32())),
	)
	attSignature, attEncoded, _, err := fraudAttestation.SignAttestation(g.GetTestContext(), g.NotaryBondedSigner, true)
	Nil(g.T(), err)

	// Before submitting the attestation, ensure that there are no disputes opened.
	err = g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
	NotNil(g.T(), err)

	// Update the agent status of the Guard and Notary.
	guardStatus, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Nil(g.T(), err)
	guardProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.GuardBondedSigner.Address())
	Nil(g.T(), err)
	err = g.DestinationDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.GuardUnbondedSigner,
		g.GuardBondedSigner,
		guardStatus,
		guardProof,
	)
	Nil(g.T(), err)

	notaryStatus, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	notaryProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	err = g.DestinationDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.NotaryUnbondedSigner,
		g.NotaryBondedSigner,
		notaryStatus,
		notaryProof,
	)
	Nil(g.T(), err)

	// Submit the attestation
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

	// Verify that the guard eventually marks the accused agent as Slashed
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
	txContextDestination := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.LightInboxMetadataOnDestination.OwnerPtr())
	g.Eventually(func() bool {
		status, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagSlashed {
			return true
		}

		// Make sure that scribe keeps producing new blocks
		g.bumpBackend(g.TestBackendSummit, g.TestContractOnSummit, txContextSummit.TransactOpts)
		g.bumpBackend(g.TestBackendDestination, g.TestContractOnDestination, txContextDestination.TransactOpts)
		return false
	})

	// Verify that a report has been submitted by the Guard by checking that a Dispute is now open.
	g.Eventually(func() bool {
		err := g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		if err != nil {
			return false
		}

		return true
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
		ChainID:               uint32(g.TestBackendDestination.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{destinationConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{destinationChainConfig},
	}

	// Start a new Guard.
	guard, err := g.getTestGuard(scribeConfig)
	Nil(g.T(), err)
	go func() {
		guardErr := guard.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), guardErr)
		}
	}()

	// Verify that the agent is marked as Active
	txContextDest := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.DestinationContractMetadata.OwnerPtr())
	status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Equal(g.T(), status.Flag(), types.AgentFlagActive)
	Nil(g.T(), err)

	// Create a fraudulent snapshot
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
	guardStatus, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Nil(g.T(), err)
	guardProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.GuardBondedSigner.Address())
	Nil(g.T(), err)
	err = g.DestinationDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.GuardUnbondedSigner,
		g.GuardBondedSigner,
		guardStatus,
		guardProof,
	)
	Nil(g.T(), err)

	notaryStatus, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	notaryProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	err = g.DestinationDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.NotaryUnbondedSigner,
		g.NotaryBondedSigner,
		notaryStatus,
		notaryProof,
	)
	Nil(g.T(), err)
	err = g.OriginDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.NotaryUnbondedSigner,
		g.NotaryBondedSigner,
		notaryStatus,
		notaryProof,
	)
	Nil(g.T(), err)

	// Submit the snapshot with a guard
	guardSnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)
	tx, err := g.SummitDomainClient.Inbox().SubmitSnapshotCtx(g.GetTestContext(), g.GuardUnbondedSigner, encodedSnapshot, guardSnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Submit the snapshot with a notary
	notarySnapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.NotaryBondedSigner)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshotCtx(g.GetTestContext(), g.NotaryUnbondedSigner, encodedSnapshot, notarySnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Submit the attestation
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

	// Verify that the guard eventually marks the accused agent as Slashed
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
	txContextDestination := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.LightInboxMetadataOnDestination.OwnerPtr())
	g.Eventually(func() bool {
		status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagFraudulent {
			return true
		}

		g.bumpBackend(g.TestBackendSummit, g.TestContractOnSummit, txContextSummit.TransactOpts)
		g.bumpBackend(g.TestBackendDestination, g.TestContractOnDestination, txContextDestination.TransactOpts)
		return false
	})

	// Verify that a report has been submitted by the Guard by checking that a Dispute is now open.
	g.Eventually(func() bool {
		err := g.SummitDomainClient.BondingManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		if err != nil {
			return false
		}

		return true
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
		ChainID:               uint32(g.TestBackendOrigin.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    g.LightInboxOnDestination.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(g.TestBackendDestination.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{destinationConfig},
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
		ChainID:               uint32(g.TestBackendSummit.GetChainID()),
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{summitConfig, bondingManagerConfig},
	}
	scribeConfig := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}

	// Start a new Guard.
	guard, err := g.getTestGuard(scribeConfig)
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

	// Submit the snapshot with a guard
	latestOriginState, err := g.OriginDomainClient.Origin().SuggestLatestState(g.GetTestContext())
	Nil(g.T(), err)
	snapshot := types.NewSnapshot([]types.State{latestOriginState})
	guardSnapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshotCtx(g.GetTestContext(), g.GuardUnbondedSigner, encodedSnapshot, guardSnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Submit the snapshot with a notary
	notarySnapshotSignature, encodedSnapshot, _, err := snapshot.SignSnapshot(g.GetTestContext(), g.NotaryBondedSigner)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitSnapshotCtx(g.GetTestContext(), g.NotaryUnbondedSigner, encodedSnapshot, notarySnapshotSignature)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Build and sign a receipt
	snapshotRoot, _, err := snapshot.SnapshotRootAndProofs()
	Nil(g.T(), err)
	notaryStatus, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	notaryProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	err = g.DestinationDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.NotaryUnbondedSigner,
		g.NotaryBondedSigner,
		notaryStatus,
		notaryProof,
	)
	Nil(g.T(), err)
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

	// Submit the receipt
	bodyHash, err := baseMessage.BodyLeaf()
	var bodyHashB32 [32]byte
	copy(bodyHashB32[:], bodyHash[:])
	headerHash, err := header.Leaf()
	Nil(g.T(), err)
	paddedTips, err := types.EncodeTipsBigInt(tips)
	Nil(g.T(), err)
	tx, err = g.SummitDomainClient.Inbox().SubmitReceipt(
		g.GetTestContext(),
		g.NotaryUnbondedSigner,
		rcptPayload,
		rcptSignature,
		paddedTips,
		headerHash,
		bodyHashB32,
	)
	Nil(g.T(), err)
	NotNil(g.T(), tx)
	fmt.Println("submitReceiptTx: ", tx.Hash().String())
	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	// Verify that the guard eventually marks the accused agent as Fraudulent
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
	txContextDestination := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.LightInboxMetadataOnDestination.OwnerPtr())
	g.Eventually(func() bool {
		status, err := g.DestinationDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == types.AgentFlagFraudulent {
			return true
		}

		g.bumpBackend(g.TestBackendSummit, g.TestContractOnSummit, txContextSummit.TransactOpts)
		g.bumpBackend(g.TestBackendDestination, g.TestContractOnDestination, txContextDestination.TransactOpts)
		return false
	})

	// TODO: Uncomment once updating agent status is implemented.
	//g.Eventually(func() bool {
	//	status, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	//	Nil(g.T(), err)
	//	if status.Flag() == types.AgentFlagSlashed {
	//		return true
	//	}
	//
	//	bumpTx, err := g.TestContractOnSummit.EmitAgentsEventA(txContextSummit.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	//	Nil(g.T(), err)
	//	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), bumpTx)
	//	return false
	//})

	// Verify that a report has been submitted by the Guard by checking that a Dispute is now open.
	g.Eventually(func() bool {
		err := g.SummitDomainClient.BondingManager().GetDispute(g.GetTestContext(), big.NewInt(0))
		if err != nil {
			return false
		}

		return true
	})
}
