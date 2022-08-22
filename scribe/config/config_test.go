package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/scribe/config"
	"github.com/synapsecns/synapse-node/testutils/utils"
)

func (c *ConfigSuite) TestConfigEncodeDecode() {
	// generate an example config
	chainID := gofakeit.Uint32()
	testConfig := config.Config{
		Chains: config.ChainConfigs{
			config.ChainConfig{
				ChainID:               chainID,
				RPCUrl:                gofakeit.URL(),
				ConfirmationThreshold: gofakeit.Uint32(),
				Contracts: config.ContractConfigs{
					"a": config.ContractConfig{
						Address:    utils.NewMockAddress().String(),
						StartBlock: gofakeit.Uint64(),
					},
					"b": config.ContractConfig{
						Address:    utils.NewMockAddress().String(),
						StartBlock: gofakeit.Uint64(),
					},
				},
			},
			config.ChainConfig{
				ChainID:               chainID + 1,
				RPCUrl:                gofakeit.URL(),
				ConfirmationThreshold: gofakeit.Uint32(),
				Contracts: config.ContractConfigs{
					"a": config.ContractConfig{
						Address:    utils.NewMockAddress().String(),
						StartBlock: gofakeit.Uint64(),
					},
				},
			},
		},
	}

	encodedConfig, err := testConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", encodedConfig)
	decodedConfig, err := config.DecodeConfig(file.Name())
	Nil(c.T(), err)

	ok, err := decodedConfig.IsValid(c.GetTestContext())
	True(c.T(), ok)
	Nil(c.T(), err)
}
