package executor_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/prysmaticlabs/prysm/shared/trieutil"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	executorCfg "github.com/synapsecns/sanguine/agents/agents/executor/config"
	types2 "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
)

func (e *ExecutorSuite) TestExecutor() {
	// TODO (joe): re-enable this later after scribe is updated with fix
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
	simulatedChainA.FundAccount(e.GetTestContext(), e.wallet.Address(), *big.NewInt(params.Ether))
	simulatedChainB.FundAccount(e.GetTestContext(), e.wallet.Address(), *big.NewInt(params.Ether))
	testContractA, testRefA := e.manager.GetTestContract(e.GetTestContext(), simulatedChainA)
	testContractB, testRefB := e.manager.GetTestContract(e.GetTestContext(), simulatedChainB)
	transactOptsA := simulatedChainA.GetTxContext(e.GetTestContext(), nil)
	transactOptsB := simulatedChainB.GetTxContext(e.GetTestContext(), nil)

	// Emit two events on each chain.
	tx, err := testRefA.EmitEventA(transactOptsA.TransactOpts, big.NewInt(1), big.NewInt(2), big.NewInt(3))
	e.Nil(err)
	simulatedChainA.WaitForConfirmation(e.GetTestContext(), tx)
	tx, err = testRefA.EmitEventB(transactOptsA.TransactOpts, []byte{4}, big.NewInt(5), big.NewInt(6))
	e.Nil(err)
	simulatedChainA.WaitForConfirmation(e.GetTestContext(), tx)
	tx, err = testRefB.EmitEventAandB(transactOptsB.TransactOpts, big.NewInt(7), big.NewInt(8), big.NewInt(9))
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

	scribe, err := node.NewScribe(e.scribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.dbPath)

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
		AttestationCollectorChainID: gofakeit.Uint32(),
	}

	exc, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient)
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
	// TODO (joe): re-enable this later after scribe is updated with fix
	e.T().Skip()
	testDone := false
	defer func() {
		testDone = true
	}()
	chainID := gofakeit.Uint32()
	simulatedChain := geth.NewEmbeddedBackendForChainID(e.GetTestContext(), e.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(e.GetTestContext(), simulatedChain.RPCAddress())
	e.Nil(err)
	simulatedChain.FundAccount(e.GetTestContext(), e.wallet.Address(), *big.NewInt(params.Ether))
	testContract, testRef := e.manager.GetTestContract(e.GetTestContext(), simulatedChain)
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

	scribe, err := node.NewScribe(e.scribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.dbPath)
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
	}

	exec, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient)
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
			tx, err := testRef.EmitEventB(transactOpts.TransactOpts, []byte{byte(i)}, big.NewInt(int64(i)), big.NewInt(int64(i)))
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
	// TODO (joe): re-enable this later after scribe is updated with fix
	e.T().Skip()
	testDone := false
	defer func() {
		testDone = true
	}()
	chainID := gofakeit.Uint32()
	deployManager := testutil.NewDeployManager(e.T())
	simulatedChain := geth.NewEmbeddedBackendForChainID(e.GetTestContext(), e.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(e.GetTestContext(), simulatedChain.RPCAddress())
	e.Nil(err)
	simulatedChain.FundAccount(e.GetTestContext(), e.wallet.Address(), *big.NewInt(params.Ether))
	originContract, originRef := deployManager.GetOriginHarness(e.GetTestContext(), simulatedChain)

	contractConfig := config.ContractConfig{
		Address:    originContract.Address().String(),
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

	scribe, err := node.NewScribe(e.scribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.dbPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		_ = scribe.Start(e.GetTestContext())
	}()

	destination := chainID + 1

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID:       chainID,
				OriginAddress: originContract.Address().String(),
			},
			{
				ChainID: destination,
			},
		},
	}

	exec, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient)
	e.Nil(err)

	_, err = exec.GetRoot(e.GetTestContext(), 1, chainID, destination)
	e.NotNil(err)

	testTree, err := trieutil.NewTrie(32)
	e.Nil(err)

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

	ownerPtr, err := originRef.OriginHarnessCaller.Owner(&bind.CallOpts{Context: e.GetTestContext()})
	e.Nil(err)

	transactOpts := simulatedChain.GetTxContext(e.GetTestContext(), &ownerPtr)

	tx, err := originRef.AddAgent(transactOpts.TransactOpts, destination, e.signer.Address())
	e.Nil(err)
	simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)

	transactOpts.Value = types.TotalTips(tips[0])

	tx, err = originRef.Dispatch(transactOpts.TransactOpts, destination, recipients[0], optimisticSeconds[0], encodedTips, messageBytes)
	e.Nil(err)
	simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := simulatedChain.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(chainID, sender.Hash(), 1, destination, recipients[0], optimisticSeconds[0])

	message := types.NewMessage(header, tips[0], messageBytes)
	e.Nil(err)

	leaf, err := message.ToLeaf()
	e.Nil(err)
	testTree.Insert(leaf[:], 0)
	testRootA := testTree.Root()

	// Start the exec.
	go func() {
		execErr := exec.Start(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	// Listen with the exec.
	go func() {
		execErr := exec.Listen(e.GetTestContext(), chainID)
		if !testDone {
			e.Nil(execErr)
		}
	}()

	waitChan := make(chan bool, 2)

	e.Eventually(func() bool {
		rootA, err := exec.GetRoot(e.GetTestContext(), 1, chainID, destination)
		if err != nil {
			return false
		}

		if testRootA == rootA {
			waitChan <- true
			return true
		}
		return false
	})

	encodedTips, err = types.EncodeTips(tips[1])
	e.Nil(err)

	transactOpts.Value = types.TotalTips(tips[1])

	tx, err = originRef.Dispatch(transactOpts.TransactOpts, destination, recipients[1], optimisticSeconds[1], encodedTips, messageBytes)
	e.Nil(err)
	simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)

	header = types.NewHeader(chainID, sender.Hash(), 2, destination, recipients[1], optimisticSeconds[1])

	message = types.NewMessage(header, tips[1], messageBytes)
	e.Nil(err)

	leaf, err = message.ToLeaf()
	e.Nil(err)
	testTree.Insert(leaf[:], 1)

	testRootB := testTree.Root()

	e.Eventually(func() bool {
		rootB, err := exec.GetRoot(e.GetTestContext(), 2, chainID, destination)
		if err != nil {
			return false
		}

		if testRootB == rootB {
			waitChan <- true
			return true
		}
		return false
	})

	<-waitChan
	<-waitChan
	exec.Stop(chainID)

	oldTreeItems := exec.GetMerkleTree(chainID, destination).Items()

	err = exec.BuildTreeFromDB(e.GetTestContext(), chainID, destination)
	e.Nil(err)

	newRoot, err := exec.GetRoot(e.GetTestContext(), 2, chainID, destination)
	e.Nil(err)

	newTreeItems := exec.GetMerkleTree(chainID, destination).Items()

	e.Equal(testRootB, newRoot)
	e.Equal(oldTreeItems, newTreeItems)
}

