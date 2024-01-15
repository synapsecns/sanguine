package preset

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
)

// GetRinkeby gets the rinkeby preset backend.
func GetRinkeby() Backend {
	chainConfig := *params.TestChainConfig
	chainConfig.ChainID = params.RinkebyChainConfig.ChainID

	return Backend{
		config:     &chainConfig,
		rpcURL:     core.GetEnv("RINEKBY_RPC_URL", "ws://0.0.0.0:8045"),
		name:       "Rinkeby",
		privateKey: os.Getenv("EXPORT_KEY"),
	}
}

// GetBSCTestnet gets the bsc backend.
func GetBSCTestnet() Backend {
	chainConfig := *params.AllCliqueProtocolChanges
	chainConfig.ChainID = client.ChapelChainConfig.ChainID

	return Backend{
		config:     &chainConfig,
		rpcURL:     core.GetEnv("BSC_TESTNET_RPC_URL", "ws://0.0.0.0:8046"),
		name:       "BSC Testnet",
		privateKey: os.Getenv("EXPORT_KEY"),
	}
}

// GetAvalancheLocal gets the avalanche local config.
// TODO: this should use avalanche.
func GetAvalancheLocal() Backend {
	chainConfig := *params.AllCliqueProtocolChanges
	chainConfig.ChainID = core.CopyBigInt(client.AvalancheLocalChainConfig.ChainID)

	return Backend{
		config:     &chainConfig,
		rpcURL:     core.GetEnv("MATIC_RPC_URL", "ws://0.0.0.0:8048"),
		name:       "Avalanche",
		privateKey: os.Getenv("EXPORT_KEY"),
	}
}

// GetMaticMumbaiFakeSynDomain gets the matic backend.
func GetMaticMumbaiFakeSynDomain() Backend {
	chainConfig := *params.AllCliqueProtocolChanges
	chainConfig.ChainID = big.NewInt(int64(10))
	// london is not activated on bsc
	chainConfig.LondonBlock = nil

	return Backend{
		config:     &chainConfig,
		rpcURL:     core.GetEnv("MATIC_RPC_URL", "ws://0.0.0.0:8049"),
		name:       "Matic",
		privateKey: os.Getenv("EXPORT_KEY"),
	}
}
