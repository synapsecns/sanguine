package api

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// MockCircleAPI is a mock CircleAPI for testing.
type MockCircleAPI struct {
	attestFunc func(context.Context, common.Hash) ([]byte, error)
}

// NewMockCircleAPI creates a new MockCircleAPI.
func NewMockCircleAPI() *MockCircleAPI {
	return &MockCircleAPI{
		attestFunc: func(context.Context, common.Hash) ([]byte, error) { return []byte{}, nil },
	}
}

// GetAttestation is a mock GetAttestation.
func (m *MockCircleAPI) GetAttestation(ctx context.Context, txHash common.Hash) (attestation []byte, err error) {
	return m.attestFunc(ctx, txHash)
}

// SetGetAttestation overrides the attestFunc.
func (m *MockCircleAPI) SetGetAttestation(attestFunc func(context.Context, common.Hash) ([]byte, error)) {
	m.attestFunc = attestFunc
}
