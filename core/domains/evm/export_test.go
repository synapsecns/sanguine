package evm

import (
	"time"
)

func init() {
	defaultMinBackoff = minBackoff
	defaultMaxBackoff = maxBackoff
}

// defaultMinBackoff is the default wait before timeout. We capture this on startup and expose it via
// a getter so it can be reset after each test. This is not thread safe.
var defaultMinBackoff time.Duration

// defaultMaxBackoff is the default max backoff time. We capture this on startup and expose it via
// a getter so it can be reset after each test. This is not thread safe.
var defaultMaxBackoff time.Duration

// DefaultMinBackoff returns the default wait before attempt so it can be reset before each test.
func DefaultMinBackoff() time.Duration {
	return defaultMinBackoff
}

// DefaultMinBackoff returns the default wait before attempt so it can be reset before each test.
func DefaultMaxBackoff() time.Duration {
	return defaultMaxBackoff
}

// SetMinBackoff sets the minimum backoff time for testing.
func SetMinBackoff(duration time.Duration) {
	minBackoff = duration
}

// SetMaxBackoff sets the maximum backoff time for testing.
func SetMaxBackoff(duration time.Duration) {
	maxBackoff = duration
}

// MaxAttempts is the maximum number of attempts filter ahead will make exported for testing.
const MaxAttempts = maxAttempts
