package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	downloadURLTemplate = "https://github.com/golangci/golangci-lint/releases/download/v%s/golangci-lint-%s-%s-%s.tar.gz"
	cacheDir           = "cache"
)

// findGitRoot returns the root directory of the git repository
func findGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error finding git root: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func main() {
	// Find repository root
	root, err := findGitRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding repository root: %v\n", err)
		os.Exit(1)
	}

	// Read version from .golangci-version in repo root
	version, err := os.ReadFile(filepath.Join(root, ".golangci-version"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading .golangci-version: %v\n", err)
		os.Exit(1)
	}
	versionStr := strings.TrimSpace(string(version))

	// Determine architecture
	arch := runtime.GOARCH
	osName := runtime.GOOS

	// Create cache path relative to the binary location
	execPath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting executable path: %v\n", err)
		os.Exit(1)
	}
	execDir := filepath.Dir(execPath)
	binaryName := fmt.Sprintf("golangci-lint-%s-%s-%s/golangci-lint", versionStr, osName, arch)
	cachePath := filepath.Join(execDir, cacheDir, binaryName)

	// Check cache and download if needed
	if _, err := os.Stat(cachePath); err != nil {
		if err := downloadAndExtract(versionStr, osName, arch, cachePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error setting up golangci-lint: %v\n", err)
			os.Exit(1)
		}
	}

	// Ensure the binary is executable
	if err := os.Chmod(cachePath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error making binary executable: %v\n", err)
		os.Exit(1)
	}

	// Check if we're being run directly or through go run
	args := os.Args[1:]

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Find nearest directory with go.mod
	workDir := cwd
	for dir := cwd; dir != "/" && dir != "."; dir = filepath.Dir(dir) {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			workDir = dir
			break
		}
	}

	// Calculate relative path from root to workDir
	relPath, err := filepath.Rel(root, workDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calculating relative path: %v\n", err)
		os.Exit(1)
	}

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
	args = append(args, "--modules-download-mode=readonly")

	// If no go.mod found, disable module-specific linters
	if workDir == cwd {
		args = append(args, "--disable=gomodguard,depguard")
	}

	// Add path to lint if not specified
	if len(args) == 1 || (len(args) > 1 && strings.HasPrefix(args[1], "-")) {
		args = append(args, "./...")
	}

	// Execute golangci-lint with arguments
	cmd := exec.Command(cachePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workDir // Run from the directory with go.mod
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}

func downloadAndExtract(version, osName, arch, destPath string) error {
	// Create temporary file for download
	tmpFile, err := os.CreateTemp("", "golangci-lint-*.tar.gz")
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Download URL
	url := fmt.Sprintf(downloadURLTemplate, version, version, osName, arch)
	fmt.Printf("Downloading golangci-lint v%s from %s\n", version, url)

	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading golangci-lint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error downloading golangci-lint: HTTP %d", resp.StatusCode)
	}

	// Copy to temp file
	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return fmt.Errorf("error saving download: %v", err)
	}
	tmpFile.Close()

	// Create cache directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("error creating cache directory: %v", err)
	}

	return extractBinary(tmpFile.Name(), destPath)
}

func extractBinary(tarPath, destPath string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return fmt.Errorf("error opening archive: %v", err)
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("error creating gzip reader: %v", err)
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading tar: %v", err)
		}

		if strings.HasSuffix(header.Name, "golangci-lint") {
			destFile, err := os.Create(destPath)
			if err != nil {
				return fmt.Errorf("error creating destination file: %v", err)
			}
			defer destFile.Close()

			if _, err := io.Copy(destFile, tr); err != nil {
				return fmt.Errorf("error extracting file: %v", err)
			}
			return nil
		}
	}

	return fmt.Errorf("golangci-lint binary not found in archive")
}
