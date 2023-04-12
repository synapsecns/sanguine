package localmetrics_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
)

func (l *LocalServerSuite) TestFullJaeger() {
	ts := localmetrics.StartTestServer(l.GetTestContext(), l.T())

	containers := l.ContainersWithLabel(localmetrics.RunIDLabel, ts.GetRunID())
	l.Require().Len(containers, 3)
}

// TestStartServerJaegerPreset tests the start server function with jaeger preset.
// only pyroscope should run.
func (l *LocalServerSuite) TestStartServerJaegerPreset() {
	l.T().Setenv(internal.JaegerEndpoint, gofakeit.URL())
	l.T().Setenv(internal.JaegerUIEndpoint, gofakeit.URL())

	ts := localmetrics.StartTestServer(l.GetTestContext(), l.T())

	containers := l.ContainersWithLabel(localmetrics.RunIDLabel, ts.GetRunID())
	l.Require().Len(containers, 1)

	l.Require().Equal(containers[0].Labels[localmetrics.AppLabel], "pyroscope")
}

func (l *LocalServerSuite) TestStartServerPyroscopePreset() {
	l.T().Setenv(internal.PyroscopeEndpoint, gofakeit.URL())
	l.T().Setenv(internal.JaegerUIEndpoint, gofakeit.URL())
	ts := localmetrics.StartTestServer(l.GetTestContext(), l.T())

	containers := l.ContainersWithLabel(localmetrics.RunIDLabel, ts.GetRunID())
	l.Require().Len(containers, 1)

	l.Require().Equal(containers[0].Labels[localmetrics.AppLabel], "jaeger")
}
