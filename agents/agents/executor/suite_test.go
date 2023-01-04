package executor_test

import (
	"testing"
	"time"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	agentsTestutil "github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	scribesqlite "github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"go.uber.org/atomic"
)

// ExecutorSuite tests the executor agent.
type ExecutorSuite struct {
	*agentsTestutil.SimulatedBackendsTestSuite
	scribeTestDB scribedb.EventDB
	testDB       db.ExecutorDB
	dbPath       string
	logIndex     atomic.Int64
	wallet       wallet.Wallet
	signer       *localsigner.Signer
}

// NewExecutorSuite creates a new executor suite.
func NewExecutorSuite(tb testing.TB) *ExecutorSuite {
	tb.Helper()

	return &ExecutorSuite{
		logIndex:                   atomic.Int64{},
		SimulatedBackendsTestSuite: agentsTestutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (e *ExecutorSuite) SetupTest() {
	e.SimulatedBackendsTestSuite.SetupTest()
	e.SetTestTimeout(time.Minute * 3)
	e.dbPath = filet.TmpDir(e.T(), "")
	scribeSqliteStore, err := scribesqlite.NewSqliteStore(e.GetTestContext(), e.dbPath)
	Nil(e.T(), err)
	e.scribeTestDB = scribeSqliteStore
	e.logIndex.Store(0)
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
