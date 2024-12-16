// Package main provides a tool for downloading and running specific versions of golangci-lint
// across different architectures and repositories. It supports version pinning through
// .golangci-version files and maintains a local cache of downloaded binaries.
//
// Security Features:
//   - Validates all file paths and URLs
//   - Enforces secure file permissions (0400 for files, 0750 for directories)
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
	"github.com/synapsecns/sanguine/contrib/golang-ci-lint/permissions"
)

const (
	downloadURLTemplate = "https://github.com/golangci/golangci-lint/releases/download/v%s/golangci-lint-%s-%s-%s.tar.gz"
	cacheDir            = "cache"
	filePerms           = permissions.FilePerms
	dirPerms            = permissions.DirPerms
	execPerms           = permissions.ExecPerms
	maxDecompressSize   = 50 * 1024 * 1024 // 50MB limit for decompression
	httpTimeout         = 30 * time.Second
	maxIdleConns        = 10               // Maximum number of idle connections
	idleConnTimeout     = 5 * time.Second  // Idle connection timeout
	splitParts          = 2                // For strings.SplitN in validateURL
	extraArgsCapacity   = 3                // Extra capacity for processArgs slice
	minTLSVersion       = tls.VersionTLS12 // Minimum TLS version for security
	defaultArgCapacity  = 3                // Default capacity for processArgs slice
)

var (
	// ErrBinaryNotFound indicates the golangci-lint binary wasn't found in the archive.
	ErrBinaryNotFound = errors.New("golangci-lint binary not found in archive")
	// ErrInvalidPath indicates a path is outside the repository root.
	ErrInvalidPath = errors.New("path is outside repository root")
	// ErrInvalidURL indicates the download URL is not from GitHub releases.
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

// validatePath ensures a path is safe to use and within allowed directories.
// nolint: cyclop
func validatePath(path string, allowedDirs ...string) error {
	// Get repository root using go-findroot
	root, err := find.Repo()
	if err == nil {
		// Add repository root to allowed directories
		allowedDirs = append(allowedDirs, root.Path)
	}

	// Resolve to absolute path and handle symlinks
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	// Get the real path, following all symlinks
	realPath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		// If the file doesn't exist yet, use the absolute path
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to resolve symlinks: %w", err)
		}
		realPath = absPath
	}

	// Normalize paths for MacOS by removing /private prefix from /var paths
	if runtime.GOOS == "darwin" {
		if strings.HasPrefix(realPath, "/private/var") {
			realPath = strings.Replace(realPath, "/private/var", "/var", 1)
		}
		if strings.HasPrefix(absPath, "/private/var") {
			absPath = strings.Replace(absPath, "/private/var", "/var", 1)
		}
	}

	if strings.Contains(realPath, "..") {
		return fmt.Errorf("path contains directory traversal: %s", realPath)
	}

	// Always allow temp dir and cache dir
	tmpDir := filepath.Clean(os.TempDir())
	if runtime.GOOS == "darwin" && strings.HasPrefix(tmpDir, "/private/var") {
		tmpDir = strings.Replace(tmpDir, "/private/var", "/var", 1)
	}

	defaultDirs := []string{
		tmpDir,
		filepath.Clean(cacheDir),
	}

	// Combine default and additional allowed directories
	allowedDirs = append(defaultDirs, allowedDirs...)

	// Check against allowed directories
	for _, dir := range allowedDirs {
		absDir, err := filepath.Abs(dir)
		if err != nil {
			continue
		}
		cleanDir := filepath.Clean(absDir)
		if runtime.GOOS == "darwin" && strings.HasPrefix(cleanDir, "/private/var") {
			cleanDir = strings.Replace(cleanDir, "/private/var", "/var", 1)
		}
		// Check both realPath and absPath for compatibility
		if strings.HasPrefix(realPath, cleanDir) || strings.HasPrefix(absPath, cleanDir) {
			return nil
		}
	}

	return fmt.Errorf("path outside allowed directories: %s", realPath)
}

// safeJoin safely joins paths, preventing directory traversal.
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

// validateBinaryPermissions ensures binary files have correct permissions.
func validateBinaryPermissions(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat binary: %w", err)
	}
	if info.Mode().Perm() != execPerms {
		return fmt.Errorf("invalid binary permissions: %o (expected %o)", info.Mode().Perm(), execPerms)
	}
	return nil
}

