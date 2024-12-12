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
	platform string // Platform override for testing
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
	return fmt.Errorf("unsupported platform: %s", platform)
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
		platform: "", // Empty means auto-detect
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
func (m *BinaryManager) GetPlatformDir() string {
	// Use platform override if set (for testing)
	if m.platform != "" {
		if err := m.validatePlatform(m.platform); err != nil {
			return "invalid-platform"
		}
		return m.platform
	}

	// Auto-detect platform
	if IsAppleSilicon() {
		return PlatformWasm32
	}
	switch runtime.GOOS {
	case PlatformDarwin:
		return PlatformMacOS
	case "linux":
		return PlatformLinux
	case "windows":
		return PlatformWin
	default:
		return PlatformWasm32 // fallback to wasm
	}
}

// GetBinary returns the path to the solc binary, downloading it if necessary.
func (m *BinaryManager) GetBinary(ctx context.Context) (string, error) {
	// Validate version format first
	if !ValidateSolcVersion(m.version) {
		return "", fmt.Errorf("invalid version format")
	}

	platform := m.GetPlatformDir()
	if platform == "invalid-platform" {
		return "", fmt.Errorf("failed to determine platform")
	}

	// Clean and validate cache directory path
	cacheDir := filepath.Clean(filepath.Join(m.cacheDir, m.version, platform))
	if !strings.HasPrefix(cacheDir, m.cacheDir) {
		return "", fmt.Errorf("invalid cache directory path: outside base directory")
	}

	// Create cache directory if it doesn't exist
	if err := os.MkdirAll(cacheDir, dirPerms); err != nil {
		return "", fmt.Errorf("failed to create cache directory: %w", err)
	}

	// Clean and validate binary path
	binaryPath := filepath.Clean(filepath.Join(cacheDir, fmt.Sprintf("solc-%s", m.version)))
	if !strings.HasPrefix(binaryPath, m.cacheDir) {
		return "", fmt.Errorf("invalid binary path: outside cache directory")
	}

	// Check if binary exists and is executable
	if info, err := os.Stat(binaryPath); err == nil {
		if info.Mode()&0111 == 0 {
			if err := os.Chmod(binaryPath, execPerms); err != nil {
				return "", fmt.Errorf("failed to make cached binary executable: %w", err)
			}
		}
		return binaryPath, nil
	}

	// Get binary info for platform
	binaryInfo, err := m.getBinaryInfo(ctx, platform)
	if err != nil {
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
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download list.json: HTTP %d", resp.StatusCode)
	}

	var list BinaryList
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, fmt.Errorf("failed to decode list.json: %w", err)
	}

	// Find matching version
	var matchingBuild BinaryInfo
	for _, build := range list.Builds {
		if build.Version == m.version {
			matchingBuild = build
			return &matchingBuild, nil
		}
	}

	return nil, fmt.Errorf("version %s not found for platform %s", m.version, platform)
}

