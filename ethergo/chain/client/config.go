package client

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

// Config contains the configuration needed to connect to the chain.
type Config struct {
	// RPCUrl is the rpc websocket url to use
	// note: this is called 'WSUrl' for historical reasons and will be updated in a future version
	RPCUrl []string `toml:"WSUrl"`

	// ChainID - name of the current chain
	ChainID int `toml:"ChainID"`

	// Type is the chain type
	Type string `toml:"Type"`

	// RequiredConfirmations is the number of confirmations required until
	// a block is considered "finalized"
	RequiredConfirmations uint `toml:"Confirmations"`

	// BridgeAddress is the address of the bridge on the chain
	BridgeAddress string `toml:"BridgeAddress"`

	// StartHeight is the start height for a given contract. This is a workaround for https://github.com/synapsecns/synapse-contracts/issues/20
	// where in the start height is always returned as 0. This can be removed after version 6 of the contracts
	StartHeight uint64 `toml:"StartHeight"`

	*LimiterConfig
}

// SetEthBridgeAddress mutates the config to set a bridge address.
func (c *Config) SetEthBridgeAddress(bridgeAddress ethCommon.Address) {
	c.BridgeAddress = bridgeAddress.String()
}

// GetEthBridgeAddress gets the bridge address cast to a ethCommon.address.
func (c *Config) GetEthBridgeAddress() ethCommon.Address {
	return ethCommon.HexToAddress(c.BridgeAddress)
}

// we need chain configs from all over the place, but we don't want to
// deal with multi-chain imports so we handle it here

var (
	// BSCChainConfig bsc mainnet config.
	BSCChainConfig = &params.ChainConfig{
		ChainID:             big.NewInt(56),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// ChapelChainConfig bsc testnet config.
	ChapelChainConfig = &params.ChainConfig{
		ChainID:             big.NewInt(97),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// RialtoChainConfig rial config.
	RialtoChainConfig = &params.ChainConfig{
		ChainID:             big.NewInt(1417),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// SimulatedConfig is the simulated config backend.
	SimulatedConfig = params.AllEthashProtocolChanges

	// MaticMumbaiConfig is the matic config.
	MaticMumbaiConfig = &params.ChainConfig{
		ChainID:             big.NewInt(80001),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// MaticMainnetConfig is the matic mainnet config.
	MaticMainnetConfig = &params.ChainConfig{
		ChainID:             big.NewInt(137),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// AvalancheMainnetChainConfig is the configuration for Avalanche Main Network
	// TODO: this and other avalanche configs should be imported directly from https://github.com/ava-labs/coreth/blob/master/params/config.go
	AvalancheMainnetChainConfig = &params.ChainConfig{
		ChainID:             big.NewInt(43114),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        big.NewInt(0),
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
	}

	// AvalancheLocalChainConfig is the configuration for the Avalanche Local Network.
	AvalancheLocalChainConfig = &params.ChainConfig{
		ChainID:             big.NewInt(43112),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        big.NewInt(0),
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
	}

	// ArbitrumMainnetConfig is the arbitrum mainnet config.
	ArbitrumMainnetConfig = &params.ChainConfig{
		ChainID:             big.NewInt(42161),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// FtmMainnetConfig contains the fantom mainnet config.
	FtmMainnetConfig = &params.ChainConfig{
		ChainID:             big.NewInt(250),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// HarmonyConfig contains the chain config for harmony
	// note: this is shard 0 only
	HarmonyConfig = &params.ChainConfig{
		ChainID:             big.NewInt(1666600000),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// BobaConfig contains the chain config for boba.
	BobaConfig = &params.ChainConfig{
		ChainID:             big.NewInt(288),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// MoonBeamConfig contains the configuration for moonriver.
	MoonBeamConfig = &params.ChainConfig{
		ChainID:             big.NewInt(1284),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// MoonRiverConfig contains the configuration for moonriver.
	MoonRiverConfig = &params.ChainConfig{
		ChainID:             big.NewInt(1285),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// OptimisticEthereum contains the configuration for optimism.
	OptimisticEthereum = &params.ChainConfig{
		ChainID:             big.NewInt(10),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// AuroraMainnet contains the configuration for aurora.
	AuroraMainnet = &params.ChainConfig{
		ChainID:             big.NewInt(1313161554),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// CronosMainnet is the cronos mainnet chain config.
	CronosMainnet = &params.ChainConfig{
		ChainID:             big.NewInt(25),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// MetisMainnet is the metis mainnet config.
	MetisMainnet = &params.ChainConfig{
		ChainID:             big.NewInt(1088),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         nil,
	}

	// DFKMainnet is the dfk mainnet contract.
	DFKMainnet = &params.ChainConfig{
		ChainID:             big.NewInt(53935),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         big.NewInt(0),
	}

	// DFKTestnet is the dfk testnet config.
	DFKTestnet = &params.ChainConfig{
		ChainID:             big.NewInt(335),
		HomesteadBlock:      big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		LondonBlock:         big.NewInt(0),
	}
)

// chainConfigs is a list of chain configs.
var chainConfigs = []*params.ChainConfig{
	params.MainnetChainConfig, params.SepoliaChainConfig,
	params.RinkebyChainConfig, params.RinkebyChainConfig, params.GoerliChainConfig,
	params.AllEthashProtocolChanges, BSCChainConfig, ChapelChainConfig, RialtoChainConfig,
	SimulatedConfig, MaticMumbaiConfig, MaticMainnetConfig, AvalancheMainnetChainConfig, AvalancheLocalChainConfig,
	ArbitrumMainnetConfig, FtmMainnetConfig, HarmonyConfig, BobaConfig, MoonRiverConfig, OptimisticEthereum,
	AuroraMainnet, MoonBeamConfig, CronosMainnet, MetisMainnet, DFKMainnet, DFKTestnet,
}

// ConfigFromID gets the chain config from the id.
func ConfigFromID(id *big.Int) *params.ChainConfig {
	// make sure we don't panic on nil
	if id == nil {
		return nil
	}

	for _, potentialConfig := range chainConfigs {
		if potentialConfig.ChainID.Uint64() == id.Uint64() {
			return potentialConfig
		}
	}
	return nil
}

// UsesLondon termines whether or not a chain uses london.
func UsesLondon(config *params.ChainConfig, currentBlock uint64) bool {
	return config.LondonBlock != nil &&
		big.NewInt(int64(currentBlock)).Cmp(config.LondonBlock) >= 0
}
