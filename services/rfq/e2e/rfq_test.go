package e2e_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
	"testing"
	"time"
)

type IntegrationSuite struct {
	*testsuite.TestSuite
	manager       *testutil.DeployManager
	originBackend backends.SimulatedTestBackend
	destBackend   backends.SimulatedTestBackend
	//omniserver is the omnirpc server address
	omniServer    string
	omniClient    omnirpcClient.RPCClient
	metrics       metrics.Handler
	apiServer     string
	relayer       *relayer.Relayer
	relayerWallet wallet.Wallet
}

func NewIntegrationSuite(tb testing.TB) *IntegrationSuite {
	tb.Helper()
	return &IntegrationSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, NewIntegrationSuite(t))
}

const (
	originBackendChainID = 1
	destBackendChainID   = 2
)

// SetupTest sets up each test in the integration suite. We need to do a few things here:
//
// 1. Create the backends
// 2. Create a bunch of different tokens on a bunch of different chains. We do this here so we can pre-generate a lot of
// the tedious configurations for both api and relayer at the same time before testing individual cases.
// 3. Create the omnirpc server: this is used by both the api and the relayer.
func (i *IntegrationSuite) SetupTest() {
	i.TestSuite.SetupTest()

	i.manager = testutil.NewDeployManager(i.T())
	// TODO: consider jaeger
	i.metrics = metrics.NewNullHandler()
	// setup backends for ethereum & omnirpc
	i.setupBackends()
	// setup the api server
	i.setupAPI()
	i.setupRelayer()

}

// getOtherBackend gets the backend that is not the current one. This is a helper
func (i *IntegrationSuite) getOtherBackend(backend backends.SimulatedTestBackend) backends.SimulatedTestBackend {
	allBackends := core.ToSlice(i.originBackend, i.destBackend)
	for _, b := range allBackends {
		if b.GetChainID() != backend.GetChainID() {
			return b
		}
	}
	return nil
}

// TODO:
func (i *IntegrationSuite) TestUSDCtoUSDC() {
	time.Sleep(time.Second * 90)
}
