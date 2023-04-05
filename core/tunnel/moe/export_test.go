package moe

import (
	"context"
	"io"
)

// Export for testing.
const MoeServer = moeServer

// ConsumeBufferUntilURL exports consumeBufferUntilURL for testing.
func ConsumeBufferUntilURL(ctx context.Context, reader io.Reader) (host string, err error) {
	return consumeBufferUntilURL(ctx, reader)
}
