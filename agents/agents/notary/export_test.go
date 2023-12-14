package notary

import "context"

func (n *Notary) EnsureNotaryActive(ctx context.Context) error {
	return n.ensureNotaryActive(ctx)
}
