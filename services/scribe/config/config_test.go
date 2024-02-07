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
				Contracts: config.ContractConfigs{
					config.ContractConfig{
						Address:    etherMocks.MockAddress().String(),
						StartBlock: gofakeit.Uint64(),
					},
				},
			},
		},
		RPCURL: gofakeit.URL(),
	}

	encodedConfig, err := testConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))
	decodedConfig, err := config.DecodeConfig(file.Name())
	Nil(c.T(), err)

	ok, err := decodedConfig.IsValid()
	True(c.T(), ok)
	Nil(c.T(), err)
}
