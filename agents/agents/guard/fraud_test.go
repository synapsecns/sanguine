package guard_test

import (
	"fmt"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/types"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
)

// TODO: Add a test for exiting the report logic early when the snapshot submitter is a guard.
func (g GuardSuite) TestReportFraudulentStateInSnapshot() {
	testDone := false
	defer func() {
		testDone = true
	}()

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

	omniRPCClient := omniClient.NewOmnirpcClient(g.TestOmniRPC, g.GuardMetrics, omniClient.WithCaptureReqRes())

	omnirpcDest := omniRPCClient.GetEndpoint(int(g.SummitDomainClient.Config().DomainID), 1)

	fmt.Println("OMNIIIIII", omnirpcDest)

	// Scribe setup.
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

	scribe, err := node.NewScribe(g.ScribeTestDB, clients, scribeConfig, g.ScribeMetrics)
	Nil(g.T(), err)
	scribeClient := client.NewEmbeddedScribe("sqlite", g.DBPath, g.ScribeMetrics)

	go func() {
		scribeErr := scribeClient.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), scribeErr)
		}
	}()
	go func() {
		scribeError := scribe.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), scribeError)
		}
	}()

	guard, err := guard.NewGuard(g.GetTestContext(), testConfig, omniRPCClient, scribeClient.ScribeClient, g.GuardTestDB, g.GuardMetrics)
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
	notaryProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.NotaryBondedSigner)
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
	Equal(g.T(), status.Flag(), uint8(1))
	Nil(g.T(), err)

	// Before submitting the attestation, ensure that there are no disputes opened.
	err = g.DestinationDomainClient.LightManager().GetDispute(g.GetTestContext(), big.NewInt(0))
	NotNil(g.T(), err)

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

	// Verify that the guard eventually marks the accused agent as Fraudulent
	g.Eventually(func() bool {
		status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)

		if status.Flag() == uint8(4) {
			return true
		}

		// Make sure that scribe keeps producing new blocks
		bumpTx, err := g.TestContractOnSummit.EmitAgentsEventA(txContextSummit.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		Nil(g.T(), err)
		g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), bumpTx)
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

func (g GuardSuite) TestZZZZZZ() {
	testDone := false
	defer func() {
		testDone = true
	}()

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

	omniRPCClient := omniClient.NewOmnirpcClient(g.TestOmniRPC, g.GuardMetrics, omniClient.WithCaptureReqRes())

	// Scribe setup.
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

	scribe, err := node.NewScribe(g.ScribeTestDB, clients, scribeConfig, g.ScribeMetrics)
	Nil(g.T(), err)
	scribeClient := client.NewEmbeddedScribe("sqlite", g.DBPath, g.ScribeMetrics)

	go func() {
		scribeErr := scribeClient.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), scribeErr)
		}
	}()
	go func() {
		scribeError := scribe.Start(g.GetTestContext())
		if !testDone {
			Nil(g.T(), scribeError)
		}
	}()

	guard, err := guard.NewGuard(g.GetTestContext(), testConfig, omniRPCClient, scribeClient.ScribeClient, g.GuardTestDB, g.GuardMetrics)
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
	Equal(g.T(), status.Flag(), uint8(1))
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
	// TODO: Some of the calls are unnecessary. Figure out which ones and nuke them.
	guardStatus, err := g.DestinationDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Nil(g.T(), err)
	guardStatus, err = g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Nil(g.T(), err)
	guardProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)
	err = g.DestinationDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.GuardUnbondedSigner,
		g.GuardBondedSigner,
		guardStatus,
		guardProof,
	)
	Nil(g.T(), err)
	guardStatus, err = g.DestinationDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner.Address())
	Nil(g.T(), err)

	notaryStatus, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
	Nil(g.T(), err)
	notaryProof, err := g.SummitDomainClient.BondingManager().GetProof(g.GetTestContext(), g.NotaryBondedSigner)
	Nil(g.T(), err)
	err = g.DestinationDomainClient.LightManager().UpdateAgentStatus(
		g.GetTestContext(),
		g.NotaryUnbondedSigner,
		g.NotaryBondedSigner,
		notaryStatus,
		notaryProof,
	)
	Nil(g.T(), err)
	notaryStatus, err = g.DestinationDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
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

	// Verify that the guard eventually marks the accused agent as Fraudulent
	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
	txContextDestination := g.TestBackendDestination.GetTxContext(g.GetTestContext(), g.LightInboxMetadataOnDestination.OwnerPtr())
	g.Eventually(func() bool {
		status, err := g.SummitDomainClient.BondingManager().GetAgentStatus(g.GetTestContext(), g.NotaryBondedSigner.Address())
		Nil(g.T(), err)
		if status.Flag() == uint8(4) {
			return true
		}

		// Make sure that scribe keeps producing new blocks
		bumpTx, err := g.TestContractOnSummit.EmitAgentsEventA(txContextSummit.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		Nil(g.T(), err)
		g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), bumpTx)
		bumpTx, err = g.TestContractOnDestination.EmitAgentsEventA(txContextDestination.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		Nil(g.T(), err)
		g.TestBackendDestination.WaitForConfirmation(g.GetTestContext(), bumpTx)
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
