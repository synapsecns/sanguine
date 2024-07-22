package relapi_test

import (
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/testutil"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	deployManager        *testutil.DeployManager
	omniRPCClient        omniClient.RPCClient
	omniRPCTestBackends  []backends.SimulatedTestBackend
	testBackendMux       sync.Mutex
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

var testWithdrawalAddress = common.BigToAddress(big.NewInt(1))

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
				RFQAddress: ethFastBridgeAddress.Hex(),
			},
			int(c.destChainID): {
				RFQAddress: arbFastBridgeAddress.Hex(),
			},
		},
		RelayerAPIPort: strconv.Itoa(port),
		Database: relconfig.DatabaseConfig{
			Type: "sqlite",
			DSN:  filet.TmpFile(c.T(), "", "").Name(),
		},
		EnableAPIWithdrawals: true,
		WithdrawalWhitelist: []string{
			testWithdrawalAddress.String(),
		},
		QuotableTokens: map[string][]string{
			// gas tokens.
			fmt.Sprintf("%d-%s", c.originChainID, chain.EthAddress): {
				// not used for this test
			},
			fmt.Sprintf("%d-%s", c.destChainID, chain.EthAddress): {
				// not used for this test
			},
			c.getMockTokenID(c.testBackends[uint64(c.originChainID)]): {
				// not used for this test
			},
			c.getMockTokenID(c.testBackends[uint64(c.destChainID)]): {
				// not used for this test
			},
		},
	}
	c.cfg = testConfig

	c.wallet, err = wallet.FromRandom()
	c.Require().NoError(err)
	signer := localsigner.NewSigner(c.wallet.PrivateKey())
	submitterCfg := &submitterConfig.Config{}
	ts := submitter.NewTransactionSubmitter(c.handler, signer, omniRPCClient, c.database.SubmitterDB(), submitterCfg)

	var wg sync.WaitGroup
	wg.Add(len(c.testBackends) * 2)
	go func() {
		// small potential for a race condition if submitter hasn't started by the time our test starts
		_ = ts.Start(c.GetTestContext())
	}()

	for _, backend := range c.testBackends {
		go func() {
			defer wg.Done()
			backend.FundAccount(c.GetTestContext(), c.wallet.Address(), *big.NewInt(params.Ether))
		}()

		go func() {
			defer wg.Done()
			mockMetadata, mockToken := c.deployManager.GetMockERC20(c.GetTestContext(), backend)
			auth := backend.GetTxContext(c.GetTestContext(), mockMetadata.OwnerPtr()).TransactOpts

			tx, err := mockToken.Mint(auth, c.wallet.Address(), big.NewInt(1000000000000000000))
			c.Require().NoError(err)
			backend.WaitForConfirmation(c.GetTestContext(), tx)
		}()
	}

	server, err := relapi.NewRelayerAPI(c.GetTestContext(), c.cfg, c.handler, c.omniRPCClient, c.database, ts)
	c.Require().NoError(err)
	c.RelayerAPIServer = server
	wg.Wait()
}

func (c *RelayerServerSuite) getMockTokenID(backend backends.SimulatedTestBackend) string {
	erc20Metadata, _ := c.deployManager.GetMockERC20(c.GetTestContext(), backend)
	return fmt.Sprintf("%d-%s", backend.GetChainID(), erc20Metadata.Address().Hex())
}

