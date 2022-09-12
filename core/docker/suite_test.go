package docker_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// DockerSuite defines the basic test suite.
type DockerSuite struct {
	*testsuite.TestSuite
}

func NewTestSuite(tb testing.TB) *DockerSuite {
	tb.Helper()
	return &DockerSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestDockerSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
