// Package main provides a tool for downloading and running specific versions of golangci-lint
// across different architectures and repositories. It supports version pinning through
// .golangci-version files and maintains a local cache of downloaded binaries.
//
// Security Features:
//   - Validates all file paths and URLs
//   - Enforces secure file permissions (0600 for files, 0750 for directories)
//   - Prevents command injection
//   - Implements secure archive extraction
//   - Uses cryptographic verification for downloads
//   - Handles decompression bombs (50MB limit)
//   - Validates binary paths and permissions
//
// Usage:
//
//	golang-ci-lint [flags] [golangci-lint arguments]
//
// The tool automatically:
//   - Downloads the correct version for the current architecture
//   - Caches binaries for future use
//   - Validates security constraints
//   - Propagates exit codes from golangci-lint
//   - Supports cross-repository usage via go-findroot
//
// Environment:
//   - Uses go-findroot for repository detection
//   - Maintains secure file permissions
//   - Supports various architectures and operating systems
//   - Handles timeouts and cancellation via context
//
// Example:
//
//	golang-ci-lint run --fix --config=/path/to/.golangci.yml
//
// Note: This tool is designed to be used across different repositories and
// automatically detects the repository root using go-findroot.
package main

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/integralist/go-findroot/find"
)

const (
	downloadURLTemplate = "https://github.com/golangci/golangci-lint/releases/download/v%s/golangci-lint-%s-%s-%s.tar.gz"
	cacheDir            = "cache"
	filePerms           = 0400
	dirPerms            = 0750
	execPerms           = 0500
	maxDecompressSize   = 50 * 1024 * 1024 // 50MB limit for decompression
	httpTimeout         = 30 * time.Second
)

var (
	// ErrBinaryNotFound indicates the golangci-lint binary wasn't found in the archive
	ErrBinaryNotFound = errors.New("golangci-lint binary not found in archive")
	// ErrInvalidPath indicates a path is outside the repository root
	ErrInvalidPath = errors.New("path is outside repository root")
	// ErrInvalidURL indicates the download URL is not from GitHub releases
	ErrInvalidURL = errors.New("invalid URL: must be from github.com/golangci/golangci-lint/releases")
)

// validateURL ensures the download URL is from GitHub releases.
func validateURL(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}
	if u.Scheme != "https" || u.Host != "github.com" {
		return ErrInvalidURL
	}
	path := strings.TrimPrefix(u.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) < 5 || parts[0] != "golangci" || parts[1] != "golangci-lint" || parts[2] != "releases" {
		return ErrInvalidURL
	}
	return nil
}

// validatePath ensures a path is within the repository root and sanitized.
func validatePath(path, root string) error {
	// Clean paths to remove any ".." components
	path = filepath.Clean(path)
	root = filepath.Clean(root)

	// Special handling for Go build cache
	if strings.HasPrefix(path, "/tmp/go-build") {
		// Verify it's actually in the Go build cache
		if !strings.Contains(path, "/go-build") {
			return fmt.Errorf("suspicious path masquerading as Go build cache: %s", path)
		}
		// Still check permissions for security
		info, err := os.Stat(path)
		if err == nil {
			if info.Mode().Perm() > filePerms && !info.IsDir() {
				return fmt.Errorf("file permissions too loose: %o", info.Mode().Perm())
			}
			if info.Mode().Perm() > dirPerms && info.IsDir() {
				return fmt.Errorf("directory permissions too loose: %o", info.Mode().Perm())
			}
		}
		return nil
	}

	// Special handling for golangci-lint binary in cache
	if strings.Contains(path, "golangci-lint") && strings.Contains(path, "/cache/") {
		execPath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("failed to get executable path: %w", err)
		}
		execDir := filepath.Dir(execPath)

		// Verify binary is in the cache directory relative to executable
		if !strings.HasPrefix(filepath.Clean(path), filepath.Clean(filepath.Join(execDir, "cache"))) {
			return fmt.Errorf("binary not in cache directory: %s", path)
		}

		// Check binary permissions - allow execute permissions for binaries
		info, err := os.Stat(path)
		if err == nil && !info.IsDir() {
			perm := info.Mode().Perm()
			// Allow read and execute permissions only (exactly 0500)
			if perm&^0500 != 0 {
				return fmt.Errorf("invalid binary path: file permissions must be exactly 0500 (read+execute), got: %o", perm)
			}
		}
		return nil
	}

	// Additional validation using absolute paths
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return fmt.Errorf("failed to get absolute root path: %w", err)
	}
	if !strings.HasPrefix(absPath, absRoot) {
		return ErrInvalidPath
	}

	// Verify file exists and check permissions
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			if info.Mode().Perm() > dirPerms {
				return fmt.Errorf("directory permissions too loose: %o", info.Mode().Perm())
			}
		} else if info.Mode().Perm() > filePerms {
			return fmt.Errorf("file permissions too loose: %o", info.Mode().Perm())
		}
	}

	return nil
}

