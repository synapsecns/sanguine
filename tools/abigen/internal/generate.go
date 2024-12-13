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

	return BuildTemplates(ctx, solVersion, solFile.Name(), pkgName, fileName, optimizerRuns, nil)
}

// BuildTemplates builds the templates. version is the solidity version to use and sol is the solidity file to use.
func BuildTemplates(ctx context.Context, version, file, pkg, filename string, optimizeRuns int, evmVersion *string) error {
	// TODO ast
	contracts, err := compileSolidity(ctx, version, file, optimizeRuns, evmVersion)
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

// compileSolidity attempts to compile the given Solidity file using either Docker or direct binary.
// nolint: cyclop
func compileSolidity(ctx context.Context, version string, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("compilation canceled: %w", ctx.Err())
	default:
	}

	// Always try Docker first regardless of platform
	contracts, dockerErr := compileWithDocker(ctx, version, filePath, optimizeRuns, evmVersion)
	if dockerErr == nil {
		return contracts, nil
	}
	// Log Docker failure but continue to binary fallback
	fmt.Printf("Docker compilation failed: %v, attempting binary compilation...\n", dockerErr)

	// Only fall back to binary if Docker fails
	// Add solc-solc- prefix for binary compilation
	formattedVersion := fmt.Sprintf("solc-solc-%s", strings.TrimPrefix(strings.TrimPrefix(version, "solc-solc-"), "solc-"))
	contracts, binaryErr := compileWithBinary(ctx, formattedVersion, filePath, optimizeRuns, evmVersion)
	if binaryErr != nil {
		// Return both Docker and binary errors for better debugging
		return nil, fmt.Errorf("solidity compilation failed:\nDocker error: %v\nBinary error: %v\nVersion: %s\nFile: %s",
			dockerErr, binaryErr, version, filePath)
	}

	return contracts, nil
}

// compileWithDocker uses Docker to compile solidity.
func compileWithDocker(ctx context.Context, version, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("docker compilation canceled: %w", ctx.Err())
	default:
	}

	// Check Docker availability first
	if err := checkForDocker(); err != nil {
		return nil, fmt.Errorf("docker not available (falling back to binary): %w", err)
	}

	// Read and validate source file first
	solContents, err := readSolFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read solidity file: %w", err)
	}
	if len(solContents) == 0 {
		return nil, fmt.Errorf("empty source file")
	}

	// Prepare solidity file in the same directory as the source
	solFile, err := prepareSolidityFile(filepath.Dir(filePath), solContents)
	if err != nil {
		return nil, err
	}

	// Execute Docker compilation with context
	return executeDockerCompilation(ctx, version, solFile, solContents, optimizeRuns, evmVersion)
}

// prepareSolidityFile prepares the solidity file for compilation and handles cleanup.
func prepareSolidityFile(sourceDir string, solContents []byte) (*os.File, error) {
	// Create temporary file in the source directory with .sol extension
	tmpFile, err := os.CreateTemp(sourceDir, "test-*.sol")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file in %s: %w", sourceDir, err)
	}

	// Write contents to temporary file
	if _, err := tmpFile.Write(solContents); err != nil {
		_ = os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("failed to write to temporary file: %w", err)
	}

	// Close file to ensure all data is written
	if err := tmpFile.Close(); err != nil {
		_ = os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("failed to close temporary file: %w", err)
	}

	return tmpFile, nil
}

// executeDockerCompilation executes the Docker compilation command and parses the output.
func executeDockerCompilation(ctx context.Context, version string, solFile *os.File, solContents []byte, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("docker execution canceled: %w", ctx.Err())
	default:
	}

	// Ensure version has the correct prefix for all compilation paths
	formattedVersion := fmt.Sprintf("solc-solc-%s", strings.TrimPrefix(strings.TrimPrefix(version, "solc-solc-"), "solc-"))

	var stderr, stdout bytes.Buffer
	cmd := prepareDockerCommand(version, solFile.Name(), optimizeRuns, evmVersion)
	if cmd == nil {
		return nil, fmt.Errorf("failed to prepare docker command")
	}

	// Verify Docker binary exists and is executable
	const dockerPath = "/usr/bin/docker"
	if _, err := os.Stat(dockerPath); err != nil {
		return nil, fmt.Errorf("docker binary not found at %s: %w", dockerPath, err)
	}

	// Create new command with validated arguments
	validatedCmd := exec.CommandContext(ctx, dockerPath, cmd.Args[1:]...)
	validatedCmd.Stderr = &stderr
	validatedCmd.Stdout = &stdout

	cmdStr := strings.Join(cmd.Args, " ")
	if err := validatedCmd.Run(); err != nil {
		if stderr.Len() > 0 {
			return nil, fmt.Errorf("docker compilation failed: %w\nCommand: %s\nstderr: %s\nVersion: %s",
				err, cmdStr, stderr.String(), formattedVersion)
		}
		return nil, fmt.Errorf("docker execution failed: %w\nCommand: %s\nVersion: %s",
			err, cmdStr, formattedVersion)
	}

	// Use formatted version for both compiler and language version
	contracts, err := compiler.ParseCombinedJSON(stdout.Bytes(), string(solContents), formattedVersion, formattedVersion, cmdStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse combined JSON output: %w\nOutput: %s\nCommand: %s\nVersion: %s",
			err, stdout.String(), cmdStr, formattedVersion)
	}
	return contracts, nil
}

