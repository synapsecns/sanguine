package listener

import (
	"context"
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

// FinalityMode represents the finality mode for block queries.

const (
	FinalityModeSafe      rpc.BlockNumber = rpc.SafeBlockNumber
	FinalityModeFinalized rpc.BlockNumber = rpc.FinalizedBlockNumber
	FinalityModeLatest    rpc.BlockNumber = rpc.FinalizedBlockNumber
)

// WithFinalityMode sets the finality mode.
func WithFinalityMode(mode rpc.BlockNumber) Option {
	return func(c *chainListener) {
		c.finalityMode = mode
	}
}

// WithBlockWait sets the block wait.
func WithBlockWait(wait uint64) Option {
	return func(c *chainListener) {
		c.blockWait = wait
	}
}
