package reldb

import (
	"context"
	"errors"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
)

type Writer interface {
	// PutLatestBlock upsers the latest block on a given chain id to be new height.
	PutLatestBlock(ctx context.Context, chainID, height uint64) error
}

type Reader interface {
	// LatestBlockForChain gets the latest block for a given chain id.
	LatestBlockForChain(ctx context.Context, chainID uint64) (uint64, error)
}

type Service interface {
	Reader
	SubmitterDB() submitterDB.Service
	Writer
}

// ErrNoLatestBlockForChainID is returned when no block exists for the chain.
var ErrNoLatestBlockForChainID = errors.New("no latest block for chainId")
