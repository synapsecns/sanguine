package api

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type MockCircleApi struct{}

func (m MockCircleApi) GetAttestation(ctx context.Context, txHash common.Hash) (attestation []byte, err error) {
	return
}
