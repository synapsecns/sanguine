package testhelper

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"github.com/synapsecns/sanguine/services/omnirpc/metadata"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
)

func makeConfig(backends []backends.SimulatedTestBackend, clientType omniHTTP.ClientType) config.Config {
	chains := make(map[uint32]config.ChainConfig)

	for _, backend := range backends {
		chains[uint32(backend.GetChainID())] = config.ChainConfig{
			RPCs:   []string{backend.RPCAddress()},
			Checks: 1,
		}
	}

	return config.Config{
		Chains:          chains,
		Port:            uint16(freeport.GetPort()),
		RefreshInterval: 0,
		ClientType:      clientType.String(),
	}
}

// NewOmnirpcServer creates a new omnirpc server with all the test backends passed in.
// since these are all mocked geth instances, these are single confirmation instances only.
// a string is returned with the base url for the omnirpc server.
//
// context is respected and the server will be killed when the context is done.
func NewOmnirpcServer(ctx context.Context, tb testing.TB, backends ...backends.SimulatedTestBackend) string {
	tb.Helper()

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	// TODO: make this optional everywhere before merging
	if useMetrics && false {
		localmetrics.SetupTestJaeger(ctx, tb)
		metricsHandler = metrics.Jaeger
	}

	handler, err := metrics.NewByType(ctx, metadata.BuildInfo(), metricsHandler)
	assert.Nil(tb, err)

	server := proxy.NewProxy(makeConfig(backends, omniHTTP.FastHTTP), handler)

	go func() {
		server.Run(ctx)
	}()

	baseHost := fmt.Sprintf("http://0.0.0.0:%d", server.Port())
	healthCheck := fmt.Sprintf("%s%s", baseHost, ginhelper.HealthCheck)

	// wait for server to start
	testsuite.Eventually(ctx, tb, func() bool {
		select {
		case <-ctx.Done():
			tb.Error(ctx.Err())
		default:
			// see below
		}

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, healthCheck, nil)
		assert.Nil(tb, err)

		res, err := http.DefaultClient.Do(request)
		if err == nil {
			defer func() {
				_ = res.Body.Close()
			}()
			return true
		}

		return false
	})

	return baseHost
}

// GetURL gets the url for a given backend given the base host.
func GetURL(baseHost string, backend backends.SimulatedTestBackend) string {
	return fmt.Sprintf("%s/rpc/%d", baseHost, backend.GetChainID())
}
