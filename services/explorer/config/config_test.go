package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	etherMocks "github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/explorer/config"
)

func (c ConfigSuite) TestConfigEncodeDecode() {
	// generate an example config
	chainID := gofakeit.Uint32()

	chain1 := config.ChainConfig{
		ChainID:                  chainID,
		FetchBlockIncrement:      3,
		StartBlock:               0,
		SynapseBridgeAddress:     etherMocks.MockAddress().String(),
		SwapFlashLoanAddresses:   []string{etherMocks.MockAddress().String(), etherMocks.MockAddress().String()},
		StartFromLastBlockStored: false,
	}
	chain2 := config.ChainConfig{
		ChainID:                  chainID + 1,
		FetchBlockIncrement:      3,
		StartBlock:               0,
		SynapseBridgeAddress:     etherMocks.MockAddress().String(),
		SwapFlashLoanAddresses:   []string{etherMocks.MockAddress().String(), etherMocks.MockAddress().String()},
		StartFromLastBlockStored: false,
	}
	chainConfigs := config.ChainConfigs{chain1, chain2}
	testConfig := config.Config{
		Chains:              chainConfigs,
		RefreshRate:         uint(gofakeit.Uint8()),
		ScribeURL:           "http://localhost:8080",
		BridgeConfigAddress: etherMocks.MockAddress().String(),
		BridgeConfigChainID: gofakeit.Uint32(),
	}

	encodedConfig, err := testConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))
	decodedConfig, err := config.DecodeConfig(file.Name())
	Nil(c.T(), err)

	ok, err := decodedConfig.IsValid(c.GetTestContext())
	True(c.T(), ok)
	Nil(c.T(), err)
}
