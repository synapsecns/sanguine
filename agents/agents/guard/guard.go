package guard

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"golang.org/x/sync/errgroup"
)

// Guard in the current version scans the attestation collector for notary signed attestations,
// signs them, and posts to destination chains.
// TODO: Note right now, I have threads for each origin-destination pair and do no batching at all
type Guard struct {
	scanners        map[string]map[string]AttestationCollectorAttestationScanner
	bondedSigner    signer.Signer
	unbondedSigner  signer.Signer
	refreshInterval time.Duration
}

// NewGuard creates a new guard.
func NewGuard(ctx context.Context, cfg config.GuardConfig) (_ Guard, err error) {
	if cfg.RefreshIntervalInSeconds == int64(0) {
		return Guard{}, fmt.Errorf("cfg.refreshInterval cannot be 0")
	}
	guard := Guard{
		scanners:        make(map[string]map[string]AttestationCollectorAttestationScanner),
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalInSeconds),
	}

	guard.bondedSigner, err = config.SignerFromConfig(cfg.BondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with bondedSigner, could not create guard: %w", err)
	}

	guard.unbondedSigner, err = config.SignerFromConfig(cfg.UnbondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with unbondedSigner, could not create guard: %w", err)
	}

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return Guard{}, fmt.Errorf("could not get legacyDB type: %w", err)
	}

	dbHandle, err := sql.NewStoreFromConfig(ctx, dbType, cfg.Database.ConnString)
	if err != nil {
		return Guard{}, fmt.Errorf("could not connect to legacyDB: %w", err)
	}

	attestationDomainClient, err := evm.NewEVM(ctx, "attestation_collector", cfg.AttestationDomain)
	if err != nil {
		return Guard{}, fmt.Errorf("failing to create evm for attestation collector, could not create guard for B: %w", err)
	}
	for originName, originDomain := range cfg.OriginDomains {
		guard.scanners[originName] = make(map[string]AttestationCollectorAttestationScanner)
		for destinationName, destinationDomain := range cfg.DestinationDomains {
			if originDomain.DomainID == destinationDomain.DomainID {
				continue
			}

			// TODO (joe): other guard workers will submit to destination but for now
			// we are commenting this out since we aren't using the destinationDomainClient yet
			/*destinationDomainClient, err := evm.NewEVM(ctx, destinationName, destinationDomain)
			if err != nil {
				return Guard{}, fmt.Errorf("failing to create evm for destination, could not create guard for: %w", err)
			}*/

			guard.scanners[originName][destinationName] = NewAttestationCollectorAttestationScanner(
				attestationDomainClient,
				originDomain.DomainID,
				destinationDomain.DomainID,
				dbHandle,
				guard.unbondedSigner,
				guard.refreshInterval)
		}
	}

	return guard, nil
}

// Start starts the guard.
func (u Guard) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	for originName, originScanners := range u.scanners {
		for destinationName := range originScanners {
			originName := originName           // capture func literal
			destinationName := destinationName // capture func literal
			g.Go(func() error {
				//nolint: wrapcheck
				return u.scanners[originName][destinationName].Start(ctx)
			})
		}
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start the guard: %w", err)
	}

	return nil
}
