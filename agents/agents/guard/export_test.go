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

// Update wraps update for the AttestationGuardSigner.
func (a AttestationGuardSigner) Update(ctx context.Context) error {
	return a.update(ctx)
}

// Update wraps update for the AttestationGuardColletorSubmitter.
func (a AttestationGuardCollectorSubmitter) Update(ctx context.Context) error {
	return a.update(ctx)
}

// Update wraps update for the AttestationGuardColletorVerifier.
func (a AttestationGuardCollectorVerifier) Update(ctx context.Context) error {
	return a.update(ctx)
}

// Update wraps update for the AttestationGuardDestinationSubmitter.
func (a AttestationGuardDestinationSubmitter) Update(ctx context.Context) error {
	return a.update(ctx)
}
