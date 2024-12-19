package localmetrics

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSetupTestJaeger(t *testing.T) {
	tests := []struct {
		name      string
		envVars   map[string]string
		opts      []Option
		mockSetup func(*MockDockerClient)
		wantNil   bool
	}{
		{
			name: "successful setup without CI",
			envVars: map[string]string{
				"CI": "",
			},
			mockSetup: func(m *MockDockerClient) {
				m.On("ContainerCreate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(container.CreateResponse{ID: "test-container"}, nil)
				m.On("ContainerStart", mock.Anything, "test-container", mock.Anything).
					Return(nil)
				m.On("ContainerLogs", mock.Anything, "test-container", mock.Anything).
					Return(io.NopCloser(strings.NewReader("container started")), nil)
			},
			wantNil: false,
		},
		{
			name: "docker rate limit error",
			envVars: map[string]string{
				"CI": "",
			},
			mockSetup: func(m *MockDockerClient) {
				m.On("ContainerCreate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(container.CreateResponse{}, errors.New("toomanyrequests: Rate exceeded"))
			},
			wantNil: true,
		},
		{
			name: "container start timeout",
			envVars: map[string]string{
				"CI": "",
			},
			mockSetup: func(m *MockDockerClient) {
				m.On("ContainerCreate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(container.CreateResponse{ID: "test-container"}, nil)
				m.On("ContainerStart", mock.Anything, "test-container", mock.Anything).
					After(2*time.Second).Return(nil)
				m.On("ContainerStop", mock.Anything, "test-container", mock.Anything).
					Return(nil)
				m.On("ContainerRemove", mock.Anything, "test-container", mock.Anything).
					Return(nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock docker client
			mockClient := new(MockDockerClient)
			if tt.mockSetup != nil {
				tt.mockSetup(mockClient)
			}

			// Setup environment
			for k, v := range tt.envVars {
				t.Setenv(k, v)
			}

			// Create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			// Run test
			server := SetupTestJaeger(ctx, t, tt.opts...)

			// Verify results
			if tt.wantNil {
				assert.Nil(t, server, "expected nil server")
			} else {
				require.NotNil(t, server, "expected non-nil server")
				assert.NotNil(t, server.jaeger, "expected non-nil jaeger")
			}

			// Verify all mock expectations were met
			mockClient.AssertExpectations(t)
		})
	}
}
