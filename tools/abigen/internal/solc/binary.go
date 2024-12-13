// Package solc provides functionality for managing Solidity compiler binaries.
// It handles downloading, caching, and verifying solc binaries for different platforms.
package solc

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

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

	// File permissions for cache directories and binaries.
	dirPerms  = 0750
	filePerms = 0600
	execPerms = 0755
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
	cacheDir string
	version  string
	Platform string // Platform override for testing, exported for test access
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

// NewBinaryManager creates a new BinaryManager instance.
func NewBinaryManager(version string) *BinaryManager {
	// Clean and validate cache directory path
	cacheDir := filepath.Clean(filepath.Join(os.Getenv("HOME"), ".cache", "solc"))
	if cacheDir == "" || cacheDir == "/" {
		cacheDir = filepath.Clean(filepath.Join(os.TempDir(), ".cache", "solc"))
	}
	return &BinaryManager{
		cacheDir: cacheDir,
		version:  version,
		Platform: "", // Empty means auto-detect
	}
}

// NewBinaryManagerWithDir creates a new BinaryManager instance with a specific cache directory.
// This is primarily used for testing.
func NewBinaryManagerWithDir(version, cacheDir string) *BinaryManager {
	return &BinaryManager{
		cacheDir: filepath.Clean(cacheDir),
		version:  version,
		Platform: "", // Empty means auto-detect
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
	if setupErr := m.setupCacheDir(cacheDir); setupErr != nil {
		fmt.Printf("DEBUG: GetBinary - setupCacheDir error: %v\n", setupErr)
		return "", setupErr // Return permission errors without wrapping
	}

	// Clean and validate binary path
	binaryPath := filepath.Clean(filepath.Join(cacheDir, fmt.Sprintf("solc-%s", m.version)))
	if !strings.HasPrefix(binaryPath, m.cacheDir) {
		return "", fmt.Errorf("invalid binary path: outside cache directory")
	}

	// Check if binary exists and is executable
	if info, statErr := os.Stat(binaryPath); statErr == nil {
		if info.Mode()&0111 == 0 {
			if chmodErr := os.Chmod(binaryPath, execPerms); chmodErr != nil {
				return "", fmt.Errorf("failed to make cached binary executable: %w", chmodErr)
			}
		}
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

	// Download binary if it doesn't exist
	return m.downloadBinary(ctx, binaryInfo, binaryPath)
}

// getBinaryInfo fetches and parses the list.json file to find matching version.
func (m *BinaryManager) getBinaryInfo(ctx context.Context, platform string) (*BinaryInfo, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("binary info fetch canceled: %w", ctx.Err())
	default:
	}

	// Platform validation already done in GetBinary
	listURL := fmt.Sprintf("https://binaries.soliditylang.org/%s/list.json", platform)
	//nolint:gosec // HTTP request to trusted domain is required for downloading solc binaries
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, listURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
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
	fmt.Printf("DEBUG: setupCacheDir - Creating directory with permissions %o\n", dirPerms)
	if mkdirErr := os.MkdirAll(cacheDir, dirPerms); mkdirErr != nil {
		fmt.Printf("DEBUG: setupCacheDir - Failed to create directory: %v\n", mkdirErr)
		if os.IsPermission(mkdirErr) {
			return fmt.Errorf("failed to create cache directory: permission denied")
		}
		return fmt.Errorf("failed to create cache directory: %w", mkdirErr)
	}
	return nil
}

// downloadFile downloads a file from the given URL and writes it to tmpFile.
func (m *BinaryManager) downloadFile(ctx context.Context, url, tmpFile string) error {
	if err := m.validateDownload(ctx, tmpFile); err != nil {
		return err
	}

	req, err := m.createRequest(ctx, url)
	if err != nil {
		return err
	}

	resp, err := m.executeRequest(req)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = fmt.Errorf("failed to close response body: %w", closeErr)
		}
	}()

	return m.writeResponseToFile(resp, tmpFile)
}

