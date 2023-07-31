package tunnel_test

import (
	"context"
	"errors"
	"fmt"
	moessh "github.com/fasmide/remotemoe/ssh"
	"path"

	moehttp "github.com/fasmide/remotemoe/http"
	"github.com/fasmide/remotemoe/routertwo"
	"github.com/fasmide/remotemoe/services"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/retry"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/synapsecns/sanguine/core/testsuite"
	"net/http"
	"os"
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
	}, retry.WithMin(time.Millisecond), retry.WithMax(time.Second), retry.WithMaxAttemptsTime(time.Second*30))

	n.Require().NoError(err)
}

// TODO: this is not context safe
func (n *TunnelSuite) startMoe() {
	routerData := "routerdata"

	if os.Getenv("STATE_DIRECTORY") != "" {
		routerData = path.Join(os.Getenv("STATE_DIRECTORY"), "routerdata")
	}

	err := os.Mkdir(routerData, 0700)

	// we are not going to be stopping on ErrExists errors
	if errors.Is(err, os.ErrExist) {
		err = nil
	}

	n.Require().NoError(err)

	router, err := routertwo.NewRouter(routerData)
	n.Require().NoError(err)

	proxy := &moehttp.Proxy{}
	proxy.Initialize(router)

	server, err := moehttp.NewServer(router.Exists)
	if err != nil {
		panic(err)
	}

	server.Handler = proxy

	services.Serve("http", server)
	services.ServeTLS("https", server)

	sshConfig, err := moessh.DefaultConfig()
	n.Require().NoError(err)

	sshServer := &moessh.Server{Config: sshConfig, Router: router}

	services.Serve("ssh", sshServer)

	// we shall be dealing with shutting down in the future :)
}

func TestTunnelSuite(t *testing.T) {
	suite.Run(t, NewTunnelSuite(t))
}
