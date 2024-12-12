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

func main() {
	// Read version from .golangci-version
	version, err := os.ReadFile("../../.golangci-version")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading .golangci-version: %v\n", err)
		os.Exit(1)
	}
	versionStr := strings.TrimSpace(string(version))

	// Determine architecture
	arch := runtime.GOARCH
	osName := runtime.GOOS

	// Create cache path
	binaryName := fmt.Sprintf("golangci-lint-%s-%s-%s/golangci-lint", versionStr, osName, arch)
	cachePath := filepath.Join(cacheDir, binaryName)

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

	// Execute golangci-lint with remaining arguments
	args := os.Args[1:]
	cmd := exec.Command(cachePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
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