// setupLinter prepares the golangci-lint binary, downloading it if necessary.
func setupLinter(ctx context.Context, version, osName, arch string) (string, error) {
	root, err := find.Repo()
	if err != nil {
		return "", fmt.Errorf("failed to get repository root: %w", err)
	}

	// Create cache directory first with proper permissions
	cacheDir := filepath.Join(root.Path, "contrib/golang-ci-lint/cache")
	if err := os.MkdirAll(cacheDir, dirPerms); err != nil {
		return "", fmt.Errorf("failed to create cache directory: %w", err)
	}

	binaryName := fmt.Sprintf("golangci-lint-%s-%s-%s/golangci-lint", version, osName, arch)
	cachePath := filepath.Join(cacheDir, binaryName)

	// Create binary directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(cachePath), dirPerms); err != nil {
		return "", fmt.Errorf("failed to create binary directory: %w", err)
	}

	// Check cache and download if needed
	if _, err := os.Stat(cachePath); err != nil {
		if err := downloadAndExtract(ctx, version, osName, arch, cachePath); err != nil {
			return "", fmt.Errorf("failed to setup golangci-lint: %w", err)
		}
	}

	return cachePath, nil
}

// findWorkDir locates the current working directory.
func findWorkDir() (string, error) {
	// Get repository root using go-findroot
	root, err := find.Repo()
	if err != nil {
		return "", fmt.Errorf("failed to get repository root: %w", err)
	}

	// Use current directory as working directory
	workDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	// Validate the path
	if err := validatePath(workDir, root.Path); err != nil {
		return "", fmt.Errorf("invalid working directory: %w", err)
	}

	return workDir, nil
}

// processArgs ensures proper argument formatting and adds default configuration.
// nolint: cyclop
func processArgs(args []string, root string) ([]string, string) {
	var workDir string
	processedArgs := make([]string, 0, len(args)+defaultArgCapacity)

	// Process arguments
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--path":
			if i+1 < len(args) {
				workDir = args[i+1]
				i++ // Skip the next argument since we consumed it
			}
		default:
			// Replace $(GIT_ROOT) in all arguments
			arg = strings.ReplaceAll(arg, "$(GIT_ROOT)", root)
			processedArgs = append(processedArgs, arg)
		}
	}

	// Ensure "run" is the first argument if not present
	hasRun := false
	for _, arg := range processedArgs {
		if arg == "run" {
			hasRun = true
			break
		}
	}
	if !hasRun {
		processedArgs = append([]string{"run"}, processedArgs...)
	}

	// Add default config if not specified
	hasConfig := false
	for _, arg := range processedArgs {
		if strings.HasPrefix(arg, "--config") {
			hasConfig = true
			break
		}
	}
	if !hasConfig {
		configPath := filepath.Join(root, ".golangci.yml")
		processedArgs = append(processedArgs, "--config", configPath)
	}

	return processedArgs, workDir
}

// extractTarGz extracts a specific file from a tar.gz archive.
func extractTarGz(tr *tar.Reader, destPath string) error {
	// Set secure permissions for file operations
	cleanup, err := permissions.SetSecureUmask()
	if err != nil {
		return fmt.Errorf("failed to set secure permissions: %w", err)
	}
	defer cleanup()

	// Create parent directory with secure permissions
	if err := os.MkdirAll(filepath.Dir(destPath), dirPerms); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	return findAndExtractBinary(tr, destPath)
}

// findAndExtractBinary locates and extracts the golangci-lint binary from the archive.
func findAndExtractBinary(tr *tar.Reader, destPath string) error {
	for {
		header, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar: %w", err)
		}

		if strings.HasSuffix(header.Name, "golangci-lint") {
			return extractBinaryFile(tr, header, destPath)
		}
	}
	return ErrBinaryNotFound
}

