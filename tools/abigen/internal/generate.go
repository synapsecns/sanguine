package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/tools/abigen/internal/etherscan"
)

// GenerateABIFromEtherscan generates the abi for an etherscan file.
func GenerateABIFromEtherscan(ctx context.Context, chainID uint32, url string, contractAddress common.Address, fileName, solVersion, pkgName string) error {
	client, err := etherscan.NewEtherscanAbiGenClientFromChain(ctx, chainID, url)
	if err != nil {
		return fmt.Errorf("could not etherscan client for chain %d", chainID)
	}

	contractSource, err := client.ContractSource(contractAddress.String())
	if err != nil {
		return fmt.Errorf("could not get contract source for address %s: %w", contractAddress, err)
	}

	solFile, err := os.Create(fmt.Sprintf("%s/%s.sol", os.TempDir(), path.Base(fileName)))
	if err != nil {
		return fmt.Errorf("could not determine wd: %w", err)
	}

	defer func() {
		_ = solFile.Close()
		_ = os.RemoveAll(solFile.Name())
	}()

	var optimizerRuns int
	for _, contract := range contractSource {
		_, err = solFile.WriteString(contract.SourceCode)
		if err != nil {
			return fmt.Errorf("could not write to temporary sol file %s: %w", solFile.Name(), err)
		}
		optimizerRuns = contract.Runs
	}

	return BuildTemplates(solVersion, solFile.Name(), pkgName, fileName, optimizerRuns, nil)
}

// BuildTemplates builds the templates. version is the solidity version to use and sol is the solidity file to use.
func BuildTemplates(version, file, pkg, filename string, optimizerRuns int, evmVersion *string) error {
	// TODO ast
	contracts, err := compileSolidity(version, file, optimizerRuns, evmVersion)
	if err != nil {
		return err
	}

	marshalledContracts, err := json.Marshal(&contracts)
	if err != nil {
		return fmt.Errorf("could not marshall contracts: %w", err)
	}

	var (
		abis    []string
		bins    []string
		types   []string
		sigs    []map[string]string
		libs    = make(map[string]string)
		aliases = make(map[string]string)
	)

	// Gather all non-excluded contract for binding
	for name, contract := range contracts {
		abi, err := json.Marshal(contract.Info.AbiDefinition) // Flatten the compiler parse
		if err != nil {
			panic(fmt.Errorf("failed to parse ABIs from compiler output: %w", err))
		}
		abis = append(abis, string(abi))
		bins = append(bins, contract.Code)
		sigs = append(sigs, contract.Hashes)
		nameParts := strings.Split(name, ":")
		types = append(types, nameParts[len(nameParts)-1])

		libPattern := crypto.Keccak256Hash([]byte(name)).String()[2:36]
		libs[libPattern] = nameParts[len(nameParts)-1]
	}

	code, err := bind.Bind(types, abis, bins, sigs, pkg, bind.LangGo, libs, aliases)
	if err != nil {
		return fmt.Errorf("could not generate abigen file: %w", err)
	}

	err = os.WriteFile(fmt.Sprintf("%s.abigen.go", filename), []byte(code), 0600)
	if err != nil {
		return fmt.Errorf("could not write abigen file: %w", err)
	}

	err = os.WriteFile(fmt.Sprintf("%s.contractinfo.json", filename), marshalledContracts, 0600)
	if err != nil {
		return fmt.Errorf("could not write contract info file: %w", err)
	}

	f, err := os.Create(fmt.Sprintf("%s.metadata.go", filename))
	if err != nil {
		return fmt.Errorf("could not create metadata file: %w", err)
	}
	_ = metadataTemplate.Execute(f, struct {
		PackageName string
		Name        string
	}{
		PackageName: pkg,
		Name:        filename,
	})
	return nil
}

