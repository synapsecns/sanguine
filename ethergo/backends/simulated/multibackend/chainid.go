package multibackend

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	util "github.com/synapsecns/sanguine/core"
)

// NewConfigWithChainID creates a new *params.ChainConfig and changes only the chain id
// everything else is taken from params.AllEthashProtocolChanges.
// we need to do this because params are global.
func NewConfigWithChainID(chainID *big.Int) *params.ChainConfig {
	return &params.ChainConfig{
		ChainID:                       util.CopyBigInt(chainID),
		HomesteadBlock:                util.CopyBigInt(params.AllEthashProtocolChanges.HomesteadBlock),
		DAOForkBlock:                  util.CopyBigInt(params.AllEthashProtocolChanges.DAOForkBlock),
		DAOForkSupport:                params.AllEthashProtocolChanges.DAOForkSupport,
		EIP150Block:                   util.CopyBigInt(params.AllEthashProtocolChanges.EIP150Block),
		EIP155Block:                   util.CopyBigInt(params.AllEthashProtocolChanges.EIP155Block),
		EIP158Block:                   util.CopyBigInt(params.AllEthashProtocolChanges.EIP158Block),
		ByzantiumBlock:                util.CopyBigInt(params.AllEthashProtocolChanges.ByzantiumBlock),
		ConstantinopleBlock:           util.CopyBigInt(params.AllEthashProtocolChanges.ConstantinopleBlock),
		PetersburgBlock:               util.CopyBigInt(params.AllEthashProtocolChanges.PetersburgBlock),
		IstanbulBlock:                 util.CopyBigInt(params.AllEthashProtocolChanges.IstanbulBlock),
		MuirGlacierBlock:              util.CopyBigInt(params.AllEthashProtocolChanges.MuirGlacierBlock),
		BerlinBlock:                   util.CopyBigInt(params.AllEthashProtocolChanges.BerlinBlock),
		LondonBlock:                   util.CopyBigInt(params.AllEthashProtocolChanges.LondonBlock),
		ArrowGlacierBlock:             util.CopyBigInt(params.AllEthashProtocolChanges.ArrowGlacierBlock),
		GrayGlacierBlock:              util.CopyBigInt(params.AllEthashProtocolChanges.GrayGlacierBlock),
		MergeNetsplitBlock:            util.CopyBigInt(params.AllEthashProtocolChanges.MergeNetsplitBlock),
		ShanghaiTime:                  util.CopyPointer(params.AllEthashProtocolChanges.ShanghaiTime),
		CancunTime:                    util.CopyPointer(params.AllEthashProtocolChanges.CancunTime),
		PragueTime:                    util.CopyPointer(params.AllEthashProtocolChanges.PragueTime),
		TerminalTotalDifficulty:       util.CopyBigInt(params.AllEthashProtocolChanges.TerminalTotalDifficulty),
		TerminalTotalDifficultyPassed: params.AllEthashProtocolChanges.TerminalTotalDifficultyPassed,
		Ethash:                        util.CopyPointer(params.AllEthashProtocolChanges.Ethash),
		Clique:                        util.CopyPointer(params.AllEthashProtocolChanges.Clique),
	}
}

// NewSimulatedBackendWithConfig creates a new simulated backend with the given chain id.
func NewSimulatedBackendWithConfig(alloc core.GenesisAlloc, gasLimit uint64, config *params.ChainConfig) *SimulatedBackend {
	database := rawdb.NewMemoryDatabase()
	triedb := trie.NewDatabase(database, nil)

	genesis := core.Genesis{Config: config, GasLimit: gasLimit, Alloc: alloc}
	genesis.MustCommit(database, triedb)
	blockchain, _ := core.NewBlockChain(database, nil, &genesis, nil, ethash.NewFaker(), vm.Config{}, nil, nil)

	backend := &SimulatedBackend{
		database:   database,
		blockchain: blockchain,
		config:     genesis.Config,
	}

	filterBackend := &filterBackend{database, blockchain, backend}
	backend.filterSystem = filters.NewFilterSystem(filterBackend, filters.Config{})
	backend.events = filters.NewEventSystem(backend.filterSystem, false)

	backend.rollback(blockchain.GetBlock(blockchain.CurrentBlock().Hash(), blockchain.CurrentBlock().Number.Uint64()))
	return backend
}
