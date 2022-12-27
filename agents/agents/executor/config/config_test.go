package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/config"
)

func configFixture() config.Config {
	chainIDA := gofakeit.Uint32()
	chainIDB := chainIDA + 1
	return config.Config{
		Chains: config.ChainConfigs{
			config.ChainConfig{
				ChainID:            chainIDA,
				OriginAddress:      gofakeit.Word(),
				DestinationAddress: gofakeit.Word(),
			},
			config.ChainConfig{
				ChainID:            chainIDB,
				OriginAddress:      gofakeit.Word(),
				DestinationAddress: gofakeit.Word(),
			},
		},
<<<<<<< HEAD
		RPCURL: gofakeit.URL(),
=======
		BaseOmnirpcURL: gofakeit.URL(),
>>>>>>> master
	}
}

func (c ConfigSuite) TestConfigEncodeDecode() {
	testConfig := configFixture()

	encodedConfig, err := testConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))
	decodedConfig, err := config.DecodeConfig(file.Name())
	Nil(c.T(), err)

	ok, err := decodedConfig.IsValid(c.GetTestContext())
	Nil(c.T(), err)
	True(c.T(), ok)
}

func (c ConfigSuite) TestInvalidAttestationInfo() {
	testConfig := configFixture()

	ok, err := testConfig.IsValid(c.GetTestContext())
	Nil(c.T(), err)
	True(c.T(), ok)

<<<<<<< HEAD
	testConfig.RPCURL = ""
=======
	testConfig.BaseOmnirpcURL = ""
>>>>>>> master

	ok, err = testConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	NotNil(c.T(), err)
}
