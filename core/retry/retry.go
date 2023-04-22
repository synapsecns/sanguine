package retry

import (
	"context"
	"errors"
	"fmt"
	"github.com/jpillora/backoff"
	"time"
)

// RetryableFunc is a function that can be retried.
type RetryableFunc func(ctx context.Context) error

// retryWithBackoffConfig holds the configuration for WithBackoff.
type retryWithBackoffConfig struct {
	factor         float64
	jitter         bool
	min            time.Duration
	max            time.Duration
	maxAttempts    int
	maxAttemptTime time.Duration
}

// WithBackoffConfigurator configures a retryWithBackoffConfig.
type WithBackoffConfigurator func(*retryWithBackoffConfig)

// WithFactor sets the backoff factor.
func WithFactor(factor float64) WithBackoffConfigurator {
	return func(c *retryWithBackoffConfig) {
		c.factor = factor
	}
}

// WithJitter enables or disables jitter.
func WithJitter(jitter bool) WithBackoffConfigurator {
	return func(c *retryWithBackoffConfig) {
		c.jitter = jitter
	}
}

// WithMin sets the minimum backoff duration.
func WithMin(min time.Duration) WithBackoffConfigurator {
	return func(c *retryWithBackoffConfig) {
		c.min = min
	}
}

// WithMax sets the maximum backoff duration.
func WithMax(max time.Duration) WithBackoffConfigurator {
	return func(c *retryWithBackoffConfig) {
		c.max = max
	}
}

// WithMaxAttempts sets the maximum number of retry attempts.
func WithMaxAttempts(maxAttempts int) WithBackoffConfigurator {
	return func(c *retryWithBackoffConfig) {
		c.maxAttempts = maxAttempts
	}
}

// WithMaxAttemptsTime sets the maximum time of all retry attempts.
func WithMaxAttemptsTime(maxAttemptTime time.Duration) WithBackoffConfigurator {
	return func(c *retryWithBackoffConfig) {
		c.maxAttemptTime = maxAttemptTime
	}
}

func defaultConfig() retryWithBackoffConfig {
	return retryWithBackoffConfig{
		factor:      2,
		jitter:      true,
		min:         200 * time.Millisecond,
		max:         5 * time.Second,
		maxAttempts: 3,
	}
}

// WithBackoff retries the given function with exponential backoff.
func WithBackoff(ctx context.Context, doFunc RetryableFunc, configurators ...WithBackoffConfigurator) error {
	config := defaultConfig()

	for _, configurator := range configurators {
		configurator(&config)
	}

	b := &backoff.Backoff{
		Factor: config.factor,
		Jitter: config.jitter,
		Min:    config.min,
		Max:    config.max,
	}

	timeout := time.Duration(0)
	attempts := 0
	for attempts < config.maxAttempts {
		select {
		case <-ctx.Done():
			return fmt.Errorf("%w while retrying", ctx.Err())
		case <-time.After(timeout):
			var funcCtx context.Context
			var cancel context.CancelFunc

			if config.maxAttemptTime > 0 {
				funcCtx, cancel = context.WithTimeout(ctx, config.maxAttemptTime)
			} else {
				funcCtx, cancel = context.WithCancel(ctx)
			}

			err := doFunc(funcCtx)
			if err != nil {
				timeout = b.Duration()
				attempts++
				cancel()
			} else {
				cancel()
				return nil
			}
		}
	}

	return ErrMaxAttempts
}

// ErrMaxAttempts is returned when the maximum number of retry attempts is reached.
var ErrMaxAttempts = errors.New("max attempts reached")
