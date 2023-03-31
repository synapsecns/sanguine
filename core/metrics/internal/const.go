// Package internal contains internal constants for metrics package
// this is done in a separate package to avoid circular dependencies
package internal

const (
	// JaegerEndpoint is the environment variable name for Jaeger endpoint.
	JaegerEndpoint = "JAEGER_ENDPOINT"
	// PyroscopeEndpoint is the environment variable name for Pyroscope endpoint.
	PyroscopeEndpoint = "PYROSCOPE_ENDPOINT"
	// JaegerUiEndpoint is the environment variable name for Jaeger UI.
	JaegerUiEndpoint = "JAEGER_UI"
	// PyroscopeStarted is the environment variable name for Pyroscope started.
	PyroscopeStarted = "PYROSCOPE_STARTED"
)
