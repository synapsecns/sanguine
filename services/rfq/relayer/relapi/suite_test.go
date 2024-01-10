package relapi_test

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"

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
	"github.com/synapsecns/sanguine/ethergo/submitter"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/connect"
	"golang.org/x/sync/errgroup"
)

// RelayerServer suite is the relayer API server test suite.
type RelayerServerSuite struct {
	*testsuite.TestSuite
	omniRPCClient        omniClient.RPCClient
	omniRPCTestBackends  []backends.SimulatedTestBackend
	testBackends         map[uint64]backends.SimulatedTestBackend
	originChainID        uint32
	destChainID          uint32
	fastBridgeAddressMap *xsync.MapOf[uint64, common.Address]
	database             reldb.Service
	cfg                  relconfig.Config
	testWallet           wallet.Wallet
	handler              metrics.Handler
	RelayerAPIServer     *relapi.RelayerAPIServer
	port                 uint16
	wallet               wallet.Wallet
}

// NewRelayerServerSuite creates a end-to-end test suite.
func NewRelayerServerSuite(tb testing.TB) *RelayerServerSuite {
	tb.Helper()
	return &RelayerServerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (c *RelayerServerSuite) SetupTest() {
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

	c.originChainID = 1
	c.destChainID = 42161

	testConfig := relconfig.Config{
		Chains: map[int]relconfig.ChainConfig{
			int(c.originChainID): {
				Bridge: ethFastBridgeAddress.Hex(),
			},
			int(c.destChainID): {
				Bridge: arbFastBridgeAddress.Hex(),
			},
		},
		RelayerAPIConfig: relconfig.RelayerAPIConfig{
			Database: relconfig.DatabaseConfig{
				Type: "sqlite",
				DSN:  filet.TmpFile(c.T(), "", "").Name(),
			},
			OmniRPCURL: testOmnirpc,
			Port:       strconv.Itoa(port),
		},
	}
	c.cfg = testConfig

	c.wallet, err = wallet.FromRandom()
	c.Require().NoError(err)
	signer := localsigner.NewSigner(c.wallet.PrivateKey())
	submitterCfg := &submitterConfig.Config{}
	ts := submitter.NewTransactionSubmitter(c.handler, signer, omniRPCClient, c.database.SubmitterDB(), submitterCfg)

	server, err := relapi.NewRelayerAPI(c.GetTestContext(), c.cfg, c.handler, c.omniRPCClient, c.database, ts)
	c.Require().NoError(err)
	c.RelayerAPIServer = server
}

func (c *RelayerServerSuite) SetupSuite() {
	c.TestSuite.SetupSuite()

	// let's create 2 mock chains
	chainIDs := []uint64{1, 42161}

	c.testBackends = make(map[uint64]backends.SimulatedTestBackend)

	g, _ := errgroup.WithContext(c.GetSuiteContext())
	for _, chainID := range chainIDs {
		// Setup Anvil backend for the suite to have RPC support
		// anvilOpts := anvil.NewAnvilOptionBuilder()
		// anvilOpts.SetChainID(chainID)
		// anvilOpts.SetBlockTime(1 * time.Second)
		// backend := anvil.NewAnvilBackend(c.GetSuiteContext(), c.T(), anvilOpts)
		backend := geth.NewEmbeddedBackendForChainID(c.GetSuiteContext(), c.T(), new(big.Int).SetUint64(chainID))

		// add the backend to the list of backends
		c.testBackends[chainID] = backend
		c.omniRPCTestBackends = append(c.omniRPCTestBackends, backend)
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
	testDB, _ := connect.Connect(c.GetSuiteContext(), dbType, filet.TmpDir(c.T(), ""), metricsHandler)
	c.database = testDB
	// setup config
}

// TestConfigSuite runs the integration test suite.
func TestRelayerServerSuite(t *testing.T) {
	suite.Run(t, NewRelayerServerSuite(t))
}
