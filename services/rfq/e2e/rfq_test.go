package e2e_test

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	apiConfig "github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db/sql"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
	"github.com/synapsecns/sanguine/services/rfq/relayer"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"
)

type IntegrationSuite struct {
	*testsuite.TestSuite
	manager       *testutil.DeployManager
	originBackend backends.SimulatedTestBackend
	destBackend   backends.SimulatedTestBackend
	//omniserver is the omnirpc server address
	omniServer string
	omniClient omnirpcClient.RPCClient
	metrics    metrics.Handler
	apiServer  string
	relayer    *relayer.Relayer
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

	rsigner, err := wallet.FromRandom()
	i.NoError(err)

	dsn := filet.TmpDir(i.T(), "")

	cfg := relconfig.Config{
		// generated ex-post facto
		Tokens: map[int][]string{},
		Bridges: map[int]relconfig.ChainConfig{
			originBackendChainID: {
				Bridge:        i.manager.Get(i.GetTestContext(), i.originBackend, testutil.FastBridgeType).Address().String(),
				Confirmations: 0,
			},
			destBackendChainID: {
				Bridge:        i.manager.Get(i.GetTestContext(), i.destBackend, testutil.FastBridgeType).Address().String(),
				Confirmations: 0,
			},
		},
		OmnirpcURL: i.omniServer,
		// TODO: need to stop hardcoding
		DBConfig: dsn,
		// generated ex-post facto
		QuotableTokens: map[string][]string{},
		RelayerAddress: rsigner.Address(),
	}

	// in the first backend, we want to deploy a bunch of different tokens
	// TODO: functionalize me.
	for _, backend := range core.ToSlice(i.originBackend, i.destBackend) {
		tokenTypes := []contracts.ContractType{testutil.DAIType, testutil.USDTType, testutil.USDCType, testutil.WETH9Type}

		for _, tokenType := range tokenTypes {
			tokenAddress := i.manager.Get(i.GetTestContext(), backend, tokenType).Address().String()
			quotableTokenID := fmt.Sprintf("%d-%s", backend.GetChainID(), tokenAddress)

			// first the simple part, add the token to the token map
			cfg.Tokens[int(backend.GetChainID())] = append(cfg.Tokens[int(backend.GetChainID())], tokenAddress)

			compatibleTokens := []contracts.ContractType{tokenType}
			// DAI/USDT are fungible
			if tokenType == testutil.DAIType || tokenType == testutil.USDCType {
				compatibleTokens = []contracts.ContractType{testutil.DAIType, testutil.USDCType}
			}

			// now we need to add the token to the quotable tokens map
			for _, token := range compatibleTokens {
				otherBackend := i.getOtherBackend(backend)
				otherToken := i.manager.Get(i.GetTestContext(), otherBackend, token).Address().String()

				cfg.QuotableTokens[quotableTokenID] = append(cfg.QuotableTokens[quotableTokenID], fmt.Sprintf("%d-%s", otherBackend.GetChainID(), otherToken))
			}
		}

	}

	// TODO: good chance we wanna leave actually starting this up to the indiividual test.
	i.relayer, err = relayer.NewRelayer(i.GetTestContext(), i.metrics, cfg)
	i.NoError(err)
	go func() {
		err = i.relayer.Start(i.GetTestContext())
	}()

	time.Sleep(time.Second * 5)
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

func (i *IntegrationSuite) setupAPI() {
	dbPath := filet.TmpDir(i.T(), "")
	apiPort, err := freeport.GetFreePort()
	i.NoError(err)

	apiStore, err := sql.Connect(i.GetTestContext(), dbcommon.Sqlite, dbPath, i.metrics)
	i.NoError(err)

	// make the api without bridges
	apiCfg := apiConfig.Config{
		Database: apiConfig.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  dbPath,
		},
		OmniRPCURL: i.omniServer,
		Bridges: map[uint32]string{
			originBackendChainID: i.manager.Get(i.GetTestContext(), i.originBackend, testutil.FastBridgeType).Address().String(),
			destBackendChainID:   i.manager.Get(i.GetTestContext(), i.destBackend, testutil.FastBridgeType).Address().String(),
		},
		Port: strconv.Itoa(apiPort),
	}
	api, err := rest.NewAPI(i.GetTestContext(), apiCfg, i.metrics, i.omniClient, apiStore)
	i.NoError(err)

	i.apiServer = fmt.Sprintf("http://localhost:%d", apiPort)

	go func() {
		err = api.Run(i.GetTestContext())
		i.NoError(err)
	}()

	// make sure api server hast started
	testsuite.Eventually(i.GetTestContext(), i.T(), func() bool {
		var req *http.Request
		req, err = http.NewRequestWithContext(i.GetTestContext(), http.MethodGet, i.apiServer, nil)
		i.NoError(err)

		//nolint: bodyclose
		_, err = http.DefaultClient.Do(req)
		if err == nil {
			return true
		}
		return false
	})
}

// setupBackends sets up the ether backends and the omnirpc client/server
func (i *IntegrationSuite) setupBackends() {
	var wg sync.WaitGroup

	// prdeploys are contracts we want to deploy before running the test to speed it up. Obviously, these can be deployed when we need them as well,
	// but this way we can do something while we're waiting for the other backend to startup.
	predeploys := []contracts.ContractType{testutil.FastBridgeType, testutil.DAIType, testutil.USDTType, testutil.USDCType, testutil.WETH9Type}

	wg.Add(2)
	go func() {
		defer wg.Done()
		i.originBackend = geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(originBackendChainID))
		i.manager.BulkDeploy(i.GetTestContext(), core.ToSlice(i.originBackend), predeploys...)

	}()
	go func() {
		defer wg.Done()
		i.destBackend = geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(destBackendChainID))
		i.manager.BulkDeploy(i.GetTestContext(), core.ToSlice(i.destBackend), predeploys...)
	}()
	wg.Wait()

	i.omniServer = testhelper.NewOmnirpcServer(i.GetTestContext(), i.T(), i.originBackend, i.destBackend)
	i.omniClient = omnirpcClient.NewOmnirpcClient(i.omniServer, i.metrics, omnirpcClient.WithCaptureReqRes())
}

// TODO:
func (i *IntegrationSuite) TestUSDCtoUSDC() {

}
