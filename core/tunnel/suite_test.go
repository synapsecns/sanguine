package tunnel_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/retry"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/core/testsuite"
	"net/http"
	"testing"
	"time"
)

// TunnelSuite defines the basic test suite.
type TunnelSuite struct {
	*testsuite.TestSuite
	// testServer is the url of the local server to test on.
	testServer string
	logger     *log.ZapEventLogger
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTunnelSuite(tb testing.TB) *TunnelSuite {
	tb.Helper()
	return &TunnelSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (n *TunnelSuite) SetupTest() {
	n.TestSuite.SetupTest()
	n.logger = log.Logger(fmt.Sprintf("test-%d-logger", n.GetTestID()))
	n.startServer(n.GetTestContext())
}

// startServer starts the test server and sets the testServer field.
func (n *TunnelSuite) startServer(ctx context.Context) {
	testServer := ginhelper.New(n.logger)
	freePort, err := freeport.GetFreePort()
	n.Require().NoError(err)

	n.testServer = fmt.Sprintf("http://localhost:%d", freePort)

	go func() {
		connection := baseServer.Server{}
		err = connection.ListenAndServe(ctx, fmt.Sprintf(":%d", freePort), testServer)
		// we expect context cancellation errors at the end of the test
		if !errors.Is(err, context.Canceled) {
			n.Require().NoError(err)
		}
	}()

	// make sure the server is running

	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", n.testServer, ginhelper.HealthCheck), nil)
		if err != nil {
			return fmt.Errorf("could not create request: %w", err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("could not send request: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		if resp.Body != nil {
			_ = resp.Body.Close()
		}

		return nil
	}, retry.WithMin(time.Millisecond), retry.WithMax(time.Second), retry.WithMaxAttemptTime(time.Second*30))

	n.Require().NoError(err)
}

func TestTunnelSuite(t *testing.T) {
	suite.Run(t, NewTunnelSuite(t))
}
