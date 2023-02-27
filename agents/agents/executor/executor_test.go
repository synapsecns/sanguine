package executor_test

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	"math/big"
	"time"

	"github.com/Flaque/filet"
	agentsConfig "github.com/synapsecns/sanguine/agents/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/merkle"

	"github.com/brianvoe/gofakeit/v6"
	executorCfg "github.com/synapsecns/sanguine/agents/agents/executor/config"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
)

func (e *ExecutorSuite) TestExecutor() {
	testDone := false
	defer func() {
		testDone = true
	}()
	chainIDA := gofakeit.Uint32()
	chainIDB := chainIDA + 1

	simulatedChainA := geth.NewEmbeddedBackendForChainID(e.GetTestContext(), e.T(), big.NewInt(int64(chainIDA)))
	simulatedChainB := geth.NewEmbeddedBackendForChainID(e.GetTestContext(), e.T(), big.NewInt(int64(chainIDB)))
	simulatedClientA, err := backfill.DialBackend(e.GetTestContext(), simulatedChainA.RPCAddress())
	e.Nil(err)
	simulatedClientB, err := backfill.DialBackend(e.GetTestContext(), simulatedChainB.RPCAddress())
	e.Nil(err)
	testContractA, testRefA := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), simulatedChainA)
	testContractB, testRefB := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), simulatedChainB)
	transactOptsA := simulatedChainA.GetTxContext(e.GetTestContext(), nil)
	transactOptsB := simulatedChainB.GetTxContext(e.GetTestContext(), nil)

	// Emit two events on each chain.
	tx, err := testRefA.EmitAgentsEventA(transactOptsA.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	e.Nil(err)
	simulatedChainA.WaitForConfirmation(e.GetTestContext(), tx)
	tx, err = testRefA.EmitAgentsEventB(transactOptsA.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	e.Nil(err)
	simulatedChainA.WaitForConfirmation(e.GetTestContext(), tx)
	tx, err = testRefB.EmitAgentsEventAandB(transactOptsB.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
	e.Nil(err)
	simulatedChainB.WaitForConfirmation(e.GetTestContext(), tx)

	contractConfigA := config.ContractConfig{
		Address:    testContractA.Address().String(),
		StartBlock: 0,
	}
	contractConfigB := config.ContractConfig{
		Address:    testContractB.Address().String(),
		StartBlock: 0,
	}
	chainConfigA := config.ChainConfig{
		ChainID:               chainIDA,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfigA},
	}
	chainConfigB := config.ChainConfig{
		ChainID:               chainIDB,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfigB},
	}
	chainConfigs := []config.ChainConfig{chainConfigA, chainConfigB}
	scribeConfig := config.Config{
		Chains: chainConfigs,
	}

	clients := map[uint32][]backfill.ScribeBackend{
		chainIDA: {simulatedClientA, simulatedClientA},
		chainIDB: {simulatedClientB, simulatedClientB},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)

	go func() {
		scribeErr := scribeClient.Start(e.GetSuiteContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetSuiteContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:            chainIDA,
				OriginAddress:      testContractA.Address().String(),
				DestinationAddress: "not_needed",
			},
			{
				ChainID:       chainIDB,
				OriginAddress: testContractB.Address().String(),
			},
		},
		BaseOmnirpcURL: simulatedChainA.RPCAddress(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		chainIDA: simulatedChainA,
		chainIDB: simulatedChainB,
	}

	urls := map[uint32]string{
		chainIDA: simulatedChainA.RPCAddress(),
		chainIDB: simulatedChainB.RPCAddress(),
	}

	exc, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	// Start the executor.
	go func() {
		excErr := exc.Start(e.GetSuiteContext())
		if !testDone {
			e.Nil(excErr)
		}
	}()

	e.Eventually(func() bool {
		if len(exc.GetLogChan(chainIDA)) == 2 && len(exc.GetLogChan(chainIDB)) == 2 {
			logA := <-exc.GetLogChan(chainIDA)
			logB := <-exc.GetLogChan(chainIDA)
			e.Assert().Less(logA.BlockNumber, logB.BlockNumber)
			logC := <-exc.GetLogChan(chainIDB)
			logD := <-exc.GetLogChan(chainIDB)
			e.Assert().LessOrEqual(logC.BlockNumber, logD.BlockNumber)
			return true
		}

		return false
	})

	e.DeferAfterTest(func() {
		exc.Stop(chainIDA)
	})
}

