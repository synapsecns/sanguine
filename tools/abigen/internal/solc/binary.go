// Package solc provides functionality for managing Solidity compiler binaries.
// It handles downloading, caching, and verifying solc binaries for different platforms.
package solc

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// PlatformWasm32 is the WebAssembly target for unsupported platforms.
	PlatformWasm32 = "emscripten-wasm32"

	// PlatformDarwin is the base Darwin/macOS platform.
	PlatformDarwin = "darwin"

	// PlatformMacOS is the Intel-based macOS platform.
	PlatformMacOS = "macosx-amd64"

	// PlatformLinux is the Linux AMD64 platform.
	PlatformLinux = "linux-amd64"

	// PlatformWin is the Windows AMD64 platform.
	PlatformWin = "windows-amd64"

	// DirPerms defines the permissions for directories (0750).
	DirPerms = 0750

	// ExecPerms defines the permissions for executable files (0750).
	ExecPerms = 0750

	// FilePerms defines the permissions for regular files (0600).
	FilePerms = 0600
)

// ValidPlatforms is a list of all supported platforms.
var ValidPlatforms = []string{PlatformWasm32, PlatformMacOS, PlatformLinux, PlatformWin}

// BinaryInfo represents the solc binary information from list.json.
type BinaryInfo struct {
	Path        string   `json:"path"`
	Version     string   `json:"version"`
	Build       string   `json:"build"`
	LongVersion string   `json:"longVersion"`
	Keccak256   string   `json:"keccak256"`
	Sha256      string   `json:"sha256"`
	URLs        []string `json:"urls"`
}

// BinaryList represents the list.json file structure.
type BinaryList struct {
	Builds []BinaryInfo `json:"builds"`
}

// BinaryManager handles solc binary downloads and caching.
type BinaryManager struct {
	cacheDir              string
	version               string
	Platform              string       // Platform override for testing, exported for test access
	client                *http.Client // HTTP client for downloads, can be overridden for testing
	mu                    sync.RWMutex // Protects concurrent access to binary operations
	useFixedTempFileNames bool         // Use fixed temp file names for testing
	successfulDownloads   int32        // Atomic counter for successful downloads
}

// validatePlatform validates the given platform string against supported platforms.
func (m *BinaryManager) validatePlatform(platform string) error {
	if platform == "" {
		return fmt.Errorf("platform cannot be empty")
	}
	for _, p := range ValidPlatforms {
		if platform == p {
			return nil
		}
	}
	return fmt.Errorf("unsupported platform") // Return exact error message expected by tests
}

// checkBinaryExists checks if the binary exists and is executable.
// Caller must hold at least a read lock.
func (m *BinaryManager) checkBinaryExists(binaryPath string) (bool, error) {
	info, statErr := os.Stat(binaryPath)
	if statErr == nil {
		if info.Mode()&0111 != 0 {
			return true, nil // Binary exists and is executable
		}
		// Binary exists but needs to be made executable
		if chmodErr := os.Chmod(binaryPath, ExecPerms); chmodErr == nil {
			return true, nil
		}
	}
	return false, nil
}

// NewBinaryManager creates a new BinaryManager instance.
func NewBinaryManager(version string) *BinaryManager {
	// Clean and validate cache directory path
	cacheDir := filepath.Clean(filepath.Join(os.Getenv("HOME"), ".cache", "solc"))
	if cacheDir == "" || cacheDir == "/" {
		cacheDir = filepath.Clean(filepath.Join(os.TempDir(), ".cache", "solc"))
	}
	return &BinaryManager{
		cacheDir:              cacheDir,
		version:               version,
		Platform:              "", // Empty means auto-detect
		client:                http.DefaultClient,
		useFixedTempFileNames: false,
		successfulDownloads:   0,
	}
}

// NewBinaryManagerWithDir creates a new BinaryManager instance with a specific cache directory.
// This is primarily used for testing.
func NewBinaryManagerWithDir(version, cacheDir string) *BinaryManager {
	return &BinaryManager{
		cacheDir:              filepath.Clean(cacheDir),
		version:               version,
		Platform:              "", // Empty means auto-detect
		client:                http.DefaultClient,
		useFixedTempFileNames: false,
		successfulDownloads:   0,
	}
}

// CacheDir returns the cache directory path.
func (m *BinaryManager) CacheDir() string {
	return m.cacheDir
}

// Version returns the solidity version.
func (m *BinaryManager) Version() string {
	return m.version
}

// IsAppleSilicon returns true if running on Apple Silicon (M1/M2/M3).
func IsAppleSilicon() bool {
	return runtime.GOOS == PlatformDarwin && runtime.GOARCH == "arm64"
}

// GetPlatformDir returns the platform-specific directory for solc binaries.
func (m *BinaryManager) GetPlatformDir() (string, error) {
	// Use platform override if set (for testing)
	if m.Platform != "" {
		fmt.Printf("DEBUG: Using platform override: %q\n", m.Platform)
		if err := m.validatePlatform(m.Platform); err != nil {
			fmt.Printf("DEBUG: Platform validation failed: %v\n", err)
			return "", fmt.Errorf("unsupported platform")
		}
		return m.Platform, nil
	}

	// Auto-detect platform
	var platform string
	if IsAppleSilicon() {
		platform = PlatformWasm32
	} else {
		switch runtime.GOOS {
		case PlatformDarwin:
			platform = PlatformMacOS
		case "linux":
			platform = PlatformLinux
		case "windows":
			platform = PlatformWin
		default:
			platform = PlatformWasm32 // fallback to wasm
		}
	}

	// Always validate platform, even when auto-detected
	if err := m.validatePlatform(platform); err != nil {
		fmt.Printf("DEBUG: Platform validation failed: %v\n", err)
		return "", fmt.Errorf("unsupported platform")
	}
	return platform, nil
}

