package solc_test

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/tools/abigen/internal/solc"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

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
	version := "0.8.20"
	manager := solc.NewBinaryManager(version)

	expectedCacheDir := filepath.Join(os.Getenv("HOME"), ".cache", "solc")
	if manager.CacheDir() != expectedCacheDir {
		t.Errorf("Expected cache dir %s, got %s", expectedCacheDir, manager.CacheDir())
	}

	if manager.Version() != version {
		t.Errorf("Expected version %s, got %s", version, manager.Version())
	}
}

func TestGetPlatformDir(t *testing.T) {
	manager := solc.NewBinaryManager("0.8.20")
	platform := manager.GetPlatformDir()

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
	manager := solc.NewBinaryManager("0.8.20")
	binary, err := manager.GetBinary()
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

	binary2, err := manager.GetBinary()
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
	defer os.RemoveAll(tmpDir)

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
			wantErr:  "version 999.999.999 not found",
		},
		{
			name:     "invalid platform",
			version:  "0.8.20",
			platform: "invalid-platform",
			wantErr:  "failed to download list.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := solc.NewBinaryManager(tt.version)
			managerValue := reflect.ValueOf(manager).Elem()
			if platformField := managerValue.FieldByName("platform"); platformField.IsValid() && platformField.CanSet() {
				platformField.SetString(tt.platform)
			}
			_, err := manager.GetBinary()
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("GetBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDownloadAndVerify(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "solc-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name     string
		version  string
		setup    func(*testing.T, string) (*solc.BinaryManager, error)
		wantErr  string
	}{
		{
			name:    "invalid permissions",
			version: "0.8.20",
			setup: func(t *testing.T, dir string) (*solc.BinaryManager, error) {
				noWriteDir := filepath.Join(dir, "no-write")
				if err := os.Mkdir(noWriteDir, 0500); err != nil {
					t.Fatalf("Failed to create no-write dir: %v", err)
				}
				if err := os.Chmod(noWriteDir, 0500); err != nil {
					t.Fatalf("Failed to set directory permissions: %v", err)
				}
				manager := solc.NewBinaryManager("0.8.20")
				managerValue := reflect.ValueOf(manager).Elem()
				if cacheDirField := managerValue.FieldByName("cacheDir"); cacheDirField.IsValid() && cacheDirField.CanSet() {
					cacheDirField.SetString(noWriteDir)
				}
				return manager, nil
			},
			wantErr: "failed to create cache directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager, err := tt.setup(t, tmpDir)
			if err != nil {
				t.Fatalf("Setup failed: %v", err)
			}
			_, err = manager.GetBinary()
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
	defer os.RemoveAll(tmpDir)

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
			info := &solc.BinaryInfo{
				Sha256:    tt.sha256,
				Keccak256: tt.keccak256,
			}
			manager := solc.NewBinaryManager("0.8.20")
			err := manager.VerifyChecksums(testFile, info)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyChecksums() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
