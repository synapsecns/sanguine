package latency_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/serivces/omnirpc/latency"
	"golang.org/x/sync/errgroup"
	"time"
)

func (r *LatencySuite) TestRPCLatency() {
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

	latencySlice := latency.GetRPCLatency(r.GetTestContext(), time.Second*3, []string{bsc.HTTPEndpoint(), avalanche.HTTPEndpoint()})
	NotEqual(r.T(), latencySlice[0].URL, latencySlice[1].URL)
	for _, latencyData := range latencySlice {
		False(r.T(), latencyData.HasError)
		Nil(r.T(), latencyData.Error)
	}
}
