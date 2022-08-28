package backfill

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
)

// ChainBackfiller is a backfiller that fetches logs for a chain. It aggregates logs
// from a slice of ContractBackfillers.
type ChainBackfiller struct {
	// contracts is the list of contracts to get logs for
	contracts []contracts.DeployedContract
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client client.EVMClient
	// contractBackfillers is the list of contract backfillers
	contractBackfillers []ContractBackfiller
	// chainConfig is the config for the backfiller
	chainConfig config.ChainConfig
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(contracts []contracts.DeployedContract, eventDB db.EventDB, client client.EVMClient, chainConfig config.ChainConfig) (*ChainBackfiller, error) {
	// initialize the list of contract backfillers
	contractBackfillers := make([]ContractBackfiller, len(contracts))
	// initialize each contract backfiller
	for i, contract := range contracts {
		contractBackfiller, err := NewContractBackfiller(contract, eventDB, client)
		if err != nil {
			return nil, fmt.Errorf("could not create contract backfiller: %w", err)
		}
		contractBackfillers[i] = *contractBackfiller
	}

	return &ChainBackfiller{
		contracts:           contracts,
		eventDB:             eventDB,
		client:              client,
		contractBackfillers: contractBackfillers,
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
		startHeight := c.chainConfig.Contracts[contractBackfiller.contract.Address().String()].StartBlock
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
