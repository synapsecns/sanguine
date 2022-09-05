package chain_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/chain"
)

// TestChainWatcher tests watching the chain for events by emitting events.
func (s *ChainSuite) TestChainWatcher() {
	s.T().Skip("TODO. This is currently covered by the end-to-end test")
}

// make sure a chain client can be started form a url.
func (s *ChainSuite) TestNewFromURL() {
	rinkeby := preset.GetRinkeby().Geth(s.GetTestContext(), s.T())
	chnRPC, err := chain.NewFromURL(s.GetTestContext(), rinkeby.RPCAddress())
	Nil(s.T(), err)

	testID, err := chnRPC.ChainID(s.GetTestContext())

	Nil(s.T(), err)
	Equal(s.T(), testID, rinkeby.GetBigChainID(), testsuite.BigIntComparer())
}
