package hardhat

// Create a test interface so we can check outputs.
type HardhatMetadataTest interface {
	GetCompiler() ContractCompiler
	GetSources() map[string]ContractSources
}

// implement the getters.
func (h hardhatMetadata) GetCompiler() ContractCompiler {
	return h.Compiler
}

func (h hardhatMetadata) GetSources() map[string]ContractSources {
	return h.Sources
}

func GetMetadata(contract HardhatContract, filePath string) (hardhatMeta HardhatMetadataTest, err error) {
	return getMetadata(contract, filePath)
}

func GetNetworks(contract HardhatContract, filePath string) (map[string]ContractNetwork, error) {
	return getNetworks(contract, filePath)
}
