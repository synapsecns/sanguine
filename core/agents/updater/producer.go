package updater

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core/db"
	pebble2 "github.com/synapsecns/sanguine/core/db/datastore/pebble"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// UpdateProducer updates a producer.
// TODO: this needs to become an interface.
type UpdateProducer struct {
	// domain allows access to the home contract
	domain domains.DomainClient
	// db contains the db object
	db db.MessageDB
	// signer is the signer
	signer signer.Signer
}

// NewUpdateProducer creates an update producer.
func NewUpdateProducer(domain domains.DomainClient, db db.MessageDB, signer signer.Signer) UpdateProducer {
	return UpdateProducer{
		domain: domain,
		db:     db,
		signer: signer,
	}
}

// FindLatestRoot finds the latest root.
func (u UpdateProducer) FindLatestRoot() (common.Hash, error) {
	latestRoot, err := u.db.RetrieveLatestRoot()
	if err != nil && errors.Is(err, pebble.ErrNotFound) {
		return common.Hash{}, nil
	} else if err != nil {
		return common.Hash{}, fmt.Errorf("could not retrieve latest root: %w", err)
	}

	return latestRoot, nil
}

// StoreProducedUpdate stores a pending update in the MessageDB for potential submission.
//
// This does not produce update meta or update the latest update messageDB value.
// It is used by update production and submission.
func (u UpdateProducer) StoreProducedUpdate(update types.SignedUpdate) error {
	existingOpt, err := u.db.RetrieveProducedUpdate(update.Update().PreviousRoot())
	if err != nil && !errors.Is(err, pebble.ErrNotFound) {
		return fmt.Errorf("could not retrieve produced update: %w", err)
	}

	if errors.Is(err, pebble.ErrNotFound) {
		return u.db.StoreProducedUpdate(update.Update().PreviousRoot(), update)
	} else {
		if existingOpt.Update().NewRoot() != update.Update().NewRoot() {
			return fmt.Errorf("updater attempted to store conflicting update. Existing update: %s. New conflicting update: %S.\"", update.Update().NewRoot(), update.Update().NewRoot())
		}
	}
	return nil
}

// Start starts the update producer.
func (u UpdateProducer) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			latestRoot, err := u.FindLatestRoot()
			if err != nil {
				return fmt.Errorf("could not find latest root: %w", err)
			}

			suggestedUpdate, err := u.domain.Home().ProduceUpdate(ctx)
			if err != nil {
				return fmt.Errorf("could not suggest update: %w", err)
			}

			if suggestedUpdate.PreviousRoot() != latestRoot {
				logger.Debugf("Local root not equal to chain root. Skipping update")
				continue
			}

			// Ensure we have not already signed a conflicting update.
			// Ignore suggested if we have.
			existing, err := u.db.RetrieveProducedUpdate(suggestedUpdate.PreviousRoot())
			if err != nil && !errors.Is(err, pebble.ErrNotFound) {
				return fmt.Errorf("could not get update: %w", err)
				// existing was found
			} else if err == nil {
				if existing.Update().NewRoot() != suggestedUpdate.NewRoot() {
					logger.Infof("Updater ignoring conflicting suggested update. Indicates chain awaiting already produced update. Existing update: %s. Suggested conflicting update: %s", existing.Update().NewRoot(), suggestedUpdate.NewRoot())
				}
				continue
			}

			// get the update to sign
			hashedUpdate, err := hashUpdate(suggestedUpdate)
			if err != nil {
				return fmt.Errorf("could not hash update: %w", err)
			}
			signature, err := u.signer.SignMessage(ctx, pebble2.ToSlice(hashedUpdate))
			if err != nil {
				return fmt.Errorf("could not sign message: %w", err)
			}

			signedUpdate := types.NewSignedUpdate(suggestedUpdate, signature)
			err = u.StoreProducedUpdate(signedUpdate)
			if err != nil {
				return err
			}
		}
	}
}

func hashUpdate(update types.Update) ([32]byte, error) {
	buf := new(bytes.Buffer)

	type DigestEncoder struct {
		HomeDomainHash, OldRoot, NewRoot [32]byte
	}

	homeHash, err := types.HomeDomainHash(update.HomeDomain())
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not get home domain hash: %w", err)
	}

	rawDigest := DigestEncoder{
		HomeDomainHash: homeHash,
		OldRoot:        update.PreviousRoot(),
		NewRoot:        update.NewRoot(),
	}

	err = binary.Write(buf, binary.BigEndian, rawDigest)
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not write digest: %w", err)
	}

	hashedDigest := crypto.Keccak256Hash(buf.Bytes())

	signedHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), hashedDigest.Bytes())
	return signedHash, nil
}
