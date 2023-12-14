package notary_test

import (
	"math/big"
	"os"
	"testing"

	"github.com/synapsecns/sanguine/agents/agents/notary"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/client"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/types"
)

func RemoveNotaryTempFile(t *testing.T, fileName string) {
	t.Helper()
	err := os.Remove(fileName)
	Nil(t, err)
}

func (u *NotarySuite) TestEnsureNotaryActive() {
	// Use a new wallet for Unknown notary.
	notaryWallet, err := wallet.FromRandom()
	u.Nil(err)

	// Fetch the owner signer for calls to BondingManager.
	_, bondingManagerHarnessContract := u.TestDeployManager.GetBondingManagerHarness(u.GetTestContext(), u.TestBackendSummit)
	bondingManagerHarnessOwnerPtr, err := bondingManagerHarnessContract.BondingManagerHarnessCaller.Owner(&bind.CallOpts{Context: u.GetTestContext()})
	u.Nil(err)
	bondingManagerHarnessOwnerAuth := u.TestBackendSummit.GetTxContext(u.GetTestContext(), &bondingManagerHarnessOwnerPtr)
	ownerWallet, err := wallet.FromPrivateKey(bondingManagerHarnessOwnerAuth.PrivateKey)
	u.Nil(err)

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
			File: filet.TmpFile(u.T(), "", notaryWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(u.T(), "", notaryWallet.PrivateKeyHex()).Name(),
		},
		OwnerSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(u.T(), "", ownerWallet.PrivateKeyHex()).Name(),
		},
		RefreshIntervalSeconds: 5,
	}

	omniRPCClient := omniClient.NewOmnirpcClient(u.TestOmniRPC, u.NotaryMetrics, omniClient.WithCaptureReqRes())

	scribeClient := client.NewEmbeddedScribe("sqlite", u.DBPath, u.ScribeMetrics)
	go func() {
		scribeErr := scribeClient.Start(u.GetTestContext())
		Nil(u.T(), scribeErr)
	}()

	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig, omniRPCClient, u.NotaryTestDB, u.NotaryMetrics)
	Nil(u.T(), err)

	// Verify that the notary is Unknown.
	verifyStatus := func(agent common.Address, flag types.AgentFlagType) {
		u.Eventually(func() bool {
			status, err := u.SummitDomainClient.BondingManager().GetAgentStatus(u.GetTestContext(), agent)
			Nil(u.T(), err)
			if status.Flag() == flag {
				return true
			}
			return false
		})
	}
	verifyStatus(notaryWallet.Address(), types.AgentFlagUnknown)

	// Ensure that the notary is active. Since it is currently Unknown, this should result in a call
	// to BondingManager.addAgent().
	err = notary.EnsureNotaryActive(u.GetTestContext())
	Nil(u.T(), err)
	verifyStatus(notaryWallet.Address(), types.AgentFlagActive)

	// Ensure that the notary is now active.
	err = notary.EnsureNotaryActive(u.GetTestContext())
	Nil(u.T(), err)
	verifyStatus(notaryWallet.Address(), types.AgentFlagActive)
}

