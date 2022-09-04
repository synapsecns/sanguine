package integration_test

import (
	"testing"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"go.uber.org/atomic"
)

// IntegrationSuite defines the basic test suite.
type IntegrationSuite struct {
	*testsuite.TestSuite
	db       db.EventDB
	dbPath   string
	logIndex atomic.Int64
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *IntegrationSuite {
	tb.Helper()
	return &IntegrationSuite{
		testsuite.NewTestSuite(tb),
		nil,
		"",
		atomic.Int64{},
	}
}

func (i *IntegrationSuite) SetupTest() {
	i.TestSuite.SetupTest()
	i.dbPath = filet.TmpDir(i.T(), "")

	sqliteStore, err := sqlite.NewSqliteStore(i.GetTestContext(), i.dbPath)
	Nil(i.T(), err)

	i.db = sqliteStore

	i.logIndex.Store(0)
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