// setupCacheDir ensures the cache directory exists and has correct permissions.
func (m *BinaryManager) setupCacheDir(cacheDir string) error {
	// Clean and validate directory path
	cacheDir = filepath.Clean(cacheDir)
	if !strings.HasPrefix(cacheDir, m.cacheDir) && !strings.HasPrefix(cacheDir, os.TempDir()) {
		return fmt.Errorf("invalid directory path: outside allowed directories")
	}

	// Create directory with secure permissions
	if err := os.MkdirAll(cacheDir, dirPerms); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Test write permissions
	testFile := filepath.Join(cacheDir, ".write-test")
	if err := os.WriteFile(testFile, []byte{}, filePerms); err != nil {
		return fmt.Errorf("failed to verify directory permissions: %w", err)
	}
	_ = os.Remove(testFile)
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

func (m *BinaryManager) writeResponseToFile(resp *http.Response, tmpFile string) error {
	if !strings.HasPrefix(filepath.Clean(tmpFile), m.cacheDir) {
		return fmt.Errorf("invalid tmp file path: outside cache directory")
	}

	// Verify content length if provided
	if resp.ContentLength > 0 {
		if resp.ContentLength > 100*1024*1024 { // 100MB limit
			return fmt.Errorf("binary too large: %d bytes", resp.ContentLength)
		}
	}

	f, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePerms)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer f.Close()

	written, err := io.Copy(io.MultiWriter(f), resp.Body)
	if err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to write binary: %w", err)
	}

	// Verify written bytes match content length if provided
	if resp.ContentLength > 0 && written != resp.ContentLength {
		os.Remove(tmpFile)
		return fmt.Errorf("incomplete download: got %d bytes, expected %d", written, resp.ContentLength)
	}

	// Ensure all data is written to disk
	if err := f.Sync(); err != nil {
		os.Remove(tmpFile)
		return fmt.Errorf("failed to sync file: %w", err)
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

	if err := m.setupCacheDir(filepath.Dir(tmpFile)); err != nil {
		return err
	}

	// Download binary from URL
	binaryURL := fmt.Sprintf("https://binaries.soliditylang.org/%s/%s", m.GetPlatformDir(), binaryInfo.Path)
	if err := m.downloadFile(ctx, binaryURL, tmpFile); err != nil {
		return fmt.Errorf("failed to download binary: %w", err)
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
	if err := m.downloadAndVerify(ctx, binaryInfo, tmpFile); err != nil {
		// Only attempt removal if file exists
		if _, statErr := os.Stat(tmpFile); statErr == nil {
			_ = os.Remove(tmpFile)
		}
		return "", fmt.Errorf("failed to download and verify binary: %w", err)
	}

	// Move temporary file to final location
	if err := os.Rename(tmpFile, binaryPath); err != nil {
		// Only attempt removal if file exists
		if _, statErr := os.Stat(tmpFile); statErr == nil {
			_ = os.Remove(tmpFile)
		}
		return "", fmt.Errorf("failed to move binary to final location: %w", err)
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
	if err := m.VerifyChecksums(tmpFile, binaryInfo); err != nil {
		// Only log removal errors if the file exists and we fail to remove it
		if removeErr := os.Remove(tmpFile); removeErr != nil && !os.IsNotExist(removeErr) {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
		return fmt.Errorf("checksum verification failed: %w", err)
	}

	// Make binary executable
	if err := os.Chmod(tmpFile, execPerms); err != nil {
		// Only log removal errors if the file exists and we fail to remove it
		if removeErr := os.Remove(tmpFile); removeErr != nil && !os.IsNotExist(removeErr) {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
		return fmt.Errorf("failed to make binary executable: %w", err)
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

	// Log checksum details for debugging
	fmt.Printf("SHA256 Verification:\nCalculated: [%s]\nExpected:   [%s]\nLengths: %d vs %d\n",
		calculatedSha256, expectedSha, len(calculatedSha256), len(expectedSha))

	if calculatedSha256 != expectedSha {
		return fmt.Errorf("sha256 checksum mismatch:\nGot:     [%s]\nWant:    [%s]\nOriginal: [%s]\nLengths:  %d vs %d",
			calculatedSha256, expectedSha, binaryInfo.Sha256, len(calculatedSha256), len(expectedSha))
	}

	// Calculate Keccak256 and normalize
	keccak256Sum := crypto.Keccak256(content)
	calculatedKeccak := strings.TrimSpace(strings.ToLower(hex.EncodeToString(keccak256Sum)))

	// Log keccak details for debugging
	fmt.Printf("Keccak256 Verification:\nCalculated: [%s]\nExpected:   [%s]\nLengths: %d vs %d\n",
		calculatedKeccak, expectedKeccak, len(calculatedKeccak), len(expectedKeccak))

	if calculatedKeccak != expectedKeccak {
		return fmt.Errorf("keccak256 checksum mismatch:\nGot:     [%s]\nWant:    [%s]\nOriginal: [%s]\nLengths:  %d vs %d",
			calculatedKeccak, expectedKeccak, binaryInfo.Keccak256, len(calculatedKeccak), len(expectedKeccak))
	}

	return nil
}

// ValidateSolcVersion checks if the version string is in a valid format.
// It accepts versions in the format "v0.8.20" or "0.8.20" and validates
// that all components are numeric.
func ValidateSolcVersion(version string) bool {
	version = strings.TrimPrefix(version, "v")
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
