package api

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type AttestationAPI interface {
	GetAttestation(context.Context, common.Hash) ([]byte, error)
}
