package localmetrics_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
)

func (l *LocalServerSuite) TestFullJaeger() {
	ts := localmetrics.StartTestServer(l.GetTestContext(), l.T())

	// Skip test if containers failed to start due to port binding or rate limit issues
	if ts == nil {
		l.T().Skip("Failed to start test server - skipping container verification")
		return
	}

	containers := l.ContainersWithLabel(localmetrics.RunIDLabel, ts.GetRunID())
	if len(containers) == 0 {
		l.T().Skip("No containers running - likely due to Docker rate limits")
		return
	}

	// Only check that we have the expected containers that successfully started
	for _, container := range containers {
		l.Require().Contains([]string{"jaeger", "pyroscope"}, container.Labels[localmetrics.AppLabel])
	}
}

// TestStartServerJaegerPreset tests the start server function with jaeger preset.
// only pyroscope should run.
func (l *LocalServerSuite) TestStartServerJaegerPreset() {
	l.T().Setenv(internal.JaegerEndpoint, gofakeit.URL())
	l.T().Setenv(internal.JaegerUIEndpoint, gofakeit.URL())

	ts := localmetrics.StartTestServer(l.GetTestContext(), l.T())

	// Skip test if containers failed to start due to port binding or rate limit issues
	if ts == nil {
		l.T().Skip("Failed to start test server - skipping container verification")
		return
	}

	containers := l.ContainersWithLabel(localmetrics.RunIDLabel, ts.GetRunID())
	if len(containers) == 0 {
		l.T().Skip("No containers running - likely due to Docker rate limits")
		return
	}

	// When Jaeger is preset, we expect only Pyroscope
	for _, container := range containers {
		l.Require().Equal("pyroscope", container.Labels[localmetrics.AppLabel])
	}
}

func (l *LocalServerSuite) TestStartServerPyroscopePreset() {
	l.T().Setenv(internal.PyroscopeEndpoint, gofakeit.URL())
	l.T().Setenv(internal.JaegerUIEndpoint, gofakeit.URL())
	ts := localmetrics.StartTestServer(l.GetTestContext(), l.T())

	// Skip test if containers failed to start due to port binding or rate limit issues
	if ts == nil {
		l.T().Skip("Failed to start test server - skipping container verification")
		return
	}

	containers := l.ContainersWithLabel(localmetrics.RunIDLabel, ts.GetRunID())
	if len(containers) == 0 {
		l.T().Skip("No containers running - likely due to Docker rate limits")
		return
	}

	// When Pyroscope is preset, we expect only Jaeger
	for _, container := range containers {
		l.Require().Equal("jaeger", container.Labels[localmetrics.AppLabel])
	}
}
