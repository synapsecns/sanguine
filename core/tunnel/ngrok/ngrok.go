package ngrok

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/core/tunnel/internal"
	"github.com/synapsecns/sanguine/core/tunnel/types"
	"golang.ngrok.com/ngrok"
	grokconfig "golang.ngrok.com/ngrok/config"
	"net/http"
	"time"
)

type ngrokProvider struct {
	options []ngrok.ConnectOption
	// checkPath is a path we override the backend with to check if the tunnel is up.
	checkPath string
	// checkChan is a channel that is closed when the checkPath is hit.
	checkChan chan bool
}

// New returns a new ngrok provider.
func New(opts ...ngrok.ConnectOption) types.Provider {
	return &ngrokProvider{
		options:   opts,
		checkPath: fmt.Sprintf("/%s", gofakeit.UUID()),
		checkChan: make(chan bool, 1),
	}
}

func (n *ngrokProvider) Start(ctx context.Context, backendURL string) (_ string, err error) {
	// add the default options
	ngrokOptions := append([]ngrok.ConnectOption{
		ngrok.WithAuthtokenFromEnv(),
	}, n.options...)

	listener, err := ngrok.Listen(ctx, grokconfig.HTTPEndpoint(), ngrokOptions...)
	if err != nil {
		return "", fmt.Errorf("could not start ngrok: %w", err)
	}

	errChan := internal.VerifiableProxy(ctx, backendURL, n.checkPath, listener, func() {
		n.checkChan <- true
	})

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
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", listener.URL(), n.checkPath), nil)
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
		case <-n.checkChan:
			return backendURL, nil
		}
	}
}