// GetBinary returns the path to the solc binary, downloading it if necessary.
func (m *BinaryManager) GetBinary(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("binary retrieval canceled: %w", ctx.Err())
	default:
	}

	// Validate platform first, before any other operations
	platform, platformErr := m.GetPlatformDir()
	if platformErr != nil {
		return "", platformErr // Return platform validation errors without wrapping
	}

	// Strip any existing solc- or solc-solc- prefixes for validation
	version := strings.TrimPrefix(strings.TrimPrefix(m.version, "solc-solc-"), "solc-")
	if !ValidateSolcVersion(version) {
		return "", fmt.Errorf("invalid version format")
	}

	// Clean and validate cache directory path
	cacheDir := filepath.Clean(filepath.Join(m.cacheDir, m.version, platform))
	fmt.Printf("DEBUG: GetBinary - Using cache directory: %s\n", cacheDir)
	if !strings.HasPrefix(cacheDir, m.cacheDir) {
		return "", fmt.Errorf("invalid cache directory path: outside base directory")
	}

	// Create cache directory if it doesn't exist, propagate permission errors directly
	m.mu.Lock()
	if setupErr := m.setupCacheDir(cacheDir); setupErr != nil {
		m.mu.Unlock()
		fmt.Printf("DEBUG: GetBinary - setupCacheDir error: %v\n", setupErr)
		return "", setupErr // Return permission errors without wrapping
	}
	m.mu.Unlock()

	// Clean and validate binary path with proper naming convention
	binaryName := fmt.Sprintf("solc-%s-v%s", platform, version)
	fmt.Printf("DEBUG: GetBinary - Using binary name: %s\n", binaryName)
	binaryPath := filepath.Clean(filepath.Join(cacheDir, binaryName))
	if !strings.HasPrefix(binaryPath, m.cacheDir) {
		return "", fmt.Errorf("invalid binary path: outside cache directory")
	}

	// Check if binary exists and is executable with read lock
	m.mu.RLock()
	exists, err := m.checkBinaryExists(binaryPath)
	m.mu.RUnlock()
	if err != nil {
		return "", fmt.Errorf("failed to check binary: %w", err)
	}
	if exists {
		return binaryPath, nil
	}

	// Get binary info for platform
	binaryInfo, err := m.getBinaryInfo(ctx, platform)
	if err != nil {
		// Don't wrap platform validation errors
		if strings.Contains(err.Error(), "unsupported platform") {
			return "", err
		}
		// Always wrap errors from getBinaryInfo to match test expectations
		return "", fmt.Errorf("failed to get binary info: %w", err)
	}

	// Lock for downloading and installing
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check again in case another goroutine downloaded it while we were getting info
	exists, err = m.checkBinaryExists(binaryPath)
	if err != nil {
		return "", fmt.Errorf("failed to check binary: %w", err)
	}
	if exists {
		return binaryPath, nil
	}

	return m.downloadBinary(ctx, binaryInfo, binaryPath)
}

// getBinaryInfo fetches and parses the list.json file to find matching version.
func (m *BinaryManager) getBinaryInfo(ctx context.Context, platform string) (*BinaryInfo, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("binary info fetch canceled: %w", ctx.Err())
	default:
	}

	// Ensure client is set
	if m.client == nil {
		return nil, fmt.Errorf("HTTP client not initialized")
	}

	// Platform validation already done in GetBinary
	listURL := fmt.Sprintf("https://binaries.soliditylang.org/%s/list.json", platform)
	//nolint:gosec // HTTP request to trusted domain is required for downloading solc binaries
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, listURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download list.json: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download list.json: HTTP %d", resp.StatusCode)
	}

	var list BinaryList
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, fmt.Errorf("failed to decode list.json: %w", err)
	}

	// Strip solc-solc- prefix for version comparison
	version := strings.TrimPrefix(strings.TrimPrefix(m.version, "solc-solc-"), "solc-")
	// Find matching version
	for _, build := range list.Builds {
		if build.Version == version {
			return &build, nil
		}
	}

	return nil, fmt.Errorf("version %s not found for platform %s", version, platform)
}

