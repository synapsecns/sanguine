package notary

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// Notary in the current version scans the origins for new messages, signs them, and posts to attestation collector.
// TODO: Note right now, I have threads for each origin-destination pair and do no batching at all
// in terms of calls to the origin.
// Right now, for this MVP, this is the simplest path and we can make improvements later.
type Notary struct {
	bondedSigner    signer.Signer
	unbondedSigner  signer.Signer
	refreshInterval time.Duration
}

// NewNotary creates a new notary.
//
//nolint:cyclop
func NewNotary(ctx context.Context, cfg config.NotaryConfig) (_ Notary, err error) {
	if cfg.RefreshIntervalInSeconds == int64(0) {
		return Notary{}, fmt.Errorf("cfg.refreshInterval cannot be 0")
	}
	notary := Notary{
		refreshInterval: time.Second * time.Duration(cfg.RefreshIntervalInSeconds),
	}

	notary.bondedSigner, err = config.SignerFromConfig(ctx, cfg.BondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with bondedSigner, could not create notary: %w", err)
	}

	notary.unbondedSigner, err = config.SignerFromConfig(ctx, cfg.UnbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error with unbondedSigner, could not create notary: %w", err)
	}

	destinationClient, err := evm.NewEVM(ctx, "destination_client", cfg.DestinationDomain)
	if err != nil {
		return Notary{}, fmt.Errorf("error with destinationClient, could not create notary for: %w", err)
	}
	summitClient, err := evm.NewEVM(ctx, "summit_client", cfg.SummitDomain)
	if err != nil {
		return Notary{}, fmt.Errorf("error with summitClient, could not create notary for: %w", err)
	}

	err = summitClient.Summit().PrimeNonce(ctx, notary.unbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error trying to PrimeNonce for summitClient, could not create notary for: %w", err)
	}

	err = destinationClient.Destination().PrimeNonce(ctx, notary.unbondedSigner)
	if err != nil {
		return Notary{}, fmt.Errorf("error trying to PrimeNonce for destinationClient, could not create notary for: %w", err)
	}

	return notary, nil
}

// Start starts the notary.
func (u Notary) Start(ctx context.Context) error {
	logger.Info("Notary exiting without error")
	return nil
}
