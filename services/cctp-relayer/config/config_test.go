package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	ethConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
)

func configFixture(c ConfigSuite) config.Config {
	chainIDA := gofakeit.Uint32()
	chainIDB := chainIDA + 1
	testWallet, err := wallet.FromRandom()
	Nil(c.T(), err)
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
		BaseOmnirpcURL: gofakeit.URL(),
		UnbondedSigner: ethConfig.SignerConfig{
			Type: ethConfig.FileType.String(),
			File: filet.TmpFile(c.T(), "", testWallet.PrivateKeyHex()).Name(),
		},
		EmbeddedScribeConfig: scribeConfig.Config{
			RPCURL: gofakeit.URL(),
		},
		HttpBackoffInitialIntervalMs: int(gofakeit.Int64()),
		HttpBackoffMaxElapsedTimeMs:  int(gofakeit.Int64()),
	}
}

func (c ConfigSuite) TestConfigEncodeDecode() {
	testConfig := configFixture(c)

	encodedConfig, err := testConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))
	decodedConfig, err := config.DecodeConfig(file.Name())
	Nil(c.T(), err)

	ok, err := decodedConfig.IsValid(c.GetTestContext())
	Nil(c.T(), err)
	True(c.T(), ok)
}
