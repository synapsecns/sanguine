package config_test

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	errorHelper "github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/synapsecns/sanguine/agents/config"
	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestUnmarshallScribeType(t *testing.T) {
	baseConfig := configFixture()

	out, err := roundtripConfig(baseConfig)
	require.Nilf(t, err, "error should be nil: %v", err)

	require.Equal(t, out, baseConfig.Type)
}

func roundtripConfig(cfg config.ScribeConfig) (*config.ScribeConfig, error) {
	encodedCfg, err := cfg.Encode()
	if err != nil {
		return nil, errorHelper.Wrapf(errCouldNotEncode, err.Error())
	}

	var decodedCfg config.ScribeConfig
	err = yaml.Unmarshal(encodedCfg, &decodedCfg)
	if err != nil {
		return nil, errorHelper.Wrapf(errCouldNotDecode, err.Error())
	}

	return &decodedCfg, nil
}

var (
	errCouldNotDecode = errors.New("could not decode config")
	errCouldNotEncode = errors.New("could not encode config")
)

func configFixture() config.ScribeConfig {
	return config.ScribeConfig{
		Type: config.EmbeddedScribeType,
		EmbeddedDBConfig: scribeConfig.DBConfig{
			Type: "mysql",
		},
		EmbeddedScribeConfig: scribeConfig.Config{
			RPCURL: gofakeit.URL(),
		},
		Port: uint(gofakeit.Uint16()),
		URL:  gofakeit.URL(),
	}
}
