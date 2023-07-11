package backfill_test

import (
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"testing"
	"time"

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

type BackfillSuite struct {
	*testsuite.TestSuite
	testDB  db.EventDB
	manager *testutil.DeployManager
	wallet  wallet.Wallet
	signer  *localsigner.Signer
	metrics metrics.Handler
}

// NewBackfillSuite creates a new backfill test suite.
func NewBackfillSuite(tb testing.TB) *BackfillSuite {
	tb.Helper()
	return &BackfillSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupTest sets up the test suite.
func (b *BackfillSuite) SetupTest() {
	b.TestSuite.SetupTest()
	b.SetTestTimeout(time.Minute * 3)
	sqliteStore, err := sqlite.NewSqliteStore(b.GetTestContext(), filet.TmpDir(b.T(), ""), b.metrics, false)
	Nil(b.T(), err)
	b.testDB = sqliteStore
	b.manager = testutil.NewDeployManager(b.T())
	b.wallet, err = wallet.FromRandom()
	Nil(b.T(), err)
	b.signer = localsigner.NewSigner(b.wallet.PrivateKey())
}

func (b *BackfillSuite) SetupSuite() {
	b.TestSuite.SetupSuite()
	localmetrics.SetupTestJaeger(b.GetSuiteContext(), b.T())

	var err error
	b.metrics, err = metrics.NewByType(b.GetSuiteContext(), metadata.BuildInfo(), metrics.Jaeger)
	Nil(b.T(), err)
}

// TestBackfillSuite tests the backfill suite.
func TestBackfillSuite(t *testing.T) {
	suite.Run(t, NewBackfillSuite(t))
}
