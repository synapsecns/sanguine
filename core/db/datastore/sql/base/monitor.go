package base

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/core/types"
)

// StoreDispatchMessage takes a committed message and stores the necessary information.
func (s Store) StoreDispatchMessage(ctx context.Context, message types.CommittedMessage) error {
	decodedMessage, err := types.DecodeMessage(message.Message())
	if err != nil {
		return fmt.Errorf("could not decode message for insertion")
	}

	dxTx := s.DB().WithContext(ctx).Create(&DispatchMessage{
		DMOrigin:      decodedMessage.OriginDomain(),
		DMSender:      decodedMessage.Sender().String(),
		DMNonce:       decodedMessage.Nonce(),
		DMDestination: decodedMessage.DestinationDomain(),
		DMRecipient:   decodedMessage.Recipient().String(),
	})
	if dxTx.Error != nil {
		return fmt.Errorf("could not insert dispatch message: %w", dxTx.Error)
	}
	return nil
}

// StoreAcceptedAttestation stores an accepted attestation from a replica.
func (s Store) StoreAcceptedAttestation(ctx context.Context, replicaDomain uint32, attestation types.Attestation) error {
	dxTx := s.DB().WithContext(ctx).Create(&AcceptedAttestation{
		AAHomeDomain:    attestation.Domain(),
		AANonce:         attestation.Nonce(),
		AARoot:          attestation.Root(),
		AAReplicaDomain: replicaDomain,
	})
	if dxTx.Error != nil {
		return fmt.Errorf("could not insert accepted attestation: %w", dxTx.Error)
	}
	return nil
}
