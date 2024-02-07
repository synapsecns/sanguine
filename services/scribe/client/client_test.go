package client_test

import (
	"fmt"
	"testing"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/scribe/client"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientSuite defines the basic test suite.
type ClientSuite struct {
	*testsuite.TestSuite
	db      string
	dbPath  string
	metrics metrics.Handler
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *ClientSuite {
	tb.Helper()
	return &ClientSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (g *ClientSuite) SetupTest() {
	g.TestSuite.SetupTest()
	g.dbPath = filet.TmpDir(g.T(), "")
	g.db = "sqlite"
}

func (g *ClientSuite) SetupSuite() {
	g.TestSuite.SetupSuite()

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(g.GetSuiteContext(), g.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	g.metrics, err = metrics.NewByType(g.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	g.Require().Nil(err)
}

func (g *ClientSuite) TestEmbeddedScribe() {
	embeddedClient := client.NewEmbeddedScribe(g.db, g.dbPath, g.metrics)

	go func() {
		Nil(g.T(), embeddedClient.Start(g.GetSuiteContext()))
	}()

	g.Eventually(func() bool {
		conn, err := grpc.DialContext(g.GetSuiteContext(), fmt.Sprintf("%s:%d", embeddedClient.URL, embeddedClient.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return false
		}

		grpcClient := pbscribe.NewScribeServiceClient(conn)

		healthCheck, err := grpcClient.Check(g.GetSuiteContext(), &pbscribe.HealthCheckRequest{})
		if err != nil {
			return false
		}

		return healthCheck.GetStatus() == pbscribe.HealthCheckResponse_SERVING
	})
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
