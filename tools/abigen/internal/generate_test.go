package internal_test

import (
	"encoding/json"
	"errors"
	"github.com/synapsecns/sanguine/core"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/tools/abigen/internal"
)

func TestCreateRunFile(t *testing.T) {
	runFile, err := internal.CreateRunFile("0.8.17")
	Nil(t, err)

	//nolint: gosec
	cmd := exec.Command("bash", "-n", runFile.Name())
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		Nil(t, err)
	}
}

func (a *AbiSuite) TestCompileSolidityImplicitEVM() {
	vals, err := internal.CompileSolidity("0.8.4", a.exampleFilePath, 1, nil)
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
	vals, err := internal.CompileSolidity("0.8.20", a.exampleFilePath, 1, core.PtrTo(testEvmVersion))
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

func TestGetSolcBinary(t *testing.T) {
	// Skip in CI to avoid unnecessary downloads
	if os.Getenv("CI") != "" {
		t.Skip("Skipping in CI environment")
	}

	// Test accessing the getSolcBinary function through the exported test function
	solcPath, err := internal.GetSolcBinary("0.8.4")
	Nil(t, err)
	NotEmpty(t, solcPath)

	// Verify the file exists
	info, err := os.Stat(solcPath)
	Nil(t, err)
	
	// Verify it's executable
	True(t, info.Mode()&0111 != 0, "solc binary should be executable")
	
	// Verify it's in the expected location
	homeDir, err := os.UserHomeDir()
	Nil(t, err)
	expectedDir := filepath.Join(homeDir, ".cache", "solc")
	True(t, filepath.HasPrefix(solcPath, expectedDir), 
		"expected solc binary to be in %s, but it was in %s", expectedDir, solcPath)
}

// TestCompileSolidityWithoutDocker tests that our fallback mechanism works end-to-end
func TestCompileSolidityWithoutDocker(t *testing.T) {
	// Skip in CI to avoid unnecessary downloads
	if os.Getenv("CI") != "" {
		t.Skip("Skipping in CI environment")
	}
	
	// Create a temporary Solidity file
	content := `
	// SPDX-License-Identifier: MIT
	pragma solidity ^0.8.4;
	
	contract SimpleStorage {
		uint256 private value;
		
		function setValue(uint256 _value) public {
			value = _value;
		}
		
		function getValue() public view returns (uint256) {
			return value;
		}
	}
	`
	
	tempDir, err := os.MkdirTemp("", "solc-test-*")
	Nil(t, err)
	defer os.RemoveAll(tempDir)
	
	tempFile := filepath.Join(tempDir, "SimpleStorage.sol")
	err = os.WriteFile(tempFile, []byte(content), 0644)
	Nil(t, err)
	
	// Force docker to fail by setting an invalid environment variable
	originalPath := os.Getenv("PATH")
	os.Setenv("DOCKER_HOST", "unix:///non-existent-path")
	defer os.Setenv("PATH", originalPath)
	
	// Compile using our fallback mechanism
	contracts, err := internal.CompileSolidity("0.8.4", tempFile, 200, nil)
	Nil(t, err)
	
	// Verify that we got the contract
	NotNil(t, contracts)
	NotEqual(t, 0, len(contracts))
	
	// Check that the compiled contract actually exists
	var found bool
	for name, contract := range contracts {
		if strings.Contains(name, "SimpleStorage") {
			found = true
			NotEmpty(t, contract.Code)
			NotNil(t, contract.Info.AbiDefinition)
		}
	}
	
	True(t, found, "SimpleStorage contract not found in compiled output")
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
