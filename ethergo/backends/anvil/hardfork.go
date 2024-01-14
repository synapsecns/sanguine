package anvil

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core"
	"math/big"
)

// Hardfork indicates which chain should be hardforked by foundry
// see: https://github.com/foundry-rs/foundry/blob/master/anvil/src/hardfork.rs#L94
// note: because of the way foundry does alias being incompatible w/ stringer, latest will have the wrong id when passed as an int
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Hardfork -linecomment
type Hardfork uint8

const (
	// Frontier is the initial release of the Ethereum protocol.
	Frontier Hardfork = iota + 1 // frontier
	// Homestead is the chain config corresponding to the homestead hardfork.
	Homestead // homestead
	// DAO is the chain config corresponding to the DAO hardfork.
	DAO // dao
	// Tangerine is the chain config corresponding to the tangerine hardfork.
	Tangerine // tangerine
	// Spurious is the chain config corresponding to the spurious hardfork.
	Spurious // spuriousdragon
	// Byzantium is the chain config corresponding to the byzantium hardfork.
	Byzantium // byzantium
	// Constantinople is the chain config corresponding to the constantinople hardfork.
	Constantinople // constantinople
	// Petersburg is the chain config corresponding to the petersburg hardfork.
	Petersburg // petersburg
	// Istanbul is the chain config corresponding to the istanbul hardfork.
	Istanbul // istanbul
	// MuirGlacier is the chain config corresponding to the muir glacier hardfork.
	MuirGlacier // muirglacier
	// Berlin is the chain config corresponding to the berlin hardfork.
	Berlin // berlin
	// London is the chain config corresponding to the london hardfork.
	London // london
	// ArrowGlacier is the chain config corresponding to the arrow glacier hardfork.
	ArrowGlacier // arrowglacier
	// GrayGlacier is the chain config corresponding to the glay glacier hardfork.
	GrayGlacier // grayglacier
	// Latest is the chain config corresponding to the latest hardfork.
	// the int of this is not correct (1 to hig), but it is not used in foundry option builder.
	Latest // latest
)

// ToChainConfig converts the hardfork to a chain config.
// it makes the assumption that the blocks start at 0 since we can't easily make assumptions
// about the underlying chain and this is used for tx submission only going forward.
// nolint: cyclop
func (h Hardfork) ToChainConfig(chainID *big.Int) *params.ChainConfig {
	baseConfig := &params.ChainConfig{
		ChainID: core.CopyBigInt(chainID),
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/homestead.md
	if h >= Homestead {
		baseConfig.HomesteadBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/dao-fork.md
	if h >= DAO {
		baseConfig.DAOForkSupport = true
		baseConfig.DAOForkBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/tangerine-whistle.md
	if h >= Tangerine {
		baseConfig.EIP150Block = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/spurious-dragon.md
	if h >= Spurious {
		baseConfig.EIP155Block = big.NewInt(0)
		baseConfig.EIP158Block = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/byzantium.md
	if h >= Byzantium {
		baseConfig.ByzantiumBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/constantinople.md
	if h >= Constantinople {
		baseConfig.ConstantinopleBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/petersburg.md
	if h >= Petersburg {
		baseConfig.PetersburgBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/istanbul.md
	if h >= Istanbul {
		baseConfig.IstanbulBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/muir-glacier.md
	if h >= MuirGlacier {
		baseConfig.MuirGlacierBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/berlin.md
	if h >= Berlin {
		baseConfig.BerlinBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/london.md
	if h >= London {
		baseConfig.LondonBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/arrow-glacier.md
	if h >= ArrowGlacier {
		baseConfig.ArrowGlacierBlock = big.NewInt(0)
	}

	// https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/gray-glacier.md
	if h >= GrayGlacier {
		baseConfig.GrayGlacierBlock = big.NewInt(0)
	}
	return baseConfig
}

// allHardforks is a list of all hardforks.
var allHardforks = []Hardfork{
	Frontier,
	Homestead,
	DAO,
	Tangerine,
	Spurious,
	Byzantium,
	Constantinople,
	Petersburg,
	Istanbul,
	MuirGlacier,
	Berlin,
	London,
	ArrowGlacier,
	GrayGlacier,
	Latest,
}

func init() {
	if len(_Hardfork_index)-1 != len(allHardforks) {
		panic("not all hardforks have been added to the stringer")
	}
}