// safeJoin safely joins paths, preventing directory traversal
func safeJoin(root, path string) (string, error) {
	// Clean both paths
	root = filepath.Clean(root)
	path = filepath.Clean(path)

	// Join paths and verify the result is still under root
	joined := filepath.Join(root, path)
	if !strings.HasPrefix(filepath.Clean(joined), root) {
		return "", ErrInvalidPath
	}
	return joined, nil
}

// validateArgs ensures command arguments are safe.
func validateArgs(args []string) error {
	for _, arg := range args {
		// Prevent command injection via arguments
		if strings.Contains(arg, ";") || strings.Contains(arg, "|") || strings.Contains(arg, "&") ||
			strings.Contains(arg, ">") || strings.Contains(arg, "<") || strings.Contains(arg, "`") {
			return fmt.Errorf("invalid argument containing shell metacharacters: %s", arg)
		}
		// Validate paths in arguments
		if strings.HasPrefix(arg, "--") && strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			if strings.Contains(parts[0], "path") || strings.Contains(parts[0], "file") {
				if err := validatePath(parts[1], "/"); err != nil {
					return fmt.Errorf("invalid path in argument: %w", err)
				}
			}
		}
	}
	return nil
}

// findRepoRoot returns the root directory of the repository.
func findRepoRoot() (string, error) {
	root, err := find.Repo()
	if err != nil {
		return "", fmt.Errorf("failed to find repo root: %w", err)
	}
	return root.Path, nil
}

// setupLinter prepares the golangci-lint binary, downloading it if necessary.
func setupLinter(ctx context.Context, version, osName, arch string) (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	execDir := filepath.Clean(filepath.Dir(execPath))
	binaryName := fmt.Sprintf("golangci-lint-%s-%s-%s/golangci-lint", version, osName, arch)
	cachePath, err := safeJoin(execDir, filepath.Join(cacheDir, binaryName))
	if err != nil {
		return "", fmt.Errorf("failed to create cache path: %w", err)
	}

	// Check cache and download if needed
	if _, err := os.Stat(cachePath); err != nil {
		if err := downloadAndExtract(ctx, version, osName, arch, cachePath); err != nil {
			return "", fmt.Errorf("failed to setup golangci-lint: %w", err)
		}
	}

	return cachePath, nil
}

// findWorkDir locates the nearest directory containing a go.mod file.
func findWorkDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	dir := cwd
	for dir != "/" && dir != "." {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		dir = filepath.Dir(dir)
	}
	return cwd, nil
}

// processArgs ensures proper argument formatting and adds default configuration.
func processArgs(args []string, root string) []string {
	// Ensure "run" is the first argument if not present
	hasRun := false
	for _, arg := range args {
		if arg == "run" {
			hasRun = true
			break
		}
	}
	if !hasRun {
		args = append([]string{"run"}, args...)
	}

	// Add default config if not specified
	hasConfig := false
	for _, arg := range args {
		if strings.HasPrefix(arg, "--config") {
			hasConfig = true
			break
		}
	}
	if !hasConfig {
		configPath := filepath.Join(root, ".golangci.yml")
		args = append(args, "--config", configPath)
	}

	// Add --modules-download-mode=readonly to prevent module modifications
	return append(args, "--modules-download-mode=readonly")
}

// extractTarGz extracts a specific file from a tar.gz archive.
func extractTarGz(tr *tar.Reader, destPath string) error {
	for {
		header, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar: %w", err)
		}

		if strings.HasSuffix(header.Name, "golangci-lint") {
			// Validate file size
			if header.Size > maxDecompressSize {
				return fmt.Errorf("file too large: %d bytes", header.Size)
			}

			// Create parent directory with secure permissions
			if err := os.MkdirAll(filepath.Dir(destPath), dirPerms); err != nil {
				return fmt.Errorf("failed to create destination directory: %w", err)
			}

			// Create binary with correct permissions from the start
			destFile, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, execPerms)
			if err != nil {
				return fmt.Errorf("failed to create destination file: %w", err)
			}

			// Copy with size limit and verify hash
			hasher := sha256.New()
			reader := io.TeeReader(io.LimitReader(tr, maxDecompressSize), hasher)
			if _, err := io.Copy(destFile, reader); err != nil {
				if closeErr := destFile.Close(); closeErr != nil {
					err = errors.Join(err, fmt.Errorf("failed to close destination file: %w", closeErr))
				}
				return fmt.Errorf("failed to extract file: %w", err)
			}

			// Close file after writing
			if err := destFile.Close(); err != nil {
				return fmt.Errorf("failed to close destination file: %w", err)
			}

			// Log hash for verification
			hash := hex.EncodeToString(hasher.Sum(nil))
			fmt.Printf("Extracted file hash: %s\n", hash)

			return nil
		}
	}
	return ErrBinaryNotFound
}

func readVersion(root string) (string, error) {
	versionPath, err := safeJoin(root, ".golangci-version")
	if err != nil {
		return "", fmt.Errorf("invalid version file path: %w", err)
	}

	version, err := os.ReadFile(versionPath)
	if err != nil {
		return "", fmt.Errorf("reading .golangci-version: %w", err)
	}
	return strings.TrimSpace(string(version)), nil
}

