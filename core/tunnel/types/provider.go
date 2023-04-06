// Package types provides common types to avoid circular dependencies.
package types

import (
	"context"
)

// Provider is a tunnel provider.
type Provider interface {
	// Start starts the tunnel provider and returns the URL of the tunnel.
	Start(ctx context.Context, backendURL string) (_ string, err error)
}
