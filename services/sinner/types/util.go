// Package types holds supplemental types for sinner.
package types

import (
	"context"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// EventParser is the interface for parsing and storing events.
type EventParser interface {
	// ParseAndStore parses and stores the event.
	ParseAndStore(ctx context.Context, log ethTypes.Log) error
	// UpdateTxMap updates the tx map for providing more information to the parser.
	UpdateTxMap(txMap map[string]TxSupplementalInfo)
}

// TxSupplementalInfo is the supplemental info for a tx.
type TxSupplementalInfo struct {
	// TxHash string
	TxHash string
	// Sender is the address of the sender
	Sender string
	// Timestamp is the timestamp of the tx
	Timestamp int
}

// MessageType is the type of message.
type MessageType string

const (
	// Origin is the origin message type.
	Origin MessageType = "origin"
	// Destination is the destination message type from the execution hub.
	Destination MessageType = "destination"
)
