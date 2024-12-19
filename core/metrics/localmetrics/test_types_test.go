package localmetrics

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	mockClient := new(MockDockerClient)
	mockClient.On("ContainerStop", mock.Anything, "test-container", mock.Anything).
		Return(nil)
	mockClient.On("ContainerRemove", mock.Anything, "test-container", mock.Anything).
		Return(nil)

	server := &TestServer{
		jaeger: &testJaeger{
			runID:       "test-cleanup",
			containerID: "test-container",
			client:      mockClient,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Cleanup(ctx)
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}
