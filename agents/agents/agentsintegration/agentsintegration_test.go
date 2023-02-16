package agentsintegration_test

import (
	"math/big"
	"time"

	executor2 "github.com/synapsecns/sanguine/agents/agents/executor"

	"github.com/Flaque/filet"
	awsTime "github.com/aws/smithy-go/time"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/stretchr/testify/assert"
	executorCfg "github.com/synapsecns/sanguine/agents/agents/executor/config"
	types2 "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/config"
	agentsConfig "github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/merkle"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
)

// TestGuardAndNotaryOnlyIntegrationE2E is an integration involving just a guard and notary
// submitting a single message to the destination.
//
//nolint:dupl
func (u AgentsIntegrationSuite) TestGuardAndNotaryOnlyIntegrationE2E() {
	notaryTestConfig := config.NotaryConfig{
		DestinationDomain: u.DestinationDomainClient.Config(),
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	guardTestConfig := config.GuardConfig{
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		DestinationDomains: map[string]config.DomainConfig{
			"destination_client": u.DestinationDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig)
	Nil(u.T(), err)

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig)
	Nil(u.T(), err)

	guardDBType, err := dbcommon.DBTypeFromString(guardTestConfig.Database.Type)
	Nil(u.T(), err)

	guardDBHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), guardDBType, guardTestConfig.Database.ConnString, "guard")
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

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		retrievedInProgressAttestation, err := guardDBHandle.RetrieveNewestInProgressAttestationIfInState(
			u.GetTestContext(),
			u.OriginDomainClient.Config().DomainID,
			u.DestinationDomainClient.Config().DomainID,
			types.AttestationStateConfirmedOnDestination)

		isTrue := err == nil &&
			retrievedInProgressAttestation != nil &&
			retrievedInProgressAttestation.SignedAttestation().Attestation().Nonce() == uint32(1) &&
			u.OriginDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Origin() &&
			u.DestinationDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Destination() &&
			[32]byte{} != retrievedInProgressAttestation.SignedAttestation().Attestation().Root() &&
			len(retrievedInProgressAttestation.SignedAttestation().NotarySignatures()) == 1 &&
			len(retrievedInProgressAttestation.SignedAttestation().GuardSignatures()) == 1 &&
			retrievedInProgressAttestation.AttestationState() == types.AttestationStateConfirmedOnDestination

		return isTrue
	})
}

// TestGuardAndNotaryOnlyMultipleMessagesIntegrationE2E is an integration involving just a guard and notary
// submitting multiple messages to the destination.
//
//nolint:dupl,gocognit,cyclop
func (u AgentsIntegrationSuite) TestGuardAndNotaryOnlyMultipleMessagesIntegrationE2E() {
	numMessages := 5

	notaryTestConfig := config.NotaryConfig{
		DestinationDomain: u.DestinationDomainClient.Config(),
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	guardTestConfig := config.GuardConfig{
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		DestinationDomains: map[string]config.DomainConfig{
			"destination_client": u.DestinationDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig)
	Nil(u.T(), err)

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig)
	Nil(u.T(), err)

	guardDBType, err := dbcommon.DBTypeFromString(guardTestConfig.Database.Type)
	Nil(u.T(), err)

	guardDBHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), guardDBType, guardTestConfig.Database.ConnString, "notary")
	Nil(u.T(), err)

	originAuth := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), nil)

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(u.T(), err)

	for i := 0; i < numMessages; i++ {
		tx, err := u.OriginContract.Dispatch(
			originAuth.TransactOpts,
			u.DestinationDomainClient.Config().DomainID,
			[32]byte{},
			gofakeit.Uint32(),
			encodedTips,
			[]byte(gofakeit.Paragraph(3, 2, 1, " ")))
		Nil(u.T(), err)
		u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)
		currRoot, currDispatchBlockNumber, err := u.OriginContract.GetHistoricalRoot(&bind.CallOpts{Context: u.GetTestContext()}, u.DestinationDomainClient.Config().DomainID, uint32(i+1))
		Nil(u.T(), err)
		Greater(u.T(), currDispatchBlockNumber.Uint64(), uint64(0))
		NotEqual(u.T(), [32]byte{}, currRoot)
	}

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		retrievedInProgressAttestation, err := guardDBHandle.RetrieveNewestInProgressAttestationIfInState(
			u.GetTestContext(),
			u.OriginDomainClient.Config().DomainID,
			u.DestinationDomainClient.Config().DomainID,
			types.AttestationStateConfirmedOnDestination)

		isTrue := err == nil &&
			retrievedInProgressAttestation != nil &&
			retrievedInProgressAttestation.SignedAttestation().Attestation().Nonce() == uint32(numMessages) &&
			u.OriginDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Origin() &&
			u.DestinationDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Destination() &&
			[32]byte{} != retrievedInProgressAttestation.SignedAttestation().Attestation().Root() &&
			len(retrievedInProgressAttestation.SignedAttestation().NotarySignatures()) == 1 &&
			len(retrievedInProgressAttestation.SignedAttestation().GuardSignatures()) == 1 &&
			retrievedInProgressAttestation.AttestationState() == types.AttestationStateConfirmedOnDestination

		if isTrue {
			i := numMessages - 1
			currNonce := uint32(i + 1)
			currInProgressAttestation, err := guardDBHandle.RetrieveInProgressAttestation(
				u.GetTestContext(),
				u.OriginDomainClient.Config().DomainID,
				u.DestinationDomainClient.Config().DomainID,
				currNonce)
			if err != nil {
				return false
			}
			if currInProgressAttestation.AttestationState() != types.AttestationStateConfirmedOnDestination {
				return false
			}
		}

		return isTrue
	})
}

