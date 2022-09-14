package db_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"testing"
)

type DBSuite struct {
	*testsuite.TestSuite
	db db.ConsumerDB
}

// NewConsumerDBSuite creates a new ConsumerDBSuite.
func NewConsumerDBSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (t *DBSuite) SetupTest() {
	t.TestSuite.SetupTest()

	consumerDB, err := sql.NewClickhouseStore(t.GetTestContext(), "localhost:9000/tcp")
	Nil(t.T(), err)

	t.db = consumerDB
}

// TestConsumerDBSuite tests the db suite.
func TestConsumerDBSuite(t *testing.T) {
	suite.Run(t, NewConsumerDBSuite(t))
}
