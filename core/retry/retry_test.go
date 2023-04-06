package retry_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/retry"
	"testing"
	"time"
)

func TestRetryWithBackoff(t *testing.T) {
	// Test a function that succeeds on the first attempt.
	t.Run("Success", func(t *testing.T) {
		err := retry.WithBackoff(context.Background(), func(ctx context.Context) error {
			return nil
		})
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	// Test a function that fails and eventually succeeds.
	t.Run("SuccessAfterRetries", func(t *testing.T) {
		var i int
		err := retry.WithBackoff(context.Background(), func(ctx context.Context) error {
			if i < 2 {
				i++
				return errors.New("simulated error")
			}
			return nil
		}, retry.WithMaxAttempts(3))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	// Test a function that always fails.
	t.Run("Failure", func(t *testing.T) {
		err := retry.WithBackoff(context.Background(), func(ctx context.Context) error {
			return errors.New("simulated error")
		}, retry.WithMaxAttempts(3))
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	// Test a function that times out.
	t.Run("Timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		err := retry.WithBackoff(ctx, func(ctx context.Context) error {
			return errors.New("simulated error")
		})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})

	t.Run("WithFactor", func(t *testing.T) {
		cfg := retry.DefaultConfig()
		factor := gofakeit.Float64Range(1, 10)
		newCfg := retry.Configurator(t, cfg, retry.WithFactor(factor))
		Equal(t, factor, newCfg.GetFactor())
	})

	t.Run("WithJitter", func(t *testing.T) {
		cfg := retry.DefaultConfig()
		jitter := gofakeit.Bool()
		newCfg := retry.Configurator(t, cfg, retry.WithJitter(jitter))
		Equal(t, jitter, newCfg.GetJitter())
	})

	t.Run("WithMin", func(t *testing.T) {
		cfg := retry.DefaultConfig()
		withMin := time.Duration(gofakeit.Uint64())
		newCfg := retry.Configurator(t, cfg, retry.WithMin(withMin))
		Equal(t, withMin, newCfg.GetMin())
	})

	t.Run("WithMax", func(t *testing.T) {
		cfg := retry.DefaultConfig()
		withMax := time.Duration(gofakeit.Uint64())
		newCfg := retry.Configurator(t, cfg, retry.WithMax(withMax))
		Equal(t, withMax, newCfg.GetMax())
	})

	t.Run("WithMaxAttemptsTime", func(t *testing.T) {
		err := retry.WithBackoff(context.Background(), func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return fmt.Errorf("context canceled: %w", ctx.Err())
			case <-time.After(time.Millisecond):
				return nil
			}
		}, retry.WithMaxAttemptsTime(1))
		NotNil(t, err)
	})
}
