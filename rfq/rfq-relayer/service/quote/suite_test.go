package quote_test

import (
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"sync/atomic"
	"testing"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/metadata"

	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
)

type QuoteSuite struct {
	*testsuite.TestSuite
	logIndex atomic.Int64 // For thread safety
	metrics  metrics.Handler
	signer   signer.Signer
}

// NewQuoteSuiteSuite creates a new QuoteSuite.
func NewQuoteSuiteSuite(tb testing.TB) *QuoteSuite {
	tb.Helper()
	return &QuoteSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupSuite sets up the rest of the test suite.
func (t *QuoteSuite) SetupSuite() {
	t.TestSuite.SetupSuite()
	t.logIndex.Store(0)

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(t.GetSuiteContext(), t.T())
		metricsHandler = metrics.Jaeger
	}
	var err error
	t.metrics, err = metrics.NewByType(t.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(t.T(), err)

	wllet, err := wallet.FromRandom()
	Nil(t.T(), err)
	t.signer = localsigner.NewSigner(wllet.PrivateKey())
}

// TestDBSuite tests the Quote suite.
func TestQuoteSuiteSuite(t *testing.T) {
	suite.Run(t, NewQuoteSuiteSuite(t))
}
