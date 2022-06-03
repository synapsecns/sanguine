package indexer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/types"
	synEvm "github.com/synapsecns/synapse-node/pkg/evm"
	"math/big"
)

// HomeIndexer indexes the home contract.
type HomeIndexer struct {
	homeAddress common.Address
	client      synEvm.Chain
}

// NewHomeIndexer creates a new indexer for the home contract.
func NewHomeIndexer(client synEvm.Chain, homeAddress common.Address) HomeIndexer {
	return HomeIndexer{
		homeAddress: homeAddress,
		client:      client,
	}
}

// FetchSortedMessages fetches sorted messages on the home contract.
func (h HomeIndexer) FetchSortedMessages(ctx context.Context, from uint32, to uint32) (messages []types.CommittedMessage, err error) {
	rangeFilter := evm.NewRangeFilter(h.homeAddress, h.client, big.NewInt(int64(from)), big.NewInt(int64(to)), 100, false)

	// blocks until done `
	err = rangeFilter.Start(ctx)
	if err != nil {
		return []types.CommittedMessage{}, fmt.Errorf("could not filter: %w", err)
	}

	filteredLogs, err := rangeFilter.Drain(ctx)
	if err != nil {
		return []types.CommittedMessage{}, fmt.Errorf("could not drain queue: %w", err)
	}

	parser, err := home.NewParser(h.homeAddress)
	if err != nil {
		return []types.CommittedMessage{}, fmt.Errorf("could not get parser: %w", err)
	}

	for _, log := range filteredLogs {
		logType, ok := parser.EventType(log)
		if !ok {
			continue
		}

		if logType == home.DispatchEvent {
			dispatchEvents, ok := parser.ParseDispatch(log)
			// TODO: this should never happen. Maybe we should return an error here?
			if !ok {
				continue
			}

			messages = append(messages, dispatchEvents)
		}
	}

	return messages, nil
}