// prepareDockerCommand creates the Docker command with all necessary arguments.
func prepareDockerCommand(version, filePath string, optimizeRuns int, evmVersion *string) *exec.Cmd {
	// Strip any prefixes for Docker image tag
	version = strings.TrimPrefix(version, "v")
	version = strings.TrimPrefix(version, "solc-")

	// Get absolute path and validate
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil
	}

	// Get source directory and ensure it exists
	sourceDir := filepath.Dir(absPath)
	if _, err := os.Stat(sourceDir); err != nil {
		return nil
	}

	// Create Docker mount path and source file path
	dockerMountPath := "/solidity"
	relPath, err := filepath.Rel(sourceDir, absPath)
	if err != nil {
		return nil
	}
	dockerSourcePath := filepath.Join(dockerMountPath, relPath)

	// Prepare base arguments with validated inputs
	args := []string{
		"run", "--rm",
		"--platform", "linux/amd64", // Ensure consistent platform
		"-v", fmt.Sprintf("%s:%s", sourceDir, dockerMountPath),
		fmt.Sprintf("ethereum/solc:%s", version),
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes",
		"--optimize",
		"--optimize-runs", strconv.Itoa(optimizeRuns),
		"--allow-paths", dockerMountPath,
		dockerSourcePath,
	}

	// Add EVM version if specified
	if evmVersion != nil {
		args = append(args[:len(args)-1], fmt.Sprintf("--evm-version=%s", *evmVersion), args[len(args)-1])
	}

	return exec.Command("docker", args...)
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
func compileWithBinary(ctx context.Context, version string, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	if err := validateBinaryCompilationInputs(ctx, version, filePath); err != nil {
		return nil, err
	}

	binaryPath, solContents, err := prepareBinaryCompilation(ctx, version, filePath)
	if err != nil {
		return nil, err
	}

	return executeBinaryCompilation(ctx, binaryPath, filePath, solContents, optimizeRuns, evmVersion)
}

// validateBinaryCompilationInputs validates the input parameters for binary compilation.
func validateBinaryCompilationInputs(ctx context.Context, version, filePath string) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("binary compilation canceled: %w", ctx.Err())
	default:
	}

	if version == "" || filePath == "" {
		return fmt.Errorf("invalid input: version and filePath must not be empty")
	}
	return nil
}

// prepareBinaryCompilation prepares the solc binary and reads the source file.
func prepareBinaryCompilation(ctx context.Context, version, filePath string) (string, []byte, error) {
	binaryManager := solc.NewBinaryManager(version)
	binaryPath, err := binaryManager.GetBinary(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get solc binary: %w", err)
	}

	if !filepath.IsAbs(binaryPath) {
		return "", nil, fmt.Errorf("solc binary path must be absolute")
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get absolute path for %s: %w", filePath, err)
	}

	//nolint:gosec // File operations required for reading Solidity source
	solContents, err := os.ReadFile(absPath)
	if err != nil {
		return "", nil, fmt.Errorf("could not read sol file %s: %w", absPath, err)
	}

	if len(solContents) == 0 {
		return "", nil, fmt.Errorf("empty source file")
	}

	return binaryPath, solContents, nil
}

// executeBinaryCompilation executes the solc binary with validated arguments.
func executeBinaryCompilation(ctx context.Context, binaryPath, filePath string, solContents []byte, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for %s: %w", filePath, err)
	}

	baseDir := filepath.Dir(absPath)
	args := prepareSolcArgs(baseDir, optimizeRuns, evmVersion)
	args = append(args, absPath)

	// Validate all paths are absolute and exist
	for i, arg := range args {
		if strings.HasPrefix(arg, "--allow-paths") {
			if !filepath.IsAbs(args[i+1]) {
				return nil, fmt.Errorf("allow-paths argument must be absolute")
			}
		}
	}

	var stderr, stdout bytes.Buffer
	//nolint:gosec // Arguments are validated above
	cmd := exec.CommandContext(ctx, binaryPath, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	cmd.Dir = baseDir

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc failed: %w\nstderr: %s", err, stderr.String())
	}

	contracts, err := compiler.ParseCombinedJSON(stdout.Bytes(), string(solContents), filepath.Base(binaryPath), filepath.Base(binaryPath), strings.Join(args, " "))
	if err != nil {
		return nil, fmt.Errorf("failed to parse solidity compilation output: %w\noutput: %s\nargs: %s\nbinary: %s",
			err, stdout.String(), strings.Join(args, " "), binaryPath)
	}
	return contracts, nil
}

// prepareSolcArgs prepares the command line arguments for solc.
func prepareSolcArgs(baseDir string, optimizeRuns int, evmVersion *string) []string {
	args := []string{
		"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes",
		"--optimize",
		"--optimize-runs", strconv.Itoa(optimizeRuns),
		"--allow-paths", baseDir,
	}

	if evmVersion != nil {
		args = append(args, fmt.Sprintf("--evm-version=%s", *evmVersion))
	}

	return args
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