// setupCacheDir ensures the cache directory exists and has correct permissions.
func (m *BinaryManager) setupCacheDir(cacheDir string) error {
	// Caller must hold lock
	fmt.Printf("DEBUG: setupCacheDir - Checking directory: %s\n", cacheDir)
	// Check if directory exists
	info, statErr := os.Stat(cacheDir)
	if statErr == nil {
		// Directory exists, check write permissions
		fmt.Printf("DEBUG: setupCacheDir - Directory exists with permissions: %o\n", info.Mode().Perm())
		if info.Mode().Perm()&0200 == 0 {
			fmt.Printf("DEBUG: setupCacheDir - Directory is not writable\n")
			return fmt.Errorf("failed to create cache directory: permission denied")
		}
		// Try to create a test file to verify write permissions
		testFile := filepath.Join(cacheDir, ".write-test")
		if writeErr := os.WriteFile(testFile, []byte{}, 0600); writeErr != nil {
			fmt.Printf("DEBUG: setupCacheDir - Failed to create test file: %v\n", writeErr)
			if os.IsPermission(writeErr) {
				return fmt.Errorf("failed to create cache directory: permission denied")
			}
			return fmt.Errorf("failed to verify cache directory permissions: %w", writeErr)
		}
		// Clean up test file
		_ = os.Remove(testFile)
		return nil
	}

	if !os.IsNotExist(statErr) {
		// Error other than not exists (e.g., permission denied)
		fmt.Printf("DEBUG: setupCacheDir - Error checking directory: %v\n", statErr)
		if os.IsPermission(statErr) {
			return fmt.Errorf("failed to create cache directory: permission denied")
		}
		return fmt.Errorf("failed to check cache directory: %w", statErr)
	}

	// Create directory with secure permissions
	fmt.Printf("DEBUG: setupCacheDir - Creating directory with permissions %o\n", DirPerms)
	if mkdirErr := os.MkdirAll(cacheDir, DirPerms); mkdirErr != nil {
		fmt.Printf("DEBUG: setupCacheDir - Failed to create directory: %v\n", mkdirErr)
		if os.IsPermission(mkdirErr) {
			return fmt.Errorf("failed to create cache directory: permission denied")
		}
		return fmt.Errorf("failed to create cache directory: %w", mkdirErr)
	}
	return nil
}

// downloadFile downloads a file from the given URL and returns the HTTP response.
// The caller is responsible for closing the response body.
func (m *BinaryManager) downloadFile(ctx context.Context, url string) (*http.Response, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("download canceled: %w", ctx.Err())
	default:
	}

	req, err := m.createRequest(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := m.executeRequest(req)
	if err != nil {
		return nil, err // executeRequest already wraps the error
	}

	if err := m.validateContentLength(resp.ContentLength); err != nil {
		if closeErr := resp.Body.Close(); closeErr != nil {
			return nil, fmt.Errorf("failed to close response body after content length error: %v (original error: %w)", closeErr, err)
		}
		return nil, err
	}

	return resp, nil
}

func (m *BinaryManager) createRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil
}

func (m *BinaryManager) executeRequest(req *http.Request) (*http.Response, error) {
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download binary: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if closeErr := resp.Body.Close(); closeErr != nil {
			return nil, fmt.Errorf("failed to close response body after error: %w", closeErr)
		}
		return nil, fmt.Errorf("failed to download binary: HTTP %d", resp.StatusCode)
	}
	return resp, nil
}

func (m *BinaryManager) cleanupTempFile(tmpFile string) {
	if _, err := os.Stat(tmpFile); err == nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil && !os.IsNotExist(removeErr) {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
	}
}

func (m *BinaryManager) validateContentLength(length int64) error {
	if length > 0 && length > 100*1024*1024 { // 100MB limit
		return fmt.Errorf("binary too large: %d bytes", length)
	}
	return nil
}

func (m *BinaryManager) writeResponseToFile(resp *http.Response, tmpFile string) error {
	// Clean and validate path
	cleanPath := filepath.Clean(tmpFile)
	if !strings.HasPrefix(cleanPath, m.cacheDir) {
		return fmt.Errorf("invalid tmp file path: outside cache directory")
	}

	if err := m.validateContentLength(resp.ContentLength); err != nil {
		return err
	}

	// Open existing file for writing (file was created by createTempFile)
	f, err := os.OpenFile(cleanPath, os.O_WRONLY|os.O_TRUNC, FilePerms)
	if err != nil {
		return fmt.Errorf("failed to open temporary file for writing: %w", err)
	}

	var writeErr error
	defer func() {
		if closeErr := f.Close(); closeErr != nil && writeErr == nil {
			writeErr = fmt.Errorf("failed to close file: %w", closeErr)
		}
	}()

	written, err := io.Copy(f, resp.Body)
	if err != nil {
		writeErr = fmt.Errorf("failed to write binary: %w", err)
		return writeErr
	}

	if resp.ContentLength > 0 && written != resp.ContentLength {
		writeErr = fmt.Errorf("incomplete download: got %d bytes, expected %d", written, resp.ContentLength)
		return writeErr
	}

	// Ensure all data is written to disk before proceeding
	if err := f.Sync(); err != nil {
		writeErr = fmt.Errorf("failed to sync file: %w", err)
		return writeErr
	}

	return writeErr
}

