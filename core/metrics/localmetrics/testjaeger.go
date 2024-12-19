package localmetrics

import (
	"context"
	"os"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type testJaeger struct {
	runID       string
	containerID string
	client      interface {
		ContainerStop(ctx context.Context, containerID string, options container.StopOptions) error
		ContainerRemove(ctx context.Context, containerID string, options container.RemoveOptions) error
	}
}

func (tj *testJaeger) cleanup(ctx context.Context) error {
	if tj.containerID != "" && tj.client != nil {
		if err := tj.client.ContainerStop(ctx, tj.containerID, container.StopOptions{}); err != nil {
			return err
		}
		return tj.client.ContainerRemove(ctx, tj.containerID, container.RemoveOptions{})
	}
	return nil
}

// SetupTestJaeger creates a new test jaeger instance. If the test fails, the instance is kept alive for 5 minutes.
// we also allow a GLOBAL_jaeger env var to be set to a jaeger url to send all traces to in order to avoid having to boot for long running tests.
func SetupTestJaeger(ctx context.Context, tb testing.TB, opts ...Option) *TestServer {
	tb.Helper()

	// disable pyroscope in CI to avoid conflicts with the pyroscope test runner
	if os.Getenv("CI") != "" {
		opts = append(opts, WithPyroscopeEnabled(false))
	}

	tj := startServer(ctx, tb, opts...)
	if tj == nil {
		tb.Logf("Failed to set up test Jaeger - continuing with limited functionality")
		return nil
	}
	return &TestServer{jaeger: tj}
}
