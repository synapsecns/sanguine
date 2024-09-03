package main

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	ctx := context.Background()

	// Find the repo root
	repoRoot, err := findRepoRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding repo root: %v\n", err)
		os.Exit(1)
	}

	// Read the .golangci-version file
	versionFile := filepath.Join(repoRoot, ".golangci-version")
	version, err := os.ReadFile(versionFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading .golangci-version file: %v\n", err)
		os.Exit(1)
	}

	// Trim any whitespace from the version
	golangciVersion := strings.TrimSpace(string(version))

	// Determine the correct binary for the current architecture
	arch := runtime.GOOS + "-" + runtime.GOARCH
	binaryURL := fmt.Sprintf("https://github.com/golangci/golangci-lint/releases/download/v%s/golangci-lint-%s-v%s.tar.gz", golangciVersion, arch, golangciVersion)

	// Create a cache directory if it doesn't exist
	cacheDir := filepath.Join(os.TempDir(), "golangci-lint-cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating cache directory: %v\n", err)
		os.Exit(1)
	}

	// Check if the binary is already cached
	cachedBinary := filepath.Join(cacheDir, fmt.Sprintf("golangci-lint-%s-v%s", arch, golangciVersion))
	if _, err := os.Stat(cachedBinary); os.IsNotExist(err) {
		// Download the binary if it's not cached
		if err := downloadFile(ctx, binaryURL, cachedBinary+".tar.gz"); err != nil {
			fmt.Fprintf(os.Stderr, "Error downloading golangci-lint: %v\n", err)
			os.Exit(1)
		}

		// Extract the binary
		if err := extractTarGz(cachedBinary+".tar.gz", cacheDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error extracting golangci-lint: %v\n", err)
			os.Exit(1)
		}

		// Remove the tar.gz file
		os.Remove(cachedBinary + ".tar.gz")
	}

	// Run golangci-lint
	lintCmd := exec.Command(filepath.Join(cachedBinary, "golangci-lint"), "run")
	lintCmd.Dir = repoRoot
	lintCmd.Stdout = os.Stdout
	lintCmd.Stderr = os.Stderr
	if err := lintCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running golangci-lint: %v\n", err)
		os.Exit(1)
	}
}

func findRepoRoot() (string, error) {
	cmd := exec.Command("repo-root", "find")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func downloadFile(ctx context.Context, url string, filepath string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func extractTarGz(tarGzPath, destDir string) error {
	file, err := os.Open(tarGzPath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(destDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(target)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}

	return nil
}
