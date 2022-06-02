package indexer

import "context"

type Contract interface {
	// FetchSortedUpdates etches stored updates
	FetchSortedUpdates(ctx context.Context, from uint32, to uint32) error
}

type HomeContract interface {
	Contract
	// FetchSortedMessages fetches sorted messageas
	FetchSortedMessages(ctx context.Context, from uint32, to uint32)
}
