package updater

import "context"

// Update wraps update and produces one if the update has not been produced.
func (u UpdateProducer) Update(ctx context.Context) error {
	return u.update(ctx)
}
