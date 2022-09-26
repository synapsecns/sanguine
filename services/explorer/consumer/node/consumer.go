package node

import "github.com/synapsecns/sanguine/services/explorer/db"

// Consumer is a live consumer that consumes event data from the Scribe's GQL data
// and stores it in the consumerDB.
type Consumer struct {
	// consumerDB is the database to store event data in
	consumerDB db.ConsumerDB
	// chainIDs is the list of chain IDs to consume
	chainIDs []uint32
	// fetchBlockIncrements is a mapping from chainID -> fetch block increments
	fetchBlockIncrements map[uint32]uint64
}

// NewConsumer creates a new consumer.
func NewConsumer(consumerDB db.ConsumerDB, chainIDs []uint32, fetchBlockIncrements map[uint32]uint64) (*Consumer, error) {
	return &Consumer{
		consumerDB:           consumerDB,
		chainIDs:             chainIDs,
		fetchBlockIncrements: fetchBlockIncrements,
	}, nil
}

// Start starts the consumer.
