package guard

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// Guard in the current version scans the attestation collector for notary signed attestations,
// signs them, and posts to destination chains.
// TODO: Note right now, I have threads for each origin-destination pair and do no batching at all.
type Guard struct {
	bondedSigner    signer.Signer
	unbondedSigner  signer.Signer
	refreshInterval time.Duration
}

// NewGuard creates a new guard.
//
//nolint:cyclop
func NewGuard(ctx context.Context, cfg config.GuardConfig) (_ Guard, err error) {
	if cfg.RefreshIntervalInSeconds == int64(0) {
		return Guard{}, fmt.Errorf("cfg.refreshInterval cannot be 0")
	}
	guard := Guard{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalInSeconds),
	}

	guard.bondedSigner, err = config.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with bondedSigner, could not create guard: %w", err)
	}

	guard.unbondedSigner, err = config.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error with unbondedSigner, could not create guard: %w", err)
	}

	summitDomainClient, err := evm.NewEVM(ctx, "summit_collector", cfg.SummitDomain)
	if err != nil {
		return Guard{}, fmt.Errorf("failing to create evm for summit, could not create guard for B: %w", err)
	}
	err = summitDomainClient.Summit().PrimeNonce(ctx, guard.unbondedSigner)
	if err != nil {
		return Guard{}, fmt.Errorf("error trying to PrimeNonce for summitClient, could not create notary for: %w", err)
	}

	/*for originName, originDomain := range cfg.OriginDomains {
		originDomainClient, err := evm.NewEVM(ctx, originName, originDomain)
		if err != nil {
			return Guard{}, fmt.Errorf("failing to create evm for origiin, could not create guard for: %w", err)
		}
	}*/

	return guard, nil
}

// Start starts the guard.
//
//nolint:cyclop
func (u Guard) Start(ctx context.Context) error {
	logger.Info("Guard exiting without error")
	return nil
}
