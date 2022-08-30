package livefill

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
)

// ChainLivefiller is a livefiller that listens for logs from a chain. It aggregates logs
// from a slice of ContractLivefillers.
type ChainLivefiller struct {
	// chainID is the chainID of the chain to listen for logs from
	chainID uint32
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client client.EVMClient
	// contractLivefillers is the list of contract livefillers
	contractLivefillers []*ContractLivefiller
	// chainConfig is the config for the livefiller
	chainConfig config.ChainConfig
}

// NewChainLivefiller creates a new livefiller for a chain.
func NewChainLivefiller(chainID uint32, eventDB db.EventDB, client client.EVMClient, chainConfig config.ChainConfig) (*ChainLivefiller, error) {
	// initialize the list of contract livefillers
	contractLivefillers := []*ContractLivefiller{}
	// initialize each contract livefiller
	for _, contract := range chainConfig.Contracts {
		contractLivefiller, err := NewContractLivefiller(chainID, contract.Address, eventDB, client)
		if err != nil {
			return nil, fmt.Errorf("could not create contract livefiller: %w", err)
		}
		contractLivefillers = append(contractLivefillers, contractLivefiller)
	}

	return &ChainLivefiller{
		chainID:             chainID,
		eventDB:             eventDB,
		client:              client,
		contractLivefillers: contractLivefillers,
		chainConfig:         chainConfig,
	}, nil
}

// Livefill iterates over each contract livefiller and calls Livefill concurrently on each one.
func (c ChainLivefiller) Livefill(ctx context.Context) error {
	// initialize the errgroup
	g, ctx := errgroup.WithContext(ctx)
	// iterate over each contract livefiller
	for _, contractLivefiller := range c.contractLivefillers {
		// capture the func literal
		contractLivefiller := contractLivefiller
		// call Livefill concurrently
		g.Go(func() error {
			err := contractLivefiller.Livefill(ctx)
			if err != nil {
				return fmt.Errorf("could not livefill contract: %w", err)
			}
			return nil
		})
	}
	// wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not livefill chain: %w", err)
	}

	return nil
}
