package api

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// MockCircleAPI is a mock CircleAPI for testing.
type MockCircleAPI struct {
	attestFunc func(context.Context, common.Hash) ([]byte, error)
}

func NewMockCircleAPI() *MockCircleAPI {
	return &MockCircleAPI{
		attestFunc: func(context.Context, common.Hash) ([]byte, error) { return []byte{}, nil },
	}
}

func (m *MockCircleAPI) GetAttestation(ctx context.Context, txHash common.Hash) (attestation []byte, err error) {
	return m.attestFunc(ctx, txHash)
}

func (m *MockCircleAPI) SetGetAttestation(attestFunc func(context.Context, common.Hash) ([]byte, error)) {
	m.attestFunc = attestFunc
}
