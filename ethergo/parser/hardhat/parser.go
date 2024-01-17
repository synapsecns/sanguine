package hardhat

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// copied from hardhat provider because it's unexported.
type hardhatMetadata struct {
	Compiler ContractCompiler           `json:"compiler"`
	Sources  map[string]ContractSources `json:"sources"`
}

// getMetadata gets the metadata for a hardhat contract.
func getMetadata(contract HardhatContract, filePath string) (hardhatMeta hardhatMetadata, err error) {
	if contract.Metadata != "" {
		err := json.Unmarshal([]byte(contract.Metadata), &hardhatMeta)
		if err != nil {
			return hardhatMeta, fmt.Errorf("failed parsing build file metadata at %s with error: %w", filePath, err)
		}
	}
	return hardhatMeta, nil
}

// updateNetworks updates the networks for a contract.
func getNetworks(contract HardhatContract, filePath string) (map[string]ContractNetwork, error) {
	if contract.Networks == nil {
		contract.Networks = make(map[string]ContractNetwork)
	}

	chainIDPath := filepath.Join(filePath, ".chainId")

	//nolint: gosec
	chainIDData, err := os.ReadFile(chainIDPath)
	if err != nil {
		return contract.Networks, fmt.Errorf("failed reading chainID file at %s: %w", chainIDPath, err)
	}

	var chainID int
	err = json.Unmarshal(chainIDData, &chainID)
	if err != nil {
		return contract.Networks, fmt.Errorf("could not parse chainid at %s: %w", filePath, err)
	}

	contract.Networks[strconv.Itoa(chainID)] = ContractNetwork{
		Address:         contract.Address,
		TransactionHash: contract.Receipt.TransactionHash,
	}

	return contract.Networks, nil
}

// GetDeployments parses all contract deployments from a directory.
//
//nolint:gocognit,cyclop
func GetDeployments(deploymentDir string) (contracts []Contract, err error) {
	var files []os.DirEntry

	files, err = os.ReadDir(deploymentDir)
	if err != nil {
		return nil, errors.Wrap(err, "failed listing build files")
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		filePath := filepath.Join(deploymentDir, file.Name())
		contractFiles, _ := os.ReadDir(filePath)

		for _, contractFile := range contractFiles {
			if contractFile.IsDir() || !strings.HasSuffix(contractFile.Name(), ".json") {
				continue
			}

			contractFilePath := filepath.Join(filePath, contractFile.Name())

			//nolint: gosec
			data, err := os.ReadFile(contractFilePath)
			if err != nil {
				logger.Debug(fmt.Sprintf("Failed reading build file at %s with error: %s", contractFilePath, err))
				break
			}

			var hardhatContract HardhatContract
			err = json.Unmarshal(data, &hardhatContract)
			if err != nil {
				logger.Debug(fmt.Sprintf("Failed parsing build file at %s with error: %s", contractFilePath, err))
				break
			}

			metadata, err := getMetadata(hardhatContract, filePath)
			if err != nil {
				logger.Debugf("failed parsing build file metadata at %s with error: %v", contractFilePath, err)
				break
			}

			contract := Contract{
				Name:             strings.Split(contractFile.Name(), ".")[0],
				Abi:              hardhatContract.Abi,
				Bytecode:         hardhatContract.Bytecode,
				DeployedBytecode: hardhatContract.DeployedBytecode,
				SourcePath:       filePath,
				Compiler: ContractCompiler{
					Name:    "",
					Version: metadata.Compiler.Version,
				},
			}

			contract.Networks, err = getNetworks(hardhatContract, filePath)
			if err != nil {
				logger.Debugf("failed parsing networks at %s with error: %v", filePath, err)
				break
			}

			contract.SourcePath = fmt.Sprintf("%s/%s.sol", contract.SourcePath, contract.Name)

			useContract := true
			for _, info := range contract.Networks {
				// if either of these fields are empty they were not deployed by us
				if info.TransactionHash == "" || info.Address == "" {
					useContract = false
				}
			}

			if useContract {
				contracts = append(contracts, contract)
			}
		}
	}

	return contracts, nil
}
