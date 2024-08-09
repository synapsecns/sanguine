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
	// RookoutToken is used for https://app.rookout.com/
	RookoutToken = "ROOKOUT_TOKEN"
	// RookoutCommit is the environment variable for the git commit in use.
	RookoutCommit = "ROOKOUT_COMMIT"
	// RookoutDebug can be used to enable rookout debug.
	RookoutDebug = "ROOKOUT_DEBUG"
	// RookoutRemoteOrigin is the environment variable for the git remote origin in use.
	RookoutRemoteOrigin = "ROOKOUT_REMOTE_ORIGIN"
	// GitRepo is the environment variable for the git repo in use.
	GitRepo = "GIT_REPO"
)