func (e *ExecutorSuite) TestVerifyMessage() {
	chainID := gofakeit.Uint32()
	destination := chainID + 1

	testTree, err := trieutil.NewTrie(32)
	e.Nil(err)

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID: chainID,
			},
			{
				ChainID: destination,
			},
		},
	}

	scribeClient := client.NewEmbeddedScribe("sqlite", e.dbPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	exec, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient)
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

	leaf0, err := message0.ToLeaf()
	e.Nil(err)
	leaf1, err := message1.ToLeaf()
	e.Nil(err)
	leaf2, err := message2.ToLeaf()
	e.Nil(err)
	leaf3, err := message3.ToLeaf()
	e.Nil(err)

	testTree.Insert(leaf0[:], 0)
	root0 := testTree.Root()
	testTree.Insert(leaf1[:], 1)
	root1 := testTree.Root()
	testTree.Insert(leaf2[:], 2)
	root2 := testTree.Root()
	testTree.Insert(leaf3[:], 3)
	root3 := testTree.Root()

	// Insert messages into the database.
	err = e.testDB.StoreMessage(e.GetTestContext(), message0, root0, blockNumbers[0])
	e.Nil(err)
	err = e.testDB.StoreMessage(e.GetTestContext(), message1, root1, blockNumbers[1])
	e.Nil(err)
	err = e.testDB.StoreMessage(e.GetTestContext(), message2, root2, blockNumbers[2])
	e.Nil(err)

	err = exec.BuildTreeFromDB(e.GetTestContext(), chainID, destination)
	e.Nil(err)

	inTree0, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[0], message0, chainID, destination)
	e.Nil(err)
	e.True(inTree0)

	inTree1, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[1], message1, chainID, destination)
	e.Nil(err)
	e.True(inTree1)

	inTree2, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[2], message2, chainID, destination)
	e.Nil(err)
	e.True(inTree2)

	inTreeFail, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[2], message1, chainID, destination)
	e.Nil(err)
	e.False(inTreeFail)

	err = e.testDB.StoreMessage(e.GetTestContext(), message3, root3, blockNumbers[3])
	e.Nil(err)

	err = exec.BuildTreeFromDB(e.GetTestContext(), chainID, destination)
	e.Nil(err)

	inTree3, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[3], message3, chainID, destination)
	e.Nil(err)
	e.True(inTree3)
}

