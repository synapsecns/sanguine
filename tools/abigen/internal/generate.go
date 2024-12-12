package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/tools/abigen/internal/etherscan"
	"github.com/synapsecns/sanguine/tools/abigen/internal/solc"
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

	//nolint:gosec // File operations required for temporary solidity file
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
func BuildTemplates(version, file, pkg, filename string, optimizeRuns int, evmVersion *string) error {
	// TODO ast
	contracts, err := compileSolidity(version, file, optimizeRuns, evmVersion)
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

	//nolint:gosec // File operations required for abigen output
	err = os.WriteFile(fmt.Sprintf("%s.abigen.go", filename), []byte(code), 0600)
	if err != nil {
		return fmt.Errorf("could not write abigen file: %w", err)
	}

	//nolint:gosec // File operations required for contract info output
	err = os.WriteFile(fmt.Sprintf("%s.contractinfo.json", filename), marshalledContracts, 0600)
	if err != nil {
		return fmt.Errorf("could not write contract info file: %w", err)
	}

	//nolint:gosec // File operations required for metadata output
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

// compileSolidity attempts to compile the given Solidity file using either Docker or direct binary.
// nolint: cyclop
func compileSolidity(version string, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	// Always try Docker first regardless of platform
	contract, err := compileWithDocker(version, filePath, optimizeRuns, evmVersion)
	if err == nil {
		return contract, nil
	}

	// Only fall back to binary if Docker fails
	return compileWithBinary(version, filePath, optimizeRuns, evmVersion)
}

// compileWithDocker uses Docker to compile solidity.
func compileWithDocker(version string, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	// Check Docker availability first
	if err := checkForDocker(); err != nil {
		return nil, fmt.Errorf("docker not available (falling back to binary): %w", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not determine working dir: %w", err)
	}

	solContents, err := readSolFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read solidity file: %w", err)
	}

	tmpPath := filepath.Join(wd, filepath.Base(filePath))
	solFile, err := prepareSolFile(tmpPath, filePath, solContents)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare solidity file: %w", err)
	}

	if !isOriginalFile(tmpPath, filePath) {
		defer func() {
			if cleanupErr := os.Remove(solFile.Name()); cleanupErr != nil {
				if err == nil {
					err = cleanupErr
				}
			}
		}()
	}

	// Run Docker compilation with direct commands
	var stderr, stdout bytes.Buffer
	args := []string{
		"run", "--rm",
		"--platform", "linux/amd64", // Ensure consistent platform
		"-v", fmt.Sprintf("%s:/solidity", filepath.Dir(solFile.Name())),
		fmt.Sprintf("ethereum/solc:%s", version),
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes",
		"--optimize",
		"--optimize-runs", strconv.Itoa(optimizeRuns),
		"--allow-paths", "/solidity",
		fmt.Sprintf("/solidity/%s", filepath.Base(solFile.Name())),
	}

	if evmVersion != nil {
		args = append(args[0:len(args)-1], fmt.Sprintf("--evm-version=%s", *evmVersion), args[len(args)-1])
	}

	cmd := exec.Command("docker", args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return nil, fmt.Errorf("docker compilation failed: %s", stderr.String())
		}
		return nil, fmt.Errorf("docker execution failed: %w", err)
	}

	contracts, err := compiler.ParseCombinedJSON(stdout.Bytes(), string(solContents), version, version, strings.Join(args, " "))
	if err != nil {
		return nil, fmt.Errorf("failed to parse combined JSON output: %w", err)
	}
	return contracts, nil
}

//nolint:gosec // File operations required for reading Solidity source
func readSolFile(filePath string) ([]byte, error) {
	solContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read sol file %s: %w", filePath, err)
	}
	return solContents, nil
}

func prepareSolFile(tmpPath, filePath string, solContents []byte) (*os.File, error) {
	originalAtTmpPath, err := filePathsAreEqual(tmpPath, filePath)
	if err != nil {
		return nil, fmt.Errorf("could not compare file paths: %w", err)
	}

	if !originalAtTmpPath {
		//nolint:gosec // File operations required for temporary solidity file
		solFile, err := os.Create(tmpPath)
		if err != nil {
			return nil, fmt.Errorf("could not create temporary sol file: %w", err)
		}
		if _, err = solFile.Write(solContents); err != nil {
			return nil, fmt.Errorf("could not write to sol tmp file at %s: %w", solFile.Name(), err)
		}
		return solFile, nil
	}

	//nolint:gosec // File operations required for reading Solidity source
	solFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read to sol file at %s: %w", filePath, err)
	}
	return solFile, nil
}

func isOriginalFile(tmpPath, filePath string) bool {
	equal, _ := filePathsAreEqual(tmpPath, filePath)
	return equal
}

// compileWithBinary uses downloaded solc binary to compile solidity.
func compileWithBinary(version string, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	binaryManager := solc.NewBinaryManager(version)
	ctx := context.Background() // TODO: Accept context from caller

	// Get solc binary
	binaryPath, err := binaryManager.GetBinary(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get solc binary: %w", err)
	}

	// Convert to absolute path and validate
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for %s: %w", filePath, err)
	}

	//nolint:gosec // File operations required for reading Solidity source
	solContents, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("could not read sol file %s: %w", absPath, err)
	}

	// Check for empty file
	if len(solContents) == 0 {
		return nil, fmt.Errorf("empty source file")
	}

	var stderr, stdout bytes.Buffer
	baseDir := filepath.Dir(absPath)
	args := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes",
		"--optimize",
		"--optimize-runs", strconv.Itoa(optimizeRuns),
		"--allow-paths", baseDir,
	}

	if evmVersion != nil {
		args = append(args, fmt.Sprintf("--evm-version=%s", *evmVersion))
	}

	//nolint:gosec // Command execution with validated solc binary is required
	cmd := exec.Command(binaryPath, append(args, absPath)...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	cmd.Dir = baseDir

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc: %w\n%s", err, stderr.Bytes())
	}

	contracts, err := compiler.ParseCombinedJSON(stdout.Bytes(), string(solContents), version, version, strings.Join(args, " "))
	if err != nil {
		return nil, fmt.Errorf("failed to parse combined JSON output: %w", err)
	}
	return contracts, nil
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

// checkForDocker verifies that Docker is available and running.
func checkForDocker() error {
	cmd := exec.Command("docker", "info")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return fmt.Errorf("docker not running: %s", stderr.String())
		}
		return fmt.Errorf("docker not installed or not in PATH: %w", err)
	}

	return nil
}
