package api

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type AttestationApi interface {
	GetAttestation(context.Context, common.Hash) ([]byte, error)
}
