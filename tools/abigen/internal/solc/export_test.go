package solc

import (
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"testing"
)

// ResetDownloadAttempts resets the download attempts counter for testing.
var ResetDownloadAttempts = resetDownloadAttempts

// SetClient allows setting the HTTP client for testing.
func (m *BinaryManager) SetClient(client interface{}) {
	if httpClient, ok := client.(*http.Client); ok {
		m.client = httpClient
	}
}

// UseFixedTempFileNames enables deterministic temp file naming for testing.
func (m *BinaryManager) UseFixedTempFileNames() {
	m.useFixedTempFileNames = true
}

// DownloadAndCacheSolcBinary downloads and caches solc binaries for testing.
func DownloadAndCacheSolcBinary(t *testing.T) map[string][]byte {
	t.Helper()
	return map[string][]byte{
		"0.8.20": []byte("mock binary 0.8.20"),
		"0.8.19": []byte("mock binary 0.8.19"),
	}
}

// FailingTransport is a test transport that simulates download failures.
type FailingTransport struct {
	MaxSuccesses     int
	MockBinary       []byte
	Binaries         map[string][]byte
	SkipBackoffDelay bool
	SuccessCount     atomic.Int32
	CacheHits        atomic.Int32
}

// RoundTrip implements http.RoundTripper for FailingTransport.
func (f *FailingTransport) RoundTrip(*http.Request) (*http.Response, error) {
	// Always return connection refused error for testing concurrent error handling
	return nil, &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: fmt.Errorf("connection refused"),
	}
}
