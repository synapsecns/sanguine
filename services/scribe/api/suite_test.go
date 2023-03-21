package api_test

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/scribe/api"
	"github.com/synapsecns/sanguine/services/scribe/grpc/client/rest"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"net/http"
	"testing"

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
	db         db.EventDB
	dbPath     string
	gqlClient  *client.Client
	grpcClient *rest.APIClient
	logIndex   atomic.Int64
	metrics    metrics.Handler
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

	metrics.SetupTestJaeger(g.T())

	var err error
	g.metrics, err = metrics.NewByType(g.GetSuiteContext(), metadata.BuildInfo(), metrics.Jaeger)
	g.Require().Nil(err)
}

func (g *APISuite) SetupTest() {
	g.TestSuite.SetupTest()
	g.dbPath = filet.TmpDir(g.T(), "")

	sqliteStore, err := sqlite.NewSqliteStore(g.GetTestContext(), g.dbPath)
	Nil(g.T(), err)

	g.db = sqliteStore

	g.logIndex.Store(0)

	port := freeport.GetPort()

	go func() {
		Nil(g.T(), api.Start(g.GetSuiteContext(), api.Config{
			Port:       uint16(port),
			Database:   "sqlite",
			Path:       g.dbPath,
			OmniRPCURL: "https://rpc.interoperability.institute/confirmations/1/rpc",
		}, g.metrics))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", port)

	g.gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

	config := rest.NewConfiguration()
	config.BasePath = baseURL
	config.Host = baseURL
	g.grpcClient = rest.NewAPIClient(config)

	// var request *http.Request
	g.Eventually(func() bool {
		request, err := http.NewRequestWithContext(g.GetTestContext(), http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		Nil(g.T(), err)
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
		res, realRes, err := g.grpcClient.ScribeServiceApi.ScribeServiceCheck(g.GetTestContext(), rest.V1HealthCheckRequest{
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
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
