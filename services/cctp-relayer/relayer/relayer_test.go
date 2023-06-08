package relayer_test

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/cctp-relayer/api"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	scribeClient "github.com/synapsecns/sanguine/services/scribe/client"
)

func (c *CCTPRelayerSuite) TestSendCircleToken() {
	// setup
	sendChain := c.testBackends[0]
	recvChain := c.testBackends[1]
	_, cctpContractRef := c.deployManager.GetSynapseCCTP(c.GetTestContext(), sendChain)
	_, mintContractRef := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), sendChain)

	// create a relayer
	sendChainId, err := sendChain.ChainID(c.GetTestContext())
	c.Nil(err)
	cfg := config.Config{
		Chains: []config.ChainConfig{
			{
				ChainID: uint32(sendChainId.Int64()),
			},
		},
	}

	// TODO clean this part up
	parsedScribe, err := url.Parse(c.testScribe)
	c.Nil(err)
	port, err := strconv.Atoi(parsedScribe.Opaque)
	c.Nil(err)

	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
	mockApi := api.NewMockCircleApi()
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, sc.ScribeClient, c.metricsHandler, mockApi)
	c.Nil(err)

	relay.SetOmnirpcClient(omniClient.NewOmnirpcClient(c.testOmnirpc, c.metricsHandler, omniClient.WithCaptureReqRes()))

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
	tx, err = cctpContractRef.SendCircleToken(opts.TransactOpts, opts.From, uint32(recvChain.GetChainID()), mintContractRef.Address(), amount, 0, []byte{})
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

	// handle send request
	err = relay.HandleSendRequest(c.GetTestContext(), tx.Hash(), uint32(sendChain.GetChainID()))
	c.Nil(err)
	recvChan := relay.GetUsdcMsgRecvChan(uint32(sendChain.GetChainID()))
	msg := <-recvChan
	// TODO(dwasse): validate rest of msg?
	c.Equal(msg.TxHash, tx.Hash())
}

func (c *CCTPRelayerSuite) TestFetchAttestation() {
	// setup
	sendChain := c.testBackends[0]

	// create a relayer
	sendChainId, err := sendChain.ChainID(c.GetTestContext())
	c.Nil(err)
	cfg := config.Config{
		Chains: []config.ChainConfig{
			{
				ChainID: uint32(sendChainId.Int64()),
			},
		},
	}

	parsedScribe, err := url.Parse(c.testScribe)
	c.Nil(err)
	port, err := strconv.Atoi(parsedScribe.Opaque)
	c.Nil(err)

	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
	mockApi := api.NewMockCircleApi()
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, sc.ScribeClient, c.metricsHandler, mockApi)
	c.Nil(err)

	// override mocked api call
	expectedSignature := "abc"
	mockApi.SetGetAttestation(func(ctx context.Context, txHash common.Hash) (attestation []byte, err error) {
		return []byte(expectedSignature), nil
	})

	// fetch attestation
	testHash := "0x5dba62229dba62f233dca8f3fd14488fdc45d2a86537da2dea7a5683b5e7f622"
	msg := relayer.UsdcMessage{
		TxHash:        common.HexToHash(testHash),
		Message:       []byte{},
		AuxillaryData: []byte{},
	}
	relay.FetchAttestation(c.GetTestContext(), uint32(sendChain.GetChainID()), &msg)
	sendChan := relay.GetUsdcMsgSendChan(uint32(sendChain.GetChainID()))
	completeMsg := <-sendChan
	fmt.Printf("completeMsg: %v\n", completeMsg)
	// TODO(dwasse): validate rest of msg?
	c.Equal(completeMsg.TxHash, msg.TxHash)
	c.Equal(completeMsg.Signature, []byte(expectedSignature))
}
