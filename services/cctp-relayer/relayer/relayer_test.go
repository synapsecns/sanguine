package relayer_test

import (
	"context"
	"math/big"
	"net/url"
	"strconv"

	"github.com/Flaque/filet"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	"github.com/synapsecns/sanguine/services/cctp-relayer/types"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	scribeClient "github.com/synapsecns/sanguine/services/scribe/client"
)

func (c *CCTPRelayerSuite) TestHandleCircleRequestSent() {
	// setup
	sendChain := c.testBackends[0]
	recvChain := c.testBackends[1]
	_, cctpContractRef := c.deployManager.GetSynapseCCTP(c.GetTestContext(), sendChain)
	_, mintContractRef := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), sendChain)

	// create a relayer
	sendChainID, err := sendChain.ChainID(c.GetTestContext())
	c.Nil(err)
	testWallet, err := wallet.FromRandom()
	cfg := config.Config{
		DBPrefix: filet.TmpDir(c.T(), ""),
		Chains: []config.ChainConfig{
			{
				ChainID: uint32(sendChainID.Int64()),
			},
		},
		BaseOmnirpcURL: c.testBackends[0].RPCAddress(),
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
		},
	}

	// TODO clean this part up
	parsedScribe, err := url.Parse(c.testScribe)
	c.Nil(err)
	port, err := strconv.Atoi(parsedScribe.Opaque)
	c.Nil(err)

	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
	mockAPI := api.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, c.testStore, sc.ScribeClient, omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// mint token
	opts := sendChain.GetTxContext(c.GetTestContext(), nil)
	amount := big.NewInt(1000000000000000000)
	tx, err := mintContractRef.MintPublic(opts.TransactOpts, opts.From, amount)
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

	// approve token
	tx, err = mintContractRef.Approve(opts.TransactOpts, cctpContractRef.Address(), amount)
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

	// send token
	tx, err = cctpContractRef.SendCircleToken(opts.TransactOpts, opts.From, big.NewInt(int64(recvChain.GetChainID())), mintContractRef.Address(), amount, 0, []byte{})
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

	// handle send request
	err = relay.HandleCircleRequestSent(c.GetTestContext(), tx.Hash(), uint32(sendChain.GetChainID()))
	c.Nil(err)
	recvChan := relay.GetUsdcMsgRecvChan(uint32(sendChain.GetChainID()))
	msg := <-recvChan
	c.Equal(msg.OriginTxHash, tx.Hash().String())
	c.Equal(msg.State, types.Pending)

	// verify that the request is stored in the db
	var storedMsg types.Message
	err = c.testStore.DB().Where("origin_tx_hash = ?", msg.OriginTxHash).First(&storedMsg).Error
	c.Nil(err)
	c.Equal(*msg, storedMsg)
}

func (c *CCTPRelayerSuite) TestFetchAttestation() {
	// setup
	sendChain := c.testBackends[0]

	// create a relayer
	sendChainID, err := sendChain.ChainID(c.GetTestContext())
	c.Nil(err)
	testWallet, err := wallet.FromRandom()
	cfg := config.Config{
		DBPrefix: filet.TmpDir(c.T(), ""),
		Chains: []config.ChainConfig{
			{
				ChainID: uint32(sendChainID.Int64()),
			},
		},
		BaseOmnirpcURL: c.testBackends[0].RPCAddress(),
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
		},
	}

	parsedScribe, err := url.Parse(c.testScribe)
	c.Nil(err)
	port, err := strconv.Atoi(parsedScribe.Opaque)
	c.Nil(err)

	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
	mockAPI := api.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, c.testStore, sc.ScribeClient, omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// override mocked api call
	expectedSignature := "abc"
	mockAPI.SetGetAttestation(func(ctx context.Context, txHash string) (attestation []byte, err error) {
		return []byte(expectedSignature), nil
	})

	// fetch attestation
	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
	msg := types.Message{
		Message:          []byte{},
		MessageHash:      testHash,
		FormattedRequest: []byte{},
	}
	err = relay.FetchAttestation(c.GetTestContext(), uint32(sendChain.GetChainID()), &msg)
	c.Nil(err)

	sendChan := relay.GetUsdcMsgSendChan(uint32(sendChain.GetChainID()))
	completeMsg := <-sendChan
	// TODO(dwasse): validate rest of msg?
	c.Equal(completeMsg.MessageHash, msg.MessageHash)
	c.Equal(completeMsg.Signature, []byte(expectedSignature))
	c.Equal(completeMsg.State, types.Attested)

	// verify that the attested request is stored in the db
	var storedMsg types.Message
	err = c.testStore.DB().Where("origin_tx_hash = ?", completeMsg.OriginTxHash).First(&storedMsg).Error
	c.Nil(err)
	c.Equal(*completeMsg, storedMsg)
}

