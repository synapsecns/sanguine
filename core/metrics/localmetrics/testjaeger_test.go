package localmetrics

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetupTestJaeger(t *testing.T) {
	tests := []struct {
		name    string
		envVars map[string]string
		opts    []Option
		setup   func(*testing.T) *dockertest.Pool
		wantNil bool
	}{
		{
			name: "successful setup without CI",
			envVars: map[string]string{
				"CI": "",
			},
			setup: func(t *testing.T) *dockertest.Pool {
				pool, err := dockertest.NewPool("")
				require.NoError(t, err)
				return pool
			},
			wantNil: false,
		},
		{
			name: "port binding error",
			envVars: map[string]string{
				"CI": "",
			},
			setup: func(t *testing.T) *dockertest.Pool {
				// Create a listener to occupy the port
				listener, err := net.Listen("tcp", ":16686")
				require.NoError(t, err)
				t.Cleanup(func() { listener.Close() })

				pool, err := dockertest.NewPool("")
				require.NoError(t, err)
				return pool
			},
			wantNil: true,
		},
		{
			name: "docker rate limit simulation",
			envVars: map[string]string{
				"CI": "",
			},
			setup: func(t *testing.T) *dockertest.Pool {
				pool, err := dockertest.NewPool("")
				require.NoError(t, err)

				// Create many containers quickly to trigger rate limit
				for i := 0; i < 10; i++ {
					_, err := pool.RunWithOptions(&dockertest.RunOptions{
						Repository: fmt.Sprintf("test-repo-%d", i),
						Tag:        "latest",
					})
					if err != nil && err.Error() == "toomanyrequests: Rate exceeded" {
						return pool
					}
				}
				return pool
			},
			wantNil: true,
		},
		{
			name: "pyroscope enabled with CI",
			envVars: map[string]string{
				"CI": "true",
			},
			opts: []Option{
				WithPyroscopeEnabled(true),
			},
			setup: func(t *testing.T) *dockertest.Pool {
				pool, err := dockertest.NewPool("")
				require.NoError(t, err)
				return pool
			},
			wantNil: false,
		},
		{
			name: "pyroscope disabled with CI",
			envVars: map[string]string{
				"CI": "true",
			},
			setup: func(t *testing.T) *dockertest.Pool {
				pool, err := dockertest.NewPool("")
				require.NoError(t, err)
				return pool
			},
			wantNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup environment
			for k, v := range tt.envVars {
				t.Setenv(k, v)
			}

			pool := tt.setup(t)
			require.NotNil(t, pool)

			// Create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Run test
			server := SetupTestJaeger(ctx, t, tt.opts...)

			// Verify results
			if tt.wantNil {
				assert.Nil(t, server, "expected nil server")
			} else {
				require.NotNil(t, server, "expected non-nil server")
				assert.NotNil(t, server.jaeger, "expected non-nil jaeger")

				// Verify Pyroscope configuration
				if tt.name == "pyroscope enabled with CI" {
					assert.True(t, server.jaeger.cfg.enablePyroscope, "pyroscope should be enabled")
				} else if tt.name == "pyroscope disabled with CI" {
					assert.False(t, server.jaeger.cfg.enablePyroscope, "pyroscope should be disabled")
				}
			}

			// Ensure cleanup happens within timeout
			cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cleanupCancel()

			done := make(chan struct{})
			go func() {
				if server != nil && server.jaeger != nil {
					server.jaeger.purgeResources()
				}
				close(done)
			}()

			select {
			case <-done:
				// Cleanup completed successfully
			case <-cleanupCtx.Done():
				t.Error("cleanup timed out")
			}
		})
	}
}
