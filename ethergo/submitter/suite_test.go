package submitter_test

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/example"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/nonce"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
)

var buildInfo = config.NewBuildInfo(config.DefaultVersion, config.DefaultCommit, "submitter", config.DefaultDate)
var fundAmount = new(big.Int).Mul(new(big.Int).SetUint64(uint64(params.Ether)), big.NewInt(10))

// SubmitterSuite is used to test individual contract deployments to make sure other tests don't break.
type SubmitterSuite struct {
	*testsuite.TestSuite
	// testBackend is the test backend
	testBackends []backends.SimulatedTestBackend
	// metrics is the metrics handler
	metrics metrics.Handler
	// signer is the signer to use for the test
	signer signer.Signer
	// store is the store to use for the test
	store db.Service
	// registry is the registry to use for the test
	deployer *manager.DeployerManager
	// localAccount is the local account to use for the test
	localAccount *keystore.Key
}

// GetClient returns a client for the given chain id.
func (s *SubmitterSuite) GetClient(ctx context.Context, chainID *big.Int) (client.EVM, error) {
	for _, backend := range s.testBackends {
		if backend.GetBigChainID().Cmp(chainID) == 0 {
			//nolint: wrapcheck
			return client.DialBackend(ctx, backend.RPCAddress(), s.metrics)
		}
	}
	return nil, fmt.Errorf("could not find client for chain id %v", chainID)
}

// SetupSuite sets up 3 backends and metrics.
func (s *SubmitterSuite) SetupSuite() {
	s.TestSuite.SetupSuite()

	testChainIDs := []uint64{1, 3, 4}
	s.testBackends = make([]backends.SimulatedTestBackend, len(testChainIDs))
	s.deployer = manager.NewDeployerManager(s.T(), example.NewCounterDeployer)

	var wg sync.WaitGroup
	// wait for all the backends to be created, add 1 to the wait group for the metrics
	wg.Add(len(testChainIDs) + 1)

	logDir := filet.TmpDir(s.T(), "")

	// create the jaeger instance
	go func() {
		defer wg.Done()
		var err error
		// don't use metrics on ci for integration tests
		isCI := core.GetEnvBool("CI", false)
		useMetrics := !isCI
		metricsHandler := metrics.Null

		if useMetrics {
			localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
			metricsHandler = metrics.Jaeger
		}
		s.metrics, err = metrics.NewByType(s.GetSuiteContext(), buildInfo, metricsHandler)
		s.Require().NoError(err)
	}()

	// create the backends
	for i, chainID := range testChainIDs {
		go func(index int, chainID uint64) {
			defer wg.Done()
			options := anvil.NewAnvilOptionBuilder()
			options.SetChainID(chainID)
			// make sure all the docker containers log to the same directory
			options.SetProcessLogOptions(processlog.WithLogFileName(fmt.Sprintf("chain-%d.log", chainID)), processlog.WithLogDir(logDir))

			s.testBackends[index] = anvil.NewAnvilBackend(s.GetSuiteContext(), s.T(), options)
			s.deployer.Get(s.GetSuiteContext(), s.testBackends[index], example.CounterType)
		}(i, chainID)
	}
	wg.Wait()

	// fallback is currently untested. For now we disable it in tests.
	// TODO: this should be fixed, or ideally, fallback can be removed.
	og := submitter.ForceNoFallbackIfZero
	submitter.SetForceNoFallback(true)
	s.T().Cleanup(func() {
		submitter.SetForceNoFallback(og)
	})
}

// SetupTest sets up the signer and funds the account with 10 eth on each backend.
func (s *SubmitterSuite) SetupTest() {
	s.TestSuite.SetupTest()
	s.localAccount = mocks.MockAccount(s.T())
	// create the local signer
	s.signer = localsigner.NewSigner(s.localAccount.PrivateKey)
	var wg sync.WaitGroup
	wg.Add(len(s.testBackends) + 1)

	// setup the db
	go func() {
		defer wg.Done()
		var err error
		s.store, err = NewSqliteStore(s.GetTestContext(), filet.TmpDir(s.T(), ""), s.metrics)
		s.Require().NoError(err)
	}()

	// fund the account on each chain
	for i := range s.testBackends {
		go func(index int) {
			defer wg.Done()

			s.testBackends[index].FundAccount(s.GetTestContext(), s.signer.Address(), *fundAmount)
		}(i)
	}
	wg.Wait()
}

var _ submitter.ClientFetcher = &SubmitterSuite{}

// NewSubmitterSuite creates a end-to-end test suite.
func NewSubmitterSuite(tb testing.TB) *SubmitterSuite {
	tb.Helper()
	return &SubmitterSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestSubmitterSuite(t *testing.T) {
	suite.Run(t, NewSubmitterSuite(t))
}

// TXSubmitterDBSuite is used to test db queries across mysql and sqlite
// this is ran here rather than in the db package to avoid having to export the sqlite
// setup query to avoid confusion about how to use the library.
type TXSubmitterDBSuite struct {
	*testsuite.TestSuite
	dbs          []db.Service
	metrics      metrics.Handler
	testBackends []backends.SimulatedTestBackend
	managers     map[uint]nonce.Manager
	mockAccounts []*keystore.Key
}

func NewTXSubmitterDBSuite(tb testing.TB) *TXSubmitterDBSuite {
	tb.Helper()
	return &TXSubmitterDBSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.Service{},
	}
}

func TestTXSubmitterDBSuite(t *testing.T) {
	suite.Run(t, NewTXSubmitterDBSuite(t))
}

