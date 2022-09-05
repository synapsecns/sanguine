package etherscan

import (
	"context"
	"github.com/pkg/errors"
)

// NewEtherscanAbiGenClientFromChain creates a new ether scan client for a given chain.
// it applies a rate limiter at a file level that is only released when the process is complete.
func NewEtherscanAbiGenClientFromChain(ctx context.Context, chainID uint32, url string) (*Client, error) {
	return newEtherscanABIClient(ctx, chainID, api, false)
}

// ErrNoClientForChain is returned if no analytics client exists for chain.
var ErrNoClientForChain = errors.New("no analytics client for chain")