func (m *BinaryManager) validateDownload(ctx context.Context, tmpFile string) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("download canceled: %w", ctx.Err())
	default:
	}

	tmpFile = filepath.Clean(tmpFile)
	if !strings.HasPrefix(tmpFile, m.cacheDir) {
		return fmt.Errorf("invalid tmp file path: outside cache directory")
	}
	return nil
}

func (m *BinaryManager) createRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	return req, nil
}

func (m *BinaryManager) executeRequest(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
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
	if removeErr := os.Remove(tmpFile); removeErr != nil && !os.IsNotExist(removeErr) {
		fmt.Printf("failed to remove temporary file: %v\n", removeErr)
	}
}

func (m *BinaryManager) validateContentLength(length int64) error {
	if length > 0 && length > 100*1024*1024 { // 100MB limit
		return fmt.Errorf("binary too large: %d bytes", length)
	}
	return nil
}

func (m *BinaryManager) writeResponseToFile(resp *http.Response, tmpFile string) error {
	if !strings.HasPrefix(filepath.Clean(tmpFile), m.cacheDir) {
		return fmt.Errorf("invalid tmp file path: outside cache directory")
	}

	if err := m.validateContentLength(resp.ContentLength); err != nil {
		return err
	}

	f, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePerms)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	var closeErr error
	defer func() {
		if cerr := f.Close(); cerr != nil {
			closeErr = cerr
		}
	}()

	written, err := io.Copy(io.MultiWriter(f), resp.Body)
	if err != nil {
		m.cleanupTempFile(tmpFile)
		return fmt.Errorf("failed to write binary: %w", err)
	}

	if resp.ContentLength > 0 && written != resp.ContentLength {
		m.cleanupTempFile(tmpFile)
		return fmt.Errorf("incomplete download: got %d bytes, expected %d", written, resp.ContentLength)
	}

	if err := f.Sync(); err != nil {
		m.cleanupTempFile(tmpFile)
		return fmt.Errorf("failed to sync file: %w", err)
	}

	if closeErr != nil {
		m.cleanupTempFile(tmpFile)
		return fmt.Errorf("failed to close file: %w", closeErr)
	}

	return nil
}

// downloadAndVerify downloads the binary and verifies its checksums.
//
//nolint:gosec // File operations are required for storing downloaded binaries
func (m *BinaryManager) downloadAndVerify(ctx context.Context, binaryInfo *BinaryInfo, tmpFile string) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("download canceled: %w", ctx.Err())
	default:
	}

	if setupErr := m.setupCacheDir(filepath.Dir(tmpFile)); setupErr != nil {
		return setupErr
	}

	// Download binary from URL
	platform, platformErr := m.GetPlatformDir()
	if platformErr != nil {
		return fmt.Errorf("failed to get platform directory: %w", platformErr)
	}
	binaryURL := fmt.Sprintf("https://binaries.soliditylang.org/%s/%s", platform, binaryInfo.Path)
	if downloadErr := m.downloadFile(ctx, binaryURL, tmpFile); downloadErr != nil {
		return fmt.Errorf("failed to download binary: %w", downloadErr)
	}

	return m.verifyAndInstall(ctx, tmpFile, binaryInfo)
}

// downloadBinary downloads and installs the solc binary for the given binary info.
func (m *BinaryManager) downloadBinary(ctx context.Context, binaryInfo *BinaryInfo, binaryPath string) (string, error) {
	// Clean and validate binary path
	binaryPath = filepath.Clean(binaryPath)
	if !strings.HasPrefix(binaryPath, m.cacheDir) {
		return "", fmt.Errorf("invalid binary path: outside cache directory")
	}

	// Download and verify binary
	tmpFile := binaryPath + ".tmp"
	if downloadErr := m.downloadAndVerify(ctx, binaryInfo, tmpFile); downloadErr != nil {
		// Only attempt removal if file exists
		if _, statErr := os.Stat(tmpFile); statErr == nil {
			_ = os.Remove(tmpFile)
		}
		return "", fmt.Errorf("failed to download and verify binary: %w", downloadErr)
	}

	// Move temporary file to final location
	if renameErr := os.Rename(tmpFile, binaryPath); renameErr != nil {
		// Only attempt removal if file exists
		if _, statErr := os.Stat(tmpFile); statErr == nil {
			_ = os.Remove(tmpFile)
		}
		return "", fmt.Errorf("failed to move binary to final location: %w", renameErr)
	}

	return binaryPath, nil
}

