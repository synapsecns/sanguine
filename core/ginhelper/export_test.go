package ginhelper

import "time"

// DefaultServerTimeout is the default timeout for the server to start
// this copies the original timeout for testing.
var DefaultServerTimeout = serverStartTimeout

// ResetServerTimeout resets the server timeout to the default value.
func ResetServerTimeout() {
	serverStartTimeout = DefaultServerTimeout
}

// SetServerTimeout sets the server timeout to the given value.
func SetServerTimeout(timeout time.Duration) {
	serverStartTimeout = timeout
}
