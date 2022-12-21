package executor_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	agentsTestutil "github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	scribesqlite "github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
	"go.uber.org/atomic"
	"math/big"
	"testing"
	"time"
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

	NotaryWallet wallet.Wallet
	GuardWallet  wallet.Wallet
	NotarySigner signer.Signer
	GuardSigner  signer.Signer
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
	e.manager = testutil.NewDeployManager(e.T())
	e.wallet, err = wallet.FromRandom()
	Nil(e.T(), err)
	e.signer = localsigner.NewSigner(e.wallet.PrivateKey())
	sqliteStore, err := sqlite.NewSqliteStore(e.GetTestContext(), filet.TmpDir(e.T(), ""))
	Nil(e.T(), err)
	e.testDB = sqliteStore

	e.chainID = gofakeit.Uint32()
	e.destination = e.chainID + 1
	e.simulatedChain = geth.NewEmbeddedBackendForChainID(e.GetTestContext(), e.T(), big.NewInt(int64(e.chainID)))
	e.simulatedClient, err = backfill.DialBackend(e.GetTestContext(), e.simulatedChain.RPCAddress())
	Nil(e.T(), err)

	e.simulatedChain.FundAccount(e.GetTestContext(), e.wallet.Address(), *big.NewInt(params.Ether))
	e.deployManager = agentsTestutil.NewDeployManager(e.T())
	e.originContract, e.originRef = e.deployManager.GetOriginHarness(e.GetTestContext(), e.simulatedChain)

}

func TestExecutorSuite(t *testing.T) {
	suite.Run(t, NewExecutorSuite(t))
}