// compileSolidity uses docker to compile solidity with a fallback to direct solc binary.
// nolint: cyclop
func compileSolidity(version string, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not determine working dir: %w", err)
	}
	//nolint: gosec
	solContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read sol file %s: %w", filePath, err)
	}

	// create a temporary sol file in the current dir so it can be referenced by docker
	tmpPath := fmt.Sprintf("%s/%s", wd, path.Base(filePath))

	var solFile *os.File

	// whether the original file is at the temporary path already
	originalAtTmpPath, err := filePathsAreEqual(tmpPath, filePath)
	if err != nil {
		return nil, fmt.Errorf("could not compare file paths: %w", err)
	}

	// we don't need to create a temporary file if it's already in our path!
	//nolint: nestif
	if !originalAtTmpPath {
		//nolint: gosec
		solFile, err = os.Create(tmpPath)
		if err != nil {
			return nil, fmt.Errorf("could not create temporary sol file: %w", err)
		}
		_, err = solFile.Write(solContents)
		if err != nil {
			return nil, fmt.Errorf("could not write to sol tmp file at %s: %w", solFile.Name(), err)
		}

		defer func() {
			if err == nil {
				err = os.Remove(solFile.Name())
			} else {
				_ = os.Remove(solFile.Name())
			}
		}()
	} else {
		// nolint: gosec
		solFile, err = os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("could not read to sol file at %s: %w", solFile.Name(), err)
		}
	}

	// compile the solidity
	var stderr, stdout bytes.Buffer
	args := []string{"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes", "--optimize", "--optimize-runs", strconv.Itoa(optimizeRuns), "--allow-paths", "., ./, ../"}

	if evmVersion != nil {
		args = append(args, fmt.Sprintf("--evm-version=%s", *evmVersion))
	}

	// First try to use Docker
	runFile, err := createRunFile(version)
	if err == nil {
		defer runFile.Close()
		
		//nolint: gosec
		cmd := exec.Command(runFile.Name(), append(args, "--", fmt.Sprintf("/solidity/%s", filepath.Base(solFile.Name())))...)
		cmd.Stderr = &stderr
		cmd.Stdout = &stdout

		err = cmd.Run()
		if err == nil {
			// Docker worked, parse the output
			contract, err := compiler.ParseCombinedJSON(stdout.Bytes(), string(solContents), version, version, strings.Join(args, " "))
			if err != nil {
				return nil, fmt.Errorf("could not parse json: %w", err)
			}
			return contract, nil
		}
		
		// If we get here, Docker failed
		fmt.Fprintf(os.Stderr, "Docker-based solc failed: %v\nFalling back to direct solc binary\n", err)
		stderr.Reset()
		stdout.Reset()
	} else {
		fmt.Fprintf(os.Stderr, "Could not create Docker run file: %v\nFalling back to direct solc binary\n", err)
	}

	// Fallback to direct solc binary
	solcPath, err := getSolcBinary(version)
	if err != nil {
		return nil, fmt.Errorf("failed to get solc binary: %w", err)
	}

	// For direct solc binary execution, we need to adjust the command arguments
	//nolint: gosec
	cmd := exec.Command(solcPath, append(args, filepath.Base(solFile.Name()))...)
	cmd.Dir = filepath.Dir(solFile.Name()) // Set working directory to where the solidity file is
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc direct binary execution failed: %w\n%s", err, stderr.Bytes())
	}

	// Parse output
	contract, err := compiler.ParseCombinedJSON(stdout.Bytes(), string(solContents), version, version, strings.Join(args, " "))
	if err != nil {
		return nil, fmt.Errorf("could not parse json: %w", err)
	}

	// Normalize the contract names to match Docker format
	// Docker uses "solidity/file.sol:Contract" format while direct binary uses "file.sol:Contract"
	normalizedContracts := make(map[string]*compiler.Contract)
	for key, value := range contract {
		// Check if the key already has the "solidity/" prefix
		if !strings.HasPrefix(key, "solidity/") {
			normalizedKey := "solidity/" + key
			
			// We also need to normalize the metadata to ensure the compilationTarget path matches
			var metadataObj map[string]interface{}
			if err := json.Unmarshal([]byte(value.Info.Metadata), &metadataObj); err == nil {
				if settings, ok := metadataObj["settings"].(map[string]interface{}); ok {
					if compilationTarget, ok := settings["compilationTarget"].(map[string]interface{}); ok {
						// Create a new compilationTarget with the normalized paths
						newTarget := make(map[string]interface{})
						for filePath, contractName := range compilationTarget {
							if !strings.HasPrefix(filePath, "solidity/") {
								newTarget["solidity/"+filePath] = contractName
							} else {
								newTarget[filePath] = contractName
							}
						}
						// Replace the old compilationTarget with the new one
						settings["compilationTarget"] = newTarget
						
						// Marshal the metadata back to string
						if newMetadata, err := json.Marshal(metadataObj); err == nil {
							value.Info.Metadata = string(newMetadata)
						}
					}
				}
			}
			
			normalizedContracts[normalizedKey] = value
		} else {
			normalizedContracts[key] = value
		}
	}
	
	return normalizedContracts, nil
}

