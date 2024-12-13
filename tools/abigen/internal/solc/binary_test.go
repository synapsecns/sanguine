package solc_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/tools/abigen/internal/solc"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func setupTestBinaryManager(t *testing.T) *solc.BinaryManager {
	t.Helper()
	return solc.NewBinaryManager("0.8.20")
}

func TestIsAppleSilicon(t *testing.T) {
	isArm := runtime.GOARCH == "arm64"
	isDarwin := runtime.GOOS == "darwin"

	result := solc.IsAppleSilicon()

	if isArm && isDarwin {
		if !result {
			t.Error("Expected true for Apple Silicon (darwin/arm64)")
		}
	} else {
		if result {
			t.Error("Expected false for non-Apple Silicon platform")
		}
	}
}

func TestNewBinaryManager(t *testing.T) {
	manager := setupTestBinaryManager(t)

	expectedCacheDir := filepath.Join(os.Getenv("HOME"), ".cache", "solc")
	if manager.CacheDir() != expectedCacheDir {
		t.Errorf("Expected cache dir %s, got %s", expectedCacheDir, manager.CacheDir())
	}

	if manager.Version() != "0.8.20" {
		t.Errorf("Expected version %s, got %s", "0.8.20", manager.Version())
	}
}

func TestGetPlatformDir(t *testing.T) {
	manager := solc.NewBinaryManager("0.8.20")
	platform, err := manager.GetPlatformDir()
	if err != nil {
		t.Fatalf("GetPlatformDir() error = %v", err)
	}

	if solc.IsAppleSilicon() {
		if platform != solc.PlatformWasm32 {
			t.Error("Expected wasm for Apple Silicon")
		}
		return
	}

	var expectedPlatform string
	switch runtime.GOOS {
	case solc.PlatformDarwin:
		expectedPlatform = solc.PlatformMacOS
	case "linux":
		expectedPlatform = solc.PlatformLinux
	case "windows":
		expectedPlatform = solc.PlatformWin
	default:
		expectedPlatform = solc.PlatformWasm32
	}

	if platform != expectedPlatform {
		t.Errorf("Expected platform %s, got %s", expectedPlatform, platform)
	}
}

