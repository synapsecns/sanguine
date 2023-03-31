package localmetrics_test

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// LocalServerSuite defines the basic test suite.
type LocalServerSuite struct {
	*testsuite.TestSuite
	client *docker.Client
}

func NewLocalServerSuite(tb testing.TB) *LocalServerSuite {
	tb.Helper()
	return &LocalServerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (l *LocalServerSuite) SetupSuite() {
	l.TestSuite.SetupSuite()
	pool, err := dockertest.NewPool("")
	l.Require().NoError(err)
	l.client = pool.Client
}

func TestCommonSuite(t *testing.T) {
	suite.Run(t, NewLocalServerSuite(t))
}

func (l *LocalServerSuite) hasContainerWithRunID(runID string) bool {
	containers := l.ContainersWithLabel(localmetrics.RunIDLabel, runID)
	return len(containers) > 0
}

func (l *LocalServerSuite) ContainersWithLabel(key, value string) []docker.APIContainers {
	containers, err := l.client.ListContainers(docker.ListContainersOptions{
		All: true,
		Filters: map[string][]string{
			"label": {fmt.Sprintf("%s=%s", key, value)},
		},
	})

	l.Require().NoError(err)
	return containers
}
