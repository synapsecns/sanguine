package attestation

import (
	"context"
)

// MockCircleAPI is a mock CircleAPI for testing.
type MockCircleAPI struct {
	attestFunc func(context.Context, string) ([]byte, error)
}

// NewMockCircleAPI creates a new MockCircleAPI.
func NewMockCircleAPI() *MockCircleAPI {
	return &MockCircleAPI{
		attestFunc: func(context.Context, string) ([]byte, error) { return []byte{}, nil },
	}
}

// GetAttestation is a mock GetAttestation.
func (m *MockCircleAPI) GetAttestation(ctx context.Context, txHash string) (attestation []byte, err error) {
	return m.attestFunc(ctx, txHash)
}

// SetGetAttestation overrides the attestFunc.
func (m *MockCircleAPI) SetGetAttestation(attestFunc func(context.Context, string) ([]byte, error)) {
	m.attestFunc = attestFunc
}
