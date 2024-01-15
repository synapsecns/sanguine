package geth

import (
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/phayes/freeport"
	"math/big"
	"testing"
)

// makeNodeConfig makes the config for a full backend.
func makeNodeConfig(tb testing.TB) *node.Config {
	tb.Helper()
	// generate node config
	nodeCfg := node.DefaultConfig
	//  see: https://github.com/ethereum/go-ethereum/blob/cc606be74c6f1f05b0b0a6226a400e734b9aac31/cmd/geth/config.go#L100
	nodeCfg.Name = "geth"
	nodeCfg.HTTPHost = "127.0.0.1"
	nodeCfg.HTTPPort = freeport.GetPort()
	nodeCfg.HTTPModules = append(nodeCfg.HTTPModules, "eth", "debug")

	nodeCfg.WSHost = "127.0.0.1"
	nodeCfg.WSPort = freeport.GetPort()
	nodeCfg.WSModules = append(nodeCfg.WSModules, "eth", "debug")
	nodeCfg.IPCPath = filet.TmpDir(tb, "")
	nodeCfg.DataDir = ""
	// see https://github.com/ethereum/go-ethereum/blob/cc606be74c6f1f05b0b0a6226a400e734b9aac31/cmd/utils/flags.go#L1194
	nodeCfg.P2P.MaxPeers = 0
	nodeCfg.P2P.ListenAddr = ""
	nodeCfg.P2P.NoDial = true
	nodeCfg.P2P.DiscoveryV5 = false
	// allow debugging via remix.
	//nolint: gosec
	nodeCfg.HTTPCors = append(nodeCfg.HTTPCors, "http://remix.ethereum.org")

	return &nodeCfg
}

// makeEthConfig gets the eth config for a mock node.
func makeEthConfig(address common.Address, config *params.ChainConfig) *ethconfig.Config {
	ethConfig := ethconfig.Defaults
	ethConfig.NetworkId = config.ChainID.Uint64()
	ethConfig.Genesis = core.DeveloperGenesisBlock(0, 10000000, address)
	ethConfig.Genesis.Config = config
	ethConfig.Miner.Etherbase = address
	ethConfig.SyncMode = downloader.FullSync
	ethConfig.TxPool.AccountSlots = 50
	ethConfig.Miner.GasPrice = big.NewInt(1)
	ethConfig.EnablePreimageRecording = true
	ethConfig.RPCTxFeeCap = params.Ether * 10
	ethConfig.NoPruning = true
	ethConfig.TxLookupLimit = 0
	ethConfig.Preimages = true
	return &ethConfig
}
