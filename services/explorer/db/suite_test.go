package db_test

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	newClickhouse "github.com/synapsecns/sanguine/agents/testutil/clickhouse"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/scribe/api"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server"
	"go.uber.org/atomic"
	"net/http"
	"testing"
)

type DBSuite struct {
	*testsuite.TestSuite
	db        db.ConsumerDB
	eventDB   scribedb.EventDB
	dbPath    string
	gqlClient *client.Client
	logIndex  atomic.Int64
	cleanup   func()
}

// NewConsumerDBSuite creates a new ConsumerDBSuite.
func NewConsumerDBSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (t *DBSuite) SetupTest() {
	t.TestSuite.SetupTest()
	t.dbPath = filet.TmpDir(t.T(), "")

	sqliteStore, err := sqlite.NewSqliteStore(t.GetTestContext(), t.dbPath)
	Nil(t.T(), err)

	t.eventDB = sqliteStore

	t.logIndex.Store(0)

	freePort := freeport.GetPort()

	go func() {
		Nil(t.T(), api.Start(t.GetSuiteContext(), api.Config{
			HTTPPort: uint16(freePort),
			Database: "sqlite",
			Path:     t.dbPath,
			GRPCPort: uint16(freeport.GetPort()),
		}))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", freePort)

	t.gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

	// var request *http.Request
	t.Eventually(func() bool {
		request, err := http.NewRequestWithContext(t.GetTestContext(), http.MethodGet, fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint), nil)
		Nil(t.T(), err)
		res, err := t.gqlClient.Client.Client.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}
		return false
	})

	cleanup, port, err := newClickhouse.NewClickhouseStore("explorer")
	if cleanup == nil || *port == 0 || err != nil {
		return
	}
	t.cleanup = cleanup
	Equal(t.T(), err, nil)
	dbUrl := "clickhouse://clickhouse_test:clickhouse_test@localhost:" + fmt.Sprintf("%d", *port) + "/clickhouse_test?read_timeout=10s&write_timeout=20s"
	consumerDB, err := sql.OpenGormClickhouse(t.GetTestContext(), dbUrl)
	Nil(t.T(), err)
	t.db = consumerDB
}

// TestConsumerDBSuite tests the db suite.
func TestConsumerDBSuite(t *testing.T) {
	suite.Run(t, NewConsumerDBSuite(t))
}