func (e *ExecutorSuite) TestLotsOfLogs() {
	testDone := false
	defer func() {
		testDone = true
	}()
	chainID := gofakeit.Uint32()
	simulatedChain := geth.NewEmbeddedBackendForChainID(e.GetTestContext(), e.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(e.GetTestContext(), simulatedChain.RPCAddress())
	e.Nil(err)

	testContract, testRef := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), simulatedChain)
	transactOpts := simulatedChain.GetTxContext(e.GetTestContext(), nil)

	contractConfig := config.ContractConfig{
		Address:    testContract.Address().String(),
		StartBlock: 0,
	}
	chainConfig := config.ChainConfig{
		ChainID:               chainID,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{chainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		chainID: {simulatedClient, simulatedClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       chainID,
				OriginAddress: testContract.Address().String(),
			},
		},
		BaseOmnirpcURL: simulatedChain.RPCAddress(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		chainID: simulatedChain,
	}

	urls := map[uint32]string{
		chainID: simulatedChain.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.Start(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	// Emit 250 events.
	go func() {
		for i := 0; i < 250; i++ {
			tx, err := testRef.EmitAgentsEventB(transactOpts.TransactOpts, []byte{byte(i)}, big.NewInt(int64(i)), big.NewInt(int64(i)))
			e.Nil(err)
			simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)
		}
	}()

	e.Eventually(func() bool {
		return len(exec.GetLogChan(chainID)) == 250
	})

	e.DeferAfterTest(func() {
		exec.Stop(chainID)
	})
}

func (e *ExecutorSuite) TestMerkleInsert() {
	testDone := false
	defer func() {
		testDone = true
	}()

	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())

	contractConfig := config.ContractConfig{
		Address:    e.OriginContractMetadata.Address().String(),
		StartBlock: 0,
	}
	chainConfig := config.ChainConfig{
		ChainID:               chainID,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{chainConfig},
	}
	simulatedClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress())
	e.Nil(err)
	clients := map[uint32][]backfill.ScribeBackend{
		chainID: {simulatedClient, simulatedClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       chainID,
				OriginAddress: e.OriginContractMetadata.Address().String(),
			},
			{
				ChainID: destination,
			},
		},
		BaseOmnirpcURL: e.TestBackendOrigin.RPCAddress(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		chainID:     e.TestBackendOrigin,
		destination: nil,
	}

	urls := map[uint32]string{
		chainID:     e.TestBackendOrigin.RPCAddress(),
		destination: e.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	_, err = exec.GetMerkleTree(chainID, destination).Root(1)
	e.NotNil(err)

	testTree := merkle.NewTree()

	recipients := [][32]byte{{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())}}
	optimisticSeconds := []uint32{gofakeit.Uint32(), gofakeit.Uint32()}
	notaryTips := []*big.Int{big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32())))}
	broadcasterTips := []*big.Int{big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32())))}
	proverTips := []*big.Int{big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32())))}
	executorTips := []*big.Int{big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32())))}
	tips := []types.Tips{
		types.NewTips(notaryTips[0], broadcasterTips[0], proverTips[0], executorTips[0]),
		types.NewTips(notaryTips[1], broadcasterTips[1], proverTips[1], executorTips[1]),
	}
	encodedTips, err := types.EncodeTips(tips[0])
	e.Nil(err)
	messageBytes := []byte{byte(gofakeit.Uint32())}

	transactOpts := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), e.OriginContractMetadata.OwnerPtr())
	transactOpts.Value = types.TotalTips(tips[0])

	tx, err := e.OriginContract.Dispatch(transactOpts.TransactOpts, destination, recipients[0], optimisticSeconds[0], encodedTips, messageBytes)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := e.TestBackendOrigin.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(chainID, sender.Hash(), 1, destination, recipients[0], optimisticSeconds[0])

	message := types.NewMessage(header, tips[0], messageBytes)
	e.Nil(err)

	leafA, err := message.ToLeaf()
	e.Nil(err)
	testTree.Insert(leafA[:])
	testRootA, err := testTree.Root(1)
	e.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.Start(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	// Listen with the exec.
	go func() {
		execErr := exec.Listen(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	waitChan := make(chan bool, 2)

	e.Eventually(func() bool {
		rootA, err := exec.GetMerkleTree(chainID, destination).Root(1)
		if err != nil {
			return false
		}

		// convert testRootA to [32]byte
		var testRootA32 [32]byte
		copy(testRootA32[:], testRootA)

		var rootA32 [32]byte
		copy(rootA32[:], rootA)

		if testRootA32 == rootA32 {
			waitChan <- true
			return true
		}
		return false
	})

	encodedTips, err = types.EncodeTips(tips[1])
	e.Nil(err)

	transactOpts.Value = types.TotalTips(tips[1])

	tx, err = e.OriginContract.Dispatch(transactOpts.TransactOpts, destination, recipients[1], optimisticSeconds[1], encodedTips, messageBytes)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	header = types.NewHeader(chainID, sender.Hash(), 2, destination, recipients[1], optimisticSeconds[1])

	message = types.NewMessage(header, tips[1], messageBytes)
	e.Nil(err)

	leafB, err := message.ToLeaf()
	e.Nil(err)
	testTree.Insert(leafB[:])
	testRootB, err := testTree.Root(2)
	e.Nil(err)

	e.Eventually(func() bool {
		rootB, err := exec.GetMerkleTree(chainID, destination).Root(2)
		if err != nil {
			return false
		}

		// convert testRootB to [32]byte
		var testRootB32 [32]byte
		copy(testRootB32[:], testRootB)

		var rootB32 [32]byte
		copy(rootB32[:], rootB)

		if testRootB32 == rootB32 {
			waitChan <- true
			return true
		}
		return false
	})

	<-waitChan
	<-waitChan
	exec.Stop(chainID)

	oldTreeItems := exec.GetMerkleTree(chainID, destination).Items()

	var newRoot []byte
	e.Eventually(func() bool {
		dbTree, err := executor.NewTreeFromDB(e.GetTestContext(), chainID, destination, e.ExecutorTestDB)
		e.Nil(err)

		exec.OverrideMerkleTree(chainID, destination, dbTree)

		newRoot, err = exec.GetMerkleTree(chainID, destination).Root(2)
		if err != nil {
			return false
		}

		waitChan <- true
		return true
	})
	<-waitChan

	newTreeItems := exec.GetMerkleTree(chainID, destination).Items()

	e.Equal(oldTreeItems, newTreeItems)

	var testRootB32 [32]byte
	copy(testRootB32[:], testRootB)

	var newRoot32 [32]byte
	copy(newRoot32[:], newRoot)

	e.Equal(testRootB32, newRoot32)
}

func (e *ExecutorSuite) TestVerifyMessage() {
	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID: chainID,
			},
			{
				ChainID: destination,
			},
		},
		BaseOmnirpcURL: e.TestBackendOrigin.RPCAddress(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	executorClients := map[uint32]executor.Backend{
		chainID:     nil,
		destination: nil,
	}

	urls := map[uint32]string{
		chainID:     e.TestBackendOrigin.RPCAddress(),
		destination: e.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	nonces := []uint32{1, 2, 3, 4}
	blockNumbers := []uint64{10, 20, 30, 40}
	recipients := [][32]byte{
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
	}
	senders := [][32]byte{
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
	}
	optimisticSeconds := []uint32{
		gofakeit.Uint32(), gofakeit.Uint32(),
		gofakeit.Uint32(), gofakeit.Uint32(),
	}
	notaryTips := []*big.Int{
		big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32()))),
		big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32()))),
	}
	broadcasterTips := []*big.Int{
		big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32()))),
		big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32()))),
	}
	proverTips := []*big.Int{
		big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32()))),
		big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32()))),
	}
	executorTips := []*big.Int{
		big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32()))),
		big.NewInt(int64(int(gofakeit.Uint32()))), big.NewInt(int64(int(gofakeit.Uint32()))),
	}
	tips := []types.Tips{
		types.NewTips(notaryTips[0], broadcasterTips[0], proverTips[0], executorTips[0]),
		types.NewTips(notaryTips[1], broadcasterTips[1], proverTips[1], executorTips[1]),
		types.NewTips(notaryTips[2], broadcasterTips[2], proverTips[2], executorTips[2]),
		types.NewTips(notaryTips[3], broadcasterTips[3], proverTips[3], executorTips[3]),
	}
	messageBytes := [][]byte{
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
	}

	header0 := types.NewHeader(chainID, senders[0], nonces[0], destination, recipients[0], optimisticSeconds[0])
	header1 := types.NewHeader(chainID, senders[1], nonces[1], destination, recipients[1], optimisticSeconds[1])
	header2 := types.NewHeader(chainID, senders[2], nonces[2], destination, recipients[2], optimisticSeconds[2])
	header3 := types.NewHeader(chainID, senders[3], nonces[3], destination, recipients[3], optimisticSeconds[3])

	message0 := types.NewMessage(header0, tips[0], messageBytes[0])
	message1 := types.NewMessage(header1, tips[1], messageBytes[1])
	message2 := types.NewMessage(header2, tips[2], messageBytes[2])
	message3 := types.NewMessage(header3, tips[3], messageBytes[3])
	failMessage := types.NewMessage(header1, tips[3], messageBytes[3])

	// Insert messages into the database.
	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message0, blockNumbers[0], false, 0)
	e.Nil(err)
	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message1, blockNumbers[1], false, 0)
	e.Nil(err)
	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message2, blockNumbers[2], false, 0)
	e.Nil(err)

	dbTree, err := executor.NewTreeFromDB(e.GetTestContext(), chainID, destination, e.ExecutorTestDB)
	e.Nil(err)

	exec.OverrideMerkleTree(chainID, destination, dbTree)

	inTree0, err := exec.VerifyMessageMerkleProof(message0)
	e.Nil(err)
	e.True(inTree0)

	inTree1, err := exec.VerifyMessageMerkleProof(message1)
	e.Nil(err)
	e.True(inTree1)

	inTree2, err := exec.VerifyMessageMerkleProof(message2)
	e.Nil(err)
	e.True(inTree2)

	inTreeFail, err := exec.VerifyMessageMerkleProof(failMessage)
	e.Nil(err)
	e.False(inTreeFail)

	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message3, blockNumbers[3], false, 0)
	e.Nil(err)

	dbTree, err = executor.NewTreeFromDB(e.GetTestContext(), chainID, destination, e.ExecutorTestDB)
	e.Nil(err)

	exec.OverrideMerkleTree(chainID, destination, dbTree)

	inTree3, err := exec.VerifyMessageMerkleProof(message3)
	e.Nil(err)
	e.True(inTree3)
}

