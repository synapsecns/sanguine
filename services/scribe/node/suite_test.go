package node_test

import (
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"testing"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
)

type LiveSuite struct {
	*testsuite.TestSuite
	testDB  db.EventDB
	manager *testutil.DeployManager
	wallet  wallet.Wallet
	signer  *localsigner.Signer
	metrics metrics.Handler
}

// NewLiveSuite creates a new live test suite.
func NewLiveSuite(tb testing.TB) *LiveSuite {
	tb.Helper()
	return &LiveSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (l *LiveSuite) SetupSuite() {
	l.TestSuite.SetupSuite()
	localmetrics.SetupTestJaeger(l.GetSuiteContext(), l.T())
	var err error

	l.metrics, err = metrics.NewByType(l.GetSuiteContext(), metadata.BuildInfo(), metrics.Jaeger)
	l.Require().Nil(err)
}

func (l *LiveSuite) SetupTest() {
	l.TestSuite.SetupTest()

	var err error
	l.testDB, err = sqlite.NewSqliteStore(l.GetTestContext(), filet.TmpDir(l.T(), ""), l.metrics, false)
	Nil(l.T(), err)

	l.manager = testutil.NewDeployManager(l.T())

	l.wallet, err = wallet.FromRandom()
	Nil(l.T(), err)
	l.signer = localsigner.NewSigner(l.wallet.PrivateKey())
}

// TestLiveSuite tests the live suite.
func TestLiveSuite(t *testing.T) {
	suite.Run(t, NewLiveSuite(t))
}