// createRunFile creates a bash file to run a command in the specified version of solidity.
func createRunFile(version string) (runFile *os.File, err error) {
	runFile, err = os.CreateTemp("", "*")
	if err != nil {
		return nil, fmt.Errorf("could not create temp file: %w", err)
	}

	// create a bash file that runs solidity with args passed to the run file
	_, err = runFile.WriteString(fmt.Sprintf("#!/bin/bash -e \n$(which docker) run --platform linux/amd64 -v $(pwd):/solidity ethereum/solc:%s \"$@\"", version))
	if err != nil {
		return nil, fmt.Errorf("could not create temp file: %w", err)
	}
	// TODO this should really be done natively. See: https://pkg.go.dev/bitbucket.org/dchapes/mode maybe?
	//nolint: gosec
	err = exec.Command("chmod", "+x", runFile.Name()).Run()
	if err != nil {
		return nil, fmt.Errorf("could not make sol runner executable")
	}

	return runFile, nil
}

// TODO consider only building on test,	as these contracts *will* get included into production binaries.
var metadataTemplate = template.Must(template.New("").Parse(`// Code generated by synapse abigen DO NOT EDIT.
package {{ .PackageName }}

import (
	_ "embed"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/compiler"
)

// rawContracts are the json we use to derive the processed contracts
//
//go:embed {{ .Name }}.contractinfo.json
var rawContracts []byte

// Contracts are unmarshalled on start
var Contracts map[string]*compiler.Contract

func init() {
	// load contract metadata
	var err error
	err = json.Unmarshal(rawContracts, &Contracts)
	if err != nil {
		panic(err)
	}
}
`))

func filePathsAreEqual(file1 string, file2 string) (equal bool, err error) {
	// get the absolute path of the files
	file1, err = filepath.Abs(file1)
	if err != nil {
		return false, fmt.Errorf("could not get absolute path of %s: %w", file1, err)
	}
	file2, err = filepath.Abs(file2)
	if err != nil {
		return false, fmt.Errorf("could not get absolute path of %s: %w", file2, err)
	}

	// check if the files are the same
	return file1 == file2, nil
}

// getSolcBinary downloads and caches a solc binary for the given version.
// It returns the path to the binary.
func getSolcBinary(version string) (string, error) {
	// Get user's home directory for cache placement
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine user's home directory: %w", err)
	}
	
	// Create cache directory if it doesn't exist
	cacheDir := filepath.Join(homeDir, ".cache", "solc")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", fmt.Errorf("could not create cache directory: %w", err)
	}
	
	// Normalize version string (remove 'v' prefix if present)
	normalizedVersion := strings.TrimPrefix(version, "v")
	
	// Path to the cached binary
	solcPath := filepath.Join(cacheDir, fmt.Sprintf("solc-%s", normalizedVersion))
	
	// Check if we already have the binary
	if _, err := os.Stat(solcPath); err == nil {
		// Binary exists, make sure it's executable
		if err := os.Chmod(solcPath, 0755); err != nil {
			return "", fmt.Errorf("could not set executable bit on cached solc: %w", err)
		}
		return solcPath, nil
	}
	
	// Use a file lock to prevent concurrent downloads
	lockFile := filepath.Join(cacheDir, fmt.Sprintf("solc-%s.lock", normalizedVersion))
	
	// Create lock file with exclusive lock
	lock, err := os.OpenFile(lockFile, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0600)
	
	// If lock file exists, another process is downloading, wait a bit and retry
	if err != nil && os.IsExist(err) {
		// Another process is downloading, wait a bit and retry checking for the binary
		time.Sleep(500 * time.Millisecond)
		
		// Check again if the binary exists now
		if _, err := os.Stat(solcPath); err == nil {
			if err := os.Chmod(solcPath, 0755); err != nil {
				return "", fmt.Errorf("could not set executable bit on cached solc: %w", err)
			}
			return solcPath, nil
		}
		
		// Binary still doesn't exist, something might be wrong with the other process
		// Try to acquire the lock again
		lock, err = os.OpenFile(lockFile, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0600)
		if err != nil {
			return "", fmt.Errorf("could not acquire lock for downloading solc: %w", err)
		}
	} else if err != nil {
		return "", fmt.Errorf("could not create lock file: %w", err)
	}
	
	// Make sure to clean up the lock file when we're done
	defer func() {
		lock.Close()
		_ = os.Remove(lockFile)
	}()
	
	// Double-check if binary exists now (might have been downloaded by another process)
	if _, err := os.Stat(solcPath); err == nil {
		if err := os.Chmod(solcPath, 0755); err != nil {
			return "", fmt.Errorf("could not set executable bit on cached solc: %w", err)
		}
		return solcPath, nil
	}
	
	// Determine platform and architecture
	platform := determinePlatform()
	
	// URL to download the binary from
	downloadURL, err := getSolcURL(normalizedVersion, platform)
	if err != nil {
		return "", err
	}
	
	fmt.Fprintf(os.Stderr, "Downloading solc %s for %s from %s\n", normalizedVersion, platform, downloadURL)
	
	// Download the binary
	resp, err := http.Get(downloadURL)
	if err != nil {
		return "", fmt.Errorf("could not download solc: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not download solc: HTTP %d", resp.StatusCode)
	}
	
	// Create a temporary file to download to
	tmpFile, err := os.CreateTemp(cacheDir, "solc-download-*")
	if err != nil {
		return "", fmt.Errorf("could not create temporary file: %w", err)
	}
	defer tmpFile.Close()
	
	// Copy the response to the temporary file
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not write downloaded solc to disk: %w", err)
	}
	
	// Make the file executable
	if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
		return "", fmt.Errorf("could not set executable bit on downloaded solc: %w", err)
	}
	
	// Move the temporary file to the final location
	if err := os.Rename(tmpFile.Name(), solcPath); err != nil {
		return "", fmt.Errorf("could not move downloaded solc to final location: %w", err)
	}
	
	return solcPath, nil
}

