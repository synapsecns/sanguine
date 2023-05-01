package localmetrics

import (
	"context"
	"os"
	"testing"
)

// SetupTestJaeger creates a new test jaeger instance. If the test fails, the instance is kept alive for 5 minutes.
// we also allow a GLOBAL_jaeger env var to be set to a jaeger url to send all traces to in order to avoid having to boot for long running tests.
func SetupTestJaeger(ctx context.Context, tb testing.TB, opts ...Option) {
	tb.Helper()

	// disable pyroscope in CI to avoid conflicts with the pyroscope test runner
	if os.Getenv("CI") != "" {
		opts = append(opts, WithPyroscopeEnabled(false))
	}

	startServer(ctx, tb, opts...)
}
