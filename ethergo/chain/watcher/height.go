package watcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"math/big"
)

// EthHeadClient defines a method for getting a subscription to the chain-tip height on geth based rpc clients.
type EthHeadClient interface {
	// HeaderByNumber gets a block by its number, If nil is passed for number, the most recent
	// block is returned.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

// blockHeightWatcher wraps EthHeadClient to return chain agnostic heads.
type blockHeightWatcher struct {
	subscriberClient EthHeadClient
}

func newBlockSubscriber(underlyingClient EthHeadClient) chainwatcher.BlockSubscriberClient {
	return blockHeightWatcher{subscriberClient: underlyingClient}
}

func (l blockHeightWatcher) LatestHeight(ctx context.Context) (uint64, error) {
	latestBlock, err := l.subscriberClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("could not get latest height: %w", err)
	}
	return latestBlock.Number.Uint64(), nil
}

var _ chainwatcher.BlockSubscriberClient = &blockHeightWatcher{}

// NewBlockHeightWatcher creates a new height watcher.
func NewBlockHeightWatcher(ctx context.Context, chainID uint64, reader EthHeadClient) chainwatcher.BlockHeightWatcher {
	return chainwatcher.NewBlockHeightWatcher(ctx, chainID, newBlockSubscriber(reader))
}