// TestAllAgentsSingleMessageIntegrationE2E is an integration involving just a guard, notary
// and executor executing a single message to the destination.
//
//nolint:dupl,cyclop,maintidx
func (u AgentsIntegrationSuite) TestAllAgentsSingleMessageIntegrationE2E() {
	notaryTestConfig := config.NotaryConfig{
		DestinationDomain: u.DestinationDomainClient.Config(),
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	guardTestConfig := config.GuardConfig{
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		DestinationDomains: map[string]config.DomainConfig{
			"destination_client": u.DestinationDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig)
	Nil(u.T(), err)

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig)
	Nil(u.T(), err)

	guardDBType, err := dbcommon.DBTypeFromString(guardTestConfig.Database.Type)
	Nil(u.T(), err)

	guardDBHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), guardDBType, guardTestConfig.Database.ConnString, "guard")
	Nil(u.T(), err)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	tips := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips, err := types.EncodeTips(tips)
	u.Nil(err)

	optimisticSeconds := uint32(10)

	testContractDest, testContractRef := u.TestDeployManager.GetAgentsTestContract(u.GetTestContext(), u.TestBackendDestination)
	testTransactOpts := u.TestBackendDestination.GetTxContext(u.GetTestContext(), nil)

	recipient := testContractDest.Address().Hash()
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}

	txContextOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)

	tx, err := u.OriginContract.Dispatch(txContextOrigin.TransactOpts, uint32(u.TestBackendDestination.GetChainID()), recipient, optimisticSeconds, encodedTips, body)
	u.Nil(err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		retrievedInProgressAttestation, err := guardDBHandle.RetrieveNewestInProgressAttestationIfInState(
			u.GetTestContext(),
			u.OriginDomainClient.Config().DomainID,
			u.DestinationDomainClient.Config().DomainID,
			types.AttestationStateConfirmedOnDestination)

		isTrue := err == nil &&
			retrievedInProgressAttestation != nil &&
			retrievedInProgressAttestation.SignedAttestation().Attestation().Nonce() == nonce &&
			u.OriginDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Origin() &&
			u.DestinationDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Destination() &&
			[32]byte{} != retrievedInProgressAttestation.SignedAttestation().Attestation().Root() &&
			len(retrievedInProgressAttestation.SignedAttestation().NotarySignatures()) == 1 &&
			len(retrievedInProgressAttestation.SignedAttestation().GuardSignatures()) == 1 &&
			retrievedInProgressAttestation.AttestationState() == types.AttestationStateConfirmedOnDestination

		return isTrue
	})

	// Beginning of executor part that was pasted.
	testDone := false
	defer func() {
		testDone = true
	}()

	originClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendOrigin.RPCAddress())
	u.Nil(err)
	destinationClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendDestination.RPCAddress())
	u.Nil(err)

	originConfig := scribeConfig.ContractConfig{
		Address:    u.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(u.TestBackendOrigin.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    u.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(u.TestBackendDestination.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{destinationConfig},
	}
	scribeConf := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		uint32(u.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(u.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
	}

	scribe, err := node.NewScribe(u.ScribeTestDB, clients, scribeConf)
	u.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", u.DBPath)
	go func() {
		scribeErr := scribeClient.Start(u.GetTestContext())
		u.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(u.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       uint32(u.TestBackendOrigin.GetChainID()),
				OriginAddress: u.OriginContract.Address().String(),
			},
			{
				ChainID:            uint32(u.TestBackendDestination.GetChainID()),
				DestinationAddress: u.DestinationContract.Address().String(),
			},
		},
		BaseOmnirpcURL: gofakeit.URL(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor2.Backend{
		uint32(u.TestBackendOrigin.GetChainID()):      u.TestBackendOrigin,
		uint32(u.TestBackendDestination.GetChainID()): u.TestBackendDestination,
	}

	urls := map[uint32]string{
		uint32(u.TestBackendOrigin.GetChainID()):      u.TestBackendOrigin.RPCAddress(),
		uint32(u.TestBackendDestination.GetChainID()): u.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor2.NewExecutorInjectedBackend(u.GetTestContext(), excCfg, u.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	u.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.Start(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	// Listen with the exec.
	go func() {
		execErr := exec.Listen(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	go func() {
		execErr := exec.SetMinimumTime(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	sender, err := u.TestBackendOrigin.Signer().Sender(tx)
	u.Nil(err)

	header := types.NewHeader(uint32(u.TestBackendOrigin.GetChainID()), sender.Hash(), nonce, uint32(u.TestBackendDestination.GetChainID()), recipient, optimisticSeconds)
	message := types.NewMessage(header, tips, body)

	tree := merkle.NewTree()

	leaf, err := message.ToLeaf()
	u.Nil(err)

	tree.Insert(leaf[:])

	root, err := tree.Root(1)
	u.Nil(err)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	continueChan := make(chan bool, 1)

	chainID := uint32(u.TestBackendOrigin.GetChainID())
	destination := uint32(u.TestBackendDestination.GetChainID())
	// Wait for message to be stored in the database.
	u.Eventually(func() bool {
		_, err = u.ExecutorTestDB.GetAttestationBlockNumber(u.GetTestContext(), types2.DBAttestation{
			ChainID:     &chainID,
			Destination: &destination,
			Nonce:       &nonce,
		})
		if err == nil {
			continueChan <- true
			return true
		}
		return false
	})

	<-continueChan

	executed, err := exec.Execute(u.GetTestContext(), message)
	u.Nil(err)
	u.False(executed)

	u.Eventually(func() bool {
		executed, err := exec.Execute(u.GetTestContext(), message)
		if err != nil {
			return false
		}
		if executed {
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

	exec.Stop(uint32(u.TestBackendOrigin.GetChainID()))
	exec.Stop(uint32(u.TestBackendDestination.GetChainID()))

	// End of executor part that was pasted.
}

// TestAllAgentsMultipleeMessagesIntegrationE2E is an integration involving just a guard, notary
// and executor executing multiple messages to the destination.
//
//nolint:dupl,cyclop,maintidx,gocognit
func (u AgentsIntegrationSuite) TestAllAgentsMultipleMessagesIntegrationE2E() {
	numMessages := 5

	notaryTestConfig := config.NotaryConfig{
		DestinationDomain: u.DestinationDomainClient.Config(),
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	guardTestConfig := config.GuardConfig{
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		DestinationDomains: map[string]config.DomainConfig{
			"destination_client": u.DestinationDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig)
	Nil(u.T(), err)

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig)
	Nil(u.T(), err)

	guardDBType, err := dbcommon.DBTypeFromString(guardTestConfig.Database.Type)
	Nil(u.T(), err)

	guardDBHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), guardDBType, guardTestConfig.Database.ConnString, "guard")
	Nil(u.T(), err)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	tips := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips, err := types.EncodeTips(tips)
	u.Nil(err)

	optimisticSeconds := uint32(10)
	txContextOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.OriginContractMetadata.OwnerPtr())
	testContractDest, testContractRef := u.TestDeployManager.GetAgentsTestContract(u.GetTestContext(), u.TestBackendDestination)
	testTransactOpts := u.TestBackendDestination.GetTxContext(u.GetTestContext(), nil)

	recipient := testContractDest.Address().Hash()
	txContextOrigin.Value = types.TotalTips(tips)

	messages := []types.Message{}
	for i := 0; i < numMessages; i++ {
		body := []byte{byte(gofakeit.Uint32())}

		tx, err := u.OriginContract.Dispatch(
			txContextOrigin.TransactOpts,
			u.DestinationDomainClient.Config().DomainID,
			recipient,
			optimisticSeconds,
			encodedTips,
			body)
		Nil(u.T(), err)
		u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)
		currRoot, currDispatchBlockNumber, err := u.OriginContract.GetHistoricalRoot(&bind.CallOpts{Context: u.GetTestContext()}, u.DestinationDomainClient.Config().DomainID, uint32(i+1))
		Nil(u.T(), err)
		Greater(u.T(), currDispatchBlockNumber.Uint64(), uint64(0))
		NotEqual(u.T(), [32]byte{}, currRoot)

		sender, err := u.TestBackendOrigin.Signer().Sender(tx)
		u.Nil(err)

		header := types.NewHeader(uint32(u.TestBackendOrigin.GetChainID()), sender.Hash(), uint32(i+1), uint32(u.TestBackendDestination.GetChainID()), recipient, optimisticSeconds)
		message := types.NewMessage(header, tips, body)
		messages = append(messages, message)
	}

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		retrievedInProgressAttestation, err := guardDBHandle.RetrieveNewestInProgressAttestationIfInState(
			u.GetTestContext(),
			u.OriginDomainClient.Config().DomainID,
			u.DestinationDomainClient.Config().DomainID,
			types.AttestationStateConfirmedOnDestination)

		nonce := uint32(numMessages)
		isTrue := err == nil &&
			retrievedInProgressAttestation != nil &&
			retrievedInProgressAttestation.SignedAttestation().Attestation().Nonce() == nonce &&
			u.OriginDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Origin() &&
			u.DestinationDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Destination() &&
			[32]byte{} != retrievedInProgressAttestation.SignedAttestation().Attestation().Root() &&
			len(retrievedInProgressAttestation.SignedAttestation().NotarySignatures()) == 1 &&
			len(retrievedInProgressAttestation.SignedAttestation().GuardSignatures()) == 1 &&
			retrievedInProgressAttestation.AttestationState() == types.AttestationStateConfirmedOnDestination

		return isTrue
	})

	// Beginning of executor part that was pasted.
	testDone := false
	defer func() {
		testDone = true
	}()

	originClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendOrigin.RPCAddress())
	u.Nil(err)
	destinationClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendDestination.RPCAddress())
	u.Nil(err)

	originConfig := scribeConfig.ContractConfig{
		Address:    u.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(u.TestBackendOrigin.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    u.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(u.TestBackendDestination.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{destinationConfig},
	}
	scribeConf := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		uint32(u.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(u.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
	}

	scribe, err := node.NewScribe(u.ScribeTestDB, clients, scribeConf)
	u.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", u.DBPath)
	go func() {
		scribeErr := scribeClient.Start(u.GetTestContext())
		u.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(u.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       uint32(u.TestBackendOrigin.GetChainID()),
				OriginAddress: u.OriginContract.Address().String(),
			},
			{
				ChainID:            uint32(u.TestBackendDestination.GetChainID()),
				DestinationAddress: u.DestinationContract.Address().String(),
			},
		},
		BaseOmnirpcURL: gofakeit.URL(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor2.Backend{
		uint32(u.TestBackendOrigin.GetChainID()):      u.TestBackendOrigin,
		uint32(u.TestBackendDestination.GetChainID()): u.TestBackendDestination,
	}

	urls := map[uint32]string{
		uint32(u.TestBackendOrigin.GetChainID()):      u.TestBackendOrigin.RPCAddress(),
		uint32(u.TestBackendDestination.GetChainID()): u.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor2.NewExecutorInjectedBackend(u.GetTestContext(), excCfg, u.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	u.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.Start(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	// Listen with the exec.
	go func() {
		execErr := exec.Listen(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	go func() {
		execErr := exec.SetMinimumTime(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	tree := merkle.NewTree()

	for i := 0; i < numMessages; i++ {
		message := messages[i]
		leaf, err := message.ToLeaf()
		u.Nil(err)

		tree.Insert(leaf[:])
	}

	root, err := tree.Root(uint32(numMessages))
	u.Nil(err)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	continueChan := make(chan bool, 1)

	chainID := uint32(u.TestBackendOrigin.GetChainID())
	destination := uint32(u.TestBackendDestination.GetChainID())
	// Wait for message to be stored in the database.
	u.Eventually(func() bool {
		trueCount := 0
		for i := 0; i < numMessages; i++ {
			nonce := uint32(i + 1)
			_, err = u.ExecutorTestDB.GetAttestationBlockNumber(u.GetTestContext(), types2.DBAttestation{
				ChainID:     &chainID,
				Destination: &destination,
				Nonce:       &nonce,
			})
			if err == nil {
				trueCount++
				continue
			}
			return false
		}
		if trueCount == numMessages {
			continueChan <- true
			return true
		}
		return false
	})

	<-continueChan

	for i := 0; i < numMessages; i++ {
		message := messages[i]
		executed, err := exec.Execute(u.GetTestContext(), message)
		u.Nil(err)
		u.False(executed)
	}

	u.Eventually(func() bool {
		trueCount := 0
		for i := 0; i < numMessages; i++ {
			message := messages[i]
			executed, err := exec.Execute(u.GetTestContext(), message)
			if err != nil {
				return false
			}
			if executed {
				trueCount++
				continue
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
		}
		return trueCount == numMessages
	})

	exec.Stop(uint32(u.TestBackendOrigin.GetChainID()))
	exec.Stop(uint32(u.TestBackendDestination.GetChainID()))

	// End of executor part that was pasted.
}

// TestAllAgentsSingleMessageWithTestClientIntegrationE2E is an integration involving just a guard, notary
// and executor executing a single message to the destination using the TestClient as the origin client and destination client.
//
//nolint:dupl,cyclop,maintidx,gocognit
func (u AgentsIntegrationSuite) TestAllAgentsSingleMessageWithTestClientIntegrationE2E() {
	notaryTestConfig := config.NotaryConfig{
		DestinationDomain: u.DestinationDomainClient.Config(),
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	guardTestConfig := config.GuardConfig{
		AttestationDomain: u.AttestationDomainClient.Config(),
		OriginDomains: map[string]config.DomainConfig{
			"origin_client": u.OriginDomainClient.Config(),
		},
		DestinationDomains: map[string]config.DomainConfig{
			"destination_client": u.DestinationDomainClient.Config(),
		},
		BondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: config.SignerConfig{
			Type: config.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.GuardUnbondedWallet.PrivateKeyHex()).Name(),
		},
		Database: config.DBConfig{
			Type:       dbcommon.Sqlite.String(),
			DBPath:     filet.TmpDir(u.T(), ""),
			ConnString: filet.TmpDir(u.T(), ""),
		},
		RefreshIntervalInSeconds: 1,
	}
	notary, err := notary.NewNotary(u.GetTestContext(), notaryTestConfig)
	Nil(u.T(), err)

	guard, err := guard.NewGuard(u.GetTestContext(), guardTestConfig)
	Nil(u.T(), err)

	guardDBType, err := dbcommon.DBTypeFromString(guardTestConfig.Database.Type)
	Nil(u.T(), err)

	guardDBHandle, err := sql.NewStoreFromConfig(u.GetTestContext(), guardDBType, guardTestConfig.Database.ConnString, "guard")
	Nil(u.T(), err)

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = notary.Start(u.GetTestContext())
	}()

	go func() {
		// we don't check errors here since this will error on cancellation at the end of the test
		_ = guard.Start(u.GetTestContext())
	}()

	tips := types.NewTips(big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)))

	optimisticSeconds := uint32(10)

	// We use the agents test contract to simply force transactions on the destination chain
	// in order to advance the time.
	_, testContractRef := u.TestDeployManager.GetAgentsTestContract(u.GetTestContext(), u.TestBackendDestination)
	testTransactOpts := u.TestBackendDestination.GetTxContext(u.GetTestContext(), nil)

	recipient := u.TestClientMetadataOnDestination.Address().Hash()
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}

	txContextOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)

	txContextTestClientOrigin := u.TestBackendOrigin.GetTxContext(u.GetTestContext(), u.TestClientMetadataOnOrigin.OwnerPtr())

	testClientOnOriginTx, err := u.TestClientOnOrigin.SendMessage(txContextTestClientOrigin.TransactOpts, uint32(u.TestBackendDestination.GetChainID()), u.TestClientMetadataOnDestination.Address(), optimisticSeconds, body)
	u.Nil(err)
	u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), testClientOnOriginTx)

	// tx, err := u.OriginContract.Dispatch(txContextOrigin.TransactOpts, uint32(u.TestBackendDestination.GetChainID()), recipient, optimisticSeconds, encodedTips, body)
	// u.Nil(err)
	// u.TestBackendOrigin.WaitForConfirmation(u.GetTestContext(), tx)

	u.Eventually(func() bool {
		_ = awsTime.SleepWithContext(u.GetTestContext(), time.Second*5)
		retrievedInProgressAttestation, err := guardDBHandle.RetrieveNewestInProgressAttestationIfInState(
			u.GetTestContext(),
			u.OriginDomainClient.Config().DomainID,
			u.DestinationDomainClient.Config().DomainID,
			types.AttestationStateConfirmedOnDestination)

		isTrue := err == nil &&
			retrievedInProgressAttestation != nil &&
			retrievedInProgressAttestation.SignedAttestation().Attestation().Nonce() == nonce &&
			u.OriginDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Origin() &&
			u.DestinationDomainClient.Config().DomainID == retrievedInProgressAttestation.SignedAttestation().Attestation().Destination() &&
			[32]byte{} != retrievedInProgressAttestation.SignedAttestation().Attestation().Root() &&
			len(retrievedInProgressAttestation.SignedAttestation().NotarySignatures()) == 1 &&
			len(retrievedInProgressAttestation.SignedAttestation().GuardSignatures()) == 1 &&
			retrievedInProgressAttestation.AttestationState() == types.AttestationStateConfirmedOnDestination

		return isTrue
	})

	// Beginning of executor part that was pasted.
	testDone := false
	defer func() {
		testDone = true
	}()

	originClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendOrigin.RPCAddress())
	u.Nil(err)
	destinationClient, err := backfill.DialBackend(u.GetTestContext(), u.TestBackendDestination.RPCAddress())
	u.Nil(err)

	originConfig := scribeConfig.ContractConfig{
		Address:    u.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(u.TestBackendOrigin.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{originConfig},
	}
	destinationConfig := scribeConfig.ContractConfig{
		Address:    u.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := scribeConfig.ChainConfig{
		ChainID:               uint32(u.TestBackendDestination.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []scribeConfig.ContractConfig{destinationConfig},
	}
	scribeConf := scribeConfig.Config{
		Chains: []scribeConfig.ChainConfig{originChainConfig, destinationChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		uint32(u.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(u.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
	}

	scribe, err := node.NewScribe(u.ScribeTestDB, clients, scribeConf)
	u.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", u.DBPath)
	go func() {
		scribeErr := scribeClient.Start(u.GetTestContext())
		u.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(u.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       uint32(u.TestBackendOrigin.GetChainID()),
				OriginAddress: u.OriginContract.Address().String(),
			},
			{
				ChainID:            uint32(u.TestBackendDestination.GetChainID()),
				DestinationAddress: u.DestinationContract.Address().String(),
			},
		},
		BaseOmnirpcURL: gofakeit.URL(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(u.T(), "", u.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor2.Backend{
		uint32(u.TestBackendOrigin.GetChainID()):      u.TestBackendOrigin,
		uint32(u.TestBackendDestination.GetChainID()): u.TestBackendDestination,
	}

	urls := map[uint32]string{
		uint32(u.TestBackendOrigin.GetChainID()):      u.TestBackendOrigin.RPCAddress(),
		uint32(u.TestBackendDestination.GetChainID()): u.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor2.NewExecutorInjectedBackend(u.GetTestContext(), excCfg, u.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	u.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.Start(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	// Listen with the exec.
	go func() {
		execErr := exec.Listen(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	go func() {
		execErr := exec.SetMinimumTime(u.GetTestContext())
		if !testDone {
			u.Nil(execErr)
		}
	}()

	sender, err := u.TestBackendOrigin.Signer().Sender(testClientOnOriginTx)
	u.Nil(err)

	header := types.NewHeader(uint32(u.TestBackendOrigin.GetChainID()), sender.Hash(), nonce, uint32(u.TestBackendDestination.GetChainID()), recipient, optimisticSeconds)
	message := types.NewMessage(header, tips, body)

	tree := merkle.NewTree()

	leaf, err := message.ToLeaf()
	u.Nil(err)

	tree.Insert(leaf[:])

	root, err := tree.Root(1)
	u.Nil(err)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	continueChan := make(chan bool, 1)

	chainID := uint32(u.TestBackendOrigin.GetChainID())
	destinationID := uint32(u.TestBackendDestination.GetChainID())
	// Wait for message to be stored in the database.
	u.Eventually(func() bool {
		_, err = u.ExecutorTestDB.GetAttestationBlockNumber(u.GetTestContext(), types2.DBAttestation{
			ChainID:     &chainID,
			Destination: &destinationID,
			Nonce:       &nonce,
		})
		if err == nil {
			continueChan <- true
			return true
		}
		return false
	})

	<-continueChan

	executed, err := exec.Execute(u.GetTestContext(), message)
	u.Nil(err)
	u.False(executed)

	u.Eventually(func() bool {
		executed, err := exec.Execute(u.GetTestContext(), message)
		if err != nil {
			return false
		}
		if executed {
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

	/*
		// Create a channel and subscription to receive TestClientMessageSent events as they are emitted from origin.
		messageSentSink := make(chan *testclient.TestClientMessageSent)
		sentMessage, err := u.TestClientOnOrigin.WatchMessageSent(&bind.WatchOpts{
			Context: u.GetTestContext()},
			messageSentSink)
		u.Nil(err)

		// Create a channel and subscription to receive TestClientMessageReceived events as they are emitted from destination.
		messageReceivedSink := make(chan *testclient.TestClientMessageReceived)
		receivedMessage, err := u.TestClientOnDestination.WatchMessageReceived(&bind.WatchOpts{
			Context: u.GetTestContext()},
			messageReceivedSink)
		u.Nil(err)

		watchCtx, cancel := context.WithTimeout(u.GetTestContext(), time.Second*10)
		defer cancel()

		select {
		// check for errors and fail
		case <-watchCtx.Done():
			u.T().Error(u.T(), fmt.Errorf("test context completed %w", u.GetTestContext().Err()))
		case <-sentMessage.Err():
			u.T().Error(u.T(), sentMessage.Err())
		// get message sent event
		case item := <-messageSentSink:
			u.NotNil(item)

			// Now sleep for a second before executing
			time.Sleep(time.Second)

			watchHandleCtx, cancel := context.WithTimeout(u.GetTestContext(), time.Second*10)
			defer cancel()

			select {
			// check for errors and fail
			case <-watchHandleCtx.Done():
				u.T().Error(u.T(), fmt.Errorf("test context completed %w", u.GetTestContext().Err()))
			case <-receivedMessage.Err():
				u.T().Error(u.T(), receivedMessage.Err())
			// get attestation accepted event
			case item := <-messageReceivedSink:
				u.NotNil(item)

				break
			}

			break
		}*/

	exec.Stop(uint32(u.TestBackendOrigin.GetChainID()))
	exec.Stop(uint32(u.TestBackendDestination.GetChainID()))

	// End of executor part that was pasted.
}
