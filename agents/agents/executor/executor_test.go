package executor_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/db"
	execConfig "github.com/synapsecns/sanguine/agents/config/executor"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/merkle"
	agentsConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
	"math/big"
	"time"
)

func (e *ExecutorSuite) TestVerifyState() {
	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())

	excCfg := execConfig.Config{
		Chains: []execConfig.ChainConfig{
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

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath, e.ScribeMetrics)
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

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls, e.ExecutorMetrics)
	e.Nil(err)

	roots := [][32]byte{
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
	}
	nonces := []uint32{
		gofakeit.Uint32(), gofakeit.Uint32(),
		gofakeit.Uint32(), gofakeit.Uint32(),
	}
	blockNumbers := []*big.Int{
		big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())),
		big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())),
	}
	timestamps := []*big.Int{
		big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())),
		big.NewInt(int64(gofakeit.Uint32())), big.NewInt(int64(gofakeit.Uint32())),
	}

	gasDatas := []types.GasData{}
	for i := 0; i < 4; i++ {
		gasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())
		gasDatas = append(gasDatas, gasData)
	}

	state0 := types.NewState(roots[0], chainID, nonces[0], blockNumbers[0], timestamps[0], gasDatas[0])
	state1 := types.NewState(roots[1], chainID, nonces[1], blockNumbers[1], timestamps[1], gasDatas[1])
	state2 := types.NewState(roots[2], chainID, nonces[2], blockNumbers[2], timestamps[2], gasDatas[2])
	state3 := types.NewState(roots[3], chainID, nonces[3], blockNumbers[3], timestamps[3], gasDatas[3])
	failState := types.NewState(roots[1], chainID+1, nonces[2], blockNumbers[3], timestamps[0], gasDatas[3])

	snapshot := types.NewSnapshot([]types.State{state0, state1, state2, state3})

	root, proofs, err := snapshot.SnapshotRootAndProofs()
	e.Nil(err)

	// Insert the states into the database.
	err = e.ExecutorTestDB.StoreStates(e.GetTestContext(), snapshot.States(), root, proofs, 5)
	e.Nil(err)

	inTree0, err := exec.VerifyStateMerkleProof(e.GetTestContext(), state0)
	e.Nil(err)
	e.True(inTree0)

	inTree1, err := exec.VerifyStateMerkleProof(e.GetTestContext(), state1)
	e.Nil(err)
	e.True(inTree1)

	inTree2, err := exec.VerifyStateMerkleProof(e.GetTestContext(), state2)
	e.Nil(err)
	e.True(inTree2)

	inTree3, err := exec.VerifyStateMerkleProof(e.GetTestContext(), state3)
	e.Nil(err)
	e.True(inTree3)

	inTreeFail, err := exec.VerifyStateMerkleProof(e.GetTestContext(), failState)
	e.Nil(err)
	e.False(inTreeFail)
}

