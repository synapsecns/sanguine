package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	"golang.org/x/exp/slices"
	"testing"
)

func TestConfig(t *testing.T) {
	testConfig := config.Config{
		Chains: map[uint32]config.ChainConfig{
			1: {
				RPCs:   []string{gofakeit.URL(), gofakeit.URL(), gofakeit.URL()},
				Checks: int(gofakeit.Uint32()),
			},
			2: {
				RPCs:   []string{gofakeit.URL(), gofakeit.URL(), gofakeit.URL()},
				Checks: int(gofakeit.Uint32()),
			},
		},
		Port:            gofakeit.Uint16(),
		RefreshInterval: gofakeit.Second(),
	}

	out, err := testConfig.Marshall()
	Nil(t, err)

	unmarshalledConfig, err := config.UnmarshallConfig(out)
	Nil(t, err)

	Equal(t, testConfig, unmarshalledConfig)
}

func TestUnmarshallMarshall(t *testing.T) {
	rpcMap, err := config.UnmarshallRPCMap(testYaml)
	Nil(t, err)

	True(t, slices.Contains(rpcMap.ChainID(1), "https://1.com"))
	True(t, slices.Contains(rpcMap.ChainID(1), "https://1.test.com"))
	True(t, slices.Contains(rpcMap.ChainID(2), "https://2.test.com"))

	newMap, err := config.UnmarshallRPCMap(config.MarshallFromMap(rpcMap))
	Nil(t, err)

	Equal(t, rpcMap.RawMap(), newMap.RawMap())
}

func TestFileUnmarshall(t *testing.T) {
	resConfig := filet.TmpFile(t, "", testYaml)
	rpcMap, err := config.UnmarshallConfigFromFile(resConfig.Name())
	Nil(t, err)

	True(t, slices.Contains(rpcMap.ChainID(2), "https://2.com"))
}

const testYaml = `
---
1:
  - "https://1.com"
  - "https://1.test.com"
2:
  - "https://2.com"
  - "https://2.test.com"`
