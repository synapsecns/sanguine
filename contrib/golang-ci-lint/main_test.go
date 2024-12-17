package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestValidatePath(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := os.TempDir()
	cacheDir := filepath.Join(tmpDir, "cache")

	tests := []struct {
		name       string
		path       string
		allowedDir []string
		wantErr    bool
		onlyDarwin bool // Skip test on non-Darwin systems
	}{
		{
			name:       "Valid temp path on MacOS",
			path:       "/private/var/folders/test",
			allowedDir: []string{"/var/folders"},
			wantErr:    false,
			onlyDarwin: true,
		},
		{
			name:       "Valid cache path on MacOS",
			path:       "/private/var/cache/golangci-lint",
			allowedDir: []string{"/var/cache"},
			wantErr:    false,
			onlyDarwin: true,
		},
		{
			name:       "Invalid path with directory traversal",
			path:       filepath.Join(tmpDir, "../outside"),
			allowedDir: nil,
			wantErr:    true,
		},
		{
			name:       "Valid path in temp directory",
			path:       filepath.Join(tmpDir, "test"),
			allowedDir: nil,
			wantErr:    false,
		},
		{
			name:       "Valid path in cache directory",
			path:       filepath.Join(cacheDir, "test"),
			allowedDir: nil,
			wantErr:    false,
		},
		{
			name:       "Path outside allowed directories",
			path:       "/usr/local/bin",
			allowedDir: nil,
			wantErr:    true,
		},
		{
			name:       "MacOS specific temp folder path",
			path:       "/private/var/folders/12/8xtw48x951g0vv4z_ctcr2hr0000gn/T/golangci-lint-3961080532.tar.gz",
			allowedDir: nil,
			wantErr:    false,
			onlyDarwin: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.onlyDarwin && runtime.GOOS != "darwin" {
				t.Skip("Skipping MacOS-specific test on non-Darwin system")
			}

			err := validatePath(tt.path, tt.allowedDir...)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
