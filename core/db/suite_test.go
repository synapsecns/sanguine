package db_test

import (
	"database/sql"
	"fmt"
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/core/db/datastore/sql/sqlite"
	"github.com/synapsecns/synapse-node/pkg/common"
	"github.com/synapsecns/synapse-node/testutils"
	"gorm.io/gorm/schema"
	"os"
	"sync"
	"testing"
	"time"
)

type DBSuite struct {
	*testutils.TestSuite
	dbs []db.SynapseDB
}

// NewTxQueueSuite creates a new transaction queue suite.
func NewTxQueueSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{
		TestSuite: testutils.NewTestSuite(tb),
		dbs:       []db.SynapseDB{},
	}
}

func (t *DBSuite) SetupTest() {
	t.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""))
	Nil(t.T(), err)

	t.dbs = []db.SynapseDB{sqliteStore}
	t.setupMysqlDB()
}

// connString gets the connection string.
func (t *DBSuite) connString(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", common.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), common.GetEnv("MYSQL_HOST", "127.0.0.1"), common.GetEnvInt("MYSQL_PORT", 3306), dbname)
}

func (t *DBSuite) setupMysqlDB() {
	// skip if mysql test disabled, this really only needs to be run in ci

	// skip if mysql test disabled
	if os.Getenv("ENABLE_MYSQL_TEST") == "" {
		return
	}
	// sets up the conn string to the default database
	connString := t.connString(os.Getenv("MYSQL_DATABASE"))
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

	// create the sql store
	mysqlStore, err := mysql.NewMysqlStore(t.GetTestContext(), connString)
	Nil(t.T(), err)
	// add the db
	t.dbs = append(t.dbs, mysqlStore)
}

func (t *DBSuite) RunOnAllDBs(testFunc func(testDB db.SynapseDB)) {
	t.T().Helper()

	wg := sync.WaitGroup{}
	for _, testDB := range t.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.SynapseDB) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

// TestDBSuite tests the db suite.
func TestTxQueueSuite(t *testing.T) {
	suite.Run(t, NewTxQueueSuite(t))
}
