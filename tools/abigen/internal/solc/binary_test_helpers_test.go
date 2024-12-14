// Package solc_test provides test utilities for the solc package.
package solc_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
	"github.com/synapsecns/sanguine/tools/abigen/internal/solc"
)

// testConcurrentDownload is a helper function that tests concurrent download handling
func testConcurrentDownload(t *testing.T, manager *solc.BinaryManager) {
	t.Helper()

	// Create a temporary file to simulate concurrent download
	tempFile, err := os.CreateTemp(manager.CacheDir(), "solc-*.tmp")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Errorf("Failed to cleanup temp file: %v", err)
		}
	})

	// Write some content to simulate a partial download
	content := []byte("partial download content")
	if err := os.WriteFile(tempFile.Name(), content, 0600); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	// Test handling of concurrent download
	exists, err := manager.CheckConcurrentDownload(tempFile.Name())
	if err != nil {
		t.Errorf("CheckConcurrentDownload() error = %v, want nil", err)
	}
	if !exists {
		t.Error("CheckConcurrentDownload() = false, want true")
	}
}

// testBackoffWithJitter is a helper function that tests the backoff with jitter functionality
func testBackoffWithJitter(t *testing.T, manager *solc.BinaryManager) {
	t.Helper()
	attempt := 2
	start := time.Now()
	delay := manager.ApplyBackoffWithJitter(attempt)

	// Calculate expected bounds
	baseDelay := time.Duration(100*(1<<uint(attempt))) * time.Millisecond
	maxJitter := baseDelay / 2

	if delay < baseDelay || delay > baseDelay+maxJitter {
		t.Errorf("ApplyBackoffWithJitter() delay = %v, want between %v and %v",
			delay, baseDelay, baseDelay+maxJitter)
	}

	// Actually wait to ensure coverage
	time.Sleep(delay)
	elapsed := time.Since(start)
	if elapsed < delay {
		t.Errorf("Sleep duration %v less than calculated delay %v", elapsed, delay)
	}
}

// testRetryableError is a helper function that tests the IsRetryableError functionality
func testRetryableError(t *testing.T) {
	t.Helper()
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "temporary error",
			err:  &os.PathError{Err: &os.PathError{Err: os.ErrExist}},
			want: true,
		},
		{
			name: "permission error",
			err:  os.ErrPermission,
			want: true,
		},
		{
			name: "context canceled",
			err:  context.Canceled,
			want: true,
		},
		{
			name: "context deadline exceeded",
			err:  context.DeadlineExceeded,
			want: true,
		},
		{
			name: "non-retryable error",
			err:  fmt.Errorf("non-retryable error"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solc.IsRetryableError(tt.err); got != tt.want {
				t.Errorf("IsRetryableError() = %v, want %v", got, tt.want)
			}
		})
	}
}
