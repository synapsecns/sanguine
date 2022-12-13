package executor_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	scribesqlite "github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"go.uber.org/atomic"
	"testing"
	"time"
)

// ExecutorSuite tests the executor agent.
type ExecutorSuite struct {
	*testsuite.TestSuite
	scribeTestDB scribedb.EventDB
	testDB       db.ExecutorDB
	dbPath       string
	logIndex     atomic.Int64
	manager      *testutil.DeployManager
	wallet       wallet.Wallet
	signer       *localsigner.Signer
}

// NewExecutorSuite creates a new executor suite.
func NewExecutorSuite(tb testing.TB) *ExecutorSuite {
	tb.Helper()

	return &ExecutorSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (e *ExecutorSuite) SetupTest() {
	e.TestSuite.SetupTest()
	e.SetTestTimeout(time.Minute * 3)
	e.dbPath = filet.TmpDir(e.T(), "")
	scribeSqliteStore, err := scribesqlite.NewSqliteStore(e.GetTestContext(), e.dbPath)
	Nil(e.T(), err)
	e.scribeTestDB = scribeSqliteStore
	e.logIndex.Store(0)
	e.manager = testutil.NewDeployManager(e.T())
	e.wallet, err = wallet.FromRandom()
	Nil(e.T(), err)
	e.signer = localsigner.NewSigner(e.wallet.PrivateKey())
	sqliteStore, err := sqlite.NewSqliteStore(e.GetTestContext(), filet.TmpDir(e.T(), ""))
	Nil(e.T(), err)
	e.testDB = sqliteStore
}

func TestExecutorSuite(t *testing.T) {
	suite.Run(t, NewExecutorSuite(t))
}