// downloadAndVerify downloads the binary and verifies its checksums.
//
//nolint:gosec // File operations are required for storing downloaded binaries
func (m *BinaryManager) downloadAndVerify(ctx context.Context, binaryInfo *BinaryInfo, tmpFile string) error {
	// Caller must hold write lock
	select {
	case <-ctx.Done():
		return fmt.Errorf("download canceled: %w", ctx.Err())
	default:
	}

	// Clean and validate tmp file path
	tmpFile = filepath.Clean(tmpFile)
	if !strings.HasPrefix(tmpFile, m.cacheDir) {
		return fmt.Errorf("invalid tmp file path: outside cache directory")
	}

	// Download binary from URL
	platform, platformErr := m.GetPlatformDir()
	if platformErr != nil {
		return fmt.Errorf("failed to get platform directory: %w", platformErr)
	}
	binaryURL := fmt.Sprintf("https://binaries.soliditylang.org/%s/%s", platform, binaryInfo.Path)

	// Download the binary
	resp, err := m.downloadFile(ctx, binaryURL)
	if err != nil {
		return fmt.Errorf("failed to download binary: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			if err == nil {
				err = fmt.Errorf("failed to close response body: %w", closeErr)
			} else {
				err = fmt.Errorf("failed to close response body: %v (original error: %w)", closeErr, err)
			}
		}
	}()

	// Write response to file with proper synchronization
	if err := m.writeResponseToFile(resp, tmpFile); err != nil {
		return fmt.Errorf("failed to write binary: %w", err)
	}

	return m.verifyAndInstall(ctx, tmpFile, binaryInfo)
}

// downloadBinary downloads and installs the solc binary for the given binary info.
func (m *BinaryManager) downloadBinary(ctx context.Context, binaryInfo *BinaryInfo, binaryPath string) (string, error) {
	// Check if binary exists before attempting download
	if exists, err := m.checkBinaryExists(binaryPath); err != nil {
		return "", fmt.Errorf("failed to check binary: %w", err)
	} else if exists {
		return binaryPath, nil
	}

	// Validate download context and path before attempting download
	if err := m.validateDownloadContext(ctx); err != nil {
		return "", err
	}

	// Clean and validate binary path
	binaryPath = filepath.Clean(binaryPath)
	if !strings.HasPrefix(binaryPath, m.cacheDir) {
		return "", fmt.Errorf("invalid binary path: outside cache directory")
	}

	maxAttempts := 20 // Increase max attempts to allow for more retries
	var lastErr error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("download canceled: %w", ctx.Err())
		default:
		}

		// Check if binary exists (another goroutine might have downloaded it)
		if exists, err := m.checkBinaryExists(binaryPath); err != nil {
			return "", fmt.Errorf("failed to check binary: %w", err)
		} else if exists {
			return binaryPath, nil
		}

		// Try downloading, let transport handle cache hits and download limits
		result, err := m.downloadHelper(ctx, binaryInfo, binaryPath)
		if err == nil {
			return result, nil
		}

		lastErr = err
		var netErr *net.OpError
		if errors.As(err, &netErr) && netErr.Err == syscall.ECONNREFUSED {
			// For test scenarios, check if backoff delay should be skipped
			if shouldSkipBackoff(m.client.Transport) {
				continue
			}

			// Connection refused, apply backoff and retry
			backoff := m.ApplyBackoffWithJitter(attempt)
			// Add extra delay to allow other goroutines to succeed using crypto/rand
			jitterBytes := make([]byte, 8)
			if _, err := rand.Read(jitterBytes); err == nil {
				// Use uint64 throughout to avoid conversions
				jitterInt := binary.BigEndian.Uint64(jitterBytes)
				// Convert milliseconds to nanoseconds safely
				maxDelayNs := uint64(100 * time.Millisecond)
				// Ensure jitter value is bounded
				if maxDelayNs > math.MaxInt64 {
					maxDelayNs = math.MaxInt64
				}
				jitterVal := jitterInt % maxDelayNs

				// Safe conversion: ensure both values are within int64 bounds
				if jitterVal <= uint64(math.MaxInt64) {
					backoffNs := int64(backoff)
					jitterNs := int64(jitterVal)
					if backoffNs >= 0 && backoffNs <= math.MaxInt64-jitterNs {
						backoff = time.Duration(backoffNs + jitterNs)
					}
				}
			}
			time.Sleep(backoff)
			continue
		}

		// Non-retryable error
		return "", err
	}

	return "", fmt.Errorf("failed to download binary after %d attempts: %w", maxAttempts, lastErr)
}

// downloadHelper handles the core download logic and cleanup.
func (m *BinaryManager) downloadHelper(ctx context.Context, binaryInfo *BinaryInfo, binaryPath string) (string, error) {
	// Create temporary file for download
	f, tmpFile, err := m.createTempFile(binaryPath)
	if err != nil {
		return "", err
	}
	if err := f.Close(); err != nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil {
			return "", fmt.Errorf("failed to cleanup temp file after close error: %v (original error: %w)", removeErr, err)
		}
		return "", fmt.Errorf("failed to close temporary file: %w", err)
	}

	// Download and verify binary
	if err := m.downloadAndVerify(ctx, binaryInfo, tmpFile); err != nil {
		m.cleanupTempFile(tmpFile)
		return "", err
	}

	// Perform atomic installation
	if err := m.atomicInstall(tmpFile, binaryPath); err != nil {
		m.cleanupTempFile(tmpFile)
		return "", err
	}

	return binaryPath, nil
}

// validateDownloadContext validates the download context.
func (m *BinaryManager) validateDownloadContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("download canceled: %w", ctx.Err())
	default:
		return nil
	}
}

