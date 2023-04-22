package anvil_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"math/big"
	"testing"
)

func TestHardfork(t *testing.T) {
	hardforks := anvil.AllHardforks()
	for _, hardfork := range hardforks {
		chainCfg := hardfork.ToChainConfig(new(big.Int).SetUint64(gofakeit.Uint64()))
		NoError(t, chainCfg.CheckConfigForkOrder())
	}
}

// nolint: dupl, thelper
func TestToChainConfig(t *testing.T) {
	chainID := big.NewInt(1)

	tests := []struct {
		name      string
		hardfork  anvil.Hardfork
		checkFunc func(t *testing.T, chainConfig *params.ChainConfig)
	}{
		{
			name:     "Frontier",
			hardfork: anvil.Frontier,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				Nil(t, chainConfig.HomesteadBlock)
			},
		},
		{
			name:     "Homestead",
			hardfork: anvil.Homestead,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				False(t, chainConfig.IsDAOFork(big.NewInt(0)))
			},
		},
		{
			name:     "DAO",
			hardfork: anvil.DAO,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				False(t, chainConfig.IsEIP150(big.NewInt(0)))
			},
		},
		{
			name:     "Tangerine",
			hardfork: anvil.Tangerine,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				False(t, chainConfig.IsEIP155(big.NewInt(0)))
			},
		},
		{
			name:     "Spurious",
			hardfork: anvil.Spurious,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				False(t, chainConfig.IsByzantium(big.NewInt(0)))
			},
		},
		{
			name:     "Byzantium",
			hardfork: anvil.Byzantium,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				False(t, chainConfig.IsConstantinople(big.NewInt(0)))
			},
		},
		{
			name:     "Constantinople",
			hardfork: anvil.Constantinople,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				True(t, chainConfig.IsConstantinople(big.NewInt(0)))
				// IsPetersburg returns whether num is either
				// - equal to or greater than the PetersburgBlock fork block,
				// - OR is nil, and Constantinople is active
				True(t, chainConfig.IsPetersburg(big.NewInt(0)))
			},
		},
		{
			name:     "Petersburg",
			hardfork: anvil.Petersburg,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				True(t, chainConfig.IsConstantinople(big.NewInt(0)))
				True(t, chainConfig.IsPetersburg(big.NewInt(0)))
				False(t, chainConfig.IsIstanbul(big.NewInt(0)))
			},
		},
		{
			name:     "Istanbul",
			hardfork: anvil.Istanbul,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				True(t, chainConfig.IsConstantinople(big.NewInt(0)))
				True(t, chainConfig.IsPetersburg(big.NewInt(0)))
				True(t, chainConfig.IsIstanbul(big.NewInt(0)))
				False(t, chainConfig.IsMuirGlacier(big.NewInt(0)))
			},
		},
		{
			name:     "MuirGlacier",
			hardfork: anvil.MuirGlacier,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				True(t, chainConfig.IsConstantinople(big.NewInt(0)))
				True(t, chainConfig.IsPetersburg(big.NewInt(0)))
				True(t, chainConfig.IsIstanbul(big.NewInt(0)))
				True(t, chainConfig.IsMuirGlacier(big.NewInt(0)))
				False(t, chainConfig.IsBerlin(big.NewInt(0)))
			},
		},
		{
			name:     "Berlin",
			hardfork: anvil.Berlin,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				True(t, chainConfig.IsConstantinople(big.NewInt(0)))
				True(t, chainConfig.IsPetersburg(big.NewInt(0)))
				True(t, chainConfig.IsIstanbul(big.NewInt(0)))
				True(t, chainConfig.IsMuirGlacier(big.NewInt(0)))
				True(t, chainConfig.IsBerlin(big.NewInt(0)))
				False(t, chainConfig.IsLondon(big.NewInt(0)))
			},
		},
		{
			name:     "London",
			hardfork: anvil.London,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				True(t, chainConfig.IsConstantinople(big.NewInt(0)))
				True(t, chainConfig.IsPetersburg(big.NewInt(0)))
				True(t, chainConfig.IsIstanbul(big.NewInt(0)))
				True(t, chainConfig.IsMuirGlacier(big.NewInt(0)))
				True(t, chainConfig.IsBerlin(big.NewInt(0)))
				True(t, chainConfig.IsLondon(big.NewInt(0)))
				False(t, chainConfig.IsArrowGlacier(big.NewInt(0)))
			},
		},
		{
			name:     "ArrowGlacier",
			hardfork: anvil.ArrowGlacier,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				True(t, chainConfig.IsConstantinople(big.NewInt(0)))
				True(t, chainConfig.IsPetersburg(big.NewInt(0)))
				True(t, chainConfig.IsIstanbul(big.NewInt(0)))
				True(t, chainConfig.IsMuirGlacier(big.NewInt(0)))
				True(t, chainConfig.IsBerlin(big.NewInt(0)))
				True(t, chainConfig.IsLondon(big.NewInt(0)))
				True(t, chainConfig.IsArrowGlacier(big.NewInt(0)))
				False(t, chainConfig.IsGrayGlacier(big.NewInt(0)))
			},
		},
		{
			name:     "GrayGlacier",
			hardfork: anvil.GrayGlacier,
			checkFunc: func(t *testing.T, chainConfig *params.ChainConfig) {
				NotNil(t, chainConfig.HomesteadBlock)
				True(t, chainConfig.IsHomestead(big.NewInt(0)))
				True(t, chainConfig.IsDAOFork(big.NewInt(0)))
				True(t, chainConfig.IsEIP150(big.NewInt(0)))
				True(t, chainConfig.IsEIP155(big.NewInt(0)))
				True(t, chainConfig.IsEIP158(big.NewInt(0)))
				True(t, chainConfig.IsByzantium(big.NewInt(0)))
				True(t, chainConfig.IsConstantinople(big.NewInt(0)))
				True(t, chainConfig.IsPetersburg(big.NewInt(0)))
				True(t, chainConfig.IsIstanbul(big.NewInt(0)))
				True(t, chainConfig.IsMuirGlacier(big.NewInt(0)))
				True(t, chainConfig.IsBerlin(big.NewInt(0)))
				True(t, chainConfig.IsLondon(big.NewInt(0)))
				True(t, chainConfig.IsArrowGlacier(big.NewInt(0)))
				True(t, chainConfig.IsGrayGlacier(big.NewInt(0)))
			},
		},
		// Add other hardfork test cases here
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			chainConfig := test.hardfork.ToChainConfig(chainID)
			Equal(t, chainID, chainConfig.ChainID)

			test.checkFunc(t, chainConfig)
		})
	}
}
