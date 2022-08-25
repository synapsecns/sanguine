package internal_test

import (
	"github.com/jarcoal/httpmock"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/tools/rpc/internal"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
	"net/http"
	"testing"
	"time"
)

func (r *RPCSuite) TestRPCLatency() {
	var bsc, avalanche *geth.Backend
	g, _ := errgroup.WithContext(r.GetTestContext())
	g.Go(func() error {
		bsc = preset.GetBSCTestnet().Geth(r.GetTestContext(), r.T())
		return nil
	})
	g.Go(func() error {
		avalanche = preset.GetAvalancheLocal().Geth(r.GetTestContext(), r.T())
		return nil
	})
	Nil(r.T(), g.Wait())

	latencySlice := internal.GetRPCLatency(r.GetTestContext(), time.Second*3, []string{bsc.HTTPEndpoint(), avalanche.HTTPEndpoint()})
	NotEqual(r.T(), latencySlice[0].URL, latencySlice[1].URL)
	for _, latencyData := range latencySlice {
		False(r.T(), latencyData.HasError)
		Nil(r.T(), latencyData.Error)
	}
}

func (r *RPCSuite) TestGetRPCMap() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, internal.RPCMapURL, httpmock.NewStringResponder(http.StatusOK, testData))

	res, err := internal.GetRPCMap(r.GetTestContext())
	Nil(r.T(), err)

	True(r.T(), slices.Contains(res[1], "https://api.mycryptoapi.com/eth"))
}

func TestGetChainRPCS(t *testing.T) {
	rpcMap, err := internal.ParseRPCMap([]byte(testData))
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