func (e *ExecutorSuite) TestVerifyOptimisticPeriod() {
	testDone := false
	defer func() {
		testDone = true
	}()
	chainID := gofakeit.Uint32()
	destination := chainID + 1
	deployManager := testutil.NewDeployManager(e.T())
	simulatedChain := geth.NewEmbeddedBackendForChainID(e.GetTestContext(), e.T(), big.NewInt(int64(chainID)))
	simulatedClient, err := backfill.DialBackend(e.GetTestContext(), simulatedChain.RPCAddress())
	e.Nil(err)
	simulatedChain.FundAccount(e.GetTestContext(), e.wallet.Address(), *big.NewInt(params.Ether))
	originContract, originRef := deployManager.GetOrigin(e.GetTestContext(), simulatedChain)

	contractConfig := config.ContractConfig{
		Address:    originContract.Address().String(),
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

	scribe, err := node.NewScribe(e.scribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.dbPath)
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
				OriginAddress: originContract.Address().String(),
			},
			{
				ChainID: destination,
			},
		},
	}

	exec, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient)
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
		execErr := exec.Listen(e.GetTestContext(), chainID)
		if !testDone {
			e.Nil(execErr)
		}
	}()

	tips := types.NewTips(big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())))
	encodedTips, err := types.EncodeTips(tips)
	e.Nil(err)

	optimisticSeconds := uint32(10)

	ownerPtr, err := originRef.OriginCaller.Owner(&bind.CallOpts{Context: e.GetTestContext()})
	e.Nil(err)

	transactOpts := simulatedChain.GetTxContext(e.GetTestContext(), &ownerPtr)

	tx, err := originRef.AddNotary(transactOpts.TransactOpts, destination, e.signer.Address())
	e.Nil(err)
	simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)

	transactOpts.Value = types.TotalTips(tips)

	recipient := [32]byte{byte(gofakeit.Uint32())}
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}

	tx, err = originRef.Dispatch(transactOpts.TransactOpts, destination, recipient, optimisticSeconds, encodedTips, body)
	e.Nil(err)

	simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := simulatedChain.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(chainID, sender.Hash(), nonce, destination, recipient, optimisticSeconds)
	message := types.NewMessage(header, tips, body)

	continueChan := make(chan bool, 1)

	var blockNum uint64

	// Wait for message to be stored in the database.`
	e.Eventually(func() bool {
		blockNum, err = e.testDB.GetBlockNumber(e.GetTestContext(), types2.DBMessage{
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

	verified, err := exec.VerifyOptimisticPeriod(e.GetTestContext(), message, blockNum)
	e.Nil(err)
	e.False(verified)

	e.Eventually(func() bool {
		verified, err = exec.VerifyOptimisticPeriod(e.GetTestContext(), message, blockNum)
		e.Nil(err)
		return verified
	})
}