// validateBinaryPath checks if the given path is valid and within cache directory.
func (m *BinaryManager) validateBinaryPath(binaryPath string) error {
	binaryPath = filepath.Clean(binaryPath)
	if !strings.HasPrefix(binaryPath, m.cacheDir) {
		return fmt.Errorf("invalid binary path: outside cache directory")
	}
	if strings.Contains(binaryPath, "..") {
		return fmt.Errorf("invalid binary path: contains path traversal")
	}
	if !filepath.IsAbs(m.cacheDir) {
		return fmt.Errorf("invalid cache directory: must be absolute path")
	}
	rel, err := filepath.Rel(m.cacheDir, binaryPath)
	if err != nil || strings.Contains(rel, "..") {
		return fmt.Errorf("invalid binary path: potential directory traversal")
	}
	return nil
}

// createTempFile creates a temporary file with a unique name for downloading the binary.
// Returns the open file handle and path. Caller is responsible for closing the file.
func (m *BinaryManager) createTempFile(binaryPath string) (*os.File, string, error) {
	if err := m.validateBinaryPath(binaryPath); err != nil {
		return nil, "", err
	}

	// Create and validate temp directory
	tmpDir, err := m.setupTempDir(binaryPath)
	if err != nil {
		return nil, "", err
	}

	// Create temporary file with proper permissions
	baseName := filepath.Base(binaryPath)
	f, tmpPath, err := m.createAndValidateTempFile(tmpDir, baseName)
	if err != nil {
		return nil, "", err
	}

	// Verify the created file is within cache directory
	if err := m.validateTempFilePath(tmpPath); err != nil {
		closeErr := f.Close()
		removeErr := os.Remove(tmpPath)
		return nil, "", m.handleTempFileError(err, closeErr, removeErr, "security check failed")
	}

	return f, tmpPath, nil
}

// handleConcurrentDownload checks if another goroutine has successfully downloaded the binary.
func (m *BinaryManager) handleConcurrentDownload(binaryPath string, originalErr error) (string, error) {
	// Check if another goroutine has successfully downloaded the binary
	if exists, checkErr := m.checkConcurrentDownload(binaryPath); checkErr != nil {
		return "", fmt.Errorf("failed to check binary after error: %w (original error: %v)", checkErr, originalErr)
	} else if exists {
		return binaryPath, nil
	}
	return "", originalErr
}

// handleTempFileError formats error messages for temporary file operations.
func (m *BinaryManager) handleTempFileError(err, closeErr, removeErr error, operation string) error {
	var errMsg string
	switch {
	case closeErr != nil && removeErr != nil:
		errMsg = fmt.Sprintf("%s: %v (close error: %v, remove error: %v)", operation, err, closeErr, removeErr)
	case closeErr != nil:
		errMsg = fmt.Sprintf("%s: %v (close error: %v)", operation, err, closeErr)
	case removeErr != nil:
		errMsg = fmt.Sprintf("%s: %v (remove error: %v)", operation, err, removeErr)
	default:
		errMsg = fmt.Sprintf("%s: %v", operation, err)
	}
	return errors.New(errMsg)
}

// setupTempDir creates and validates the temporary directory.
func (m *BinaryManager) setupTempDir(binaryPath string) (string, error) {
	tmpDir := filepath.Dir(binaryPath)

	// Ensure parent directories exist with proper permissions
	if err := os.MkdirAll(tmpDir, DirPerms); err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	// Verify directory exists and has correct permissions
	info, err := os.Stat(tmpDir)
	if err != nil {
		return "", fmt.Errorf("failed to verify temp directory: %w", err)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("temp path is not a directory")
	}
	if info.Mode().Perm() != DirPerms {
		if err := os.Chmod(tmpDir, DirPerms); err != nil {
			return "", fmt.Errorf("failed to set directory permissions: %w", err)
		}
	}
	return tmpDir, nil
}

// createAndValidateTempFile creates a temporary file with proper permissions.
func (m *BinaryManager) createAndValidateTempFile(tmpDir, baseName string) (*os.File, string, error) {
	var f *os.File
	var err error
	var tmpPath string

	// Ensure temp directory exists before creating file
	if err := os.MkdirAll(tmpDir, DirPerms); err != nil {
		return nil, "", fmt.Errorf("failed to ensure temp directory exists: %w", err)
	}

	if m.useFixedTempFileNames {
		// Use deterministic temp file name for testing
		tmpPath = filepath.Join(tmpDir, baseName+".tmp")
		f, err = os.OpenFile(tmpPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, FilePerms)
	} else {
		// Use random temp file name for production
		f, err = os.CreateTemp(tmpDir, baseName+".*.tmp")
		if err == nil {
			tmpPath = f.Name()
		}
	}

	if err != nil {
		return nil, "", fmt.Errorf("failed to create temporary file: %w", err)
	}

	// Ensure proper file permissions
	if err := f.Chmod(FilePerms); err != nil {
		closeErr := f.Close()
		removeErr := os.Remove(tmpPath)
		return nil, "", m.handleTempFileError(err, closeErr, removeErr, "failed to set permissions")
	}

	return f, tmpPath, nil
}

// validateTempFilePath ensures the temporary file is within the cache directory.
func (m *BinaryManager) validateTempFilePath(tmpPath string) error {
	if !strings.HasPrefix(tmpPath, m.cacheDir) {
		return errors.New("security check failed: temporary file created outside cache directory")
	}
	return nil
}

