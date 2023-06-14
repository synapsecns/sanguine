package relayer_test

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

func (c *CCTPRelayerSuite) TestHandleCircleRequestSent() {
	// setup
	originChain := c.testBackends[0]
	destChain := c.testBackends[1]
	_, originSynapseCCTP := c.deployManager.GetSynapseCCTP(c.GetTestContext(), originChain)
	_, originMockUsdc := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), originChain)

	// create a new relayer
	mockAPI := api.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), c.GetTestConfig(), c.testStore, c.GetTestScribe(), omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// mint token
	opts := originChain.GetTxContext(c.GetTestContext(), nil)
	amount := big.NewInt(1000000000000000000)
	tx, err := originMockUsdc.MintPublic(opts.TransactOpts, opts.From, amount)
	c.Nil(err)
	originChain.WaitForConfirmation(c.GetTestContext(), tx)

	// approve token
	tx, err = originMockUsdc.Approve(opts.TransactOpts, originSynapseCCTP.Address(), amount)
	c.Nil(err)
	originChain.WaitForConfirmation(c.GetTestContext(), tx)

	// send token
	tx, err = originSynapseCCTP.SendCircleToken(opts.TransactOpts, opts.From, big.NewInt(int64(destChain.GetChainID())), originMockUsdc.Address(), amount, 0, []byte{})
	c.Nil(err)
	originChain.WaitForConfirmation(c.GetTestContext(), tx)

	// handle send request
	err = relay.HandleCircleRequestSent(c.GetTestContext(), tx.Hash(), uint32(originChain.GetChainID()))
	c.Nil(err)
	recvChan := relay.GetUsdcMsgRecvChan(uint32(originChain.GetChainID()))
	msg := <-recvChan
	c.Equal(msg.OriginTxHash, tx.Hash().String())
	c.Equal(msg.State, relayTypes.Pending)

	// verify that the request is stored in the db
	var storedMsg relayTypes.Message
	err = c.testStore.DB().Where("origin_tx_hash = ?", msg.OriginTxHash).First(&storedMsg).Error
	c.Nil(err)
	c.Equal(*msg, storedMsg)
}

func (c *CCTPRelayerSuite) TestFetchAttestation() {
	// create a new relayer
	mockAPI := api.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), c.GetTestConfig(), c.testStore, c.GetTestScribe(), omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// override mocked api call
	expectedSignature := "abc"
	mockAPI.SetGetAttestation(func(ctx context.Context, txHash string) (attestation []byte, err error) {
		return []byte(expectedSignature), nil
	})

	// fetch attestation
	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
	msg := relayTypes.Message{
		Message:          []byte{},
		MessageHash:      testHash,
		FormattedRequest: []byte{},
	}
	originChain := c.testBackends[0]
	err = relay.FetchAttestation(c.GetTestContext(), uint32(originChain.GetChainID()), &msg)
	c.Nil(err)

	sendChan := relay.GetUsdcMsgSendChan(uint32(originChain.GetChainID()))
	completeMsg := <-sendChan
	c.Equal(completeMsg.MessageHash, msg.MessageHash)
	c.Equal(completeMsg.Attestation, []byte(expectedSignature))
	c.Equal(completeMsg.State, relayTypes.Attested)

	// verify that the attested request is stored in the db
	var storedMsg relayTypes.Message
	err = c.testStore.DB().Where("origin_tx_hash = ?", completeMsg.OriginTxHash).First(&storedMsg).Error
	c.Nil(err)
	c.Equal(*completeMsg, storedMsg)
}