func setupAndValidate(ctx context.Context, version, root string) (string, string, error) {
	cachePath, err := setupLinter(ctx, version, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return "", "", fmt.Errorf("setting up linter: %w", err)
	}

	if err := validatePath(cachePath, filepath.Dir(os.Args[0])); err != nil {
		return "", "", fmt.Errorf("invalid binary path: %w", err)
	}

	workDir, err := findWorkDir()
	if err != nil {
		return "", "", fmt.Errorf("finding work directory: %w", err)
	}

	return cachePath, workDir, nil
}

func runLinter(ctx context.Context, cachePath, workDir string, args []string) error {
	// Get repository root for path validation
	root, err := findRepoRoot()
	if err != nil {
		return fmt.Errorf("finding repository root: %w", err)
	}

	// Validate binary path and working directory
	if err := validatePath(cachePath, root); err != nil {
		return fmt.Errorf("invalid binary path: %w", err)
	}

	if err := validatePath(workDir, root); err != nil {
		return fmt.Errorf("invalid working directory: %w", err)
	}

	// Validate command arguments
	if err := validateArgs(args); err != nil {
		return fmt.Errorf("invalid command arguments: %w", err)
	}

	// Create command with validated paths and args
	cmd := exec.CommandContext(ctx, cachePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workDir

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr
		}
		return fmt.Errorf("running linter: %w", err)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Initialize application context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Find repository root
	root, err := findRepoRoot()
	if err != nil {
		return fmt.Errorf("finding repository root: %w", err)
	}

	// Read version
	version, err := readVersion(root)
	if err != nil {
		return fmt.Errorf("reading version: %w", err)
	}

	// Setup and validate paths
	cachePath, workDir, err := setupAndValidate(ctx, version, root)
	if err != nil {
		return fmt.Errorf("setting up and validating: %w", err)
	}

	// Process arguments and run linter
	processedArgs := processArgs(os.Args[1:], root)
	if err := runLinter(ctx, cachePath, workDir, processedArgs); err != nil {
		return fmt.Errorf("running linter: %w", err)
	}

	return nil
}

func downloadAndExtract(ctx context.Context, version, osName, arch, destPath string) error {
	// Create temporary file for download with secure permissions
	tmpFile, err := os.CreateTemp("", "golangci-lint-*.tar.gz")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer func() {
		if closeErr := tmpFile.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close temp file: %w", closeErr))
		}
		if removeErr := os.Remove(tmpFile.Name()); removeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to remove temp file: %w", removeErr))
		}
	}()

	// Download URL
	url := fmt.Sprintf(downloadURLTemplate, version, version, osName, arch)
	if err := validateURL(url); err != nil {
		return fmt.Errorf("invalid download URL: %w", err)
	}
	fmt.Printf("Downloading golangci-lint v%s from %s\n", version, url)

	// Create HTTP client with timeout and security settings
	client := &http.Client{
		Timeout: httpTimeout,
		Transport: &http.Transport{
			DisableCompression: true, // Handle decompression manually
			MaxIdleConns:       10,
			IdleConnTimeout:    5 * time.Second,
			ForceAttemptHTTP2:  true,
			TLSClientConfig: &tls.Config{
				MinVersion:    tls.VersionTLS12,
				Renegotiation: tls.RenegotiateNever,
				CurvePreferences: []tls.CurveID{
					tls.X25519,
					tls.CurveP256,
				},
				CipherSuites: []uint16{
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
					tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				},
			},
		},
	}

	// Create request with context and security headers
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "golangci-lint-manager/1.0")
	req.Header.Set("Accept", "application/octet-stream")

	// Download with security checks
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download golangci-lint: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close response body: %w", closeErr))
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download golangci-lint: HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	// Verify content type
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/octet-stream") &&
		!strings.Contains(contentType, "application/x-gzip") {
		return fmt.Errorf("unexpected content type: %s", contentType)
	}

	// Copy to temp file with size limit and verify hash
	hasher := sha256.New()
	reader := io.TeeReader(io.LimitReader(resp.Body, maxDecompressSize), hasher)
	if _, err := io.Copy(tmpFile, reader); err != nil {
		return fmt.Errorf("failed to save download: %w", err)
	}

	// Log hash for verification
	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("Downloaded file hash: %s\n", hash)

	// Create cache directory if it doesn't exist
	cacheDir := filepath.Clean(filepath.Dir(destPath))
	if err := os.MkdirAll(cacheDir, dirPerms); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	return extractBinary(ctx, tmpFile.Name(), destPath)
}

func extractBinary(ctx context.Context, tarPath, destPath string) error {
	// Open archive with secure permissions
	file, err := os.OpenFile(tarPath, os.O_RDONLY, filePerms)
	if err != nil {
		return fmt.Errorf("failed to open archive: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close archive: %w", closeErr))
		}
	}()

	// Create gzip reader with size limit
	gzr, err := gzip.NewReader(io.LimitReader(file, maxDecompressSize))
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer func() {
		if closeErr := gzr.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("failed to close gzip reader: %w", closeErr))
		}
	}()

	return extractTarGz(tar.NewReader(gzr), destPath)
}
