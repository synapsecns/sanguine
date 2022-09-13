package test_test

import (
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"

	"github.com/stretchr/testify/suite"
)

// SQLSuite is a test suite used for testing sql
// note: in general, the interface methods should be tested together in the
// db package. this is made for testing sql/gorm implementation specific
// cases.
type SQLSuite struct {
	*testsuite.TestSuite
}

// NewDBSuite creates a new chain testing suite.
func NewSQLSuite(tb testing.TB) *SQLSuite {
	tb.Helper()
	return &SQLSuite{TestSuite: testsuite.NewTestSuite(tb)}
}

func TestSqlSuite(t *testing.T) {
	suite.Run(t, NewSQLSuite(t))
}
