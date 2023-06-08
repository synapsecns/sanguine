package relayer_test

import (
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/services/cctp-relayer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/cctp-relayer/testutil"
)

func (c *CCTPRelayerSuite) TestSendCircleToken() {
	sendChain := c.testBackends[0]
	recvChain := c.testBackends[1]
	_, contractRef := manager.GetContract[*cctp.SynapseCCTPRef](c.GetTestContext(), c.T(), c.deployManager, sendChain, testutil.SynapseCCTPType)

	opts := sendChain.GetTxContext(c.GetTestContext(), nil)
	contractRef.SendCircleToken(opts, opts.From)
}
