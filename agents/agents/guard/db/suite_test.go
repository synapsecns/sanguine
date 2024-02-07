package db_test

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm/schema"

	"github.com/Flaque/filet"
	"github.com/synapsecns/sanguine/agents/agents/guard/db"
	"github.com/synapsecns/sanguine/agents/agents/guard/db/sql/mysql"
	"github.com/synapsecns/sanguine/agents/agents/guard/db/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/agents/guard/metadata"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
)

type DBSuite struct {
	*testsuite.TestSuite
	dbs     []db.GuardDB
	metrics metrics.Handler
}

// NewEventDBSuite creates a new EventDBSuite.
func NewEventDBSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.GuardDB{},
	}
}

func (t *DBSuite) SetupSuite() {
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
	t.metrics, err = metrics.NewByType(t.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(t.T(), err)
}

func (t *DBSuite) SetupTest() {
	t.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""), t.metrics, false)
	Nil(t.T(), err)

	t.dbs = []db.GuardDB{sqliteStore}
	t.setupMysqlDB()
}

func (t *DBSuite) setupMysqlDB() {
	// skip if mysql test disabled, this really only needs to be run in ci

	// skip if mysql test disabled
	if os.Getenv(dbcommon.EnableMysqlTestVar) == "" {
		return
	}
	// sets up the conn string to the default database
	connString := dbcommon.GetTestConnString()
	// sets up the mysql db
	testDB, err := sql.Open("mysql", connString)
	Nil(t.T(), err)
	// close the db once the connection is done
	defer func() {
		Nil(t.T(), testDB.Close())
	}()

	// override the naming strategy to prevent tests from messing with each other.
	// todo this should be solved via a proper teardown process or transactions.
	mysql.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("test%d_%d_", t.GetTestID(), time.Now().Unix()),
	}

	mysql.MaxIdleConns = 10

	// create the sql store
	mysqlStore, err := mysql.NewMysqlStore(t.GetTestContext(), connString, t.metrics)
	Nil(t.T(), err)
	// add the db
	t.dbs = append(t.dbs, mysqlStore)
}

func (t *DBSuite) RunOnAllDBs(testFunc func(testDB db.GuardDB)) {
	t.T().Helper()

	wg := sync.WaitGroup{}
	for _, testDB := range t.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.GuardDB) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

// TestDBSuite tests the db suite.
func TestEventDBSuite(t *testing.T) {
	suite.Run(t, NewEventDBSuite(t))
}
