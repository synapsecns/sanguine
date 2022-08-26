package config_test

import (
	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/serivces/omnirpc/config"
	"golang.org/x/exp/slices"
	"testing"
)

func TestUnmarshallMarshall(t *testing.T) {
	rpcMap, err := config.UnmarshallConfig(testYaml)
	Nil(t, err)

	True(t, slices.Contains(rpcMap.ChainID(1), "https://1.com"))
	True(t, slices.Contains(rpcMap.ChainID(1), "https://1.test.com"))
	True(t, slices.Contains(rpcMap.ChainID(2), "https://2.test.com"))

	newMap, err := config.UnmarshallConfig(config.MarshallFromMap(rpcMap))
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
