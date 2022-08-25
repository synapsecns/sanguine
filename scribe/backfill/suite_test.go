package backfill_test

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/scribe/db"
	"github.com/synapsecns/sanguine/scribe/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/synapse-node/pkg/common"
	"gorm.io/gorm/schema"
)

type BackfillSuite struct {
	*testsuite.TestSuite
	dbs     []db.EventDB
	manager *testutil.DeployManager
	wallet  wallet.Wallet
	signer  *localsigner.Signer
}

// NewBackfillSuite creates a new backfill test suite.
func NewBackfillSuite(tb testing.TB) *BackfillSuite {
	tb.Helper()
	return &BackfillSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		dbs:       []db.EventDB{},
	}
}

func (b *BackfillSuite) SetupTest() {
	b.TestSuite.SetupTest()

	sqliteStore, err := sqlite.NewSqliteStore(b.GetTestContext(), filet.TmpDir(b.T(), ""))
	Nil(b.T(), err)

	b.dbs = []db.EventDB{sqliteStore}
	b.setupMysqlDB()

	b.manager = testutil.NewDeployManager(b.T())

	b.wallet, err = wallet.FromRandom()
	Nil(b.T(), err)
	b.signer = localsigner.NewSigner(b.wallet.PrivateKey())
}

// connString gets the connection string.
func (b *BackfillSuite) connString(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", common.GetEnv("MYSQL_USER", "root"), os.Getenv("MYSQL_PASSWORD"), common.GetEnv("MYSQL_HOST", "127.0.0.1"), common.GetEnvInt("MYSQL_PORT", 3306), dbname)
}

func (b *BackfillSuite) setupMysqlDB() {
	// skip if mysql test disabled, this really only needs to be run in ci

	// skip if mysql test disabled
	if os.Getenv("ENABLE_MYSQL_TEST") == "" {
		return
	}
	// sets up the conn string to the default database
	connString := b.connString(os.Getenv("MYSQL_DATABASE"))
	// sets up the myqsl db
	testDB, err := sql.Open("mysql", connString)
	Nil(b.T(), err)
	// close the db once the ocnnection is odne
	defer func() {
		Nil(b.T(), testDB.Close())
	}()

	// override the naming strategy to prevent tests from messing with each other.
	// todo this should be solved via a proper teardown process or transactions.
	mysql.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("test%d_%d_", b.GetTestID(), time.Now().Unix()),
	}

	mysql.MaxIdleConns = 10

	// create the sql store
	mysqlStore, err := mysql.NewMysqlStore(b.GetTestContext(), connString)
	Nil(b.T(), err)
	// add the db
	b.dbs = append(b.dbs, mysqlStore)
}

func (b *BackfillSuite) RunOnAllDBs(testFunc func(testDB db.EventDB)) {
	b.T().Helper()

	wg := sync.WaitGroup{}
	for _, testDB := range b.dbs {
		wg.Add(1)
		// capture the value
		go func(testDB db.EventDB) {
			defer wg.Done()
			testFunc(testDB)
		}(testDB)
	}
	wg.Wait()
}

// TestBackfillSuite tests the backfill suite.
func TestBackfillSuite(t *testing.T) {
	suite.Run(t, NewBackfillSuite(t))
}
