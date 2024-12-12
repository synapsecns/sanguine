package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
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

	// Check cache
	if _, err := os.Stat(cachePath); err == nil {
		// Ensure the binary is executable
		if err := os.Chmod(cachePath, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error making cached binary executable: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Using cached golangci-lint v%s\n", versionStr)
		os.Exit(0)
	}

	// Download URL
	url := fmt.Sprintf(downloadURLTemplate, versionStr, versionStr, osName, arch)
	fmt.Printf("Downloading golangci-lint v%s from %s\n", versionStr, url)

	// Create temporary file for download
	tmpFile, err := os.CreateTemp("", "golangci-lint-*.tar.gz")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating temp file: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile.Name())

	// Download the file
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading golangci-lint: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error downloading golangci-lint: HTTP %d\n", resp.StatusCode)
		os.Exit(1)
	}

	// Copy to temp file
	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving download: %v\n", err)
		os.Exit(1)
	}
	tmpFile.Close()

	// Create cache directory if it doesn't exist
	cacheParentDir := filepath.Dir(cachePath)
	if err := os.MkdirAll(cacheParentDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating cache directory: %v\n", err)
		os.Exit(1)
	}

	// Extract the binary
	if err := extractBinary(tmpFile.Name(), cachePath); err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting binary: %v\n", err)
		os.Exit(1)
	}

	// Make binary executable
	if err := os.Chmod(cachePath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error making binary executable: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully installed golangci-lint v%s\n", versionStr)
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