func (e *ExecutorSuite) TestVerifyOptimisticPeriod() {
	testDone := false
	defer func() {
		testDone = true
	}()

	originClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress())
	e.Nil(err)
	destinationClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendDestination.RPCAddress())
	e.Nil(err)

	_, testContractRef := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), e.TestBackendDestination)
	testTransactOpts := e.TestBackendDestination.GetTxContext(e.GetTestContext(), nil)

	originConfig := config.ContractConfig{
		Address:    e.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := config.ChainConfig{
		ChainID:               uint32(e.TestBackendOrigin.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{originConfig},
	}
	destinationConfig := config.ContractConfig{
		Address:    e.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := config.ChainConfig{
		ChainID:               uint32(e.TestBackendDestination.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{destinationConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{originChainConfig, destinationChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		uint32(e.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(e.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       uint32(e.TestBackendOrigin.GetChainID()),
				OriginAddress: e.OriginContract.Address().String(),
			},
			{
				ChainID:            uint32(e.TestBackendDestination.GetChainID()),
				DestinationAddress: e.DestinationContract.Address().String(),
			},
		},
		BaseOmnirpcURL: e.TestBackendOrigin.RPCAddress(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		uint32(e.TestBackendOrigin.GetChainID()):      e.TestBackendOrigin,
		uint32(e.TestBackendDestination.GetChainID()): e.TestBackendDestination,
	}

	urls := map[uint32]string{
		uint32(e.TestBackendOrigin.GetChainID()):      e.TestBackendOrigin.RPCAddress(),
		uint32(e.TestBackendDestination.GetChainID()): e.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.Start(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	// Listen with the exec.
	go func() {
		execErr := exec.Listen(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	go func() {
		execErr := exec.SetMinimumTime(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	tips := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips, err := types.EncodeTips(tips)
	e.Nil(err)

	optimisticSeconds := uint32(10)

	recipient := [32]byte{byte(gofakeit.Uint32())}
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}

	txContextOrigin := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), e.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)

	tx, err := e.OriginContract.Dispatch(txContextOrigin.TransactOpts, uint32(e.TestBackendDestination.GetChainID()), recipient, optimisticSeconds, encodedTips, body)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := e.TestBackendOrigin.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(uint32(e.TestBackendOrigin.GetChainID()), sender.Hash(), nonce, uint32(e.TestBackendDestination.GetChainID()), recipient, optimisticSeconds)
	message := types.NewMessage(header, tips, body)

	attestKey := types.AttestationKey{
		Origin:      uint32(e.TestBackendOrigin.GetChainID()),
		Destination: uint32(e.TestBackendDestination.GetChainID()),
		Nonce:       nonce,
	}
	root := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), root)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	e.Nil(err)

	guardSignature, err := e.GuardBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	notarySignature, err := e.NotaryBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	e.Nil(err)

	txContextDestination := e.TestBackendDestination.GetTxContext(e.GetTestContext(), e.DestinationContractMetadata.OwnerPtr())

	tx, err = e.DestinationContract.SubmitAttestation(txContextDestination.TransactOpts, rawSignedAttestation)
	e.Nil(err)
	e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)

	continueChan := make(chan bool, 1)

	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())
	// Wait for message to be stored in the database.
	e.Eventually(func() bool {
		_, err = e.ExecutorTestDB.GetAttestationBlockNumber(e.GetTestContext(), types2.DBAttestation{
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

	returnedNonce, err := exec.VerifyMessageOptimisticPeriod(e.GetTestContext(), message)
	e.Nil(err)
	e.Nil(returnedNonce)

	e.Eventually(func() bool {
		returnedNonce, err = exec.VerifyMessageOptimisticPeriod(e.GetTestContext(), message)
		if err != nil {
			return false
		}
		if returnedNonce != nil {
			return true
		}
		// Need to create a tx and wait for it to be confirmed to continue adding blocks, and therefore
		// increase the `time`.
		testTx, err := testContractRef.IncrementCounter(testTransactOpts.TransactOpts)
		e.Nil(err)
		e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), testTx)
		return false
	})
}

//nolint:maintidx
func (e *ExecutorSuite) TestExecute() {
	testDone := false
	defer func() {
		testDone = true
	}()

	testContractDest, testContractRef := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), e.TestBackendDestination)
	testTransactOpts := e.TestBackendDestination.GetTxContext(e.GetTestContext(), nil)

	originClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress())
	e.Nil(err)
	destinationClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendDestination.RPCAddress())
	e.Nil(err)

	originConfig := config.ContractConfig{
		Address:    e.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := config.ChainConfig{
		ChainID:               uint32(e.TestBackendOrigin.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{originConfig},
	}
	destinationConfig := config.ContractConfig{
		Address:    e.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := config.ChainConfig{
		ChainID:               uint32(e.TestBackendDestination.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{destinationConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{originChainConfig, destinationChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		uint32(e.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(e.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       uint32(e.TestBackendOrigin.GetChainID()),
				OriginAddress: e.OriginContract.Address().String(),
			},
			{
				ChainID:            uint32(e.TestBackendDestination.GetChainID()),
				DestinationAddress: e.DestinationContract.Address().String(),
			},
		},
		BaseOmnirpcURL: gofakeit.URL(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		uint32(e.TestBackendOrigin.GetChainID()):      e.TestBackendOrigin,
		uint32(e.TestBackendDestination.GetChainID()): e.TestBackendDestination,
	}

	urls := map[uint32]string{
		uint32(e.TestBackendOrigin.GetChainID()):      e.TestBackendOrigin.RPCAddress(),
		uint32(e.TestBackendDestination.GetChainID()): e.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.Start(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	// Listen with the exec.
	go func() {
		execErr := exec.Listen(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	go func() {
		execErr := exec.SetMinimumTime(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	tips := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips, err := types.EncodeTips(tips)
	e.Nil(err)

	optimisticSeconds := uint32(10)

	recipient := testContractDest.Address().Hash()
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}

	txContextOrigin := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), e.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)

	tx, err := e.OriginContract.Dispatch(txContextOrigin.TransactOpts, uint32(e.TestBackendDestination.GetChainID()), recipient, optimisticSeconds, encodedTips, body)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := e.TestBackendOrigin.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(uint32(e.TestBackendOrigin.GetChainID()), sender.Hash(), nonce, uint32(e.TestBackendDestination.GetChainID()), recipient, optimisticSeconds)
	message := types.NewMessage(header, tips, body)

	attestKey := types.AttestationKey{
		Origin:      uint32(e.TestBackendOrigin.GetChainID()),
		Destination: uint32(e.TestBackendDestination.GetChainID()),
		Nonce:       nonce,
	}

	tree := merkle.NewTree()

	leaf, err := message.ToLeaf()
	e.Nil(err)

	tree.Insert(leaf[:])

	root, err := tree.Root(1)
	e.Nil(err)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), rootB32)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	e.Nil(err)

	guardSignature, err := e.GuardBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	notarySignature, err := e.NotaryBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	e.Nil(err)

	txContextDestination := e.TestBackendDestination.GetTxContext(e.GetTestContext(), e.DestinationContractMetadata.OwnerPtr())

	tx, err = e.DestinationContract.SubmitAttestation(txContextDestination.TransactOpts, rawSignedAttestation)
	e.Nil(err)
	e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)

	continueChan := make(chan bool, 2)

	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())
	// Wait for message to be stored in the database.
	e.Eventually(func() bool {
		_, err = e.ExecutorTestDB.GetAttestationBlockNumber(e.GetTestContext(), types2.DBAttestation{
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

	executed, err := exec.Execute(e.GetTestContext(), message)
	e.Nil(err)
	e.False(executed)

	e.Eventually(func() bool {
		executed, err := exec.Execute(e.GetTestContext(), message)
		if err != nil {
			return false
		}
		if executed {
			return true
		}
		// Need to create a tx and wait for it to be confirmed to continue adding blocks, and therefore
		// increase the `time`.
		countBeforeIncrement, err := testContractRef.GetCount(&bind.CallOpts{Context: e.GetTestContext()})
		e.Nil(err)
		testTx, err := testContractRef.IncrementCounter(testTransactOpts.TransactOpts)
		e.Nil(err)
		e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), testTx)
		countAfterIncrement, err := testContractRef.GetCount(&bind.CallOpts{Context: e.GetTestContext()})
		e.Nil(err)
		e.Greater(countAfterIncrement.Uint64(), countBeforeIncrement.Uint64())
		return false
	})

	tips2 := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips2, err := types.EncodeTips(tips2)
	e.Nil(err)

	optimisticSeconds2 := uint32(5)

	nonce2 := uint32(2)
	body2 := []byte{byte(gofakeit.Uint32())}

	txContextOrigin.Value = types.TotalTips(tips2)

	tx, err = e.OriginContract.Dispatch(txContextOrigin.TransactOpts, uint32(e.TestBackendDestination.GetChainID()), recipient, optimisticSeconds2, encodedTips2, body2)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	header2 := types.NewHeader(uint32(e.TestBackendOrigin.GetChainID()), sender.Hash(), nonce2, uint32(e.TestBackendDestination.GetChainID()), recipient, optimisticSeconds2)
	message2 := types.NewMessage(header2, tips2, body2)

	attestKey2 := types.AttestationKey{
		Origin:      uint32(e.TestBackendOrigin.GetChainID()),
		Destination: uint32(e.TestBackendDestination.GetChainID()),
		Nonce:       nonce2,
	}

	leaf2, err := message2.ToLeaf()
	e.Nil(err)

	tree.Insert(leaf2[:])

	root2, err := tree.Root(2)
	e.Nil(err)

	var root2B32 [32]byte
	copy(root2B32[:], root2)

	unsignedAttestation2 := types.NewAttestation(attestKey2.GetRawKey(), root2B32)
	hashedAttestation2, err := types.Hash(unsignedAttestation2)
	e.Nil(err)

	guardSignature2, err := e.GuardBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation2), false)
	e.Nil(err)

	notarySignature2, err := e.NotaryBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation2), false)
	e.Nil(err)

	signedAttestation2 := types.NewSignedAttestation(unsignedAttestation2, []types.Signature{guardSignature2}, []types.Signature{notarySignature2})

	rawSignedAttestation2, err := types.EncodeSignedAttestation(signedAttestation2)
	e.Nil(err)

	tx, err = e.DestinationContract.SubmitAttestation(txContextDestination.TransactOpts, rawSignedAttestation2)
	e.Nil(err)
	e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)

	e.Eventually(func() bool {
		_, err = e.ExecutorTestDB.GetAttestationBlockNumber(e.GetTestContext(), types2.DBAttestation{
			ChainID:     &chainID,
			Destination: &destination,
			Nonce:       &nonce2,
		})
		if err == nil {
			continueChan <- true
			return true
		}
		return false
	})

	<-continueChan

	executed, err = exec.Execute(e.GetTestContext(), message2)
	e.Nil(err)
	e.False(executed)

	e.Eventually(func() bool {
		executed, err = exec.Execute(e.GetTestContext(), message2)
		if err != nil {
			return false
		}
		if executed {
			return true
		}
		// Need to create a tx and wait for it to be confirmed to continue adding blocks, and therefore
		// increase the `time`.
		testTx, err := testContractRef.IncrementCounter(testTransactOpts.TransactOpts)
		e.Nil(err)
		e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), testTx)
		return false
	})

	exec.Stop(uint32(e.TestBackendOrigin.GetChainID()))
	exec.Stop(uint32(e.TestBackendDestination.GetChainID()))
}

// TestDestinationExecute test executing on destination.
func (e *ExecutorSuite) TestDestinationExecute() {
	var err error

	testContractDest, testContractDestRef := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), e.TestBackendDestination)

	originDomain := uint32(e.TestBackendOrigin.GetBigChainID().Uint64())
	destinationDomain := uint32(e.TestBackendDestination.GetBigChainID().Uint64())
	nonce := uint32(1)

	origins := []uint32{originDomain}
	nonces := []uint32{nonce}

	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	attestationSink := make(chan *destinationharness.DestinationHarnessAttestationAccepted)
	subAttestation, err := e.DestinationContract.WatchAttestationAccepted(&bind.WatchOpts{
		Context: e.GetTestContext()},
		attestationSink)
	e.Nil(err)

	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	iMessageHandledSink := make(chan *agentstestcontract.AgentsTestContractIMessageReceipientHandleEvent)
	subMessageHandled, err := testContractDestRef.WatchIMessageReceipientHandleEvent(&bind.WatchOpts{
		Context: e.GetTestContext()},
		iMessageHandledSink, origins, nonces)
	e.Nil(err)

	txContextOrigin := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), nil)
	txContextDestination := e.TestBackendDestination.GetTxContext(e.GetTestContext(), e.DestinationContractMetadata.OwnerPtr())

	messageBody := []byte("This is a test message")

	recipient := testContractDest.Address().Hash()

	tips := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips, err := types.EncodeTips(tips)
	e.Nil(err)

	optimisticSeconds := uint32(1)
	// Dispatch an event from the Origin contract to be accepted on the Destination contract.
	tx, err := e.OriginContract.Dispatch(txContextOrigin.TransactOpts, destinationDomain, recipient, optimisticSeconds, encodedTips, messageBody)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := e.TestBackendOrigin.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(originDomain, sender.Hash(), nonce, destinationDomain, recipient, optimisticSeconds)

	message := types.NewMessage(header, tips, messageBody)
	e.Nil(err)
	encodedMessage, err := types.EncodeMessage(message)
	e.Nil(err)
	allMessages := []types.Message{message}
	rawMessages := make([][]byte, len(allMessages))
	for i, message := range allMessages {
		rawMessage, err := message.ToLeaf()
		e.Nil(err)

		rawMessages[i] = rawMessage[:]
	}

	historicalMerkleTree := merkle.NewTreeFromItems(rawMessages)

	rawProof, err := historicalMerkleTree.MerkleProof(0, 1)
	e.Nil(err)
	var proofToUse [32][32]byte
	for i := 0; i < int(merkle.TreeDepth); i++ {
		copy(proofToUse[i][:], rawProof[i][:32])
	}

	attestationKey := types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain,
		Nonce:       nonce,
	}

	rawRoot, err := historicalMerkleTree.Root(1)
	e.Nil(err)
	var root [32]byte
	copy(root[:], rawRoot[:32])

	unsignedAttestation := types.NewAttestation(attestationKey.GetRawKey(), root)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	e.Nil(err)

	notarySignature, err := e.NotaryBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	guardSignature, err := e.GuardBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	e.Nil(err)

	tx, err = e.DestinationContract.SubmitAttestation(txContextDestination.TransactOpts, rawSignedAttestation)
	e.Nil(err)

	e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)

	watchCtx, cancel := context.WithTimeout(e.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		e.T().Error(e.T(), fmt.Errorf("test context completed %w", e.GetTestContext().Err()))
	case <-subAttestation.Err():
		e.T().Error(e.T(), subAttestation.Err())
	// get attestation accepted event
	case item := <-attestationSink:
		parser, err := destination.NewParser(e.DestinationContract.Address())
		e.Nil(err)

		// Check to see if the event was an AttestationAccepted event.
		eventType, ok := parser.EventType(item.Raw)
		e.True(ok)
		e.Equal(eventType, destination.AttestationAcceptedEvent)

		emittedSignedAttestation, err := types.DecodeSignedAttestation(item.Attestation)
		e.Nil(err)

		e.Equal(e.OriginDomainClient.Config().DomainID, emittedSignedAttestation.Attestation().Origin())
		e.Equal(e.DestinationDomainClient.Config().DomainID, emittedSignedAttestation.Attestation().Destination())
		e.Equal(nonce, emittedSignedAttestation.Attestation().Nonce())
		e.Equal(root, emittedSignedAttestation.Attestation().Root())

		// Now sleep for a second before executing
		time.Sleep(time.Second)
		index := big.NewInt(int64(0))

		tx, err = e.DestinationContract.Execute(txContextDestination.TransactOpts, encodedMessage, proofToUse, index)
		e.Nil(err)

		e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)

		watchHandleCtx, cancel := context.WithTimeout(e.GetTestContext(), time.Second*10)
		defer cancel()

		select {
		// check for errors and fail
		case <-watchHandleCtx.Done():
			e.T().Error(e.T(), fmt.Errorf("test context completed %w", e.GetTestContext().Err()))
		case <-subMessageHandled.Err():
			e.T().Error(e.T(), subMessageHandled.Err())
		// get attestation accepted event
		case item := <-iMessageHandledSink:
			e.NotNil(item)

			break
		}

		break
	}
}

