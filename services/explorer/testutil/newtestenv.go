package testutil

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe/client"
	"net/http"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	gqlServer "github.com/synapsecns/sanguine/services/explorer/graphql/server"
	"github.com/synapsecns/sanguine/services/explorer/testutil/clickhouse"
	"github.com/synapsecns/sanguine/services/scribe/api"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"go.uber.org/atomic"
	"k8s.io/apimachinery/pkg/util/wait"
)

// NewTestEnvDB sets up the test env with a database.
func NewTestEnvDB(ctx context.Context, t *testing.T, handler metrics.Handler) (db db.ConsumerDB, eventDB scribedb.EventDB, gqlClient *client.Client, logIndex atomic.Int64, cleanup func(), testBackend backends.SimulatedTestBackend, deployManager *DeployManager) {
	t.Helper()
	dbPath := filet.TmpDir(t, "")

	sqliteStore, err := sqlite.NewSqliteStore(ctx, dbPath, handler, false)
	assert.Nil(t, err)

	eventDB = sqliteStore

	logIndex.Store(0)

	freePort := freeport.GetPort()

	go func() {
		assert.Nil(t, api.Start(ctx, api.Config{
			Port:     uint16(freePort),
			Database: "sqlite",
			Path:     dbPath,
		}, handler))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", freePort)

	gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, gqlServer.GraphqlEndpoint))

	checkConnection := func() bool {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		assert.Nil(t, err)
		res, err := gqlClient.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	}
	cancellableCtx, cancel := context.WithCancel(ctx)
	isTrue := false
	wait.UntilWithContext(cancellableCtx, func(cancellableCtx context.Context) {
		if checkConnection() {
			isTrue = true
			cancel()
		}
	}, time.Millisecond)

	// make sure the context didn't cancel
	if !isTrue {
		t.Errorf("expected %T to be true before test context timed out", checkConnection)
	}

	cleanup, port, err := clickhouse.NewClickhouseStore("explorer")
	if cleanup == nil {
		panic("Clickhouse spin up failure, no open port found.")
	}
	if port == nil || err != nil {
		cleanup()
		panic("Clickhouse spin up failure, destroying container...")
	}
	assert.Equal(t, err, nil)
	dbURL := "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test"
	consumerDB, err := sql.OpenGormClickhouse(ctx, dbURL, false, handler)
	assert.Nil(t, err)
	db = consumerDB

	testBackend = geth.NewEmbeddedBackend(ctx, t)
	deployManager = NewDeployManager(t)

	return db, eventDB, gqlClient, logIndex, cleanup, testBackend, deployManager
}
