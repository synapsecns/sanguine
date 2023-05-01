package tunnel_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/tunnel"
	"github.com/synapsecns/sanguine/core/tunnel/internal"
	"net"
	"net/http"
	"time"
)

// this module is currently broken due to an issue with moe.
// this is caused by two issues:
// 1. Moe returns localhost as the hostname: https://github.com/fasmide/remotemoe/blob/13a9ba0f5ddadffdf4fb395ebff7c366b88a0745/ssh/session.go#L363
// 2. Go assumes the response returned by forwarded-tcpip is a tcp address: https://github.com/golang/crypto/blob/master/ssh/tcpip.go#L211
// A PR has been made to fix this: https://github.com/fasmide/remotemoe/pull/18 and this module will be ready when that's merged
// should the PR be unmergable for an extended period of time, we can intercept the requests in the net.Conn through a reverse ssh proxy.
func (n *TunnelSuite) TestMoe() {
	n.T().Skip("moe is currently broken, waiting on https://github.com/fasmide/remotemoe/pull/18")
	remoteURL, err := tunnel.StartTunnel(n.GetTestContext(), n.testServer, tunnel.WithProvider(tunnel.Moe))
	time.Sleep(time.Hour)
	n.Require().NoError(err)

	n.checkTunnel(n.GetTestContext(), remoteURL)
}

func (n *TunnelSuite) TestNgrok() {
	remoteURL, err := tunnel.StartTunnel(n.GetTestContext(), n.testServer, tunnel.WithNgrokOptions(), tunnel.WithProvider(tunnel.Ngrok))
	n.Require().NoError(err)

	n.checkTunnel(n.GetTestContext(), remoteURL)
}

func (n *TunnelSuite) checkTunnel(ctx context.Context, remoteURL string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", remoteURL, ginhelper.HealthCheck), nil)
	n.Require().NoError(err)

	resp, err := http.DefaultClient.Do(req)
	n.Require().NoError(err)
	n.Require().Equal(http.StatusOK, resp.StatusCode)

	if resp.Body != nil {
		_ = resp.Body.Close()
	}
}

func (n *TunnelSuite) TestVerifiableProxy() {
	var lc net.ListenConfig
	port := freeport.GetPort()

	hostname := fmt.Sprintf("localhost:%d", port)
	listener, err := lc.Listen(n.GetTestContext(), "tcp", hostname)
	n.Require().NoError(err)

	checkPath := fmt.Sprintf("/%s", gofakeit.Word())

	pathChan := make(chan bool, 1)
	errChan := internal.VerifiableProxy(n.GetTestContext(), n.testServer, checkPath, listener, func() {
		pathChan <- true
	})

	req, err := http.NewRequestWithContext(n.GetTestContext(), http.MethodGet, fmt.Sprintf("http://%s%s", hostname, checkPath), nil)
	n.Require().NoError(err)

	resp, err := http.DefaultClient.Do(req)
	n.Require().NoError(err)

	if resp.Body != nil {
		_ = resp.Body.Close()
	}

	select {
	case <-n.GetTestContext().Done():
		n.Require().NoError(n.GetTestContext().Err())
	case err := <-errChan:
		n.Require().NoError(err)
	case <-pathChan:
		return
	}
}