func TestGetBinary(t *testing.T) {
	manager := setupTestBinaryManager(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	binary, err := manager.GetBinary(ctx)
	if err != nil {
		t.Fatalf("Failed to get binary: %v", err)
	}

	if _, err := os.Stat(binary); os.IsNotExist(err) {
		t.Error("Binary file does not exist")
	}

	info, err := os.Stat(binary)
	if err != nil {
		t.Fatalf("Failed to stat binary: %v", err)
	}
	if info.Mode()&0111 == 0 {
		t.Error("Binary is not executable")
	}

	binary2, err := manager.GetBinary(ctx)
	if err != nil {
		t.Fatalf("Failed to get cached binary: %v", err)
	}
	if binary != binary2 {
		t.Error("Cache not reused")
	}
}

func TestGetBinaryInfo(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "solc-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Errorf("failed to cleanup test directory: %v", err)
		}
	}()

	tests := []struct {
		name     string
		version  string
		platform string
		wantErr  string
	}{
		{
			name:     "invalid version",
			version:  "999.999.999",
			platform: "linux-amd64",
			wantErr:  "failed to get binary info: version 999.999.999 not found for platform linux-amd64",
		},
		{
			name:     "invalid platform",
			version:  "0.8.20",
			platform: "invalid-platform",
			wantErr:  "unsupported platform",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			manager := solc.NewBinaryManager(tt.version)
			managerValue := reflect.ValueOf(manager).Elem()
			if platformField := managerValue.FieldByName("Platform"); platformField.IsValid() && platformField.CanSet() {
				platformField.SetString(tt.platform)
				t.Logf("DEBUG: Set platform field to %q\n", tt.platform)
			} else {
				t.Error("Failed to set platform field")
			}
			_, err := manager.GetBinary(context.Background())
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("GetBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func restoreWritePermissions(t *testing.T, path string) error {
	t.Helper()
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk path %s: %w", path, err)
		}
		if err := os.Chmod(path, 0600); err != nil {
			return fmt.Errorf("failed to chmod %s: %w", path, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to restore write permissions: %w", err)
	}
	return nil
}

func setupTestDir(t *testing.T) string {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", "solc-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	t.Cleanup(func() {
		// Restore write permissions before cleanup
		if err := restoreWritePermissions(t, tmpDir); err != nil {
			t.Errorf("failed to restore write permissions: %v", err)
		}
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Errorf("failed to cleanup test directory: %v", err)
		}
	})
	return tmpDir
}

func setupNoWriteDir(t *testing.T, baseDir string) string {
	t.Helper()
	// Create the base directory with initial write permissions
	noWriteDir := filepath.Join(baseDir, "no-write")
	if err := os.MkdirAll(noWriteDir, 0600); err != nil {
		t.Fatalf("Failed to create base dir: %v", err)
	}

	// Create the version and platform subdirectories
	versionDir := filepath.Join(noWriteDir, "0.8.20")
	platformDir := filepath.Join(versionDir, "linux-amd64")
	if err := os.MkdirAll(platformDir, 0600); err != nil {
		t.Fatalf("Failed to create platform dir: %v", err)
	}

	// Remove write permissions recursively on the entire directory tree
	err := filepath.Walk(noWriteDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk path %s: %w", path, err)
		}
		if err := os.Chmod(path, 0400); err != nil {
			return fmt.Errorf("failed to chmod %s: %w", path, err)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Failed to set directory permissions: %v", err)
	}

	fmt.Printf("DEBUG: setupNoWriteDir - Created directory structure under %s with permissions 0400\n", noWriteDir)
	return noWriteDir
}

func TestDownloadAndVerify(t *testing.T) {
	tmpDir := setupTestDir(t)

	tests := []struct {
		name    string
		version string
		setup   func(t *testing.T) (*solc.BinaryManager, error)
		wantErr string
	}{
		{
			name:    "invalid permissions",
			version: "0.8.20",
			setup: func(t *testing.T) (*solc.BinaryManager, error) {
				t.Helper()
				noWriteDir := setupNoWriteDir(t, tmpDir)
				manager := solc.NewBinaryManagerWithDir("0.8.20", noWriteDir)
				managerValue := reflect.ValueOf(manager).Elem()
				if platformField := managerValue.FieldByName("Platform"); platformField.IsValid() && platformField.CanSet() {
					platformField.SetString("linux-amd64")
				}
				return manager, nil
			},
			wantErr: "failed to create cache directory: permission denied",
		},
		{
			name:    "invalid version format",
			version: "invalid.version",
			setup: func(t *testing.T) (*solc.BinaryManager, error) {
				t.Helper()
				manager := solc.NewBinaryManagerWithDir("invalid.version", tmpDir)
				managerValue := reflect.ValueOf(manager).Elem()
				if platformField := managerValue.FieldByName("Platform"); platformField.IsValid() && platformField.CanSet() {
					platformField.SetString("linux-amd64")
				}
				return manager, nil
			},
			wantErr: "invalid version format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			manager, err := tt.setup(t)
			if err != nil {
				t.Fatalf("Setup failed: %v", err)
			}
			_, err = manager.GetBinary(context.Background())
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("GetBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerifyChecksums(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "solc-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Errorf("failed to cleanup test directory: %v", err)
		}
	}()

	testContent := []byte("test content")
	testFile := filepath.Join(tmpDir, "test-binary")
	if err := os.WriteFile(testFile, testContent, 0600); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	sha256Sum := sha256.Sum256(testContent)
	keccak256Sum := crypto.Keccak256(testContent)

	tests := []struct {
		name      string
		sha256    string
		keccak256 string
		wantErr   bool
	}{
		{
			name:      "valid checksums",
			sha256:    hex.EncodeToString(sha256Sum[:]),
			keccak256: hex.EncodeToString(keccak256Sum),
			wantErr:   false,
		},
		{
			name:      "valid checksums with 0x prefix",
			sha256:    "0x" + hex.EncodeToString(sha256Sum[:]),
			keccak256: "0x" + hex.EncodeToString(keccak256Sum),
			wantErr:   false,
		},
		{
			name:      "invalid sha256",
			sha256:    "invalid",
			keccak256: hex.EncodeToString(keccak256Sum),
			wantErr:   true,
		},
		{
			name:      "invalid keccak256",
			sha256:    hex.EncodeToString(sha256Sum[:]),
			keccak256: "invalid",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			info := &solc.BinaryInfo{
				Sha256:    tt.sha256,
				Keccak256: tt.keccak256,
			}
			manager := setupTestBinaryManager(t)
			err := manager.VerifyChecksums(testFile, info)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyChecksums() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
