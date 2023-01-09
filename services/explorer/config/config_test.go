package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	etherMocks "github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"math/big"
)

func (c ConfigSuite) TestConfigEncodeDecode() {
	// Create the chain configs
	chain1 := config.ChainConfig{
		ChainID:             gofakeit.Uint32(),
		FetchBlockIncrement: 100,
		MaxGoroutines:       10,
		Contracts:           []config.ContractConfig{makeContractConfig(), makeContractConfig()},
	}
	chain2 := config.ChainConfig{
		ChainID:             gofakeit.Uint32(),
		FetchBlockIncrement: 100,
		MaxGoroutines:       10,
		Contracts:           []config.ContractConfig{makeContractConfig(), makeContractConfig()},
	}

	// Put all the chain configs together
	chainConfigs := config.ChainConfigs{chain1, chain2}

	// Put everything into one Config
	testConfig := config.Config{
		RefreshRate:         int(gofakeit.Uint8()),
		ScribeURL:           gofakeit.URL(),
		RPCURL:              gofakeit.URL(),
		BridgeConfigAddress: etherMocks.MockAddress().String(),
		BridgeConfigChainID: gofakeit.Uint32(),
		Chains:              chainConfigs,
	}

	encodedConfig, err := testConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))
	decodedConfig, err := config.DecodeConfig(file.Name())
	Nil(c.T(), err)

	// Check the validity of the decoded config
	ok, err := decodedConfig.IsValid(c.GetTestContext())
	True(c.T(), ok)
	Nil(c.T(), err)
}

func makeContractConfig() config.ContractConfig {
	return config.ContractConfig{
		ContractType: gofakeit.UUID(),
		Address:      common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		StartBlock:   gofakeit.Int64(),
	}
}
