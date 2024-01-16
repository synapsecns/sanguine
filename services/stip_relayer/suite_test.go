package stip_relayer_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/stip_relayer/db"
)

type STIPRelayerSuite struct {
	*testsuite.TestSuite
	omniRPCClient       omniClient.RPCClient
	OmniRPCTestBackends []backends.SimulatedTestBackend
	database            db.STIPDB
	cfg                 config.Config
	testWallet          wallet.Wallet
	handler             metrics.Handler
}

func NewSTIPRelayerSuite(tb testing.TB) *STIPRelayerSuite {
	tb.Helper()
	return &STIPRelayerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (c *STIPRelayerSuite) SetupTest() {
	c.TestSuite.SetupTest()

	// c.cfg = config.Config{
	// 	DatabaseURL: c.DatabaseURL(),
	// }

	// c.database = db.NewSTIPDB(c.GetTestContext(), c.T(), c.cfg.DatabaseURL)
	// c.handler = metrics.NewNoopHandler()

	// c.testWallet = wallet.NewWallet(c.GetTestContext(), c.T(), c.OmniRPCTestBackends[0], c.handler)

	// c.omniRPCClient = omniClient.NewOmnirpcClient(c.OmniRPCTestBackends[0], c.handler, omniClient.WithCaptureReqRes())
}

func (c *STIPRelayerSuite) SetupSuite() {
	c.TestSuite.SetupSuite()

	// c.port = uint16(freeport.GetPort())
	// c.cfg = config.Config{
	// 	DatabaseURL: c.DatabaseURL(),
	// 	APIPort:     c.port,
	// }

	// c.database = sql.NewAPIDB(c.GetTestContext(), c.T(), c.cfg.DatabaseURL)
	// c.handler = metrics.NewNoopHandler()

	// c.testWallet = wallet.NewWallet(c.GetTestContext(), c.T(), c.OmniRPCTestBackends[0], c.handler)

	// c.omniRPCClient = omniClient.NewOmnirpcClient(c.OmniRPCTestBackends[0], c.handler, omniClient.WithCaptureReqRes())
}

// TestConfigSuite runs the integration test suite.
func TestSTIPRelayerSuite(t *testing.T) {
	suite.Run(t, NewSTIPRelayerSuite(t))
}
