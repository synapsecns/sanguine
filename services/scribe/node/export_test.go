package node

import "context"

// ProcessRange exports filtering logs for testing.
func (s Scribe) ProcessRange(ctx context.Context, chainID uint32, requiredConfirmations uint32) error {
	return s.confirmBlocks(ctx, chainID, requiredConfirmations)
}
