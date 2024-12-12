package internal_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/tools/abigen/internal"
)

func TestCheckForDocker(t *testing.T) {
	t.Helper()

	// Test Docker availability
	err := internal.CheckForDocker()
	if err != nil {
		t.Logf("Docker not available: %v", err)
		return
	}
	Nil(t, err, "Docker should be available in test environment")
}

func (a *AbiSuite) TestCompileSolidityImplicitEVM() {
	ctx := context.Background()
	vals, err := internal.CompileSolidity(ctx, "0.8.4", a.exampleFilePath, 1, nil)
	Nil(a.T(), err)

	Len(a.T(), vals, 1)
	for _, value := range vals {
		Equal(a.T(), value.Info.CompilerVersion, "0.8.4")
		Equal(a.T(), value.Info.LanguageVersion, "0.8.4")
	}
}

func (a *AbiSuite) TestCompileSolidityExplicitEVM() {
	// default would be shnghai
	const testEvmVersion = "istanbul"
	ctx := context.Background()
	vals, err := internal.CompileSolidity(ctx, "0.8.20", a.exampleFilePath, 1, core.PtrTo(testEvmVersion))
	Nil(a.T(), err)

	Len(a.T(), vals, 1)
	for _, value := range vals {
		Equal(a.T(), value.Info.CompilerVersion, "0.8.20")
		Equal(a.T(), value.Info.LanguageVersion, "0.8.20")

		var metadata ContractMetadata
		err = json.Unmarshal([]byte(value.Info.Metadata), &metadata)
		a.Require().NoError(err)

		if metadata.Settings.EvmVersion != testEvmVersion {
			a.T().Errorf("expected %s to be %s", metadata.Language, testEvmVersion)
		}
	}
}

func TestFilePathsAreEqual(t *testing.T) {
	tests := []struct {
		file1 string
		file2 string
		want  bool
		err   error
	}{
		{"path/to/file1.txt", "path/to/file2.txt", false, nil},
		{"path/to/file1.txt", "path/to/file1.txt", true, nil},
		{"path/to/file2.txt", "path/to/file2.txt", true, nil},
		{"path/to/file1.txt", "", false, filepath.ErrBadPattern},
		{"", "path/to/file2.txt", false, filepath.ErrBadPattern},
		{"nonexistent/file.txt", "path/to/file.txt", false, nil},
	}

	for _, tt := range tests {
		got, err := internal.FilePathsAreEqual(tt.file1, tt.file2)

		if got != tt.want {
			t.Errorf("filePathsAreEqual(%v, %v) got %v, want %v", tt.file1, tt.file2, got, tt.want)
		}

		if err != nil && !errors.Is(err, tt.err) {
			t.Errorf("filePathsAreEqual(%v, %v) error got %v, want %v", tt.file1, tt.file2, err, tt.err)
		}
	}
}

func TestCompileSolidityBinaryFallback(t *testing.T) {
	ctx := context.Background()
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.sol")
	content := []byte("pragma solidity ^0.8.0; contract Test {}")
	err := os.WriteFile(testFile, content, 0600)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name       string
		version    string
		evmVersion string
		wantErr    bool
	}{
		{
			name:    "valid version",
			version: "0.8.20",
			wantErr: false,
		},
		{
			name:       "with evm version",
			version:    "0.8.20",
			evmVersion: "london",
			wantErr:    false,
		},
		{
			name:    "invalid version",
			version: "999.999.999",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var evmPtr *string
			if tt.evmVersion != "" {
				evmPtr = &tt.evmVersion
			}

			_, err := internal.CompileSolidity(ctx, tt.version, testFile, 200, evmPtr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompileSolidity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompileSolidityErrors(t *testing.T) {
	ctx := context.Background()
	tmpDir := t.TempDir()

	tests := []struct {
		name    string
		content string
		version string
		wantErr string
	}{
		{
			name:    "syntax error",
			content: "pragma solidity ^0.8.0; contract Test { function invalid ",
			version: "0.8.20",
			wantErr: "exit status 1",
		},
		{
			name:    "empty file",
			content: "",
			version: "0.8.20",
			wantErr: "empty source file",
		},
		{
			name:    "invalid pragma",
			content: "pragma solidity ^999.999.999; contract Test {}",
			version: "0.8.20",
			wantErr: "Source file requires different compiler version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testFile := filepath.Join(tmpDir, fmt.Sprintf("test_%s.sol", tt.name))
			err := os.WriteFile(testFile, []byte(tt.content), 0600)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			_, err = internal.CompileSolidity(ctx, tt.version, testFile, 200, nil)
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("CompileSolidity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompileSolidityDockerStrategy(t *testing.T) {
	t.Helper()
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.sol")
	content := []byte("pragma solidity ^0.8.0; contract Test {}")
	err := os.WriteFile(testFile, content, 0600)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	ctx := context.Background()

	// Test Docker compilation scenarios
	t.Run("docker scenarios", func(t *testing.T) {
		if err := internal.CheckForDocker(); err != nil {
			t.Skip("Docker not available")
		}

		t.Run("successful compilation", func(t *testing.T) {
			_, err := internal.CompileSolidity(ctx, "0.8.20", testFile, 200, nil)
			if err != nil {
				t.Errorf("Expected successful Docker compilation, got error: %v", err)
			}
		})

		t.Run("context cancellation", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := internal.CompileSolidity(ctx, "0.8.20", testFile, 200, nil)
			if err == nil {
				t.Error("Expected error due to canceled context, got nil")
			}
		})
	})

	// Test binary fallback
	t.Run("binary fallback", func(t *testing.T) {
		tmpPath := filepath.Join(tmpDir, "fake-path")
		if err := os.Mkdir(tmpPath, 0700); err != nil {
			t.Fatalf("Failed to create fake PATH dir: %v", err)
		}
		if err := os.WriteFile(filepath.Join(tmpPath, "docker"), []byte("#!/bin/sh\nexit 1"), 0600); err != nil {
			t.Fatalf("Failed to create fake docker binary: %v", err)
		}

		t.Setenv("PATH", tmpPath)
		origPath := os.Getenv("PATH")
		t.Cleanup(func() { t.Setenv("PATH", origPath) })

		t.Run("successful fallback", func(t *testing.T) {
			_, err := internal.CompileSolidity(ctx, "0.8.20", testFile, 200, nil)
			if err != nil {
				t.Errorf("Expected successful binary fallback, got error: %v", err)
			}
		})

		t.Run("context cancellation", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := internal.CompileSolidity(ctx, "0.8.20", testFile, 200, nil)
			if err == nil {
				t.Error("Expected error due to canceled context during fallback, got nil")
			}
		})
	})
}

// ContractSettings outed by solc.
type ContractSettings struct {
	CompilationTarget map[string]string `json:"compilationTarget"`
	EvmVersion        string            `json:"evmVersion"`
	// TODO implement w/ ast
	Libraries struct{}          `json:"libraries"`
	Metadata  map[string]string `json:"metadata"`
	Optimizer struct {
		Enabled bool `json:"enabled"`
		Runs    int  `json:"runs"`
	} `json:"optimizer"`
	Remappings []interface{} `json:"remappings"`
}

// ContractMetadata is metadata produced by solc.
type ContractMetadata struct {
	Compiler struct {
		Version string `json:"version"`
	} `json:"compiler"`
	Language string                 `json:"language"`
	Output   interface{}            `json:"output"`
	Settings ContractSettings       `json:"settings"`
	Sources  map[string]interface{} `json:"sources"`
	Version  int                    `json:"version"`
}
