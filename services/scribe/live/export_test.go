package live

import "context"

func (s Scribe) ProcessRange(ctx context.Context, chainID uint32) error {
	return s.processRange(ctx, chainID)
}
