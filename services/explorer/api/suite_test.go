package api_test

import (
	"fmt"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/services/explorer/api"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil/clickhouse"
	"net/http"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/graphql/client"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server"
	"go.uber.org/atomic"
)

// APISuite defines the basic test suite.
type APISuite struct {
	*testsuite.TestSuite
	db        db.ConsumerDB
	gqlClient *client.Client
	// grpcClient *rest.APIClient
	logIndex atomic.Int64
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

func (g *APISuite) SetupTest() {
	g.TestSuite.SetupTest()

	g.logIndex.Store(0)

	httpport := freeport.GetPort()
	cleanup, port, err := clickhouse.NewClickhouseStore("explorer")
	NotNil(g.T(), cleanup)
	Nil(g.T(), err)
	if port == nil || err != nil {
		cleanup()
	}
	// NotN

	address := "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test"
	g.db, err = sql.OpenGormClickhouse(g.GetTestContext(), address)
	Nil(g.T(), err)

	go func() {
		Nil(g.T(), api.Start(g.GetSuiteContext(), api.Config{
			HTTPPort: uint16(httpport),
			Address:  address,
		}))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", httpport)

	g.gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

	// config := rest.NewConfiguration()
	//config.BasePath = baseURL
	//config.Host = baseURL
	//g.grpcClient = rest.NewAPIClient(config)

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
	//
	// g.Eventually(func() bool {
	//	res, realRes, err := g.grpcClient.ScribeServiceApi.ScribeServiceCheck(g.GetTestContext(), rest.V1HealthCheckRequest{
	//		Service: "any",
	//	})
	//	if err == nil {
	//		defer func() {
	//			_ = realRes.Body.Close()
	//		}()
	//
	//		return *res.Status == rest.SERVING_HealthCheckResponseServingStatus
	//	}
	//
	//	return false
	//})
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
