package executor_test

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor"
	executorCfg "github.com/synapsecns/sanguine/agents/agents/executor/config"
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

func (e *ExecutorSuite) TestExecutor() {
	testDone := false
	defer func() {
		testDone = true
	}()

	chainID := uint32(e.TestBackendOrigin.GetChainID())
	destination := uint32(e.TestBackendDestination.GetChainID())
	summit := uint32(e.TestBackendSummit.GetChainID())

	testContractDest, _ := e.TestDeployManager.GetAgentsTestContract(e.GetTestContext(), e.TestBackendDestination)
	// testTransactOpts := e.TestBackendDestination.GetTxContext(e.GetTestContext(), nil)

	// _, _, _ = testContractDest, testContractRef, testTransactOpts

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
	summitConfig := config.ContractConfig{
		Address:    e.SummitContract.Address().String(),
		StartBlock: 0,
	}
	summitChainConfig := config.ChainConfig{
		ChainID:               summit,
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

	exec.OverrideMerkleTree(chainID, tree)

	var rootB32 [32]byte
	copy(rootB32[:], root)

	originState := types.NewState(rootB32, chainID, nonce, big.NewInt(1), big.NewInt(1))
	randomState := types.NewState(common.BigToHash(big.NewInt(gofakeit.Int64())), chainID+1, gofakeit.Uint32(), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
	originSnapshot := types.NewSnapshot([]types.State{originState, randomState})

	snapshotRoot, proofs, err := originSnapshot.SnapshotRootAndProofs()
	fmt.Println("PROOF 1", proofs[0])
	fmt.Println("PROOF 2", proofs[1])
	e.Nil(err)

	err = e.ExecutorTestDB.StoreMessage(e.GetTestContext(), message, 1, false, 0)
	e.Nil(err)

	err = e.ExecutorTestDB.StoreStates(e.GetTestContext(), []types.State{originState, randomState}, snapshotRoot, proofs, originSnapshot.TreeHeight())
	e.Nil(err)

	destinationAttestation := types.NewAttestation(snapshotRoot, uint8(originSnapshot.TreeHeight()), uint32(1), big.NewInt(1), big.NewInt(1))

	err = e.ExecutorTestDB.StoreAttestation(e.GetTestContext(), destinationAttestation, destination, 1, 1)
	e.Nil(err)

	time.Sleep(60 * time.Second)
}
