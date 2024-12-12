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
	cacheDir := filepath.Join(os.Getenv("HOME"), ".cache", "solc")
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
	platform := m.GetPlatformDir()
	if platform == "invalid-platform" {
		return "", fmt.Errorf("failed to determine platform")
	}

	cacheDir := filepath.Join(m.cacheDir, m.version, platform)

	// Create cache directory if it doesn't exist
	if err := os.MkdirAll(cacheDir, dirPerms); err != nil {
		return "", fmt.Errorf("failed to create cache directory: %w", err)
	}

	// Check if binary exists in cache and is executable
	binaryPath := filepath.Join(cacheDir, fmt.Sprintf("solc-%s-v%s", platform, m.version))
	if info, err := os.Stat(binaryPath); err == nil {
		// Verify the binary is executable
		if info.Mode()&0111 == 0 {
			if err := os.Chmod(binaryPath, execPerms); err != nil {
				return "", fmt.Errorf("failed to make cached binary executable: %w", err)
			}
		}
		return binaryPath, nil
	}

	// Download binary with provided context
	return m.downloadBinary(ctx, platform, binaryPath)
}

// downloadBinary downloads the solc binary for the given platform.
func (m *BinaryManager) downloadBinary(ctx context.Context, platform, binaryPath string) (string, error) {
	// Platform validation already done in GetBinary
	binaryInfo, err := m.getBinaryInfo(ctx, platform)
	if err != nil {
		return "", err
	}

	// Download binary
	tmpFile := binaryPath + ".tmp"
	if err := m.downloadAndVerify(ctx, binaryInfo, tmpFile); err != nil {
		return "", err
	}

	// Move temporary file to final location and make it executable
	if err := os.Chmod(tmpFile, execPerms); err != nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
		return "", fmt.Errorf("failed to make binary executable: %w", err)
	}
	if err := os.Rename(tmpFile, binaryPath); err != nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
		return "", fmt.Errorf("failed to move binary to final location: %w", err)
	}

	return binaryPath, nil
}

// getBinaryInfo fetches and parses the list.json file to find matching version.
func (m *BinaryManager) getBinaryInfo(ctx context.Context, platform string) (*BinaryInfo, error) {
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

// setupCacheDir ensures the cache directory exists and is writable.
func (m *BinaryManager) setupCacheDir(cacheDir string) error {
	if err := os.MkdirAll(cacheDir, dirPerms); err != nil {
		return fmt.Errorf("failed to create cache directory")
	}

	// Test write permissions
	testFile := filepath.Join(cacheDir, ".write-test")
	if err := os.WriteFile(testFile, []byte{}, filePerms); err != nil {
		return fmt.Errorf("failed to create cache directory")
	}
	_ = os.Remove(testFile)
	return nil
}

// downloadFile downloads a file from the given URL and writes it to tmpFile.
func (m *BinaryManager) downloadFile(ctx context.Context, url, tmpFile string) error {
	//nolint:gosec // HTTP request to trusted domain is required for downloading solc binaries
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download binary: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download binary: HTTP %d", resp.StatusCode)
	}

	//nolint:gosec // File operations are required for storing downloaded binaries
	f, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, execPerms)
	if err != nil {
		return fmt.Errorf("failed to create cache directory")
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}()

	h := sha256.New()
	if _, err := io.Copy(io.MultiWriter(f, h), resp.Body); err != nil {
		return fmt.Errorf("failed to write binary: %w", err)
	}

	return nil
}

//nolint:gosec // File operations are required for storing downloaded binaries
// downloadAndVerify downloads the binary and verifies its checksums.
func (m *BinaryManager) downloadAndVerify(ctx context.Context, binaryInfo *BinaryInfo, tmpFile string) error {
	if err := m.setupCacheDir(filepath.Dir(tmpFile)); err != nil {
		return err
	}
	if err := m.downloadBinary(ctx, binaryInfo, tmpFile); err != nil {
		return err
	}
	return m.verifyAndInstall(tmpFile, binaryInfo)
}

// downloadBinary downloads the solc binary from the official repository.
func (m *BinaryManager) downloadBinary(ctx context.Context, binaryInfo *BinaryInfo, tmpFile string) error {
	platform := m.GetPlatformDir()
	if err := m.validatePlatform(platform); err != nil {
		return fmt.Errorf("failed to download list.json")
	}

	binaryURL := fmt.Sprintf("https://binaries.soliditylang.org/%s/%s", platform, binaryInfo.Path)
	if err := m.downloadFile(ctx, binaryURL, tmpFile); err != nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
		return err
	}
	return nil
}

// verifyAndInstall verifies the downloaded binary's checksums and installs it.
func (m *BinaryManager) verifyAndInstall(tmpFile string, binaryInfo *BinaryInfo) error {
	if err := m.VerifyChecksums(tmpFile, binaryInfo); err != nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil {
			fmt.Printf("failed to remove temporary file: %v\n", removeErr)
		}
		return err
	}
	return nil
}

func (m *BinaryManager) VerifyChecksums(tmpFile string, binaryInfo *BinaryInfo) error {
    // For test files, allow paths in /tmp
    cleanPath := filepath.Clean(tmpFile)
    if !strings.HasPrefix(cleanPath, m.cacheDir) && !strings.HasPrefix(cleanPath, os.TempDir()) {
        return fmt.Errorf("invalid binary path: attempted to access file outside allowed directories")
    }

    content, err := os.ReadFile(tmpFile)
    if err != nil {
        return fmt.Errorf("failed to read binary for verification: %w", err)
    }

    sha256Sum := sha256.Sum256(content)
    if hex.EncodeToString(sha256Sum[:]) != binaryInfo.Sha256 {
        return fmt.Errorf("sha256 checksum mismatch")
    }

    keccak256Sum := crypto.Keccak256(content)
    if hex.EncodeToString(keccak256Sum) != binaryInfo.Keccak256 {
        return fmt.Errorf("keccak256 checksum mismatch")
    }

    return nil
}
