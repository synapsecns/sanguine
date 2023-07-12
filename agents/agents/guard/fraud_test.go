package guard_test

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/guard"
	"github.com/synapsecns/sanguine/agents/agents/notary"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/types"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/client"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/node"
)

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

	omniiiii := omniRPCClient.GetEndpoint(int(g.SummitDomainClient.Config().DomainID), 1)
	fmt.Println("OMNIIIIIIII", omniiiii)

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

	notaryTestConfig := config.AgentConfig{
		Domains: map[string]config.DomainConfig{
			"origin_client":      g.OriginDomainClient.Config(),
			"destination_client": g.DestinationDomainClient.Config(),
			"summit_client":      g.SummitDomainClient.Config(),
		},
		DomainID:       g.DestinationDomainClient.Config().DomainID,
		SummitDomainID: g.SummitDomainClient.Config().DomainID,
		BondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(g.T(), "", g.NotaryBondedWallet.PrivateKeyHex()).Name(),
		},
		UnbondedSigner: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(g.T(), "", g.NotaryUnbondedWallet.PrivateKeyHex()).Name(),
		},
		RefreshIntervalSeconds: 5,
	}

	notary, err := notary.NewNotary(g.GetTestContext(), notaryTestConfig, omniRPCClient, g.NotaryTestDB, g.NotaryMetrics)
	Nil(g.T(), err)

	_ = notary

	//var badStateRoot [32]byte
	//for i := 0; i < 32; i++ {
	//	badStateRoot[i] = byte(i)
	//}
	//badGasData := types.NewGasData(uint16(7), uint16(7), uint16(7), uint16(7), uint16(7), uint16(7))
	//badBlockNumberBigInt := new(big.Int).SetUint64(10)
	//badTimeStampBigInt := new(big.Int).SetUint64(10)
	//originID := g.OriginDomainClient.Config().DomainID
	//badStateNonce := uint32(100)
	//
	//badState := types.NewState(badStateRoot, originID, badStateNonce, badBlockNumberBigInt, badTimeStampBigInt, badGasData)
	//
	//badSnapshotStates := make([]types.State, 0, 1)
	//badSnapshotStates = append(badSnapshotStates, badState)
	//
	//badSnapshot := types.NewSnapshot(badSnapshotStates)
	//
	//badSnapshotSignature, badEncodedSnapshot, _, err := badSnapshot.SignSnapshot(g.GetTestContext(), g.NotaryBondedSigner)
	//Nil(g.T(), err)
	//
	//badSnapshotRawSig, err := types.EncodeSignature(badSnapshotSignature)
	//Nil(g.T(), err)
	//
	//txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())
	//badSummitTx, err := g.InboxOnSummit.SubmitSnapshot(
	//	txContextSummit.TransactOpts,
	//	badEncodedSnapshot,
	//	badSnapshotRawSig,
	//)
	//
	//Nil(g.T(), err)
	//_ = badSummitTx

	gasData := types.NewGasData(gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16(), gofakeit.Uint16())

	gasDataEncoded, err := types.EncodeGasData(gasData)
	Nil(g.T(), err)

	fmt.Println("GAS DATA LENGTH: ", len(gasDataEncoded))

	fraudulentState := types.NewState(
		common.BigToHash(big.NewInt(gofakeit.Int64())),
		g.OriginDomainClient.Config().DomainID,
		1,
		big.NewInt(int64(gofakeit.Int32())),
		big.NewInt(int64(gofakeit.Int32())),
		gasData,
	)

	encodedState, err := types.EncodeState(fraudulentState)
	Nil(g.T(), err)

	fmt.Println("STATE LENGTH: ", len(encodedState))

	fraudulentSnapshot := types.NewSnapshot([]types.State{fraudulentState})
	encodedFraudulentSnapshot, err := types.EncodeSnapshot(fraudulentSnapshot)
	Nil(g.T(), err)

	fmt.Println("SNAPSHOT LENGTH FROM ENCODE: ", len(encodedFraudulentSnapshot))

	snapshotSignature, encodedSnapshot, _, err := fraudulentSnapshot.SignSnapshot(g.GetTestContext(), g.GuardBondedSigner)
	Nil(g.T(), err)

	fmt.Println("SNAPSHOT LENGTH: ", len(encodedSnapshot))

	// Checking if stateAmount (1) * state length (62) == snapshot.length
	// Also checks if stateAmount != 0 && <= 32

	//opts, err := g.NotaryUnbondedSigner.GetTransactor(g.GetTestContext(), g.TestBackendSummit.GetBigChainID())
	//Nil(g.T(), err)

	txContextSummit := g.TestBackendSummit.GetTxContext(g.GetTestContext(), g.SummitMetadata.OwnerPtr())

	// transactOpts := bind.NewKeyedTransactor(g.NotaryUnbondedWallet.PrivateKey())
	status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner)
	Equal(g.T(), status.Flag(), uint8(1))
	Nil(g.T(), err)
	fmt.Printf("status: %v\n", status)

	tx, err := g.SummitDomainClient.Inbox().SubmitSnapshot(txContextSummit.TransactOpts, g.NotaryBondedSigner, encodedSnapshot, snapshotSignature)
	fmt.Printf("TXXXXXXXx: %v\n", tx)

	Nil(g.T(), err)
	NotNil(g.T(), tx)

	g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), tx)

	txReceipt, err := g.TestBackendSummit.TransactionReceipt(g.GetTestContext(), tx.Hash())
	Nil(g.T(), err)

	fmt.Println("TXRECEIPTTTT", txReceipt.Status)

	fmt.Println("TXHASHHHHHHH", tx.Hash().String())

	fmt.Println("STOP")
	// time.Sleep(15 * time.Minute)

	g.Eventually(func() bool {
		time.Sleep(5 * time.Second)

		status, err := g.OriginDomainClient.LightManager().GetAgentStatus(g.GetTestContext(), g.GuardBondedSigner)
		Nil(g.T(), err)
		fmt.Printf("status: %v\n", status)

		if status.Flag() == uint8(4) {
			fmt.Println("SUCCESS")
			return true
		}

		bumpTx, err := g.TestContractOnSummit.EmitAgentsEventA(txContextSummit.TransactOpts, big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()), big.NewInt(gofakeit.Int64()))
		Nil(g.T(), err)
		g.TestBackendSummit.WaitForConfirmation(g.GetTestContext(), bumpTx)
		fmt.Println("FALSE")
		return false
	})
}

func getNewUint40() *big.Int {
	// Max random value, a 130-bits integer, i.e 2^96 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(40), nil).Sub(max, big.NewInt(1))

	// Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	return n
}
