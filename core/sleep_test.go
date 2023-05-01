package core_test

import (
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"testing"
	"time"
)

func TestSleepWithContext(t *testing.T) {
	ctx := &FakeContext{DoneCh: make(chan struct{})}

	err := core.SleepWithContext(ctx, 1*time.Millisecond)
	if err != nil {
		t.Errorf("expect context to not be canceled, got %v", err)
	}
}

func TestSleepWithContext_Canceled(t *testing.T) {
	ctx := &FakeContext{DoneCh: make(chan struct{})}

	expectErr := fmt.Errorf("context canceled")

	ctx.Error = expectErr
	close(ctx.DoneCh)

	err := core.SleepWithContext(ctx, 10*time.Second)
	if err == nil {
		t.Fatalf("expect error, did not get one")
	}

	if e, a := expectErr, err; !errors.Is(e, a) {
		t.Errorf("expect %v error, got %v", e, a)
	}
}

// A FakeContext provides a simple stub implementation of a Context.
type FakeContext struct {
	Error  error
	DoneCh chan struct{}
}

// Deadline always will return not set.
func (c *FakeContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Done returns a read channel for listening to the Done event.
func (c *FakeContext) Done() <-chan struct{} {
	return c.DoneCh
}

// Err returns the error, is nil if not set.
func (c *FakeContext) Err() error {
	return c.Error
}

// Value ignores the Value and always returns nil.
func (c *FakeContext) Value(key interface{}) interface{} {
	return nil
}
