package relayer_test

import (
	"context"
	"math/big"
	"net/url"
	"strconv"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	relayTypes "github.com/synapsecns/sanguine/services/cctp-relayer/types"
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
			File: filet.TmpFile(c.T(), "", c.testWallet.PrivateKeyHex()).Name(),
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
	c.Equal(msg.State, relayTypes.Pending)

	// verify that the request is stored in the db
	var storedMsg relayTypes.Message
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
			File: filet.TmpFile(c.T(), "", c.testWallet.PrivateKeyHex()).Name(),
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
	msg := relayTypes.Message{
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
	c.Equal(completeMsg.State, relayTypes.Attested)

	// verify that the attested request is stored in the db
	var storedMsg relayTypes.Message
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
			File: filet.TmpFile(c.T(), "", c.testWallet.PrivateKeyHex()).Name(),
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
	msg := relayTypes.Message{
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
	var storedMsg relayTypes.Message
	err = c.testStore.DB().Where("origin_tx_hash = ?", msg.OriginTxHash).First(&storedMsg).Error
	c.Nil(err)
	msg.State = relayTypes.Complete
	msg.DestTxHash = storedMsg.DestTxHash
	c.Equal(msg, storedMsg)
}

func getTestConfig(c *CCTPRelayerSuite, backends []backends.SimulatedTestBackend) config.Config {
	cfg := config.Config{
		DBPrefix:       filet.TmpDir(c.T(), ""),
		BaseOmnirpcURL: c.testBackends[0].RPCAddress(),
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(c.T(), "", c.testWallet.PrivateKeyHex()).Name(),
		},
	}
	chains := []config.ChainConfig{}
	for _, backend := range backends {
		_, handle := c.deployManager.GetSynapseCCTP(c.GetTestContext(), backend)
		chains = append(chains, config.ChainConfig{
			ChainID:            uint32(backend.GetChainID()),
			SynapseCCTPAddress: handle.Address().String(),
		})
	}
	cfg.Chains = chains
	return cfg
}

func (c *CCTPRelayerSuite) TestBridgeUSDC() {
	// setup
	sendChain := c.testBackends[0]
	recvChain := c.testBackends[1]

	// create a relayer
	sendChainID, err := sendChain.ChainID(c.GetTestContext())
	recvChainID, err := recvChain.ChainID(c.GetTestContext())
	c.Nil(err)
	cfg := getTestConfig(c, []backends.SimulatedTestBackend{sendChain, recvChain})

	parsedScribe, err := url.Parse(c.testScribe)
	c.Nil(err)
	port, err := strconv.Atoi(parsedScribe.Opaque)
	c.Nil(err)

	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
	mockAPI := api.NewMockCircleAPI()
	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, c.testStore, sc.ScribeClient, omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// start relayer
	ctx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()
	go relay.Run(ctx)

	// mint some USDC on send chain
	_, sendMockUsdc := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), sendChain)
	sendTxOpts, err := bind.NewKeyedTransactorWithChainID(c.testWallet.PrivateKey(), sendChainID)
	c.Nil(err)
	bridgeAmount := big.NewInt(1000000000) // 1000 USDC
	tx, err := sendMockUsdc.MintPublic(sendTxOpts, c.testWallet.Address(), bridgeAmount)
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

	// approve USDC for spending
	_, sendSynapseCCTP := c.deployManager.GetSynapseCCTP(c.GetTestContext(), sendChain)
	tx, err = sendMockUsdc.Approve(sendTxOpts, sendSynapseCCTP.Address(), bridgeAmount)
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

	// send USDC from sendChain
	tx, err = sendSynapseCCTP.SendCircleToken(sendTxOpts, c.testWallet.Address(), recvChainID, sendMockUsdc.Address(), bridgeAmount, 0, []byte{})
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

	// TODO: figure out why log is not streamed properly by relayer.
	// for now, inject the log manually
	receipt, err := sendChain.TransactionReceipt(c.GetTestContext(), tx.Hash())
	c.Nil(err)
	var sentLog *types.Log
	for _, log := range receipt.Logs {
		if log.Topics[0] == cctp.CircleRequestSentTopic {
			sentLog = log
			break
		}
	}
	err = relay.HandleLog(c.GetTestContext(), sentLog, uint32(sendChainID.Int64()))
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
	// _, recvMockUsdc := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), recvChain)
	// c.Nil(err)
	// expectedBalance := bridgeAmount
	// c.Eventually(func() bool {
	// 	balance, err := recvMockUsdc.BalanceOf(nil, c.testWallet.Address())
	// 	c.Nil(err)
	// 	return c.Equal(expectedBalance, balance)
	// })
}