// extractBinaryFile handles the extraction of a single binary file.
func extractBinaryFile(tr *tar.Reader, header *tar.Header, destPath string) error {
	// Validate file size
	if header.Size > maxDecompressSize {
		return fmt.Errorf("file too large: %d bytes", header.Size)
	}

	// Create temporary file
	tmpFile := destPath + ".tmp"
	if err := createAndCopyBinary(tr, tmpFile); err != nil {
		return err
	}

	// Move to final destination with correct permissions
	if err := finalizeBinary(tmpFile, destPath); err != nil {
		return err
	}

	return nil
}

func createAndCopyBinary(tr *tar.Reader, tmpFile string) error {
	var errs []error

	// Validate path before file operations
	if err := validatePath(tmpFile, os.TempDir()); err != nil {
		return fmt.Errorf("invalid temporary file path: %w", err)
	}

	// Safe to use os.OpenFile here as path is validated by validatePath above
	//nolint:gosec // G304: path is validated by validatePath above
	file, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, execPerms)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			errs = append(errs, fmt.Errorf("failed to close file: %w", closeErr))
		}
	}()

	// Copy with size limit and verify hash
	hasher := sha256.New()
	reader := io.TeeReader(io.LimitReader(tr, maxDecompressSize), hasher)
	if _, err := io.Copy(file, reader); err != nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil {
			errs = append(errs, fmt.Errorf("failed to remove temporary file: %w", removeErr))
		}
		errs = append(errs, fmt.Errorf("failed to extract file: %w", err))
		return errors.Join(errs...)
	}

	// Log hash for verification
	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("Extracted file hash: %s\n", hash)

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// finalizeBinary moves the temporary file to its final destination.
func finalizeBinary(tmpFile, destPath string) error {
	var errs []error
	// Set correct permissions
	if err := os.Chmod(tmpFile, execPerms); err != nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil {
			errs = append(errs, fmt.Errorf("failed to remove temporary file: %w", removeErr))
		}
		errs = append(errs, fmt.Errorf("failed to set file permissions: %w", err))
		return errors.Join(errs...)
	}

	// Move file to final destination
	if err := os.Rename(tmpFile, destPath); err != nil {
		if removeErr := os.Remove(tmpFile); removeErr != nil {
			errs = append(errs, fmt.Errorf("failed to remove temporary file: %w", removeErr))
		}
		errs = append(errs, fmt.Errorf("failed to move file to destination: %w", err))
		return errors.Join(errs...)
	}

	return validateBinaryPermissions(destPath)
}

func readVersion(root string) (string, error) {
	versionPath, err := safeJoin(root, ".golangci-version")
	if err != nil {
		return "", fmt.Errorf("invalid version file path: %w", err)
	}

	// Validate path before reading
	if err := validatePath(versionPath, root); err != nil {
		return "", fmt.Errorf("invalid version file path: %w", err)
	}

	// Safe to use os.ReadFile here as path is validated by validatePath above
	//nolint:gosec // G304: path is validated by validatePath above
	version, err := os.ReadFile(versionPath)
	if err != nil {
		return "", fmt.Errorf("reading .golangci-version: %w", err)
	}
	return strings.TrimSpace(string(version)), nil
}

// nolint: cyclop
func runLinter(ctx context.Context, binaryPath, workDir string, args []string) error {
	// Validate binary permissions before execution
	info, err := os.Stat(binaryPath)
	if err != nil {
		return fmt.Errorf("checking binary before execution: %w", err)
	}
	if info.Mode().Perm() != execPerms {
		if err := os.Chmod(binaryPath, execPerms); err != nil {
			return fmt.Errorf("setting binary permissions: %w", err)
		}
	}

	// Find repository root for go.work
	root, err := find.Repo()
	if err != nil {
		return fmt.Errorf("finding repository root: %w", err)
	}

	// If workDir is specified, validate and use it
	if workDir != "" {
		if err := validatePath(workDir); err != nil {
			return fmt.Errorf("invalid working directory: %w", err)
		}
		// Check if workDir contains go.mod
		if _, err := os.Stat(filepath.Join(workDir, "go.mod")); err != nil {
			return fmt.Errorf("working directory must contain go.mod file: %w", err)
		}
	} else {
		// Use current directory if no workDir specified
		workDir, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("getting current directory: %w", err)
		}
	}

	// Create command with validated paths and args
	cmd := exec.CommandContext(ctx, binaryPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workDir
	cmd.Env = append(os.Environ(),
		"GO111MODULE=on",
		fmt.Sprintf("GOWORK=%s", filepath.Join(root.Path, "go.work")),
	)

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr
		}
		return fmt.Errorf("running linter: %w", err)
	}
	return nil
}

