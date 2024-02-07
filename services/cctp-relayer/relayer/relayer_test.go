package relayer_test

import (
	"context"
	"math/big"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/cctp-relayer/attestation"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

func (s *CCTPRelayerSuite) mockMessage(originChainID, destinationChainID, blockNumber uint32) relayTypes.Message {
	return relayTypes.Message{
		OriginTxHash:     mocks.NewMockHash(s.T()).String(),
		DestTxHash:       mocks.NewMockHash(s.T()).String(),
		OriginChainID:    originChainID,
		DestChainID:      destinationChainID,
		Message:          []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		MessageHash:      mocks.NewMockHash(s.T()).String(),
		Attestation:      []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		RequestVersion:   0,
		FormattedRequest: []byte(gofakeit.Paragraph(10, 10, 10, " ")),
		RequestID:        strings.TrimPrefix(mocks.NewMockHash(s.T()).String(), "0x"),
		BlockNumber:      uint64(blockNumber),
		State:            relayTypes.Submitted,
	}
}

func (s *CCTPRelayerSuite) TestHandleCircleRequestSent() {
	// setup
	originChain := s.testBackends[0]
	destChain := s.testBackends[1]
	_, originSynapseCCTP := s.deployManager.GetSynapseCCTP(s.GetTestContext(), originChain)
	_, originMockUsdc := s.deployManager.GetMockMintBurnTokenType(s.GetTestContext(), originChain)

	// create a new relayer
	mockAPI := attestation.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(s.testOmnirpc, s.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(s.GetTestContext(), s.GetTestConfig(), s.testStore, s.GetTestScribe(), omniRPCClient, s.metricsHandler, mockAPI)
	s.Nil(err)

	// mint token
	opts := originChain.GetTxContext(s.GetTestContext(), nil)
	amount := big.NewInt(1000000000000000000)
	tx, err := originMockUsdc.MintPublic(opts.TransactOpts, opts.From, amount)
	s.Nil(err)
	originChain.WaitForConfirmation(s.GetTestContext(), tx)

	// approve token
	tx, err = originMockUsdc.Approve(opts.TransactOpts, originSynapseCCTP.Address(), amount)
	s.Nil(err)
	originChain.WaitForConfirmation(s.GetTestContext(), tx)

	// send token
	tx, err = originSynapseCCTP.SendCircleToken(opts.TransactOpts, opts.From, big.NewInt(int64(destChain.GetChainID())), originMockUsdc.Address(), amount, 0, []byte{})
	s.Nil(err)
	originChain.WaitForConfirmation(s.GetTestContext(), tx)

	// handle send request
	msg, err := relay.HandleCircleRequestSent(s.GetTestContext(), tx.Hash(), uint32(originChain.GetChainID()))
	s.Nil(err)
	s.Equal(msg.OriginTxHash, tx.Hash().String())
	s.Equal(msg.State, relayTypes.Pending)

	// verify that the request is stored in the db
	var storedMsg relayTypes.Message
	err = s.testStore.DB().Where("origin_tx_hash = ?", msg.OriginTxHash).First(&storedMsg).Error
	s.Nil(err)
	s.Equal(*msg, storedMsg)
}

func (s *CCTPRelayerSuite) TestStoreCircleRequestFulfilled() {
	// setup
	originChain := s.testBackends[0]
	destChain := s.testBackends[1]

	// create a new relayer
	mockAPI := attestation.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(s.testOmnirpc, s.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(s.GetTestContext(), s.GetTestConfig(), s.testStore, s.GetTestScribe(), omniRPCClient, s.metricsHandler, mockAPI)
	s.Nil(err)

	// inject a fulfilled event to be handled by relayer
	blockNumber, err := destChain.BlockNumber(s.GetTestContext())
	s.Nil(err)
	message1 := s.mockMessage(uint32(originChain.GetChainID()), uint32(destChain.GetChainID()), uint32(blockNumber))
	fulfilledLog := &types.Log{
		TxHash:      common.HexToHash(message1.DestTxHash),
		BlockNumber: blockNumber,
	}
	var requestID [32]byte
	requestIDBytes := common.Hex2Bytes(message1.RequestID)
	copy(requestID[:], requestIDBytes)
	fulfilledEvent := &cctp.SynapseCCTPEventsCircleRequestFulfilled{
		OriginDomain: message1.OriginChainID,
		RequestID:    requestID,
	}
	err = relay.StoreCircleRequestFulfilled(s.GetTestContext(), fulfilledLog, fulfilledEvent, uint32(destChain.GetChainID()))
	s.Nil(err)

	// verify that the request is stored as Complete in the db
	var storedMsg relayTypes.Message
	err = s.testStore.DB().Where("request_id = ?", message1.RequestID).First(&storedMsg).Error
	s.Nil(err)
	s.Equal(message1.DestTxHash, storedMsg.DestTxHash)
	s.Equal(message1.OriginChainID, storedMsg.OriginChainID)
	s.Equal(message1.DestChainID, storedMsg.DestChainID)
	s.Equal(message1.RequestID, storedMsg.RequestID)
	s.Equal(message1.BlockNumber, storedMsg.BlockNumber)
	s.Equal(relayTypes.Complete, storedMsg.State)

	// store a Submitted message
	message2 := s.mockMessage(uint32(originChain.GetChainID()), uint32(destChain.GetChainID()), uint32(blockNumber))
	err = s.testStore.DB().Create(&message2).Error
	s.Nil(err)

	// process the corresponding fulfilled event
	requestIDBytes = common.Hex2Bytes(message2.RequestID)
	copy(requestID[:], requestIDBytes)
	fulfilledEvent = &cctp.SynapseCCTPEventsCircleRequestFulfilled{
		RequestID: requestID,
	}
	err = relay.StoreCircleRequestFulfilled(s.GetTestContext(), &types.Log{TxHash: common.HexToHash(message2.DestTxHash)}, fulfilledEvent, uint32(destChain.GetChainID()))
	s.Nil(err)

	// verify that the message is stored as Complete in the db
	err = s.testStore.DB().Where("request_id = ?", message2.RequestID).First(&storedMsg).Error
	s.Nil(err)
	message2.State = relayTypes.Complete
	s.Equal(message2, storedMsg)
}

func (s *CCTPRelayerSuite) TestFetchAttestation() {
	// create a new relayer
	mockAPI := attestation.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(s.testOmnirpc, s.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(s.GetTestContext(), s.GetTestConfig(), s.testStore, s.GetTestScribe(), omniRPCClient, s.metricsHandler, mockAPI)
	s.Nil(err)

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
	completeMsg, err := relay.FetchAttestation(s.GetTestContext(), &msg)
	s.Nil(err)

	s.Equal(completeMsg.MessageHash, msg.MessageHash)
	s.Equal(completeMsg.Attestation, []byte(expectedSignature))
	s.Equal(completeMsg.State, relayTypes.Attested)

	// verify that the attested request is stored in the db
	var storedMsg relayTypes.Message
	err = s.testStore.DB().Where("origin_tx_hash = ?", completeMsg.OriginTxHash).First(&storedMsg).Error
	s.Nil(err)
	s.Equal(*completeMsg, storedMsg)
}

func (s *CCTPRelayerSuite) TestSubmitReceiveCircleToken() {
	// create a new relayer
	mockAPI := attestation.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(s.testOmnirpc, s.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(s.GetTestContext(), s.GetTestConfig(), s.testStore, s.GetTestScribe(), omniRPCClient, s.metricsHandler, mockAPI)
	s.Nil(err)

	// build test msg
	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
	originChain := s.testBackends[0]
	originChainID, err := originChain.ChainID(s.GetTestContext())
	s.Nil(err)
	destChain := s.testBackends[1]
	destChainID, err := destChain.ChainID(s.GetTestContext())
	s.Nil(err)
	msg := relayTypes.Message{
		OriginTxHash:     testHash,
		OriginChainID:    uint32(originChainID.Int64()),
		DestChainID:      uint32(destChainID.Int64()),
		Message:          []byte{},
		MessageHash:      testHash,
		FormattedRequest: []byte{},
	}

	// submit ReceiveCircleToken()
	// nolint: wrapcheck
	err = relay.SubmitReceiveCircleToken(s.GetTestContext(), &msg)
	s.Nil(err)

	// verify that the attested request is stored in the db
	var storedMsg relayTypes.Message
	err = s.testStore.DB().Where("origin_tx_hash = ?", msg.OriginTxHash).First(&storedMsg).Error
	s.Nil(err)
	msg.State = relayTypes.Submitted
	s.Equal(msg, storedMsg)
}

func (s *CCTPRelayerSuite) TestBridgeUSDC() {
	// create a new relayer
	mockAPI := attestation.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(s.testOmnirpc, s.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(s.GetTestContext(), s.GetTestConfig(), s.testStore, s.GetTestScribe(), omniRPCClient, s.metricsHandler, mockAPI)
	s.Nil(err)

	// start relayer
	ctx, cancel := context.WithCancel(s.GetTestContext())
	defer cancel()
	//nolint:errcheck
	go relay.Run(ctx)

	// mint some USDC on send chain
	originChain := s.testBackends[0]
	destChain := s.testBackends[1]
	_, originMockUsdc := s.deployManager.GetMockMintBurnTokenType(s.GetTestContext(), originChain)
	originChainID, err := originChain.ChainID(s.GetTestContext())
	s.Nil(err)
	bridgeAmount := big.NewInt(1000000000) // 1000 USDC
	opts := originChain.GetTxContext(s.GetTestContext(), nil)
	tx, err := originMockUsdc.MintPublic(opts.TransactOpts, opts.From, bridgeAmount)
	s.Nil(err)
	originChain.WaitForConfirmation(s.GetTestContext(), tx)

	// approve USDC for spending
	_, originSynapseCCTP := s.deployManager.GetSynapseCCTP(s.GetTestContext(), originChain)
	tx, err = originMockUsdc.Approve(opts.TransactOpts, originSynapseCCTP.Address(), bridgeAmount)
	s.Nil(err)
	originChain.WaitForConfirmation(s.GetTestContext(), tx)

	// send USDC from originChain
	destChainID, err := destChain.ChainID(s.GetTestContext())
	s.Nil(err)
	tx, err = originSynapseCCTP.SendCircleToken(opts.TransactOpts, opts.From, destChainID, originMockUsdc.Address(), bridgeAmount, 0, []byte{})
	s.Nil(err)
	originChain.WaitForConfirmation(s.GetTestContext(), tx)

	// TODO: figure out why log is not streamed properly by relayer.
	// for now, inject the log manually
	receipt, err := originChain.TransactionReceipt(s.GetTestContext(), tx.Hash())
	s.Nil(err)
	var sentLog *types.Log
	for _, log := range receipt.Logs {
		if log.Topics[0] == cctp.CircleRequestSentTopic {
			sentLog = log
			break
		}
	}
	err = relay.HandleLog(s.GetTestContext(), sentLog, uint32(originChainID.Int64()))
	s.Require().Nil(err)

	// verify that the confirmed request is stored in the backend
	s.Eventually(func() bool {
		var storedMsg relayTypes.Message
		// TODO: shuld make this check for completion
		err = s.testStore.DB().Where("state = ?", relayTypes.Submitted).Last(&storedMsg).Error
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
