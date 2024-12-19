package localmetrics

import (
	"context"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTestServer_GetRunID(t *testing.T) {
	tests := []struct {
		name    string
		server  *TestServer
		wantID  string
		wantNil bool
	}{
		{
			name: "valid server with run ID",
			server: &TestServer{
				jaeger: &testJaeger{
					runID: "test-123",
				},
			},
			wantID: "test-123",
		},
		{
			name: "nil jaeger",
			server: &TestServer{
				jaeger: nil,
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantNil {
				assert.Panics(t, func() {
					tt.server.GetRunID()
				}, "expected panic when jaeger is nil")
				return
			}

			gotID := tt.server.GetRunID()
			assert.Equal(t, tt.wantID, gotID)
		})
	}
}

func TestTestServer_Cleanup(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "jaegertracing/all-in-one",
		Tag:        "1.22",
	})
	require.NoError(t, err)

	server := &TestServer{
		jaeger: &testJaeger{
			runID: "test-cleanup",
			pool:  pool,
			jaegerResource: &uiResource{
				Resource: resource,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.jaeger.purgeResources()
	assert.NoError(t, err)
}
