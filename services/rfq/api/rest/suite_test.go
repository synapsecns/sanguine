package rest_test

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
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/db/sql"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"golang.org/x/sync/errgroup"
)

// Server suite is the main API server test suite.
type ServerSuite struct {
	*testsuite.TestSuite
	omniRPCClient        omniClient.RPCClient
	omniRPCTestBackends  []backends.SimulatedTestBackend
	testBackends         map[uint64]backends.SimulatedTestBackend
	fastBridgeAddressMap *xsync.MapOf[uint64, common.Address]
	database             db.APIDB
	cfg                  config.Config
	testWallet           wallet.Wallet
	relayerWallets       []wallet.Wallet
	handler              metrics.Handler
	QuoterAPIServer      *rest.QuoterAPIServer
	port                 uint16
	originChainID        int
	destChainID          int
}

// NewServerSuite creates a end-to-end test suite.
func NewServerSuite(tb testing.TB) *ServerSuite {
	tb.Helper()
	return &ServerSuite{
		TestSuite:      testsuite.NewTestSuite(tb),
		relayerWallets: []wallet.Wallet{},
	}
}

//nolint:gosec
func (c *ServerSuite) SetupTest() {
	c.TestSuite.SetupTest()

	c.setDB()
	testOmnirpc := omnirpcHelper.NewOmnirpcServer(c.GetTestContext(), c.T(), c.omniRPCTestBackends...)
	omniRPCClient := omniClient.NewOmnirpcClient(testOmnirpc, c.handler, omniClient.WithCaptureReqRes())
	c.omniRPCClient = omniRPCClient

	c.originChainID = 1
	c.destChainID = 42161
	arbFastBridgeAddress, ok := c.fastBridgeAddressMap.Load(uint64(c.destChainID))
	c.True(ok)
	ethFastBridgeAddress, ok := c.fastBridgeAddressMap.Load(uint64(c.originChainID))
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
		Port:        fmt.Sprintf("%d", port),
		MaxQuoteAge: 15 * time.Minute,
	}
	c.cfg = testConfig

	QuoterAPIServer, err := rest.NewAPI(c.GetTestContext(), c.cfg, c.handler, c.omniRPCClient, c.database)
	c.Require().NoError(err)

	c.QuoterAPIServer = QuoterAPIServer

	// go func() {
	// 	err := c.QuoterAPIServer.Run(c.GetTestContext())
	// 	c.Require().NoError(err)
	// }()
	// time.Sleep(2 * time.Second) // Wait for the server to start.
}

func (c *ServerSuite) SetupSuite() {
	c.TestSuite.SetupSuite()

	// let's create 2 mock chains
	chainIDs := []uint64{1, 42161}

	c.testBackends = make(map[uint64]backends.SimulatedTestBackend)

	g, _ := errgroup.WithContext(c.GetSuiteContext())
	for _, chainID := range chainIDs {
		chainID := chainID // capture func literal
		mux := sync.Mutex{}
		g.Go(func() error {
			// Setup Anvil backend for the suite to have RPC support
			// anvilOpts := anvil.NewAnvilOptionBuilder()
			// anvilOpts.SetChainID(chainID)
			// anvilOpts.SetBlockTime(1 * time.Second)
			// backend := anvil.NewAnvilBackend(c.GetSuiteContext(), c.T(), anvilOpts)
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
	c.relayerWallets = []wallet.Wallet{c.testWallet}
	for range [2]int{} {
		relayerWallet, err := wallet.FromRandom()
		c.Require().NoError(err)
		c.relayerWallets = append(c.relayerWallets, relayerWallet)
	}
	for _, backend := range c.testBackends {
		for _, relayerWallet := range c.relayerWallets {
			backend.FundAccount(c.GetSuiteContext(), relayerWallet.Address(), *big.NewInt(params.Ether))
		}
	}

	c.fastBridgeAddressMap = xsync.NewIntegerMapOf[uint64, common.Address]()

	g, _ = errgroup.WithContext(c.GetSuiteContext())
	for _, backend := range c.testBackends {
		backend := backend
		// TODO: functionalize me
		g.Go(func() error {
			chainID, err := backend.ChainID(c.GetSuiteContext())
			if err != nil {
				return fmt.Errorf("could not get chain id: %w", err)
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
			relayerRole, err := fastBridgeInstance.RELAYERROLE(&bind.CallOpts{Context: c.GetTestContext()})
			c.NoError(err)

			// Grant relayer role to all relayer wallets
			for _, relayerWallet := range c.relayerWallets {
				tx, err = fastBridgeInstance.GrantRole(auth, relayerRole, relayerWallet.Address())
				c.Require().NoError(err)
				backend.WaitForConfirmation(c.GetSuiteContext(), tx)
			}

			return nil
		})
	}

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		c.T().Fatal(err)
	}
	// setup config
}

func (c *ServerSuite) setDB() {
	dbType, err := dbcommon.DBTypeFromString("sqlite")
	c.Require().NoError(err)
	metricsHandler := metrics.NewNullHandler()
	c.handler = metricsHandler
	// TODO use temp file / in memory sqlite3 to not create in directory files
	testDB, _ := sql.Connect(c.GetSuiteContext(), dbType, filet.TmpDir(c.T(), ""), metricsHandler)
	c.database = testDB
}

// TestConfigSuite runs the integration test suite.
func TestServerSuite(t *testing.T) {
	suite.Run(t, NewServerSuite(t))
}
