package attestation

import (
	"context"
)

// CCTPAPI is the interface for the Attestation API.
type CCTPAPI interface {
	GetAttestation(context.Context, string) ([]byte, error)
}
