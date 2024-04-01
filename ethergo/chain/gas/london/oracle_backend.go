package london

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/backend"
)

// NewOracleBackendFromHeight creates a fee oracle for deterministically generating gas prices from a given height
// that hijacks the header by number method (when it's passed -1) and returns the block height passed in the constructor
// it also overrides PendingBlockAndReceipts to always return nil since these cannot deterministically be used cross-client.
func NewOracleBackendFromHeight(chain backend.OracleBackendChain, height uint64) HeightOracleBackend {
	return HeightOracleBackend{
		chain:  chain,
		height: height,
	}
}

// HeightOracleBackend is an oracle that deterministically gets gas + fee prices for a given backend.
type HeightOracleBackend struct {
	chain  backend.OracleBackendChain
	height uint64
}

// SubscribeChainHeadEvent is not used, but required for analytics compatibility).
func (h HeightOracleBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return nil
}

// HeaderByNumber wraps oracle baccend.
//
//nolint:wrapcheck
func (h HeightOracleBackend) HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error) {
	if number == rpc.LatestBlockNumber {
		return h.chain.HeaderByNumber(ctx, big.NewInt(int64(h.height)))
	}
	return h.chain.HeaderByNumber(ctx, big.NewInt(int64(number)))
}

// BlockByNumber wraps the oracle backend.
//
//nolint:wrapcheck
func (h HeightOracleBackend) BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	return h.chain.BlockByNumber(ctx, big.NewInt(int64(number)))
}

// GetReceipts gets receipts for a block in a single rpc call.
//
//nolint:wrapcheck
func (h HeightOracleBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	block, err := h.chain.BlockByHash(ctx, hash)
	if err != nil {
		return nil, fmt.Errorf("could not fetch block with hash: %s: %w", hash, err)
	}

	var receipts types.Receipts

	err = receipts.DeriveFields(h.ChainConfig(), hash, block.NumberU64(), block.Time(), block.BaseFee(), block.Transactions())
	if err != nil {
		return nil, fmt.Errorf("could not derive receipts from block: %w", err)
	}

	return receipts, nil
}

// PendingBlockAndReceipts always returns nil since we can't use this cross-client in a deterministic way.
func (h HeightOracleBackend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	return nil, nil
}

// ChainConfig gets the chainconfig for the chain.
func (h HeightOracleBackend) ChainConfig() *params.ChainConfig {
	return h.chain.ChainConfig()
}

var _ gasprice.OracleBackend = &HeightOracleBackend{}
