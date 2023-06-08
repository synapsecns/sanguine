package relayer_test

import (
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	scribeClient "github.com/synapsecns/sanguine/services/scribe/client"
	"math/big"
	"net/url"
	"strconv"
)

func (c *CCTPRelayerSuite) TestSendCircleToken() {
	// setup
	sendChain := c.testBackends[0]
	recvChain := c.testBackends[1]
	_, cctpContractRef := c.deployManager.GetSynapseCCTP(c.GetTestContext(), sendChain)
	_, mintContractRef := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), sendChain)

	// create a relayer
	// TODO: figure out
	cfg := config.Config{}

	// TODO clean this part up
	parsedScribe, err := url.Parse(c.testScribe)
	c.Nil(err)
	port, err := strconv.Atoi(parsedScribe.Opaque)
	c.Nil(err)

	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, c.metricsHandler)
	relay, err := relayer.NewCCTPRelayer(c.GetTestContext(), cfg, sc.ScribeClient, c.metricsHandler)
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

	err = relay.HandleSendRequest(c.GetTestContext(), tx.Hash(), uint32(sendChain.GetChainID()))
	c.Nil(err)

}
