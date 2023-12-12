package listener_test

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/listener"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/testutil"
	relayerTypes "github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"

	"math/big"
	"strings"
	"time"
)

func (t *ListenerSuite) TestNewListener() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		testCtx := t.GetTestContext()
		testChainID := uint32(42161)
		anvilBackend := testutil.NewAnvilBackend(testCtx, t.T(), testChainID)
		evmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilBackend, t.metrics)
		Nil(t.T(), err)

		// Wallet
		testWallet, _ := wallet.FromRandom()

		// FastBridge ABI
		parsedABI, err := abi.JSON(strings.NewReader(bindings.FastBridgeMetaData.ABI))
		Nil(t.T(), err)

		testContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilBackend, testWallet, testChainID)
		Nil(t.T(), err)
		NotNil(t.T(), testContractHandler)

		// Create a new listener
		listenerConfig := &listener.ChainListenerConfig{
			ChainID:         testChainID,
			StartBlock:      0,
			BridgeAddress:   testContractHandler.FastBridgeAddress(),
			Client:          evmClient,
			PollInterval:    1,
			MaxGetLogsRange: 1000,
			Confirmations:   10,
			ABI:             parsedABI,
		}
		// Create channels
		eventChan := make(chan relayerTypes.WrappedLog)
		seenChan := make(chan relayerTypes.WrappedLog)
		chainListener, err := listener.NewChainListener(listenerConfig, testDB, eventChan, seenChan)
		Nil(t.T(), err)
		NotNil(t.T(), chainListener)
	})
}

// nolint: cyclop
func (t *ListenerSuite) TestIterateThroughLogs() {
	testCtx := t.GetTestContext()
	numLogs := 10
	lastUnconfirmedBlock := uint64(6)
	numUnconfirmedLogs := 0

	// Create test logs
	var testLogs []types.Log
	for i := 0; i < numLogs; i++ {
		blockNumber := i + 1
		newLog := testutil.GenerateTestLog()
		newLog.BlockNumber = uint64(blockNumber)
		testLogs = append(testLogs, *newLog)
		if blockNumber >= int(lastUnconfirmedBlock) {
			numUnconfirmedLogs++
		}
	}
	// FastBridge ABI
	parsedABI, err := abi.JSON(strings.NewReader(bindings.FastBridgeMetaData.ABI))
	Nil(t.T(), err)

	testChainID := uint32(42161)
	anvilBackend := testutil.NewAnvilBackend(testCtx, t.T(), testChainID)
	evmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilBackend, t.metrics)
	Nil(t.T(), err)

	// Create a new listener
	listenerConfig := &listener.ChainListenerConfig{
		ChainID:         gofakeit.Uint32(),
		StartBlock:      0,
		BridgeAddress:   common.HexToAddress(big.NewInt(gofakeit.Int64()).String()),
		Client:          evmClient,
		PollInterval:    1,
		MaxGetLogsRange: 1000,
		Confirmations:   10,
		ABI:             parsedABI,
	}

	eventChan := make(chan relayerTypes.WrappedLog, service.MaxEventChanSize)
	seenChan := make(chan relayerTypes.WrappedLog, service.MaxSeenChanSize)
	chainListener, err := listener.NewChainListener(listenerConfig, t.dbs[0], eventChan, seenChan)
	Nil(t.T(), err)
	NotNil(t.T(), chainListener)

	// Iterate through logs
	err = chainListener.IterateThroughLogs(testLogs, lastUnconfirmedBlock)
	Nil(t.T(), err)

	// Check that the logs were added to the correct channels
	confirmedLogsCount := 0
	unconfirmedLogsCount := 0
