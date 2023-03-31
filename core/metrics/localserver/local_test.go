package localserver_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/metrics/localserver"
)

func (l *LocalServerSuite) TestFullJaeger() {
	ts := localserver.StartTestServer(l.GetTestContext(), l.T())

	containers := l.ContainersWithLabel(localserver.RunIDLabel, ts.GetRunID())
	l.Require().Len(containers, 3)
}

// TestStartServerJaegerPreset tests the start server function with jaeger preset.
// only pyroscope should run.
func (l *LocalServerSuite) TestStartServerJaegerPreset() {
	l.T().Setenv(internal.JaegerEndpoint, gofakeit.URL())
	l.T().Setenv(internal.JaegerUIEndpoint, gofakeit.URL())

	ts := localserver.StartTestServer(l.GetTestContext(), l.T())

	containers := l.ContainersWithLabel(localserver.RunIDLabel, ts.GetRunID())
	l.Require().Len(containers, 1)

	l.Require().Equal(containers[0].Labels[localserver.AppLabel], "pyroscope")
}

func (l *LocalServerSuite) TestStartServerPyroscopePreset() {
	l.T().Setenv(internal.PyroscopeEndpoint, gofakeit.URL())
	l.T().Setenv(internal.JaegerUIEndpoint, gofakeit.URL())
	ts := localserver.StartTestServer(l.GetTestContext(), l.T())

	containers := l.ContainersWithLabel(localserver.RunIDLabel, ts.GetRunID())
	l.Require().Len(containers, 1)

	l.Require().Equal(containers[0].Labels[localserver.AppLabel], "jaeger")
}
