package solc_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/synapsecns/sanguine/tools/abigen/internal/solc"
)

func setupTestBinaryManager(t *testing.T) *solc.BinaryManager {
	t.Helper()
	return solc.NewBinaryManager("0.8.20")
}

func TestIsAppleSilicon(t *testing.T) {
	isArm := runtime.GOARCH == "arm64"
	isDarwin := runtime.GOOS == "darwin"

	result := solc.IsAppleSilicon()

	if isArm && isDarwin {
		if !result {
			t.Error("Expected true for Apple Silicon (darwin/arm64)")
		}
	} else {
		if result {
			t.Error("Expected false for non-Apple Silicon platform")
		}
	}
}

func TestNewBinaryManager(t *testing.T) {
	manager := setupTestBinaryManager(t)

	expectedCacheDir := filepath.Join(os.Getenv("HOME"), ".cache", "solc")
	if manager.CacheDir() != expectedCacheDir {
		t.Errorf("Expected cache dir %s, got %s", expectedCacheDir, manager.CacheDir())
	}

	if manager.Version() != "0.8.20" {
		t.Errorf("Expected version %s, got %s", "0.8.20", manager.Version())
	}
}

func TestGetPlatformDir(t *testing.T) {
	manager := solc.NewBinaryManager("0.8.20")
	platform, err := manager.GetPlatformDir()
	if err != nil {
		t.Fatalf("GetPlatformDir() error = %v", err)
	}

	if solc.IsAppleSilicon() {
		if platform != solc.PlatformWasm32 {
			t.Error("Expected wasm for Apple Silicon")
		}
		return
	}

	var expectedPlatform string
	switch runtime.GOOS {
	case solc.PlatformDarwin:
		expectedPlatform = solc.PlatformMacOS
	case "linux":
		expectedPlatform = solc.PlatformLinux
	case "windows":
		expectedPlatform = solc.PlatformWin
	default:
		expectedPlatform = solc.PlatformWasm32
	}

	if platform != expectedPlatform {
		t.Errorf("Expected platform %s, got %s", expectedPlatform, platform)
	}
}

func TestGetBinary(t *testing.T) {
	manager := setupTestBinaryManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	binary, err := manager.GetBinary(ctx)
	if err != nil {
		t.Fatalf("Failed to get binary: %v", err)
	}

	if _, statErr := os.Stat(binary); os.IsNotExist(statErr) {
		t.Error("Binary file does not exist")
	}

	info, err := os.Stat(binary)
	if err != nil {
		t.Fatalf("Failed to stat binary: %v", err)
	}
	if info.Mode()&0111 == 0 {
		t.Error("Binary is not executable")
	}

	binary2, err := manager.GetBinary(ctx)
	if err != nil {
		t.Fatalf("Failed to get cached binary: %v", err)
	}
	if binary != binary2 {
		t.Error("Cache not reused")
	}
}

func TestGetBinaryInfo(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "solc-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Errorf("failed to cleanup test directory: %v", err)
		}
	}()

	tests := []struct {
		name     string
		version  string
		platform string
		wantErr  string
	}{
		{
			name:     "invalid version",
			version:  "999.999.999",
			platform: "linux-amd64",
			wantErr:  "failed to get binary info: version 999.999.999 not found for platform linux-amd64",
		},
		{
			name:     "invalid platform",
			version:  "0.8.20",
			platform: "invalid-platform",
			wantErr:  "unsupported platform",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			manager := solc.NewBinaryManager(tt.version)
			managerValue := reflect.ValueOf(manager).Elem()
			if platformField := managerValue.FieldByName("Platform"); platformField.IsValid() && platformField.CanSet() {
				platformField.SetString(tt.platform)
				t.Logf("DEBUG: Set platform field to %q\n", tt.platform)
			} else {
				t.Error("Failed to set platform field")
			}
			_, err := manager.GetBinary(context.Background())
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("GetBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func restoreWritePermissions(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk path %s: %w", path, err)
		}
		// Use 0700 for directories (rwx) and 0600 (rw) for files
		mode := os.FileMode(0600)
		if info != nil && info.IsDir() {
			mode = os.FileMode(0700)
		}
		if err := os.Chmod(path, mode); err != nil {
			return fmt.Errorf("failed to chmod %s: %w", path, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to restore write permissions: %w", err)
	}
	return nil
}

func setupTestDir(t *testing.T) string {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", "solc-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	t.Cleanup(func() {
		// Restore write permissions before cleanup
		if err := restoreWritePermissions(tmpDir); err != nil {
			t.Errorf("failed to restore write permissions: %v", err)
		}
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Errorf("failed to cleanup test directory: %v", err)
		}
	})
	return tmpDir
}