DoneDrainingA:
	for {
		select {
		case wrappedLog := <-eventChan:
			log := wrappedLog.Log
			confirmedLogsCount++
			Less(t.T(), log.BlockNumber, lastUnconfirmedBlock)
		case seenWrappedLog := <-seenChan:
			log := seenWrappedLog.Log
			unconfirmedLogsCount++
			GreaterOrEqual(t.T(), log.BlockNumber, lastUnconfirmedBlock)
		default:
			break DoneDrainingA
		}
	}
	Equal(t.T(), numUnconfirmedLogs, unconfirmedLogsCount)
	Equal(t.T(), numLogs-numUnconfirmedLogs, confirmedLogsCount)

	// Check cache functionality
	// Remove the first three logs from the testLogs array, and then add one more log to the end
	// This will simulate a new range of logs being fetched. We will increment the lastUnconfirmedBlock
	// by one as well. After iterating through the logs there should only be one new log added to the seen chan and
	// 1 in the event chan (due to the cache).
	newTestLogs := testLogs[3:]
	newTestLog := testutil.GenerateTestLog()
	newTestLog.BlockNumber = uint64(len(newTestLogs))
	newTestLogs = append(newTestLogs, *newTestLog)
	newLastUnconfirmedBlock := lastUnconfirmedBlock + 1
	err = chainListener.IterateThroughLogs(newTestLogs, newLastUnconfirmedBlock)
	Nil(t.T(), err)

	// Check that the logs were added to the correct channels
	confirmedLogsCount = 0
	unconfirmedLogsCount = 0
DoneDrainingB:
	for {
		select {
		case wrappedLog := <-eventChan:
			log := wrappedLog.Log
			confirmedLogsCount++
			Less(t.T(), log.BlockNumber, newLastUnconfirmedBlock)
		case seenWrappedLog := <-seenChan:
			log := seenWrappedLog.Log
			unconfirmedLogsCount++
			GreaterOrEqual(t.T(), log.BlockNumber, lastUnconfirmedBlock)
		default:
			break DoneDrainingB
		}
	}
	Equal(t.T(), 1, unconfirmedLogsCount)
	Equal(t.T(), 1, confirmedLogsCount)
}

// TestListenActive tests listening for events on an origin and destination chain. This test also checks for
// the correct topic filtering and that the logs are added to the correct channels.
// nolint: cyclop
func (t *ListenerSuite) TestListenE2E() {
	testCtx := t.GetTestContext()
	numBridgeRequestLogs := 5

	// Wallet
	testWallet, _ := wallet.FromRandom()

	// FastBridge ABI
	parsedABI, err := abi.JSON(strings.NewReader(bindings.FastBridgeMetaData.ABI))
	Nil(t.T(), err)

	// Origin Chain --------
	originChainID := uint32(42161)
	anvilOriginBackend := testutil.NewAnvilBackend(testCtx, t.T(), originChainID)
	originContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilOriginBackend, testWallet, originChainID)
	Nil(t.T(), err)
	NotNil(t.T(), originContractHandler)
	originEvmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilOriginBackend, t.metrics)
	Nil(t.T(), err)
	originTokens := originContractHandler.Tokens()
	originListenerConfig := &listener.ChainListenerConfig{
		ChainID:         originChainID,
		StartBlock:      0,
		BridgeAddress:   originContractHandler.FastBridgeAddress(),
		Client:          originEvmClient,
		PollInterval:    1,
		MaxGetLogsRange: 1000,
		Confirmations:   10,
		ABI:             parsedABI,
	}
	originEventChan := make(chan relayerTypes.WrappedLog, service.MaxEventChanSize)
	originSeenChan := make(chan relayerTypes.WrappedLog, service.MaxSeenChanSize)
	originListener, err := listener.NewChainListener(originListenerConfig, t.dbs[0], originEventChan, originSeenChan)
	Nil(t.T(), err)
	NotNil(t.T(), originListener)

	// Destination Chain --------
	destinationChainID := uint32(1)
	anvilDestinationBackend := testutil.NewAnvilBackend(testCtx, t.T(), destinationChainID)
	destinationContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilDestinationBackend, testWallet, destinationChainID)
	Nil(t.T(), err)
	NotNil(t.T(), destinationContractHandler)
	destinationEvmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilOriginBackend, t.metrics)
	Nil(t.T(), err)
	destinationTokens := destinationContractHandler.Tokens()
	destinationListenerConfig := &listener.ChainListenerConfig{
		ChainID:         destinationChainID,
		StartBlock:      0,
		BridgeAddress:   destinationContractHandler.FastBridgeAddress(),
		Client:          destinationEvmClient,
		PollInterval:    1,
		MaxGetLogsRange: 1000,
		Confirmations:   0,
		ABI:             parsedABI,
	}
	destinationEventChan := make(chan relayerTypes.WrappedLog, service.MaxEventChanSize)
	destinationSeenChan := make(chan relayerTypes.WrappedLog, service.MaxSeenChanSize)
	destinationListener, err := listener.NewChainListener(destinationListenerConfig, t.dbs[0], destinationEventChan, destinationSeenChan)
	Nil(t.T(), err)
	NotNil(t.T(), destinationListener)

	// Trigger all on-chain events on both origin and destination chains
	for i := int64(0); i < int64(numBridgeRequestLogs); i++ {
		bridgeParams := bindings.IFastBridgeBridgeParams{
			DstChainId:   destinationChainID,
			To:           testWallet.Address(),
			OriginToken:  originTokens[0].Erc20Address,
			DestToken:    destinationTokens[0].Erc20Address,
			OriginAmount: big.NewInt(params.GWei + i),
			DestAmount:   big.NewInt(params.GWei + i),
			Deadline:     big.NewInt(time.Now().Unix() + 4000),
		}
		tx, txErr := originContractHandler.FBExecuteBridge(testCtx, bridgeParams)
		Nil(t.T(), txErr)
		NotNil(t.T(), tx)

		// Get the transaction receipt to get logs (to get request for executing relay)
		receipt, txErr := anvilOriginBackend.TransactionReceipt(testCtx, tx.Hash())
		Nil(t.T(), txErr)
		event := new(bindings.FastBridgeBridgeRequested)
		for _, log := range receipt.Logs {
			// Check if the log is a BridgeRequested event
			if log.Topics[0] == originListener.ABI().Events["BridgeRequested"].ID {
				// Unpack the event
				logErr := originListener.ABI().UnpackIntoInterface(event, "BridgeRequested", log.Data)
				Nil(t.T(), logErr)
				continue
			}
		}

		// Execute relay, should emit BridgeRelayed
		destTx, txErr := destinationContractHandler.FBExecuteRelay(testCtx, event.Request)
		Nil(t.T(), txErr)
		NotNil(t.T(), destTx)
	}

	listenCtx, cancelListening := context.WithCancel(testCtx)
	go func() {
		// Listen on origin chain
		listenErr := originListener.StartListening(listenCtx)
		fmt.Println("origin listener terminated", listenErr)
	}()

	go func() {
		// Listen on destination chain
		listenErr := destinationListener.StartListening(listenCtx)
		fmt.Println("destination listener terminated", listenErr)
	}()

	// Wait for listener to get all logs
	<-time.After(15 * time.Second)

	// Cancel listening
	cancelListening()

	// Drain all logs and check that the logs were added to the correct channels
	originCount := 0
	destinationCount := 0
	requestCount := 0
	relayCount := 0
