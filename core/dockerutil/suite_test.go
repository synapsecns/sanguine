package dockerutil_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// DockerSuite defines the basic test suite.
type DockerSuite struct {
	*testsuite.TestSuite
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *DockerSuite {
	tb.Helper()
	return &DockerSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
