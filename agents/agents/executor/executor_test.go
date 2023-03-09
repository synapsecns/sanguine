package executor_test

import (
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	agentsConfig "github.com/synapsecns/sanguine/agents/config"
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
	// TODO (joeallen): FIX ME
	e.T().Skip()
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
		excErr := exc.Start(e.GetTestContext())
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
	// TODO (joeallen): FIX ME
	e.T().Skip()
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
	// TODO (joeallen): FIX ME
	e.T().Skip()
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

	_, err = exec.GetMerkleTree(chainID).Root(1)
	e.NotNil(err)

	testTree := merkle.NewTree(merkle.MessageTreeDepth)

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
		rootA, err := exec.GetMerkleTree(chainID).Root(1)
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
		rootB, err := exec.GetMerkleTree(chainID).Root(2)
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

	oldTreeItems := exec.GetMerkleTree(chainID).Items()

	var newRoot []byte
	e.Eventually(func() bool {
		dbTree, err := executor.NewTreeFromDB(e.GetTestContext(), chainID, e.ExecutorTestDB)
		e.Nil(err)

		exec.OverrideMerkleTree(chainID, dbTree)

		newRoot, err = exec.GetMerkleTree(chainID).Root(2)
		if err != nil {
			time.Sleep(2 * time.Second)
			return false
		}

		waitChan <- true
		return true
	})
	<-waitChan

	newTreeItems := exec.GetMerkleTree(chainID).Items()

	e.Equal(oldTreeItems, newTreeItems)

	var testRootB32 [32]byte
	copy(testRootB32[:], testRootB)

	var newRoot32 [32]byte
	copy(newRoot32[:], newRoot)

	e.Equal(testRootB32, newRoot32)
}

func (e *ExecutorSuite) TestVerifyMessage() {
	// TODO (joeallen): FIX ME
	e.T().Skip()
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

	dbTree, err := executor.NewTreeFromDB(e.GetTestContext(), chainID, e.ExecutorTestDB)
	e.Nil(err)

	exec.OverrideMerkleTree(chainID, dbTree)

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

	dbTree, err = executor.NewTreeFromDB(e.GetTestContext(), chainID, e.ExecutorTestDB)
	e.Nil(err)

	exec.OverrideMerkleTree(chainID, dbTree)

	inTree3, err := exec.VerifyMessageMerkleProof(message3)
	e.Nil(err)
	e.True(inTree3)
}