DoneDraining:
	for {
		select {
		case originWrappedLog := <-originEventChan:
			originCount++
			log := originWrappedLog.Log
			if log.Topics[0] == originListener.ABI().Events["BridgeRequested"].ID {
				requestCount++
			} else if log.Topics[0] == originListener.ABI().Events["BridgeRelayed"].ID {
				relayCount++
			}
		case destinationWrappedLog := <-destinationEventChan:
			destinationCount++
			log := destinationWrappedLog.Log
			if log.Topics[0] == originListener.ABI().Events["BridgeRequested"].ID {
				requestCount++
			} else if log.Topics[0] == originListener.ABI().Events["BridgeRelayed"].ID {
				relayCount++
			}
		default:
			break DoneDraining
		}
	}
	Equal(t.T(), numBridgeRequestLogs, originCount)
	Equal(t.T(), numBridgeRequestLogs, destinationCount)
}

// TestListenActive simulates listening for events occurring while a listener is active.
// nolint: cyclop
func (t *ListenerSuite) TestListenActive() {
	testCtx := t.GetTestContext()
	numBridgeRequestLogs := 5

	// Wallet
	testWallet, _ := wallet.FromRandom()

	// FastBridge ABI
	parsedABI, err := abi.JSON(strings.NewReader(bindings.FastBridgeMetaData.ABI))
	Nil(t.T(), err)

	// Init
	originChainID := uint32(42161)
	destinationChainID := uint32(1)
	anvilOriginBackend := testutil.NewAnvilBackend(testCtx, t.T(), originChainID)
	originContractHandler, err := testutil.NewTestContractHandlerImpl(testCtx, anvilOriginBackend, testWallet, originChainID)
	Nil(t.T(), err)
	NotNil(t.T(), originContractHandler)
	originEvmClient, err := testutil.NewEVMClientFromAnvil(testCtx, anvilOriginBackend, t.metrics)
	Nil(t.T(), err)
	originTokens := originContractHandler.Tokens()
	originListenerConfig := &listener.ChainListenerConfig{
		ChainID:         originChainID,
		StartBlock:      0,
		BridgeAddress:   originContractHandler.FastBridgeAddress(),
		Client:          originEvmClient,
		PollInterval:    1,
		MaxGetLogsRange: 1000,
		Confirmations:   1,
		ABI:             parsedABI,
	}
	originEventChan := make(chan relayerTypes.WrappedLog, service.MaxEventChanSize)
	originSeenChan := make(chan relayerTypes.WrappedLog, service.MaxSeenChanSize)
	originListener, err := listener.NewChainListener(originListenerConfig, t.dbs[0], originEventChan, originSeenChan)
	Nil(t.T(), err)
	NotNil(t.T(), originListener)

	// Trigger initial bridge request logs
	for i := int64(0); i < int64(numBridgeRequestLogs); i++ {
		bridgeParams := bindings.IFastBridgeBridgeParams{
			DstChainId:   destinationChainID,
			To:           testWallet.Address(),
			OriginToken:  originTokens[0].Erc20Address,
			DestToken:    originTokens[1].Erc20Address,
			OriginAmount: big.NewInt(params.GWei + i),
			DestAmount:   big.NewInt(params.GWei + i),
			Deadline:     big.NewInt(time.Now().Unix() + 4000),
		}
		tx, txErr := originContractHandler.FBExecuteBridge(testCtx, bridgeParams)
		Nil(t.T(), txErr)
		NotNil(t.T(), tx)
	}
	lastBlock, err := originEvmClient.BlockNumber(testCtx)
	Nil(t.T(), err)

	// Before starting the listener we want to make sure that the chain's head is
	// greater than the last block the event was in + confirmations. This will allow us
	// to do a deterministic test in which the events emitted after the listener starts
	// will go into the unconfirmed channel.
DoneWaitingForConfirmations:
	for {
		select {
		case <-testCtx.Done():
			t.FailNow("context was canceled")
		case <-time.After(1 * time.Second):
			// Check that the current chain head - confirmations is greater the block the last event was in.
			currentHead, waitConfErr := originEvmClient.BlockNumber(testCtx)
			Nil(t.T(), waitConfErr)
			if currentHead-originListenerConfig.Confirmations > lastBlock {
				break DoneWaitingForConfirmations
			}
			continue
		}
	}

	listenCtx, cancelListening := context.WithCancel(testCtx)
	go func() {
		listenErr := originListener.StartListening(listenCtx)
		fmt.Println("origin listener terminated", listenErr)
	}()

	// Wait for listener to get all logs + confirmations
	<-time.After(10 * time.Second)

	// Add more logs to chain
	// Trigger initial bridge request logs
	for i := int64(0); i < int64(numBridgeRequestLogs); i++ {
		bridgeParams := bindings.IFastBridgeBridgeParams{
			DstChainId:   destinationChainID,
			To:           testWallet.Address(),
			OriginToken:  originTokens[0].Erc20Address,
			DestToken:    originTokens[0].Erc20Address,
			OriginAmount: big.NewInt(params.GWei + i),
			DestAmount:   big.NewInt(params.GWei + i),
			Deadline:     big.NewInt(time.Now().Unix() + 4000),
		}
		tx, txErr := originContractHandler.FBExecuteBridge(testCtx, bridgeParams)
		Nil(t.T(), txErr)
		NotNil(t.T(), tx)
	}

	// Wait for listener to get all logs + confirmations
	<-time.After(10 * time.Second)

	// Cancel listening
	cancelListening()

	// Drain all logs and check that the logs were added to the correct channels
	confirmedCount := 0
	unconfirmedCount := 0
DoneDraining:
	for {
		select {
		case <-originEventChan:
			confirmedCount++
		case <-originSeenChan:
			unconfirmedCount++
		default:
			break DoneDraining
		}
	}
	Equal(t.T(), numBridgeRequestLogs*2, confirmedCount)
	Equal(t.T(), numBridgeRequestLogs, unconfirmedCount)
}
