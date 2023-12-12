package service_test

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/config"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/testutil"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"golang.org/x/sync/errgroup"

	"github.com/Flaque/filet"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/sql/mysql"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/sql/sqlite"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/metadata"

	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"

	"gorm.io/gorm/schema"
)

type RelayerSuite struct {
	*testsuite.TestSuite
	dbs                 []db.TestDB
	logIndex            atomic.Int64 // For thread safety
	metrics             metrics.Handler
	testBackends        map[uint32]backends.SimulatedTestBackend
	contractHandlers    map[uint32]testutil.ITestContractHandler
	omniRpcTestBackends []backends.SimulatedTestBackend
	config              *config.Config
	chainIDs            []uint32
	wallet              wallet.Wallet
}

// NewRelayerSuiteSuite creates a new RelayerSuite.
func NewRelayerSuiteSuite(tb testing.TB) *RelayerSuite {
	tb.Helper()
	return &RelayerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.TestDB{},
	}
}

// SetupTest sets up the databases.
func (t *RelayerSuite) SetupTest() {
	t.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)

	t.dbs = []db.TestDB{sqliteStore}
	t.setupMysqlDB()
}

// SetupSuite sets up the rest of the test suite.
func (t *RelayerSuite) SetupSuite() {
	t.TestSuite.SetupSuite()
	t.logIndex.Store(0)

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(t.GetSuiteContext(), t.T())
		metricsHandler = metrics.Jaeger
	}
	var err error
	t.metrics, err = metrics.NewByType(t.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(t.T(), err)

	// Wallet
	t.wallet, _ = wallet.FromRandom()

	// Create test backends
	t.chainIDs = []uint32{1, 42161, 137, 56}
	t.testBackends = make(map[uint32]backends.SimulatedTestBackend)
	g, _ := errgroup.WithContext(t.GetSuiteContext())
	for _, chainID := range t.chainIDs {
		currChainID := chainID // capture func literal
		g.Go(func() error {
			backend := testutil.NewAnvilBackend(t.GetSuiteContext(), currChainID, t.T())
			t.testBackends[currChainID] = backend
			t.omniRpcTestBackends = append(t.omniRpcTestBackends, backend)
			return nil
		})
	}

	// wait for all backends to be ready
	err = g.Wait()
	Nil(t.T(), err)

	// Init all test contract handlers
	t.initAllContractHandlers()

	// Generate config
	t.generateConfig()

}

func (t *RelayerSuite) setupMysqlDB() {
	// If we are going to use mysql, 100% we should test using it on CI
	// This function will never run since we don't have the CI testing for mysql
	// set up in this repo, but we should set that up if we are going to use mysql for the relayer.
	if os.Getenv(dbcommon.EnableMysqlTestVar) != "true" {
		return
	}
	// Init connection
	connString := dbcommon.GetTestConnString()
	testDB, err := sql.Open("mysql", connString)
	Nil(t.T(), err)

	// Close the DB when the test is done
	defer func() {
		Nil(t.T(), testDB.Close())
	}()
	if err := testDB.Ping(); err != nil {
		fmt.Println("error connecting to MySQL database: %w", err)
	}

	// Override the naming strategy to prevent tests from messing with each other.
	mysql.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("test%d_%d_", t.GetTestID(), time.Now().Unix()),
	}

	mysql.MaxIdleConns = 10
	mysql.MaxOpenConns = 10

	// Create the sql store
	mysqlStore, err := mysql.NewMysqlStore(t.GetTestContext(), connString, t.metrics, false)
	fmt.Println("mysqlStore", mysqlStore, err)

	Nil(t.T(), err)
	// Add the db to the Relayer suite
	t.dbs = append(t.dbs, mysqlStore)
}

// RunOnAllDBs runs the test function on all dbs available.
func (t *RelayerSuite) RunOnAllDBs(testFunc func(testDB db.TestDB)) {
	t.T().Helper()
	wg := sync.WaitGroup{}
	for _, testDB := range t.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.TestDB) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

func (t *RelayerSuite) initAllContractHandlers() {
	contractHandlers := make(map[uint32]testutil.ITestContractHandler)
	for _, chainID := range t.chainIDs {
		originContractHandler, err := testutil.NewTestContractHandlerImpl(t.GetSuiteContext(), t.testBackends[chainID], t.wallet, chainID)
		Nil(t.T(), err)
		NotNil(t.T(), originContractHandler)
		contractHandlers[chainID] = originContractHandler
	}
	t.contractHandlers = contractHandlers
}

func (t *RelayerSuite) generateConfig() {
	chains := make(map[uint32]config.ChainConfig)
	var assets []config.AssetConfig
	for _, chainID := range t.chainIDs {
		chains[chainID] = config.ChainConfig{
			ChainID:                 chainID,
			FastBridgeAddress:       t.contractHandlers[chainID].FastBridgeAddress().String(),
			FastBridgeBlockDeployed: 0,
		}

		// Add assets
		chainTokens := t.contractHandlers[chainID].Tokens()
		for _, token := range chainTokens {
			assets = append(assets, config.AssetConfig{
				ChainID: chainID,
				Address: token.Erc20Address.String(),
			})
		}
	}
	t.config = &config.Config{
		Chains:         chains,
		Database:       config.DatabaseConfig{},
		Assets:         assets,
		RelayerAddress: t.wallet.Address().String(),
		Signer: signerConfig.SignerConfig{
			Type: "File",
			File: "./test-signer.txt",
		},
		SubmitterConfig: submitterConfig.Config{},
		OmnirpcURL:      "",
	}
	testOmnirpc := omnirpcHelper.NewOmnirpcServer(t.GetSuiteContext(), t.T(), t.omniRpcTestBackends...)
	t.config.OmnirpcURL = testOmnirpc
}

// TestDBSuite tests the Relayer suite.
func TestRelayerSuiteSuite(t *testing.T) {
	suite.Run(t, NewRelayerSuiteSuite(t))
}
