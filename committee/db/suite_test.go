package db_test

import (
	dbSQL "database/sql"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/ipfs/go-datastore"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/committee/db"
	"github.com/synapsecns/sanguine/committee/db/connect"
	"github.com/synapsecns/sanguine/committee/db/mysql"
	"github.com/synapsecns/sanguine/committee/metadata"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"gorm.io/gorm/schema"
	"os"
	"sync"
	"testing"
)

type DBSuite struct {
	*testsuite.TestSuite
	dbs     map[dbcommon.DBType]db.Service
	dss     map[dbcommon.DBType]datastore.Batching
	metrics metrics.Handler
}

// NewDBSuite creates a new DBSuite.
func NewDBSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       make(map[dbcommon.DBType]db.Service),
		dss:       make(map[dbcommon.DBType]datastore.Batching),
	}
}

func (d *DBSuite) SetupSuite() {
	d.TestSuite.SetupSuite()

	var err error
	d.metrics, err = metrics.NewByType(d.GetSuiteContext(), metadata.BuildInfo(), metrics.Null)
	d.NoError(err)
}

func (d *DBSuite) SetupTest() {
	d.TestSuite.SetupTest()

	sqliteStore, err := connect.Connect(d.GetTestContext(), dbcommon.Sqlite, filet.TmpDir(d.T(), ""), d.metrics)
	d.NoError(err)

	d.dbs[dbcommon.Sqlite] = sqliteStore
	d.setupMysqlDB()

	// make datastores
	for name, testDB := range d.dbs {
		ds, err := testDB.GlobalDatastore()
		d.NoError(err)

		d.dss[name] = ds
	}
}

func (d *DBSuite) setupMysqlDB() {
	if os.Getenv(dbcommon.EnableMysqlTestVar) != "true" {
		return
	}

	mysql.SetNamingStrategy(schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("committee_%d", d.GetTestID()),
	})

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

	d.dbs[dbcommon.Mysql] = mysqlStore
}

func (d *DBSuite) RunOnAllDBs(testFunc func(testDB db.Service)) {
	runOnAll[db.Service](d, d.dbs, testFunc)
}

func (d *DBSuite) RunOnAllDatastores(testFunc func(testStore datastore.Batching)) {
	runOnAll[datastore.Batching](d, d.dss, testFunc)
}

func runOnAll[T any](d *DBSuite, testMap map[dbcommon.DBType]T, testFunc func(testStore T)) {
	d.T().Helper()
	// note: d.T().Parallel() can't be called here, see: https://github.com/stretchr/testify/issues/187

	wg := sync.WaitGroup{}
	for name, testDB := range testMap {
		wg.Add(1)
		// capture the value
		go func(dbType dbcommon.DBType, testDB T) {
			defer wg.Done()
			d.T().Run(dbType.String(), func(t *testing.T) {
				testFunc(testDB)
			})
		}(name, testDB)
	}
	wg.Wait()
}

func TestDBSuite(t *testing.T) {
	suite.Run(t, NewDBSuite(t))
}
