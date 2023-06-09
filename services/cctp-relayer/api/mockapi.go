package api

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// MockCircleApi is a mock CircleApi for testing.
type MockCircleApi struct {
	attestFunc func(context.Context, common.Hash) ([]byte, error)
}

func NewMockCircleApi() *MockCircleApi {
	return &MockCircleApi{
		attestFunc: func(context.Context, common.Hash) ([]byte, error) { return []byte{}, nil },
	}
}

func (m *MockCircleApi) GetAttestation(ctx context.Context, txHash common.Hash) (attestation []byte, err error) {
	return m.attestFunc(ctx, txHash)
}

func (m *MockCircleApi) SetGetAttestation(attestFunc func(context.Context, common.Hash) ([]byte, error)) {
	m.attestFunc = attestFunc
}
