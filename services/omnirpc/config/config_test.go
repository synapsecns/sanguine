package config_test

import (
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
				Checks: gofakeit.Uint16(),
			},
			2: {
				RPCs:   []string{gofakeit.URL(), gofakeit.URL(), gofakeit.URL()},
				Checks: gofakeit.Uint16(),
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
	rpcConf, err := config.UnmarshallConfig([]byte(testYaml))
	Nil(t, err)

	True(t, slices.Contains(rpcConf.Chains[1].RPCs, "https://api.mycryptoapi.com/eth"))
	True(t, slices.Contains(rpcConf.Chains[1].RPCs, "https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet"))
	True(t, slices.Contains(rpcConf.Chains[2].RPCs, "https://node.eggs.cool"))
}

const testYaml = `
chains:
    0:
        rpcs:
            - https://rpc.kardiachain.io/
        confirmations: 1
    1:
        rpcs:
            - https://api.mycryptoapi.com/eth
            - https://rpc.flashbots.net/
            - https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79
            - https://cloudflare-eth.com/
            - https://mainnet-nethermind.blockscout.com/
            - https://nodes.mewapi.io/rpc/eth
            - https://main-rpc.linkpool.io/
            - https://mainnet.eth.cloud.ava.do/
            - https://ethereumnodelight.app.runonflux.io
            - https://rpc.ankr.com/eth
            - https://eth-rpc.gateway.pokt.network
            - https://main-light.eth.linkpool.io
            - https://eth-mainnet.public.blastapi.io
            - http://18.211.207.34:8545
            - https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7
            - wss://eth-mainnet.nodereal.io/ws/v1/1659dfb40aa24bbb8153a677b98064d7
            - https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet
        confirmations: 1
    2:
        rpcs:
            - https://node.eggs.cool
            - https://node.expanse.tech
        confirmations: 1`
