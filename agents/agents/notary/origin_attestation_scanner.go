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
	// domain allows access to the origin contract
	domain domains.DomainClient
	// domain of target destination
	destinationID uint32
	// db is the synapse db
	db db.SynapseDB
	// signer is the signer
	signer signer.Signer
	// interval waits for an interval
	interval time.Duration
}

// NewOriginAttestationScanner creates a new origin attestation scanner.
func NewOriginAttestationScanner(domain domains.DomainClient, destinationID uint32, db db.SynapseDB, signer signer.Signer, interval time.Duration) OriginAttestationScanner {
	return OriginAttestationScanner{
		domain:        domain,
		destinationID: destinationID,
		db:            db,
		signer:        signer,
		interval:      interval,
	}
}

// Start starts the OriginAttestationScanner.
func (a OriginAttestationScanner) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(a.interval): // TODO: a.interval
			err := a.update(ctx)
			if err != nil {
				return err
			}
		}
	}
}

// FindLatestNonce fetches the latest nonce for a given chain.
func (a OriginAttestationScanner) FindLatestNonce(ctx context.Context) (nonce uint32, err error) {
	latestNonce, err := a.db.RetrieveLatestCachedNonce(ctx, a.domain.Config().DomainID, a.destinationID)
	if err != nil {
		if errors.Is(err, db.ErrNoNonceForDomain) {
			return 0, nil
		}
		return 0, fmt.Errorf("could not find latest root: %w", err)
	}
	return latestNonce, nil
}

// update runs the job of the scanner
// nolint: cyclop
func (a OriginAttestationScanner) update(ctx context.Context) error {
	latestNonce, err := a.FindLatestNonce(ctx)
	if err != nil {
		return fmt.Errorf("could not find latest root: %w", err)
	}

	// TODO (joe): Currently we are scanning all nonces in order. Later, we really want to get the latest
	// attestation after the latestNonce if any exists.
	nextNonce := latestNonce + 1
	attestation, dispatchBlockNumber, err := a.domain.Origin().GetHistoricalAttestation(ctx, a.destinationID, nextNonce)
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
