package consumer

import (
	"github.com/synapsecns/sanguine/services/explorer/db"
)

type Consumer struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// bridgeParser is the parser to use to parse bridge events
	bridgeParser BridgeParser
}

func NewConsumer(consumerDB db.ConsumerDB, bridgeParser BridgeParser) *Consumer {
	return &Consumer{consumerDB, bridgeParser}
}

//func (c Consumer) getLogsRange(ctx context.Context, chainID uint32, fromBlock, toBlock uint64) ([]ethTypes.Log, error) {
//
//}
