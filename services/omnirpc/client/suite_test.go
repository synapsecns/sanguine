package client_test

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/omnirpc/metadata"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"golang.org/x/sync/errgroup"
)

type TestClientSuite struct {
	*testsuite.TestSuite
	// testBackends contins a list of all test backends
	testBackends []backends.SimulatedTestBackend
	endpoint     string
	client       client.RPCClient
	metrics      metrics.Handler
}

func NewTestClientSuite(tb testing.TB) *TestClientSuite {
	tb.Helper()
	return &TestClientSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestTestClientSuite(t *testing.T) {
	suite.Run(t, NewTestClientSuite(t))
}

func (s *TestClientSuite) SetupSuite() {
	s.TestSuite.SetupSuite()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		s.SetupBackends(s.GetSuiteContext())
	}()

	go func() {
		defer wg.Done()
		s.SetupJaeger()
	}()

	wg.Wait()
}

func (s *TestClientSuite) SetupTest() {
	s.TestSuite.SetupTest()

	s.endpoint = s.getHostURL(testhelper.NewOmnirpcServer(s.GetTestContext(), s.T(), s.testBackends...))
	s.client = client.NewOmnirpcClient(s.endpoint, s.metrics, client.WithCaptureReqRes())
}

func (s *TestClientSuite) SetupJaeger() {
	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	s.metrics, err = metrics.NewByType(s.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	s.Require().Nil(err)
}

// SetupBackends sets up the test backends that are used for the tests. These need to be setup as embedded backends since
// scribe requires rpc addresses, so we employ some paraellism to speed up the test process.
//
// This can either be done per suite or per test. This is done per suite do to the cost of spinning up fake geth nodes.
func (s *TestClientSuite) SetupBackends(ctx context.Context) {
	// let's create 3 mock chains
	chainIDs := []uint64{1, 2, 3}

	// preallocate a slice for testbackends to the size of chainIDs
	// this way we can avoid non-deterministic order + needing to acquire/release a lock
	s.testBackends = make([]backends.SimulatedTestBackend, len(chainIDs))

	// TODO: can we use a waitgroup here instead?
	g, gCtx := errgroup.WithContext(ctx)
	for i, chainID := range chainIDs {
		pos := i           // get position of chain id in array
		chainID := chainID // capture func literal
		g.Go(func() error {
			// we need to use the embedded backend here, because the simulated backend doesn't support rpcs required by scribe
			backend := geth.NewEmbeddedBackendForChainID(ctx, s.T(), new(big.Int).SetUint64(chainID))

			// make sure we mine at least 1 block
			backend.GetFundedAccount(gCtx, big.NewInt(1000000000000000000))
			// add the backend to the list of backends
			s.testBackends[pos] = backend
			return nil
		})
	}

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *TestClientSuite) getHostURL(urlString string) string {
	parsedURL, err := url.Parse(urlString)
	s.Require().NoError(err)

	// Extract the protocol, hostname, and port
	protocol := parsedURL.Scheme
	hostname := parsedURL.Hostname()
	port := parsedURL.Port()

	// If the port is not specified, use the default port for the protocol
	if port == "" {
		if protocol == "http" {
			port = "80"
		} else if protocol == "https" {
			port = "443"
		}
	}

	// Construct the new URL string with just the protocol, hostname, and port
	newURLString := fmt.Sprintf("%s://%s:%s", protocol, hostname, port)
	return newURLString
}
