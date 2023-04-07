package core

import (
	"context"
	"time"
)

// SleepWithContext will wait for the timer duration to expire, or the context
// is canceled. Which ever happens first. If the context is canceled the Context's
// error will be returned.
//
// Expects Context to always return a non-nil error if the Done channel is closed.
// Note: this was taknen from the aws sdk, but modified to avoid the dependency in core.
func SleepWithContext(ctx context.Context, dur time.Duration) error {
	t := time.NewTimer(dur)
	defer t.Stop()

	select {
	case <-t.C:
		break
	case <-ctx.Done():
		//nolint: wrapcheck
		return ctx.Err()
	}

	return nil
}
