package solc_test

import (
	"github.com/synapsecns/sanguine/tools/abigen/internal/solc"
	"os"
	"path/filepath"
	"runtime"
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