func (e *ExecutorSuite) TestDestinationBadProofExecute() {
	var err error

	testContractDest, testContractDestRef := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), e.TestBackendDestination)

	originDomain := uint32(e.TestBackendOrigin.GetBigChainID().Uint64())
	destinationDomain := uint32(e.TestBackendDestination.GetBigChainID().Uint64())
	nonce := uint32(1)

	origins := []uint32{originDomain}
	nonces := []uint32{nonce}

	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	attestationSink := make(chan *destinationharness.DestinationHarnessAttestationAccepted)
	subAttestation, err := e.DestinationContract.WatchAttestationAccepted(&bind.WatchOpts{
		Context: e.GetTestContext()},
		attestationSink)
	e.Nil(err)

	// Create a channel and subscription to receive AttestationAccepted events as they are emitted.
	iMessageHandledSink := make(chan *agentstestcontract.AgentsTestContractIMessageReceipientHandleEvent)
	subMessageHandled, err := testContractDestRef.WatchIMessageReceipientHandleEvent(&bind.WatchOpts{
		Context: e.GetTestContext()},
		iMessageHandledSink, origins, nonces)
	e.Nil(err)

	txContextOrigin := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), nil)
	txContextDestination := e.TestBackendDestination.GetTxContext(e.GetTestContext(), e.DestinationContractMetadata.OwnerPtr())

	messageBody := []byte("This is a test message")

	recipient := testContractDest.Address().Hash()

	tips := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips, err := types.EncodeTips(tips)
	e.Nil(err)

	optimisticSeconds := uint32(1)
	// Dispatch an event from the Origin contract to be accepted on the Destination contract.
	tx, err := e.OriginContract.Dispatch(txContextOrigin.TransactOpts, destinationDomain, recipient, optimisticSeconds, encodedTips, messageBody)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := e.TestBackendOrigin.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(originDomain, sender.Hash(), nonce, destinationDomain, recipient, optimisticSeconds)

	message := types.NewMessage(header, tips, messageBody)
	e.Nil(err)
	encodedMessage, err := types.EncodeMessage(message)
	e.Nil(err)
	allMessages := []types.Message{message}
	rawMessages := make([][]byte, len(allMessages))
	for i, message := range allMessages {
		rawMessage, err := message.ToLeaf()
		e.Nil(err)

		rawMessages[i] = rawMessage[:]
	}

	historicalMerkleTree := merkle.NewTreeFromItems(rawMessages)

	_, err = historicalMerkleTree.MerkleProof(0, 1)
	e.Nil(err)
	var badProofToUse [32][32]byte
	for i := 0; i < int(merkle.TreeDepth); i++ {
		for j := 0; j < int(merkle.TreeDepth); j++ {
			badProofToUse[i][j] = 1
		}
	}

	attestationKey := types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain,
		Nonce:       nonce,
	}

	rawRoot, err := historicalMerkleTree.Root(1)
	e.Nil(err)
	var root [32]byte
	copy(root[:], rawRoot[:32])

	unsignedAttestation := types.NewAttestation(attestationKey.GetRawKey(), root)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	e.Nil(err)

	notarySignature, err := e.NotaryBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	guardSignature, err := e.GuardBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	e.Nil(err)

	tx, err = e.DestinationContract.SubmitAttestation(txContextDestination.TransactOpts, rawSignedAttestation)
	e.Nil(err)

	e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)

	watchCtx, cancel := context.WithTimeout(e.GetTestContext(), time.Second*10)
	defer cancel()

	select {
	// check for errors and fail
	case <-watchCtx.Done():
		e.T().Error(e.T(), fmt.Errorf("test context completed %w", e.GetTestContext().Err()))
	case <-subAttestation.Err():
		e.T().Error(e.T(), subAttestation.Err())
	// get attestation accepted event
	case item := <-attestationSink:
		parser, err := destination.NewParser(e.DestinationContract.Address())
		e.Nil(err)

		// Check to see if the event was an AttestationAccepted event.
		eventType, ok := parser.EventType(item.Raw)
		e.True(ok)
		e.Equal(eventType, destination.AttestationAcceptedEvent)

		emittedSignedAttesation, err := types.DecodeSignedAttestation(item.Attestation)
		e.Nil(err)

		e.Equal(e.OriginDomainClient.Config().DomainID, emittedSignedAttesation.Attestation().Origin())
		e.Equal(e.DestinationDomainClient.Config().DomainID, emittedSignedAttesation.Attestation().Destination())
		e.Equal(nonce, emittedSignedAttesation.Attestation().Nonce())
		e.Equal(root, emittedSignedAttesation.Attestation().Root())

		// Now sleep for a second before executing
		time.Sleep(time.Second)
		index := big.NewInt(int64(0))

		tx, err = e.DestinationContract.Execute(txContextDestination.TransactOpts, encodedMessage, badProofToUse, index)
		e.Nil(err)

		e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)

		watchHandleCtx, cancel := context.WithTimeout(e.GetTestContext(), time.Second*10)
		defer cancel()

		didFailAsExpected := false
		select {
		// check for errors and fail
		case <-watchHandleCtx.Done():
			didFailAsExpected = true
		case <-subMessageHandled.Err():
			e.FailNow("Should not be here")
		// get attestation accepted event
		case <-iMessageHandledSink:
			e.FailNow("Should not be here")

			break
		}
		e.True(didFailAsExpected)

		break
	}
}

