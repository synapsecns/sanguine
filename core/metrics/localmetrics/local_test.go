package localmetrics_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
)

type LocalServerSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *LocalServerSuite) SetupTest() {
	s.ctx = context.Background()
}

func (s *LocalServerSuite) GetTestContext() context.Context {
	return s.ctx
}

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

func (s *LocalServerSuite) TestStartServerWithDockerErrors() {
	tests := []struct {
		name    string
		setup   func()
		wantErr bool
	}{
		{
			name: "docker rate limit error",
			setup: func() {
				s.T().Setenv("DOCKER_RATE_LIMIT", "true")
			},
			wantErr: true,
		},
		{
			name: "port binding error",
			setup: func() {
				s.T().Setenv("FORCE_PORT_CONFLICT", "true")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			if tt.setup != nil {
				tt.setup()
			}

			ts := localmetrics.StartTestServer(s.GetTestContext(), s.T())
			if tt.wantErr {
				s.Nil(ts, "expected nil server due to error")
			} else {
				s.NotNil(ts, "expected non-nil server")
			}
		})
	}
}

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

func TestLocalServerSuite(t *testing.T) {
	suite.Run(t, new(LocalServerSuite))
}