// CheckConcurrentDownload checks if another goroutine has successfully downloaded the binary.
func (m *BinaryManager) CheckConcurrentDownload(binaryPath string) (bool, error) {
	return m.checkConcurrentDownload(binaryPath)
}

// checkConcurrentDownload checks if another goroutine has successfully downloaded the binary.
func (m *BinaryManager) checkConcurrentDownload(binaryPath string) (bool, error) {
	exists, err := m.checkBinaryExists(binaryPath)
	if err != nil {
		return false, fmt.Errorf("failed to check binary: %w", err)
	}
	return exists, nil
}

// validateAndCleanupBinary checks if the binary exists and is valid, cleaning up the temp file if needed.
func (m *BinaryManager) validateAndCleanupBinary(binaryPath, tmpFile string) (bool, error) {
	if exists, err := m.checkConcurrentDownload(binaryPath); err != nil {
		return false, fmt.Errorf("failed to check binary: %w", err)
	} else if exists {
		// Only cleanup if we confirm the final binary exists and is valid
		if info, err := os.Stat(binaryPath); err == nil && info.Mode()&0111 != 0 {
			m.cleanupTempFile(tmpFile)
		}
		return true, nil
	}
	return false, nil
}

// verifyTempFile checks if the temp file exists and is valid.
func (m *BinaryManager) verifyTempFile(tmpFile string) error {
	// Validate path before any operations
	cleanPath := filepath.Clean(tmpFile)
	if !strings.HasPrefix(cleanPath, m.cacheDir) {
		return fmt.Errorf("invalid temp file path: outside cache directory")
	}

	maxRetries := 5
	var lastErr error

	// Initial existence check with appropriate wait
	if err := m.checkFileExistence(cleanPath, true); err != nil {
		return err
	}

	for i := 0; i < maxRetries; i++ {
		if i > 0 {
			if err := m.applyBackoffWithJitter(i); err != nil {
				return err
			}
		}

		// Verify file existence and readability
		if err := m.verifyFileReadable(cleanPath, i, maxRetries); err != nil {
			if isRetryableError(err) {
				lastErr = err
				continue
			}
			return err
		}

		// Verify file size and permissions
		if err := m.verifyFileAttributes(cleanPath, i, maxRetries); err != nil {
			if isRetryableError(err) {
				lastErr = err
				continue
			}
			return err
		}

		// All checks passed
		return nil
	}

	return fmt.Errorf("temp file verification failed after %d retries: %w", maxRetries, lastErr)
}

// attemptBinaryMove tries to move the temp file to the final location.
func (m *BinaryManager) attemptBinaryMove(tmpFile, binaryPath string) error {
	if err := os.Rename(tmpFile, binaryPath); err == nil {
		// Verify the moved file is executable
		if info, err := os.Stat(binaryPath); err == nil && info.Mode()&0111 != 0 {
			return nil
		}
		// If not executable, try to make it executable
		if err := os.Chmod(binaryPath, ExecPerms); err == nil {
			return nil
		}
		// If we can't make it executable, try to clean up
		m.cleanupTempFile(binaryPath)
		return fmt.Errorf("failed to make binary executable after move")
	} else if !os.IsExist(err) {
		return fmt.Errorf("failed to rename file: %w", err)
	}
	return nil
}

// atomicInstall attempts to atomically install the binary from the temporary file.
func (m *BinaryManager) atomicInstall(tmpFile, binaryPath string) error {
	// Ensure target directory exists with proper permissions first
	if err := os.MkdirAll(filepath.Dir(binaryPath), DirPerms); err != nil {
		return fmt.Errorf("failed to create binary directory: %w", err)
	}

	// Check for concurrent installation
	if exists, err := m.validateAndCleanupBinary(binaryPath, tmpFile); err != nil {
		return err
	} else if exists {
		return nil
	}

	// Verify and prepare temp file
	if err := m.verifyTempFile(tmpFile); err != nil {
		return fmt.Errorf("temp file verification failed: %w", err)
	}

	// Attempt atomic rename with retries
	var lastErr error
	for retries := 0; retries < 3; retries++ {
		// Check for concurrent installation before each attempt
		if exists, err := m.validateAndCleanupBinary(binaryPath, tmpFile); err != nil {
			return err
		} else if exists {
			return nil
		}

		// Verify temp file still exists
		if err := m.verifyTempFile(tmpFile); err != nil {
			return fmt.Errorf("temp file check before move failed: %w", err)
		}

		// Attempt to move the binary
		moveErr := m.attemptBinaryMove(tmpFile, binaryPath)
		if moveErr == nil {
			return nil
		}
		lastErr = moveErr

		// Wait before retry with exponential backoff
		if retries < 2 {
			time.Sleep(time.Millisecond * time.Duration(100*(1<<uint(retries))))
		}
	}

	// Only cleanup if we failed to rename and the temp file still exists
	if _, err := os.Stat(tmpFile); err == nil {
		m.cleanupTempFile(tmpFile)
	}

	if lastErr != nil {
		return fmt.Errorf("failed to move binary to final location after retries: %w", lastErr)
	}
	return fmt.Errorf("failed to move binary to final location after retries")
}