func (c *CCTPRelayerSuite) TestSubmitReceiveCircleToken() {
	// create a new relayer
	mockAPI := api.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), c.GetTestConfig(), c.testStore, c.GetTestScribe(), omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// build test msg
	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
	originChain := c.testBackends[0]
	originChainID, err := originChain.ChainID(c.GetTestContext())
	c.Nil(err)
	destChain := c.testBackends[1]
	destChainID, err := destChain.ChainID(c.GetTestContext())
	c.Nil(err)
	msg := relayTypes.Message{
		OriginTxHash:     testHash,
		OriginChainID:    uint32(originChainID.Int64()),
		DestChainID:      uint32(destChainID.Int64()),
		Message:          []byte{},
		MessageHash:      testHash,
		FormattedRequest: []byte{},
	}

	// submit ReceiveCircleToken()
	err = relay.SubmitReceiveCircleToken(c.GetTestContext(), &msg)
	c.Nil(err)

	// verify that the attested request is stored in the db
	var storedMsg relayTypes.Message
	err = c.testStore.DB().Where("origin_tx_hash = ?", msg.OriginTxHash).First(&storedMsg).Error
	c.Nil(err)
	msg.State = relayTypes.Complete
	msg.DestTxHash = storedMsg.DestTxHash
	c.Equal(msg, storedMsg)
}

func (c *CCTPRelayerSuite) TestBridgeUSDC() {
	// create a new relayer
	mockAPI := api.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), c.GetTestConfig(), c.testStore, c.GetTestScribe(), omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// start relayer
	ctx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()
	//nolint:errcheck
	go relay.Run(ctx)

	// mint some USDC on send chain
	originChain := c.testBackends[0]
	destChain := c.testBackends[1]
	_, originMockUsdc := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), originChain)
	originChainID, err := originChain.ChainID(c.GetTestContext())
	c.Nil(err)
	bridgeAmount := big.NewInt(1000000000) // 1000 USDC
	opts := originChain.GetTxContext(c.GetTestContext(), nil)
	tx, err := originMockUsdc.MintPublic(opts.TransactOpts, opts.From, bridgeAmount)
	c.Nil(err)
	originChain.WaitForConfirmation(c.GetTestContext(), tx)

	// approve USDC for spending
	_, originSynapseCCTP := c.deployManager.GetSynapseCCTP(c.GetTestContext(), originChain)
	tx, err = originMockUsdc.Approve(opts.TransactOpts, originSynapseCCTP.Address(), bridgeAmount)
	c.Nil(err)
	originChain.WaitForConfirmation(c.GetTestContext(), tx)

	// send USDC from originChain
	destChainID, err := destChain.ChainID(c.GetTestContext())
	c.Nil(err)
	tx, err = originSynapseCCTP.SendCircleToken(opts.TransactOpts, opts.From, destChainID, originMockUsdc.Address(), bridgeAmount, 0, []byte{})
	c.Nil(err)
	originChain.WaitForConfirmation(c.GetTestContext(), tx)

	// TODO: figure out why log is not streamed properly by relayer.
	// for now, inject the log manually
	receipt, err := originChain.TransactionReceipt(c.GetTestContext(), tx.Hash())
	c.Nil(err)
	var sentLog *types.Log
	for _, log := range receipt.Logs {
		if log.Topics[0] == cctp.CircleRequestSentTopic {
			sentLog = log
			break
		}
	}
	err = relay.HandleLog(c.GetTestContext(), sentLog, uint32(originChainID.Int64()))
	c.Nil(err)

	// verify that the confirmed request is stored in the backend
	c.Eventually(func() bool {
		var storedMsg relayTypes.Message
		err = c.testStore.DB().Where("state = ?", relayTypes.Complete).Last(&storedMsg).Error
		if err != nil {
			return false
		}
		return storedMsg.OriginTxHash == tx.Hash().String()
	})

	// TODO: verify USDC is credited on recv chain
	// _, recvMockUsdc := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), destChain)
	// c.Nil(err)
	// expectedBalance := bridgeAmount
	// c.Eventually(func() bool {
	// 	balance, err := recvMockUsdc.BalanceOf(nil, opts.From)
	// 	c.Nil(err)
	// 	return c.Equal(expectedBalance, balance)
	// })
}
