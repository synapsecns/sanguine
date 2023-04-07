// Package tunnel provides a simple interface to start a tunnel to a backend URL.
package tunnel

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/tunnel/moe"
	ngrok "github.com/synapsecns/sanguine/core/tunnel/ngrok"
	"github.com/synapsecns/sanguine/core/tunnel/types"
)

// StartTunnel starts a tunnel to the backend URL.
func StartTunnel(ctx context.Context, backendURL string, opts ...Option) (string, error) {
	cfg, err := makeConfig(opts)
	if err != nil {
		return "", err
	}

	var provider types.Provider

	switch cfg.provider {
	case Moe:
		provider = moe.New()
	case Ngrok:
		provider = ngrok.New(cfg.ngrokOptions...)
	}
	tunnelURL, err := provider.Start(ctx, backendURL)
	if err != nil {
		return "", fmt.Errorf("could not start tunnel: %w", err)
	}
	return tunnelURL, nil
}
