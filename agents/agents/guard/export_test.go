package guard

import "context"

// Update wraps update for the AttestationCollectorAttestationScanner.
func (a AttestationCollectorAttestationScanner) Update(ctx context.Context) error {
	return a.update(ctx)
}

// Update wraps update for the AttestationCollectorAttestationScanner.
func (a AttestationDoubleCheckOnOriginVerifier) Update(ctx context.Context) error {
	return a.update(ctx)
}
