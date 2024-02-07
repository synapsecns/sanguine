package db_test

import (
	dbSQL "database/sql"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql/mysql"
	"github.com/synapsecns/sanguine/services/cctp-relayer/metadata"
	"gorm.io/gorm/schema"
)

type DBSuite struct {
	*testsuite.TestSuite
	dbs     []db.CCTPRelayerDB
	metrics metrics.Handler
}

// NewDBSuite creates a new DBSuite.
func NewDBSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.CCTPRelayerDB{},
	}
}
func (d *DBSuite) SetupSuite() {
	d.TestSuite.SetupSuite()

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(d.GetSuiteContext(), d.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	d.metrics, err = metrics.NewByType(d.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(d.T(), err)
}

func (d *DBSuite) SetupTest() {
	d.TestSuite.SetupTest()

	sqliteStore, err := sql.Connect(d.GetTestContext(), dbcommon.Sqlite, filet.TmpDir(d.T(), ""), d.metrics)
	Nil(d.T(), err)

	d.dbs = []db.CCTPRelayerDB{sqliteStore}
	d.setupMysqlDB()
}

func (d *DBSuite) setupMysqlDB() {
	if os.Getenv(dbcommon.EnableMysqlTestVar) != "true" {
		return
	}

	mysql.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("cctp_%d", d.GetTestID()),
	}

	// sets up the conn string to the default database
	connString := dbcommon.GetTestConnString()
	// sets up the myqsl db
	testDB, err := dbSQL.Open("mysql", connString)
	d.Require().NoError(err)
	// close the db once the connection is don
	defer func() {
		d.Require().NoError(testDB.Close())
	}()

	mysqlStore, err := mysql.NewMysqlStore(d.GetTestContext(), connString, d.metrics)
	d.Require().NoError(err)

	d.dbs = append(d.dbs, mysqlStore)
}

func (d *DBSuite) RunOnAllDBs(testFunc func(testDB db.CCTPRelayerDB)) {
	d.T().Helper()

	wg := sync.WaitGroup{}
	for _, testDB := range d.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.CCTPRelayerDB) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

func TestDBSuite(t *testing.T) {
	suite.Run(t, NewDBSuite(t))
}
