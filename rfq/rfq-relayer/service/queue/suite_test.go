package queue_test

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

type QueueSuite struct {
	*testsuite.TestSuite
	dbs      []db.TestDB
	logIndex atomic.Int64 // For thread safety
	metrics  metrics.Handler
}

// NewQueueSuiteSuite creates a new QueueSuite.
func NewQueueSuiteSuite(tb testing.TB) *QueueSuite {
	tb.Helper()
	return &QueueSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.TestDB{},
	}
}

// SetupTest sets up the databases.
func (t *QueueSuite) SetupTest() {
	t.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)

	t.dbs = []db.TestDB{sqliteStore}
	t.setupMysqlDB()
}

// SetupSuite sets up the rest of the test suite.
func (t *QueueSuite) SetupSuite() {
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
}

func (t *QueueSuite) setupMysqlDB() {
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
	// Add the db to the queue suite
	t.dbs = append(t.dbs, mysqlStore)
}

// RunOnAllDBs runs the test function on all dbs available.
func (t *QueueSuite) RunOnAllDBs(testFunc func(testDB db.TestDB)) {
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

// TestDBSuite tests the queue suite.
func TestQueueSuiteSuite(t *testing.T) {
	suite.Run(t, NewQueueSuiteSuite(t))
}
