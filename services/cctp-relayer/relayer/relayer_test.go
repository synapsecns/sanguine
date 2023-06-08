package relayer_test

import "math/big"

func (c *CCTPRelayerSuite) TestSendCircleToken() {
	// setup
	sendChain := c.testBackends[0]
	recvChain := c.testBackends[1]
	_, cctpContractRef := c.deployManager.GetSynapseCCTP(c.GetTestContext(), sendChain)
	_, mintContractRef := c.deployManager.GetMockMintBurnTokenType(c.GetTestContext(), sendChain)

	// mint token
	opts := sendChain.GetTxContext(c.GetTestContext(), nil)
	amount := big.NewInt(1000000000000000000)
	mintContractRef.MintPublic(opts.TransactOpts, opts.From, amount)

	// approve token
	tx, err := mintContractRef.Approve(opts.TransactOpts, cctpContractRef.Address(), amount)
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)

	// send token
	tx, err = cctpContractRef.SendCircleToken(opts.TransactOpts, opts.From, uint32(recvChain.GetChainID()), mintContractRef.Address(), amount, 0, []byte{})
	c.Nil(err)
	sendChain.WaitForConfirmation(c.GetTestContext(), tx)
}
