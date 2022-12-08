package notary

import "context"

// Update wraps update and produces one if the update has not been produced.
func (a AttestationProducer) Update(ctx context.Context) error {
	return a.update(ctx)
}

// Update wraps update for OriginAttestationScanner.
func (a OriginAttestationScanner) Update(ctx context.Context) error {
	return a.update(ctx)
}

// Update wraps update for OriginAttestationSigner.
func (a OriginAttestationSigner) Update(ctx context.Context) error {
	return a.update(ctx)
}

// Update wraps update for OriginAttestationSubmitter.
func (a OriginAttestationSubmitter) Update(ctx context.Context) error {
	return a.update(ctx)
}

// Update wraps update for OriginAttestationVerifier.
func (a OriginAttestationVerifier) Update(ctx context.Context) error {
	return a.update(ctx)
}
