package updater

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/indexer"
)

// Updater updates the updater.
type Updater struct {
	indexers map[string]indexer.DomainIndexer
}

// NewUpdater creates a new updater.
func NewUpdater(ctx context.Context, cfg config.Config) (Updater, error) {
	updater := Updater{
		indexers: make(map[string]indexer.DomainIndexer),
	}
	for name, domain := range cfg.Domains {
		domainClient, err := evm.NewEVM(ctx, name, domain)
		if err != nil {
			return Updater{}, fmt.Errorf("could not create updater for: %w", err)
		}

		dbHandle, err := db.NewDB(cfg.DBPath, name)
		if err != nil {
			return Updater{}, fmt.Errorf("can not create db: %w", err)
		}

		updater.indexers[name] = indexer.NewDomainIndexer(dbHandle, domainClient)
	}

	return updater, nil
}

// Start starts the updater.{.
func (u Updater) Start(ctx context.Context) {
	// returnDBPath, name)
	//if err != nil
	//	name := name
	//go func() {
	//	_ = updater.indexers[name].SyncMessages(ctx)
	//}()
}
