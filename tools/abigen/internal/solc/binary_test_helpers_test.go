// Package solc_test provides test utilities for the solc package.
package solc_test

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/tools/abigen/internal/solc"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// testConcurrentDownload is a helper function that tests concurrent download handling.
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

// testBackoffWithJitter is a helper function that tests the backoff with jitter functionality.
func testBackoffWithJitter(t *testing.T, manager *solc.BinaryManager) {
	t.Helper()
	attempt := 2

	// Ensure attempt is non-negative
	if attempt < 0 {
		attempt = 0
	}

	start := time.Now()
	delay := manager.ApplyBackoffWithJitter(attempt)

	// Calculate expected bounds with safe conversion
	const baseMs = 100
	baseDelay := time.Duration(baseMs) * time.Millisecond
	if attempt > 0 {
		baseDelay = baseDelay << uint(attempt)
	}
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

// testRetryableError is a helper function that tests the IsRetryableError functionality.
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

// testConcurrentDownloadHandling tests the concurrent download handling through GetBinary.
func testConcurrentDownloadHandling(t *testing.T, manager *solc.BinaryManager) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get the actual cache directory path with correct structure
	platformDir := filepath.Join(manager.CacheDir(), "0.8.20", "linux-amd64")
	binaryPath := filepath.Join(platformDir, "solc-linux-amd64-v0.8.20")

	// Create directory structure
	if err := os.MkdirAll(platformDir, 0750); err != nil {
		t.Fatalf("Failed to create platform directory: %v", err)
	}

	// Create a binary file to simulate successful concurrent download
	if err := os.WriteFile(binaryPath, []byte("test binary"), 0600); err != nil {
		t.Fatalf("Failed to create test binary: %v", err)
	}

	// GetBinary should succeed and return the existing binary
	path, err := manager.GetBinary(ctx)
	if err != nil {
		t.Errorf("GetBinary() error = %v, want nil", err)
	}
	if path != binaryPath {
		t.Errorf("GetBinary() path = %v, want %v", path, binaryPath)
	}

	// Remove binary to force download attempt
	if err := os.Remove(binaryPath); err != nil {
		t.Fatalf("Failed to remove test binary: %v", err)
	}

	// GetBinary should handle the missing binary case
	path, err = manager.GetBinary(ctx)
	if err != nil {
		t.Errorf("GetBinary() error = %v, want nil", err)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("GetBinary() did not create missing binary")
	}
}

// testTempFileErrorHandling tests error handling for temp files through GetBinary.
func testTempFileErrorHandling(t *testing.T, manager *solc.BinaryManager) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get the actual cache directory path with correct structure
	platformDir := filepath.Join(manager.CacheDir(), "0.8.20", "linux-amd64")
	if err := os.MkdirAll(platformDir, 0750); err != nil {
		t.Fatalf("Failed to create platform directory: %v", err)
	}

	// Remove any existing binary to force download attempt
	binaryPath := filepath.Join(platformDir, "solc-linux-amd64-v0.8.20")
	if err := os.Remove(binaryPath); err != nil && !os.IsNotExist(err) {
		t.Fatalf("Failed to remove existing binary: %v", err)
	}

	// Make the platform directory read-only to trigger permission error
	if err := os.Chmod(platformDir, 0400); err != nil {
		t.Fatalf("Failed to make directory read-only: %v", err)
	}
	defer func() {
		// Restore permissions to allow cleanup
		if err := os.Chmod(platformDir, 0600); err != nil {
			t.Errorf("Failed to restore directory permissions: %v", err)
		}
	}()

	// GetBinary should fail with permission error
	_, err := manager.GetBinary(ctx)
	if err == nil {
		t.Error("GetBinary() succeeded, want permission error")
	} else if !os.IsPermission(err) && !strings.Contains(err.Error(), "permission denied") {
		t.Errorf("GetBinary() error = %v, want permission denied error", err)
	}
}
