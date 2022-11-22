package origin_scanner

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// OriginDestinationScanner scans for new messages from given origin to given destination.
type OriginDestinationScanner struct {
	signer      signer.Signer
	origin      uint32
	destination uint32
	// interval waits for an interval (in seconds)
	interval time.Duration
}

// NewOriginDestinationScanner creates a new origin-destination scanner.
func NewOriginDestinationScanner(ctx context.Context, signer signer.Signer, origin, destination uint32, interval time.Duration) (_ OriginDestinationScanner, err error) {
	originDestinationScanner := OriginDestinationScanner{
		signer:      signer,
		origin:      origin,
		destination: destination,
		interval:    interval,
	}

	return originDestinationScanner, nil
}

// Start starts the notary.{.
func (o OriginDestinationScanner) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(o.interval):
			err := o.checkMessagesFromOriginToDestination(ctx)
			if err != nil {
				return err
			}
		}
	}
}

// update runs the update producer to produce an update.
// nolint: cyclop
func (o OriginDestinationScanner) checkMessagesFromOriginToDestination(ctx context.Context) error {

	// 1) Most basic thread for a given origin-destination pair
	// a) Call origin and get latest nonce and root
	// b) Call AttestationCollector and get latest nonce for this signer for origin-destination
	// c) If Origin has a later nonce, then we want to sign
	// 	i) Sign for (origin, destination, latestNonce, latestRoot)
	//  ii) Submit attestion to attestion collector
	// d) We are done and can terminate thread

	fmt.Printf("CRONIN Origin<%d>-Destination<%d> Scanner \n", o.origin, o.destination)
	return nil

	// TODO (joe): we want to go through and update attestations for each destination.
	/*latestNonce, err := a.FindLatestNonce(ctx)
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
	signature, err := a.signer.SignMessage(ctx, core.BytesToSlice(hashedUpdate), false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	signedAttestation := types.NewSignedAttestation(suggestedAttestation, signature)
	err = a.db.StoreSignedAttestations(ctx, signedAttestation)
	if err != nil {
		return fmt.Errorf("could not store signed attestations: %w", err)
	}
	return nil*/
}
