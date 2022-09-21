package consumer

import (
	"github.com/synapsecns/sanguine/services/explorer/db"
)

// Consumer is the consumer for the events.
type Consumer struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// bridgeParser is the parser to use to parse bridge events
	bridgeParser BridgeParser
}
