package backfill_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	eventdb "github.com/synapsecns/sanguine/services/scribe/db"
	"go.uber.org/atomic"
	"testing"
)

type BackfillSuite struct {
	*testsuite.TestSuite
	db        db.ConsumerDB
	eventDB   eventdb.EventDB
	gqlClient *client.Client
	logIndex  atomic.Int64
	cleanup   func()
}

// NewBackfillSuite creates a new backfill test suite.
func NewBackfillSuite(tb testing.TB) *BackfillSuite {
	tb.Helper()
	return &BackfillSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (t *BackfillSuite) SetupTest() {
	t.TestSuite.SetupTest()

	t.db, t.eventDB, t.gqlClient, t.logIndex, t.cleanup = testutil.SetupDB(t.TestSuite)
}

// TestBackfillSuite tests the backfill suite.
func TestBackfillSuite(t *testing.T) {
	suite.Run(t, NewBackfillSuite(t))
}
