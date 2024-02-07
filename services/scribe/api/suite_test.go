package api_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/grpc/client/rest"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/client"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"go.uber.org/atomic"
)

// APISuite defines the basic test suite.
type APISuite struct {
	*testsuite.TestSuite
	db             db.EventDB
	dbPath         string
	gqlClient      *client.Client
	grpcRestClient *rest.APIClient
	grpcClient     pbscribe.ScribeServiceClient
	logIndex       atomic.Int64
	metrics        metrics.Handler
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *APISuite {
	tb.Helper()
	return &APISuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (g *APISuite) SetupSuite() {
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

func (g *APISuite) TearDownSuite() {
	g.TestSuite.TearDownSuite()
	time.Sleep(time.Second * 10)
}

func (g *APISuite) SetupTest() {
	g.TestSuite.SetupTest()
	g.dbPath = filet.TmpDir(g.T(), "")
	g.SetTestTimeout(time.Minute * 3)

	sqliteStore, err := sqlite.NewSqliteStore(g.GetTestContext(), g.dbPath, g.metrics, false)
	Nil(g.T(), err)

	g.db = sqliteStore

	g.logIndex.Store(0)

	port := freeport.GetPort()

	go func() {
		Nil(g.T(), api.Start(g.GetSuiteContext(), api.Config{
			Port:           uint16(port),
			Database:       "sqlite",
			Path:           g.dbPath,
			OmniRPCURL:     "https://rpc.omnirpc.io/confirmations/1/rpc",
			SkipMigrations: true,
		}, g.metrics))
	}()

	hostName := fmt.Sprintf("localhost:%d", port)
	baseURL := fmt.Sprintf("http://%s", hostName)

	g.gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

	config := rest.NewConfiguration()
	config.BasePath = baseURL
	config.Host = hostName

	g.grpcRestClient = rest.NewAPIClient(config)
	rawGrpcClient, err := grpc.DialContext(g.GetTestContext(), hostName, grpc.WithTransportCredentials(insecure.NewCredentials()))
	g.NoError(err)

	g.grpcClient = pbscribe.NewScribeServiceClient(rawGrpcClient)

	// var request *http.Request
	g.Eventually(func() bool {
		request, err := http.NewRequestWithContext(g.GetTestContext(), http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		g.NoError(err)
		res, err := g.gqlClient.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	})

	g.Eventually(func() bool {
		res, realRes, err := g.grpcRestClient.ScribeServiceApi.ScribeServiceCheck(g.GetTestContext(), rest.V1HealthCheckRequest{
			Service: "any",
		})

		if err == nil {
			defer func() {
				_ = realRes.Body.Close()
			}()

			return *res.Status == rest.SERVING_HealthCheckResponseServingStatus
		}

		return false
	})

	g.Eventually(func() bool {
		res, err := g.grpcClient.Check(g.GetTestContext(), &pbscribe.HealthCheckRequest{
			Service: "any",
		})
		if err == nil {
			return res.Status == pbscribe.HealthCheckResponse_SERVING
		}

		return false
	})
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
