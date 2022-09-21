package db_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"go.uber.org/atomic"
	"testing"
)

type DBSuite struct {
	*testsuite.TestSuite
	db            db.ConsumerDB
	eventDB       scribedb.EventDB
	gqlClient     *client.Client
	logIndex      atomic.Int64
	cleanup       func()
	testBackend   backends.SimulatedTestBackend
	deployManager *testutil.DeployManager
}

// NewDBSuite creates a new ConsumerDBSuite.
func NewDBSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (t *DBSuite) SetupTest() {
	t.TestSuite.SetupTest()

	t.db, t.eventDB, t.gqlClient, t.logIndex, t.cleanup, t.testBackend, t.deployManager = testutil.NewTestEnvDB(t.TestSuite)
}

// TestDBSuite tests the db suite.
func TestDBSuite(t *testing.T) {
	suite.Run(t, NewDBSuite(t))
}
