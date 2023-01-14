package notary

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// OriginAttestationScanner fetches merkle roots for particular origin-destination pair.
// TODO: this needs to become an interface.
type OriginAttestationScanner struct {
	// originDomain allows access to the origin contract on the origin chain
	originDomain domains.DomainClient
	// attestationDomain allows access to the attestation contract on the SYN chain
	attestationDomain domains.DomainClient
	// destinationDomain allows access to the destination contract on the destination chain
	destinationDomain domains.DomainClient
	// db is the synapse db
	db db.SynapseDB
	// bondedSigner is the attestation signer that must be a bonded agent
	bondedSigner signer.Signer
	// unbondedSigner is the signer for submitting transactions
	unbondedSigner signer.Signer
	// interval waits for an interval
	interval time.Duration
}

// NewOriginAttestationScanner creates a new origin attestation scanner.
func NewOriginAttestationScanner(
	originDomain domains.DomainClient,
	attestationDomain domains.DomainClient,
	destinationDomain domains.DomainClient,
	db db.SynapseDB,
	bondedSigner signer.Signer,
	unbondedSigner signer.Signer,
	interval time.Duration) OriginAttestationScanner {
	return OriginAttestationScanner{
		originDomain:      originDomain,
		attestationDomain: attestationDomain,
		destinationDomain: destinationDomain,
		db:                db,
		bondedSigner:      bondedSigner,
		unbondedSigner:    unbondedSigner,
		interval:          interval,
	}
}

// Start starts the OriginAttestationScanner.
func (a OriginAttestationScanner) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Info("Notary OriginAttestationScanner exiting without error")
			return nil
		case <-time.After(a.interval):
			err := a.update(ctx)
			if err != nil {
				logger.Errorf("Notary OriginAttestationScanner exiting with error: %v", err)
				return err
			}
		}
	}
}

// FindLatestNonce fetches the latest nonce for a given chain.
func (a OriginAttestationScanner) FindLatestNonce(ctx context.Context) (nonce uint32, err error) {
	latestNonce, err := a.db.RetrieveLatestCachedNonce(ctx, a.originDomain.Config().DomainID, a.destinationDomain.Config().DomainID)
	if err != nil {
		if errors.Is(err, db.ErrNoNonceForDomain) {
			return 0, nil
		}
		return 0, fmt.Errorf("could not find latest root: %w", err)
	}
	return latestNonce, nil
}

// update runs the job of the scanner
//
//nolint:cyclop
func (a OriginAttestationScanner) update(ctx context.Context) error {
	latestNonce, err := a.FindLatestNonce(ctx)
	if err != nil {
		return fmt.Errorf("could not find latest root: %w", err)
	}

	// TODO (joe): Currently we are scanning all nonces in order. Later, we really want to get the latest
	// attestation after the latestNonce if any exists.
	nextNonce := latestNonce + 1
	attestation, dispatchBlockNumber, err := a.originDomain.Origin().GetHistoricalAttestation(ctx, a.destinationDomain.Config().DomainID, nextNonce)
	if errors.Is(err, domains.ErrNoUpdate) {
		// no update produced this time
		return nil
	}

	if err != nil {
		return fmt.Errorf("could not get historical attestation: %w", err)
	}

	if !a.isConfirmed(dispatchBlockNumber) {
		// not yet confirmed so skip
		return nil
	}

	err = a.db.StoreNewInProgressAttestation(ctx, attestation, dispatchBlockNumber)
	if err != nil {
		return fmt.Errorf("could not store in-progress attestations: %w", err)
	}

	return nil
}

func (a OriginAttestationScanner) isConfirmed(txBlockNum uint64) bool {
	// TODO (joe): figure this out
	return true
}
