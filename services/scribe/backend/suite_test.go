package backend_test

import (
	"testing"
	"time"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/services/scribe/metadata"

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

type BackendSuite struct {
	*testsuite.TestSuite
	testDB  db.EventDB
	manager *testutil.DeployManager
	wallet  wallet.Wallet
	signer  *localsigner.Signer
	metrics metrics.Handler
}

// NewBackendSuite creates a new backfill test suite.
func NewBackendSuite(tb testing.TB) *BackendSuite {
	tb.Helper()
	return &BackendSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupTest sets up the test suite.
func (b *BackendSuite) SetupTest() {
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

func (b *BackendSuite) SetupSuite() {
	b.TestSuite.SetupSuite()

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(b.GetSuiteContext(), b.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	b.metrics, err = metrics.NewByType(b.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(b.T(), err)
}

// TestBackendSuite tests the backfill suite.
func TestBackendSuite(t *testing.T) {
	suite.Run(t, NewBackendSuite(t))
}
