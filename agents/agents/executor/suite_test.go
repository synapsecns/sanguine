package executor_test

import (
	"testing"
	"time"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	agentsTestutil "github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	scribesqlite "github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"go.uber.org/atomic"
)

// ExecutorSuite tests the executor agent.
type ExecutorSuite struct {
	*agentsTestutil.SimulatedBackendsTestSuite
	scribeTestDB    scribedb.EventDB
	testDB          db.ExecutorDB
	dbPath          string
	logIndex        atomic.Int64
	manager         *testutil.DeployManager
	wallet          wallet.Wallet
	signer          *localsigner.Signer
	chainID         uint32
	destination     uint32
	simulatedChain  *geth.Backend
	simulatedClient backfill.ScribeBackend
	deployManager   *agentsTestutil.DeployManager
	originContract  contracts.DeployedContract
	originRef       *originharness.OriginHarnessRef
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
	sqliteStore, err := sqlite.NewSqliteStore(e.GetTestContext(), filet.TmpDir(e.T(), ""))
	Nil(e.T(), err)
	e.testDB = sqliteStore
}

func TestExecutorSuite(t *testing.T) {
	suite.Run(t, NewExecutorSuite(t))
}
