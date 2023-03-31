package api_test

import (
	"fmt"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/services/explorer/api"
	explorerclient "github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"github.com/synapsecns/sanguine/services/explorer/testutil/clickhouse"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	gqlServer "github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
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
	db     db.ConsumerDB
	client *client.Client
	// grpcClient *rest.APIClient
	eventDB       scribedb.EventDB
	gqlClient     *explorerclient.Client
	logIndex      atomic.Int64
	cleanup       func()
	testBackend   backends.SimulatedTestBackend
	deployManager *testutil.DeployManager
	chainIDs      []uint32
	scribeMetrics metrics.Handler
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
	metrics.SetupTestJaeger(g.GetSuiteContext(), g.T())

	var err error
	g.scribeMetrics, err = metrics.NewByType(g.GetSuiteContext(), metadata.BuildInfo(), metrics.Jaeger)
	g.Require().Nil(err)
}

func (g *APISuite) SetupTest() {
	g.TestSuite.SetupTest()

	g.db, g.eventDB, g.gqlClient, g.logIndex, g.cleanup, g.testBackend, g.deployManager = testutil.NewTestEnvDB(g.GetTestContext(), g.T(), g.scribeMetrics)

	httpport := freeport.GetPort()
	cleanup, port, err := clickhouse.NewClickhouseStore("explorer")
	NotNil(g.T(), cleanup)
	NotNil(g.T(), port)
	Nil(g.T(), err)
	if port == nil || err != nil {
		g.TearDownTest()
		return
	}

	address := "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test"
	g.db, err = sql.OpenGormClickhouse(g.GetTestContext(), address, false)
	Nil(g.T(), err)
	g.chainIDs = []uint32{1, 10, 25, 56, 137}
	go func() {
		Nil(g.T(), api.Start(g.GetSuiteContext(), api.Config{
			HTTPPort:  uint16(httpport),
			Address:   address,
			ScribeURL: g.gqlClient.Client.BaseURL,
		}))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", httpport)

	g.client = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, gqlServer.GraphqlEndpoint))

	g.Eventually(func() bool {
		request, err := http.NewRequestWithContext(g.GetTestContext(), http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		Nil(g.T(), err)
		res, err := g.client.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	})
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
