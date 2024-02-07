package service_test

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

type ScribeSuite struct {
	*testsuite.TestSuite
	testDB        db.EventDB
	manager       *testutil.DeployManager
	wallet        wallet.Wallet
	signer        *localsigner.Signer
	metrics       metrics.Handler
	nullMetrics   metrics.Handler
	runVolumeTest bool
}

// NewScribeSuite creates a new scribe test suite.
func NewScribeSuite(tb testing.TB) *ScribeSuite {
	tb.Helper()
	return &ScribeSuite{
		TestSuite:     testsuite.NewTestSuite(tb),
		runVolumeTest: true,
	}
}

// SetupTest sets up the test suite.
func (s *ScribeSuite) SetupTest() {
	s.TestSuite.SetupTest()
	s.SetTestTimeout(time.Minute * 10)
	sqliteStore, err := sqlite.NewSqliteStore(s.GetTestContext(), filet.TmpDir(s.T(), ""), s.metrics, false)
	Nil(s.T(), err)
	s.testDB = sqliteStore
	s.manager = testutil.NewDeployManager(s.T())
	s.wallet, err = wallet.FromRandom()
	Nil(s.T(), err)
	s.signer = localsigner.NewSigner(s.wallet.PrivateKey())
}

func (s *ScribeSuite) SetupSuite() {
	s.TestSuite.SetupSuite()

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	s.metrics, err = metrics.NewByType(s.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(s.T(), err)

	s.nullMetrics, err = metrics.NewByType(s.GetSuiteContext(), metadata.BuildInfo(), metrics.Null)
	Nil(s.T(), err)
}

// TestScribeSuite tests the scribe suite.
func TestScribeSuite(t *testing.T) {
	suite.Run(t, NewScribeSuite(t))
}
