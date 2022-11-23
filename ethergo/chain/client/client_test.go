package client_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"go.uber.org/atomic"
	"os"
	"time"
)

// TestAttemptReconnect tests the rehupping process. You may notice that in the below test, if the first connection
// had broken no new blocks would come in on the subcription. This is expected and PoolClient works around this
// we're specifically testing if our rehup process breaks any existing subscriptions if the connection didn't break
// this stands in for all other types of one off requests.
func (c *ClientSuite) TestAttemptReconnect() {
	if os.Getenv("ATTEMPT_CONNECT_ENABLED") != "" {
		c.T().Skip("deprecated")
	}
	testBackend := preset.GetRinkeby().Geth(c.GetTestContext(), c.T())

	// make sure the client doesn't reset on us during the test
	client.SetResetTimeout(time.Hour)

	// create a subscription before rehupping to make sure we don't panic/break the original subscription
	testClient, err := client.NewTestClient(c.GetTestContext(), c.T(), testBackend.WSEndpoint())
	Nil(c.T(), err)

	ogSubscription := c.maxHeightSubscription(testClient)

	// reconnect to see if ogSubscription breaks
	err = testClient.AttemptReconnect()
	Nil(c.T(), err)

	postRehupSubscription := c.maxHeightSubscription(testClient)
	Nil(c.T(), err)

	// mock some blocks to make sure our subscription gets updated
	mocks.MockBlocksOnBackend(c.GetTestContext(), c.T(), testBackend, 10)

	c.Eventually(func() bool {
		return ogSubscription.MaxHeight() == postRehupSubscription.MaxHeight()
	})
}

// maxHeightTracker fetches the max height from haed.
type maxHeightTracker struct {
	maxHeight atomic.Uint64
}

// MaxHeight gets the current max height in a thread safe manner.
func (m *maxHeightTracker) MaxHeight() uint64 {
	return m.maxHeight.Load()
}

func (c *ClientSuite) maxHeightSubscription(testClient client.TestClient) *maxHeightTracker {
	tracker := maxHeightTracker{}

	// headerChan is a channel containing headerChan from the subscription on the client
	headerChan := make(chan *types.Header)
	sub, err := testClient.SubscribeNewHead(c.GetTestContext(), headerChan)
	Nil(c.T(), err)

	// start a goroutine that updates the max height
	go func() {
		for {
			select {
			case <-c.GetTestContext().Done():
				return
			case err := <-sub.Err():
				// should not error because of a broken client
				Nil(c.T(), err)
			case header := <-headerChan:
				tracker.maxHeight.Store(header.Number.Uint64())
			}
		}
	}()
	return &tracker
}

func (c *ClientSuite) TestBatch() {
	backend := geth.NewEmbeddedBackend(c.GetTestContext(), c.T())

	caller, err := client.NewClient(c.GetTestContext(), backend.RPCAddress())
	Nil(c.T(), err)

	const callCount = 10

	calls := make([]w3types.Caller, callCount)
	res := make([]uint64, callCount)

	for i := 0; i < callCount; i++ {
		calls[i] = eth.ChainID().Returns(&res[i])
	}

	err = caller.BatchContext(c.GetTestContext(), calls...)
	Nil(c.T(), err)

	for _, chainID := range res {
		Equal(c.T(), chainID, uint64(backend.GetChainID()))
	}
}
