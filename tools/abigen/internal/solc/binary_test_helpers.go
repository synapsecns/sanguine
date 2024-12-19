package solc

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
)

var (
	downloadAttempts uint32
)

// createFixedTempFile creates a temporary file with a fixed name for testing concurrent access.
func (m *BinaryManager) createFixedTempFile(tmpDir, baseName string) (*os.File, string, error) {
	// Simulate more realistic concurrent access patterns
	attempts := atomic.AddUint32(&downloadAttempts, 1)

	// Fail on even-numbered attempts to create a predictable pattern
	// but only if we haven't exceeded the maximum retries
	if attempts%2 == 0 && attempts <= 8 { // 8 is maxRetries * 2
		return nil, "", os.ErrExist
	}

	// Use fixed name without random suffix to force collisions
	tmpPath := filepath.Join(tmpDir, baseName+".tmp")
	f, err := os.OpenFile(tmpPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, FilePerms)
	if err != nil {
		return nil, "", err
	}
	return f, tmpPath, nil
}

// resetDownloadAttempts resets the download attempts counter for testing.
func resetDownloadAttempts() {
	atomic.StoreUint32(&downloadAttempts, 0)
}

// shouldSkipBackoff checks if backoff delay should be skipped for testing.
func shouldSkipBackoff(transport http.RoundTripper) bool {
	if t, ok := transport.(*failingTransport); ok {
		return t.skipBackoffDelay
	}
	return false
}

// failingTransport is a mock http.RoundTripper that fails some requests.
type failingTransport struct {
	maxSuccesses     int32             // Maximum number of total successes (downloads + cache hits)
	successCount     atomic.Int32      // Total number of successful downloads
	cacheHits        atomic.Int32      // Number of cache hits
	successUrls      sync.Map          // URLs that have been successfully downloaded
	mockBinary       []byte            // Mock binary content for testing
	binaries         map[string][]byte // Version-specific binaries
	skipBackoffDelay bool              // Skip backoff delay for faster test execution
}

// downloadAndCacheSolcBinary downloads the actual solc binaries for testing.
func downloadAndCacheSolcBinary(t testing.TB) map[string][]byte {
	t.Helper()
	type binaryInfo struct {
		url    string
		commit string
	}
	info := map[string]binaryInfo{
		"0.8.20": {
			url:    "https://binaries.soliditylang.org/linux-amd64/solc-linux-amd64-v0.8.20+commit.a1b79de6",
			commit: "a1b79de6",
		},
		"0.8.19": {
			url:    "https://binaries.soliditylang.org/linux-amd64/solc-linux-amd64-v0.8.19+commit.7dd6d404",
			commit: "7dd6d404",
		},
	}
	binaries := make(map[string][]byte)
	hashes := make(map[string]string)

	for version, bi := range info {
		// Create a temporary directory for each version
		tmpDir := filepath.Join(os.TempDir(), fmt.Sprintf("solc-test-%s", version))
		if err := os.MkdirAll(tmpDir, 0750); err != nil {
			t.Fatalf("Failed to create temp directory for version %s: %v", version, err)
		}

		// Download binary
		resp, err := http.Get(bi.url)
		if err != nil {
			t.Fatalf("Failed to download solc binary %s: %v", version, err)
		}

		binary, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			t.Fatalf("Failed to read solc binary %s: %v", version, err)
		}

		// Write binary to temp file to ensure it's properly saved
		tmpFile := filepath.Join(tmpDir, fmt.Sprintf("solc-%s", version))
		if err := os.WriteFile(tmpFile, binary, 0750); err != nil {
			t.Fatalf("Failed to write temp binary %s: %v", version, err)
		}

		// Read back the binary to ensure it was written correctly
		binary, err = os.ReadFile(tmpFile)
		if err != nil {
			t.Fatalf("Failed to read temp binary %s: %v", version, err)
		}

		binaries[version] = binary

		// Calculate and log hashes
		sha256Hash := sha256.Sum256(binary)
		keccakHash := crypto.Keccak256(binary)
		hashes[version] = hex.EncodeToString(sha256Hash[:])
		t.Logf("DEBUG: SHA256 hash for version %s: %s", version, hashes[version])
		t.Logf("DEBUG: Keccak256 hash for version %s: %s", version, hex.EncodeToString(keccakHash))

		// Clean up temp file
		os.RemoveAll(tmpDir)
	}
	return binaries
}

// RoundTrip implements http.RoundTripper for failingTransport.
func (t *failingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	url := req.URL.String()

	// Always allow list.json requests without counting them towards success limit
	if strings.HasSuffix(url, "list.json") {
		// Build list.json response with both versions
		var builds []string
		var releases []string
		versions := []struct {
			version string
			commit  string
		}{
			{"0.8.20", "a1b79de6"},
			{"0.8.19", "7dd6d404"},
		}

		for _, v := range versions {
			binary := t.binaries[v.version]
			if binary == nil {
				return nil, fmt.Errorf("binary not found for version %s", v.version)
			}

			// Calculate SHA256 hash
			sha256Hash := sha256.Sum256(binary)
			sha256Str := hex.EncodeToString(sha256Hash[:])

			// Calculate Keccak256 hash using go-ethereum's crypto package
			keccakHash := crypto.Keccak256(binary)
			keccakStr := hex.EncodeToString(keccakHash)

			build := fmt.Sprintf(`{
				"path": "solc-linux-amd64-v%[1]s+commit.%[2]s",
				"version": "%[1]s",
				"build": "commit.%[2]s",
				"longVersion": "%[1]s+commit.%[2]s",
				"sha256": "%[3]s",
				"keccak256": "%[4]s",
				"urls": ["solc-linux-amd64-v%[1]s+commit.%[2]s"]
			}`, v.version, v.commit, sha256Str, keccakStr)
			builds = append(builds, build)

			release := fmt.Sprintf(`"%[1]s": "solc-linux-amd64-v%[1]s+commit.%[2]s"`,
				v.version, v.commit)
			releases = append(releases, release)
		}

		listJSON := fmt.Sprintf(`{
			"builds": [%s],
			"releases": {%s}
		}`, strings.Join(builds, ","), strings.Join(releases, ","))

		return &http.Response{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(listJSON)),
			Header:     make(http.Header),
		}, nil
	}

	// All binary download requests should fail as maxSuccesses is 0
	fmt.Printf("DEBUG: Connection refused for URL: %s (maxSuccesses: %d, skipBackoffDelay: %v)\n",
		url, t.maxSuccesses, t.skipBackoffDelay)

	// Return connection refused error immediately if skipBackoffDelay is true
	if t.skipBackoffDelay {
		return nil, &net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: syscall.ECONNREFUSED,
		}
	}

	// For non-skipped cases, return connection refused error
	return nil, &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: syscall.ECONNREFUSED,
	}
}

// createBinaryResponse creates a mock binary response for testing.
func (t *failingTransport) createBinaryResponse(url string) *http.Response {
	// Find the correct binary based on version in URL
	var binary []byte
	for version, versionBinary := range t.binaries {
		if strings.Contains(url, fmt.Sprintf("v%s", version)) {
			binary = versionBinary
			break
		}
	}

	// If no version-specific binary found, use mockBinary
	if binary == nil {
		binary = t.mockBinary
	}

	// Ensure we have valid binary content
	if len(binary) == 0 {
		panic(fmt.Sprintf("No binary content available for URL: %s", url))
	}

	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(binary)),
		Header:     make(http.Header),
	}
	resp.Header.Set("Content-Length", fmt.Sprintf("%d", len(binary)))
	return resp
}
