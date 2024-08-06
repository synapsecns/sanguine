package relconfig_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/metadata"
)

func TestValidateDecimalsSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}

type ValidateDecimalsSuite struct {
	*testsuite.TestSuite
	// testBackends contains a list of all test backends
	metricsHandler metrics.Handler
	omniClient     omniClient.RPCClient
}

// NewTestSuite creates a new test suite.
func NewTestSuite(tb testing.TB) *ValidateDecimalsSuite {
	tb.Helper()
	return &ValidateDecimalsSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (v *ValidateDecimalsSuite) SetupSuite() {
	v.TestSuite.SetupSuite()

	var err error
	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	metricsHandler := metrics.Null

	if !isCI {
		localmetrics.SetupTestJaeger(v.GetSuiteContext(), v.T())
		metricsHandler = metrics.Jaeger
	}
	v.metricsHandler, err = metrics.NewByType(v.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	v.Require().NoError(err)

	v.omniClient = omniClient.NewOmnirpcClient("https://rpc.omnirpc.io", v.metricsHandler, omniClient.WithCaptureReqRes())
}
