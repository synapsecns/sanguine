package api

import (
	"context"
)

// AttestationAPI is the interface for the Attestation API.
type AttestationAPI interface {
	GetAttestation(context.Context, string) ([]byte, error)
}
