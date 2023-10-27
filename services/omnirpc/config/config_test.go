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
				RPCs: []config.RPCConfig{{
					RPC:     gofakeit.URL(),
					RPCType: "stable",
				},
					{
						RPC:     gofakeit.URL(),
						RPCType: "stable",
					},
					{
						RPC:     gofakeit.URL(),
						RPCType: "stable",
					}},
				Checks: gofakeit.Uint16(),
			},
			2: {
				RPCs: []config.RPCConfig{{
					RPC:     gofakeit.URL(),
					RPCType: "stable",
				},
					{
						RPC:     gofakeit.URL(),
						RPCType: "stable",
					},
					{
						RPC:     gofakeit.URL(),
						RPCType: "stable",
					}},
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

	True(t, slices.Contains(rpcConf.Chains[1].RPCs, config.RPCConfig{
		RPC:     "https://api.mycryptoapi.com/eth",
		RPCType: "stable",
	}))
	True(t, slices.Contains(rpcConf.Chains[1].RPCs,
		config.RPCConfig{
			RPC:     "https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet",
			RPCType: "auxiliary",
		}))
	True(t, slices.Contains(rpcConf.Chains[2].RPCs,
		config.RPCConfig{
			RPC:     "https://node.eggs.cool",
			RPCType: "stable",
		}))
}

const testYaml = `
chains:
  0:
    confirmations: 1
    rpcs:
      - rpc: 'https://rpc.kardiachain.io/'
        rpc_type: stable
  1:
    confirmations: 1
    rpcs:
      - rpc: 'https://api.mycryptoapi.com/eth'
        rpc_type: stable
      - rpc: 'https://rpc.flashbots.net/'
        rpc_type: stable
      - rpc: 'https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79'
        rpc_type: stable
      - rpc: 'https://cloudflare-eth.com/'
        rpc_type: stable
      - rpc: 'https://mainnet-nethermind.blockscout.com/'
        rpc_type: stable
      - rpc: 'https://nodes.mewapi.io/rpc/eth'
        rpc_type: stable
      - rpc: 'https://main-rpc.linkpool.io/'
        rpc_type: stable
      - rpc: 'https://mainnet.eth.cloud.ava.do/'
        rpc_type: stable
      - rpc: 'https://ethereumnodelight.app.runonflux.io'
        rpc_type: stable
      - rpc: 'https://rpc.ankr.com/eth'
        rpc_type: stable
      - rpc: 'https://eth-rpc.gateway.pokt.network'
        rpc_type: stable
      - rpc: 'https://main-light.eth.linkpool.io'
        rpc_type: stable
      - rpc: 'https://eth-mainnet.public.blastapi.io'
        rpc_type: stable
      - rpc: 'http://18.211.207.34:8545'
        rpc_type: stable
      - rpc: 'https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7'
        rpc_type: stable
      - rpc: 'wss://eth-mainnet.nodereal.io/ws/v1/1659dfb40aa24bbb8153a677b98064d7'
        rpc_type: stable
      - rpc: 'https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet'
        rpc_type: auxiliary
  2:
    confirmations: 1
    rpcs:
      - rpc: 'https://node.eggs.cool'
        rpc_type: stable
      - rpc: 'https://node.expanse.tech'
        rpc_type: stable
`
