package client_test

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/rfq/quoting-api/client"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/config"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/rest"
	"github.com/synapsecns/sanguine/rfq/quoting-api/internal/testutil"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"testing"
	"time"
)

type ClientSuite struct {
	*testsuite.TestSuite
	// anvilBackend is currently required because we are unable to trivially inject
	// *bind.ContractBackend into the rest api server.
	anvilBackend  backends.SimulatedTestBackend
	testHandler   testutil.ITestContractHandler
	restAPIServer *rest.APIServer
	client        client.Client
}

func NewTestClientSuite(tb testing.TB) *ClientSuite {
	tb.Helper()

	return &ClientSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

const chainID = 100

func (c *ClientSuite) SetupSuite() {
	c.TestSuite.SetupSuite()
	options := anvil.NewAnvilOptionBuilder()
	options.SetChainID(chainID)
	c.anvilBackend = anvil.NewAnvilBackend(c.GetSuiteContext(), c.T(), options)
}

func (c *ClientSuite) SetupTest() {
	c.TestSuite.SetupTest()

	testWallet, err := wallet.FromRandom()
	assert.NoError(c.T(), err)

	c.testHandler, err = testutil.NewTestContractHandlerImpl(c.GetTestContext(), c.anvilBackend, testWallet, chainID)
	assert.NoError(c.T(), err)

	testServer := testhelper.NewOmnirpcServer(c.GetTestContext(), c.T(), c.anvilBackend, c.anvilBackend)
	port, err := freeport.GetFreePort()
	assert.NoError(c.T(), err)

	c.restAPIServer, err = rest.NewRestAPIServer(c.GetTestContext(), &config.Config{
		AuthExpiryDelta: 60000,
		Port:            uint16(port),
		DBType:          "sqlite", // TODO: use constant
		DSN:             filet.TmpFile(c.T(), "", "").Name(),
		OmniRPCURL:      testServer,
		Bridges: map[uint32]string{
			chainID: c.testHandler.FastBridgeAddress().String(),
		},
	})
	assert.NoError(c.T(), err)

	c.restAPIServer.Setup()
	c.client, err = client.NewClient(fmt.Sprintf("http://127.0.0.1:%d", port), localsigner.NewSigner(testWallet.PrivateKey()), testWallet)

	go func() {
		err = c.restAPIServer.Run(c.GetTestContext())
		// TODO: make sure this doesn't error rnormally
		assert.NoError(c.T(), err)
	}()
	// TODO: may cause races
	time.Sleep(time.Second)

}

func TestClientSuite(t *testing.T) {
	suite.Run(t, NewTestClientSuite(t))
}
