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

// compileSolidity uses docker to compile solidity.
// nolint: cyclop
func compileSolidity(version string, filePath string, optimizeRuns int, evmVersion *string) (map[string]*compiler.Contract, error) {
	runFile, err := createRunFile(version)
	if err != nil {
		return nil, err
	}

	_ = runFile.Close()

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

	//nolint: gosec
	cmd := exec.Command(runFile.Name(), append(args, "--", fmt.Sprintf("/solidity/%s", filepath.Base(solFile.Name())))...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc: %w\n%s", err, stderr.Bytes())
	}
	contract, err := compiler.ParseCombinedJSON(stdout.Bytes(), string(solContents), version, version, strings.Join(args, " "))
	if err != nil {
		return nil, fmt.Errorf("could not parse json: %w", err)
	}
	return contract, nil
}

// createRunFile creates a bash file to run a command in the specified version of solidity.
func createRunFile(version string) (runFile *os.File, err error) {
	runFile, err = os.CreateTemp("", "*")
	if err != nil {
		return nil, fmt.Errorf("could not create temp file: %w", err)
	}

	// create a bash file that runs solidity with args passed to the run file
	_, err = runFile.WriteString(fmt.Sprintf("#!/bin/bash -e \n$(which docker) run -v $(pwd):/solidity ethereum/solc:%s \"$@\"", version))
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

// rawContracts are the json we use to dervive the processed contracts
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
