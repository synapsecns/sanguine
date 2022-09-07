package db

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
)

// SwapDBWriter writes bridge events to the db interface
type SwapDBWriter interface {
	// StoreTokenSwap stores a token swap event
	StoreTokenSwap(ctx context.Context, rawLog types.Log, swap swap.SwapUtilsTokenSwap, chainID uint32)
}

type EventDB interface {
	SwapDBWriter
}
