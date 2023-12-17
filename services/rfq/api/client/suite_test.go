package client_test

import (
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/phayes/freeport"
	"github.com/puzpuzpuz/xsync/v2"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/db/sql"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"golang.org/x/sync/errgroup"
)

type ClientSuite struct {
	*testsuite.TestSuite
	omniRPCClient        omniClient.RPCClient
	omniRPCTestBackends  []backends.SimulatedTestBackend
	testBackends         map[uint64]backends.SimulatedTestBackend
	fastBridgeAddressMap *xsync.MapOf[uint64, common.Address]
	database             db.APIDB
	cfg                  config.Config
	testWallet           wallet.Wallet
	handler              metrics.Handler
	APIServer            *rest.APIServer
	port                 uint16
	client               client.AuthenticatedClient
}

func NewTestClientSuite(tb testing.TB) *ClientSuite {
	tb.Helper()

	return &ClientSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (c *ClientSuite) SetupTest() {
	c.TestSuite.SetupTest()

	testOmnirpc := omnirpcHelper.NewOmnirpcServer(c.GetTestContext(), c.T(), c.omniRPCTestBackends...)
	omniRPCClient := omniClient.NewOmnirpcClient(testOmnirpc, c.handler, omniClient.WithCaptureReqRes())
	c.omniRPCClient = omniRPCClient

	arbFastBridgeAddress, ok := c.fastBridgeAddressMap.Load(42161)
	c.True(ok)
	ethFastBridgeAddress, ok := c.fastBridgeAddressMap.Load(1)
	c.True(ok)
	port, err := freeport.GetFreePort()
	c.port = uint16(port)
	c.Require().NoError(err)

	testConfig := config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  filet.TmpFile(c.T(), "", "").Name(),
		},
		OmniRPCURL: testOmnirpc,
		Bridges: map[uint32]string{
			1:     ethFastBridgeAddress.Hex(),
			42161: arbFastBridgeAddress.Hex(),
		},
		Port: fmt.Sprintf("%d", port),
	}
	c.cfg = testConfig

	APIServer, err := rest.NewAPI(c.GetTestContext(), c.cfg, c.handler, c.omniRPCClient, c.database)
	c.Require().NoError(err)
	c.APIServer = APIServer

	go func() {
		err := c.APIServer.Run(c.GetTestContext())
		c.Require().NoError(err)
	}()
	time.Sleep(2 * time.Second) // Wait for the server to start.

	c.client, err = client.NewAuthenticatedClient(metrics.Get(), fmt.Sprintf("http://127.0.0.1:%d", port), localsigner.NewSigner(c.testWallet.PrivateKey()))
	c.Require().NoError(err)
}

func (c *ClientSuite) SetupSuite() {
	c.TestSuite.SetupSuite()

	// let's create 2 mock chains
	chainIDs := []uint64{1, 42161}

	c.testBackends = make(map[uint64]backends.SimulatedTestBackend)
	var mux sync.Mutex

	g, _ := errgroup.WithContext(c.GetSuiteContext())
	for _, chainID := range chainIDs {
		chainID := chainID // capture func literal
		g.Go(func() error {
			// Setup backend for the suite to have RPC support
			backend := geth.NewEmbeddedBackendForChainID(c.GetSuiteContext(), c.T(), new(big.Int).SetUint64(chainID))

			// add the backend to the list of backends
			mux.Lock()
			defer mux.Unlock()
			c.testBackends[chainID] = backend
			c.omniRPCTestBackends = append(c.omniRPCTestBackends, backend)
			return nil
		})
	}

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		c.T().Fatal(err)
	}

	testWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	c.testWallet = testWallet
	for _, backend := range c.testBackends {
		backend.FundAccount(c.GetSuiteContext(), c.testWallet.Address(), *big.NewInt(params.Ether))
	}

	c.fastBridgeAddressMap = xsync.NewIntegerMapOf[uint64, common.Address]()

	g, _ = errgroup.WithContext(c.GetSuiteContext())
	for _, backend := range c.testBackends {
		backend := backend
		g.Go(func() error {
			chainID, err := backend.ChainID(c.GetSuiteContext())
			if err != nil {
				return fmt.Errorf("could not get chainID: %w", err)
			}
			// Create an auth to interact with the blockchain
			auth, err := bind.NewKeyedTransactorWithChainID(c.testWallet.PrivateKey(), chainID)
			c.Require().NoError(err)

			// Deploy the FastBridge contract
			fastBridgeAddress, tx, _, err := fastbridge.DeployFastBridge(auth, backend, c.testWallet.Address())
			c.Require().NoError(err)
			backend.WaitForConfirmation(c.GetSuiteContext(), tx)

			// Save the contracts to the map
			c.fastBridgeAddressMap.Store(chainID.Uint64(), fastBridgeAddress)

			fastBridgeInstance, err := fastbridge.NewFastBridge(fastBridgeAddress, backend)
			c.Require().NoError(err)
			tx, err = fastBridgeInstance.AddRelayer(auth, c.testWallet.Address())
			c.Require().NoError(err)
			backend.WaitForConfirmation(c.GetSuiteContext(), tx)

			return nil
		})
	}

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		c.T().Fatal(err)
	}

	dbType, err := dbcommon.DBTypeFromString("sqlite")
	c.Require().NoError(err)
	metricsHandler := metrics.NewNullHandler()
	c.handler = metricsHandler
	// TODO use temp file / in memory sqlite3 to not create in directory files
	testDB, err := sql.Connect(c.GetSuiteContext(), dbType, filet.TmpDir(c.T(), ""), metricsHandler)
	c.Require().NoError(err)
	c.database = testDB
	// setup config
}

// TestConfigSuite runs the integration test suite.
func TestClientSuite(t *testing.T) {
	suite.Run(t, NewTestClientSuite(t))
}