func (u *NotarySuite) TestLoadStates() {
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
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		RefreshIntervalSeconds: 5,
	}

	omniRPCClient := omniClient.NewOmnirpcClient(u.TestOmniRPC, u.NotaryMetrics, omniClient.WithCaptureReqRes())

	scribeClient := client.NewEmbeddedScribe("sqlite", u.DBPath, u.ScribeMetrics)
	go func() {
		scribeErr := scribeClient.Start(u.GetTestContext())
		Nil(u.T(), scribeErr)
	}()

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig, omniRPCClient, scribeClient.ScribeClient, u.GuardTestDB, u.GuardMetrics)
	Nil(u.T(), err)

	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig, omniRPCClient, u.NotaryTestDB, u.NotaryMetrics)
	Nil(u.T(), err)

	// Fetch the guard's latest states.
	notary.LoadGuardLatestStates(u.GetTestContext())
	guardLatestStates := notary.GuardLatestStates(u.GetTestContext())
	expectedStates := map[uint32]types.State{}
	u.Equal(expectedStates, guardLatestStates)

	// Fetch the notary's latest states.
	notary.LoadMyLatestStates(u.GetTestContext())
	myLatestStates := notary.MyLatestStates(u.GetTestContext())
	u.Equal(expectedStates, myLatestStates)

	// Start the agents.
	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()
	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	// Send a message, which should generate a new State.
	tips := types.NewTips(big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)))
	optimisticSeconds := uint32(10)
	body := []byte{byte(gofakeit.Uint32())}
	txContextOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)
	txContextTestClientOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.TestClientMetadataOnOrigin.OwnerPtr())
	gasLimit := uint64(10000000)
	version := uint32(1)
	testClientOnOriginTx, err := u.TestClientOnOrigin.SendMessage(
		txContextTestClientOrigin.TransactOpts,
		uint32(u.TestBackendDestination.GetChainID()),
		u.TestClientMetadataOnDestination.Address(),
		optimisticSeconds,
		gasLimit,
		version,
		body)

	u.Nil(err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), testClientOnOriginTx)

	// Verify that the newly generated guard state can be loaded by the Notary.
	u.Eventually(func() bool {
		notary.LoadGuardLatestStates(u.GetTestContext())
		guardLatestStates := notary.GuardLatestStates(u.GetTestContext())
		originState, ok := guardLatestStates[u.OriginDomainClient.Config().DomainID]
		if !ok {
			return false
		}
		return originState.Nonce() == 1
	})

	// Verify that the newly generated notary state can be loaded by the Notary.
	u.Eventually(func() bool {
		notary.LoadMyLatestStates(u.GetTestContext())
		myLatestStates := notary.MyLatestStates(u.GetTestContext())
		originState, ok := myLatestStates[u.OriginDomainClient.Config().DomainID]
		if !ok {
			return false
		}
		return originState.Nonce() == 1
	})
}

//nolint:maintidx
func (u *NotarySuite) TestSubmitAttestation() {
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

	omniRPCClient := omniClient.NewOmnirpcClient(u.TestOmniRPC, u.NotaryMetrics, omniClient.WithCaptureReqRes())

	scribeClient := client.NewEmbeddedScribe("sqlite", u.DBPath, u.ScribeMetrics)
	go func() {
		scribeErr := scribeClient.Start(u.GetTestContext())
		Nil(u.T(), scribeErr)
	}()

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig, omniRPCClient, scribeClient.ScribeClient, u.GuardTestDB, u.GuardMetrics)
	Nil(u.T(), err)

	// Send a message.
	tips := types.NewTips(big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)))
	optimisticSeconds := uint32(10)
	body := []byte{byte(gofakeit.Uint32())}
	txContextOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)
	txContextTestClientOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.TestClientMetadataOnOrigin.OwnerPtr())
	gasLimit := uint64(10000000)
	version := uint32(1)
	testClientOnOriginTx, err := u.TestClientOnOrigin.SendMessage(
		txContextTestClientOrigin.TransactOpts,
		uint32(u.TestBackendDestination.GetChainID()),
		u.TestClientMetadataOnDestination.Address(),
		optimisticSeconds,
		gasLimit,
		version,
		body)
	u.Nil(err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), testClientOnOriginTx)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	// Verify that the guard has submitted a snapshot.
	u.Eventually(func() bool {
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

	// Verify that the message has resulted in a new state for the origin on summit.
	u.Eventually(func() bool {
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

	// Verify that the notary is Unknown on destination.
	agentStatus, err := u.DestinationContract.AgentStatus(&bind.CallOpts{Context: u.GetTestContext()}, u.NotaryBondedSigner.Address())
	Nil(u.T(), err)
	Equal(u.T(), types.AgentFlagUnknown, types.AgentFlagType(agentStatus.Flag))

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	// Verify that the notary is able to register itself as Active on the destination.
	u.Eventually(func() bool {
		agentStatus, err := u.DestinationContract.AgentStatus(&bind.CallOpts{Context: u.GetTestContext()}, u.NotaryBondedSigner.Address())
		Nil(u.T(), err)
		return types.AgentFlagType(agentStatus.Flag) == types.AgentFlagActive
	})

	// Verify that the notary has submitted a snapshot containing the new origin state.
	u.Eventually(func() bool {
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

	// Verify that the notary finally submits an attestation.
	u.Eventually(func() bool {
		attestationsAmount, err := u.DestinationContract.AttestationsAmount(&bind.CallOpts{Context: u.GetTestContext()})
		Nil(u.T(), err)

		return attestationsAmount != nil && attestationsAmount.Uint64() >= uint64(1)
	})
}
