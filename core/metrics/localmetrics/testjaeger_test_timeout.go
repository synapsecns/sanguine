package localmetrics

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStartServer_Timeouts(t *testing.T) {
	tests := []struct {
		name      string
		timeout   time.Duration
		setupPort bool
		wantNil   bool
	}{
		{
			name:    "context timeout during startup",
			timeout: 1 * time.Millisecond,
			wantNil: true,
		},
		{
			name:      "port conflict with short timeout",
			timeout:   1 * time.Second,
			setupPort: true,
			wantNil:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupPort {
				listener, err := net.Listen("tcp", ":16686")
				require.NoError(t, err)
				defer listener.Close()
			}

			ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
			defer cancel()


			server := startServer(ctx, t)
			assert.Equal(t, tt.wantNil, server == nil)

			if server != nil {
				cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cleanupCancel()

				done := make(chan struct{})
				go func() {
					server.purgeResources()
					close(done)
				}()

				select {
				case <-done:
					// Cleanup completed successfully
				case <-cleanupCtx.Done():
					t.Error("cleanup timed out")
				}
			}
		})
	}
}
