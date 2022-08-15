package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/synapse-node/testutils/debug/etherscan"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// GenerateABIFromEtherscan generates the abi for an etherscan file.
func GenerateABIFromEtherscan(ctx context.Context, chainID uint, contractAddress common.Address, fileName, solVersion, pkgName string) error {
	client, err := etherscan.NewEtherscanAbiGenClientFromChain(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not etherscan client for chain %d", chainID)
	}

	contractSource, err := client.ContractSource(contractAddress.String())
	if err != nil {
		return fmt.Errorf("could not get contract source for address %s: %w", contractAddress, err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not determine wd: %w", err)
	}

	// create a temporary sol file in the current dir so it can be referenced by docker
	solFile, err := os.Create(fmt.Sprintf("%s/%s.sol", wd, path.Base(fileName)))
	if err != nil {
		return fmt.Errorf("could not create temporary sol file: %w", err)
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

	return BuildTemplates(solVersion, solFile.Name(), pkgName, fileName, optimizerRuns)
}

// BuildTemplates builds the templates. version is the solidity version to use and sol is the solidity file to use.
func BuildTemplates(version, file, pkg, filename string, optimizerRuns int) error {
	// TODO ast
	contracts, err := compileSolidity(version, file, optimizerRuns)
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

	err = ioutil.WriteFile(fmt.Sprintf("%s.abigen.go", filename), []byte(code), 0600)
	if err != nil {
		return fmt.Errorf("could not write abigen file: %w", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s.contractinfo.json", filename), marshalledContracts, 0600)
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
func compileSolidity(version string, filePath string, optimizeRuns int) (map[string]*compiler.Contract, error) {
	runFile, err := createRunFile(version)
	if err != nil {
		return nil, err
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not determine working dir: %w", err)
	}
	//nolint: gosec
	solContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read sol file %s: %w", filePath, err)
	}

	// create a temporary sol file in the current dir so it can be referenced by docker
	solFile, err := os.Create(fmt.Sprintf("%s/%s", wd, path.Base(filePath)))
	if err != nil {
		return nil, fmt.Errorf("could not create temporary sol file: %w", err)
	}
	_, err = solFile.Write(solContents)
	if err != nil {
		return nil, fmt.Errorf("could not write to sol tmp file at %s: %w", solFile.Name(), err)
	}

	err = solFile.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close solidity file handle %s: %w", solFile.Name(), err)
	}

	defer func() {
		if err == nil {
			err = os.Remove(solFile.Name())
		} else {
			_ = os.Remove(solFile.Name())
		}
	}()

	args := []string{"--combined-json", "bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes", "--optimize", "--optimize-runs", strconv.Itoa(optimizeRuns), "--allow-paths", "., ./, ../"}

	nbusy := 0
	var stderr, stdout bytes.Buffer
	for {
		// reset
		stderr = bytes.Buffer{}
		stdout = bytes.Buffer{}

		// compile the solidity
		//nolint: gosec
		cmd := exec.Command(runFile.Name(), append(args, "--", fmt.Sprintf("/solidity/%s", filepath.Base(solFile.Name())))...)
		cmd.Stderr = &stderr
		cmd.Stdout = &stdout

		yo := gofakeit.Word()
		fmt.Println(yo)
		err = cmd.Run()
		fmt.Println(yo)
		// cmd.Run will fail on Unix if some other process has the binary
		// we want to run open for writing.  This can happen here because
		// we build and install the cgo command and then run it.
		// If another command was kicked off while we were writing the
		// cgo binary, the child process for that command may be holding
		// a reference to the fd, keeping us from running exec.
		//
		// But, you might reasonably wonder, how can this happen?
		// The cgo fd, like all our fds, is close-on-exec, so that we need
		// not worry about other processes inheriting the fd accidentally.
		// The answer is that running a command is fork and exec.
		// A child forked while the cgo fd is open inherits that fd.
		// Until the child has called exec, it holds the fd open and the
		// kernel will not let us run cgo.  Even if the child were to close
		// the fd explicitly, it would still be open from the time of the fork
		// until the time of the explicit close, and the race would remain.
		//
		// On Unix systems, this results in ETXTBSY, which formats
		// as "text file busy".  Rather than hard-code specific error cases,
		// we just look for that string.  If this happens, sleep a little
		// and try again.  We let this happen three times, with increasing
		// sleep lengths: 100+200+400 ms = 0.7 seconds.
		//
		// An alternate solution might be to split the cmd.Run into
		// separate cmd.Start and cmd.Wait, and then use an RWLock
		// to make sure that copyFile only executes when no cmd.Start
		// call is in progress.  However, cmd.Start (really syscall.forkExec)
		// only guarantees that when it returns, the exec is committed to
		// happen and succeed.  It uses a close-on-exec file descriptor
		// itself to determine this, so we know that when cmd.Start returns,
		// at least one close-on-exec file descriptor has been closed.
		// However, we cannot be sure that all of them have been closed,
		// so the program might still encounter ETXTBSY even with such
		// an RWLock.  The race window would be smaller, perhaps, but not
		// guaranteed to be gone.
		//
		// Sleeping when we observe the race seems to be the most reliable
		// option we have.
		//
		// http://golang.org/issue/3001
		//
		fmt.Println("nbusy " + strconv.Itoa(nbusy))
		if err != nil {
			fmt.Println("I got an error")
			fmt.Println(err.Error())
		} else {
			fmt.Println("no err")
		}
		if err != nil && nbusy < 3 && strings.Contains(err.Error(), "text file busy") {
			fmt.Println("true")
			time.Sleep(100 * time.Millisecond << uint(nbusy))
			nbusy++
			continue
		}
		break
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

	err = runFile.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close temp file: %w", err)
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