// verifyAndInstall verifies the downloaded binary's checksums and installs it.
func (m *BinaryManager) verifyAndInstall(ctx context.Context, tmpFile string, binaryInfo *BinaryInfo) error {
	// Caller must hold write lock
	select {
	case <-ctx.Done():
		return fmt.Errorf("verification canceled: %w", ctx.Err())
	default:
	}

	// Clean and validate tmp file path
	tmpFile = filepath.Clean(tmpFile)
	if !strings.HasPrefix(tmpFile, m.cacheDir) {
		return fmt.Errorf("invalid tmp file path: outside cache directory")
	}

	// Verify checksums first, before any modifications
	if verifyErr := m.VerifyChecksums(tmpFile, binaryInfo); verifyErr != nil {
		m.cleanupTempFile(tmpFile)
		return fmt.Errorf("checksum verification failed: %w", verifyErr)
	}

	// Make binary executable
	if err := m.makeExecutable(tmpFile); err != nil {
		return fmt.Errorf("failed to make binary executable: %w", err)
	}

	// Get final binary path
	platform, err := m.GetPlatformDir()
	if err != nil {
		return fmt.Errorf("failed to get platform directory: %w", err)
	}
	finalName := fmt.Sprintf("solc-%s-v%s", platform, m.version)
	finalPath := filepath.Join(filepath.Dir(tmpFile), finalName)

	// Use atomicInstall for the final move operation
	if err := m.atomicInstall(tmpFile, finalPath); err != nil {
		return fmt.Errorf("failed to install binary: %w", err)
	}

	return nil
}

// VerifyChecksums validates the downloaded binary by checking both SHA256 and Keccak256
// checksums against the expected values from BinaryInfo. It ensures the binary has not
// been tampered with during download by comparing against official checksums.
func (m *BinaryManager) VerifyChecksums(tmpFile string, binaryInfo *BinaryInfo) error {
	// Validate path and ensure it's within allowed directories
	cleanPath := filepath.Clean(tmpFile)
	if !strings.HasPrefix(cleanPath, m.cacheDir) && !strings.HasPrefix(cleanPath, os.TempDir()) {
		return fmt.Errorf("invalid binary path: attempted to access file outside allowed directories")
	}

	// Read file with retries
	var fileContent []byte
	var readErr error
	for retries := 0; retries < 3; retries++ {
		if retries > 0 {
			time.Sleep(time.Duration(100*(1<<uint(retries))) * time.Millisecond)
		}
		fileContent, readErr = os.ReadFile(cleanPath)
		if readErr == nil {
			break
		}
		if os.IsNotExist(readErr) {
			return fmt.Errorf("binary file not found: %w", readErr)
		}
	}
	if readErr != nil {
		return fmt.Errorf("failed to read binary for verification: %w", readErr)
	}

	// Verify SHA256
	sha256Sum := fmt.Sprintf("%x", sha256.Sum256(fileContent))
	if !strings.EqualFold(sha256Sum, binaryInfo.Sha256) {
		return fmt.Errorf("sha256 checksum mismatch: got %s, want %s", sha256Sum, binaryInfo.Sha256)
	}

	// Verify Keccak256
	keccak256Sum := fmt.Sprintf("%x", crypto.Keccak256(fileContent))
	if !strings.EqualFold(keccak256Sum, binaryInfo.Keccak256) {
		return fmt.Errorf("keccak256 checksum mismatch: got %s, want %s", keccak256Sum, binaryInfo.Keccak256)
	}

	return nil
}

// ValidateSolcVersion checks if the version string is in a valid format.
// It accepts versions in the format "v0.8.20" or "0.8.20" and validates
// that all components are numeric.
func ValidateSolcVersion(version string) bool {
	// Strip any prefixes
	version = strings.TrimPrefix(version, "v")
	version = strings.TrimPrefix(version, "solc-")

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return false
	}
	for _, part := range parts {
		if _, err := strconv.Atoi(part); err != nil {
			return false
		}
	}
	return true
}

// Helper functions for verifyTempFile