func (e *ExecutorSuite) TestSetMinimumTime() {
	// Put some messages without a minimum time in the database.
	originDomain := uint32(e.TestBackendOrigin.GetBigChainID().Uint64())
	destinationDomain := uint32(e.TestBackendDestination.GetBigChainID().Uint64())

	for i := uint32(1); i <= 5; i++ {
		header := types.NewHeader(originDomain, common.BigToHash(big.NewInt(gofakeit.Int64())), i, destinationDomain, common.BigToHash(big.NewInt(gofakeit.Int64())), i*2)
		tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		message := types.NewMessage(header, tips, []byte{})

		err := e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message, uint64(i), false, 0)
		e.Nil(err)
	}

	// Sanity check that the messages are in the database.
	messageMask := types2.DBMessage{
		ChainID:     &originDomain,
		Destination: &destinationDomain,
	}
	messages, err := e.ExecutorTestDB.GetMessages(e.GetTestContext(), messageMask, 1)
	e.Nil(err)
	e.Len(messages, 5)

	originClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress())
	e.Nil(err)
	destinationClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendDestination.RPCAddress())
	e.Nil(err)

	originConfig := config.ContractConfig{
		Address:    e.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := config.ChainConfig{
		ChainID:               uint32(e.TestBackendOrigin.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{originConfig},
	}
	destinationConfig := config.ContractConfig{
		Address:    e.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := config.ChainConfig{
		ChainID:               uint32(e.TestBackendDestination.GetChainID()),
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{destinationConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{originChainConfig, destinationChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		uint32(e.TestBackendOrigin.GetChainID()):      {originClient, originClient},
		uint32(e.TestBackendDestination.GetChainID()): {destinationClient, destinationClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       uint32(e.TestBackendOrigin.GetChainID()),
				OriginAddress: e.OriginContract.Address().String(),
			},
			{
				ChainID:            uint32(e.TestBackendDestination.GetChainID()),
				DestinationAddress: e.DestinationContract.Address().String(),
			},
		},
		BaseOmnirpcURL: e.TestBackendOrigin.RPCAddress(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		uint32(e.TestBackendOrigin.GetChainID()):      e.TestBackendOrigin,
		uint32(e.TestBackendDestination.GetChainID()): e.TestBackendDestination,
	}

	urls := map[uint32]string{
		uint32(e.TestBackendOrigin.GetChainID()):      e.TestBackendOrigin.RPCAddress(),
		uint32(e.TestBackendDestination.GetChainID()): e.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	go func() {
		_ = exec.SetMinimumTime(e.GetTestContext())
	}()

	// Inject one attestation into the database for nonce 1.
	attestKey := types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain,
		Nonce:       1,
	}
	root := common.BigToHash(big.NewInt(gofakeit.Int64()))
	attestation := types.NewAttestation(attestKey.GetRawKey(), root)
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation, 1, 10)
	e.Nil(err)

	// Store a dud attestation for nonce 1 on a different destination.
	attestKey = types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain + 1,
		Nonce:       1,
	}
	root = common.BigToHash(big.NewInt(gofakeit.Int64()))
	attestation = types.NewAttestation(attestKey.GetRawKey(), root)
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation, 1, 0)
	e.Nil(err)

	waitChan := make(chan bool, 1)

	e.Eventually(func() bool {
		trueVal := true
		minimumTime := uint64(10 + 2) // attestation block time + optimistic period of nonce 1 message
		retrieveMessageMask := types2.DBMessage{
			MinimumTimeSet: &trueVal,
			MinimumTime:    &minimumTime,
		}
		messages, err := e.ExecutorTestDB.GetMessages(e.GetTestContext(), retrieveMessageMask, 1)
		e.Nil(err)
		if len(messages) != 1 {
			return false
		}
		waitChan <- true
		return true
	})
	<-waitChan

	// Inject two attestations into the database. One for nonce 2, and one for nonce 4.
	attestKey = types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain,
		Nonce:       2,
	}
	root = common.BigToHash(big.NewInt(gofakeit.Int64()))
	attestation = types.NewAttestation(attestKey.GetRawKey(), root)
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation, 2, 20)
	e.Nil(err)

	attestKey = types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain,
		Nonce:       4,
	}
	root = common.BigToHash(big.NewInt(gofakeit.Int64()))
	attestation = types.NewAttestation(attestKey.GetRawKey(), root)
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation, 4, 40)
	e.Nil(err)

	e.Eventually(func() bool {
		trueVal := true
		retrieveMessageMask := types2.DBMessage{
			MinimumTimeSet: &trueVal,
		}
		messages, err := e.ExecutorTestDB.GetMessages(e.GetTestContext(), retrieveMessageMask, 1)
		e.Nil(err)
		if len(messages) != 4 {
			return false
		}
		e.Equal(uint32(2), messages[1].Nonce())
		e.Equal(uint32(3), messages[2].Nonce())
		e.Equal(uint32(4), messages[3].Nonce())

		waitChan <- true
		return true
	})
	<-waitChan
}

