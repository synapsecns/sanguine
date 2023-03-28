package retry

import (
	. "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// RetryWithBackoffConfig holds the configuration for WithBackoff.
type RetryWithBackoffConfig interface {
	GetFactor() float64
	GetJitter() bool
	GetMin() time.Duration
	GetMax() time.Duration
	GetMaxAttempts() int
}

func (r *retryWithBackoffConfig) GetFactor() float64 {
	return r.factor
}

func (r *retryWithBackoffConfig) GetJitter() bool {
	return r.jitter
}

func (r *retryWithBackoffConfig) GetMin() time.Duration {
	return r.min
}

func (r *retryWithBackoffConfig) GetMax() time.Duration {
	return r.max
}

func (r *retryWithBackoffConfig) GetMaxAttempts() int {
	return r.maxAttempts
}

// DefaultConfig returns a default RetryWithBackoffConfig.
func DefaultConfig() RetryWithBackoffConfig {
	cfg := defaultConfig()
	return &cfg
}

// Configurator returns a RetryWithBackoffConfig with the given configurators applied.
func Configurator(tb testing.TB, configurator RetryWithBackoffConfig, apply ...WithBackoffConfigurator) RetryWithBackoffConfig {
	tb.Helper()

	for _, cfg := range apply {
		res, ok := configurator.(*retryWithBackoffConfig)
		True(tb, ok, "configurator is not a *retryWithBackoffConfig")

		cfg(res)
	}
	return configurator
}
