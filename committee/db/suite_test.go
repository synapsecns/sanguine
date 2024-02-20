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
	dbs     []db.Service
	metrics metrics.Handler
}

// NewDBSuite creates a new DBSuite.
func NewDBSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.Service{},
	}
}

func (d *DBSuite) TestDB() {
	d.RunOnAllDBs(func(testDB db.Service) {
		ds, err := testDB.GlobalDatastore()
		d.NoError(err)

		err = ds.Put(d.GetTestContext(), datastore.NewKey("key"), []byte("value"))
		d.NoError(err)

		value, err := ds.Get(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)
		d.Equal([]byte("value"), value)

		var has bool
		has, err = ds.Has(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)
		d.True(has)

		_, err = ds.GetSize(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)

		err = ds.Delete(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)

		has, err = ds.Has(d.GetTestContext(), datastore.NewKey("key"))
		d.NoError(err)
		d.False(has)

	})
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

	d.dbs = []db.Service{sqliteStore}
	d.setupMysqlDB()
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

	d.dbs = append(d.dbs, mysqlStore)
}

func (d *DBSuite) RunOnAllDBs(testFunc func(testDB db.Service)) {
	d.T().Helper()

	wg := sync.WaitGroup{}
	for _, testDB := range d.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.Service) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

func TestDBSuite(t *testing.T) {
	suite.Run(t, NewDBSuite(t))
}
