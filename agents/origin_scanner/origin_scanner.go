package origin_scanner

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"golang.org/x/sync/errgroup"
)

// OriginScanner scans all origins for updated messages.
type OriginScanner struct {
	signer   signer.Signer
	config   config.Config
	interval time.Duration
}

// NewOriginScanner creates a new origin scanner.
func NewOriginScanner(ctx context.Context, cfg config.Config, interval time.Duration) (_ OriginScanner, err error) {
	originScanner := OriginScanner{
		config:   cfg,
		interval: interval,
	}

	originScanner.signer, err = config.SignerFromConfig(cfg.Signer)
	if err != nil {
		return OriginScanner{}, fmt.Errorf("could not create origin scanner: %w", err)
	}

	return originScanner, nil
}

// Start starts the notary.{.
func (o OriginScanner) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// 1) Call Attestation Collector to get domains from the GlobalNotaryRegistry (if we add a new chain, just restart)
	// 2) Check what kind of address Signer is (Notary vs Guard) by querying the GlobalNotaryRegistry which should also have some sort of GuardRegistry.
	// 3) If its a Notary, we only care about our destination

	for _, originDomain := range o.config.Domains {
		originDomainID := originDomain.DomainID
		for _, destinationDomain := range o.config.Domains {
			destinationDomainID := destinationDomain.DomainID
			if originDomainID == destinationDomainID {
				continue
			}
			originDestinationScanner, err := NewOriginDestinationScanner(ctx, o.signer, originDomainID, destinationDomainID, o.interval)
			if err != nil {
				return fmt.Errorf("could not create scanner for origin %d and destination %d: %w", originDomainID, destinationDomainID, err)
			}
			g.Go(func() error {
				//nolint: wrapcheck
				return originDestinationScanner.Start(ctx)
			})
		}
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start the origin scanner: %w", err)
	}

	return nil
}
