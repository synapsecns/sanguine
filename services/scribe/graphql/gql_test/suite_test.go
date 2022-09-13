package gql_test

import (
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/api"
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

// GQLSuite defines the basic test suite.
type GQLSuite struct {
	*testsuite.TestSuite
	db        db.EventDB
	dbPath    string
	gqlClient *client.Client
	logIndex  atomic.Int64
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *GQLSuite {
	tb.Helper()
	return &GQLSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (g *GQLSuite) SetupTest() {
	g.TestSuite.SetupTest()
	g.dbPath = filet.TmpDir(g.T(), "")

	sqliteStore, err := sqlite.NewSqliteStore(g.GetTestContext(), g.dbPath)
	Nil(g.T(), err)

	g.db = sqliteStore

	g.logIndex.Store(0)

	port := freeport.GetPort()

	go func() {
		Nil(g.T(), api.Start(g.GetSuiteContext(), uint16(port), "sqlite", g.dbPath))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", port)

	g.gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))

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

	// go func() {
	// 	<-g.GetSuiteContext().Done()
	// 	err = request.Body.Close()
	// 	Nil(g.T(), err)
	// }()
}

func TestGQLSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
