package api_test

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"net/http"

	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/sinner/api"
	serverConfig "github.com/synapsecns/sanguine/services/sinner/config/server"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/graphql/client"
	"github.com/synapsecns/sanguine/services/sinner/graphql/server"

	"sync/atomic"
	"testing"

	"github.com/synapsecns/sanguine/services/sinner/db/sql/sqlite"

	"github.com/synapsecns/sanguine/services/sinner/metadata"
)

type APISuite struct {
	*testsuite.TestSuite
	db                 db.TestEventDB
	dbPath             string
	logIndex           atomic.Int64
	metrics            metrics.Handler
	sinnerAPI          *client.Client
	originChainID      uint32
	destinationChainID uint32
}

// NewEventAPISuite creates a new EventAPISuite.
func NewEventAPISuite(tb testing.TB) *APISuite {
	tb.Helper()

	return &APISuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (t *APISuite) SetupSuite() {
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
	t.dbPath = filet.TmpDir(t.T(), "")

	sqliteStore, err := sqlite.NewSqliteStore(t.GetSuiteContext(), t.dbPath, t.metrics, false)
	Nil(t.T(), err)

	t.db = sqliteStore
	t.originChainID = 1
	t.destinationChainID = 2

	httpPort := freeport.GetPort()
	hostName := fmt.Sprintf("localhost:%d", httpPort)
	baseURL := fmt.Sprintf("http://%s", hostName)
	config := serverConfig.Config{
		HTTPPort:       uint16(httpPort),
		DBPath:         t.dbPath,
		DBType:         dbcommon.Sqlite.String(),
		SkipMigrations: true,
	}

	go func() {
		Nil(t.T(), api.Start(t.GetSuiteContext(), config, t.metrics))
	}()
	t.sinnerAPI = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))
}

// TestAPISuite tests the db suite.
func TestEventAPISuite(t *testing.T) {
	suite.Run(t, NewEventAPISuite(t))
}