// verifyAndInstall verifies the downloaded binary's checksums and installs it.
func (m *BinaryManager) verifyAndInstall(ctx context.Context, tmpFile string, binaryInfo *BinaryInfo) error {
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

	// Verify checksums
	if verifyErr := m.VerifyChecksums(tmpFile, binaryInfo); verifyErr != nil {
		// Only log removal errors if the file exists and we fail to remove it
		if removeErr := os.Remove(tmpFile); removeErr != nil && !os.IsNotExist(removeErr) {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
		return fmt.Errorf("checksum verification failed: %w", verifyErr)
	}

	// Make binary executable
	if chmodErr := os.Chmod(tmpFile, execPerms); chmodErr != nil {
		// Only log removal errors if the file exists and we fail to remove it
		if removeErr := os.Remove(tmpFile); removeErr != nil && !os.IsNotExist(removeErr) {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
		return fmt.Errorf("failed to make binary executable: %w", chmodErr)
	}

	return nil
}

// VerifyChecksums validates the downloaded binary by checking both SHA256 and Keccak256
// checksums against the expected values from BinaryInfo. It ensures the binary has not
// been tampered with during download by comparing against official checksums.
func (m *BinaryManager) VerifyChecksums(tmpFile string, binaryInfo *BinaryInfo) error {
	cleanPath := filepath.Clean(tmpFile)
	if !strings.HasPrefix(cleanPath, m.cacheDir) && !strings.HasPrefix(cleanPath, os.TempDir()) {
		return fmt.Errorf("invalid binary path: attempted to access file outside allowed directories")
	}

	content, err := os.ReadFile(tmpFile)
	if err != nil {
		return fmt.Errorf("failed to read binary for verification: %w", err)
	}

	// Normalize expected checksums: trim spaces, remove 0x prefix, convert to lowercase
	expectedSha := strings.TrimSpace(strings.ToLower(strings.TrimPrefix(binaryInfo.Sha256, "0x")))
	expectedKeccak := strings.TrimSpace(strings.ToLower(strings.TrimPrefix(binaryInfo.Keccak256, "0x")))

	// Calculate SHA256 and normalize
	sha256Sum := sha256.Sum256(content)
	calculatedSha256 := strings.TrimSpace(strings.ToLower(hex.EncodeToString(sha256Sum[:])))

	// Compare SHA256 checksums with detailed error message
	if calculatedSha256 != expectedSha {
		return fmt.Errorf("sha256 checksum mismatch: got %s, want %s", calculatedSha256, expectedSha)
	}

	// Calculate Keccak256 and normalize
	keccak256Sum := crypto.Keccak256(content)
	calculatedKeccak := strings.TrimSpace(strings.ToLower(hex.EncodeToString(keccak256Sum)))

	// Compare Keccak256 checksums with detailed error message
	if calculatedKeccak != expectedKeccak {
		// Convert both checksums to byte arrays for detailed comparison
		calcBytes, _ := hex.DecodeString(calculatedKeccak)
		expBytes, _ := hex.DecodeString(expectedKeccak)
		return fmt.Errorf("keccak256 checksum mismatch: got %x, want %x (lengths: calc=%d, exp=%d)",
			calcBytes, expBytes, len(calculatedKeccak), len(expectedKeccak))
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