func (c *RelayerServerSuite) SetupSuite() {
	c.TestSuite.SetupSuite()

	c.deployManager = testutil.NewDeployManager(c.T())

	// let's create 2 mock chains
	chainIDs := []uint64{1, 42161}

	c.testBackends = make(map[uint64]backends.SimulatedTestBackend)

	testWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	c.testWallet = testWallet

	g, _ := errgroup.WithContext(c.GetSuiteContext())
	for _, chainID := range chainIDs {
		// Setup Anvil backend for the suite to have RPC support
		chainID := chainID
		g.Go(func() error {
			backend := geth.NewEmbeddedBackendForChainID(c.GetSuiteContext(), c.T(), new(big.Int).SetUint64(chainID))

			backend.FundAccount(c.GetSuiteContext(), c.testWallet.Address(), *big.NewInt(params.Ether))

			// add the backend to the list of backends
			c.testBackendMux.Lock()
			defer c.testBackendMux.Unlock()
			c.testBackends[chainID] = backend
			c.omniRPCTestBackends = append(c.omniRPCTestBackends, backend)

			return nil
		})
	}

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		c.T().Fatal(err)
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

			fastBridgeMetadata, _ := c.deployManager.GetFastBridge(c.GetSuiteContext(), backend)

			// Save the contracts to the map
			c.fastBridgeAddressMap.Store(chainID.Uint64(), fastBridgeMetadata.Address())

			fastBridgeInstance, err := fastbridge.NewFastBridge(fastBridgeMetadata.Address(), backend)
			c.Require().NoError(err)

			relayerRole, err := fastBridgeInstance.RELAYERROLE(&bind.CallOpts{Context: c.GetSuiteContext()})
			c.NoError(err)

			auth := backend.GetTxContext(c.GetSuiteContext(), fastBridgeMetadata.OwnerPtr()).TransactOpts

			tx, err := fastBridgeInstance.GrantRole(auth, relayerRole, c.testWallet.Address())
			c.Require().NoError(err)
			backend.WaitForConfirmation(c.GetSuiteContext(), tx)

			return nil
		})

		g.Go(func() error {
			mockERC20Metadata, mockERC20 := c.deployManager.GetMockERC20(c.GetSuiteContext(), backend)
			if err != nil {
				return fmt.Errorf("could not get mock ERC20: %w", err)
			}

			auth := backend.GetTxContext(c.GetSuiteContext(), mockERC20Metadata.OwnerPtr()).TransactOpts

			mintTx, err := mockERC20.Mint(auth, c.testWallet.Address(), big.NewInt(1000000000000000000))
			c.Require().NoError(err)

			backend.WaitForConfirmation(c.GetSuiteContext(), mintTx)
			return nil
		})
	}

	dbType, err := dbcommon.DBTypeFromString("sqlite")
	c.Require().NoError(err)
	metricsHandler := metrics.NewNullHandler()
	c.handler = metricsHandler
	// TODO use temp file / in memory sqlite3 to not create in directory files
	testDB, _ := connect.Connect(c.GetSuiteContext(), dbType, filet.TmpDir(c.T(), ""), metricsHandler)
	c.database = testDB
	// setup config

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		c.T().Fatal(err)
	}
}

// TestConfigSuite runs the integration test suite.
func TestRelayerServerSuite(t *testing.T) {
	suite.Run(t, NewRelayerServerSuite(t))
}

type RelayerClientSuite struct {
	*testsuite.TestSuite
	underlying *RelayerServerSuite
	Client     relapi.RelayerClient
}

// NewRelayerClientSuite creates a new relayer client suite.
func NewRelayerClientSuite(tb testing.TB) *RelayerClientSuite {
	tb.Helper()
	underlying := NewRelayerServerSuite(tb)

	return &RelayerClientSuite{
		TestSuite:  underlying.TestSuite,
		underlying: underlying,
	}
}
func (c *RelayerClientSuite) SetupSuite() {
	c.underlying.SetupSuite()
}

func (c *RelayerClientSuite) SetupTest() {
	c.underlying.SetupTest()
	c.underlying.startQuoterAPIServer()
	c.Client = relapi.NewRelayerClient(c.underlying.handler, fmt.Sprintf("http://localhost:%d", c.underlying.port))
}

// TestConfigSuite runs the integration test suite.
func TestRelayerClientSuite(t *testing.T) {
	suite.Run(t, NewRelayerClientSuite(t))
}
