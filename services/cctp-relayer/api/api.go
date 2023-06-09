package api

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// AttestationAPI is the interface for the Attestation API.
type AttestationAPI interface {
	GetAttestation(context.Context, common.Hash) ([]byte, error)
}
