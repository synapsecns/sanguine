package integration_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Flaque/filet"
	"github.com/phayes/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/scribe/client"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/server"
	"go.uber.org/atomic"
)

// IntegrationSuite defines the basic test suite.
type IntegrationSuite struct {
	*testsuite.TestSuite
	db        db.EventDB
	dbPath    string
	gqlClient *client.Client
	logIndex  atomic.Int64
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *IntegrationSuite {
	tb.Helper()
	return &IntegrationSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (i *IntegrationSuite) SetupTest() {
	i.TestSuite.SetupTest()
	i.dbPath = filet.TmpDir(i.T(), "")

	sqliteStore, err := sqlite.NewSqliteStore(i.GetTestContext(), i.dbPath)
	Nil(i.T(), err)

	i.db = sqliteStore

	i.logIndex.Store(0)

	port := freeport.GetPort()

	go func() {
		Nil(i.T(), server.Start(uint16(port), "sqlite", i.dbPath))
	}()

	baseURL := fmt.Sprintf("http://127.0.0.1:%d", port)

	i.Eventually(func() bool {
		// TODO: use context here
		_, err := http.Get(fmt.Sprintf("%s%s", baseURL, server.GraphiqlEndpoint))
		return err == nil
	})

	i.gqlClient = client.NewClient(http.DefaultClient, fmt.Sprintf("%s%s", baseURL, server.GraphqlEndpoint))
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
