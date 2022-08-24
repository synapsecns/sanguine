package etherscan

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	synapseCommon "github.com/synapsecns/synapse-node/pkg/common"
)

// NewEtherscanAbiGenClientFromChain creates a new ether scan client for a given chain.
// it applies a rate limiter at a file level that is only released when the process is complete.
func NewEtherscanAbiGenClientFromChain(ctx context.Context, chainID uint) (*Client, error) {
	chain := synapseCommon.ChainIDs.ChainByID(chainID)

	api, hasAPI := APIForChain(chainID)
	if !hasAPI {
		return nil, fmt.Errorf("%w: chain id %d", ErrNoClientForChain, chainID)
	}
	return newEtherscanABIClient(ctx, chain, api, false)
}

// APIForChain gets the analytics for a given chain that matches the etherscan standard
// returns false if none matches.
func APIForChain(chainID uint) (api synapseCommon.API, hasAPI bool) {
	chain := synapseCommon.ChainIDs.ChainByID(chainID)

	for _, api = range chain.APIs {
		if api.Standard != synapseCommon.EtherscanAPIStandard {
			continue
		}
		return api, true
	}
	return synapseCommon.API{}, false
}

// ErrNoClientForChain is returned if no analytics client exists for chain.
var ErrNoClientForChain = errors.New("no analytics client for chain")
