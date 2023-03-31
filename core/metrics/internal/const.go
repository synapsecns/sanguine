// Package internal contains internal constants for metrics package
// this is done in a separate package to avoid circular dependencies
package internal

const (
	// JAEGER_ENDPOINT is the environment variable name for Jaeger endpoint.
	JAEGER_ENDPOINT = "JAEGER_ENDPOINT"
	// PYROSCOPE_ENDPOINT is the environment variable name for Pyroscope endpoint.
	PYROSCOPE_ENDPOINT = "PYROSCOPE_ENDPOINT"
	// JAEGER_UI_ENDPOINT is the environment variable name for Jaeger UI.
	JAEGER_UI_ENDPOINT = "JAEGER_UI"
)
