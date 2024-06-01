package listener

import "context"

// Option is a functional option for chainListener.
type Option func(*chainListener)

// NewBlockHandler is a function that is called when a new block is detected.
type NewBlockHandler func(ctx context.Context, block uint64) error

// WithNewBlockHandler sets the new block handler.
func WithNewBlockHandler(handler NewBlockHandler) Option {
	return func(c *chainListener) {
		c.newBlockHandler = handler
	}
}