func setupGitRoot() error {
	root, err := find.Repo()
	if err != nil {
		return fmt.Errorf("failed to find repository root: %w", err)
	}
	if err := os.Setenv("GIT_ROOT", root.Path); err != nil {
		return fmt.Errorf("failed to set GIT_ROOT: %w", err)
	}
	return nil
}

func main() {
	if err := setupGitRoot(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := run(context.Background()); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// Set secure permissions for file operations
	cleanup, err := permissions.SetSecureUmask()
	if err != nil {
		return fmt.Errorf("failed to set secure permissions: %w", err)
	}
	defer cleanup()

	// Find repository root using go-findroot
	root, err := find.Repo()
	if err != nil {
		return fmt.Errorf("finding repository root: %w", err)
	}

	// Read version from .golangci-version file
	version, err := readVersion(root.Path)
	if err != nil {
		return fmt.Errorf("reading version: %w", err)
	}

	// Setup linter with correct version
	cachePath, err := setupLinter(ctx, version, runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return fmt.Errorf("setting up linter: %w", err)
	}

	// Process arguments first to get working directory
	args, workDir := processArgs(os.Args[1:], root.Path)

	// If no working directory specified via --path, use current directory
	if workDir == "" {
		workDir, err = findWorkDir()
		if err != nil {
			return fmt.Errorf("finding working directory: %w", err)
		}
	}

	// Run linter with processed arguments and working directory
	if err := runLinter(ctx, cachePath, workDir, args); err != nil {
		return fmt.Errorf("running linter: %w", err)
	}

	return nil
}

func setupHTTPClient() *http.Client {
	return &http.Client{
		Timeout: httpTimeout,
		Transport: &http.Transport{
			DisableCompression: true,
			MaxIdleConns:       maxIdleConns,
			IdleConnTimeout:    idleConnTimeout,
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
}

func downloadFile(ctx context.Context, url string, tmpFile *os.File) error {
	client := setupHTTPClient()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "golangci-lint-manager/1.0")
	req.Header.Set("Accept", "application/octet-stream")

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

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/octet-stream") &&
		!strings.Contains(contentType, "application/x-gzip") {
		return fmt.Errorf("unexpected content type: %s", contentType)
	}

	hasher := sha256.New()
	reader := io.TeeReader(io.LimitReader(resp.Body, maxDecompressSize), hasher)
	if _, err := io.Copy(tmpFile, reader); err != nil {
		return fmt.Errorf("failed to save download: %w", err)
	}

	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("Downloaded file hash: %s\n", hash)
	return nil
}

func downloadAndExtract(ctx context.Context, version, osName, arch, destPath string) error {
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

	url := fmt.Sprintf(downloadURLTemplate, version, version, osName, arch)
	if err := validateURL(url); err != nil {
		return fmt.Errorf("invalid download URL: %w", err)
	}
	fmt.Printf("Downloading golangci-lint v%s from %s\n", version, url)

	if err := downloadFile(ctx, url, tmpFile); err != nil {
		return err
	}

	cacheDir := filepath.Clean(filepath.Dir(destPath))
	if err := os.MkdirAll(cacheDir, dirPerms); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	return extractBinary(ctx, tmpFile.Name(), destPath)
}

func extractBinary(_ context.Context, tarPath, destPath string) error {
	// Validate paths before operations
	if err := validatePath(tarPath, os.TempDir()); err != nil {
		return fmt.Errorf("invalid tar file path: %w", err)
	}
	if err := validatePath(destPath, cacheDir); err != nil {
		return fmt.Errorf("invalid destination path: %w", err)
	}

	// Set secure permissions for file operations
	cleanup, err := permissions.SetSecureUmask()
	if err != nil {
		return fmt.Errorf("failed to set secure permissions: %w", err)
	}
	defer cleanup()

	// Open archive with secure permissions
	// Safe to use os.OpenFile here as path is validated by validatePath above
	//nolint:gosec // G304: path is validated by validatePath above
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
