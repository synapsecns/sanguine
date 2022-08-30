package live

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
)

// Scribe is a live scribe that logs all event data.
type Scribe struct {
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// clients is a mapping of chain IDs -> clients
	clients map[uint32]client.EVMClient
	// scribeBackfiller is the backfiller for the scribe
	scribeBackfiller *backfill.ScribeBackfiller
	// config is the config for the scribe
	config config.Config
}

// NewScribe creates a new scribe.
func NewScribe(eventDB db.EventDB, clients []client.EVMClient, config config.Config) (*Scribe, error) {
	// set up the clients mapping
	clientsMap := make(map[uint32]client.EVMClient)
	for _, client := range clients {
		chainID, err := client.ChainID(context.Background())
		if err != nil {
			return nil, fmt.Errorf("could not get chain ID: %w", err)
		}
		clientsMap[uint32(chainID.Uint64())] = client
	}
	// initialize the scribe backfiller
	scribeBackfiller, err := backfill.NewScribeBackfiller(eventDB, clients, config)
	if err != nil {
		return nil, fmt.Errorf("could not create scribe backfiller: %w", err)
	}

	return &Scribe{
		eventDB:          eventDB,
		clients:          clientsMap,
		scribeBackfiller: scribeBackfiller,
		config:           config,
	}, nil
}
