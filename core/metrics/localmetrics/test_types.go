package localmetrics

// TestServer represents a test server.
type TestServer struct {
	jaeger *testJaeger
}

// GetRunID returns the run ID of the test server.
func (ts *TestServer) GetRunID() string {
	return ts.jaeger.runID
}
