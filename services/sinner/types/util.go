package types

import (
	"context"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/sinner/contracts/destination"
	"github.com/synapsecns/sanguine/services/sinner/contracts/origin"
)

// Parsers holds all the parsers for a given chain
type Parsers struct {
	// ChainID is the chain these parsers are for.
	ChainID uint32
	// OriginParser parses logs from the origin contract.
	OriginParser *origin.ParserImpl
	// DestinationParser parses logs from the execution hub contract.
	DestinationParser *destination.ParserImpl
}

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