func (c *CCTPRelayerSuite) TestSubmitReceiveCircleToken() {
	// setup
	sendChain := c.testBackends[0]
	recvChain := c.testBackends[1]

	// create a relayer
	sendChainID, err := sendChain.ChainID(c.GetTestContext())
	recvChainID, err := recvChain.ChainID(c.GetTestContext())
	c.Nil(err)
	testWallet, err := wallet.FromRandom()
	cfg := config.Config{
		DBPrefix: filet.TmpDir(c.T(), ""),
		Chains: []config.ChainConfig{
			{
				ChainID: uint32(sendChainID.Int64()),
			},
			{
				ChainID: uint32(recvChainID.Int64()),
			},
		},
		BaseOmnirpcURL: c.testBackends[0].RPCAddress(),
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
		},
	}

	parsedScribe, err := url.Parse(c.testScribe)
	c.Nil(err)
	port, err := strconv.Atoi(parsedScribe.Opaque)
	c.Nil(err)

	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
	mockAPI := api.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, c.testStore, sc.ScribeClient, omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// submit receive circle token
	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
	msg := types.Message{
		OriginTxHash:     testHash,
		OriginChainID:    uint32(sendChainID.Int64()),
		DestChainID:      uint32(recvChainID.Int64()),
		Message:          []byte{},
		MessageHash:      testHash,
		FormattedRequest: []byte{},
	}
	err = relay.SubmitReceiveCircleToken(c.GetTestContext(), &msg)
	c.Nil(err)

	// verify that the attested request is stored in the db
	var storedMsg types.Message
	err = c.testStore.DB().Where("origin_tx_hash = ?", msg.OriginTxHash).First(&storedMsg).Error
	c.Nil(err)
	msg.State = types.Complete
	msg.DestTxHash = storedMsg.DestTxHash
	c.Equal(msg, storedMsg)
}

// func (c *CCTPRelayerSuite) TestBridgeUSDC() {
// 	// setup
// 	sendChain := c.testBackends[0]
// 	recvChain := c.testBackends[1]

// 	// create a relayer
// 	sendChainID, err := sendChain.ChainID(c.GetTestContext())
// 	recvChainID, err := recvChain.ChainID(c.GetTestContext())
// 	c.Nil(err)
// 	testWallet, err := wallet.FromRandom()
// 	cfg := config.Config{
// 		DBPrefix: filet.TmpDir(c.T(), ""),
// 		Chains: []config.ChainConfig{
// 			{
// 				ChainID: uint32(sendChainID.Int64()),
// 			},
// 			{
// 				ChainID: uint32(recvChainID.Int64()),
// 			},
// 		},
// 		BaseOmnirpcURL: c.testBackends[0].RPCAddress(),
// 		Signer: signerConfig.SignerConfig{
// 			Type: signerConfig.FileType.String(),
// 			File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
// 		},
// 	}

// 	parsedScribe, err := url.Parse(c.testScribe)
// 	c.Nil(err)
// 	port, err := strconv.Atoi(parsedScribe.Opaque)
// 	c.Nil(err)

// 	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
// 	mockAPI := api.NewMockCircleAPI()
// 	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
// 	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, sc.ScribeClient, omniRPCClient, c.metricsHandler, mockAPI)
// 	c.Nil(err)

// 	// start relayer
// 	ctx, cancel := context.WithCancel(c.GetTestContext())
// 	defer cancel()
// 	relay.Run(ctx)

// 	// mint some USDC on send chain
// 	_, sendMockUsdc := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), sendChain)
// 	sendTxOpts, err := bind.NewKeyedTransactorWithChainID(testWallet.PrivateKey(), sendChainID)
// 	c.Nil(err)
// 	bridgeAmount := big.NewInt(1000000000) // 1000 USDC
// 	tx, err := sendMockUsdc.Mint(sendTxOpts, testWallet.Address(), bridgeAmount)
// 	c.Nil(err)
// 	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

// 	// send USDC from sendChain
// 	_, sendSynapseCCTP := c.deployManager.GetSynapseCCTP(c.GetTestContext(), sendChain)
// 	c.Nil(err)
// 	tx, err = sendSynapseCCTP.SendCircleToken(sendTxOpts, testWallet.Address(), recvChainID, sendMockUsdc.Address(), bridgeAmount, 0, []byte{})
// 	c.Nil(err)
// 	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

// 	// verify USDC is credited on recv chain
// 	_, recvMockUsdc := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), recvChain)
// 	c.Nil(err)
// 	expectedBalance := bridgeAmount
// 	c.Eventually(func() bool {
// 		balance, err := recvMockUsdc.BalanceOf(nil, testWallet.Address())
// 		c.Nil(err)
// 		return c.Equal(expectedBalance, balance)
// 	})
// }
