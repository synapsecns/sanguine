package notary

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/synapse-node/contracts/bridge"
	"time"
)

// AttestationProducer updates a producer.
// TODO: this needs to become an interface.
type AttestationProducer struct {
	// domain allows access to the origin contract
	domain domains.DomainClient
	// db is the synapse db
	db db.SynapseDB

	// signer is the signer
	signer signer.Signer
	// interval waits for an interval
	interval time.Duration
}

// NewAttestationProducer creates an attestation producer.
func NewAttestationProducer(domain domains.DomainClient, db db.SynapseDB, signer signer.Signer, interval time.Duration) AttestationProducer {
	return AttestationProducer{
		domain:   domain,
		db:       db,
		signer:   signer,
		interval: interval,
	}
}

// Start starts the update producer.
func (a AttestationProducer) Start(ctx context.Context) error {
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
func (a AttestationProducer) FindLatestNonce(ctx context.Context) (nonce uint32, err error) {
	latestNonce, err := a.db.RetrieveLatestCommittedMessageNonce(ctx, a.domain.Config().DomainID)
	if err != nil {
		if errors.Is(err, db.ErrNoNonceForDomain) {
			return 0, nil
		}
		return 0, fmt.Errorf("could not find latest root: %w", err)
	}
	return latestNonce, nil
}

// update runs the update producer to produce an update.
//nolint: cyclop
func (a AttestationProducer) update(ctx context.Context) error {
	latestNonce, err := a.FindLatestNonce(ctx)
	if err != nil {
		return fmt.Errorf("could not find latest root: %w", err)
	}

	suggestedAttestation, err := a.domain.Origin().ProduceAttestation(ctx)
	if errors.Is(err, domains.ErrNoUpdate) {
		// no update produced this time
		return nil
	}
	if err != nil {
		return fmt.Errorf("could not suggest update: %w", err)
	}

	// TODO: let's figure out if we need to keep track of non-sequential updates?
	if suggestedAttestation.Nonce() < latestNonce {
		logger.Debugf("Local root not more then chain root. Skipping update")
		return nil
	}

	// Ensure we have not already signed a conflicting update.
	// Ignore suggested if we have.
	existing, err := a.db.RetrieveSignedAttestationByNonce(ctx, a.domain.Config().DomainID, suggestedAttestation.Nonce())
	if err != nil && !errors.Is(err, db.ErrNotFound) {
		return fmt.Errorf("could not get update: %w", err)
		// existing was found
	} else if err == nil {
		if existing.Attestation().Root() != suggestedAttestation.Root() {
			logger.Infof("Notary ignoring conflicting suggested update. Indicates chain awaiting already produced update. Existing update: %s. Suggested conflicting update: %s", existing.Attestation().Root(), suggestedAttestation.Root())
		}
		return nil
	}

	// get the update to sign
	hashedUpdate, err := HashAttestation(suggestedAttestation)
	if err != nil {
		return fmt.Errorf("could not hash update: %w", err)
	}
	signature, err := a.signer.SignMessage(ctx, bridge.KappaToSlice(hashedUpdate), false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	signedAttestation := types.NewSignedAttestation(suggestedAttestation, signature)
	err = a.db.StoreSignedAttestations(ctx, signedAttestation)
	if err != nil {
		return fmt.Errorf("could not store signed attestations: %w", err)
	}
	return nil
}

// HashAttestation hashes an attestation.
func HashAttestation(attestation types.Attestation) ([32]byte, error) {
	encodedAttestation, err := types.EncodeAttestation(attestation)
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not encode attestation: %w", err)
	}

	hashedDigest := crypto.Keccak256Hash(encodedAttestation)

	signedHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), hashedDigest.Bytes())
	return signedHash, nil
}
