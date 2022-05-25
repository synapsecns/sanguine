package db_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

// DBSuite is the db test suite.
type DBSuite struct {
	*testutils.TestSuite
}

// NewDBSuite creates a db test suite.
func NewDBSuite(tb testing.TB) *DBSuite {
	tb.Helper()
	return &DBSuite{TestSuite: testutils.NewTestSuite(tb)}
}

// TestDBSuite tests the db suite.
func TestDBSuite(t *testing.T) {
	suite.Run(t, NewDBSuite(t))
}
