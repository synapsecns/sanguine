package backfill

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
)

// ChainBackfiller is a backfiller that fetches logs for a chain. It aggregates logs
// from a slice of ContractBackfillers.
type ChainBackfiller struct {
	// chainID is the chain ID of the chain
	chainID uint32
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client client.EVMClient
	// contractBackfillers is the list of contract backfillers
	contractBackfillers []*ContractBackfiller
	// startHeights is a map from address -> start height
	startHeights map[string]uint64
	// chainConfig is the config for the backfiller
	chainConfig config.ChainConfig
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(chainID uint32, eventDB db.EventDB, client client.EVMClient, chainConfig config.ChainConfig) (*ChainBackfiller, error) {
	// initialize the list of contract backfillers
	contractBackfillers := []*ContractBackfiller{}
	// initialize each contract backfiller and start heights
	startHeights := make(map[string]uint64)
	for _, contract := range chainConfig.Contracts {
		contractBackfiller, err := NewContractBackfiller(chainConfig.ChainID, contract.Address, eventDB, client)
		if err != nil {
			return nil, fmt.Errorf("could not create contract backfiller: %w", err)
		}
		contractBackfillers = append(contractBackfillers, contractBackfiller)
		startHeights[contract.Address] = contract.StartBlock
	}

	return &ChainBackfiller{
		chainID:             chainID,
		eventDB:             eventDB,
		client:              client,
		contractBackfillers: contractBackfillers,
		startHeights:        startHeights,
		chainConfig:         chainConfig,
	}, nil
}

// Backfill iterates over each contract backfiller and calls Backfill concurrently on each one.
func (c ChainBackfiller) Backfill(ctx context.Context, endHeight uint64) error {
	// initialize the errgroup
	g, ctx := errgroup.WithContext(ctx)
	// iterate over each contract backfiller
	for _, contractBackfiller := range c.contractBackfillers {
		// capture func literal
		contractBackfiller := contractBackfiller
		// get the start height for the backfill
		startHeight := c.startHeights[contractBackfiller.address]
		// call Backfill concurrently
		g.Go(func() error {
			err := contractBackfiller.Backfill(ctx, startHeight, endHeight)
			if err != nil {
				return fmt.Errorf("could not backfill contract: %w", err)
			}
			return nil
		})
	}
	// wait for all of the backfillers to finish
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	return nil
}
