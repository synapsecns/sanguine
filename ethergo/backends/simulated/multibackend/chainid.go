package multibackend

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	util "github.com/synapsecns/sanguine/core"
)

// NewConfigWithChainID creates a new *params.ChainConfig and changes only the chain id
// everything else is taken from params.AllEthashProtocolChanges.
// we need to do this because params are global.
func NewConfigWithChainID(chainID *big.Int) *params.ChainConfig {
	return &params.ChainConfig{
		ChainID:                 util.CopyBigInt(chainID),
		HomesteadBlock:          util.CopyBigInt(params.AllEthashProtocolChanges.HomesteadBlock),
		DAOForkBlock:            util.CopyBigInt(params.AllEthashProtocolChanges.DAOForkBlock),
		DAOForkSupport:          params.AllEthashProtocolChanges.DAOForkSupport,
		EIP150Block:             util.CopyBigInt(params.AllEthashProtocolChanges.EIP150Block),
		EIP155Block:             util.CopyBigInt(params.AllEthashProtocolChanges.EIP155Block),
		EIP158Block:             util.CopyBigInt(params.AllEthashProtocolChanges.EIP158Block),
		ByzantiumBlock:          util.CopyBigInt(params.AllEthashProtocolChanges.ByzantiumBlock),
		ConstantinopleBlock:     util.CopyBigInt(params.AllEthashProtocolChanges.ConstantinopleBlock),
		PetersburgBlock:         util.CopyBigInt(params.AllEthashProtocolChanges.PetersburgBlock),
		IstanbulBlock:           util.CopyBigInt(params.AllEthashProtocolChanges.IstanbulBlock),
		MuirGlacierBlock:        util.CopyBigInt(params.AllEthashProtocolChanges.MuirGlacierBlock),
		BerlinBlock:             util.CopyBigInt(params.AllEthashProtocolChanges.BerlinBlock),
		LondonBlock:             util.CopyBigInt(params.AllEthashProtocolChanges.LondonBlock),
		ArrowGlacierBlock:       util.CopyBigInt(params.AllEthashProtocolChanges.ArrowGlacierBlock),
		GrayGlacierBlock:        util.CopyBigInt(params.AllEthashProtocolChanges.GrayGlacierBlock),
		MergeNetsplitBlock:      util.CopyBigInt(params.AllEthashProtocolChanges.MergeNetsplitBlock),
		ShanghaiTime:            util.CopyPointer(params.AllEthashProtocolChanges.ShanghaiTime),
		CancunTime:              util.CopyPointer(params.AllEthashProtocolChanges.CancunTime),
		PragueTime:              util.CopyPointer(params.AllEthashProtocolChanges.PragueTime),
		VerkleTime:              util.CopyPointer(params.AllEthashProtocolChanges.VerkleTime),
		TerminalTotalDifficulty: util.CopyBigInt(params.AllEthashProtocolChanges.TerminalTotalDifficulty),
		Ethash:                  util.CopyPointer(params.AllEthashProtocolChanges.Ethash),
		Clique:                  util.CopyPointer(params.AllEthashProtocolChanges.Clique),
	}
}

// NewSimulatedBackendWithConfig creates a new simulated backend with the given chain id.
func NewSimulatedBackendWithConfig(alloc core.GenesisAlloc, gasLimit uint64, config *params.ChainConfig) *SimulatedBackend {
	genesis := &core.Genesis{
		Config:     config,
		GasLimit:   gasLimit,
		Alloc:      alloc,
		Difficulty: big.NewInt(0),
	}

	// Configure the simulated backend with custom chain config
	backend := simulated.NewBackend(alloc,
		simulated.WithBlockGasLimit(gasLimit),
		func(nodeConf *node.Config, ethConf *ethconfig.Config) {
			ethConf.Genesis = genesis
		},
	)
	return &SimulatedBackend{Backend: backend}
}
