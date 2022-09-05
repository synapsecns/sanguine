package mapmutex_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// MapMutexSuite defines the basic test suite.
type MapMutexSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *MapMutexSuite {
	tb.Helper()
	return &MapMutexSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestMapMutexSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
