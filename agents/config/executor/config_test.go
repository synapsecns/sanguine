package executor_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/config/executor"
	agentsConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
)

func configFixture(c ConfigSuite) executor.Config {
	chainIDA := gofakeit.Uint32()
	chainIDB := chainIDA + 1
	testWallet, err := wallet.FromRandom()
	Nil(c.T(), err)
	return executor.Config{
		DBConfig: config.DBConfig{
			Type:   "sqlite",
			Source: gofakeit.Word(),
		},
		ScribeConfig: config.ScribeConfig{
			Type: "embedded",
			EmbeddedDBConfig: scribeConfig.DBConfig{
				Type: "mysql",
			},
			EmbeddedScribeConfig: scribeConfig.Config{
				RPCURL: gofakeit.URL(),
			},
			Port: uint(gofakeit.Uint16()),
			URL:  gofakeit.URL(),
		},
		Chains: executor.ChainConfigs{
			executor.ChainConfig{
				ChainID:            chainIDA,
				OriginAddress:      gofakeit.Word(),
				DestinationAddress: gofakeit.Word(),
			},
			executor.ChainConfig{
				ChainID:            chainIDB,
				OriginAddress:      gofakeit.Word(),
				DestinationAddress: gofakeit.Word(),
			},
		},
		BaseOmnirpcURL: gofakeit.URL(),
		UnbondedSigner: agentsConfig.SignerConfig{
			Type: agentsConfig.FileType.String(),
			File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
		},
	}
}

func (c ConfigSuite) TestConfigEncodeDecode() {
	testConfig := configFixture(c)

	encodedConfig, err := testConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))
	decodedConfig, err := executor.DecodeConfig(file.Name())
	Nil(c.T(), err)

	ok, err := decodedConfig.IsValid(c.GetTestContext())
	Nil(c.T(), err)
	True(c.T(), ok)
}

func (c ConfigSuite) TestInvalidAttestationInfo() {
	testConfig := configFixture(c)

	ok, err := testConfig.IsValid(c.GetTestContext())
	Nil(c.T(), err)
	True(c.T(), ok)

	testConfig.BaseOmnirpcURL = ""

	ok, err = testConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	NotNil(c.T(), err)
}
