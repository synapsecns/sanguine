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
// TODO: Note right now, I have threads for each origin-destination pair and do no batching at all.
type Guard struct {
	scanners                 map[string]map[string]AttestationCollectorAttestationScanner
	originDoubleCheckers     map[string]map[string]AttestationDoubleCheckOnOriginVerifier
	guardSigners             map[string]map[string]AttestationGuardSigner
	guardCollectorSubmitters map[string]map[string]AttestationGuardCollectorSubmitter
	guardCollectorVerifiers  map[string]map[string]AttestationGuardCollectorVerifier
	bondedSigner             signer.Signer
	unbondedSigner           signer.Signer
	refreshInterval          time.Duration
}

// NewGuard creates a new guard.
//
//nolint:cyclop
func NewGuard(ctx context.Context, cfg config.GuardConfig) (_ Guard, err error) {
	if cfg.RefreshIntervalInSeconds == int64(0) {
		return Guard{}, fmt.Errorf("cfg.refreshInterval cannot be 0")
	}
	guard := Guard{
		scanners:                 make(map[string]map[string]AttestationCollectorAttestationScanner),
		originDoubleCheckers:     make(map[string]map[string]AttestationDoubleCheckOnOriginVerifier),
		guardSigners:             make(map[string]map[string]AttestationGuardSigner),
		guardCollectorSubmitters: make(map[string]map[string]AttestationGuardCollectorSubmitter),
		guardCollectorVerifiers:  make(map[string]map[string]AttestationGuardCollectorVerifier),
		refreshInterval:          time.Second * time.Duration(cfg.RefreshIntervalInSeconds),
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
		originDomainClient, err := evm.NewEVM(ctx, originName, originDomain)
		if err != nil {
			return Guard{}, fmt.Errorf("failing to create evm for origiin, could not create guard for: %w", err)
		}
		guard.scanners[originName] = make(map[string]AttestationCollectorAttestationScanner)
		guard.originDoubleCheckers[originName] = make(map[string]AttestationDoubleCheckOnOriginVerifier)
		guard.guardSigners[originName] = make(map[string]AttestationGuardSigner)
		guard.guardCollectorSubmitters[originName] = make(map[string]AttestationGuardCollectorSubmitter)
		guard.guardCollectorVerifiers[originName] = make(map[string]AttestationGuardCollectorVerifier)
		for destinationName, destinationDomain := range cfg.DestinationDomains {
			if originDomain.DomainID == destinationDomain.DomainID {
				continue
			}

			// TODO (joe): other guard workers will submit to destination but for now
			// we are commenting this out since we aren't using the destinationDomainClient yet
			destinationDomainClient, err := evm.NewEVM(ctx, destinationName, destinationDomain)
			if err != nil {
				return Guard{}, fmt.Errorf("failing to create evm for destination, could not create guard for: %w", err)
			}

			guard.scanners[originName][destinationName] = NewAttestationCollectorAttestationScanner(
				attestationDomainClient,
				originDomain.DomainID,
				destinationDomain.DomainID,
				dbHandle,
				guard.unbondedSigner,
				guard.refreshInterval)

			guard.originDoubleCheckers[originName][destinationName] = NewAttestationDoubleCheckOnOriginVerifier(
				originDomainClient,
				attestationDomainClient,
				destinationDomainClient,
				dbHandle,
				guard.bondedSigner,
				guard.unbondedSigner,
				guard.refreshInterval)

			guard.guardSigners[originName][destinationName] = NewAttestationGuardSigner(
				originDomainClient,
				attestationDomainClient,
				destinationDomainClient,
				dbHandle,
				guard.bondedSigner,
				guard.unbondedSigner,
				guard.refreshInterval)

			guard.guardCollectorSubmitters[originName][destinationName] = NewAttestationGuardCollectorSubmitter(
				originDomainClient,
				attestationDomainClient,
				destinationDomainClient,
				dbHandle,
				guard.bondedSigner,
				guard.unbondedSigner,
				guard.refreshInterval)

			guard.guardCollectorVerifiers[originName][destinationName] = NewAttestationGuardCollectorVerifier(
				originDomainClient,
				attestationDomainClient,
				destinationDomainClient,
				dbHandle,
				guard.bondedSigner,
				guard.unbondedSigner,
				guard.refreshInterval)
		}
	}

	return guard, nil
}

// Start starts the guard.
//
//nolint:cyclop
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

	for originName, originToAllDestinationsDoubleCheckers := range u.originDoubleCheckers {
		for destinationName := range originToAllDestinationsDoubleCheckers {
			originName := originName           // capture func literal
			destinationName := destinationName // capture func literal
			g.Go(func() error {
				//nolint: wrapcheck
				return u.originDoubleCheckers[originName][destinationName].Start(ctx)
			})
		}
	}

	for originName, allDestinationGuardSigners := range u.guardSigners {
		for destinationName := range allDestinationGuardSigners {
			originName := originName           // capture func literal
			destinationName := destinationName // capture func literal
			g.Go(func() error {
				//nolint: wrapcheck
				return u.guardSigners[originName][destinationName].Start(ctx)
			})
		}
	}

	for originName, allDestinationGuardCollectorSubmitters := range u.guardCollectorSubmitters {
		for destinationName := range allDestinationGuardCollectorSubmitters {
			originName := originName           // capture func literal
			destinationName := destinationName // capture func literal
			g.Go(func() error {
				//nolint: wrapcheck
				return u.guardCollectorSubmitters[originName][destinationName].Start(ctx)
			})
		}
	}

	for originName, allDestinationGuardCollectorVerifiers := range u.guardCollectorVerifiers {
		for destinationName := range allDestinationGuardCollectorVerifiers {
			originName := originName           // capture func literal
			destinationName := destinationName // capture func literal
			g.Go(func() error {
				//nolint: wrapcheck
				return u.guardCollectorVerifiers[originName][destinationName].Start(ctx)
			})
		}
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start the guard: %w", err)
	}

	return nil
}
