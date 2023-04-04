package ngrok

import (
	"context"
	"fmt"
	"github.com/jpillora/backoff"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// ProxyURL exposes a server to the internet using ngrok.
func ProxyURL(parentCtx context.Context, backendURL string, options ...ngrok.ConnectOption) (string, error) {
	// reset after 30 seconds
	ctx, cancel := context.WithTimeout(parentCtx, time.Second*30)
	defer cancel()

	// add the default options
	ngrokOptions := append([]ngrok.ConnectOption{
		ngrok.WithAuthtokenFromEnv(),
	}, options...)

	listener, err := ngrok.Listen(ctx, config.HTTPEndpoint(), ngrokOptions...)
	if err != nil {
		return "", fmt.Errorf("could not start ngrok: %w", err)
	}

	errChan := startProxy(ctx, backendURL, listener)

	timeout := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    200 * time.Millisecond,
		Max:    time.Second,
	}

	duration := time.Duration(0)
	for {
		select {
		case <-ctx.Done():
			return "", nil
		case err := <-errChan:
			return "", fmt.Errorf("could not serve ngrok: %w", err)
		case <-time.After(duration):
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, listener.URL(), nil)
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

		}
	}
}

// startProxy creates a proxy to the backend URL.
func startProxy(ctx context.Context, backendURL string, listener net.Listener) (_ <-chan error) {
	errChan := make(chan error, 1)

	parsedURL, err := url.Parse(backendURL)
	if err != nil {
		errChan <- fmt.Errorf("could not parse backend URL: %w", err)
		return errChan
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	go func() {
		err = http.Serve(listener, proxy)
		if err != nil {
			select {
			case errChan <- fmt.Errorf("could not serve ngrok: %w", err):
				return
			case <-ctx.Done():
				return

			}
		}
	}()

	return errChan
}
