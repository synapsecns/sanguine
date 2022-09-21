package testutil

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/testutil/clickhouse"
	"github.com/synapsecns/sanguine/services/scribe/api"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"go.uber.org/atomic"
	"net/http"
)

// NewTestEnvDB sets up the test env with a database.
func NewTestEnvDB(t *testsuite.TestSuite) (db db.ConsumerDB, eventDB scribedb.EventDB, gqlClient *client.Client, logIndex atomic.Int64, cleanup func(), testBackend backends.SimulatedTestBackend, deployManager *DeployManager) {
	dbPath := filet.TmpDir(t.T(), "")

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), dbPath)
	assert.Nil(t.T(), err)

	eventDB = sqliteStore

	logIndex.Store(0)

	freePort := freeport.GetPort()

	go func() {
		assert.Nil(t.T(), api.Start(t.GetSuiteContext(), api.Config{
			HTTPPort: uint16(freePort),
			Database: "sqlite",
			Path:     dbPath,
			GRPCPort: uint16(freeport.GetPort()),
		}))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", freePort)

	gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

	// var request *http.Request
	t.Eventually(func() bool {
		request, err := http.NewRequestWithContext(t.GetTestContext(), http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		assert.Nil(t.T(), err)
		res, err := gqlClient.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	})

	cleanup, port, err := clickhouse.NewClickhouseStore("explorer")
	if cleanup == nil || *port == 0 || err != nil {
		return
	}
	assert.Equal(t.T(), err, nil)
	dbURL := "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test"
	consumerDB, err := sql.OpenGormClickhouse(t.GetTestContext(), dbURL)
	assert.Nil(t.T(), err)
	db = consumerDB

	// maybe newSimulatedBackendWithChainID?
	testBackend = simulated.NewSimulatedBackend(t.GetTestContext(), t.T())
	deployManager = NewDeployManager(t.T())

	return db, eventDB, gqlClient, logIndex, cleanup, testBackend, deployManager
}
