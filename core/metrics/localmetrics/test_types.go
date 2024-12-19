package localmetrics

import (
	"context"
)

// TestServer represents a test server.
type TestServer struct {
	jaeger *testJaeger
}

// GetRunID returns the run ID of the test server.
func (ts *TestServer) GetRunID() string {
	return ts.jaeger.runID
}

// Cleanup cleans up the test server's resources.
func (ts *TestServer) Cleanup(ctx context.Context) error {
	if ts.jaeger != nil {
		return ts.jaeger.cleanup(ctx)
	}
	return nil
}
