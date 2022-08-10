package notary

import "context"

// Update wraps update and produces one if the update has not been produced.
func (a AttestationProducer) Update(ctx context.Context) error {
	return a.update(ctx)
}