func (e *ExecutorSuite) TestExecuteExecutable() {
	testDone := false
	defer func() {
		testDone = true
	}()

	originDomain := uint32(e.TestBackendOrigin.GetBigChainID().Uint64())
	destinationDomain := uint32(e.TestBackendDestination.GetBigChainID().Uint64())

	testContractDest, testContractRef := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), e.TestBackendDestination)
	testTransactOpts := e.TestBackendDestination.GetTxContext(e.GetTestContext(), nil)

	originClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress())
	e.Nil(err)
	destinationClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendDestination.RPCAddress())
	e.Nil(err)

	originConfig := config.ContractConfig{
		Address:    e.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := config.ChainConfig{
		ChainID:               originDomain,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{originConfig},
	}
	destinationConfig := config.ContractConfig{
		Address:    e.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := config.ChainConfig{
		ChainID:               destinationDomain,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{destinationConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{originChainConfig, destinationChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		originDomain:      {originClient, originClient},
		destinationDomain: {destinationClient, destinationClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       originDomain,
				OriginAddress: e.OriginContract.Address().String(),
			},
			{
				ChainID:            destinationDomain,
				DestinationAddress: e.DestinationContract.Address().String(),
			},
		},
		BaseOmnirpcURL: gofakeit.URL(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		originDomain:      e.TestBackendOrigin,
		destinationDomain: e.TestBackendDestination,
	}

	urls := map[uint32]string{
		originDomain:      e.TestBackendOrigin.RPCAddress(),
		destinationDomain: e.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.Start(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	// Listen with the exec.
	go func() {
		execErr := exec.Listen(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	go func() {
		execErr := exec.SetMinimumTime(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	go func() {
		execErr := exec.ExecuteExecutable(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	tips := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips, err := types.EncodeTips(tips)
	e.Nil(err)

	optimisticSeconds := uint32(10)

	recipient := testContractDest.Address().Hash()
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}

	txContextOrigin := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), e.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)

	tx, err := e.OriginContract.Dispatch(txContextOrigin.TransactOpts, destinationDomain, recipient, optimisticSeconds, encodedTips, body)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := e.TestBackendOrigin.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(originDomain, sender.Hash(), nonce, destinationDomain, recipient, optimisticSeconds)
	message := types.NewMessage(header, tips, body)

	attestKey := types.AttestationKey{
		Origin:      originDomain,
		Destination: destinationDomain,
		Nonce:       uint32(1),
	}

	tree := merkle.NewTree()

	leaf, err := message.ToLeaf()
	e.Nil(err)

	tree.Insert(leaf[:])

	root, err := tree.Root(1)
	e.Nil(err)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	unsignedAttestation := types.NewAttestation(attestKey.GetRawKey(), rootB32)
	hashedAttestation, err := types.Hash(unsignedAttestation)
	e.Nil(err)

	guardSignature, err := e.GuardBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	notarySignature, err := e.NotaryBondedSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	signedAttestation := types.NewSignedAttestation(unsignedAttestation, []types.Signature{guardSignature}, []types.Signature{notarySignature})

	rawSignedAttestation, err := types.EncodeSignedAttestation(signedAttestation)
	e.Nil(err)

	txContextDestination := e.TestBackendDestination.GetTxContext(e.GetTestContext(), e.DestinationContractMetadata.OwnerPtr())

	tx, err = e.DestinationContract.SubmitAttestation(txContextDestination.TransactOpts, rawSignedAttestation)
	e.Nil(err)
	e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)

	waitChan := make(chan bool, 1)

	// Make sure there is one executable message in the database.
	e.Eventually(func() bool {
		mask := types2.DBMessage{
			ChainID: &originDomain,
		}
		executableMessages, err := e.ExecutorTestDB.GetExecutableMessages(e.GetTestContext(), mask, uint64(time.Now().Unix()), 1)
		e.Nil(err)
		if len(executableMessages) == 0 {
			return false
		}
		waitChan <- true
		return true
	})

	<-waitChan

	// Ensure that the message that was previously executable in the database has been executed.
	e.Eventually(func() bool {
		mask := types2.DBMessage{
			ChainID: &originDomain,
		}
		executableMessages, err := e.ExecutorTestDB.GetExecutableMessages(e.GetTestContext(), mask, uint64(time.Now().Unix()), 1)
		e.Nil(err)
		if len(executableMessages) == 0 {
			return true
		}
		if len(executableMessages) > 1 {
			e.Fail("too many executable messages")
		}
		// Need to create a tx and wait for it to be confirmed to continue adding blocks, and therefore
		// increase the `time`.
		countBeforeIncrement, err := testContractRef.GetCount(&bind.CallOpts{Context: e.GetTestContext()})
		e.Nil(err)
		testTx, err := testContractRef.IncrementCounter(testTransactOpts.TransactOpts)
		e.Nil(err)
		e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), testTx)
		countAfterIncrement, err := testContractRef.GetCount(&bind.CallOpts{Context: e.GetTestContext()})
		e.Nil(err)
		e.Greater(countAfterIncrement.Uint64(), countBeforeIncrement.Uint64())
		return false
	})
}

func (e *ExecutorSuite) TestSetMinimumTimes() {
	chainID := uint32(e.TestBackendOrigin.GetBigChainID().Uint64())
	destination := uint32(e.TestBackendDestination.GetBigChainID().Uint64())

	originClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress())
	e.Nil(err)
	destinationClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendDestination.RPCAddress())
	e.Nil(err)

	originConfig := config.ContractConfig{
		Address:    e.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := config.ChainConfig{
		ChainID:               chainID,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{originConfig},
	}
	destinationConfig := config.ContractConfig{
		Address:    e.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := config.ChainConfig{
		ChainID:               destination,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{destinationConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{originChainConfig, destinationChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		chainID:     {originClient, originClient},
		destination: {destinationClient, destinationClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetTestContext())
	}()

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       chainID,
				OriginAddress: e.OriginContract.Address().String(),
			},
			{
				ChainID:            destination,
				DestinationAddress: e.DestinationContract.Address().String(),
			},
		},
		BaseOmnirpcURL: gofakeit.URL(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(e.T(), "", e.ExecutorUnbondedWallet.PrivateKeyHex()).Name(),
		},
	}

	executorClients := map[uint32]executor.Backend{
		chainID:     e.TestBackendOrigin,
		destination: e.TestBackendDestination,
	}

	urls := map[uint32]string{
		chainID:     e.TestBackendOrigin.RPCAddress(),
		destination: e.TestBackendDestination.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	message := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()

	for i := 1; i <= 6; i++ {
		header := types.NewHeader(chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), uint32(i), destination, common.BigToHash(big.NewInt(gofakeit.Int64())), 0)
		tips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		typesMessage := types.NewMessage(header, tips, message)

		err := e.ExecutorTestDB.StoreMessage(e.GetTestContext(), typesMessage, uint64(i), false, 0)
		e.Nil(err)
	}

	// This test tests 4 main cases:
	// 1. The message has an associated attestation with equal nonce (and lowest block number).
	// 2. The message has an associated attestation with greater nonce (and lowest block number).
	// 3. The message has an associated attestation with greater nonce (with an attestation of equal nonce, but greater block number).
	// 4. The message has no associated attestation.

	// Case 1
	blockNumber := uint64(1)
	attestKey := types.AttestationKey{
		Origin:      chainID,
		Destination: destination,
		Nonce:       1,
	}
	root := common.BigToHash(big.NewInt(gofakeit.Int64()))
	attestation := types.NewAttestation(attestKey.GetRawKey(), root)

	// Attestation with nonce 1 should be used by message 1.
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation, blockNumber, blockNumber)
	e.Nil(err)

	// Case 2
	blockNumber = uint64(3)
	attestKey = types.AttestationKey{
		Origin:      chainID,
		Destination: destination,
		Nonce:       3,
	}
	root = common.BigToHash(big.NewInt(gofakeit.Int64()))
	attestation = types.NewAttestation(attestKey.GetRawKey(), root)

	// Attestation with nonce 3 should be used by messages 2 and 3.
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation, blockNumber, blockNumber)
	e.Nil(err)

	// Case 3
	blockNumber = uint64(4)
	attestKey = types.AttestationKey{
		Origin:      chainID,
		Destination: destination,
		Nonce:       5,
	}
	root = common.BigToHash(big.NewInt(gofakeit.Int64()))
	attestation = types.NewAttestation(attestKey.GetRawKey(), root)

	// Attestation with nonce 5 should be used by messages 4 and 5.
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation, blockNumber, blockNumber)
	e.Nil(err)

	blockNumber = uint64(5)
	attestKey = types.AttestationKey{
		Origin:      chainID,
		Destination: destination,
		Nonce:       4,
	}
	root = common.BigToHash(big.NewInt(gofakeit.Int64()))
	attestation = types.NewAttestation(attestKey.GetRawKey(), root)

	// Attestation with nonce 4 should not be used by any messages.
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation, blockNumber, blockNumber)
	e.Nil(err)

	// Get the unset messages and attestations.
	messageMask := types2.DBMessage{
		ChainID: &chainID,
	}

	messages, err := e.ExecutorTestDB.GetUnsetMinimumTimeMessages(e.GetTestContext(), messageMask, 1)
	e.Nil(err)

	e.Len(messages, 6)

	minNonce := messages[0].Nonce()
	attestationMask := types2.DBAttestation{
		ChainID: &chainID,
	}
	attestations, err := e.ExecutorTestDB.GetAttestationsAboveOrEqualNonce(e.GetTestContext(), attestationMask, minNonce, 1)
	e.Nil(err)

	e.Len(attestations, 4)

	// Set the messages times.
	err = exec.SetMinimumTimes(e.GetTestContext(), messages, attestations)
	e.Nil(err)

	blockNumber = uint64(1)
	messageMask = types2.DBMessage{
		ChainID:     &chainID,
		BlockNumber: &blockNumber,
	}
	minTime, err := e.ExecutorTestDB.GetMessageMinimumTime(e.GetTestContext(), messageMask)
	e.Nil(err)
	e.Equal(uint64(1), *minTime)

	blockNumber = uint64(2)
	messageMask = types2.DBMessage{
		ChainID:     &chainID,
		BlockNumber: &blockNumber,
	}
	minTime, err = e.ExecutorTestDB.GetMessageMinimumTime(e.GetTestContext(), messageMask)
	e.Nil(err)
	e.Equal(uint64(3), *minTime)

	blockNumber = uint64(3)
	messageMask = types2.DBMessage{
		ChainID:     &chainID,
		BlockNumber: &blockNumber,
	}
	minTime, err = e.ExecutorTestDB.GetMessageMinimumTime(e.GetTestContext(), messageMask)
	e.Nil(err)
	e.Equal(uint64(3), *minTime)

	blockNumber = uint64(4)
	messageMask = types2.DBMessage{
		ChainID:     &chainID,
		BlockNumber: &blockNumber,
	}
	minTime, err = e.ExecutorTestDB.GetMessageMinimumTime(e.GetTestContext(), messageMask)
	e.Nil(err)
	e.Equal(uint64(4), *minTime)

	blockNumber = uint64(5)
	messageMask = types2.DBMessage{
		ChainID:     &chainID,
		BlockNumber: &blockNumber,
	}
	minTime, err = e.ExecutorTestDB.GetMessageMinimumTime(e.GetTestContext(), messageMask)
	e.Nil(err)
	e.Equal(uint64(4), *minTime)

	blockNumber = uint64(6)
	messageMask = types2.DBMessage{
		ChainID:     &chainID,
		BlockNumber: &blockNumber,
	}
	minTime, err = e.ExecutorTestDB.GetMessageMinimumTime(e.GetTestContext(), messageMask)
	e.Nil(err)
	e.Nil(minTime)
}
