package consumer

import (
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/bridge"
)

type Consumer struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// bridgeParser is the parser to use to parse bridge events
	bridgeParser bridge.Parser
}

func NewConsumer(consumerDB db.ConsumerDB, bridgeParser bridge.Parser) *Consumer {
	return &Consumer{consumerDB, bridgeParser}
}

//func (c Consumer) getLogsRange(ctx context.Context, chainID uint32, fromBlock, toBlock uint64) ([]ethTypes.Log, error) {
//
//}
