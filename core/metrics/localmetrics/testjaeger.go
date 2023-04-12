package localmetrics

import (
	"context"
	"testing"
)

// TestSuite defines a testSuite that exposes functions for running after a suite
type TestSuite interface {
	// DeferAfterSuite adds a function to be run after the test
	DeferAfterSuite(newFunc func())
	// HasFailures returns true if the test suite has any failures
	HasFailures() bool
	// T is the testing.T interface
	T() *testing.T
}

// SetupTestJaeger creates a new test jaeger instance. If the test fails, the instance is kept alive for 5 minutes.
// we also allow a GLOBAL_jaeger env var to be set to a jaeger url to send all traces to in order to avoid having to boot for long running tests.
func SetupTestJaeger(ctx context.Context, ts TestSuite, opts ...Option) {
	startServer(ctx, ts, opts...)
}