//nolint:maintidx
func (e *ExecutorSuite) TestMerkleInsert() {
	// TODO (joe and lex): FIX ME
	// e.T().Skip()
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
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{contractConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{chainConfig},
	}
	simulatedClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress(), e.ScribeMetrics)
	e.Nil(err)
	clients := map[uint32][]backfill.ScribeBackend{
		chainID: {simulatedClient, simulatedClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig, e.ScribeMetrics)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath, e.ScribeMetrics)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		scribeError := scribe.Start(e.GetTestContext())
		if !testDone {
			e.Nil(scribeError)
		}
	}()

	excCfg := execConfig.Config{
		Chains: []execConfig.ChainConfig{
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

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls, e.ExecutorMetrics)
	e.Nil(err)

	_, err = exec.GetMerkleTree(chainID).Root(1)
	e.NotNil(err)

	testTree := merkle.NewTree(merkle.MessageTreeHeight)

	recipients := [][32]byte{{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())}}
	optimisticSeconds := []uint32{gofakeit.Uint32(), gofakeit.Uint32()}

	messageBytes := []byte{byte(gofakeit.Uint32()), byte(gofakeit.Uint32()), byte(gofakeit.Uint32()), byte(gofakeit.Uint32()), byte(gofakeit.Uint32())}
	msgTips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))

	transactOpts := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), e.OriginContractMetadata.OwnerPtr())
	transactOpts.Value = types.TotalTips(msgTips)

	paddedRequest := big.NewInt(0)
	tx, err := e.OriginContract.SendBaseMessage(
		transactOpts.TransactOpts,
		destination,
		recipients[0],
		optimisticSeconds[0],
		paddedRequest,
		messageBytes,
	)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	header := types.NewHeader(types.MessageFlagBase, chainID, 1, destination, optimisticSeconds[0])

	msgSender := common.BytesToHash(transactOpts.TransactOpts.From.Bytes())
	msgRequest := types.NewRequest(uint32(0), uint64(0), big.NewInt(0))
	baseMessage := types.NewBaseMessage(msgSender, recipients[0], msgTips, msgRequest, messageBytes)
	message, err := types.NewMessageFromBaseMessage(header, baseMessage)
	e.Nil(err)

	leafA, err := message.ToLeaf()
	e.Nil(err)

	testTree.Insert(leafA[:])
	testRootA, err := testTree.Root(1)
	e.Nil(err)

	// Start the exec.
	go func() {
		execErr := exec.StartAndListenOrigin(e.GetTestContext(), chainID, e.OriginContractMetadata.Address().String())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	waitChan := make(chan bool, 2)

	e.Eventually(func() bool {
		rootA, err := exec.GetMerkleTree(chainID).Root(1)
		if err != nil {
			// This transaction is needed to get the simulated chain's block number to increase by 1, since StreamLogs will
			// do lastBlockNumber - 1.
			tx, err = e.TestContractOnOrigin.EmitAgentsEventA(transactOpts.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
			e.Nil(err)
			e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

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

	transactOpts.Value = types.TotalTips(msgTips)
	// paddedRequest = big.NewInt(0)
	tx, err = e.OriginContract.SendBaseMessage(transactOpts.TransactOpts, destination, recipients[1], optimisticSeconds[1], paddedRequest, messageBytes)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	// Advance block again.
	tx, err = e.TestContractOnOrigin.EmitAgentsEventA(transactOpts.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	header = types.NewHeader(types.MessageFlagBase, chainID, 2, destination, optimisticSeconds[1])

	msgRequest = types.NewRequest(uint32(0), uint64(0), big.NewInt(0))
	baseMessage = types.NewBaseMessage(msgSender, recipients[1], msgTips, msgRequest, messageBytes)
	message, err = types.NewMessageFromBaseMessage(header, baseMessage)
	e.Nil(err)

	leafB, err := message.ToLeaf()
	e.Nil(err)
	testTree.Insert(leafB[:])
	testRootB, err := testTree.Root(2)
	e.Nil(err)

	e.Eventually(func() bool {
		rootB, err := exec.GetMerkleTree(chainID).Root(2)
		if err != nil {
			// This transaction is needed to get the simulated chain's block number to increase by 1, since StreamLogs will
			// do lastBlockNumber - 1.
			tx, err = e.TestContractOnOrigin.EmitAgentsEventA(transactOpts.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
			e.Nil(err)
			e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)
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

func (e *ExecutorSuite) TestVerifyMessageMerkleProof() {
	// TODO (joe and lex): FIX ME
	// e.T().Skip()
	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())

	excCfg := execConfig.Config{
		Chains: []execConfig.ChainConfig{
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

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath, e.ScribeMetrics)
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

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls, e.ExecutorMetrics)
	e.Nil(err)

	nonces := []uint32{1, 2, 3, 4}
	blockNumbers := []uint64{10, 20, 30, 40}
	optimisticSeconds := []uint32{
		gofakeit.Uint32(), gofakeit.Uint32(),
		gofakeit.Uint32(), gofakeit.Uint32(),
	}
	messageBytes := [][]byte{
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
		{byte(gofakeit.Uint32())}, {byte(gofakeit.Uint32())},
	}

	header0 := types.NewHeader(types.MessageFlagManager, chainID, nonces[0], destination, optimisticSeconds[0])
	header1 := types.NewHeader(types.MessageFlagManager, chainID, nonces[1], destination, optimisticSeconds[1])
	header2 := types.NewHeader(types.MessageFlagManager, chainID, nonces[2], destination, optimisticSeconds[2])
	header3 := types.NewHeader(types.MessageFlagManager, chainID, nonces[3], destination, optimisticSeconds[3])

	message0 := types.NewMessage(header0, nil, messageBytes[0])
	message1 := types.NewMessage(header1, nil, messageBytes[1])
	message2 := types.NewMessage(header2, nil, messageBytes[2])
	message3 := types.NewMessage(header3, nil, messageBytes[3])
	failMessage := types.NewMessage(header1, nil, messageBytes[3])

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

func (e *ExecutorSuite) TestExecutor() {
	// e.T().Skip()
	testDone := false
	defer func() {
		testDone = true
	}()

	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())
	summit := uint32(e.TestBackendSummit.GetChainID())

	testContractDest, _ := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), e.TestBackendDestination)

	originClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress(), e.ScribeMetrics)
	e.Nil(err)
	destinationClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendDestination.RPCAddress(), e.ScribeMetrics)
	e.Nil(err)
	summitClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendSummit.RPCAddress(), e.ScribeMetrics)
	e.Nil(err)

	originConfig := config.ContractConfig{
		Address:    e.OriginContract.Address().String(),
		StartBlock: 0,
	}
	originChainConfig := config.ChainConfig{
		ChainID:               chainID,
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{originConfig},
	}
	destinationConfig := config.ContractConfig{
		Address:    e.DestinationContract.Address().String(),
		StartBlock: 0,
	}
	destinationChainConfig := config.ChainConfig{
		ChainID:               destination,
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{destinationConfig},
	}
	summitConfig := config.ContractConfig{
		Address:    e.SummitContract.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := config.ChainConfig{
		ChainID:               summit,
		BlockTimeChunkSize:    1,
		ContractSubChunkSize:  1,
		RequiredConfirmations: 0,
		Contracts:             []config.ContractConfig{summitConfig},
	}
	scribeConfig := config.Config{
		Chains: []config.ChainConfig{originChainConfig, destinationChainConfig, summitChainConfig},
	}
	clients := map[uint32][]backfill.ScribeBackend{
		chainID:     {originClient, originClient},
		destination: {destinationClient, destinationClient},
		summit:      {summitClient, summitClient},
	}

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig, e.ScribeMetrics)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath, e.ScribeMetrics)
	go func() {
		scribeErr := scribeClient.Start(e.GetTestContext())
		e.Nil(scribeErr)
	}()

	// Start the Scribe.
	go func() {
		scribeError := scribe.Start(e.GetTestContext())
		if !testDone {
			e.Nil(scribeError)
		}
	}()

	excCfg := execConfig.Config{
		SummitChainID: summit,
		SummitAddress: e.SummitContract.Address().String(),
		Chains: []execConfig.ChainConfig{
			{
				ChainID:       chainID,
				OriginAddress: e.OriginContract.Address().String(),
			},
			{
				ChainID:            destination,
				DestinationAddress: e.DestinationContract.Address().String(),
			},
			{
				ChainID: summit,
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
		destination: e.TestBackendDestination,
		summit:      e.TestBackendSummit,
	}

	urls := map[uint32]string{
		chainID:     e.TestBackendOrigin.RPCAddress(),
		destination: e.TestBackendDestination.RPCAddress(),
		summit:      e.TestBackendSummit.RPCAddress(),
	}

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls, e.ExecutorMetrics)
	e.Nil(err)

	go func() {
		execErr := exec.Run(e.GetTestContext())
		if !testDone {
			e.Nil(execErr)
		}
	}()

	tips := types.NewTips(big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)), big.NewInt(int64(0)))

	optimisticSeconds := uint32(10)

	recipient := testContractDest.Address().Hash()
	nonce := uint32(1)
	body := []byte{byte(gofakeit.Uint32())}

	txContextOrigin := e.TestBackendOrigin.GetTxContext(e.GetTestContext(), e.OriginContractMetadata.OwnerPtr())
	txContextOrigin.Value = types.TotalTips(tips)

	paddedRequest := big.NewInt(0)
	// txContextOrigin.TransactOpts.Value = big.NewInt(0)
	tx, err := e.OriginContract.SendBaseMessage(
		txContextOrigin.TransactOpts,
		uint32(e.TestBackendDestination.GetChainID()),
		recipient,
		optimisticSeconds,
		paddedRequest,
		body)
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	// This transaction is needed to get the simulated chain's block number to increase by 1, since StreamLogs will
	// do lastBlockNumber - 1.
	tx, err = e.TestContractOnOrigin.EmitAgentsEventA(txContextOrigin.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	e.Nil(err)
	e.TestBackendOrigin.WaitForConfirmation(e.GetTestContext(), tx)

	tree := merkle.NewTree(merkle.MessageTreeHeight)

	header := types.NewHeader(types.MessageFlagBase, uint32(e.TestBackendOrigin.GetChainID()), nonce, uint32(e.TestBackendDestination.GetChainID()), optimisticSeconds)

	var msgSender [32]byte
	copy(msgSender[:], txContextOrigin.TransactOpts.From.Bytes())
	msgTips := types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
	msgRequest := types.NewRequest(uint32(1), uint64(0), big.NewInt(0))
	baseMessage := types.NewBaseMessage(msgSender, recipient, msgTips, msgRequest, body)
	message := types.NewMessage(header, baseMessage, []byte{})
	leaf, err := message.ToLeaf()
	e.Nil(err)

	tree.Insert(leaf[:])

	root, err := tree.Root(1)
	e.Nil(err)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	gasData := types.NewGasData(uint16(1), uint16(1), uint16(1), uint16(1), uint16(1), uint16(1))
	originState := types.NewState(rootB32, chainID, nonce, big.NewInt(1), big.NewInt(1), gasData)
	randomGasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())
	randomState := types.NewState(common.BigToHash(big.NewInt(gofakeit.Int64())), chainID+1, gofakeit.Uint32(), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), randomGasData)
	originSnapshot := types.NewSnapshot([]types.State{originState, randomState})

	snapshotRoot, proofs, err := originSnapshot.SnapshotRootAndProofs()
	e.Nil(err)

	err = e.ExecutorTestDB.StoreStates(e.GetTestContext(), []types.State{originState, randomState}, snapshotRoot, proofs, 5)
	e.Nil(err)

	destinationAttestation := types.NewAttestation(snapshotRoot, [32]byte{}, uint32(1), big.NewInt(1), big.NewInt(1))

	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), destinationAttestation, destination, 1, 1)
	e.Nil(err)

	mask := execTypes.DBMessage{
		ChainID:     &chainID,
		Destination: &destination,
	}
	executableMessages, err := e.ExecutorTestDB.GetExecutableMessages(e.GetTestContext(), mask, uint64(time.Now().Unix()), 1)
	e.Nil(err)
	e.Len(executableMessages, 0)

	// Make sure there is one executable message in the database.
	e.Eventually(func() bool {
		mask := execTypes.DBMessage{
			ChainID:     &chainID,
			Destination: &destination,
		}
		executableMessages, err := e.ExecutorTestDB.GetExecutableMessages(e.GetTestContext(), mask, uint64(time.Now().Unix()), 1)
		e.Nil(err)
		return len(executableMessages) == 1
	})
}

