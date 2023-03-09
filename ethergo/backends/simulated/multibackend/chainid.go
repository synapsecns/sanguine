package multibackend

import (
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

// NewConfigWithChainID creates a new *params.ChainConfig and changes only the chain id
// everything else is taken from params.AllEthashProtocolChanges.
// we need to do this because params are global.
func NewConfigWithChainID(chainID *big.Int) *params.ChainConfig {
	return &params.ChainConfig{
		ChainID:                       chainID,
		HomesteadBlock:                params.AllEthashProtocolChanges.HomesteadBlock,
		DAOForkBlock:                  params.AllEthashProtocolChanges.DAOForkBlock,
		DAOForkSupport:                params.AllEthashProtocolChanges.DAOForkSupport,
		EIP150Block:                   params.AllEthashProtocolChanges.EIP150Block,
		EIP150Hash:                    params.AllEthashProtocolChanges.EIP150Hash,
		EIP155Block:                   params.AllEthashProtocolChanges.EIP155Block,
		EIP158Block:                   params.AllEthashProtocolChanges.EIP158Block,
		ByzantiumBlock:                params.AllEthashProtocolChanges.ByzantiumBlock,
		ConstantinopleBlock:           params.AllEthashProtocolChanges.ConstantinopleBlock,
		PetersburgBlock:               params.AllEthashProtocolChanges.PetersburgBlock,
		IstanbulBlock:                 params.AllEthashProtocolChanges.IstanbulBlock,
		MuirGlacierBlock:              params.AllEthashProtocolChanges.MuirGlacierBlock,
		BerlinBlock:                   params.AllEthashProtocolChanges.BerlinBlock,
		LondonBlock:                   params.AllEthashProtocolChanges.LondonBlock,
		ArrowGlacierBlock:             params.AllEthashProtocolChanges.ArrowGlacierBlock,
		GrayGlacierBlock:              params.AllEthashProtocolChanges.GrayGlacierBlock,
		MergeNetsplitBlock:            params.AllEthashProtocolChanges.MergeNetsplitBlock,
		ShanghaiTime:                  params.AllEthashProtocolChanges.ShanghaiTime,
		CancunTime:                    params.AllEthashProtocolChanges.CancunTime,
		PragueTime:                    params.AllEthashProtocolChanges.PragueTime,
		TerminalTotalDifficulty:       params.AllEthashProtocolChanges.TerminalTotalDifficulty,
		TerminalTotalDifficultyPassed: params.AllEthashProtocolChanges.TerminalTotalDifficultyPassed,
		Ethash:                        params.AllEthashProtocolChanges.Ethash,
		Clique:                        params.AllEthashProtocolChanges.Clique,
	}
}

// NewSimulatedBackendWithConfig creates a new simulated backend with the given chain id.
func NewSimulatedBackendWithConfig(alloc core.GenesisAlloc, gasLimit uint64, config *params.ChainConfig) *SimulatedBackend {
	database := rawdb.NewMemoryDatabase()

	genesis := core.Genesis{Config: config, GasLimit: gasLimit, Alloc: alloc}
	genesis.MustCommit(database)
	blockchain, _ := core.NewBlockChain(database, nil, &genesis, nil, ethash.NewFaker(), vm.Config{}, nil, nil)

	backend := &SimulatedBackend{
		database:   database,
		blockchain: blockchain,
		config:     genesis.Config,
	}

	filterBackend := &filterBackend{database, blockchain, backend}
	backend.filterSystem = filters.NewFilterSystem(filterBackend, filters.Config{})
	backend.events = filters.NewEventSystem(backend.filterSystem, false)

	header := backend.blockchain.CurrentBlock()
	block := backend.blockchain.GetBlock(header.Hash(), header.Number.Uint64())
	backend.rollback(block)

	return backend
}