func (t *TXSubmitterDBSuite) SetupSuite() {
	t.TestSuite.SetupSuite()
	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(t.GetSuiteContext(), t.T())
		metricsHandler = metrics.Jaeger
	}
	var err error
	t.metrics, err = metrics.NewByType(t.GetSuiteContext(), buildInfo, metricsHandler)
	t.Require().NoError(err)

	og := submitter.ForceNoFallbackIfZero
	submitter.SetForceNoFallback(true)
	t.T().Cleanup(func() {
		submitter.SetForceNoFallback(og)
	})
}

func (t *TXSubmitterDBSuite) SetupTest() {
	t.TestSuite.SetupTest()

	t.dbs = []db.Service{}
	t.testBackends = []backends.SimulatedTestBackend{}

	t.setupMysqlDB()

	sqliteStore, err := NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""), t.metrics)
	t.Require().NoError(err)

	t.dbs = append(t.dbs, sqliteStore)

	// setup a signer for each chain
	t.managers = make(map[uint]nonce.Manager)

	// fund the accounts on each chain
	const mockAccounts = 5
	t.mockAccounts = make([]*keystore.Key, mockAccounts)
	for i := 0; i < mockAccounts; i++ {
		t.mockAccounts[i] = mocks.MockAccount(t.T())
	}

	chainIDs := []uint{1, 3, 4}
	var backendMux sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(chainIDs))

	for _, chainID := range chainIDs {
		go func(chainID uint) {
			defer wg.Done()

			backend := simulated.NewSimulatedBackendWithChainID(t.GetTestContext(), t.T(), big.NewInt(int64(chainID)))

			backendMux.Lock()
			t.testBackends = append(t.testBackends, backend)
			t.managers[chainID] = nonce.NewNonceManager(t.GetTestContext(), backend, backend.GetBigChainID())
			backendMux.Unlock()

			for _, account := range t.mockAccounts {
				backend.FundAccount(t.GetTestContext(), account.Address, *fundAmount)
			}
		}(chainID)
	}
	wg.Wait()
}

func (t *TXSubmitterDBSuite) RunOnAllDBs(testFunc func(db db.Service)) {
	wg := sync.WaitGroup{}
	for _, testDB := range t.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.Service) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

func (t *TXSubmitterDBSuite) setupMysqlDB() {
	// skip if mysql test disabled, this really only needs to be run in ci

	// skip if mysql test disabled
	if os.Getenv(common_base.EnableMysqlTestVar) == "" {
		return
	}
	// sets up the conn string to the default database
	connString := t.connString(os.Getenv(common_base.MysqlDatabaseVar))
	// sets up the myqsl db
	testDB, err := sql.Open("mysql", connString)
	t.Require().NoError(err)
	// close the db once the connection is don
	defer func() {
		t.Require().NoError(testDB.Close())
	}()

	mysqlStore, err := NewMysqlStore(t.GetTestContext(), connString, t.metrics)
	t.Require().NoError(err)

	t.dbs = append(t.dbs, mysqlStore)
}

// connString gets the mysql connection string.
func (t *TXSubmitterDBSuite) connString(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", core.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), core.GetEnv("MYSQL_HOST", "127.0.0.1"), core.GetEnvInt("MYSQL_PORT", 3306), dbname)
}

// Store wraps the store. Since tx submitter is a library and not a standalone service, we simulate db creation here.
// and then proceed as we would with any other db test.
type Store struct {
	*txdb.Store
}

// NewMysqlStore creates a new mysql data store. It emulates the way another caller would create a store.
func NewMysqlStore(parentCtx context.Context, dbURL string, handler metrics.Handler) (_ *Store, err error) {
	logger := log.Logger("mysql-store")

	ctx, span := handler.Tracer().Start(parentCtx, "start-mysql")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	namingStrategy := schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("test%d_%d_", gofakeit.Int64(), time.Now().Unix()),
	}

	gdb, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger:                 common_base.GetGormLogger(logger),
		FullSaveAssociations:   true,
		NamingStrategy:         namingStrategy,
		NowFunc:                time.Now,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create mysql connection: %w", err)
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, fmt.Errorf("could not get sql db: %w", err)
	}

	// fixes a timeout issue https://stackoverflow.com/a/42146536
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	handler.AddGormCallbacks(gdb)

	// migrate in a transaction since we skip this by default
	err = gdb.Transaction(func(tx *gorm.DB) error {
		//nolint: wrapcheck
		return gdb.WithContext(ctx).AutoMigrate(txdb.GetAllModels()...)
	})

	if err != nil {
		return nil, fmt.Errorf("could not migrate on mysql: %w", err)
	}
	return &Store{txdb.NewTXStore(gdb, handler)}, nil
}

// NewSqliteStore creates a new sqlite data store.
func NewSqliteStore(parentCtx context.Context, dbPath string, handler metrics.Handler) (_ *Store, err error) {
	logger := log.Logger("sqlite-store")

	logger.Debugf("creating sqlite store at %s", dbPath)

	ctx, span := handler.Tracer().Start(parentCtx, "start-sqlite")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// create the directory to the store if it doesn't exist
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not create sqlite store")
	}

	logger.Warnf("submitter database is at %s/synapse.db", dbPath)

	namingStrategy := schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("test%d_%d_", gofakeit.Int64(), time.Now().Unix()),
	}

	gdb, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/%s", dbPath, "synapse.db")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   common_base.GetGormLogger(logger),
		FullSaveAssociations:                     true,
		SkipDefaultTransaction:                   true,
		NamingStrategy:                           namingStrategy,
	})
	if err != nil {
		return nil, fmt.Errorf("could not connect to db %s: %w", dbPath, err)
	}

	handler.AddGormCallbacks(gdb)

	err = gdb.WithContext(ctx).AutoMigrate(txdb.GetAllModels()...)
	if err != nil {
		return nil, fmt.Errorf("could not migrate models: %w", err)
	}
	return &Store{txdb.NewTXStore(gdb, handler)}, nil
}
