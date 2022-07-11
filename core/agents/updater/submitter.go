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

// UpdateSubmitter submits updates continuously.
type UpdateSubmitter struct {
	// domain allows access to the home contract
	domain domains.DomainClient
	// messageDB contains the messageDB object
	messageDB db.MessageDB
	// txQueueDB contains the transaction queue db
	txQueueDB db.TxQueueDB
	// signer is the signer
	signer signer.Signer
	// interval is count in seconds
	interval time.Duration
}

// NewUpdateSubmitter creates an update producer.
func NewUpdateSubmitter(domain domains.DomainClient, messageDB db.MessageDB, txDB db.TxQueueDB, signer signer.Signer, interval time.Duration) UpdateSubmitter {
	return UpdateSubmitter{
		domain:    domain,
		messageDB: messageDB,
		txQueueDB: txDB,
		signer:    signer,
		interval:  interval,
	}
}

// Start runs the update submitter
// todo: next up you need to borrow the tx loop from synapse-node and test well
// myabe in ethergo? Should be agnostic and utilize nonce manager.
func (u UpdateSubmitter) Start(ctx context.Context) error {
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
