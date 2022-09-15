package db_test

import (
	"fmt"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	newClickhouse "github.com/synapsecns/sanguine/agents/testutil/clickhouse"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"testing"
)

type DBSuite struct {
	*testsuite.TestSuite
	db      db.ConsumerDB
	cleanup func()
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
