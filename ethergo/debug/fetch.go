package debug

import (
	"github.com/tenderly/tenderly-cli/providers"
)

// FetchLocalContracts gets local contracts.
func (c *ContractSource) FetchLocalContracts() (resContracts []providers.Contract) {
	for _, contract := range c.localContracts {
		resContracts = append(resContracts, *contract.Contract)
	}
	return resContracts
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
