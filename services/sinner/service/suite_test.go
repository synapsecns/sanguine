package service_test

import (
	"context"
	"database/sql"
	"github.com/ethereum/go-ethereum/common"

	"fmt"
	"github.com/alecthomas/assert"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"k8s.io/apimachinery/pkg/util/wait"

	"math/big"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/scribe/api"
	scribeSqlite "github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/sql/mysql"
	"github.com/synapsecns/sanguine/services/sinner/db/sql/sqlite"
	"github.com/synapsecns/sanguine/services/sinner/fetcher"
	"github.com/synapsecns/sanguine/services/sinner/fetcher/client"
	gqlServer "github.com/synapsecns/sanguine/services/sinner/graphql/server"
	"github.com/synapsecns/sanguine/services/sinner/metadata"
	"gorm.io/gorm/schema"
)

// ServiceSuite is the test suite for the db package.
type ServiceSuite struct {
	*testsuite.TestSuite
	dbs                    []db.EventDB
	logIndex               atomic.Int64
	scribeDB               scribedb.EventDB
	metrics                metrics.Handler
	scribeFetcher          fetcher.ScribeFetcher
	testBackend            backends.SimulatedTestBackend
	destinationTestBackend backends.SimulatedTestBackend
	originChainID          uint32
	destinationChainID     uint32
	originTestLog          types.Log
	desTestLog             types.Log
}

// NewEventServiceSuite creates a new EventServiceSuite.
func NewEventServiceSuite(tb testing.TB) *ServiceSuite {
	tb.Helper()

	return &ServiceSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.EventDB{},
	}
}

func (t *ServiceSuite) SetupTest() {
	t.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)

	t.dbs = []db.EventDB{sqliteStore}
	t.setupMysqlDB()
}

func (t *ServiceSuite) SetupSuite() {
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
	t.Require().Nil(err)

	t.originChainID = 1
	t.destinationChainID = 2
	t.scribeDB, t.scribeFetcher = t.CreateScribeFetcher(t.GetSuiteContext())
	t.testBackend = simulated.NewSimulatedBackendWithChainID(t.GetSuiteContext(), t.T(), big.NewInt(int64(t.originChainID)))
	t.destinationTestBackend = simulated.NewSimulatedBackendWithChainID(t.GetSuiteContext(), t.T(), big.NewInt(int64(t.destinationChainID)))

	t.originChainID = 1
	t.destinationChainID = 2
	sentTopic := common.HexToHash("0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87")
	sentTopic2 := common.HexToHash("0xc6e19a3538fbd9b7a4f9bd8d45e08a95ff23e7e03b6a3bc9d9db9b8869b55c94")
	sentTopic3 := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")
	sentTopic4 := common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000002c")

	executeTopic := common.HexToHash("0x39c48fd1b2185b07007abc7904a8cdf782cfe449fd0e9bba1c2223a691e15f0b")
	executeTopic2 := common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000002b")
	executeTopic3 := common.HexToHash("0x481244fb9db711b88ab9bfe081311cbed0b50dc547a71151aef55a38871fc9bd")

	t.originTestLog = types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		BlockNumber: gofakeit.Uint64(),
		Topics:      []common.Hash{sentTopic, sentTopic2, sentTopic3, sentTopic4},
		Data:        []byte{},
		TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
		TxIndex:     uint(gofakeit.Int8()),
		BlockHash:   common.HexToHash(big.NewInt(gofakeit.Int64()).String()),
		Index:       uint(gofakeit.Int8()),
		Removed:     false,
	}
	t.desTestLog = types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		BlockNumber: gofakeit.Uint64(),
		Topics:      []common.Hash{executeTopic, executeTopic2, executeTopic3},
		Data:        []byte{},
		TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
		TxIndex:     uint(gofakeit.Int8()),
		BlockHash:   common.HexToHash(big.NewInt(gofakeit.Int64()).String()),
		Index:       uint(gofakeit.Int8()),
		Removed:     false,
	}
}

func (t *ServiceSuite) setupMysqlDB() {
	// skip if mysql test disabled, this really only needs to be run in ci
	// skip if mysql test disabled
	if os.Getenv(dbcommon.EnableMysqlTestVar) == "" {
		return
	}
	// sets up the conn string to the default database
	connString := dbcommon.GetTestConnString()
	// sets up the myqsl db
	testDB, err := sql.Open("mysql", connString)
	Nil(t.T(), err)
	// close the db once the ocnnection is odne
	defer func() {
		Nil(t.T(), testDB.Close())
	}()

	// override the naming strategy to prevent tests from messing with each other.
	// todo this should be solved via a proper teardown process or transactions.
	mysql.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("test%d_%d_", t.GetTestID(), time.Now().Unix()),
	}

	mysql.MaxIdleConns = 10
	mysql.MaxOpenConns = 10

	// create the sql store
	mysqlStore, err := mysql.NewMysqlStore(t.GetTestContext(), connString, t.metrics, false)
	Nil(t.T(), err)
	// add the db
	t.dbs = append(t.dbs, mysqlStore)
}

func (t *ServiceSuite) RunOnAllDBs(testFunc func(testDB db.EventDB)) {
	t.T().Helper()
	wg := sync.WaitGroup{}
	for _, testDB := range t.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.EventDB) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

// TestServiceSuite tests the db suite.
func TestEventServiceSuite(t *testing.T) {
	suite.Run(t, NewEventServiceSuite(t))
}

func (t *ServiceSuite) CreateScribeFetcher(ctx context.Context) (scribeDB scribedb.EventDB, scribeFetcher fetcher.ScribeFetcher) {
	t.T().Helper()
	dbPath := filet.TmpDir(t.T(), "")
	sqliteStore, err := scribeSqlite.NewSqliteStore(ctx, dbPath, t.metrics, false)
	Nil(t.T(), err)

	scribeDB = sqliteStore

	freePort := freeport.GetPort()

	go func() {
		err = api.Start(ctx, api.Config{
			Port:     uint16(freePort),
			Database: "sqlite",
			Path:     dbPath,
		}, t.metrics)
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", freePort)
	gqlClient := client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, gqlServer.GraphqlEndpoint))

	checkConnection := func() bool {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		assert.Nil(t.T(), err)
		res, err := gqlClient.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	}
	cancellableCtx, cancel := context.WithCancel(ctx)
	isTrue := false
	wait.UntilWithContext(cancellableCtx, func(cancellableCtx context.Context) {
		if checkConnection() {
			isTrue = true
			cancel()
		}
	}, time.Millisecond)

	// make sure the context didn't cancel
	if !isTrue {
		t.T().Errorf("expected %T to be true before test context timed out", checkConnection)
	}

	scribeFetcher = fetcher.NewFetcher(gqlClient, t.metrics)

	return scribeDB, scribeFetcher
}
