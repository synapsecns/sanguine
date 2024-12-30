// Package internal contains internal constants for metrics package
// this is done in a separate package to avoid circular dependencies
package internal

const (
	// JaegerEndpoint is the environment variable name for Jaeger endpoint.
	JaegerEndpoint = "JAEGER_ENDPOINT"
	// PyroscopeEndpoint is the environment variable name for Pyroscope endpoint.
	PyroscopeEndpoint = "PYROSCOPE_ENDPOINT"
	// JaegerUIEndpoint is the environment variable name for Jaeger UI.
	JaegerUIEndpoint = "JAEGER_UI"
	// PyroscopeJaegerUIEnabled is the environment variable name for enabling Jaeger UI.
	PyroscopeJaegerUIEnabled = "PYROSCOPE_JAEGER_UI_ENABLED"
)
