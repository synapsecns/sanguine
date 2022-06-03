package indexer

import (
	"context"
	"github.com/synapsecns/sanguine/core/types"
)

// Contract is the contract interface.
type Contract interface {
	// FetchSortedUpdates etches stored updates
	FetchSortedUpdates(ctx context.Context, from uint32, to uint32) (updates []types.Update, err error)
}

// HomeContract is the interface for the home contract.
type HomeContract interface {
	Contract
	// FetchSortedMessages fetches sorted messageas
	FetchSortedMessages(ctx context.Context, from uint32, to uint32) (updates []types.Update, err error)
}
