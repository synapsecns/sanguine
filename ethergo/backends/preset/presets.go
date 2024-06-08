package preset

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
)

// GetSepolia gets the rinkeby preset backend.
func GetSepolia() Backend {
	chainConfig := *params.AllDevChainProtocolChanges
	chainConfig.ChainID = params.SepoliaChainConfig.ChainID

	return Backend{
		config:     &chainConfig,
		rpcURL:     core.GetEnv("SEPOLIA_RPC_URL", "ws://0.0.0.0:8045"),
		name:       "Sepolia",
		privateKey: os.Getenv("EXPORT_KEY"),
	}
}

// GetBSCTestnet gets the bsc backend.
func GetBSCTestnet() Backend {
	chainConfig := *params.AllDevChainProtocolChanges
	chainConfig.ChainID = client.ChapelChainConfig.ChainID

	return Backend{
		config:     &chainConfig,
		rpcURL:     core.GetEnv("BSC_TESTNET_RPC_URL", "ws://0.0.0.0:8046"),
		name:       "BSC Testnet",
		privateKey: os.Getenv("EXPORT_KEY"),
	}
}

// GetMaticMumbai gets the matic backend.
func GetMaticMumbai() Backend {
	chainConfig := *params.AllDevChainProtocolChanges
	chainConfig.ChainID = client.MaticMainnetConfig.ChainID
	// london is not activated on bsc
	chainConfig.LondonBlock = big.NewInt(0)

	return Backend{
		config:     &chainConfig,
		rpcURL:     core.GetEnv("MATIC_RPC_URL", "ws://0.0.0.0:8047"),
		name:       "Matic",
		privateKey: os.Getenv("EXPORT_KEY"),
	}
}

// GetMaticMumbaiFakeSynDomain gets the matic backend.
func GetMaticMumbaiFakeSynDomain() Backend {
	chainConfig := *params.AllDevChainProtocolChanges
	chainConfig.ChainID = big.NewInt(int64(10))
	// london is not activated on bsc

	return Backend{
		config:     &chainConfig,
		rpcURL:     core.GetEnv("MATIC_RPC_URL", "ws://0.0.0.0:8049"),
		name:       "Matic",
		privateKey: os.Getenv("EXPORT_KEY"),
	}
}
