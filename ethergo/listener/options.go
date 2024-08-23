package listener

import (
	"context"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

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

// WithPollInterval sets the poll interval.
func WithPollInterval(interval time.Duration) Option {
	return func(c *chainListener) {
		c.pollIntervalSetting = interval
	}
}

// WithFinalityMode sets the finality mode.
func WithFinalityMode(mode string) Option {
	return func(c *chainListener) {
		switch strings.ToLower(mode) {
		case "latest":
			c.finalityMode = rpc.LatestBlockNumber
		case "safe":
			c.finalityMode = rpc.SafeBlockNumber
		case "finalized":
			c.finalityMode = rpc.FinalizedBlockNumber
		default:
			c.finalityMode = rpc.LatestBlockNumber
		}
	}
}

// WithBlockWait sets the block wait.
func WithBlockWait(wait uint64) Option {
	return func(c *chainListener) {
		c.blockWait = wait
	}
}

// WithName sets the listener name.
func WithName(name string) Option {
	return func(c *chainListener) {
		c.name = name
		c.store.SetListenerName(name)
	}
}