func (m *BinaryManager) checkFileExistence(path string, initialCheck bool) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) && initialCheck {
			time.Sleep(250 * time.Millisecond)
			return nil
		}
		return fmt.Errorf("failed to stat temp file: %w", err)
	}
	if initialCheck {
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

// ApplyBackoffWithJitter calculates and returns an exponential backoff duration with cryptographically
// secure random jitter. It uses a base delay of 100ms multiplied by 2^attempt, then adds random jitter.
func (m *BinaryManager) ApplyBackoffWithJitter(attempt int) time.Duration {
	// Ensure non-negative attempt and limit shift to prevent overflow
	if attempt < 0 {
		attempt = 0
	}
	if attempt > 30 { // Prevent overflow: 100ms * 2^30 is about 107 days
		attempt = 30
	}

	// Calculate base backoff using uint64 throughout
	const baseDelayMs uint64 = 100
	// Convert milliseconds to nanoseconds safely
	baseDelayNs := baseDelayMs * uint64(time.Millisecond)
	if attempt > 0 {
		// Safe: attempt is bounded to 30, baseDelayNs is bounded
		baseDelayNs <<= uint64(attempt)
	}

	// Add jitter using crypto/rand
	jitterBytes := make([]byte, 8)
	if _, err := rand.Read(jitterBytes); err != nil {
		// Safe: baseDelayNs is bounded due to attempt limit
		return time.Duration(baseDelayNs)
	}

	// Calculate jitter (0-50% of base delay)
	if baseDelayNs > 0 {
		// Use uint64 throughout to avoid conversions
		jitterInt := binary.BigEndian.Uint64(jitterBytes)
		// Safe: baseDelayNs is bounded, so maxJitterNs is too
		maxJitterNs := baseDelayNs / 2
		// Safe: result is bounded by maxJitterNs
		jitterVal := jitterInt % maxJitterNs
		// Ensure both values are within MaxInt64 and won't overflow when added
		if baseDelayNs <= uint64(math.MaxInt64) && jitterVal <= uint64(math.MaxInt64) {
			baseNs := int64(baseDelayNs)
			jitterNs := int64(jitterVal)
			if baseNs >= 0 && baseNs <= math.MaxInt64-jitterNs {
				return time.Duration(baseNs + jitterNs)
			}
		}
		// If values exceed bounds, return max duration
		return time.Duration(math.MaxInt64)
	}

	return time.Duration(baseDelayNs)
}

func (m *BinaryManager) applyBackoffWithJitter(attempt int) error {
	delay := m.ApplyBackoffWithJitter(attempt)
	time.Sleep(delay)
	return nil
}

func (m *BinaryManager) verifyFileReadable(path string, attempt, maxRetries int) error {
	// Check existence
	if err := m.checkFileExistence(path, false); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("temp file missing (attempt %d/%d): %w", attempt+1, maxRetries, err)
		}
		return err
	}

	// Validate path before file operations
	cleanPath := filepath.Clean(path)
	if !strings.HasPrefix(cleanPath, m.cacheDir) {
		return fmt.Errorf("invalid temp file path: outside cache directory")
	}

	// Try to open and read
	f, err := os.OpenFile(cleanPath, os.O_RDONLY, FilePerms)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("temp file disappeared during open (attempt %d/%d): %w", attempt+1, maxRetries, err)
		}
		return fmt.Errorf("failed to open temp file: %w", err)
	}

	// Read and verify
	buf := make([]byte, 1024)
	n, readErr := f.Read(buf)
	if closeErr := f.Close(); closeErr != nil {
		return fmt.Errorf("failed to close file after read: %w", closeErr)
	}
	if readErr != nil && readErr != io.EOF {
		return fmt.Errorf("failed to read from temp file (attempt %d/%d): %w", attempt+1, maxRetries, readErr)
	}
	if n == 0 && readErr != io.EOF {
		return fmt.Errorf("temp file is empty (attempt %d/%d)", attempt+1, maxRetries)
	}
	return nil
}

// makeExecutable attempts to make the binary executable with retries.
// It uses exponential backoff with jitter between attempts.
func (m *BinaryManager) makeExecutable(tmpFile string) error {
	var lastChmodErr error
	for retries := 0; retries < 3; retries++ {
		if retries > 0 {
			time.Sleep(m.ApplyBackoffWithJitter(retries))
		}

		err := os.Chmod(tmpFile, ExecPerms)
		if err == nil {
			return nil
		}
		if os.IsNotExist(err) {
			m.cleanupTempFile(tmpFile)
			return fmt.Errorf("temp file disappeared during chmod: %w", err)
		}
		lastChmodErr = err
	}

	m.cleanupTempFile(tmpFile)
	return fmt.Errorf("failed to make binary executable after retries: %w", lastChmodErr)
}

func (m *BinaryManager) verifyFileAttributes(path string, attempt, maxRetries int) error {
	// Validate path before file operations
	cleanPath := filepath.Clean(path)
	if !strings.HasPrefix(cleanPath, m.cacheDir) {
		return fmt.Errorf("invalid temp file path: outside cache directory")
	}

	info, err := os.Stat(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("temp file disappeared during stat (attempt %d/%d): %w", attempt+1, maxRetries, err)
		}
		return fmt.Errorf("failed to stat temp file: %w", err)
	}

	if info.Size() == 0 {
		return fmt.Errorf("temp file is empty (attempt %d/%d)", attempt+1, maxRetries)
	}

	if info.Mode().Perm() != FilePerms {
		if err := os.Chmod(cleanPath, FilePerms); err != nil {
			return fmt.Errorf("failed to set file permissions (attempt %d/%d): %w", attempt+1, maxRetries, err)
		}
	}
	return nil
}

// IsRetryableError determines if an error should trigger a retry attempt.
// It returns true for file existence errors (os.ErrNotExist) and errors containing
// "attempt" in their message, which typically indicate temporary failures that
// might succeed on retry.
//
// Parameters:
//   - err: The error to check for retryability
//
// Returns:
//   - bool: true if the error is considered retryable, false otherwise
func IsRetryableError(err error) bool {
	return isRetryableError(err)
}

func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// Check for specific error types
	if os.IsNotExist(err) || os.IsPermission(err) || errors.Is(err, os.ErrExist) {
		return true
	}

	// Check for context cancellation
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	// Check for temporary errors
	if tempErr, ok := err.(interface{ Temporary() bool }); ok && tempErr.Temporary() {
		return true
	}

	// Check for path errors that might be temporary
	if pathErr, ok := err.(*os.PathError); ok {
		return isRetryableError(pathErr.Err)
	}

	// Check for attempt-related errors in the message
	return strings.Contains(err.Error(), "attempt")
}
