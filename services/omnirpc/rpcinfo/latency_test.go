package rpcinfo_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcinfo"
	"golang.org/x/sync/errgroup"
	"net/http"
	"net/http/httptest"
	"sync"
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

	latencySlice := rpcinfo.GetRPCLatency(r.GetTestContext(), time.Second*3, []string{bsc.HTTPEndpoint(), avalanche.HTTPEndpoint()}, metrics.NewNullHandler())
	NotEqual(r.T(), latencySlice[0].URL, latencySlice[1].URL)
	for _, latencyData := range latencySlice {
		False(r.T(), latencyData.HasError)
		Nil(r.T(), latencyData.Error)
	}
}

var statusCodes = []int{
	http.StatusBadRequest,
	http.StatusUnauthorized,
	http.StatusPaymentRequired,
	http.StatusNotFound,
	http.StatusMethodNotAllowed,
	http.StatusNotAcceptable,
	http.StatusRequestTimeout,
	http.StatusGone,
	http.StatusLengthRequired,
	http.StatusPreconditionRequired,
	http.StatusPreconditionFailed,
	http.StatusRequestEntityTooLarge,
	http.StatusRequestEntityTooLarge,
	// just for fun
	http.StatusTeapot,
	http.StatusMisdirectedRequest,
	http.StatusTooManyRequests,
	http.StatusRequestHeaderFieldsTooLarge,
	http.StatusInternalServerError,
	http.StatusNotImplemented,
	http.StatusGatewayTimeout,
	http.StatusBadGateway,
	http.StatusBadRequest,
	http.StatusLoopDetected,
	http.StatusNotExtended,
	http.StatusNetworkAuthenticationRequired,
}

func (r *LatencySuite) TestGetLatencyError() {
	ctx, cancel := context.WithTimeout(r.GetTestContext(), time.Minute)
	defer cancel()

	wg := &sync.WaitGroup{}

	for _, status := range statusCodes {
		// capture func literal
		status := status

		wg.Add(1)

		go func() {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
			}))

			res := rpcinfo.GetLatency(ctx, server.URL)

			// there's an error, these should be nil
			Equal(r.T(), res.Latency, time.Duration(0))
			Equal(r.T(), res.BlockAge, time.Duration(0))

			server.Close()
			wg.Done()
		}()
	}

	select {
	case <-wrapWait(wg):
		// all good
	case <-ctx.Done():
		// context canceled
		r.T().Fail()
	}
}

// helper function to allow using WaitGroup in a select
// see: https://stackoverflow.com/a/70050009
func wrapWait(wg *sync.WaitGroup) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		wg.Wait()
		out <- struct{}{}
	}()
	return out
}