func setupNoWriteDir(t *testing.T, baseDir string) string {
	t.Helper()
	// Create the base directory with writable permissions initially
	noWriteDir := filepath.Join(baseDir, "no-write")
	if err := os.MkdirAll(noWriteDir, 0700); err != nil {
		t.Fatalf("Failed to create base dir: %v", err)
	}

	// Set both parent and test directory to read-only (0400)
	// This satisfies the linter requirement and ensures the test will fail
	// when BinaryManager tries to create subdirectories
	if err := os.Chmod(noWriteDir, 0400); err != nil {
		t.Fatalf("Failed to set directory permissions: %v", err)
	}

	fmt.Printf("DEBUG: setupNoWriteDir - Created directory %s with read-only permissions\n", noWriteDir)
	return noWriteDir
}

func TestDownloadAndVerify(t *testing.T) {
	tmpDir := setupTestDir(t)

	tests := []struct {
		name    string
		version string
		setup   func(t *testing.T) (*solc.BinaryManager, error)
		wantErr string
	}{
		{
			name:    "invalid permissions",
			version: "0.8.20",
			setup: func(t *testing.T) (*solc.BinaryManager, error) {
				t.Helper()
				noWriteDir := setupNoWriteDir(t, tmpDir)
				manager := solc.NewBinaryManagerWithDir("0.8.20", noWriteDir)
				managerValue := reflect.ValueOf(manager).Elem()
				if platformField := managerValue.FieldByName("Platform"); platformField.IsValid() && platformField.CanSet() {
					platformField.SetString("linux-amd64")
				}
				return manager, nil
			},
			wantErr: "failed to create cache directory: permission denied",
		},
		{
			name:    "invalid version format",
			version: "invalid.version",
			setup: func(t *testing.T) (*solc.BinaryManager, error) {
				t.Helper()
				manager := solc.NewBinaryManagerWithDir("invalid.version", tmpDir)
				managerValue := reflect.ValueOf(manager).Elem()
				if platformField := managerValue.FieldByName("Platform"); platformField.IsValid() && platformField.CanSet() {
					platformField.SetString("linux-amd64")
				}
				return manager, nil
			},
			wantErr: "invalid version format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			manager, err := tt.setup(t)
			if err != nil {
				t.Fatalf("Setup failed: %v", err)
			}
			_, err = manager.GetBinary(context.Background())
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("GetBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerifyChecksums(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "solc-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Errorf("failed to cleanup test directory: %v", err)
		}
	}()

	testContent := []byte("test content")
	testFile := filepath.Join(tmpDir, "test-binary")
	if err := os.WriteFile(testFile, testContent, 0600); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	sha256Sum := sha256.Sum256(testContent)
	keccak256Sum := crypto.Keccak256(testContent)

	tests := []struct {
		name      string
		sha256    string
		keccak256 string
		wantErr   bool
	}{
		{
			name:      "valid checksums",
			sha256:    hex.EncodeToString(sha256Sum[:]),
			keccak256: hex.EncodeToString(keccak256Sum),
			wantErr:   false,
		},
		{
			name:      "valid checksums with 0x prefix",
			sha256:    "0x" + hex.EncodeToString(sha256Sum[:]),
			keccak256: "0x" + hex.EncodeToString(keccak256Sum),
			wantErr:   false,
		},
		{
			name:      "invalid sha256",
			sha256:    "invalid",
			keccak256: hex.EncodeToString(keccak256Sum),
			wantErr:   true,
		},
		{
			name:      "invalid keccak256",
			sha256:    hex.EncodeToString(sha256Sum[:]),
			keccak256: "invalid",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			info := &solc.BinaryInfo{
				Sha256:    tt.sha256,
				Keccak256: tt.keccak256,
			}
			manager := setupTestBinaryManager(t)
			err := manager.VerifyChecksums(testFile, info)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyChecksums() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConcurrentDownloads(t *testing.T) {
	tmpDir := setupTestDir(t)

	// Create cache directories for both versions with proper permissions
	for _, version := range []string{"0.8.20", "0.8.19"} {
		cacheDir := filepath.Join(tmpDir, version, "linux-amd64")
		if err := os.MkdirAll(cacheDir, 0750); err != nil {
			t.Fatalf("Failed to create cache directory for version %s: %v", version, err)
		}
		t.Logf("DEBUG: Created cache directory: %s", cacheDir)
	}

	manager := solc.NewBinaryManagerWithDir("0.8.20", tmpDir)
	manager.UseFixedTempFileNames()

	// Download and cache the actual solc binaries for testing
	binaries := solc.DownloadAndCacheSolcBinary(t)

	// Reset download attempts counter before test
	solc.ResetDownloadAttempts()
	// Clean up any existing binary before test
	binaryPath := filepath.Join(tmpDir, "0.8.20", "linux-amd64", "solc-linux-amd64-v0.8.20")
	if err := os.RemoveAll(filepath.Dir(binaryPath)); err != nil {
		t.Fatalf("Failed to clean up test directory: %v", err)
	}
	const numConcurrent = 4
	var wg sync.WaitGroup
	errorCount := int32(0)

	// Create shorter context timeout for overall test since we expect immediate failures
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a failing transport that will fail all requests
	transport := &solc.FailingTransport{
		MaxSuccesses:     0,                  // All downloads should fail
		MockBinary:       binaries["0.8.20"], // Use version-specific binary
		Binaries:         binaries,           // Store all binaries for different versions
		SkipBackoffDelay: true,               // Skip backoff delay for faster test execution
	}
	transport.SuccessCount.Store(0) // Initialize atomic counter
	transport.CacheHits.Store(0)    // Initialize cache hits counter
	client := &http.Client{Transport: transport}
	manager.SetClient(client)
	// Start all goroutines at the same time
	ready := make(chan struct{})
	results := make(chan error, numConcurrent)
	completed := make(chan int, numConcurrent) // Track completed goroutines

	// Create two different URLs to test proper caching behavior
	urls := []string{
		"https://binaries.soliditylang.org/linux-amd64/solc-linux-amd64-v0.8.20+commit.a1b79de6",
		"https://binaries.soliditylang.org/linux-amd64/solc-linux-amd64-v0.8.19+commit.7dd6d404",
	}

	for i := 0; i < numConcurrent; i++ {
		wg.Add(1)
		go func(id int) {
			defer func() {
				wg.Done()
				completed <- id
				t.Logf("DEBUG: Goroutine %d completed", id)
			}()

			t.Logf("DEBUG: Goroutine %d waiting for start signal", id)
			<-ready
			t.Logf("DEBUG: Goroutine %d starting download", id)

			// Use different URLs for different goroutines
			urlIndex := id % len(urls)
			manager := solc.NewBinaryManagerWithDir(fmt.Sprintf("0.8.%d", 20-urlIndex), tmpDir)
			manager.SetClient(client)
			manager.UseFixedTempFileNames()

			// Single attempt with short timeout since we expect immediate failure
			attemptCtx, attemptCancel := context.WithTimeout(ctx, 2*time.Second)
			defer attemptCancel()

			_, err := manager.GetBinary(attemptCtx)
			if err == nil {
				t.Logf("DEBUG: Goroutine %d succeeded", id)
				// Let the transport track cache hits vs actual downloads
				results <- nil // Signal success
				return
			}

			// Only count connection refused errors as expected failures
			if strings.Contains(err.Error(), "connection refused") {
				atomic.AddInt32(&errorCount, 1)
				t.Logf("DEBUG: Goroutine %d failed with connection refused", id)
			} else {
				t.Logf("DEBUG: Goroutine %d failed with unexpected error: %v", id, err)
			}

			// Always send error to results channel
			results <- fmt.Errorf("goroutine %d: %w", id, err)
		}(i)
	}

	t.Log("DEBUG: Starting all goroutines")
	close(ready) // Signal all goroutines to start

	t.Log("DEBUG: Waiting for goroutines to complete")
	// Wait for all goroutines to complete or timeout
	completedCount := 0
	for completedCount < numConcurrent {
		select {
		case id := <-completed:
			completedCount++
			t.Logf("DEBUG: Received completion signal from goroutine %d (%d/%d done)",
				id, completedCount, numConcurrent)
		case <-ctx.Done():
			t.Fatalf("Test timed out waiting for goroutines (%d/%d completed)",
				completedCount, numConcurrent)
		}
	}

	close(results)
	t.Log("DEBUG: All goroutines completed, checking results")

	// Count successes and collect errors
	var errs []error
	for err := range results {
		if err != nil {
			errs = append(errs, err)
			t.Logf("Download error: %v", err)
		}
	}

	// Use transport's success count and cache hits for verification
	actualDownloads := int(transport.SuccessCount.Load())
	cacheHits := int(transport.CacheHits.Load())
	totalSuccesses := actualDownloads + cacheHits

	t.Logf("Results: %d actual downloads, %d cache hits, %d total successes, %d errors",
		actualDownloads, cacheHits, totalSuccesses, len(errs))

	// All downloads should fail
	if actualDownloads != 0 {
		t.Errorf("Expected 0 successful downloads, got %d actual downloads", actualDownloads)
	}

	if cacheHits != 0 {
		t.Errorf("Expected 0 cache hits, got %d", cacheHits)
	}

	// Verify all attempts resulted in errors
	if len(errs) != numConcurrent {
		t.Errorf("Expected %d errors (all downloads should fail), got %d", numConcurrent, len(errs))
	}

	// Verify all errors are connection refused
	for i, err := range errs {
		if !strings.Contains(err.Error(), "connection refused") {
			t.Errorf("Error %d: expected connection refused, got: %v", i, err)
		}
	}

	// Verify no binary files exist in cache directory
	files, err := os.ReadDir(filepath.Join(tmpDir, "0.8.20", "linux-amd64"))
	if err != nil {
		t.Fatalf("Failed to read cache directory: %v", err)
	}

	binaryCount := 0
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".tmp") {
			binaryCount++
		}
	}

	if binaryCount != 0 {
		t.Errorf("Expected no binary files, found %d", binaryCount)
	}
}

func TestNetworkErrorHandling(t *testing.T) {
	tmpDir := t.TempDir()
	manager := solc.NewBinaryManagerWithDir("0.8.20", tmpDir)

	// Create a client with the failing transport that always fails
	transport := &solc.FailingTransport{MaxSuccesses: 0, SkipBackoffDelay: true}
	client := &http.Client{Transport: transport}
	manager.SetClient(client)

	_, err := manager.GetBinary(context.Background())
	if err == nil {
		t.Error("Expected network error, got nil")
		return
	}

	if !strings.Contains(err.Error(), "connection refused") {
		t.Errorf("Expected connection refused error, got: %v", err)
	}
}

func TestBackoffWithJitter(t *testing.T) {
	manager := solc.NewBinaryManagerWithDir("0.8.20", t.TempDir())
	attempts := 0
	start := time.Now()

	for i := 0; i < 3; i++ {
		attempts++
		delay := manager.ApplyBackoffWithJitter(i)
		expectedBase := time.Duration(100*(1<<uint(i))) * time.Millisecond
		if delay < expectedBase || delay > expectedBase*2 {
			t.Errorf("Attempt %d: delay %v outside expected range [%v, %v]",
				i, delay, expectedBase, expectedBase*2)
		}
		time.Sleep(delay)
	}

	duration := time.Since(start)
	minDuration := 700 * time.Millisecond
	if duration < minDuration {
		t.Errorf("Total duration %v less than minimum expected %v", duration, minDuration)
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

func TestTempFileErrorHandling(t *testing.T) {
	tmpDir := setupTestDir(t)
	noWriteDir := setupNoWriteDir(t, tmpDir)
	manager := solc.NewBinaryManagerWithDir("0.8.20", noWriteDir)

	_, err := manager.GetBinary(context.Background())
	if err == nil || !strings.Contains(err.Error(), "permission denied") {
		t.Errorf("Expected permission denied error, got: %v", err)
	}
}