func (e *ExecutorSuite) TestSetMinimumTime() {
	// TODO (joe and lex): FIX ME
	// e.T().Skip()
	testDone := false
	defer func() {
		testDone = true
	}()
	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())

	// Store 5 messages in the database.
	for i := 1; i <= 5; i++ {
		nonce := uint32(i)
		optimisticSeconds := i
		body := []byte{byte(gofakeit.Uint32())}

		message := types.NewMessage(types.NewHeader(types.MessageFlagManager, chainID, nonce, destination, uint32(optimisticSeconds)), nil, body)

		err := e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message, uint64(i), false, 0)
		e.Nil(err)
	}

	// Ensure that there are 5 unset messages in the database.
	mask := execTypes.DBMessage{
		ChainID:     &chainID,
		Destination: &destination,
	}
	messages, err := e.ExecutorTestDB.GetUnsetMinimumTimeMessages(e.GetTestContext(), mask, 0)
	e.Nil(err)
	e.Len(messages, 5)

	// Store some states (as snapshots with length 1) in the database.
	state0 := types.NewState(
		common.BigToHash(big.NewInt(gofakeit.Int64())),
		chainID,
		1,
		big.NewInt(gofakeit.Int64()),
		big.NewInt(gofakeit.Int64()),
		types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16()))
	state1 := types.NewState(
		common.BigToHash(big.NewInt(gofakeit.Int64())),
		chainID,
		2,
		big.NewInt(gofakeit.Int64()),
		big.NewInt(gofakeit.Int64()),
		types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16()))
	state2 := types.NewState(
		common.BigToHash(big.NewInt(gofakeit.Int64())),
		chainID,
		5,
		big.NewInt(gofakeit.Int64()),
		big.NewInt(gofakeit.Int64()),
		types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16()))

	snapshot0 := types.NewSnapshot([]types.State{state0})
	snapshot1 := types.NewSnapshot([]types.State{state1})
	snapshot2 := types.NewSnapshot([]types.State{state2})

	snapshotRoot0, proofs0, err := snapshot0.SnapshotRootAndProofs()
	e.Nil(err)
	snapshotRoot1, proofs1, err := snapshot1.SnapshotRootAndProofs()
	e.Nil(err)
	snapshotRoot2, proofs2, err := snapshot2.SnapshotRootAndProofs()
	e.Nil(err)

	err = e.ExecutorTestDB.StoreStates(e.GetTestContext(), snapshot0.States(), snapshotRoot0, proofs0, 0)
	e.Nil(err)
	err = e.ExecutorTestDB.StoreStates(e.GetTestContext(), snapshot1.States(), snapshotRoot1, proofs1, 1)
	e.Nil(err)
	err = e.ExecutorTestDB.StoreStates(e.GetTestContext(), snapshot2.States(), snapshotRoot2, proofs2, 2)
	e.Nil(err)

	potentialSnapshotRoots, err := e.ExecutorTestDB.GetPotentialSnapshotRoots(e.GetTestContext(), chainID, 1)
	e.Nil(err)
	e.Len(potentialSnapshotRoots, 3)
	potentialSnapshotRoots, err = e.ExecutorTestDB.GetPotentialSnapshotRoots(e.GetTestContext(), chainID, 2)
	e.Nil(err)
	e.Len(potentialSnapshotRoots, 2)
	potentialSnapshotRoots, err = e.ExecutorTestDB.GetPotentialSnapshotRoots(e.GetTestContext(), chainID, 3)
	e.Nil(err)
	e.Len(potentialSnapshotRoots, 1)

	// Store an attestation for the first and last state's snapshot root.
	attestation0 := types.NewAttestation(snapshotRoot0, [32]byte{}, 1, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	attestation2 := types.NewAttestation(snapshotRoot2, [32]byte{}, 2, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))

	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation0, destination, 10, 10)
	e.Nil(err)
	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), attestation2, destination, 20, 20)
	e.Nil(err)

	excCfg := execConfig.Config{
		Chains: []execConfig.ChainConfig{
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

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath, e.ScribeMetrics)
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

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls, e.ExecutorMetrics)
	e.Nil(err)

	go func() {
		setMinErr := exec.SetMinimumTime(e.GetTestContext())
		if !testDone {
			e.Nil(setMinErr)
		}
	}()

	e.Eventually(func() bool {
		messages, err := e.ExecutorTestDB.GetUnsetMinimumTimeMessages(e.GetTestContext(), mask, 1)
		e.Nil(err)
		return len(messages) == 0
	})

	// Ensure that the correct attestation was used for the messages.
	for i := uint32(1); i <= 5; i++ {
		nonce := i
		messageMask := execTypes.DBMessage{
			ChainID:     &chainID,
			Destination: &destination,
			Nonce:       &nonce,
		}
		minTime, err := e.ExecutorTestDB.GetMessageMinimumTime(e.GetTestContext(), messageMask)
		e.Nil(err)
		// TODO: Check this using attestation nonce, as this is added in messaging-0.0.3.
		if i == 1 {
			e.Equal(*minTime, uint64(10+(i)))
		} else {
			e.Equal(*minTime, uint64(20+(i)))
		}
	}
}
