package executor_test

import (
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/synapsecns/sanguine/agents/agents/executor/types"
	"github.com/synapsecns/sanguine/core"
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/prysmaticlabs/prysm/shared/trieutil"
	"github.com/synapsecns/sanguine/agents/agents/executor"
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
		AttestationCollectorAddress: gofakeit.Word(),
	}

	executorClients := map[uint32]executor.Backend{
		chainIDA: simulatedChainA,
		chainIDB: simulatedChainB,
	}

	exc, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient, executorClients)
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

	executorClients := map[uint32]executor.Backend{
		chainID: simulatedChain,
	}

	exec, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient, executorClients)
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
	testDone := false
	defer func() {
		testDone = true
	}()

	contractConfig := config.ContractConfig{
		Address:    e.originContract.Address().String(),
		StartBlock: 0,
	}
	chainConfig := config.ChainConfig{
		ChainID:               e.chainID,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{chainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		e.chainID: {e.simulatedClient, e.simulatedClient},
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
				ChainID:       e.chainID,
				OriginAddress: e.originContract.Address().String(),
			},
			{
				ChainID: e.destination,
			},
		},
	}

	executorClients := map[uint32]executor.Backend{
		e.chainID:     e.simulatedChain,
		e.destination: nil,
	}

	exec, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient, executorClients)
	e.Nil(err)

	_, err = exec.GetRoot(e.GetTestContext(), 1, e.chainID, e.destination)
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

	ownerPtr, err := e.originRef.OriginHarnessCaller.Owner(&bind.CallOpts{Context: e.GetTestContext()})
	e.Nil(err)

	originOwnerAuth := e.simulatedChain.GetTxContext(e.GetTestContext(), &ownerPtr)

	tx, err := e.originRef.AddAgent(originOwnerAuth.TransactOpts, e.destination, e.signer.Address())
	e.Nil(err)
	e.simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)

	notaries, err := e.originRef.AllAgents(&bind.CallOpts{Context: e.GetTestContext()}, e.destination)
	e.Nil(err)
	e.Len(notaries, 1)

	transactOpts := e.simulatedChain.GetTxContext(e.GetTestContext(), nil)
	transactOpts.Value = types.TotalTips(tips[0])

	tx, err = e.originRef.Dispatch(transactOpts.TransactOpts, e.destination, recipients[0], optimisticSeconds[0], encodedTips, messageBytes)
	e.Nil(err)
	e.simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)

	sender, err := e.simulatedChain.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(e.chainID, sender.Hash(), 1, e.destination, recipients[0], optimisticSeconds[0])

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
		execErr := exec.Listen(e.GetTestContext(), e.chainID)
		if !testDone {
			e.Nil(execErr)
		}
	}()

	waitChan := make(chan bool, 2)

	e.Eventually(func() bool {
		rootA, err := exec.GetRoot(e.GetTestContext(), 1, e.chainID, e.destination)
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

	tx, err = e.originRef.Dispatch(transactOpts.TransactOpts, e.destination, recipients[1], optimisticSeconds[1], encodedTips, messageBytes)
	e.Nil(err)
	e.simulatedChain.WaitForConfirmation(e.GetTestContext(), tx)

	header = types.NewHeader(e.chainID, sender.Hash(), 2, e.destination, recipients[1], optimisticSeconds[1])

	message = types.NewMessage(header, tips[1], messageBytes)
	e.Nil(err)

	leaf, err = message.ToLeaf()
	e.Nil(err)
	testTree.Insert(leaf[:], 1)

	testRootB := testTree.Root()

	e.Eventually(func() bool {
		rootB, err := exec.GetRoot(e.GetTestContext(), 2, e.chainID, e.destination)
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
	exec.Stop(e.chainID)

	oldTreeItems := exec.GetMerkleTree(e.chainID, e.destination).Items()

	err = exec.BuildTreeFromDB(e.GetTestContext(), e.chainID, e.destination)
	e.Nil(err)

	newRoot, err := exec.GetRoot(e.GetTestContext(), 2, e.chainID, e.destination)
	e.Nil(err)

	newTreeItems := exec.GetMerkleTree(e.chainID, e.destination).Items()

	e.Equal(testRootB, newRoot)
	e.Equal(oldTreeItems, newTreeItems)
}

func (e *ExecutorSuite) TestVerifyMessage() {
	testTree, err := trieutil.NewTrie(32)
	e.Nil(err)

	excCfg := executorCfg.Config{
		Chains: []executorCfg.ChainConfig{
			{
				ChainID: e.chainID,
			},
			{
				ChainID: e.destination,
			},
		},
	}

	scribeClient := client.NewEmbeddedScribe("sqlite", e.dbPath)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	executorClients := map[uint32]executor.Backend{
		e.chainID:     nil,
		e.destination: nil,
	}

	exec, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient, executorClients)
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

	header0 := types.NewHeader(e.chainID, senders[0], nonces[0], e.destination, recipients[0], optimisticSeconds[0])
	header1 := types.NewHeader(e.chainID, senders[1], nonces[1], e.destination, recipients[1], optimisticSeconds[1])
	header2 := types.NewHeader(e.chainID, senders[2], nonces[2], e.destination, recipients[2], optimisticSeconds[2])
	header3 := types.NewHeader(e.chainID, senders[3], nonces[3], e.destination, recipients[3], optimisticSeconds[3])

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

	err = exec.BuildTreeFromDB(e.GetTestContext(), e.chainID, e.destination)
	e.Nil(err)

	inTree0, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[0], message0, e.chainID, e.destination)
	e.Nil(err)
	e.True(inTree0)

	inTree1, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[1], message1, e.chainID, e.destination)
	e.Nil(err)
	e.True(inTree1)

	inTree2, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[2], message2, e.chainID, e.destination)
	e.Nil(err)
	e.True(inTree2)

	inTreeFail, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[2], message1, e.chainID, e.destination)
	e.Nil(err)
	e.False(inTreeFail)

	err = e.testDB.StoreMessage(e.GetTestContext(), message3, root3, blockNumbers[3])
	e.Nil(err)

	err = exec.BuildTreeFromDB(e.GetTestContext(), e.chainID, e.destination)
	e.Nil(err)

	inTree3, err := exec.VerifyMessageNonce(e.GetTestContext(), nonces[3], message3, e.chainID, e.destination)
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

	_, passBlockRef := e.deployManager.GetOriginHarness(e.GetTestContext(), e.TestBackendDestination)
	_ = passBlockRef
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
				ChainID:       uint32(e.TestBackendOrigin.GetChainID()),
				OriginAddress: e.OriginContract.Address().String(),
			},
			{
				ChainID:            uint32(e.TestBackendDestination.GetChainID()),
				DestinationAddress: e.DestinationContract.Address().String(),
			},
		},
	}

	executorClients := map[uint32]executor.Backend{
		uint32(e.TestBackendOrigin.GetChainID()):      e.TestBackendOrigin,
		uint32(e.TestBackendDestination.GetChainID()): e.TestBackendDestination,
	}

	exec, err := executor.NewExecutor(e.GetTestContext(), excCfg, e.testDB, scribeClient.ScribeClient, executorClients)
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
		execErr := exec.Listen(e.GetTestContext(), uint32(e.TestBackendOrigin.GetChainID()))
		if !testDone {
			e.Nil(execErr)
		}
	}()
	go func() {
		execErr := exec.Listen(e.GetTestContext(), uint32(e.TestBackendDestination.GetChainID()))
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

	guardSignature, err := e.GuardSigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
	e.Nil(err)

	notarySignature, err := e.NotarySigner.SignMessage(e.GetTestContext(), core.BytesToSlice(hashedAttestation), false)
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
		_, err = e.testDB.GetAttestationBlockNumber(e.GetTestContext(), types2.DBAttestation{
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

	verified, err := exec.VerifyOptimisticPeriod(e.GetTestContext(), message)
	e.Nil(err)
	e.False(verified)

	e.Eventually(func() bool {
		verified, err = exec.VerifyOptimisticPeriod(e.GetTestContext(), message)
		if err != nil {
			return false
		}
		if verified {
			return true
		}
		// Need to create a tx and wait for it to be confirmed to continue adding blocks, and therefore
		// increase the `time`.
		tx, err = passBlockRef.Dispatch(txContextDestination.TransactOpts, gofakeit.Uint32(), recipient, optimisticSeconds, encodedTips, body)
		e.Nil(err)
		e.TestBackendDestination.WaitForConfirmation(e.GetTestContext(), tx)
		return false
	})
}
