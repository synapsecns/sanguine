package types

import (
	"context"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type EventParser interface {
	ParseAndStore(ctx context.Context, log ethTypes.Log) error
	UpdateTxMap(txMap map[string]TxSupplementalInfo)
}

type TxSupplementalInfo struct {
	// TxHash string
	TxHash string
	// Sender is the address of the sender
	Sender string
	// Timestamp is the timestamp of the tx
	Timestamp int
}

type MessageType string

const (
	Origin      MessageType = "origin"
	Destination MessageType = "destination"
)
