package executor_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	executorCfg "github.com/synapsecns/sanguine/agents/agents/executor/config"
	execTypes "github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsConfig "github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/merkle"
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

	state0 := types.NewState(roots[0], chainID, nonces[0], blockNumbers[0], timestamps[0])
	state1 := types.NewState(roots[1], chainID, nonces[1], blockNumbers[1], timestamps[1])
	state2 := types.NewState(roots[2], chainID, nonces[2], blockNumbers[2], timestamps[2])
	state3 := types.NewState(roots[3], chainID, nonces[3], blockNumbers[3], timestamps[3])
	failState := types.NewState(roots[1], chainID+1, nonces[2], blockNumbers[3], timestamps[0])

	snapshot := types.NewSnapshot([]types.State{state0, state1, state2, state3})

	root, proofs, err := snapshot.SnapshotRootAndProofs()
	e.Nil(err)

	// Insert the states into the database.
	err = e.ExecutorTestDB.StoreStates(e.GetTestContext(), snapshot.States(), root, proofs, snapshot.TreeHeight())
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
		execErr := exec.StartAndListenOrigin(e.GetTestContext(), chainID, e.OriginContractMetadata.Address().String())
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

func (e *ExecutorSuite) TestVerifyMessageMerkleProof() {
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
	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message0, blockNumbers[0], false, false, 0)
	e.Nil(err)
	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message1, blockNumbers[1], false, false, 0)
	e.Nil(err)
	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message2, blockNumbers[2], false, false, 0)
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

	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message3, blockNumbers[3], false, false, 0)
	e.Nil(err)

	dbTree, err = executor.NewTreeFromDB(e.GetTestContext(), chainID, e.ExecutorTestDB)
	e.Nil(err)

	exec.OverrideMerkleTree(chainID, dbTree)

	inTree3, err := exec.VerifyMessageMerkleProof(message3)
	e.Nil(err)
	e.True(inTree3)
}

func (e *ExecutorSuite) TestVerifyStateMerkleProof() {

}

func (e *ExecutorSuite) TestExecutor() {
	testDone := false
	defer func() {
		testDone = true
	}()

	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())
	summit := uint32(e.TestBackendSummit.GetChainID())

	testContractDest, _ := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), e.TestBackendDestination)

	originClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendOrigin.RPCAddress())
	e.Nil(err)
	destinationClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendDestination.RPCAddress())
	e.Nil(err)
	summitClient, err := backfill.DialBackend(e.GetTestContext(), e.TestBackendSummit.RPCAddress())
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

	scribe, err := node.NewScribe(e.ScribeTestDB, clients, scribeConfig)
	e.Nil(err)

	scribeClient := client.NewEmbeddedScribe("sqlite", e.DBPath)
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

	excCfg := executorCfg.Config{
		SummitChainID: summit,
		SummitAddress: e.SummitContract.Address().String(),
		Chains: []executorCfg.ChainConfig{
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

	exec, err := executor.NewExecutorInjectedBackend(e.GetTestContext(), excCfg, e.ExecutorTestDB, scribeClient.ScribeClient, executorClients, urls)
	e.Nil(err)

	go func() {
		execErr := exec.Run(e.GetTestContext())
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

	tree := merkle.NewTree(merkle.MessageTreeDepth)

	sender, err := e.TestBackendOrigin.Signer().Sender(tx)
	e.Nil(err)

	header := types.NewHeader(uint32(e.TestBackendOrigin.GetChainID()), sender.Hash(), nonce, uint32(e.TestBackendDestination.GetChainID()), recipient, optimisticSeconds)
	message := types.NewMessage(header, tips, body)
	leaf, err := message.ToLeaf()
	e.Nil(err)

	tree.Insert(leaf[:])

	root, err := tree.Root(1)
	e.Nil(err)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	originState := types.NewState(rootB32, chainID, nonce, big.NewInt(1), big.NewInt(1))
	randomState := types.NewState(common.BigToHash(big.NewInt(gofakeit.Int64())), chainID+1, gofakeit.Uint32(), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	originSnapshot := types.NewSnapshot([]types.State{originState, randomState})

	snapshotRoot, proofs, err := originSnapshot.SnapshotRootAndProofs()
	e.Nil(err)

	err = e.ExecutorTestDB.StoreStates(e.GetTestContext(), []types.State{originState, randomState}, snapshotRoot, proofs, originSnapshot.TreeHeight())
	e.Nil(err)

	destinationAttestation := types.NewAttestation(snapshotRoot, uint8(originSnapshot.TreeHeight()), uint32(1), big.NewInt(1), big.NewInt(1))

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