// determinePlatform returns the platform-specific string for downloading solc
func determinePlatform() string {
	// Check if this is Apple Silicon
	var isAppleSilicon bool
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("uname", "-m")
		output, err := cmd.Output()
		if err == nil && strings.TrimSpace(string(output)) == "arm64" {
			isAppleSilicon = true
		}
	}
	
	switch {
	case isAppleSilicon:
		return "macosx-aarch64"
	case runtime.GOOS == "darwin":
		return "macosx-amd64"
	case runtime.GOOS == "linux":
		return "linux-amd64"
	case runtime.GOOS == "windows":
		return "windows-amd64"
	default:
		// Default to Linux, which is most commonly supported
		return "linux-amd64"
	}
}

// getSolcURL returns the URL to download the solc binary for the given version and platform
func getSolcURL(version, platform string) (string, error) {
	// Base URL for solc binaries
	baseURL := "https://github.com/ethereum/solidity/releases/download/v" + version
	
	// For apple silicon (arm64), we need to use a slightly different approach.
	// The apple silicon binaries are only available starting with more recent solc versions
	if platform == "macosx-aarch64" {
		// For Apple Silicon, check if we're using a version that has native binaries
		if isSemverGreaterOrEqual(version, "0.8.5") {
			// The correct filename is "solc-macos" with no architecture suffix
			return fmt.Sprintf("%s/solc-macos", baseURL), nil
		}
		
		// Fall back to x86_64 version that will run using Rosetta 2
		platform = "macosx-amd64"
	}
	
	// Handle specific platform formatting
	switch platform {
	case "macosx-amd64":
		return fmt.Sprintf("%s/solc-macos", baseURL), nil
	case "linux-amd64":
		return fmt.Sprintf("%s/solc-static-linux", baseURL), nil
	case "windows-amd64":
		return fmt.Sprintf("%s/solc-windows.exe", baseURL), nil
	default:
		// Default to Linux AMD64 as fallback
		return fmt.Sprintf("%s/solc-static-linux", baseURL), nil
	}
}

// isSemverGreaterOrEqual checks if version a is greater than or equal to version b
func isSemverGreaterOrEqual(a, b string) bool {
	aParts := strings.Split(a, ".")
	bParts := strings.Split(b, ".")
	
	// Pad with zeros to ensure both have same number of parts
	for len(aParts) < 3 {
		aParts = append(aParts, "0")
	}
	for len(bParts) < 3 {
		bParts = append(bParts, "0")
	}
	
	// Compare major version
	aMajor, _ := strconv.Atoi(aParts[0])
	bMajor, _ := strconv.Atoi(bParts[0])
	if aMajor > bMajor {
		return true
	}
	if aMajor < bMajor {
		return false
	}
	
	// Compare minor version
	aMinor, _ := strconv.Atoi(aParts[1])
	bMinor, _ := strconv.Atoi(bParts[1])
	if aMinor > bMinor {
		return true
	}
	if aMinor < bMinor {
		return false
	}
	
	// Compare patch version
	aPatch, _ := strconv.Atoi(aParts[2])
	bPatch, _ := strconv.Atoi(bParts[2])
	return aPatch >= bPatch
}
