package moe

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ipfs/go-log"
	"github.com/jpillora/backoff"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core/tunnel/internal"
	"github.com/synapsecns/sanguine/core/tunnel/types"
	"net"
	"net/http"
	"time"
)

var logger = log.Logger("moe")

type moeProvider struct {
	// serverURL is the URL of the moe server.
	serverURL string
	// checkPath is a path we override the backend with to check if the tunnel is up.
	checkPath string
	// checkChan is a channel that is closed when the checkPath is hit.
	checkChan chan bool
	cfg       *config
}

// New returns a new moe provider.
func New(opts ...Option) types.Provider {
	cfg := makeConfig(opts)

	return &moeProvider{
		cfg:       cfg,
		checkPath: fmt.Sprintf("/%s", gofakeit.UUID()),
		checkChan: make(chan bool, 1),
	}
}

type hostInfo struct {
	hostname string
	port     int
}

// makeListener returns the hostname and port of a listener.
func (m *moeProvider) makeListener() hostInfo {
	port := freeport.GetPort()
	return hostInfo{
		port:     port,
		hostname: fmt.Sprintf("localhost:%d", port),
	}
}

// nolint: cyclop
func (m *moeProvider) Start(ctx context.Context, backendURL string) (_ string, err error) {
	var lc net.ListenConfig

	// make the verifiable proxy and listener
	verifiableProxy := m.makeListener()
	vpListener, err := lc.Listen(ctx, "tcp", verifiableProxy.hostname)
	if err != nil {
		return "", fmt.Errorf("could not start moe: %w", err)
	}

	errChan := internal.VerifiableProxy(ctx, backendURL, m.checkPath, vpListener, func() {
		m.checkChan <- true
	})

	// make the backend proxy and listener
	host, err := createTunnel(ctx, m.serverURL, verifiableProxy.port, m.cfg.remotePort)
	if err != nil {
		return "", fmt.Errorf("could not get tunnel")
	}

	// TODO: this needs to be deduped w/ moe
	timeout := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    200 * time.Millisecond,
		Max:    time.Second,
	}

	// give the server a max of 30 seconds to start
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	duration := time.Duration(0)
	for {
		select {
		case <-ctx.Done():
			return "", nil
		case err := <-errChan:
			return "", fmt.Errorf("could not serve ngrok: %w", err)
		case <-time.After(duration):
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", host, m.checkPath), nil)
			if err != nil {
				return "", fmt.Errorf("could not create request: %w", err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				duration = timeout.Duration()
				continue
			}

			if resp != nil {
				_ = resp.Body.Close()
			}
		case <-m.checkChan:
			return backendURL, nil
		}
	}
}
