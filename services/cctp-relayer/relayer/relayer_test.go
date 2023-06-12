package relayer_test

import (
	"net/url"
	"strconv"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	scribeClient "github.com/synapsecns/sanguine/services/scribe/client"
)

// func (c *CCTPRelayerSuite) TestHandleCircleRequestSent() {
// 	// setup
// 	sendChain := c.testBackends[0]
// 	recvChain := c.testBackends[1]
// 	_, cctpContractRef := c.deployManager.GetSynapseCCTP(c.GetTestContext(), sendChain)
// 	_, mintContractRef := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), sendChain)

// 	// create a relayer
// 	sendChainID, err := sendChain.ChainID(c.GetTestContext())
// 	c.Nil(err)
// 	testWallet, err := wallet.FromRandom()
// 	cfg := config.Config{
// 		DBPrefix: filet.TmpDir(c.T(), ""),
// 		Chains: []config.ChainConfig{
// 			{
// 				ChainID: uint32(sendChainID.Int64()),
// 			},
// 		},
// 		BaseOmnirpcURL: c.testBackends[0].RPCAddress(),
// 		Signer: signerConfig.SignerConfig{
// 			Type: signerConfig.FileType.String(),
// 			File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
// 		},
// 	}

// 	// TODO clean this part up
// 	parsedScribe, err := url.Parse(c.testScribe)
// 	c.Nil(err)
// 	port, err := strconv.Atoi(parsedScribe.Opaque)
// 	c.Nil(err)

// 	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
// 	mockAPI := api.NewMockCircleAPI()
// 	omniRPCClient := omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes())
// 	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, sc.ScribeClient, omniRPCClient, c.metricsHandler, mockAPI)
// 	c.Nil(err)

// 	// mint token
// 	opts := sendChain.GetTxContext(c.GetTestContext(), nil)
// 	amount := big.NewInt(1000000000000000000)
// 	tx, err := mintContractRef.MintPublic(opts.TransactOpts, opts.From, amount)
// 	c.Nil(err)
// 	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

// 	// approve token
// 	tx, err = mintContractRef.Approve(opts.TransactOpts, cctpContractRef.Address(), amount)
// 	c.Nil(err)
// 	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

// 	// send token
// 	tx, err = cctpContractRef.SendCircleToken(opts.TransactOpts, opts.From, big.NewInt(int64(recvChain.GetChainID())), mintContractRef.Address(), amount, 0, []byte{})
// 	c.Nil(err)
// 	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

// 	// handle send request
// 	err = relay.HandleCircleRequestSent(c.GetTestContext(), tx.Hash(), uint32(sendChain.GetChainID()))
// 	c.Nil(err)
// 	recvChan := relay.GetUsdcMsgRecvChan(uint32(sendChain.GetChainID()))
// 	msg := <-recvChan
// 	// TODO(dwasse): validate rest of msg?
// 	c.Equal(msg.MessageHash, tx.Hash())
// }

// func (c *CCTPRelayerSuite) TestFetchAttestation() {
// 	// setup
// 	sendChain := c.testBackends[0]

// 	// create a relayer
// 	sendChainID, err := sendChain.ChainID(c.GetTestContext())
// 	c.Nil(err)
// 	testWallet, err := wallet.FromRandom()
// 	cfg := config.Config{
// 		DBPrefix: filet.TmpDir(c.T(), ""),
// 		Chains: []config.ChainConfig{
// 			{
// 				ChainID: uint32(sendChainID.Int64()),
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

// 	// override mocked api call
// 	expectedSignature := "abc"
// 	mockAPI.SetGetAttestation(func(ctx context.Context, txHash common.Hash) (attestation []byte, err error) {
// 		return []byte(expectedSignature), nil
// 	})

// 	// fetch attestation
// 	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
// 	msg := relayer.UsdcMessage{
// 		Message:          []byte{},
// 		MessageHash:      common.HexToHash(testHash),
// 		FormattedRequest: []byte{},
// 	}
// 	err = relay.FetchAttestation(c.GetTestContext(), uint32(sendChain.GetChainID()), &msg)
// 	c.Nil(err)

// 	sendChan := relay.GetUsdcMsgSendChan(uint32(sendChain.GetChainID()))
// 	completeMsg := <-sendChan
// 	fmt.Printf("completeMsg: %v\n", completeMsg)
// 	// TODO(dwasse): validate rest of msg?
// 	c.Equal(completeMsg.MessageHash, msg.MessageHash)
// 	c.Equal(completeMsg.Signature, []byte(expectedSignature))
// }

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
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, sc.ScribeClient, omniRPCClient, c.metricsHandler, mockAPI)
	c.Nil(err)

	// submit receive circle token
	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
	msg := relayer.UsdcMessage{
		ChainID:          uint32(recvChainID.Int64()),
		Message:          []byte{},
		MessageHash:      common.HexToHash(testHash),
		FormattedRequest: []byte{},
	}
	err = relay.SubmitReceiveCircleToken(c.GetTestContext(), &msg)
	c.Nil(err)
}
