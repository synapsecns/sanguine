package hardhat

import (
	"github.com/tenderly/tenderly-cli/hardhat"
	"github.com/tenderly/tenderly-cli/providers"
)

// Create a test interface so we can check outputs.
type HardhatMetadataTest interface {
	GetCompiler() providers.ContractCompiler
	GetSources() map[string]providers.ContractSources
}

// implement the getters.
func (h hardhatMetadata) GetCompiler() providers.ContractCompiler {
	return h.Compiler
}

func (h hardhatMetadata) GetSources() map[string]providers.ContractSources {
	return h.Sources
}

func GetMetadata(contract hardhat.HardhatContract, filePath string) (hardhatMeta HardhatMetadataTest, err error) {
	return getMetadata(contract, filePath)
}

func GetNetworks(contract hardhat.HardhatContract, filePath string) (map[string]providers.ContractNetwork, error) {
	return getNetworks(contract, filePath)
}
