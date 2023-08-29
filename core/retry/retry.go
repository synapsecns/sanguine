package retry

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/jpillora/backoff"
	errorUtil "github.com/pkg/errors"
)

// RetryableFunc is a function that can be retried.
type RetryableFunc func(ctx context.Context) error

// FuncWrapper wraps a function to be generic. This should be checked to be a function before calling.
type FuncWrapper any

// retryWithBackoffConfig holds the configuration for WithBackoff.
type retryWithBackoffConfig struct {
	factor float64
	jitter bool
	min    time.Duration
	max    time.Duration
	// maxAttempts sets the maximum number of retry attempts.
	// if this is negative it is ignored
	maxAttempts    int
	maxAttemptTime time.Duration
	// maxAllAttempts sets the maximum time for all attempts.
	// if this is negative it is ignored
	maxAllAttemptsTime time.Duration
}

// returns true if the number of attempts exceeds the maximum number of attempts.
func (r *retryWithBackoffConfig) exceedsMaxAttempts(attempts int) bool {
	return r.maxAttempts > 0 && attempts > r.maxAttempts
}

// returns true if the time for all attempts exceeds the maximum time for all attempts.
func (r *retryWithBackoffConfig) exceedsMaxTime(startTime time.Time) bool {
	return r.maxAllAttemptsTime > 0 && time.Since(startTime) > r.maxAllAttemptsTime
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

// WithMaxAttemptTime sets the maximum time of all retry attempts.
func WithMaxAttemptTime(maxAttemptTime time.Duration) WithBackoffConfigurator {
	return func(c *retryWithBackoffConfig) {
		c.maxAttemptTime = maxAttemptTime
	}
}

// WithMaxTotalTime sets the maximum time of all retry attempts combined.
func WithMaxTotalTime(maxTotalTime time.Duration) WithBackoffConfigurator {
	return func(c *retryWithBackoffConfig) {
		c.maxAllAttemptsTime = maxTotalTime
	}
}

func defaultConfig() retryWithBackoffConfig {
	return retryWithBackoffConfig{
		factor:             2,
		jitter:             true,
		min:                200 * time.Millisecond,
		max:                5 * time.Second,
		maxAttempts:        -1,
		maxAllAttemptsTime: time.Second * 30,
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
	startTime := time.Now()

	attempts := 0
	for !config.exceedsMaxAttempts(attempts) && !config.exceedsMaxTime(startTime) {
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

	if config.exceedsMaxAttempts(attempts) {
		return errorUtil.Wrapf(ErrMaxAttempts, "after %d attempts", attempts)
	}
	if config.exceedsMaxTime(startTime) {
		return errorUtil.Wrapf(ErrMaxTime, "after %s (max was %s)", time.Since(startTime).String(), config.maxAllAttemptsTime.String())
	}

	return ErrUnknown
}

// WithBackoffOneArg retries the given one arg function with exponential backoff.
func WithBackoffOneArg(ctx context.Context, doFunc FuncWrapper, arg interface{}, configurators ...WithBackoffConfigurator) (val interface{}, err error) {
	funcValue := reflect.ValueOf(doFunc)

	if funcValue.Kind() != reflect.Func {
		return nil, fmt.Errorf("doFunc is not a function")
	}

	funcType := reflect.TypeOf(doFunc)

	// Ensure the function has the expected signature.
	if funcType.NumIn() != 2 || funcType.NumOut() != 2 {
		return nil, fmt.Errorf("function does not have the expected signature")
	}

	if !reflect.TypeOf(ctx).AssignableTo(funcType.In(0)) {
		return nil, fmt.Errorf("context argument is not assignable to function's first parameter type")
	}

	if !reflect.TypeOf(arg).AssignableTo(funcType.In(1)) {
		return nil, fmt.Errorf("provided argument is not assignable to function's second parameter type")
	}

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
	startTime := time.Now()

	attempts := 0
	for !config.exceedsMaxAttempts(attempts) && !config.exceedsMaxTime(startTime) {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("%w while retrying", ctx.Err())
		case <-time.After(timeout):
			var funcCtx context.Context
			var cancel context.CancelFunc

			if config.maxAttemptTime > 0 {
				funcCtx, cancel = context.WithTimeout(ctx, config.maxAttemptTime)
			} else {
				funcCtx, cancel = context.WithCancel(ctx)
			}

			// Create input args.
			inputArgs := []reflect.Value{
				reflect.ValueOf(funcCtx),
				reflect.ValueOf(arg),
			}

			// Call the function.
			output := funcValue.Call(inputArgs)

			// Extract the return values.
			val = output[0].Interface()
			err := output[1].Interface().(error)

			if err != nil {
				timeout = b.Duration()
				attempts++
				cancel()
			} else {
				cancel()
				return val, nil
			}
		}
	}

	if config.exceedsMaxAttempts(attempts) {
		return nil, errorUtil.Wrapf(ErrMaxAttempts, "after %d attempts", attempts)
	}
	if config.exceedsMaxTime(startTime) {
		return nil, errorUtil.Wrapf(ErrMaxTime, "after %s (max was %s)", time.Since(startTime).String(), config.maxAllAttemptsTime.String())
	}

	return nil, ErrUnknown
}

// ErrMaxAttempts is returned when the maximum number of retry attempts is reached.
var ErrMaxAttempts = errors.New("max attempts reached")

// ErrMaxTime is returned when the maximum time for all retry attempts is reached.
var ErrMaxTime = errors.New("max time reached")

// ErrUnknown is returned when an unknown error occurs.
var ErrUnknown = errors.New("unknown error")
