package rpcmap_test

import (
	"github.com/jarcoal/httpmock"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/serivces/omnirpc/rpcmap"
	"golang.org/x/exp/slices"
	"net/http"
	"testing"
)

func (r *RPCSuite) TestGetRPCMap() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, rpcmap.RPCMapURL, httpmock.NewStringResponder(http.StatusOK, testData))

	res, err := rpcmap.GetRPCMap(r.GetTestContext())
	Nil(r.T(), err)

	True(r.T(), slices.Contains(res[1], "https://api.mycryptoapi.com/eth"))
}

func TestGetChainRPCS(t *testing.T) {
	rpcMap, err := rpcmap.ParseRPCMap([]byte(testData))
	Nil(t, err)

	True(t, slices.Contains(rpcMap[1], "https://api.mycryptoapi.com/eth"))
	True(t, slices.Contains(rpcMap[1], "https://cloudflare-eth.com/"))
	True(t, slices.Contains(rpcMap[2], "https://node.eggs.cool"))
}

// first two rpcs from https://raw.githubusercontent.com/DefiLlama/chainlist/main/constants/extraRpcs.json
const testData = `
{
  "1": {
    "rpcs": [
      "https://api.mycryptoapi.com/eth",
      "https://rpc.flashbots.net/",
      "https://eth-mainnet.gateway.pokt.network/v1/5f3453978e354ab992c4da79",
      "https://cloudflare-eth.com/",
      "https://mainnet-nethermind.blockscout.com/",
      "https://nodes.mewapi.io/rpc/eth",
      "https://main-rpc.linkpool.io/",
      "https://mainnet.eth.cloud.ava.do/",
      "https://ethereumnodelight.app.runonflux.io",
      "https://rpc.ankr.com/eth",
      "https://eth-rpc.gateway.pokt.network",
      "https://main-light.eth.linkpool.io",
      "https://eth-mainnet.public.blastapi.io",
      "http://18.211.207.34:8545",
      "https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7",
      "wss://eth-mainnet.nodereal.io/ws/v1/1659dfb40aa24bbb8153a677b98064d7",
      "https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet"
    ]
  },
  "2": {
    "rpcs": [
      "https://node.eggs.cool",
      "https://node.expanse.tech"
    ]
  }
}
`
