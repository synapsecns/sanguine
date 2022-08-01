package updater

import (
	"context"
	"errors"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/synapsecns/sanguine/core/db"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"time"
)

// AttestationSubmitter submits updates continuously.
type AttestationSubmitter struct {
	// domain allows access to the home contract
	domain domains.DomainClient
	// db contains the transaction queue legacyDB
	db db.SynapseDB
	// signer is the signer
	signer signer.Signer
	// interval is count in seconds
	interval time.Duration
}

// NewAttestationSubmitter creates an update producer.
func NewAttestationSubmitter(domain domains.DomainClient, db db.SynapseDB, signer signer.Signer, interval time.Duration) AttestationSubmitter {
	return AttestationSubmitter{
		domain:   domain,
		db:       db,
		signer:   signer,
		interval: interval,
	}
}

// Start runs the update submitter
func (u AttestationSubmitter) Start(ctx context.Context) error {
	committedRoot, err := u.domain.Home().CommittedRoot(ctx)
	if err != nil {
		return fmt.Errorf("could not get committed root: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(u.interval):
			signed, err := u.messageDB.RetrieveProducedUpdate(committedRoot)
			if errors.Is(err, pebble.ErrNotFound) {
				logger.Infof("No produced update to submit for committed_root: %s", committedRoot)
				continue
			} else if err != nil {
				return fmt.Errorf("could not retrieve produced update: %w", err)
			}

			err = u.domain.Home().Update(ctx, u.signer, signed)
			if err != nil {
				return fmt.Errorf("could not produce update: %w", err)
			}

			committedRoot = signed.Update().NewRoot()
		}
	}
}
