package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	etherMocks "github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/scribe/config"
)

func (c ConfigSuite) TestConfigEncodeDecode() {
	// generate an example config
	chainID := gofakeit.Uint32()
	testConfig := config.Config{
		Chains: config.ChainConfigs{
			config.ChainConfig{
				ChainID: chainID,
				RPCUrl:  gofakeit.URL(),
				Contracts: config.ContractConfigs{
					config.ContractConfig{
						Address:    etherMocks.MockAddress().String(),
						StartBlock: gofakeit.Uint64(),
					},
					config.ContractConfig{
						Address:    etherMocks.MockAddress().String(),
						StartBlock: gofakeit.Uint64(),
					},
				},
			},
			config.ChainConfig{
				ChainID: chainID + 1,
				RPCUrl:  gofakeit.URL(),
				Contracts: config.ContractConfigs{
					config.ContractConfig{
						Address:    etherMocks.MockAddress().String(),
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
